// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.6
// source: grpc/proto.proto

package proto

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
	AuctionService_Bid_FullMethodName            = "/AuctionService/Bid"
	AuctionService_Result_FullMethodName         = "/AuctionService/Result"
	AuctionService_UpdateBid_FullMethodName      = "/AuctionService/UpdateBid"
	AuctionService_CheckPulse_FullMethodName     = "/AuctionService/CheckPulse"
	AuctionService_AddNode_FullMethodName        = "/AuctionService/AddNode"
	AuctionService_Join_FullMethodName           = "/AuctionService/Join"
	AuctionService_GetLogicalTime_FullMethodName = "/AuctionService/GetLogicalTime"
	AuctionService_SetLeader_FullMethodName      = "/AuctionService/SetLeader"
	AuctionService_RunElection_FullMethodName    = "/AuctionService/RunElection"
	AuctionService_RemoveNode_FullMethodName     = "/AuctionService/RemoveNode"
	AuctionService_StartAuction_FullMethodName   = "/AuctionService/StartAuction"
)

// AuctionServiceClient is the client API for AuctionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionServiceClient interface {
	Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*Reply, error)
	Result(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResultResponse, error)
	UpdateBid(ctx context.Context, in *BidUpdateMessage, opts ...grpc.CallOption) (*Reply, error)
	CheckPulse(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reply, error)
	AddNode(ctx context.Context, in *JoinMessage, opts ...grpc.CallOption) (*JoinMessage, error)
	Join(ctx context.Context, in *JoinMessage, opts ...grpc.CallOption) (*JoinResponse, error)
	GetLogicalTime(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TimeMessage, error)
	SetLeader(ctx context.Context, in *PortMessage, opts ...grpc.CallOption) (*Empty, error)
	RunElection(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	RemoveNode(ctx context.Context, in *PortMessage, opts ...grpc.CallOption) (*Empty, error)
	StartAuction(ctx context.Context, in *TimeMessage, opts ...grpc.CallOption) (*Empty, error)
}

type auctionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionServiceClient(cc grpc.ClientConnInterface) AuctionServiceClient {
	return &auctionServiceClient{cc}
}

func (c *auctionServiceClient) Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*Reply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reply)
	err := c.cc.Invoke(ctx, AuctionService_Bid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) Result(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ResultResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResultResponse)
	err := c.cc.Invoke(ctx, AuctionService_Result_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) UpdateBid(ctx context.Context, in *BidUpdateMessage, opts ...grpc.CallOption) (*Reply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reply)
	err := c.cc.Invoke(ctx, AuctionService_UpdateBid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) CheckPulse(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Reply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Reply)
	err := c.cc.Invoke(ctx, AuctionService_CheckPulse_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) AddNode(ctx context.Context, in *JoinMessage, opts ...grpc.CallOption) (*JoinMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinMessage)
	err := c.cc.Invoke(ctx, AuctionService_AddNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) Join(ctx context.Context, in *JoinMessage, opts ...grpc.CallOption) (*JoinResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, AuctionService_Join_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) GetLogicalTime(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TimeMessage, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TimeMessage)
	err := c.cc.Invoke(ctx, AuctionService_GetLogicalTime_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) SetLeader(ctx context.Context, in *PortMessage, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AuctionService_SetLeader_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) RunElection(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AuctionService_RunElection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) RemoveNode(ctx context.Context, in *PortMessage, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AuctionService_RemoveNode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) StartAuction(ctx context.Context, in *TimeMessage, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AuctionService_StartAuction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServiceServer is the server API for AuctionService service.
// All implementations must embed UnimplementedAuctionServiceServer
// for forward compatibility.
type AuctionServiceServer interface {
	Bid(context.Context, *BidMessage) (*Reply, error)
	Result(context.Context, *Empty) (*ResultResponse, error)
	UpdateBid(context.Context, *BidUpdateMessage) (*Reply, error)
	CheckPulse(context.Context, *Empty) (*Reply, error)
	AddNode(context.Context, *JoinMessage) (*JoinMessage, error)
	Join(context.Context, *JoinMessage) (*JoinResponse, error)
	GetLogicalTime(context.Context, *Empty) (*TimeMessage, error)
	SetLeader(context.Context, *PortMessage) (*Empty, error)
	RunElection(context.Context, *Empty) (*Empty, error)
	RemoveNode(context.Context, *PortMessage) (*Empty, error)
	StartAuction(context.Context, *TimeMessage) (*Empty, error)
	mustEmbedUnimplementedAuctionServiceServer()
}

// UnimplementedAuctionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuctionServiceServer struct{}

func (UnimplementedAuctionServiceServer) Bid(context.Context, *BidMessage) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedAuctionServiceServer) Result(context.Context, *Empty) (*ResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedAuctionServiceServer) UpdateBid(context.Context, *BidUpdateMessage) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBid not implemented")
}
func (UnimplementedAuctionServiceServer) CheckPulse(context.Context, *Empty) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPulse not implemented")
}
func (UnimplementedAuctionServiceServer) AddNode(context.Context, *JoinMessage) (*JoinMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNode not implemented")
}
func (UnimplementedAuctionServiceServer) Join(context.Context, *JoinMessage) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedAuctionServiceServer) GetLogicalTime(context.Context, *Empty) (*TimeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogicalTime not implemented")
}
func (UnimplementedAuctionServiceServer) SetLeader(context.Context, *PortMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLeader not implemented")
}
func (UnimplementedAuctionServiceServer) RunElection(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunElection not implemented")
}
func (UnimplementedAuctionServiceServer) RemoveNode(context.Context, *PortMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveNode not implemented")
}
func (UnimplementedAuctionServiceServer) StartAuction(context.Context, *TimeMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartAuction not implemented")
}
func (UnimplementedAuctionServiceServer) mustEmbedUnimplementedAuctionServiceServer() {}
func (UnimplementedAuctionServiceServer) testEmbeddedByValue()                        {}

// UnsafeAuctionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServiceServer will
// result in compilation errors.
type UnsafeAuctionServiceServer interface {
	mustEmbedUnimplementedAuctionServiceServer()
}

func RegisterAuctionServiceServer(s grpc.ServiceRegistrar, srv AuctionServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuctionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuctionService_ServiceDesc, srv)
}

func _AuctionService_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_Bid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).Bid(ctx, req.(*BidMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_Result_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).Result(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_UpdateBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidUpdateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).UpdateBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_UpdateBid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).UpdateBid(ctx, req.(*BidUpdateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_CheckPulse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).CheckPulse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_CheckPulse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).CheckPulse(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_AddNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).AddNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_AddNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).AddNode(ctx, req.(*JoinMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_Join_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).Join(ctx, req.(*JoinMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_GetLogicalTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).GetLogicalTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_GetLogicalTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).GetLogicalTime(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_SetLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).SetLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_SetLeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).SetLeader(ctx, req.(*PortMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_RunElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).RunElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_RunElection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).RunElection(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_RemoveNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).RemoveNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_RemoveNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).RemoveNode(ctx, req.(*PortMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_StartAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimeMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).StartAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_StartAuction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).StartAuction(ctx, req.(*TimeMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// AuctionService_ServiceDesc is the grpc.ServiceDesc for AuctionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuctionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuctionService",
	HandlerType: (*AuctionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bid",
			Handler:    _AuctionService_Bid_Handler,
		},
		{
			MethodName: "Result",
			Handler:    _AuctionService_Result_Handler,
		},
		{
			MethodName: "UpdateBid",
			Handler:    _AuctionService_UpdateBid_Handler,
		},
		{
			MethodName: "CheckPulse",
			Handler:    _AuctionService_CheckPulse_Handler,
		},
		{
			MethodName: "AddNode",
			Handler:    _AuctionService_AddNode_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _AuctionService_Join_Handler,
		},
		{
			MethodName: "GetLogicalTime",
			Handler:    _AuctionService_GetLogicalTime_Handler,
		},
		{
			MethodName: "SetLeader",
			Handler:    _AuctionService_SetLeader_Handler,
		},
		{
			MethodName: "RunElection",
			Handler:    _AuctionService_RunElection_Handler,
		},
		{
			MethodName: "RemoveNode",
			Handler:    _AuctionService_RemoveNode_Handler,
		},
		{
			MethodName: "StartAuction",
			Handler:    _AuctionService_StartAuction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}
