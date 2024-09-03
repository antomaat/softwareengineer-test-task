// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.6.1
// source: ticket_score/ticket_score.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TicketScores_GetAggregatedCategoryScores_FullMethodName    = "/protos.TicketScores/GetAggregatedCategoryScores"
	TicketScores_GetScoresByTicket_FullMethodName              = "/protos.TicketScores/GetScoresByTicket"
	TicketScores_GetOverallQualityScores_FullMethodName        = "/protos.TicketScores/GetOverallQualityScores"
	TicketScores_GetPeriodOverPeriodScoreChange_FullMethodName = "/protos.TicketScores/GetPeriodOverPeriodScoreChange"
)

// TicketScoresClient is the client API for TicketScores service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TicketScoresClient interface {
	GetAggregatedCategoryScores(ctx context.Context, in *GetAggregatedCategoryScoresRequest, opts ...grpc.CallOption) (*GetAggregatedCategoryScoresResponse, error)
	GetScoresByTicket(ctx context.Context, in *GetScoresByTicketRequest, opts ...grpc.CallOption) (*GetScoresByTicketResponse, error)
	GetOverallQualityScores(ctx context.Context, in *GetOverallQualityScoresRequest, opts ...grpc.CallOption) (*GetOverallQualityScoresResponse, error)
	GetPeriodOverPeriodScoreChange(ctx context.Context, in *GetPeriodOverPeriodScoreChangeRequest, opts ...grpc.CallOption) (*GetPeriodOverPeriodScoreChangeResponse, error)
}

type ticketScoresClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketScoresClient(cc grpc.ClientConnInterface) TicketScoresClient {
	return &ticketScoresClient{cc}
}

func (c *ticketScoresClient) GetAggregatedCategoryScores(ctx context.Context, in *GetAggregatedCategoryScoresRequest, opts ...grpc.CallOption) (*GetAggregatedCategoryScoresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAggregatedCategoryScoresResponse)
	err := c.cc.Invoke(ctx, TicketScores_GetAggregatedCategoryScores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketScoresClient) GetScoresByTicket(ctx context.Context, in *GetScoresByTicketRequest, opts ...grpc.CallOption) (*GetScoresByTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetScoresByTicketResponse)
	err := c.cc.Invoke(ctx, TicketScores_GetScoresByTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketScoresClient) GetOverallQualityScores(ctx context.Context, in *GetOverallQualityScoresRequest, opts ...grpc.CallOption) (*GetOverallQualityScoresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOverallQualityScoresResponse)
	err := c.cc.Invoke(ctx, TicketScores_GetOverallQualityScores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketScoresClient) GetPeriodOverPeriodScoreChange(ctx context.Context, in *GetPeriodOverPeriodScoreChangeRequest, opts ...grpc.CallOption) (*GetPeriodOverPeriodScoreChangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPeriodOverPeriodScoreChangeResponse)
	err := c.cc.Invoke(ctx, TicketScores_GetPeriodOverPeriodScoreChange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketScoresServer is the server API for TicketScores service.
// All implementations must embed UnimplementedTicketScoresServer
// for forward compatibility.
type TicketScoresServer interface {
	GetAggregatedCategoryScores(context.Context, *GetAggregatedCategoryScoresRequest) (*GetAggregatedCategoryScoresResponse, error)
	GetScoresByTicket(context.Context, *GetScoresByTicketRequest) (*GetScoresByTicketResponse, error)
	GetOverallQualityScores(context.Context, *GetOverallQualityScoresRequest) (*GetOverallQualityScoresResponse, error)
	GetPeriodOverPeriodScoreChange(context.Context, *GetPeriodOverPeriodScoreChangeRequest) (*GetPeriodOverPeriodScoreChangeResponse, error)
	mustEmbedUnimplementedTicketScoresServer()
}

// UnimplementedTicketScoresServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTicketScoresServer struct{}

func (UnimplementedTicketScoresServer) GetAggregatedCategoryScores(context.Context, *GetAggregatedCategoryScoresRequest) (*GetAggregatedCategoryScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAggregatedCategoryScores not implemented")
}
func (UnimplementedTicketScoresServer) GetScoresByTicket(context.Context, *GetScoresByTicketRequest) (*GetScoresByTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScoresByTicket not implemented")
}
func (UnimplementedTicketScoresServer) GetOverallQualityScores(context.Context, *GetOverallQualityScoresRequest) (*GetOverallQualityScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOverallQualityScores not implemented")
}
func (UnimplementedTicketScoresServer) GetPeriodOverPeriodScoreChange(context.Context, *GetPeriodOverPeriodScoreChangeRequest) (*GetPeriodOverPeriodScoreChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeriodOverPeriodScoreChange not implemented")
}
func (UnimplementedTicketScoresServer) mustEmbedUnimplementedTicketScoresServer() {}
func (UnimplementedTicketScoresServer) testEmbeddedByValue()                      {}

// UnsafeTicketScoresServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketScoresServer will
// result in compilation errors.
type UnsafeTicketScoresServer interface {
	mustEmbedUnimplementedTicketScoresServer()
}

func RegisterTicketScoresServer(s grpc.ServiceRegistrar, srv TicketScoresServer) {
	// If the following call pancis, it indicates UnimplementedTicketScoresServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TicketScores_ServiceDesc, srv)
}

func _TicketScores_GetAggregatedCategoryScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAggregatedCategoryScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketScoresServer).GetAggregatedCategoryScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketScores_GetAggregatedCategoryScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketScoresServer).GetAggregatedCategoryScores(ctx, req.(*GetAggregatedCategoryScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketScores_GetScoresByTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScoresByTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketScoresServer).GetScoresByTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketScores_GetScoresByTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketScoresServer).GetScoresByTicket(ctx, req.(*GetScoresByTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketScores_GetOverallQualityScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOverallQualityScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketScoresServer).GetOverallQualityScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketScores_GetOverallQualityScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketScoresServer).GetOverallQualityScores(ctx, req.(*GetOverallQualityScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketScores_GetPeriodOverPeriodScoreChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeriodOverPeriodScoreChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketScoresServer).GetPeriodOverPeriodScoreChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketScores_GetPeriodOverPeriodScoreChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketScoresServer).GetPeriodOverPeriodScoreChange(ctx, req.(*GetPeriodOverPeriodScoreChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketScores_ServiceDesc is the grpc.ServiceDesc for TicketScores service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketScores_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.TicketScores",
	HandlerType: (*TicketScoresServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAggregatedCategoryScores",
			Handler:    _TicketScores_GetAggregatedCategoryScores_Handler,
		},
		{
			MethodName: "GetScoresByTicket",
			Handler:    _TicketScores_GetScoresByTicket_Handler,
		},
		{
			MethodName: "GetOverallQualityScores",
			Handler:    _TicketScores_GetOverallQualityScores_Handler,
		},
		{
			MethodName: "GetPeriodOverPeriodScoreChange",
			Handler:    _TicketScores_GetPeriodOverPeriodScoreChange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticket_score/ticket_score.proto",
}
