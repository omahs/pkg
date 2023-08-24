package ww

import (
	"context"
	"crypto/rand"
	_ "embed"
	"io"
	"runtime"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/rpc"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"

	"github.com/wetware/pkg/rom"
	"github.com/wetware/pkg/system"
)

const (
	Version = "0.1.0"
)

// Ww is the execution context for WebAssembly (WASM) bytecode,
// allowing it to interact with (1) the local host and (2) the
// cluster environment.
type Ww[T ~capnp.ClientKind] struct {
	NS     string
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	Client T
}

// String returns the cluster namespace in which the wetware is
// executing. If ww.NS has been assigned a non-empty string, it
// returns the string unchanged.  Else, it defaults to "ww".
func (ww *Ww[T]) String() string {
	if ww.NS != "" {
		return ww.NS
	}

	return "ww"
}

// Exec compiles and runs the ww instance's ROM in a WASM runtime.
// It returns any error produced by the compilation or execution of
// the ROM.
func (ww Ww[T]) Exec(ctx context.Context, rom rom.ROM) error {
	// Spawn a new runtime.
	r := wazero.NewRuntimeWithConfig(ctx, wazero.
		NewRuntimeConfigCompiler().
		WithCloseOnContextDone(true))
	defer r.Close(ctx)

	// Instantiate WASI.
	c, err := wasi_snapshot_preview1.Instantiate(ctx, r)
	if err != nil {
		return err
	}
	defer c.Close(ctx)

	// Instantiate Wetware host module.
	sock, err := system.Instantiate(ctx, r)
	if err != nil {
		return err
	}
	defer sock.Close()

	// Compile guest module.
	compiled, err := r.CompileModule(sock.Ctx(), rom.Bytecode)
	if err != nil {
		return err
	}
	defer compiled.Close(ctx)

	// Instantiate the guest module, and configure host exports.
	mod, err := r.InstantiateModule(sock.Ctx(), compiled, wazero.NewModuleConfig().
		WithOsyield(runtime.Gosched). // runtime.Gosched
		WithRandSource(rand.Reader).
		WithStartFunctions(). // don't automatically call _start while instanitating.
		WithSysNanosleep().
		WithSysNanotime().
		WithSysWalltime().
		WithArgs(rom.String()). // TODO(soon):  use content id
		WithEnv("ns", ww.String()).
		WithName(rom.String()).
		WithStdin(ww.Stdin). // notice:  we connect stdio to host process' stdio
		WithStdout(ww.Stdout).
		WithStderr(ww.Stderr))
	if err != nil {
		return err
	}
	defer mod.Close(ctx)

	conn := rpc.NewConn(sock.RPCTransport(), &rpc.Options{
		BootstrapClient: capnp.Client(ww.Client),
	})
	defer conn.Close()

	return sock.Bind(mod)
}
