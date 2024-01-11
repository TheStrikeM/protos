// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/ticket.proto

package __

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ComplainClient is the client API for Complain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComplainClient interface {
	AllTickets(ctx context.Context, in *AllTicketsRequest, opts ...grpc.CallOption) (*AllTicketsResponse, error)
	CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error)
	ResolveTicket(ctx context.Context, in *ResolveTicketRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	TicketByTrader(ctx context.Context, in *TicketByTraderRequest, opts ...grpc.CallOption) (*TicketByTraderResponse, error)
}

type complainClient struct {
	cc grpc.ClientConnInterface
}

func NewComplainClient(cc grpc.ClientConnInterface) ComplainClient {
	return &complainClient{cc}
}

func (c *complainClient) AllTickets(ctx context.Context, in *AllTicketsRequest, opts ...grpc.CallOption) (*AllTicketsResponse, error) {
	out := new(AllTicketsResponse)
	err := c.cc.Invoke(ctx, "/ticket_proto.Complain/AllTickets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complainClient) CreateTicket(ctx context.Context, in *CreateTicketRequest, opts ...grpc.CallOption) (*CreateTicketResponse, error) {
	out := new(CreateTicketResponse)
	err := c.cc.Invoke(ctx, "/ticket_proto.Complain/CreateTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complainClient) ResolveTicket(ctx context.Context, in *ResolveTicketRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ticket_proto.Complain/ResolveTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *complainClient) TicketByTrader(ctx context.Context, in *TicketByTraderRequest, opts ...grpc.CallOption) (*TicketByTraderResponse, error) {
	out := new(TicketByTraderResponse)
	err := c.cc.Invoke(ctx, "/ticket_proto.Complain/TicketByTrader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComplainServer is the server API for Complain service.
// All implementations must embed UnimplementedComplainServer
// for forward compatibility
type ComplainServer interface {
	AllTickets(context.Context, *AllTicketsRequest) (*AllTicketsResponse, error)
	CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error)
	ResolveTicket(context.Context, *ResolveTicketRequest) (*empty.Empty, error)
	TicketByTrader(context.Context, *TicketByTraderRequest) (*TicketByTraderResponse, error)
	mustEmbedUnimplementedComplainServer()
}

// UnimplementedComplainServer must be embedded to have forward compatible implementations.
type UnimplementedComplainServer struct {
}

func (UnimplementedComplainServer) AllTickets(context.Context, *AllTicketsRequest) (*AllTicketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllTickets not implemented")
}
func (UnimplementedComplainServer) CreateTicket(context.Context, *CreateTicketRequest) (*CreateTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTicket not implemented")
}
func (UnimplementedComplainServer) ResolveTicket(context.Context, *ResolveTicketRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveTicket not implemented")
}
func (UnimplementedComplainServer) TicketByTrader(context.Context, *TicketByTraderRequest) (*TicketByTraderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TicketByTrader not implemented")
}
func (UnimplementedComplainServer) mustEmbedUnimplementedComplainServer() {}

// UnsafeComplainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComplainServer will
// result in compilation errors.
type UnsafeComplainServer interface {
	mustEmbedUnimplementedComplainServer()
}

func RegisterComplainServer(s grpc.ServiceRegistrar, srv ComplainServer) {
	s.RegisterService(&Complain_ServiceDesc, srv)
}

func _Complain_AllTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllTicketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplainServer).AllTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket_proto.Complain/AllTickets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplainServer).AllTickets(ctx, req.(*AllTicketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Complain_CreateTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplainServer).CreateTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket_proto.Complain/CreateTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplainServer).CreateTicket(ctx, req.(*CreateTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Complain_ResolveTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplainServer).ResolveTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket_proto.Complain/ResolveTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplainServer).ResolveTicket(ctx, req.(*ResolveTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Complain_TicketByTrader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketByTraderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComplainServer).TicketByTrader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ticket_proto.Complain/TicketByTrader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComplainServer).TicketByTrader(ctx, req.(*TicketByTraderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Complain_ServiceDesc is the grpc.ServiceDesc for Complain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Complain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticket_proto.Complain",
	HandlerType: (*ComplainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllTickets",
			Handler:    _Complain_AllTickets_Handler,
		},
		{
			MethodName: "CreateTicket",
			Handler:    _Complain_CreateTicket_Handler,
		},
		{
			MethodName: "ResolveTicket",
			Handler:    _Complain_ResolveTicket_Handler,
		},
		{
			MethodName: "TicketByTrader",
			Handler:    _Complain_TicketByTrader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ticket.proto",
}
