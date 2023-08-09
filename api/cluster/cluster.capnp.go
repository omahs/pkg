// Code generated by capnpc-go. DO NOT EDIT.

package cluster

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	fc "capnproto.org/go/capnp/v3/flowcontrol"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
	anchor "github.com/wetware/ww/api/anchor"
	process "github.com/wetware/ww/api/process"
	pubsub "github.com/wetware/ww/api/pubsub"
	registry "github.com/wetware/ww/api/registry"
)

type Host capnp.Client

// Host_TypeID is the unique identifier for the type Host.
const Host_TypeID = 0x957cbefc645fd307

func (c Host) View(ctx context.Context, params func(Host_view_Params) error) (Host_view_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      0,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "view",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Host_view_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Host_view_Results_Future{Future: ans.Future()}, release

}

func (c Host) PubSub(ctx context.Context, params func(Host_pubSub_Params) error) (Host_pubSub_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      1,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "pubSub",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Host_pubSub_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Host_pubSub_Results_Future{Future: ans.Future()}, release

}

func (c Host) Root(ctx context.Context, params func(Host_root_Params) error) (Host_root_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      2,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "root",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Host_root_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Host_root_Results_Future{Future: ans.Future()}, release

}

func (c Host) Registry(ctx context.Context, params func(Host_registry_Params) error) (Host_registry_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      3,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "registry",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Host_registry_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Host_registry_Results_Future{Future: ans.Future()}, release

}

func (c Host) Executor(ctx context.Context, params func(Host_executor_Params) error) (Host_executor_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      4,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "executor",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Host_executor_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Host_executor_Results_Future{Future: ans.Future()}, release

}

func (c Host) WaitStreaming() error {
	return capnp.Client(c).WaitStreaming()
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c Host) String() string {
	return "Host(" + capnp.Client(c).String() + ")"
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c Host) AddRef() Host {
	return Host(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c Host) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c Host) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c Host) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (Host) DecodeFromPtr(p capnp.Ptr) Host {
	return Host(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c Host) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c Host) IsSame(other Host) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c Host) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c Host) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
}

// A Host_Server is a Host with a local implementation.
type Host_Server interface {
	View(context.Context, Host_view) error

	PubSub(context.Context, Host_pubSub) error

	Root(context.Context, Host_root) error

	Registry(context.Context, Host_registry) error

	Executor(context.Context, Host_executor) error
}

// Host_NewServer creates a new Server from an implementation of Host_Server.
func Host_NewServer(s Host_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Host_Methods(nil, s), s, c)
}

// Host_ServerToClient creates a new Client from an implementation of Host_Server.
// The caller is responsible for calling Release on the returned Client.
func Host_ServerToClient(s Host_Server) Host {
	return Host(capnp.NewClient(Host_NewServer(s)))
}

// Host_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Host_Methods(methods []server.Method, s Host_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 5)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      0,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "view",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.View(ctx, Host_view{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      1,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "pubSub",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.PubSub(ctx, Host_pubSub{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      2,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "root",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Root(ctx, Host_root{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      3,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "registry",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Registry(ctx, Host_registry{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0x957cbefc645fd307,
			MethodID:      4,
			InterfaceName: "cluster.capnp:Host",
			MethodName:    "executor",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Executor(ctx, Host_executor{call})
		},
	})

	return methods
}

// Host_view holds the state for a server call to Host.view.
// See server.Call for documentation.
type Host_view struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Host_view) Args() Host_view_Params {
	return Host_view_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Host_view) AllocResults() (Host_view_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_view_Results(r), err
}

// Host_pubSub holds the state for a server call to Host.pubSub.
// See server.Call for documentation.
type Host_pubSub struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Host_pubSub) Args() Host_pubSub_Params {
	return Host_pubSub_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Host_pubSub) AllocResults() (Host_pubSub_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_pubSub_Results(r), err
}

// Host_root holds the state for a server call to Host.root.
// See server.Call for documentation.
type Host_root struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Host_root) Args() Host_root_Params {
	return Host_root_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Host_root) AllocResults() (Host_root_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_root_Results(r), err
}

// Host_registry holds the state for a server call to Host.registry.
// See server.Call for documentation.
type Host_registry struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Host_registry) Args() Host_registry_Params {
	return Host_registry_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Host_registry) AllocResults() (Host_registry_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_registry_Results(r), err
}

// Host_executor holds the state for a server call to Host.executor.
// See server.Call for documentation.
type Host_executor struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Host_executor) Args() Host_executor_Params {
	return Host_executor_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Host_executor) AllocResults() (Host_executor_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_executor_Results(r), err
}

// Host_List is a list of Host.
type Host_List = capnp.CapList[Host]

// NewHost creates a new list of Host.
func NewHost_List(s *capnp.Segment, sz int32) (Host_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Host](l), err
}

type Host_view_Params capnp.Struct

// Host_view_Params_TypeID is the unique identifier for the type Host_view_Params.
const Host_view_Params_TypeID = 0xa404c24b5375b9e4

func NewHost_view_Params(s *capnp.Segment) (Host_view_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_view_Params(st), err
}

func NewRootHost_view_Params(s *capnp.Segment) (Host_view_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_view_Params(st), err
}

func ReadRootHost_view_Params(msg *capnp.Message) (Host_view_Params, error) {
	root, err := msg.Root()
	return Host_view_Params(root.Struct()), err
}

func (s Host_view_Params) String() string {
	str, _ := text.Marshal(0xa404c24b5375b9e4, capnp.Struct(s))
	return str
}

func (s Host_view_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_view_Params) DecodeFromPtr(p capnp.Ptr) Host_view_Params {
	return Host_view_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_view_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_view_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_view_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_view_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Host_view_Params_List is a list of Host_view_Params.
type Host_view_Params_List = capnp.StructList[Host_view_Params]

// NewHost_view_Params creates a new list of Host_view_Params.
func NewHost_view_Params_List(s *capnp.Segment, sz int32) (Host_view_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Host_view_Params](l), err
}

// Host_view_Params_Future is a wrapper for a Host_view_Params promised by a client call.
type Host_view_Params_Future struct{ *capnp.Future }

func (f Host_view_Params_Future) Struct() (Host_view_Params, error) {
	p, err := f.Future.Ptr()
	return Host_view_Params(p.Struct()), err
}

type Host_view_Results capnp.Struct

// Host_view_Results_TypeID is the unique identifier for the type Host_view_Results.
const Host_view_Results_TypeID = 0x8f58928e854cd4f5

func NewHost_view_Results(s *capnp.Segment) (Host_view_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_view_Results(st), err
}

func NewRootHost_view_Results(s *capnp.Segment) (Host_view_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_view_Results(st), err
}

func ReadRootHost_view_Results(msg *capnp.Message) (Host_view_Results, error) {
	root, err := msg.Root()
	return Host_view_Results(root.Struct()), err
}

func (s Host_view_Results) String() string {
	str, _ := text.Marshal(0x8f58928e854cd4f5, capnp.Struct(s))
	return str
}

func (s Host_view_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_view_Results) DecodeFromPtr(p capnp.Ptr) Host_view_Results {
	return Host_view_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_view_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_view_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_view_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_view_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Host_view_Results) View() capnp.Client {
	p, _ := capnp.Struct(s).Ptr(0)
	return p.Interface().Client()
}

func (s Host_view_Results) HasView() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Host_view_Results) SetView(c capnp.Client) error {
	if !c.IsValid() {
		return capnp.Struct(s).SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().CapTable().Add(c))
	return capnp.Struct(s).SetPtr(0, in.ToPtr())
}

// Host_view_Results_List is a list of Host_view_Results.
type Host_view_Results_List = capnp.StructList[Host_view_Results]

// NewHost_view_Results creates a new list of Host_view_Results.
func NewHost_view_Results_List(s *capnp.Segment, sz int32) (Host_view_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Host_view_Results](l), err
}

// Host_view_Results_Future is a wrapper for a Host_view_Results promised by a client call.
type Host_view_Results_Future struct{ *capnp.Future }

func (f Host_view_Results_Future) Struct() (Host_view_Results, error) {
	p, err := f.Future.Ptr()
	return Host_view_Results(p.Struct()), err
}
func (p Host_view_Results_Future) View() capnp.Client {
	return p.Future.Field(0, nil).Client()
}

type Host_pubSub_Params capnp.Struct

// Host_pubSub_Params_TypeID is the unique identifier for the type Host_pubSub_Params.
const Host_pubSub_Params_TypeID = 0xe5b5227505fcaa99

func NewHost_pubSub_Params(s *capnp.Segment) (Host_pubSub_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_pubSub_Params(st), err
}

func NewRootHost_pubSub_Params(s *capnp.Segment) (Host_pubSub_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_pubSub_Params(st), err
}

func ReadRootHost_pubSub_Params(msg *capnp.Message) (Host_pubSub_Params, error) {
	root, err := msg.Root()
	return Host_pubSub_Params(root.Struct()), err
}

func (s Host_pubSub_Params) String() string {
	str, _ := text.Marshal(0xe5b5227505fcaa99, capnp.Struct(s))
	return str
}

func (s Host_pubSub_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_pubSub_Params) DecodeFromPtr(p capnp.Ptr) Host_pubSub_Params {
	return Host_pubSub_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_pubSub_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_pubSub_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_pubSub_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_pubSub_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Host_pubSub_Params_List is a list of Host_pubSub_Params.
type Host_pubSub_Params_List = capnp.StructList[Host_pubSub_Params]

// NewHost_pubSub_Params creates a new list of Host_pubSub_Params.
func NewHost_pubSub_Params_List(s *capnp.Segment, sz int32) (Host_pubSub_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Host_pubSub_Params](l), err
}

// Host_pubSub_Params_Future is a wrapper for a Host_pubSub_Params promised by a client call.
type Host_pubSub_Params_Future struct{ *capnp.Future }

func (f Host_pubSub_Params_Future) Struct() (Host_pubSub_Params, error) {
	p, err := f.Future.Ptr()
	return Host_pubSub_Params(p.Struct()), err
}

type Host_pubSub_Results capnp.Struct

// Host_pubSub_Results_TypeID is the unique identifier for the type Host_pubSub_Results.
const Host_pubSub_Results_TypeID = 0xdc88f975f5090eee

func NewHost_pubSub_Results(s *capnp.Segment) (Host_pubSub_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_pubSub_Results(st), err
}

func NewRootHost_pubSub_Results(s *capnp.Segment) (Host_pubSub_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_pubSub_Results(st), err
}

func ReadRootHost_pubSub_Results(msg *capnp.Message) (Host_pubSub_Results, error) {
	root, err := msg.Root()
	return Host_pubSub_Results(root.Struct()), err
}

func (s Host_pubSub_Results) String() string {
	str, _ := text.Marshal(0xdc88f975f5090eee, capnp.Struct(s))
	return str
}

func (s Host_pubSub_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_pubSub_Results) DecodeFromPtr(p capnp.Ptr) Host_pubSub_Results {
	return Host_pubSub_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_pubSub_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_pubSub_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_pubSub_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_pubSub_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Host_pubSub_Results) PubSub() pubsub.Router {
	p, _ := capnp.Struct(s).Ptr(0)
	return pubsub.Router(p.Interface().Client())
}

func (s Host_pubSub_Results) HasPubSub() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Host_pubSub_Results) SetPubSub(v pubsub.Router) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().CapTable().Add(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(0, in.ToPtr())
}

// Host_pubSub_Results_List is a list of Host_pubSub_Results.
type Host_pubSub_Results_List = capnp.StructList[Host_pubSub_Results]

// NewHost_pubSub_Results creates a new list of Host_pubSub_Results.
func NewHost_pubSub_Results_List(s *capnp.Segment, sz int32) (Host_pubSub_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Host_pubSub_Results](l), err
}

// Host_pubSub_Results_Future is a wrapper for a Host_pubSub_Results promised by a client call.
type Host_pubSub_Results_Future struct{ *capnp.Future }

func (f Host_pubSub_Results_Future) Struct() (Host_pubSub_Results, error) {
	p, err := f.Future.Ptr()
	return Host_pubSub_Results(p.Struct()), err
}
func (p Host_pubSub_Results_Future) PubSub() pubsub.Router {
	return pubsub.Router(p.Future.Field(0, nil).Client())
}

type Host_root_Params capnp.Struct

// Host_root_Params_TypeID is the unique identifier for the type Host_root_Params.
const Host_root_Params_TypeID = 0x828b2823e5eeb7be

func NewHost_root_Params(s *capnp.Segment) (Host_root_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_root_Params(st), err
}

func NewRootHost_root_Params(s *capnp.Segment) (Host_root_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_root_Params(st), err
}

func ReadRootHost_root_Params(msg *capnp.Message) (Host_root_Params, error) {
	root, err := msg.Root()
	return Host_root_Params(root.Struct()), err
}

func (s Host_root_Params) String() string {
	str, _ := text.Marshal(0x828b2823e5eeb7be, capnp.Struct(s))
	return str
}

func (s Host_root_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_root_Params) DecodeFromPtr(p capnp.Ptr) Host_root_Params {
	return Host_root_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_root_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_root_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_root_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_root_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Host_root_Params_List is a list of Host_root_Params.
type Host_root_Params_List = capnp.StructList[Host_root_Params]

// NewHost_root_Params creates a new list of Host_root_Params.
func NewHost_root_Params_List(s *capnp.Segment, sz int32) (Host_root_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Host_root_Params](l), err
}

// Host_root_Params_Future is a wrapper for a Host_root_Params promised by a client call.
type Host_root_Params_Future struct{ *capnp.Future }

func (f Host_root_Params_Future) Struct() (Host_root_Params, error) {
	p, err := f.Future.Ptr()
	return Host_root_Params(p.Struct()), err
}

type Host_root_Results capnp.Struct

// Host_root_Results_TypeID is the unique identifier for the type Host_root_Results.
const Host_root_Results_TypeID = 0xcabb5c85a457450b

func NewHost_root_Results(s *capnp.Segment) (Host_root_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_root_Results(st), err
}

func NewRootHost_root_Results(s *capnp.Segment) (Host_root_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_root_Results(st), err
}

func ReadRootHost_root_Results(msg *capnp.Message) (Host_root_Results, error) {
	root, err := msg.Root()
	return Host_root_Results(root.Struct()), err
}

func (s Host_root_Results) String() string {
	str, _ := text.Marshal(0xcabb5c85a457450b, capnp.Struct(s))
	return str
}

func (s Host_root_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_root_Results) DecodeFromPtr(p capnp.Ptr) Host_root_Results {
	return Host_root_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_root_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_root_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_root_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_root_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Host_root_Results) Root() anchor.Anchor {
	p, _ := capnp.Struct(s).Ptr(0)
	return anchor.Anchor(p.Interface().Client())
}

func (s Host_root_Results) HasRoot() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Host_root_Results) SetRoot(v anchor.Anchor) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().CapTable().Add(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(0, in.ToPtr())
}

// Host_root_Results_List is a list of Host_root_Results.
type Host_root_Results_List = capnp.StructList[Host_root_Results]

// NewHost_root_Results creates a new list of Host_root_Results.
func NewHost_root_Results_List(s *capnp.Segment, sz int32) (Host_root_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Host_root_Results](l), err
}

// Host_root_Results_Future is a wrapper for a Host_root_Results promised by a client call.
type Host_root_Results_Future struct{ *capnp.Future }

func (f Host_root_Results_Future) Struct() (Host_root_Results, error) {
	p, err := f.Future.Ptr()
	return Host_root_Results(p.Struct()), err
}
func (p Host_root_Results_Future) Root() anchor.Anchor {
	return anchor.Anchor(p.Future.Field(0, nil).Client())
}

type Host_registry_Params capnp.Struct

// Host_registry_Params_TypeID is the unique identifier for the type Host_registry_Params.
const Host_registry_Params_TypeID = 0x89ec8e1ef0f263f3

func NewHost_registry_Params(s *capnp.Segment) (Host_registry_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_registry_Params(st), err
}

func NewRootHost_registry_Params(s *capnp.Segment) (Host_registry_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_registry_Params(st), err
}

func ReadRootHost_registry_Params(msg *capnp.Message) (Host_registry_Params, error) {
	root, err := msg.Root()
	return Host_registry_Params(root.Struct()), err
}

func (s Host_registry_Params) String() string {
	str, _ := text.Marshal(0x89ec8e1ef0f263f3, capnp.Struct(s))
	return str
}

func (s Host_registry_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_registry_Params) DecodeFromPtr(p capnp.Ptr) Host_registry_Params {
	return Host_registry_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_registry_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_registry_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_registry_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_registry_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Host_registry_Params_List is a list of Host_registry_Params.
type Host_registry_Params_List = capnp.StructList[Host_registry_Params]

// NewHost_registry_Params creates a new list of Host_registry_Params.
func NewHost_registry_Params_List(s *capnp.Segment, sz int32) (Host_registry_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Host_registry_Params](l), err
}

// Host_registry_Params_Future is a wrapper for a Host_registry_Params promised by a client call.
type Host_registry_Params_Future struct{ *capnp.Future }

func (f Host_registry_Params_Future) Struct() (Host_registry_Params, error) {
	p, err := f.Future.Ptr()
	return Host_registry_Params(p.Struct()), err
}

type Host_registry_Results capnp.Struct

// Host_registry_Results_TypeID is the unique identifier for the type Host_registry_Results.
const Host_registry_Results_TypeID = 0xbe186003ae0f0429

func NewHost_registry_Results(s *capnp.Segment) (Host_registry_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_registry_Results(st), err
}

func NewRootHost_registry_Results(s *capnp.Segment) (Host_registry_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_registry_Results(st), err
}

func ReadRootHost_registry_Results(msg *capnp.Message) (Host_registry_Results, error) {
	root, err := msg.Root()
	return Host_registry_Results(root.Struct()), err
}

func (s Host_registry_Results) String() string {
	str, _ := text.Marshal(0xbe186003ae0f0429, capnp.Struct(s))
	return str
}

func (s Host_registry_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_registry_Results) DecodeFromPtr(p capnp.Ptr) Host_registry_Results {
	return Host_registry_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_registry_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_registry_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_registry_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_registry_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Host_registry_Results) Registry() registry.Registry {
	p, _ := capnp.Struct(s).Ptr(0)
	return registry.Registry(p.Interface().Client())
}

func (s Host_registry_Results) HasRegistry() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Host_registry_Results) SetRegistry(v registry.Registry) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().CapTable().Add(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(0, in.ToPtr())
}

// Host_registry_Results_List is a list of Host_registry_Results.
type Host_registry_Results_List = capnp.StructList[Host_registry_Results]

// NewHost_registry_Results creates a new list of Host_registry_Results.
func NewHost_registry_Results_List(s *capnp.Segment, sz int32) (Host_registry_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Host_registry_Results](l), err
}

// Host_registry_Results_Future is a wrapper for a Host_registry_Results promised by a client call.
type Host_registry_Results_Future struct{ *capnp.Future }

func (f Host_registry_Results_Future) Struct() (Host_registry_Results, error) {
	p, err := f.Future.Ptr()
	return Host_registry_Results(p.Struct()), err
}
func (p Host_registry_Results_Future) Registry() registry.Registry {
	return registry.Registry(p.Future.Field(0, nil).Client())
}

type Host_executor_Params capnp.Struct

// Host_executor_Params_TypeID is the unique identifier for the type Host_executor_Params.
const Host_executor_Params_TypeID = 0xbe5314ed29d84c52

func NewHost_executor_Params(s *capnp.Segment) (Host_executor_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_executor_Params(st), err
}

func NewRootHost_executor_Params(s *capnp.Segment) (Host_executor_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Host_executor_Params(st), err
}

func ReadRootHost_executor_Params(msg *capnp.Message) (Host_executor_Params, error) {
	root, err := msg.Root()
	return Host_executor_Params(root.Struct()), err
}

func (s Host_executor_Params) String() string {
	str, _ := text.Marshal(0xbe5314ed29d84c52, capnp.Struct(s))
	return str
}

func (s Host_executor_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_executor_Params) DecodeFromPtr(p capnp.Ptr) Host_executor_Params {
	return Host_executor_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_executor_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_executor_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_executor_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_executor_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Host_executor_Params_List is a list of Host_executor_Params.
type Host_executor_Params_List = capnp.StructList[Host_executor_Params]

// NewHost_executor_Params creates a new list of Host_executor_Params.
func NewHost_executor_Params_List(s *capnp.Segment, sz int32) (Host_executor_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Host_executor_Params](l), err
}

// Host_executor_Params_Future is a wrapper for a Host_executor_Params promised by a client call.
type Host_executor_Params_Future struct{ *capnp.Future }

func (f Host_executor_Params_Future) Struct() (Host_executor_Params, error) {
	p, err := f.Future.Ptr()
	return Host_executor_Params(p.Struct()), err
}

type Host_executor_Results capnp.Struct

// Host_executor_Results_TypeID is the unique identifier for the type Host_executor_Results.
const Host_executor_Results_TypeID = 0x9e8120f9bb059602

func NewHost_executor_Results(s *capnp.Segment) (Host_executor_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_executor_Results(st), err
}

func NewRootHost_executor_Results(s *capnp.Segment) (Host_executor_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Host_executor_Results(st), err
}

func ReadRootHost_executor_Results(msg *capnp.Message) (Host_executor_Results, error) {
	root, err := msg.Root()
	return Host_executor_Results(root.Struct()), err
}

func (s Host_executor_Results) String() string {
	str, _ := text.Marshal(0x9e8120f9bb059602, capnp.Struct(s))
	return str
}

func (s Host_executor_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Host_executor_Results) DecodeFromPtr(p capnp.Ptr) Host_executor_Results {
	return Host_executor_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Host_executor_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Host_executor_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Host_executor_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Host_executor_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Host_executor_Results) Executor() process.Executor {
	p, _ := capnp.Struct(s).Ptr(0)
	return process.Executor(p.Interface().Client())
}

func (s Host_executor_Results) HasExecutor() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Host_executor_Results) SetExecutor(v process.Executor) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().CapTable().Add(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(0, in.ToPtr())
}

// Host_executor_Results_List is a list of Host_executor_Results.
type Host_executor_Results_List = capnp.StructList[Host_executor_Results]

// NewHost_executor_Results creates a new list of Host_executor_Results.
func NewHost_executor_Results_List(s *capnp.Segment, sz int32) (Host_executor_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Host_executor_Results](l), err
}

// Host_executor_Results_Future is a wrapper for a Host_executor_Results promised by a client call.
type Host_executor_Results_Future struct{ *capnp.Future }

func (f Host_executor_Results_Future) Struct() (Host_executor_Results, error) {
	p, err := f.Future.Ptr()
	return Host_executor_Results(p.Struct()), err
}
func (p Host_executor_Results_Future) Executor() process.Executor {
	return process.Executor(p.Future.Field(0, nil).Client())
}

const schema_fcf6ac08e448a6ac = "x\xda\x8cRMh\x13]\x14\xbdw\xde\x9bo\x9aL" +
	"\xda\xf22\xfd\xa4(\x88\xd4.l\xc1b\x11\x11\xb2i" +
	"\\\xa8\xb5v\x91i\xf0g!\xd64\x0eRh\x9d\x90" +
	"\x99\xa9\x06\xdcT\x08(\x85\x88\x82\x82.\\U\x17R" +
	"t!\"Z\x08\xe8N]\xf9\xb30 \x82\x90\x08\x82" +
	"\x0d\xa2T\x88\xa4<y\x93\xcedZj\xdb\xdd\x0c\xf7" +
	"\x9c\xf3\xee9\xe7\xee\x99!q\xda\xdf:\xba\x05\xa4$" +
	"E\xf9?^|Z\xad\xec\xdc5s\x19\x98\x86\x00T" +
	"\x01\xd0\xfeW\xff\x00\xe5\xbf\xd2?\x7fl/|\xbf\xda" +
	"\x1c\xec\xad\x85\xb7\"P\xbe\xf8a8_\xb8q\xf2Z" +
	"c\"\xa3\x18}\x09K\x08\xa8U\xc2\x03\x80\\y?" +
	"z\xb6^\xbct\x13X\x1b\xe1s\xf7\x07\xcb-s\xbf" +
	"\xeb\x00\xa8\xc9\xea\x1d\xadU\x15\xf8\x90z\x18\xb5\x92\xf8" +
	"\xe4\xd2-y\xbe\xb6c\xfanP\xee\x85\xda%\xe4^" +
	"\xa9B\xae\xfc\xccI\x1e}Ig\x03+.\xb8+\xf6" +
	"\xd0\xf6\x87\xe4Lg1\xc8|\xdb`\x96\\\xe6\xc8\xf0" +
	"\xc7\x9e\x85\x8ed1\xe0aIu=\xa8\x07O\xcc\xe6" +
	"O\xcd\xbf\x0eR+\xaa\xeb\xe1\x9bK\xad\xb6\x85\x16\x9d" +
	"\xda\x95OA@(\x12\x15\x00\x16\x11\x80\xdb\x0f\xea\xb2" +
	"\xd3\xf5\xa4\x12\xd0\xee\x8f\x84\x118p8\xce\xd3\x13\x8e" +
	"e\x1b\xd9>)\x9d\xca\x9c\xcf\xc4\x06M\xcb\xee\xcb\x9a" +
	"\xa6\xdd=\x90HeS\x93\x96\x0f A\x80qn\xdc" +
	"\xb2\xb3\xb9n\x81!\xff\x00M\x8d\x1b\x17\xbaG\x0c\xcb" +
	"\x99\xb0-\xd0)\xa1\x00\x14\x01Xk/\x80\xdeBP" +
	"\xef\x90\xb0]\x800J\x09 F\x01}\x1d\xf4t\x88" +
	"e'\x10\xf5\x0e\"\x03\xf8\xf9\xa2W,\xbb\xde\x0b\x12" +
	"\xcb+\xd8t\x89^\x1e,\x17\x03\x89M*(\xf9\xa7" +
	"\x83^\x98,%x\xc7\x14$\xfe\xf5\xa0\xd7\x11;2" +
	"\x04\x12;\xa0 \xf5[A\xafy\xb6O\xccv+\xee" +
	"\xdaq\x1c\xc88cIg,\x8e\xed\"\xb18r/" +
	"\x17\x00\x88#7.\x1ai\xc76\xb3\xee_\x02q\xcd" +
	"\x94<\x90\x9b\x942a[\xc1\xa4\x86\x00\xf4\x08A\xbd" +
	"SZ\xa1\x86\x8c\xbfs\xa6\xef=?\xdd\xf7\x08\x00\x91" +
	"\x05\x82\x93V\x17\xb0\xa9\x1a7z;\xe0\x0b\x19?\xf4" +
	"&\x976\x95\xea\xd2\xea\xb7\xd7\xb4\xb5\xde\x85\xb8w\xb6" +
	"\xd1\x85\x08\x102\xdeU*\x84\xca\xfb\xa3\xe5\xf5\x1em" +
	"\xd4\xb1\xac\x88+\xdc\xc4\x9a\x8a\xcb\xad!\xe31sj" +
	"\xdb\xd7\xc7\x89\xcf\x9b\xd0l\xa4\x08\xf07\x00\x00\xff\xff" +
	"89Z\xa7"

func RegisterSchema(reg *schemas.Registry) {
	reg.Register(&schemas.Schema{
		String: schema_fcf6ac08e448a6ac,
		Nodes: []uint64{
			0x828b2823e5eeb7be,
			0x89ec8e1ef0f263f3,
			0x8f58928e854cd4f5,
			0x957cbefc645fd307,
			0x9e8120f9bb059602,
			0xa404c24b5375b9e4,
			0xbe186003ae0f0429,
			0xbe5314ed29d84c52,
			0xcabb5c85a457450b,
			0xdc88f975f5090eee,
			0xe5b5227505fcaa99,
		},
		Compressed: true,
	})
}
