// Code generated by capnpc-go. DO NOT EDIT.

package proc

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	stream "capnproto.org/go/capnp/v3/std/capnp/stream"
	context "context"
	channel "github.com/wetware/ww/internal/api/channel"
)

type Executor struct{ Client *capnp.Client }

// Executor_TypeID is the unique identifier for the type Executor.
const Executor_TypeID = 0xe8bb307fa2f406fb

func (c Executor) Exec(ctx context.Context, params func(Executor_exec_Params) error) (Executor_exec_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xe8bb307fa2f406fb,
			MethodID:      0,
			InterfaceName: "proc.capnp:Executor",
			MethodName:    "exec",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Executor_exec_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Executor_exec_Results_Future{Future: ans.Future()}, release
}

func (c Executor) AddRef() Executor {
	return Executor{
		Client: c.Client.AddRef(),
	}
}

func (c Executor) Release() {
	c.Client.Release()
}

// A Executor_Server is a Executor with a local implementation.
type Executor_Server interface {
	Exec(context.Context, Executor_exec) error
}

// Executor_NewServer creates a new Server from an implementation of Executor_Server.
func Executor_NewServer(s Executor_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Executor_Methods(nil, s), s, c, policy)
}

// Executor_ServerToClient creates a new Client from an implementation of Executor_Server.
// The caller is responsible for calling Release on the returned Client.
func Executor_ServerToClient(s Executor_Server, policy *server.Policy) Executor {
	return Executor{Client: capnp.NewClient(Executor_NewServer(s, policy))}
}

// Executor_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Executor_Methods(methods []server.Method, s Executor_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe8bb307fa2f406fb,
			MethodID:      0,
			InterfaceName: "proc.capnp:Executor",
			MethodName:    "exec",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Exec(ctx, Executor_exec{call})
		},
	})

	return methods
}

// Executor_exec holds the state for a server call to Executor.exec.
// See server.Call for documentation.
type Executor_exec struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Executor_exec) Args() Executor_exec_Params {
	return Executor_exec_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c Executor_exec) AllocResults() (Executor_exec_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Executor_exec_Results{Struct: r}, err
}

type Executor_exec_Params struct{ capnp.Struct }

// Executor_exec_Params_TypeID is the unique identifier for the type Executor_exec_Params.
const Executor_exec_Params_TypeID = 0xaf67b0a40b1c2bea

func NewExecutor_exec_Params(s *capnp.Segment) (Executor_exec_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Executor_exec_Params{st}, err
}

func NewRootExecutor_exec_Params(s *capnp.Segment) (Executor_exec_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Executor_exec_Params{st}, err
}

func ReadRootExecutor_exec_Params(msg *capnp.Message) (Executor_exec_Params, error) {
	root, err := msg.Root()
	return Executor_exec_Params{root.Struct()}, err
}

func (s Executor_exec_Params) String() string {
	str, _ := text.Marshal(0xaf67b0a40b1c2bea, s.Struct)
	return str
}

func (s Executor_exec_Params) Param() (capnp.Ptr, error) {
	return s.Struct.Ptr(0)
}

func (s Executor_exec_Params) HasParam() bool {
	return s.Struct.HasPtr(0)
}

func (s Executor_exec_Params) SetParam(v capnp.Ptr) error {
	return s.Struct.SetPtr(0, v)
}

// Executor_exec_Params_List is a list of Executor_exec_Params.
type Executor_exec_Params_List = capnp.StructList[Executor_exec_Params]

// NewExecutor_exec_Params creates a new list of Executor_exec_Params.
func NewExecutor_exec_Params_List(s *capnp.Segment, sz int32) (Executor_exec_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Executor_exec_Params]{List: l}, err
}

// Executor_exec_Params_Future is a wrapper for a Executor_exec_Params promised by a client call.
type Executor_exec_Params_Future struct{ *capnp.Future }

func (p Executor_exec_Params_Future) Struct() (Executor_exec_Params, error) {
	s, err := p.Future.Struct()
	return Executor_exec_Params{s}, err
}

func (p Executor_exec_Params_Future) Param() *capnp.Future {
	return p.Future.Field(0, nil)
}

type Executor_exec_Results struct{ capnp.Struct }

// Executor_exec_Results_TypeID is the unique identifier for the type Executor_exec_Results.
const Executor_exec_Results_TypeID = 0x8d124035fd940437

func NewExecutor_exec_Results(s *capnp.Segment) (Executor_exec_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Executor_exec_Results{st}, err
}

func NewRootExecutor_exec_Results(s *capnp.Segment) (Executor_exec_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Executor_exec_Results{st}, err
}

func ReadRootExecutor_exec_Results(msg *capnp.Message) (Executor_exec_Results, error) {
	root, err := msg.Root()
	return Executor_exec_Results{root.Struct()}, err
}

func (s Executor_exec_Results) String() string {
	str, _ := text.Marshal(0x8d124035fd940437, s.Struct)
	return str
}

func (s Executor_exec_Results) Proc() P {
	p, _ := s.Struct.Ptr(0)
	return P{Client: p.Interface().Client()}
}

func (s Executor_exec_Results) HasProc() bool {
	return s.Struct.HasPtr(0)
}

func (s Executor_exec_Results) SetProc(v P) error {
	if !v.Client.IsValid() {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Executor_exec_Results_List is a list of Executor_exec_Results.
type Executor_exec_Results_List = capnp.StructList[Executor_exec_Results]

// NewExecutor_exec_Results creates a new list of Executor_exec_Results.
func NewExecutor_exec_Results_List(s *capnp.Segment, sz int32) (Executor_exec_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Executor_exec_Results]{List: l}, err
}

// Executor_exec_Results_Future is a wrapper for a Executor_exec_Results promised by a client call.
type Executor_exec_Results_Future struct{ *capnp.Future }

func (p Executor_exec_Results_Future) Struct() (Executor_exec_Results, error) {
	s, err := p.Future.Struct()
	return Executor_exec_Results{s}, err
}

func (p Executor_exec_Results_Future) Proc() P {
	return P{Client: p.Future.Field(0, nil).Client()}
}

type P struct{ Client *capnp.Client }

// P_TypeID is the unique identifier for the type P.
const P_TypeID = 0xe19c553506f6045b

func (c P) Wait(ctx context.Context, params func(P_wait_Params) error) (P_wait_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xe19c553506f6045b,
			MethodID:      0,
			InterfaceName: "proc.capnp:P",
			MethodName:    "wait",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(P_wait_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return P_wait_Results_Future{Future: ans.Future()}, release
}

func (c P) AddRef() P {
	return P{
		Client: c.Client.AddRef(),
	}
}

func (c P) Release() {
	c.Client.Release()
}

// A P_Server is a P with a local implementation.
type P_Server interface {
	Wait(context.Context, P_wait) error
}

// P_NewServer creates a new Server from an implementation of P_Server.
func P_NewServer(s P_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(P_Methods(nil, s), s, c, policy)
}

// P_ServerToClient creates a new Client from an implementation of P_Server.
// The caller is responsible for calling Release on the returned Client.
func P_ServerToClient(s P_Server, policy *server.Policy) P {
	return P{Client: capnp.NewClient(P_NewServer(s, policy))}
}

// P_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func P_Methods(methods []server.Method, s P_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe19c553506f6045b,
			MethodID:      0,
			InterfaceName: "proc.capnp:P",
			MethodName:    "wait",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Wait(ctx, P_wait{call})
		},
	})

	return methods
}

// P_wait holds the state for a server call to P.wait.
// See server.Call for documentation.
type P_wait struct {
	*server.Call
}

// Args returns the call's arguments.
func (c P_wait) Args() P_wait_Params {
	return P_wait_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c P_wait) AllocResults() (P_wait_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return P_wait_Results{Struct: r}, err
}

type P_wait_Params struct{ capnp.Struct }

// P_wait_Params_TypeID is the unique identifier for the type P_wait_Params.
const P_wait_Params_TypeID = 0xbf26ab53506a6190

func NewP_wait_Params(s *capnp.Segment) (P_wait_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return P_wait_Params{st}, err
}

func NewRootP_wait_Params(s *capnp.Segment) (P_wait_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return P_wait_Params{st}, err
}

func ReadRootP_wait_Params(msg *capnp.Message) (P_wait_Params, error) {
	root, err := msg.Root()
	return P_wait_Params{root.Struct()}, err
}

func (s P_wait_Params) String() string {
	str, _ := text.Marshal(0xbf26ab53506a6190, s.Struct)
	return str
}

// P_wait_Params_List is a list of P_wait_Params.
type P_wait_Params_List = capnp.StructList[P_wait_Params]

// NewP_wait_Params creates a new list of P_wait_Params.
func NewP_wait_Params_List(s *capnp.Segment, sz int32) (P_wait_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[P_wait_Params]{List: l}, err
}

// P_wait_Params_Future is a wrapper for a P_wait_Params promised by a client call.
type P_wait_Params_Future struct{ *capnp.Future }

func (p P_wait_Params_Future) Struct() (P_wait_Params, error) {
	s, err := p.Future.Struct()
	return P_wait_Params{s}, err
}

type P_wait_Results struct{ capnp.Struct }

// P_wait_Results_TypeID is the unique identifier for the type P_wait_Results.
const P_wait_Results_TypeID = 0xa9eb6dedccb8d3ff

func NewP_wait_Results(s *capnp.Segment) (P_wait_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return P_wait_Results{st}, err
}

func NewRootP_wait_Results(s *capnp.Segment) (P_wait_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return P_wait_Results{st}, err
}

func ReadRootP_wait_Results(msg *capnp.Message) (P_wait_Results, error) {
	root, err := msg.Root()
	return P_wait_Results{root.Struct()}, err
}

func (s P_wait_Results) String() string {
	str, _ := text.Marshal(0xa9eb6dedccb8d3ff, s.Struct)
	return str
}

// P_wait_Results_List is a list of P_wait_Results.
type P_wait_Results_List = capnp.StructList[P_wait_Results]

// NewP_wait_Results creates a new list of P_wait_Results.
func NewP_wait_Results_List(s *capnp.Segment, sz int32) (P_wait_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[P_wait_Results]{List: l}, err
}

// P_wait_Results_Future is a wrapper for a P_wait_Results promised by a client call.
type P_wait_Results_Future struct{ *capnp.Future }

func (p P_wait_Results_Future) Struct() (P_wait_Results, error) {
	s, err := p.Future.Struct()
	return P_wait_Results{s}, err
}

type Unix struct{ Client *capnp.Client }

// Unix_TypeID is the unique identifier for the type Unix.
const Unix_TypeID = 0x85f7549a53596cef

func (c Unix) Exec(ctx context.Context, params func(Executor_exec_Params) error) (Executor_exec_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xe8bb307fa2f406fb,
			MethodID:      0,
			InterfaceName: "proc.capnp:Executor",
			MethodName:    "exec",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Executor_exec_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Executor_exec_Results_Future{Future: ans.Future()}, release
}

func (c Unix) AddRef() Unix {
	return Unix{
		Client: c.Client.AddRef(),
	}
}

func (c Unix) Release() {
	c.Client.Release()
}

// A Unix_Server is a Unix with a local implementation.
type Unix_Server interface {
	Exec(context.Context, Executor_exec) error
}

// Unix_NewServer creates a new Server from an implementation of Unix_Server.
func Unix_NewServer(s Unix_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Unix_Methods(nil, s), s, c, policy)
}

// Unix_ServerToClient creates a new Client from an implementation of Unix_Server.
// The caller is responsible for calling Release on the returned Client.
func Unix_ServerToClient(s Unix_Server, policy *server.Policy) Unix {
	return Unix{Client: capnp.NewClient(Unix_NewServer(s, policy))}
}

// Unix_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Unix_Methods(methods []server.Method, s Unix_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe8bb307fa2f406fb,
			MethodID:      0,
			InterfaceName: "proc.capnp:Executor",
			MethodName:    "exec",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Exec(ctx, Executor_exec{call})
		},
	})

	return methods
}

type Unix_Command struct{ capnp.Struct }

// Unix_Command_TypeID is the unique identifier for the type Unix_Command.
const Unix_Command_TypeID = 0x8e898dedb95cdee4

func NewUnix_Command(s *capnp.Segment) (Unix_Command, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 7})
	return Unix_Command{st}, err
}

func NewRootUnix_Command(s *capnp.Segment) (Unix_Command, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 7})
	return Unix_Command{st}, err
}

func ReadRootUnix_Command(msg *capnp.Message) (Unix_Command, error) {
	root, err := msg.Root()
	return Unix_Command{root.Struct()}, err
}

func (s Unix_Command) String() string {
	str, _ := text.Marshal(0x8e898dedb95cdee4, s.Struct)
	return str
}

func (s Unix_Command) Path() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Unix_Command) HasPath() bool {
	return s.Struct.HasPtr(0)
}

func (s Unix_Command) PathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Unix_Command) SetPath(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Unix_Command) Dir() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Unix_Command) HasDir() bool {
	return s.Struct.HasPtr(1)
}

func (s Unix_Command) DirBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Unix_Command) SetDir(v string) error {
	return s.Struct.SetText(1, v)
}

func (s Unix_Command) Args() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.TextList{List: p.List()}, err
}

func (s Unix_Command) HasArgs() bool {
	return s.Struct.HasPtr(2)
}

func (s Unix_Command) SetArgs(v capnp.TextList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewArgs sets the args field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Unix_Command) NewArgs(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s Unix_Command) Env() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(3)
	return capnp.TextList{List: p.List()}, err
}

func (s Unix_Command) HasEnv() bool {
	return s.Struct.HasPtr(3)
}

func (s Unix_Command) SetEnv(v capnp.TextList) error {
	return s.Struct.SetPtr(3, v.List.ToPtr())
}

// NewEnv sets the env field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Unix_Command) NewEnv(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(3, l.List.ToPtr())
	return l, err
}

func (s Unix_Command) Stdin() Unix_InputStream {
	p, _ := s.Struct.Ptr(4)
	return Unix_InputStream{Client: p.Interface().Client()}
}

func (s Unix_Command) HasStdin() bool {
	return s.Struct.HasPtr(4)
}

func (s Unix_Command) SetStdin(v Unix_InputStream) error {
	if !v.Client.IsValid() {
		return s.Struct.SetPtr(4, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(4, in.ToPtr())
}

func (s Unix_Command) Stdout() Unix_OutputStream {
	p, _ := s.Struct.Ptr(5)
	return Unix_OutputStream{Client: p.Interface().Client()}
}

func (s Unix_Command) HasStdout() bool {
	return s.Struct.HasPtr(5)
}

func (s Unix_Command) SetStdout(v Unix_OutputStream) error {
	if !v.Client.IsValid() {
		return s.Struct.SetPtr(5, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(5, in.ToPtr())
}

func (s Unix_Command) Stderr() Unix_OutputStream {
	p, _ := s.Struct.Ptr(6)
	return Unix_OutputStream{Client: p.Interface().Client()}
}

func (s Unix_Command) HasStderr() bool {
	return s.Struct.HasPtr(6)
}

func (s Unix_Command) SetStderr(v Unix_OutputStream) error {
	if !v.Client.IsValid() {
		return s.Struct.SetPtr(6, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(6, in.ToPtr())
}

// Unix_Command_List is a list of Unix_Command.
type Unix_Command_List = capnp.StructList[Unix_Command]

// NewUnix_Command creates a new list of Unix_Command.
func NewUnix_Command_List(s *capnp.Segment, sz int32) (Unix_Command_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 7}, sz)
	return capnp.StructList[Unix_Command]{List: l}, err
}

// Unix_Command_Future is a wrapper for a Unix_Command promised by a client call.
type Unix_Command_Future struct{ *capnp.Future }

func (p Unix_Command_Future) Struct() (Unix_Command, error) {
	s, err := p.Future.Struct()
	return Unix_Command{s}, err
}

func (p Unix_Command_Future) Stdin() Unix_InputStream {
	return Unix_InputStream{Client: p.Future.Field(4, nil).Client()}
}

func (p Unix_Command_Future) Stdout() Unix_OutputStream {
	return Unix_OutputStream{Client: p.Future.Field(5, nil).Client()}
}

func (p Unix_Command_Future) Stderr() Unix_OutputStream {
	return Unix_OutputStream{Client: p.Future.Field(6, nil).Client()}
}

type Unix_InputStream struct{ Client *capnp.Client }

// Unix_InputStream_TypeID is the unique identifier for the type Unix_InputStream.
const Unix_InputStream_TypeID = 0xd134f1942f6607f0

func (c Unix_InputStream) Recv(ctx context.Context, params func(channel.Recver_recv_Params) error) (channel.Recver_recv_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xdf05a90d671c0c07,
			MethodID:      0,
			InterfaceName: "channel.capnp:Recver",
			MethodName:    "recv",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(channel.Recver_recv_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return channel.Recver_recv_Results_Future{Future: ans.Future()}, release
}

func (c Unix_InputStream) AddRef() Unix_InputStream {
	return Unix_InputStream{
		Client: c.Client.AddRef(),
	}
}

func (c Unix_InputStream) Release() {
	c.Client.Release()
}

// A Unix_InputStream_Server is a Unix_InputStream with a local implementation.
type Unix_InputStream_Server interface {
	Recv(context.Context, channel.Recver_recv) error
}

// Unix_InputStream_NewServer creates a new Server from an implementation of Unix_InputStream_Server.
func Unix_InputStream_NewServer(s Unix_InputStream_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Unix_InputStream_Methods(nil, s), s, c, policy)
}

// Unix_InputStream_ServerToClient creates a new Client from an implementation of Unix_InputStream_Server.
// The caller is responsible for calling Release on the returned Client.
func Unix_InputStream_ServerToClient(s Unix_InputStream_Server, policy *server.Policy) Unix_InputStream {
	return Unix_InputStream{Client: capnp.NewClient(Unix_InputStream_NewServer(s, policy))}
}

// Unix_InputStream_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Unix_InputStream_Methods(methods []server.Method, s Unix_InputStream_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xdf05a90d671c0c07,
			MethodID:      0,
			InterfaceName: "channel.capnp:Recver",
			MethodName:    "recv",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Recv(ctx, channel.Recver_recv{call})
		},
	})

	return methods
}

type Unix_OutputStream struct{ Client *capnp.Client }

// Unix_OutputStream_TypeID is the unique identifier for the type Unix_OutputStream.
const Unix_OutputStream_TypeID = 0x82ce9b5bf259d963

func (c Unix_OutputStream) Send(ctx context.Context, params func(channel.Sender_send_Params) error) (stream.StreamResult_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xe8bbed1438ea16ee,
			MethodID:      0,
			InterfaceName: "channel.capnp:Sender",
			MethodName:    "send",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		s.PlaceArgs = func(s capnp.Struct) error { return params(channel.Sender_send_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return stream.StreamResult_Future{Future: ans.Future()}, release
}
func (c Unix_OutputStream) Close(ctx context.Context, params func(channel.Closer_close_Params) error) (channel.Closer_close_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xfad0e4b80d3779c3,
			MethodID:      0,
			InterfaceName: "channel.capnp:Closer",
			MethodName:    "close",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(channel.Closer_close_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return channel.Closer_close_Results_Future{Future: ans.Future()}, release
}

func (c Unix_OutputStream) AddRef() Unix_OutputStream {
	return Unix_OutputStream{
		Client: c.Client.AddRef(),
	}
}

func (c Unix_OutputStream) Release() {
	c.Client.Release()
}

// A Unix_OutputStream_Server is a Unix_OutputStream with a local implementation.
type Unix_OutputStream_Server interface {
	Send(context.Context, channel.Sender_send) error

	Close(context.Context, channel.Closer_close) error
}

// Unix_OutputStream_NewServer creates a new Server from an implementation of Unix_OutputStream_Server.
func Unix_OutputStream_NewServer(s Unix_OutputStream_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Unix_OutputStream_Methods(nil, s), s, c, policy)
}

// Unix_OutputStream_ServerToClient creates a new Client from an implementation of Unix_OutputStream_Server.
// The caller is responsible for calling Release on the returned Client.
func Unix_OutputStream_ServerToClient(s Unix_OutputStream_Server, policy *server.Policy) Unix_OutputStream {
	return Unix_OutputStream{Client: capnp.NewClient(Unix_OutputStream_NewServer(s, policy))}
}

// Unix_OutputStream_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Unix_OutputStream_Methods(methods []server.Method, s Unix_OutputStream_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe8bbed1438ea16ee,
			MethodID:      0,
			InterfaceName: "channel.capnp:Sender",
			MethodName:    "send",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Send(ctx, channel.Sender_send{call})
		},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xfad0e4b80d3779c3,
			MethodID:      0,
			InterfaceName: "channel.capnp:Closer",
			MethodName:    "close",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.Close(ctx, channel.Closer_close{call})
		},
	})

	return methods
}

const schema_d78885a0de56b292 = "x\xda|T]h\x1cU\x18=\xe7\xce\x9d\xb9\xd6\xee" +
	"t\xf7f\xa6\x14!\x10\x10\x15\xac$\xd6\xd6P\x08H" +
	"\x13E\x8a/\xba\xd74B\xa8>\x8c\x9b5&t\x7f" +
	"\x98\x9d\xd5\xf4A\x0a\x85\x82>TE\xfaT}\x90\"" +
	"\x94\xe2\x83E\x14\x1aT\xb4B_\xa4\x08QA|\xa9" +
	"V\x1aH\xaa\xc1*\xad\"\"S\xeevwg7[" +
	"\xfa6|\xf7\xcc\xf9\xcew\xcew\xef\xae\x191)\x1f" +
	"\xf1Oo\x810u\xd7KK?\xcd\xfeu\xf0\xddo" +
	"\x8fB\x17\x9c\xf4\x8fC\xb3\xd3'\x0f\xfcs\x0c`\xb0" +
	",\xaf\x06\x17\xa4\x02\x82\xf3R\x05\xe7\xe5\x0e ;\xd7" +
	"[\x9d\xf4\x9d\x8f\x9f\xbb\xf4\xfe\xb1\xd7\x7f\xb4\xe0\xcb\xf2" +
	"h\xb0*\xf7\x02\x81\xef\xaa\xc0w-x\xaf<\xf1\xff" +
	"\xf8\xe4\xd0q\xe8\x90\x80KU\xe0\x9eQ\xf7n\x82\xc1" +
	"\xb8\xbb\x0fL\xaf\\z~y\xe3\xf8\x1boB\x17\x98" +
	"Q\xbb\xca\xf6\x9cq\xbf\x0e^p\x15\xb0g\xd6}\x8b" +
	"`\x9a~\x7f\xee\xe2F\xe5\xb73\xd0>\x81\x96.\xaa" +
	"\xef \xd3\xab\x0f\x0do\xfd\xe0\xec\xfcG\xbdmV=" +
	"a\xdb\xac{\xb6\xcd\xdb\xd1bq\xfa\xc3\x07\xbe\xec\xf9" +
	"s\x8b\xfa\x062\xbd\xa6^z\xf8\xc4\x9f\x8f\xae\x0c\x8c" +
	"\xbe\xe1\xfd\x1a\xfc\xebY\xe0\x0dO\x057<;\xcdA" +
	"\xf9\xb77>\xf3\xde\xe5\x81\xd1\xb7\xab\xc5\xe0\x9e\x96\xe6" +
	"\xedj\x7f\xf0\x98\xfdJ\xff\xf3\xae\x9f:\xb2\xeb\xb3\xb5" +
	"M`\x14\x18\xdc\xafN\x05\xa3-\xf8\x83j\x7fpX" +
	"\xa9\xe0\xb0\xca\xe3BZ\x8fk\xa5\xb1RT\x17\xd5\xfa" +
	"\xc4Luai\xec\x99fRo&\xd3I>.G" +
	"\x95\"Yt\\#\xc9\xf4\xe9\xf5\xf0\xb5\x93+\xa7\xd7" +
	"1I\xcd\x11#EO\x09\xd0\xdcaQ\x16N\xfa\x10" +
	"]^\xb6yar\xec\xb5^?\xde\xe3\x83\xffb\xcf" +
	">\xf8\x8bG\x9e\xa8U*Qu.}\xaa\xda\x92\x12" +
	"C\x95\xa3J\xdaQ\x86[\xd2\xda\xba\xba3g\xba\xba" +
	"\xa5M\xba\x0aY\x7f\xb0\xab\xd0\xa9\xd6'\x9e\\*\x97" +
	"\x9aI-\x1e+/\x95K\xf7=[n4\x0f%\x0d" +
	"\xc0HG\x02\x92\x80\xf6w\x02\xe6.\x87&\x14\xcc\xdb" +
	"_\xa9\xb3l@\xea\x1e\xc6\xae\x97\xed9P$\xcdp" +
	"\x97\xeaSKu\xd6\xa1\xf9\\P\x93!mq\xf9^" +
	"\xc0|\xe2\xd0|%\xa8\x85\x08)\x00\xfd\x85E\x9es" +
	"h~\x10\xd4\x8e\x13\xd2\x01\xf4\x8aE^th\xd6\x04" +
	"\xb5\x94!%\xa0Ww\x03\xe6\x17\x87\xe6wA\xed\xba" +
	"!]@\xafO\x00\xe6\x8aCsMP{^H\x0f" +
	"\xd0\x1b\xb6\xb8\xe6\xd0\\\xb7\x83D\xc9\xcb\xccA0\x07" +
	"\xaa\xb9\x85\xb8\xf3\x9d\x8f\xe2\xf9\x06\xb7\xa1e\x9b\xadm" +
	"\x03U\xb9\xfa\xca\xa6\xd2H#\x99[\xa8RgQ\xde" +
	"rb_#\x99\xab5\x13\xea,\xd6\xec\xa0\x1c\xc7\x83" +
	"\x07}\xde\x15\xc7^\x8d\x16\x92v\x0cl\xdc!\xa9b" +
	"\x14G\x95F_N\xbb\xb3\x9cF\xea\xf6\x98C\xec[" +
	"\x09\x0e\xdd\xbe]\x9b\x0b\x831v\xd6P\xf5\xdf\x08\x95" +
	"\x1b\x9e\xf7\xcf\xb8?\xf7l^\xb7t\x87\x1b\x81\xbc\xed" +
	"iWB:.\xd0}'\xd8yj\xb4\xde\x09\xa1]" +
	"\x95\xb7\xb2&Y$\xfbnS\xcb\x01\x95\xd4\xe2\x8c\xa2" +
	"\xf3\x16\xb1\xf3\xf6Y\x0aL\xe585L=\xaa\xf2\xd6" +
	"\xab\x81\xcb\xd1Z\xbb\xdb\x15[-\xa7$5\x87x\xe0" +
	"f\x00\x00\x00\xff\xff$\xfa\x9b\xbf"

func init() {
	schemas.Register(schema_d78885a0de56b292,
		0x82ce9b5bf259d963,
		0x85f7549a53596cef,
		0x8d124035fd940437,
		0x8e898dedb95cdee4,
		0xa9eb6dedccb8d3ff,
		0xaf67b0a40b1c2bea,
		0xbf26ab53506a6190,
		0xd134f1942f6607f0,
		0xe19c553506f6045b,
		0xe8bb307fa2f406fb)
}
