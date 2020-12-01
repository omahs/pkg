package lang

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/spy16/slurp"
	"github.com/wetware/ww/internal/api"
	ww "github.com/wetware/ww/pkg"
	"github.com/wetware/ww/pkg/lang/core"
	capnp "zombiezen.com/go/capnproto2"
)

func parseDo(a core.Analyzer, env core.Env, args core.Seq) (core.Expr, error) {
	var de DoExpr
	err := core.ForEach(args, func(item ww.Any) (bool, error) {
		expr, err := a.Analyze(env, item)
		if err != nil {
			return true, err
		}
		de.Exprs = append(de.Exprs, expr)
		return false, nil
	})
	if err != nil {
		return nil, err
	}
	return de, nil
}

func parseIf(a core.Analyzer, env core.Env, args core.Seq) (core.Expr, error) {
	count, err := args.Count()
	if err != nil {
		return nil, err
	} else if count != 2 && count != 3 {
		return nil, core.Error{
			Cause:   fmt.Errorf("%w: if", slurp.ErrParseSpecial),
			Message: fmt.Sprintf("requires 2 or 3 arguments, got %d", count),
		}
	}

	exprs := [3]core.Expr{}
	for i := 0; i < count; i++ {
		f, err := args.First()
		if err != nil {
			return nil, err
		}

		expr, err := a.Analyze(env, f)
		if err != nil {
			return nil, err
		}
		exprs[i] = expr

		args, err = args.Next()
		if err != nil {
			return nil, err
		}
	}

	return IfExpr{
		Test: exprs[0],
		Then: exprs[1],
		Else: exprs[2],
	}, nil
}
func parseQuote(a core.Analyzer, _ core.Env, args core.Seq) (core.Expr, error) {
	if count, err := args.Count(); err != nil {
		return nil, err
	} else if count != 1 {
		return nil, core.Error{
			Cause:   fmt.Errorf("%w: quote", slurp.ErrParseSpecial),
			Message: fmt.Sprintf("requires exactly 1 argument, got %d", count),
		}
	}

	first, err := args.First()
	if err != nil {
		return nil, err
	}

	return QuoteExpr{Form: first}, nil
}

func parseDef(a core.Analyzer, env core.Env, args core.Seq) (core.Expr, error) {
	e := core.Error{Cause: fmt.Errorf("%w: def", slurp.ErrParseSpecial)}

	if args == nil {
		return nil, e.With("requires exactly 2 args, got 0")
	}

	if count, err := args.Count(); err != nil {
		return nil, err
	} else if count != 2 {
		return nil, e.With(fmt.Sprintf(
			"requires exactly 2 arguments, got %d", count))
	}

	first, err := args.First()
	if err != nil {
		return nil, err
	}

	sym, ok := first.(core.Symbol)
	if !ok {
		return nil, e.With(fmt.Sprintf(
			"first arg must be symbol, not '%s'", reflect.TypeOf(first)))
	}

	symStr, err := sym.Symbol()
	if err != nil {
		return nil, err
	}

	rest, err := args.Next()
	if err != nil {
		return nil, err
	}

	second, err := rest.First()
	if err != nil {
		return nil, err
	}

	res, err := a.Analyze(env, second)
	if err != nil {
		return nil, err
	}

	return DefExpr{
		Name:  symStr,
		Value: res,
	}, nil
}

// parseFn parses the (fn name? [<params>*] <body>*) or
// (fn name? ([<params>*] <body>*)+) special forms and returns a function value.
func parseFn(a core.Analyzer, env core.Env, args core.Seq) (core.Expr, error) {
	fn, err := parseFnDef(env, args, false)
	return ConstExpr{fn}, err
}

// parseFn parses the (macro name? [<params>*] <body>*) or
// (macro name? ([<params>*] <body>*)+) special forms and returns a macro value.
func parseMacro(_ core.Analyzer, env core.Env, args core.Seq) (core.Expr, error) {
	fn, err := parseFnDef(env, args, true)
	return ConstExpr{fn}, err
}

func parseFnDef(env core.Env, args core.Seq, macro bool) (core.Fn, error) {
	if args == nil {
		return core.Fn{}, errors.New("nil argument sequence")
	}

	cnt, err := args.Count()
	if err != nil {
		return core.Fn{}, err
	} else if cnt < 1 {
		return core.Fn{}, fmt.Errorf("%w: got %d, want at-least 1", core.ErrArity, cnt)
	}

	var b core.FuncBuilder
	b.Start(capnp.SingleSegment(nil))
	b.SetMacro(macro)

	first, err := args.First()
	if err != nil {
		return core.Fn{}, err
	}

	// Set function name?
	if sym := first.MemVal(); sym.Type() == api.Value_Which_symbol {
		name, err := sym.Raw.Symbol()
		if err != nil {
			return core.Fn{}, err
		}

		b.SetName(name)

		// Advance the sequence ...
		if args, err = args.Next(); err != nil {
			return core.Fn{}, err
		}

		if first, err = args.First(); err != nil {
			return core.Fn{}, err
		}
	}

	// Set call signatures.
	switch mv := first.MemVal(); mv.Type() {
	case api.Value_Which_vector:
		b.AddTarget(args)

	case api.Value_Which_list:
		if err = core.ForEach(args, func(item ww.Any) (bool, error) {
			if t, ok := item.(core.Seq); ok {
				b.AddTarget(t)
				return false, nil
			}

			return false, fmt.Errorf("expected core.Seq, got %s", item.MemVal().Type())
		}); err != nil {
			return core.Fn{}, err
		}

	default:
		return core.Fn{}, errors.New("syntax error")

	}

	return b.Commit()
}

func lsParser(root ww.Anchor) SpecialParser {
	return func(a core.Analyzer, _ core.Env, seq core.Seq) (core.Expr, error) {
		args, err := core.ToSlice(seq)
		if err != nil {
			return nil, err
		}

		pexpr := PathExpr{Root: root, Path: core.RootPath}
		for _, arg := range args {
			if arg.MemVal().Type() == api.Value_Which_path {
				pexpr.Path = args[0].(core.Path)
				args = args[1:]
			}

			break
		}

		// TODO(enhancement):  other args like `:long` or `:recursive`

		return PathListExpr{
			PathExpr: pexpr,
			Args:     args,
		}, nil
	}
}

func parseGo(a core.Analyzer, env core.Env, seq core.Seq) (core.Expr, error) {
	return nil, errors.New("parseGo NOT IMPLEMENTED")
	// args, err := core.ToSlice(seq)
	// if err != nil {
	// 	return nil, err
	// }

	// if len(args) == 0 {
	// 	return nil, errors.Errorf("expected at least one argument, got %d", len(args))
	// }

	// if p, ok := procArgs(args).Remote(); ok {
	// 	return RemoteGoExpr{
	// 		Root: a.root,
	// 		Path: p,
	// 		Args: procArgs(args).Args(),
	// 	}, nil
	// }

	// return LocalGoExpr{
	// 	Args: procArgs(args).Args(),
	// }, nil
}
