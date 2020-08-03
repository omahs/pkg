package client

import (
	"context"
	"time"

	"go.uber.org/fx"

	"github.com/ipfs/go-datastore"
	"github.com/libp2p/go-eventbus"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/event"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/pnet"
	"github.com/libp2p/go-libp2p-core/routing"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/config"
	"github.com/pkg/errors"

	ctxutil "github.com/wetware/ww/internal/util/ctx"
	hostutil "github.com/wetware/ww/internal/util/host"
	"github.com/wetware/ww/pkg/boot"
	"github.com/wetware/ww/pkg/internal/p2p"
	"github.com/wetware/ww/pkg/runtime"
	"github.com/wetware/ww/pkg/runtime/service"
)

const (
	kmin     = 3
	kmax     = 64
	timestep = time.Millisecond * 100
)

func services(cfg serviceConfig) runtime.ServiceBundle {
	return runtime.Bundle(
		service.Ticker(cfg.Host.EventBus(), timestep),
		service.ConnTracker(cfg.Host),
		service.Neighborhood(cfg.Host.EventBus(), kmin, kmax),
		service.Bootstrap(cfg.Host, cfg.Boot),
		// service.Discover(cfg.EventBus, cfg.Namespace, cfg.Discovery),  // TODO:  initial advertisement
		service.Graph(cfg.Host),
		service.Joiner(cfg.Host),
	)
}

// Config contains user-supplied parameters used by Dial.
type Config struct {
	ns  string
	psk pnet.PSK
	ds  datastore.Batching
	d   boot.Strategy
}

func (cfg Config) assemble(ctx context.Context, c *Client) {
	c.app = fx.New(
		fx.NopLogger,
		fx.Populate(c),
		fx.Provide(
			cfg.options,
			p2p.New,
			services,
			newClient,
		),
		fx.Invoke(
			runtime.Register,
			dial,
		),
	)
}

func (cfg Config) options(lx fx.Lifecycle) (mod module, err error) {
	mod.Ctx = ctxutil.WithLifecycle(context.Background(), lx) // libp2p lifecycle
	mod.Namespace = cfg.ns
	mod.Datastore = cfg.ds
	mod.Boot = cfg.d

	// options for host.Host
	mod.HostOpt = []config.Option{
		hostutil.MaybePrivate(cfg.psk),
		libp2p.Ping(false),
		libp2p.NoListenAddrs, // also disables relay
		libp2p.UserAgent("ww-client"),
	}

	// options for DHT
	mod.DHTOpt = []dht.Option{
		dht.Datastore(cfg.ds),
		dht.Mode(dht.ModeClient),
	}

	return
}

type module struct {
	fx.Out

	Ctx       context.Context
	Namespace string `name:"ns"`

	Datastore datastore.Batching
	Boot      boot.Strategy

	HostOpt []config.Option
	DHTOpt  []dht.Option
}

type serviceConfig struct {
	fx.In

	Namespace string `name:"ns"`
	Host      host.Host
	Discovery discovery.Discovery
	Boot      boot.Strategy
}

func dial(h host.Host, dht routing.Routing, lx fx.Lifecycle) error {
	e, err := h.EventBus().Emitter(new(p2p.EvtNetworkReady), eventbus.Stateful)
	if err != nil {
		return err
	}
	defer e.Close()

	lx.Append(dialhook(h.EventBus(), dht))

	return e.Emit(netready(h))
}

func dialhook(bus event.Bus, dht interface{ Bootstrap(context.Context) error }) fx.Hook {
	return fx.Hook{
		OnStart: func(ctx context.Context) error {
			sub, err := bus.Subscribe(new(service.EvtNeighborhoodChanged))
			if err != nil {
				return err
			}
			defer sub.Close()

			for {
				select {
				case v := <-sub.Out():
					if v.(service.EvtNeighborhoodChanged).To == service.PhaseOrphaned {
						continue
					}

					return errors.Wrap(dht.Bootstrap(ctx), "dht bootstrap") // async call
				case <-ctx.Done():
					return errors.Wrap(ctx.Err(), "join cluster")
				}
			}
		},
	}
}

func netready(h host.Host) p2p.EvtNetworkReady {
	return p2p.EvtNetworkReady{Network: h.Network()}
}
