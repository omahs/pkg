package start

import (
	"context"
	"io"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/metrics"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	libp2pquic "github.com/libp2p/go-libp2p-quic-transport"
	"github.com/lthibault/log"
	"github.com/thejerf/suture/v4"
	"github.com/urfave/cli/v2"

	"go.uber.org/fx"

	"github.com/wetware/casm/pkg/cluster"
	"github.com/wetware/ww/internal/runtime"
	logutil "github.com/wetware/ww/internal/util/log"
	serviceutil "github.com/wetware/ww/internal/util/service"
	statsdutil "github.com/wetware/ww/internal/util/statsd"
	"github.com/wetware/ww/pkg/server"
)

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "ns",
		Usage:   "cluster namespace",
		Value:   "ww",
		EnvVars: []string{"WW_NS"},
	},
	&cli.StringSliceFlag{
		Name:    "listen",
		Aliases: []string{"a"},
		Usage:   "host listen address",
		Value: cli.NewStringSlice(
			"/ip4/0.0.0.0/udp/2020/quic",
			"/ip6/::0/udp/2020/quic"),
		EnvVars: []string{"WW_LISTEN"},
	},
	&cli.StringFlag{
		Name:    "discover",
		Aliases: []string{"d"},
		Usage:   "bootstrap discovery addr (cidr url)",
		Value:   "tcp://127.0.0.1:8822/24", // TODO:  this should default to mudp
		EnvVars: []string{"WW_DISCOVER"},
	},
}

// Command constructor
func Command() *cli.Command {
	return &cli.Command{
		Name:   "start",
		Usage:  "start a host process",
		Flags:  flags,
		Action: run(),
	}
}

func run() cli.ActionFunc {
	return func(c *cli.Context) error {
		var (
			logger log.Logger
			node   *server.Node
			app    = fx.New(fx.NopLogger,
				fx.Populate(&logger, &node),
				bind(c))
		)

		if err := start(c, app); err != nil {
			return err
		}

		logger.With(node).Info("wetware loaded")
		<-app.Done() // TODO:  process OS signals in a loop here.

		return shutdown(app)
	}
}

func bind(c *cli.Context) fx.Option {
	return fx.Options(
		runtime.Bind(),
		fx.Supply(c),
		fx.Provide(
			statsdutil.NewBandwidthCounter,
			statsdutil.New,
			logutil.New,
			supervisor,
			localhost,
			node))
}

//
// Dependency declarations
//

func supervisor(c *cli.Context) *suture.Supervisor {
	return suture.New("runtime", suture.Spec{
		EventHook: serviceutil.NewEventHook(c, "runtime"),
	})
}

func localhost(c *cli.Context, lx fx.Lifecycle, b *metrics.BandwidthCounter) (host.Host, error) {
	h, err := libp2p.New(c.Context,
		libp2p.NoTransports,
		libp2p.Transport(libp2pquic.NewTransport),
		libp2p.ListenAddrStrings(c.StringSlice("listen")...),
		libp2p.BandwidthReporter(b))
	if err == nil {
		lx.Append(closer(h))
	}

	return h, err
}

type serverConfig struct {
	fx.In

	Log       log.Logger
	Host      host.Host
	PubSub    *pubsub.PubSub
	Lifecycle fx.Lifecycle
}

func node(c *cli.Context, config serverConfig) (*server.Node, error) {
	n, err := server.New(c.Context, config.Host, config.PubSub,
		server.WithLogger(config.Log),
		server.WithNamespace(c.String("ns")),
		server.WithClusterConfig(
			cluster.WithMeta(nil) /* TODO */))

	if err == nil {
		config.Lifecycle.Append(closer(n))
	}

	return n, err
}

func start(c *cli.Context, app *fx.App) error {
	ctx, cancel := context.WithTimeout(c.Context, time.Second*15)
	defer cancel()

	return app.Start(ctx)
}

func shutdown(app *fx.App) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	if err = app.Stop(ctx); err == context.Canceled {
		err = nil
	}

	return
}

func closer(c io.Closer) fx.Hook {
	return fx.Hook{
		OnStop: func(context.Context) error {
			return c.Close()
		},
	}
}
