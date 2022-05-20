package cluster

import (
	"context"
	"time"

	"capnproto.org/go/capnp/v3"
	"capnproto.org/go/capnp/v3/server"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/wetware/casm/pkg/cluster/routing"
	chan_api "github.com/wetware/ww/internal/api/channel"
	api "github.com/wetware/ww/internal/api/cluster"
	"github.com/wetware/ww/pkg/cap/channel"
	"github.com/wetware/ww/pkg/vat"
	"golang.org/x/sync/semaphore"
)

var (
	ViewCapability = vat.BasicCap{
		"view/packed",
		"view"}
)

const (
	batchSize   = 64
	maxInFlight = 8
)

var defaultPolicy = server.Policy{
	MaxConcurrentCalls: 64,
}

// RoutingTable provides a global view of namespace peers.
type RoutingTable interface {
	Iter() routing.Iterator
	Lookup(peer.ID) (routing.Record, bool)
}

type ViewServer struct {
	View RoutingTable
}

func (f ViewServer) NewClient(policy *server.Policy) View {
	if policy == nil {
		policy = &defaultPolicy
	}

	return View(api.View_ServerToClient(f, policy))
}

func (f ViewServer) Client() *capnp.Client {
	return api.View_ServerToClient(f, &defaultPolicy).Client
}

func (f ViewServer) Iter(ctx context.Context, call api.View_iter) error {
	call.Ack()

	b := newBatcher(call.Args())

	for it := f.View.Iter(); it.Record() != nil; it.Next() {
		if err := b.Send(ctx, it.Record(), it.Deadline()); err != nil {
			it.Finish()
			return err
		}
	}

	return b.Wait(ctx)
}

func (f ViewServer) Lookup(_ context.Context, call api.View_lookup) error {
	peerID, err := call.Args().PeerID()
	if err != nil {
		return err
	}
	capRec, ok := f.View.Lookup(peer.ID(peerID))
	results, err := call.AllocResults()
	if err != nil {
		return err
	}
	rec, err := results.NewRecord()
	if err != nil {
		return err
	}

	results.SetOk(ok)

	if ok {
		rec.SetPeer(string(capRec.Peer()))
		rec.SetTtl(int64(capRec.TTL()))
		rec.SetSeq(capRec.Seq())
	}
	return nil
}

type View api.View

func (v View) Iter(ctx context.Context) (*RecordStream, capnp.ReleaseFunc) {
	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan Record, maxInFlight)

	rs, release := newRecordStream(ctx, api.View(v), ch)
	return rs, func() {
		cancel()
		release()
	}
}

func peerID(id peer.ID) func(api.View_lookup_Params) error {
	return func(ps api.View_lookup_Params) error {
		return ps.SetPeerID(string(id))
	}
}

func (v View) Lookup(ctx context.Context, id peer.ID) (FutureRecord, capnp.ReleaseFunc) {
	f, release := api.View(v).Lookup(ctx, peerID(id))
	return FutureRecord(f), release
}

type FutureRecord api.View_lookup_Results_Future

func (f FutureRecord) Record() (Record, error) {
	res, err := api.View_lookup_Results_Future(f).Struct()
	if err != nil {
		return Record{}, err
	}

	r, err := res.Record()
	if err != nil {
		return Record{}, err
	}

	return Record(r), Record(r).Validate()
}

func (f FutureRecord) Await(ctx context.Context) (Record, error) {
	select {
	case <-f.Done():
		return f.Record()

	case <-ctx.Done():
		return Record{}, ctx.Err()
	}
}

type Record api.View_Record

func (r Record) Validate() error {
	_, err := r.ID()
	return err
}

func (r Record) ID() (peer.ID, error) {
	s, err := api.View_Record(r).Peer()
	if err != nil {
		return "", err
	}

	return peer.IDFromString(s)
}

func (r Record) Peer() peer.ID {
	id, err := r.ID()
	if err != nil {
		panic(err)
	}

	return id
}

func (r Record) TTL() time.Duration {
	return time.Duration(api.View_Record(r).Ttl())
}

func (r Record) Seq() uint64 {
	return api.View_Record(r).Seq()
}

// RecordStream is a routing.Iterator that receives iterates
// over an asynchronous stream of records.
type RecordStream struct {
	h <-chan Record
	f channel.Future

	rec Record
	Err error
}

func newRecordStream(ctx context.Context, r api.View, ch chan Record) (*RecordStream, capnp.ReleaseFunc) {
	f, release := r.Iter(ctx, handler(ch))
	return &RecordStream{
		h: ch,
		f: channel.Future{Future: f.Future},
	}, release
}

func (rs *RecordStream) Next(ctx context.Context) (more bool) {
	for {
		select {
		case rs.rec, more = <-rs.h:
			return

		case <-rs.done():
			rs.Err = rs.f.Err()
			rs.f.Future = nil

		case <-ctx.Done():
			return false
		}
	}
}

func (rs *RecordStream) Record() routing.Record { return rs.rec }

func (rs *RecordStream) done() <-chan struct{} {
	if rs.f.Future != nil {
		return rs.f.Done()
	}

	return nil
}

type sender chan<- Record

func handler(s sender) func(ps api.View_iter_Params) error {
	return func(ps api.View_iter_Params) error {
		return ps.SetHandler(chan_api.Sender_ServerToClient(s, &server.Policy{
			MaxConcurrentCalls: cap(s),
		}))
	}
}

func (s sender) Shutdown() { close(s) }

func (s sender) Send(ctx context.Context, call chan_api.Sender_send) error {
	rs, err := handlerSendParams(call.Args()).Records()
	if err != nil {
		return err
	}

	var r Record
	for i := 0; i < rs.Len(); i++ {
		r = Record(rs.At(i))
		if err = r.Validate(); err != nil {
			return err
		}

		select {
		case s <- r:
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}

type handlerSendParams chan_api.Sender_send_Params

func (ps handlerSendParams) Records() (api.View_Record_List, error) {
	ptr, err := chan_api.Sender_send_Params(ps).Value()
	return capnp.StructList[api.View_Record]{List: ptr.List()}, err
}

func (ps handlerSendParams) NewRecords(size int32) (api.View_Record_List, error) {
	return api.NewView_Record_List(ps.Segment(), size)
}

type batcher struct {
	handler chan_api.Sender
	fs      map[*capnp.Future]capnp.ReleaseFunc // in-flight
	batch   batch
}

func newBatcher(p api.View_iter_Params) batcher {
	h := p.Handler()
	h.Client.SetFlowLimiter(newLimiter())

	return batcher{
		handler: h,
		fs:      make(map[*capnp.Future]capnp.ReleaseFunc),
		batch:   newBatch(),
	}
}

func (b *batcher) Send(ctx context.Context, r routing.Record, dl time.Time) error {
	// batch is full?
	if b.batch.Add(r, dl) {
		return b.Flush(ctx, false)
	}

	return nil
}

func (b *batcher) Flush(ctx context.Context, force bool) error {
	if b.batch.FilterExpired() || force {
		f, release := b.handler.Send(ctx, b.batch.Flush())
		b.fs[f.Future] = func() {
			delete(b.fs, f.Future)
			release()
		}
	}

	// release any resolved futures and return their errors, if any
	for f, release := range b.fs {
		select {
		case <-f.Done():
			defer release()
			if _, err := f.Struct(); err != nil {
				return err
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}

	return nil
}

func (b *batcher) Wait(ctx context.Context) (err error) {
	if err = b.Flush(ctx, true); err != nil {
		return
	}

	for f, release := range b.fs {
		// This is a rare case in which the use of 'defer' in
		// a loop is NOT a bug.
		//
		// We iterate over the whole map in order to schedule
		// a deferred call to 'release' for all pending RPC
		// calls.
		//
		// In principle, this is not necessary since resources
		// will be released when the handler for Iter returns.
		// We do it anyway to guard against bugs and/or changes
		// in the capnp API.
		defer release()

		// We want to abort early if any future encounters an
		// error, but as per the previous comment, we also want
		// to defer a call to 'release' for each future.
		if err == nil {
			// We're waiting until all futures resolve, so it's
			// okay to block on any given 'f'.
			select {
			case <-f.Done():
				_, err = f.Struct()

			case <-ctx.Done():
				err = ctx.Err()
			}
		}
	}

	return
}

type batch struct {
	t  time.Time
	rs []batchRecord
}

func newBatch() batch {
	return batch{
		rs: make([]batchRecord, 0, batchSize),
	}
}

func (b *batch) Full() bool { return len(b.rs) == cap(b.rs) }
func (b *batch) Len() int32 { return int32(len(b.rs)) }

func (b *batch) Add(r routing.Record, dl time.Time) bool {
	b.rs = append(b.rs, batchRecord{
		ID:       r.Peer(),
		Seq:      r.Seq(),
		Deadline: dl.Truncate(time.Millisecond),
	})

	return b.Full()
}

func (b *batch) Flush() func(chan_api.Sender_send_Params) error {
	return func(p chan_api.Sender_send_Params) error {
		defer func() {
			b.rs = b.rs[:0]
		}()

		rs, err := handlerSendParams(p).NewRecords(b.Len())
		if err != nil {
			return err
		}

		for i, r := range b.rs {
			if err = r.SetParam(b.t, rs.At(i)); err != nil {
				break
			}
		}

		return err
	}
}

func (b *batch) FilterExpired() bool {
	b.t = time.Now().Truncate(time.Millisecond)

	current := b.rs[:]
	b.rs = b.rs[:0]
	for _, r := range current {
		if b.t.Before(r.Deadline) {
			b.rs = append(b.rs, r)
		}
	}

	return b.Full()
}

type batchRecord struct {
	ID       peer.ID
	Seq      uint64
	Deadline time.Time
}

func (r batchRecord) SetParam(t time.Time, rec api.View_Record) error {
	rec.SetSeq(r.Seq)
	rec.SetTtl(r.Deadline.Sub(t).Microseconds())
	return rec.SetPeer(string(r.ID))
}

type limiter semaphore.Weighted

func newLimiter() *limiter {
	return (*limiter)(semaphore.NewWeighted(maxInFlight))
}

func (l *limiter) StartMessage(ctx context.Context, _ uint64) (gotResponse func(), err error) {
	if err = l.Acquire(ctx); err == nil {
		gotResponse = l.Release
	}

	return
}

func (l *limiter) Acquire(ctx context.Context) error {
	return (*semaphore.Weighted)(l).Acquire(ctx, 1)
}

func (l *limiter) Release() {
	(*semaphore.Weighted)(l).Release(1)
}
