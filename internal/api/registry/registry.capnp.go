// Code generated by capnpc-go. DO NOT EDIT.

package registry

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	fc "capnproto.org/go/capnp/v3/flowcontrol"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
	channel "github.com/wetware/ww/internal/api/channel"
	pubsub "github.com/wetware/ww/internal/api/pubsub"
	strconv "strconv"
)

type Registry capnp.Client

// Registry_TypeID is the unique identifier for the type Registry.
const Registry_TypeID = 0xfdee076f6379cb46

func (c Registry) Provide(ctx context.Context, params func(Registry_provide_Params) error) (Registry_provide_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xfdee076f6379cb46,
			MethodID:      0,
			InterfaceName: "registry.capnp:Registry",
			MethodName:    "provide",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Registry_provide_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Registry_provide_Results_Future{Future: ans.Future()}, release

}

func (c Registry) FindProviders(ctx context.Context, params func(Registry_findProviders_Params) error) (Registry_findProviders_Results_Future, capnp.ReleaseFunc) {

	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xfdee076f6379cb46,
			MethodID:      1,
			InterfaceName: "registry.capnp:Registry",
			MethodName:    "findProviders",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 2}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Registry_findProviders_Params(s)) }
	}

	ans, release := capnp.Client(c).SendCall(ctx, s)
	return Registry_findProviders_Results_Future{Future: ans.Future()}, release

}

func (c Registry) WaitStreaming() error {
	return capnp.Client(c).WaitStreaming()
}

// String returns a string that identifies this capability for debugging
// purposes.  Its format should not be depended on: in particular, it
// should not be used to compare clients.  Use IsSame to compare clients
// for equality.
func (c Registry) String() string {
	return "Registry(" + capnp.Client(c).String() + ")"
}

// AddRef creates a new Client that refers to the same capability as c.
// If c is nil or has resolved to null, then AddRef returns nil.
func (c Registry) AddRef() Registry {
	return Registry(capnp.Client(c).AddRef())
}

// Release releases a capability reference.  If this is the last
// reference to the capability, then the underlying resources associated
// with the capability will be released.
//
// Release will panic if c has already been released, but not if c is
// nil or resolved to null.
func (c Registry) Release() {
	capnp.Client(c).Release()
}

// Resolve blocks until the capability is fully resolved or the Context
// expires.
func (c Registry) Resolve(ctx context.Context) error {
	return capnp.Client(c).Resolve(ctx)
}

func (c Registry) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Client(c).EncodeAsPtr(seg)
}

func (Registry) DecodeFromPtr(p capnp.Ptr) Registry {
	return Registry(capnp.Client{}.DecodeFromPtr(p))
}

// IsValid reports whether c is a valid reference to a capability.
// A reference is invalid if it is nil, has resolved to null, or has
// been released.
func (c Registry) IsValid() bool {
	return capnp.Client(c).IsValid()
}

// IsSame reports whether c and other refer to a capability created by the
// same call to NewClient.  This can return false negatives if c or other
// are not fully resolved: use Resolve if this is an issue.  If either
// c or other are released, then IsSame panics.
func (c Registry) IsSame(other Registry) bool {
	return capnp.Client(c).IsSame(capnp.Client(other))
}

// Update the flowcontrol.FlowLimiter used to manage flow control for
// this client. This affects all future calls, but not calls already
// waiting to send. Passing nil sets the value to flowcontrol.NopLimiter,
// which is also the default.
func (c Registry) SetFlowLimiter(lim fc.FlowLimiter) {
	capnp.Client(c).SetFlowLimiter(lim)
}

// Get the current flowcontrol.FlowLimiter used to manage flow control
// for this client.
func (c Registry) GetFlowLimiter() fc.FlowLimiter {
	return capnp.Client(c).GetFlowLimiter()
}

// A Registry_Server is a Registry with a local implementation.
type Registry_Server interface {
	Provide(context.Context, Registry_provide) error

	FindProviders(context.Context, Registry_findProviders) error
}

// Registry_NewServer creates a new Server from an implementation of Registry_Server.
func Registry_NewServer(s Registry_Server) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Registry_Methods(nil, s), s, c)
}

// Registry_ServerToClient creates a new Client from an implementation of Registry_Server.
// The caller is responsible for calling Release on the returned Client.
func Registry_ServerToClient(s Registry_Server) Registry {
	return Registry(capnp.NewClient(Registry_NewServer(s)))
}

// Registry_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Registry_Methods(methods []server.Method, s Registry_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xfdee076f6379cb46,
			MethodID:      0,
			InterfaceName: "registry.capnp:Registry",
			MethodName:    "provide",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Provide(ctx, Registry_provide{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xfdee076f6379cb46,
			MethodID:      1,
			InterfaceName: "registry.capnp:Registry",
			MethodName:    "findProviders",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.FindProviders(ctx, Registry_findProviders{call})
		},
	})

	return methods
}

// Registry_provide holds the state for a server call to Registry.provide.
// See server.Call for documentation.
type Registry_provide struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Registry_provide) Args() Registry_provide_Params {
	return Registry_provide_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Registry_provide) AllocResults() (Registry_provide_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Registry_provide_Results(r), err
}

// Registry_findProviders holds the state for a server call to Registry.findProviders.
// See server.Call for documentation.
type Registry_findProviders struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Registry_findProviders) Args() Registry_findProviders_Params {
	return Registry_findProviders_Params(c.Call.Args())
}

// AllocResults allocates the results struct.
func (c Registry_findProviders) AllocResults() (Registry_findProviders_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Registry_findProviders_Results(r), err
}

// Registry_List is a list of Registry.
type Registry_List = capnp.CapList[Registry]

// NewRegistry creates a new list of Registry.
func NewRegistry_List(s *capnp.Segment, sz int32) (Registry_List, error) {
	l, err := capnp.NewPointerList(s, sz)
	return capnp.CapList[Registry](l), err
}

type Registry_provide_Params capnp.Struct

// Registry_provide_Params_TypeID is the unique identifier for the type Registry_provide_Params.
const Registry_provide_Params_TypeID = 0xbf9edfd4684337f6

func NewRegistry_provide_Params(s *capnp.Segment) (Registry_provide_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Registry_provide_Params(st), err
}

func NewRootRegistry_provide_Params(s *capnp.Segment) (Registry_provide_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Registry_provide_Params(st), err
}

func ReadRootRegistry_provide_Params(msg *capnp.Message) (Registry_provide_Params, error) {
	root, err := msg.Root()
	return Registry_provide_Params(root.Struct()), err
}

func (s Registry_provide_Params) String() string {
	str, _ := text.Marshal(0xbf9edfd4684337f6, capnp.Struct(s))
	return str
}

func (s Registry_provide_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Registry_provide_Params) DecodeFromPtr(p capnp.Ptr) Registry_provide_Params {
	return Registry_provide_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Registry_provide_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Registry_provide_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Registry_provide_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Registry_provide_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Registry_provide_Params) Topic() pubsub.Topic {
	p, _ := capnp.Struct(s).Ptr(0)
	return pubsub.Topic(p.Interface().Client())
}

func (s Registry_provide_Params) HasTopic() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Registry_provide_Params) SetTopic(v pubsub.Topic) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().CapTable().Add(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(0, in.ToPtr())
}

func (s Registry_provide_Params) Envelope() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return []byte(p.Data()), err
}

func (s Registry_provide_Params) HasEnvelope() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Registry_provide_Params) SetEnvelope(v []byte) error {
	return capnp.Struct(s).SetData(1, v)
}

// Registry_provide_Params_List is a list of Registry_provide_Params.
type Registry_provide_Params_List = capnp.StructList[Registry_provide_Params]

// NewRegistry_provide_Params creates a new list of Registry_provide_Params.
func NewRegistry_provide_Params_List(s *capnp.Segment, sz int32) (Registry_provide_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return capnp.StructList[Registry_provide_Params](l), err
}

// Registry_provide_Params_Future is a wrapper for a Registry_provide_Params promised by a client call.
type Registry_provide_Params_Future struct{ *capnp.Future }

func (f Registry_provide_Params_Future) Struct() (Registry_provide_Params, error) {
	p, err := f.Future.Ptr()
	return Registry_provide_Params(p.Struct()), err
}
func (p Registry_provide_Params_Future) Topic() pubsub.Topic {
	return pubsub.Topic(p.Future.Field(0, nil).Client())
}

type Registry_provide_Results capnp.Struct

// Registry_provide_Results_TypeID is the unique identifier for the type Registry_provide_Results.
const Registry_provide_Results_TypeID = 0xd9ef66060e1157d3

func NewRegistry_provide_Results(s *capnp.Segment) (Registry_provide_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Registry_provide_Results(st), err
}

func NewRootRegistry_provide_Results(s *capnp.Segment) (Registry_provide_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Registry_provide_Results(st), err
}

func ReadRootRegistry_provide_Results(msg *capnp.Message) (Registry_provide_Results, error) {
	root, err := msg.Root()
	return Registry_provide_Results(root.Struct()), err
}

func (s Registry_provide_Results) String() string {
	str, _ := text.Marshal(0xd9ef66060e1157d3, capnp.Struct(s))
	return str
}

func (s Registry_provide_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Registry_provide_Results) DecodeFromPtr(p capnp.Ptr) Registry_provide_Results {
	return Registry_provide_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Registry_provide_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Registry_provide_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Registry_provide_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Registry_provide_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Registry_provide_Results_List is a list of Registry_provide_Results.
type Registry_provide_Results_List = capnp.StructList[Registry_provide_Results]

// NewRegistry_provide_Results creates a new list of Registry_provide_Results.
func NewRegistry_provide_Results_List(s *capnp.Segment, sz int32) (Registry_provide_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Registry_provide_Results](l), err
}

// Registry_provide_Results_Future is a wrapper for a Registry_provide_Results promised by a client call.
type Registry_provide_Results_Future struct{ *capnp.Future }

func (f Registry_provide_Results_Future) Struct() (Registry_provide_Results, error) {
	p, err := f.Future.Ptr()
	return Registry_provide_Results(p.Struct()), err
}

type Registry_findProviders_Params capnp.Struct

// Registry_findProviders_Params_TypeID is the unique identifier for the type Registry_findProviders_Params.
const Registry_findProviders_Params_TypeID = 0xd589c56f3d6a445e

func NewRegistry_findProviders_Params(s *capnp.Segment) (Registry_findProviders_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Registry_findProviders_Params(st), err
}

func NewRootRegistry_findProviders_Params(s *capnp.Segment) (Registry_findProviders_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Registry_findProviders_Params(st), err
}

func ReadRootRegistry_findProviders_Params(msg *capnp.Message) (Registry_findProviders_Params, error) {
	root, err := msg.Root()
	return Registry_findProviders_Params(root.Struct()), err
}

func (s Registry_findProviders_Params) String() string {
	str, _ := text.Marshal(0xd589c56f3d6a445e, capnp.Struct(s))
	return str
}

func (s Registry_findProviders_Params) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Registry_findProviders_Params) DecodeFromPtr(p capnp.Ptr) Registry_findProviders_Params {
	return Registry_findProviders_Params(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Registry_findProviders_Params) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Registry_findProviders_Params) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Registry_findProviders_Params) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Registry_findProviders_Params) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Registry_findProviders_Params) Topic() pubsub.Topic {
	p, _ := capnp.Struct(s).Ptr(0)
	return pubsub.Topic(p.Interface().Client())
}

func (s Registry_findProviders_Params) HasTopic() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Registry_findProviders_Params) SetTopic(v pubsub.Topic) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().CapTable().Add(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(0, in.ToPtr())
}

func (s Registry_findProviders_Params) Chan() channel.Sender {
	p, _ := capnp.Struct(s).Ptr(1)
	return channel.Sender(p.Interface().Client())
}

func (s Registry_findProviders_Params) HasChan() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Registry_findProviders_Params) SetChan(v channel.Sender) error {
	if !v.IsValid() {
		return capnp.Struct(s).SetPtr(1, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().CapTable().Add(capnp.Client(v)))
	return capnp.Struct(s).SetPtr(1, in.ToPtr())
}

// Registry_findProviders_Params_List is a list of Registry_findProviders_Params.
type Registry_findProviders_Params_List = capnp.StructList[Registry_findProviders_Params]

// NewRegistry_findProviders_Params creates a new list of Registry_findProviders_Params.
func NewRegistry_findProviders_Params_List(s *capnp.Segment, sz int32) (Registry_findProviders_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return capnp.StructList[Registry_findProviders_Params](l), err
}

// Registry_findProviders_Params_Future is a wrapper for a Registry_findProviders_Params promised by a client call.
type Registry_findProviders_Params_Future struct{ *capnp.Future }

func (f Registry_findProviders_Params_Future) Struct() (Registry_findProviders_Params, error) {
	p, err := f.Future.Ptr()
	return Registry_findProviders_Params(p.Struct()), err
}
func (p Registry_findProviders_Params_Future) Topic() pubsub.Topic {
	return pubsub.Topic(p.Future.Field(0, nil).Client())
}

func (p Registry_findProviders_Params_Future) Chan() channel.Sender {
	return channel.Sender(p.Future.Field(1, nil).Client())
}

type Registry_findProviders_Results capnp.Struct

// Registry_findProviders_Results_TypeID is the unique identifier for the type Registry_findProviders_Results.
const Registry_findProviders_Results_TypeID = 0xd86baa632daef690

func NewRegistry_findProviders_Results(s *capnp.Segment) (Registry_findProviders_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Registry_findProviders_Results(st), err
}

func NewRootRegistry_findProviders_Results(s *capnp.Segment) (Registry_findProviders_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Registry_findProviders_Results(st), err
}

func ReadRootRegistry_findProviders_Results(msg *capnp.Message) (Registry_findProviders_Results, error) {
	root, err := msg.Root()
	return Registry_findProviders_Results(root.Struct()), err
}

func (s Registry_findProviders_Results) String() string {
	str, _ := text.Marshal(0xd86baa632daef690, capnp.Struct(s))
	return str
}

func (s Registry_findProviders_Results) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Registry_findProviders_Results) DecodeFromPtr(p capnp.Ptr) Registry_findProviders_Results {
	return Registry_findProviders_Results(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Registry_findProviders_Results) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Registry_findProviders_Results) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Registry_findProviders_Results) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Registry_findProviders_Results) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}

// Registry_findProviders_Results_List is a list of Registry_findProviders_Results.
type Registry_findProviders_Results_List = capnp.StructList[Registry_findProviders_Results]

// NewRegistry_findProviders_Results creates a new list of Registry_findProviders_Results.
func NewRegistry_findProviders_Results_List(s *capnp.Segment, sz int32) (Registry_findProviders_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Registry_findProviders_Results](l), err
}

// Registry_findProviders_Results_Future is a wrapper for a Registry_findProviders_Results promised by a client call.
type Registry_findProviders_Results_Future struct{ *capnp.Future }

func (f Registry_findProviders_Results_Future) Struct() (Registry_findProviders_Results, error) {
	p, err := f.Future.Ptr()
	return Registry_findProviders_Results(p.Struct()), err
}

type Message capnp.Struct
type Message_Which uint16

const (
	Message_Which_request  Message_Which = 0
	Message_Which_response Message_Which = 1
)

func (w Message_Which) String() string {
	const s = "requestresponse"
	switch w {
	case Message_Which_request:
		return s[0:7]
	case Message_Which_response:
		return s[7:15]

	}
	return "Message_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Message_TypeID is the unique identifier for the type Message.
const Message_TypeID = 0xd2afeaf36c70c91f

func NewMessage(s *capnp.Segment) (Message, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Message(st), err
}

func NewRootMessage(s *capnp.Segment) (Message, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Message(st), err
}

func ReadRootMessage(msg *capnp.Message) (Message, error) {
	root, err := msg.Root()
	return Message(root.Struct()), err
}

func (s Message) String() string {
	str, _ := text.Marshal(0xd2afeaf36c70c91f, capnp.Struct(s))
	return str
}

func (s Message) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Message) DecodeFromPtr(p capnp.Ptr) Message {
	return Message(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Message) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}

func (s Message) Which() Message_Which {
	return Message_Which(capnp.Struct(s).Uint16(0))
}
func (s Message) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Message) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Message) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Message) SetRequest() {
	capnp.Struct(s).SetUint16(0, 0)

}

func (s Message) Response() ([]byte, error) {
	if capnp.Struct(s).Uint16(0) != 1 {
		panic("Which() != response")
	}
	p, err := capnp.Struct(s).Ptr(0)
	return []byte(p.Data()), err
}

func (s Message) HasResponse() bool {
	if capnp.Struct(s).Uint16(0) != 1 {
		return false
	}
	return capnp.Struct(s).HasPtr(0)
}

func (s Message) SetResponse(v []byte) error {
	capnp.Struct(s).SetUint16(0, 1)
	return capnp.Struct(s).SetData(0, v)
}

// Message_List is a list of Message.
type Message_List = capnp.StructList[Message]

// NewMessage creates a new list of Message.
func NewMessage_List(s *capnp.Segment, sz int32) (Message_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return capnp.StructList[Message](l), err
}

// Message_Future is a wrapper for a Message promised by a client call.
type Message_Future struct{ *capnp.Future }

func (f Message_Future) Struct() (Message, error) {
	p, err := f.Future.Ptr()
	return Message(p.Struct()), err
}

type Location capnp.Struct
type Location_Which uint16

const (
	Location_Which_maddrs Location_Which = 0
	Location_Which_anchor Location_Which = 1
	Location_Which_custom Location_Which = 2
)

func (w Location_Which) String() string {
	const s = "maddrsanchorcustom"
	switch w {
	case Location_Which_maddrs:
		return s[0:6]
	case Location_Which_anchor:
		return s[6:12]
	case Location_Which_custom:
		return s[12:18]

	}
	return "Location_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Location_TypeID is the unique identifier for the type Location.
const Location_TypeID = 0xe61540af32cf81b6

func NewLocation(s *capnp.Segment) (Location, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Location(st), err
}

func NewRootLocation(s *capnp.Segment) (Location, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return Location(st), err
}

func ReadRootLocation(msg *capnp.Message) (Location, error) {
	root, err := msg.Root()
	return Location(root.Struct()), err
}

func (s Location) String() string {
	str, _ := text.Marshal(0xe61540af32cf81b6, capnp.Struct(s))
	return str
}

func (s Location) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Location) DecodeFromPtr(p capnp.Ptr) Location {
	return Location(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Location) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}

func (s Location) Which() Location_Which {
	return Location_Which(capnp.Struct(s).Uint16(0))
}
func (s Location) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Location) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Location) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Location) Service() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Location) HasService() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Location) ServiceBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Location) SetService(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Location) Meta() (capnp.TextList, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return capnp.TextList(p.List()), err
}

func (s Location) HasMeta() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Location) SetMeta(v capnp.TextList) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewMeta sets the meta field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Location) NewMeta(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(capnp.Struct(s).Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}
func (s Location) Maddrs() (capnp.DataList, error) {
	if capnp.Struct(s).Uint16(0) != 0 {
		panic("Which() != maddrs")
	}
	p, err := capnp.Struct(s).Ptr(2)
	return capnp.DataList(p.List()), err
}

func (s Location) HasMaddrs() bool {
	if capnp.Struct(s).Uint16(0) != 0 {
		return false
	}
	return capnp.Struct(s).HasPtr(2)
}

func (s Location) SetMaddrs(v capnp.DataList) error {
	capnp.Struct(s).SetUint16(0, 0)
	return capnp.Struct(s).SetPtr(2, v.ToPtr())
}

// NewMaddrs sets the maddrs field to a newly
// allocated capnp.DataList, preferring placement in s's segment.
func (s Location) NewMaddrs(n int32) (capnp.DataList, error) {
	capnp.Struct(s).SetUint16(0, 0)
	l, err := capnp.NewDataList(capnp.Struct(s).Segment(), n)
	if err != nil {
		return capnp.DataList{}, err
	}
	err = capnp.Struct(s).SetPtr(2, l.ToPtr())
	return l, err
}
func (s Location) Anchor() (string, error) {
	if capnp.Struct(s).Uint16(0) != 1 {
		panic("Which() != anchor")
	}
	p, err := capnp.Struct(s).Ptr(2)
	return p.Text(), err
}

func (s Location) HasAnchor() bool {
	if capnp.Struct(s).Uint16(0) != 1 {
		return false
	}
	return capnp.Struct(s).HasPtr(2)
}

func (s Location) AnchorBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return p.TextBytes(), err
}

func (s Location) SetAnchor(v string) error {
	capnp.Struct(s).SetUint16(0, 1)
	return capnp.Struct(s).SetText(2, v)
}

func (s Location) Custom() (capnp.Ptr, error) {
	if capnp.Struct(s).Uint16(0) != 2 {
		panic("Which() != custom")
	}
	return capnp.Struct(s).Ptr(2)
}

func (s Location) HasCustom() bool {
	if capnp.Struct(s).Uint16(0) != 2 {
		return false
	}
	return capnp.Struct(s).HasPtr(2)
}

func (s Location) SetCustom(v capnp.Ptr) error {
	capnp.Struct(s).SetUint16(0, 2)
	return capnp.Struct(s).SetPtr(2, v)
}

// Location_List is a list of Location.
type Location_List = capnp.StructList[Location]

// NewLocation creates a new list of Location.
func NewLocation_List(s *capnp.Segment, sz int32) (Location_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return capnp.StructList[Location](l), err
}

// Location_Future is a wrapper for a Location promised by a client call.
type Location_Future struct{ *capnp.Future }

func (f Location_Future) Struct() (Location, error) {
	p, err := f.Future.Ptr()
	return Location(p.Struct()), err
}
func (p Location_Future) Custom() *capnp.Future {
	return p.Future.Field(2, nil)
}

const schema_fcba4f486a351ac3 = "x\xda\x8cSKh$U\x14=\xe7\xbd\xaa.\x07\xd2" +
	"i\xdfT\xfb\x85\xa6\x11\x12\x98\x88\x19g2\x88C@" +
	"\xba\x89\x1fD\x14\xfbe\xe3N(\xaak&\x9dIW" +
	"\x95U\x9d\xe8\x80C\x18P\xc4]\\\x08Y\xb9\x124" +
	"\xa2&\x0b\x11\x8c\x0b1\xe2B\\\x08~@W\xeeT" +
	"\x02FD\x0c\x08&O^\xff\xd2\xc4\x09d\xd9\xe7\x9e" +
	"~\xe7s\xeb^\x98\x11u\xe7b\xf1\x9d3\x10:u" +
	"\x0bf\xff\xe1G\x17\xbe\xff\xf9\xad\xcf\xa0\xee&\xe0\x0a" +
	"\x0f\xb8\xf4\x893K\xd0\xff\xdcy\x114\xd5\xaf\xd2\xa5" +
	"\xbfv7\xbf\x85.\x91\xe6\x8b{\x1fZ|\xf2\xd9\xed" +
	"\x7f\xf18=\x01\xf8\xf7\xb9\xdb\xfe\x94\xeb\x01\xfe\xa4k" +
	"\xd9\xcf?\xb6\xf8H\xf2\xe5\xeb?\x8c>\xf7\xaa\xbbh" +
	"\x9f{\xa3KX\xdb\xffp:|\xef\xda\x8f=\x82c" +
	"\xe7\x95BF8\xe6\xbb\xe7\xd4x\xe1\xca\x1f?\x8dL" +
	"\xce\x14\xe6\xec\xe4\xe3\x9b\xdf\xccl\xd6\xef\xf8\xe5\xb8\x05" +
	"\xe9I\xc0\xff\xd3\xdd\xf1\xff\xb1\x16.\xfd\xedV\x09\x9a" +
	"'\xbe\xbe\x1e&\xde\xde\x01TI\x1e\xb1A\x7f\xd2\xdb" +
	"\xf1\xa7=kv\xca{\xcd\xbf\xe1y\xf8\xd5d\xd1\xd5" +
	"V\xde\xc9\xae\xcb\xf3a\x90\xc6\xe9\xec|\xff\xf7\xf94" +
	"KVZ\xcdh\xa2\xd6\x08\xb2\xa0\x9d\xeb\xdb\xa4\x038" +
	"\x04\xd4\xd4\x0c\xa0'$\xf5\x05AE\x96i\xc1\xe9\xa7" +
	"\x00\xfd\x80\xa4\xbe,X\xed$i+\xa42[\xd7n" +
	"\x7f\xf0\xdcF\xbc\x0e\x90\x0a4Q\xbc\x12-%i\x04" +
	"\x80E\x08\x16\xc1\xa1\x01\xf6\x0d<S\x8b\xf2<\xb8\x1a" +
	"5H\xab9fLOt\xeeH\xb4\xc8C\xf3\x7f\xd5" +
	"\xd5,za9\xca;(\x98,\xca\xd3$\xceo\xa9" +
	"\xe3\x1c\x0fz\xa5\x157\x1b\xbd\xb0Y>\xd1\x08J\xa7" +
	"\x8a{?\xa0\xcfI\xea\xe6\xc9qK\xe1B\x10S\x99" +
	"\xbd;w/\x97\x7f\xff\xf47\x00u*V\xb5#8" +
	"\x0a*\xde\xa5\x1d\x92lHv\xed\xaa\xd3\xdb\x9d\x8f\xaa" +
	"\xf9\xf2R'\x1f\xf2\xbd\x93\xf68\x1fu\x89\x18\x10\xfb" +
	"\xbc\xa7\x930\xe8\xb4\x92\x18\xb6\xef\xf20\xf4\x0d[\xf7" +
	"K\x92\xfa\x95\x91\xd07m\xe8\x97%\xf5\xba`E\x18" +
	"#\xca\x14\x80zs\x16\xd0k\x92\xfa}\xc1\x8a<\xb4" +
	"\xb0\x04\xd4\x86\x85\xdf\x96\xd4[\x82\x15\xe7\xc0\xc2\x0e\xa0" +
	">\xb0\xf0\xbb\x92\xfa#\xc1\xd5<\xcaVZa\xc41" +
	"\x08\x8e\x81\xa5v\xd4\x098\x8en\x13\x16\x1b\x07k\xed" +
	"\xa0\xd9\xcc\xf2\x01Z\xec\xa3A\x1c.$\xd9\xe0\x8f\xb5" +
	"p9\xef$m\x9e\x85\xe0\xd9\x91\xf6\xc4\xb16\xd0\xfb" +
	"\xaa\\`x\xf8\x1c\xdc\x9d\xba8\x07\xa1&=\x1e]" +
	"1\x07\xd7\xaa\xee\xc9 \x94\xf2V\xfbm\xd6i\x06{" +
	"@\xb5\xbb\x89:\x1b\xe4\x7f\x01\x00\x00\xff\xff\x7f\xc2(" +
	"-"

func RegisterSchema(reg *schemas.Registry) {
	reg.Register(&schemas.Schema{
		String: schema_fcba4f486a351ac3,
		Nodes: []uint64{
			0xbf9edfd4684337f6,
			0xd2afeaf36c70c91f,
			0xd589c56f3d6a445e,
			0xd86baa632daef690,
			0xd9ef66060e1157d3,
			0xe61540af32cf81b6,
			0xfdee076f6379cb46,
		},
		Compressed: true,
	})
}
