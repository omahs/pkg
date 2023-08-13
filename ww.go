package ww

import (
	"context"
	"crypto/rand"
	_ "embed"
	"errors"
	"io"
	"runtime"

	"capnproto.org/go/capnp/v3"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/sys"
	"github.com/wetware/pkg/system"
	"golang.org/x/exp/slog"
)

const (
	Version = "0.1.0"
	Codec   = 2020
)

// Ww is the execution context for WebAssembly (WASM) bytecode,
// allowing it to interact with (1) the local host and (2) the
// cluster environment.
type Ww[T ~capnp.ClientKind] struct {
	NS              string
	Stdin           io.Reader
	Stdout          io.Writer
	Stderr          io.Writer
	BootstrapClient T
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
func (ww Ww[T]) Exec(ctx context.Context, rom ROM) error {
	// Spawn a new runtime.
	r := wazero.NewRuntimeWithConfig(ctx, wazero.
		NewRuntimeConfigCompiler().
		WithCloseOnContextDone(true))
	defer r.Close(ctx)

	// Initialize  system modules.
	sys, ctx, err := system.Init[T](ctx, r, ww.BootstrapClient)
	if err != nil {
		return err
	}
	defer sys.Close(ctx)

	// Build the guest module.
	// Start by compiling the ROM bytecode.
	compiled, err := r.CompileModule(ctx, rom.bytecode)
	if err != nil {
		return err
	}
	defer compiled.Close(ctx)

	// Instantiate the guest module, and configure host exports.
	mod, err := r.InstantiateModule(ctx, compiled, wazero.NewModuleConfig().
		WithOsyield(runtime.Gosched).
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

	return ww.run(ctx, mod)
}

func (ww Ww[T]) run(ctx context.Context, mod api.Module) error {
	// Grab the the main() function and call it with the system context.
	fn := mod.ExportedFunction("_start")
	if fn == nil {
		return errors.New("missing export: _start")
	}

	// TODO(performance):  fn.CallWithStack(ctx, nil)
	_, err := fn.Call(ctx)
	switch err.(*sys.ExitError).ExitCode() {
	case 0:
	case sys.ExitCodeContextCanceled:
		return context.Canceled
	case sys.ExitCodeDeadlineExceeded:
		return context.DeadlineExceeded
	default:
		slog.Default().Debug(err.Error(),
			"version", Version,
			"ns", ww.String(),
			"rom", mod.Name())
	}

	return nil
}
