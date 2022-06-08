package channel

import (
	"context"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/server"
	"github.com/wetware/ww/internal/api/channel"
	"github.com/wetware/ww/pkg/ocap"
)

type Value func(channel.Sender_send_Params) error

func Ptr(ptr capnp.Ptr) Value {
	return func(ps channel.Sender_send_Params) error {
		return ps.SetValue(ptr)
	}
}

func Data(b []byte) Value {
	return func(ps channel.Sender_send_Params) error {
		return ps.SetData(0, b)
	}
}

func Text(s string) Value {
	return func(ps channel.Sender_send_Params) error {
		return ps.SetText(0, s)
	}
}

type Chan channel.Chan

func New(s Server, p *server.Policy) Chan {
	return Chan(channel.Chan_ServerToClient(s, p))
}

func (c Chan) Close(ctx context.Context) error {
	return Closer(c).Close(ctx)
}

func (c Chan) Send(ctx context.Context, v Value) (ocap.Future, capnp.ReleaseFunc) {
	return Sender(c).Send(ctx, v)
}

func (c Chan) Recv(ctx context.Context) (ocap.FuturePtr, capnp.ReleaseFunc) {
	return Recver(c).Recv(ctx)
}

func (c Chan) AddRef() Chan {
	return Chan{Client: c.Client.AddRef()}
}

func (c Chan) Release() {
	c.Client.Release()
}

type PeekableChan channel.Chan

func (c PeekableChan) Close(ctx context.Context) error {
	return Closer(c).Close(ctx)
}

func NewPeekableChan(s PeekableServer, p *server.Policy) PeekableChan {
	return PeekableChan(channel.PeekableChan_ServerToClient(s, p))
}

func (c PeekableChan) Send(ctx context.Context, v Value) (ocap.Future, capnp.ReleaseFunc) {
	return Sender(c).Send(ctx, v)
}

func (c PeekableChan) Recv(ctx context.Context) (ocap.FuturePtr, capnp.ReleaseFunc) {
	return Recver(c).Recv(ctx)
}

func (c PeekableChan) AddRef() PeekableChan {
	return PeekableChan{Client: c.Client.AddRef()}
}

func (c PeekableChan) Release() {
	c.Client.Release()
}

type SendCloser Chan

func NewSendCloser(sc SendCloseServer, p *server.Policy) SendCloser {
	return SendCloser(channel.SendCloser_ServerToClient(sc, p))
}

func (sc SendCloser) Close(ctx context.Context) error {
	return Closer(sc).Close(ctx)
}

func (sc SendCloser) Send(ctx context.Context, v Value) (ocap.Future, capnp.ReleaseFunc) {
	return Sender(sc).Send(ctx, v)
}

func (sc SendCloser) AddRef() SendCloser {
	return SendCloser{Client: sc.Client.AddRef()}
}

func (sc SendCloser) Release() {
	sc.Client.Release()
}

type PeekRecver Chan

func NewPeekRecver(pr PeekRecvServer, p *server.Policy) PeekRecver {
	return PeekRecver(channel.PeekRecver_ServerToClient(pr, p))
}

func (pr PeekRecver) Peek(ctx context.Context) (ocap.FuturePtr, capnp.ReleaseFunc) {
	return Peeker(pr).Peek(ctx)
}

func (pr PeekRecver) Recv(ctx context.Context) (ocap.FuturePtr, capnp.ReleaseFunc) {
	return Recver(pr).Recv(ctx)
}

func (pr PeekRecver) AddRef() PeekRecver {
	return PeekRecver{Client: pr.Client.AddRef()}
}

func (pr PeekRecver) Release() {
	pr.Client.Release()
}

type Sender Chan

func NewSender(s SendServer, p *server.Policy) Sender {
	return Sender(channel.Sender_ServerToClient(s, p))
}

func (s Sender) Send(ctx context.Context, v Value) (ocap.Future, capnp.ReleaseFunc) {
	f, release := channel.Sender(s).Send(ctx, v)
	return ocap.Future(f), release
}

func (s Sender) AddRef() Sender {
	return Sender{Client: s.Client.AddRef()}
}

func (s Sender) Release() {
	s.Client.Release()
}

type Peeker Chan

func NewPeeker(p PeekServer, q *server.Policy) Peeker {
	return Peeker(channel.Peeker_ServerToClient(p, q))
}

func (p Peeker) Peek(ctx context.Context) (ocap.FuturePtr, capnp.ReleaseFunc) {
	f, release := channel.Peeker(p).Peek(ctx, nil)
	return ocap.FuturePtr(f), release
}

func (p Peeker) AddRef() Peeker {
	return Peeker{Client: p.Client.AddRef()}
}

func (p Peeker) Release() {
	p.Client.Release()
}

type Recver Chan

func NewRecver(r RecvServer, p *server.Policy) Recver {
	return Recver(channel.Recver_ServerToClient(r, p))
}

func (r Recver) Recv(ctx context.Context) (ocap.FuturePtr, capnp.ReleaseFunc) {
	f, release := channel.Recver(r).Recv(ctx, nil)
	return ocap.FuturePtr(f), release
}

func (r Recver) AddRef() Recver {
	return Recver{Client: r.Client.AddRef()}
}

func (r Recver) Release() {
	r.Client.Release()
}

type Closer Chan

func NewCloser(c CloseServer, p *server.Policy) Closer {
	return Closer(channel.Closer_ServerToClient(c, p))
}

func (c Closer) Close(ctx context.Context) error {
	f, release := channel.Closer(c).Close(ctx, nil)
	defer release()

	_, err := f.Struct()
	return err
}

func (c Closer) AddRef() Closer {
	return Closer{Client: c.Client.AddRef()}
}

func (c Closer) Release() {
	c.Client.Release()
}
