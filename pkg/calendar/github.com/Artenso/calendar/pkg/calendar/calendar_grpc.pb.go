// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: calendar.proto

package calendar

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Calendar_AddEvent_FullMethodName        = "/github.com.Artenso.calendar.api.calendar.Calendar/AddEvent"
	Calendar_GetEventByID_FullMethodName    = "/github.com.Artenso.calendar.api.calendar.Calendar/GetEventByID"
	Calendar_RemoveEvent_FullMethodName     = "/github.com.Artenso.calendar.api.calendar.Calendar/RemoveEvent"
	Calendar_GetFromToEvents_FullMethodName = "/github.com.Artenso.calendar.api.calendar.Calendar/GetFromToEvents"
	Calendar_GetAllEvents_FullMethodName    = "/github.com.Artenso.calendar.api.calendar.Calendar/GetAllEvents"
	Calendar_EditEvent_FullMethodName       = "/github.com.Artenso.calendar.api.calendar.Calendar/EditEvent"
)

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalendarClient interface {
	AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventResponse, error)
	GetEventByID(ctx context.Context, in *GetEventByIDRequest, opts ...grpc.CallOption) (*GetEventByIDResponse, error)
	RemoveEvent(ctx context.Context, in *RemoveEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetFromToEvents(ctx context.Context, in *GetFromToEventsRequest, opts ...grpc.CallOption) (*GetFromToEventsResponse, error)
	GetAllEvents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllEventsResponse, error)
	EditEvent(ctx context.Context, in *EditEventRequest, opts ...grpc.CallOption) (*EditEventResponse, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventResponse, error) {
	out := new(AddEventResponse)
	err := c.cc.Invoke(ctx, Calendar_AddEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetEventByID(ctx context.Context, in *GetEventByIDRequest, opts ...grpc.CallOption) (*GetEventByIDResponse, error) {
	out := new(GetEventByIDResponse)
	err := c.cc.Invoke(ctx, Calendar_GetEventByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) RemoveEvent(ctx context.Context, in *RemoveEventRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Calendar_RemoveEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetFromToEvents(ctx context.Context, in *GetFromToEventsRequest, opts ...grpc.CallOption) (*GetFromToEventsResponse, error) {
	out := new(GetFromToEventsResponse)
	err := c.cc.Invoke(ctx, Calendar_GetFromToEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetAllEvents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllEventsResponse, error) {
	out := new(GetAllEventsResponse)
	err := c.cc.Invoke(ctx, Calendar_GetAllEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) EditEvent(ctx context.Context, in *EditEventRequest, opts ...grpc.CallOption) (*EditEventResponse, error) {
	out := new(EditEventResponse)
	err := c.cc.Invoke(ctx, Calendar_EditEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
// All implementations must embed UnimplementedCalendarServer
// for forward compatibility
type CalendarServer interface {
	AddEvent(context.Context, *AddEventRequest) (*AddEventResponse, error)
	GetEventByID(context.Context, *GetEventByIDRequest) (*GetEventByIDResponse, error)
	RemoveEvent(context.Context, *RemoveEventRequest) (*emptypb.Empty, error)
	GetFromToEvents(context.Context, *GetFromToEventsRequest) (*GetFromToEventsResponse, error)
	GetAllEvents(context.Context, *emptypb.Empty) (*GetAllEventsResponse, error)
	EditEvent(context.Context, *EditEventRequest) (*EditEventResponse, error)
	mustEmbedUnimplementedCalendarServer()
}

// UnimplementedCalendarServer must be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (UnimplementedCalendarServer) AddEvent(context.Context, *AddEventRequest) (*AddEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (UnimplementedCalendarServer) GetEventByID(context.Context, *GetEventByIDRequest) (*GetEventByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventByID not implemented")
}
func (UnimplementedCalendarServer) RemoveEvent(context.Context, *RemoveEventRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveEvent not implemented")
}
func (UnimplementedCalendarServer) GetFromToEvents(context.Context, *GetFromToEventsRequest) (*GetFromToEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFromToEvents not implemented")
}
func (UnimplementedCalendarServer) GetAllEvents(context.Context, *emptypb.Empty) (*GetAllEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEvents not implemented")
}
func (UnimplementedCalendarServer) EditEvent(context.Context, *EditEventRequest) (*EditEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditEvent not implemented")
}
func (UnimplementedCalendarServer) mustEmbedUnimplementedCalendarServer() {}

// UnsafeCalendarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalendarServer will
// result in compilation errors.
type UnsafeCalendarServer interface {
	mustEmbedUnimplementedCalendarServer()
}

func RegisterCalendarServer(s grpc.ServiceRegistrar, srv CalendarServer) {
	s.RegisterService(&Calendar_ServiceDesc, srv)
}

func _Calendar_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_AddEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).AddEvent(ctx, req.(*AddEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetEventByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetEventByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_GetEventByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetEventByID(ctx, req.(*GetEventByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_RemoveEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).RemoveEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_RemoveEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).RemoveEvent(ctx, req.(*RemoveEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetFromToEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFromToEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetFromToEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_GetFromToEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetFromToEvents(ctx, req.(*GetFromToEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetAllEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetAllEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_GetAllEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetAllEvents(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_EditEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).EditEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Calendar_EditEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).EditEvent(ctx, req.(*EditEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Calendar_ServiceDesc is the grpc.ServiceDesc for Calendar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calendar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.Artenso.calendar.api.calendar.Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddEvent",
			Handler:    _Calendar_AddEvent_Handler,
		},
		{
			MethodName: "GetEventByID",
			Handler:    _Calendar_GetEventByID_Handler,
		},
		{
			MethodName: "RemoveEvent",
			Handler:    _Calendar_RemoveEvent_Handler,
		},
		{
			MethodName: "GetFromToEvents",
			Handler:    _Calendar_GetFromToEvents_Handler,
		},
		{
			MethodName: "GetAllEvents",
			Handler:    _Calendar_GetAllEvents_Handler,
		},
		{
			MethodName: "EditEvent",
			Handler:    _Calendar_EditEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calendar.proto",
}
