package system

import (
	"context"
	"errors"

	rpccp "capnproto.org/go/capnp/v3/std/capnp/rpc"
	"zenhack.net/go/util/rc"
)

var ErrNotReady = errors.New("not ready")

type Pipe struct {
	closed chan struct{}
	buffer chan *rc.Ref[rpccp.Message]
}

func NewPipe() *Pipe {
	return &Pipe{
		closed: make(chan struct{}),
		buffer: make(chan *rc.Ref[rpccp.Message], 1),
	}
}

func (p *Pipe) Push(ctx context.Context, ref *rc.Ref[rpccp.Message]) error {
	select {
	case <-p.closed:
		return errors.New("closed")
	default:
	}

	select {
	case p.buffer <- ref:
		// fast path; we have a message waiting in the buffer
		return nil

	case <-p.closed:
		return errors.New("closed")

	case <-ctx.Done():
		return ctx.Err()

	default:
		return ErrNotReady
	}
}

func (p *Pipe) Pop(ctx context.Context) (*rc.Ref[rpccp.Message], error) {
	select {
	case ref := <-p.buffer:
		return ref, nil

	case <-p.closed:
		return nil, errors.New("closed")

	case <-ctx.Done():
		return nil, ctx.Err()

	default:
		return nil, ErrNotReady
	}
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (p *Pipe) Close() error {
	select {
	case <-p.closed:
		return errors.New("closed")
	default:
		close(p.closed)
		return nil
	}
}
