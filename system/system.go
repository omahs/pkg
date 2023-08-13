package system

import (
	"context"
	"fmt"
	"net"
	"time"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/jpillora/backoff"
	"github.com/stealthrocket/wazergo"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/experimental/sock"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/wetware/pkg/cap/view"
	"github.com/wetware/pkg/log"
	"go.uber.org/multierr"
	"golang.org/x/exp/slog"
)

type Closer struct {
	api.Closer
	Next *Closer
}

func (c Closer) Close(ctx context.Context) (err error) {
	for c, next := c.Closer, c.Next; c != nil; c, next = next.Closer, next.Next {
		err = multierr.Append(err, c.Close(ctx))
	}

	return err
}

type HostModule[T ~capnp.ClientKind] interface {
	Instantiate(ctx context.Context, r wazero.Runtime, t T) (api.Closer, context.Context, error)
}

func Init[T ~capnp.ClientKind](ctx context.Context, r wazero.Runtime, t T) (c Closer, out context.Context, err error) {
	for name, module := range map[string]HostModule[T]{
		"wasi": wasi[T]{},
		"ww":   wetware[T]{},
		"view": view.HostModule[T]{},
	} {
		if c.Closer, ctx, err = module.Instantiate(ctx, r, t); err != nil {
			err = Error{Module: name, Cause: err}
		}

		slog.Debug("instantiated", "module", name)
		c = Closer{
			Closer: nil, // available
			Next: &Closer{
				Closer: c.Closer,
				Next:   c.Next,
			},
		}
	}

	out = ctx
	return
}

type wasi[T ~capnp.ClientKind] struct{}

func (wasi[T]) Instantiate(ctx context.Context, r wazero.Runtime, t T) (api.Closer, context.Context, error) {
	c, err := wasi_snapshot_preview1.Instantiate(ctx, r)
	return c, ctx, err
}

type wetware[T ~capnp.ClientKind] struct{}

func (wetware[T]) Instantiate(ctx context.Context, r wazero.Runtime, t T) (api.Closer, context.Context, error) {
	return Instantiate[T](ctx, r, t)
}

// module for wetware Host
var module wazergo.HostModule[*Host] = functions{
	// "answer": F0((*Module).Answer),
	// "double": F1((*Module).Double),
}

// Instantiate the system host module.  If instantiation fails, the
// returned context is expired, and the ctx.Err() method returns the
// offending error.
func Instantiate[T ~capnp.ClientKind](ctx context.Context, r wazero.Runtime, t T) (*wazergo.ModuleInstance[*Host], context.Context, error) {
	l, err := net.Listen("tcp", ":0") // TODO:  localhost?
	if err != nil {
		return nil, ctx, fmt.Errorf("net: listen: %w", err)
	}
	defer l.Close()

	addr := l.Addr().(*net.TCPAddr)

	// Instantiate the host module and bind it to the context.
	instance, err := wazergo.Instantiate(ctx, r, module,
		logger(slog.Default()),
		transport(addr),
		bootstrap(t))
	if err == nil {
		// Bind the module instance to the context, so that the caller can
		// access it.
		ctx = wazergo.WithModuleInstance(ctx, instance)

		// The system socket enables the creation of non-blocking TCP conns
		// inside of the WASM module.  The host will pre-open the TCP port
		// and pass it to the guest through a file descriptor.
		ctx = sock.WithConfig(ctx, sock.NewConfig().
			WithTCPListener("", addr.Port))
	}

	return instance, ctx, err

}

type Option = wazergo.Option[*Host]

func logger(log log.Logger) Option {
	return wazergo.OptionFunc(func(h *Host) {
		h.Logger = log
	})
}

func transport(addr net.Addr) Option {
	return wazergo.OptionFunc(func(h *Host) {
		h.Addr = addr
	})
}

func bootstrap[T ~capnp.ClientKind](t T) Option {
	return wazergo.OptionFunc(func(h *Host) {
		h.BootstrapClient = capnp.Client(t)
	})
}

// The `functions` type impements `Module[*Module]`, providing the
// module name, map of exported functions, and the ability to create
// instances of the module type
type functions wazergo.Functions[*Host]

func (f functions) Name() string {
	return "ww"
}

func (f functions) Functions() wazergo.Functions[*Host] {
	return (wazergo.Functions[*Host])(f)
}

func (f functions) Instantiate(ctx context.Context, opts ...Option) (out *Host, err error) {
	wazergo.Configure(new(Host), append(opts, wazergo.OptionFunc(func(h *Host) {
		var b = backoff.Backoff{
			Min:    time.Millisecond * 1,
			Max:    time.Minute,
			Jitter: true,
		}

		// retry in a loop until context is canceled; back-off exponentially.
		for err = h.dial(ctx); err != nil; err = h.dial(ctx) {
			h.Logger.Debug("failed to dial host socket",
				"error", err,
				"attempt", b.Attempt(),
				"backoff", b.ForAttempt(b.Attempt()))

			select {
			case <-time.After(b.Duration()):
				h.Logger.Debug("resuming",
					"attempt", b.Attempt())

			case <-ctx.Done():
				err = ctx.Err()
			}
		}

		out = h // pass up the call stack
	}))...)

	return
}

type Host struct {
	Addr            net.Addr
	Logger          log.Logger
	BootstrapClient capnp.Client

	conn *rpc.Conn
}

func (h *Host) Close(context.Context) error {
	h.BootstrapClient.Release()

	return h.conn.Close()
}

func (h *Host) dial(ctx context.Context) error {
	raw, err := dial(ctx, h.Addr)
	if err != nil {
		return err
	}

	h.conn = rpc.NewConn(rpc.NewStreamTransport(raw), &rpc.Options{
		ErrorReporter:   ErrorReporter{Logger: h.Logger},
		BootstrapClient: h.BootstrapClient,
	})

	return nil
}

func dial(ctx context.Context, addr net.Addr) (net.Conn, error) {
	dialer := net.Dialer{}
	return dialer.DialContext(ctx, addr.Network(), addr.String())
}
