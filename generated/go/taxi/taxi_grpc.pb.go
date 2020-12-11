// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package taxi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TaxiClient is the client API for Taxi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaxiClient interface {
	// The endpoint Topup it's called by the user without account and
	// are willing to pay with Lightning Network.
	//
	// The process will returs a BOLT11 Lightning Invoice to be paid
	//
	// Once the LN invoice has been paid by the user the PSET with
	// added LBTC fees could be found in the GetTopupResponse
	Topup(ctx context.Context, in *TopupRequest, opts ...grpc.CallOption) (*TopupReply, error)
	// The status of the topup can be fetched via GetTopup rpc.
	GetTopup(ctx context.Context, in *GetTopupRequest, opts ...grpc.CallOption) (*GetTopupReply, error)
	// Recurrent users could signup for an account and obtains an
	// API KEY that allows to programmatically invoke topups.
	TopupWithKey(ctx context.Context, in *TopupWithKeyRequest, opts ...grpc.CallOption) (*TopupWithKeyReply, error)
	// A subset of supported stablecoins could be accepted as payment
	// for topups. Before calling the actual TopupWithAsset the user could
	// call the GetAssetEstimate rpc
	GetAssetEstimate(ctx context.Context, in *GetAssetEstimateRequest, opts ...grpc.CallOption) (*GetAssetEstimateReply, error)
	TopupWithAsset(ctx context.Context, in *TopupWithAssetRequest, opts ...grpc.CallOption) (*TopupWithAssetReply, error)
}

type taxiClient struct {
	cc grpc.ClientConnInterface
}

func NewTaxiClient(cc grpc.ClientConnInterface) TaxiClient {
	return &taxiClient{cc}
}

func (c *taxiClient) Topup(ctx context.Context, in *TopupRequest, opts ...grpc.CallOption) (*TopupReply, error) {
	out := new(TopupReply)
	err := c.cc.Invoke(ctx, "/Taxi/Topup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxiClient) GetTopup(ctx context.Context, in *GetTopupRequest, opts ...grpc.CallOption) (*GetTopupReply, error) {
	out := new(GetTopupReply)
	err := c.cc.Invoke(ctx, "/Taxi/GetTopup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxiClient) TopupWithKey(ctx context.Context, in *TopupWithKeyRequest, opts ...grpc.CallOption) (*TopupWithKeyReply, error) {
	out := new(TopupWithKeyReply)
	err := c.cc.Invoke(ctx, "/Taxi/TopupWithKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxiClient) GetAssetEstimate(ctx context.Context, in *GetAssetEstimateRequest, opts ...grpc.CallOption) (*GetAssetEstimateReply, error) {
	out := new(GetAssetEstimateReply)
	err := c.cc.Invoke(ctx, "/Taxi/GetAssetEstimate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxiClient) TopupWithAsset(ctx context.Context, in *TopupWithAssetRequest, opts ...grpc.CallOption) (*TopupWithAssetReply, error) {
	out := new(TopupWithAssetReply)
	err := c.cc.Invoke(ctx, "/Taxi/TopupWithAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaxiServer is the server API for Taxi service.
// All implementations must embed UnimplementedTaxiServer
// for forward compatibility
type TaxiServer interface {
	// The endpoint Topup it's called by the user without account and
	// are willing to pay with Lightning Network.
	//
	// The process will returs a BOLT11 Lightning Invoice to be paid
	//
	// Once the LN invoice has been paid by the user the PSET with
	// added LBTC fees could be found in the GetTopupResponse
	Topup(context.Context, *TopupRequest) (*TopupReply, error)
	// The status of the topup can be fetched via GetTopup rpc.
	GetTopup(context.Context, *GetTopupRequest) (*GetTopupReply, error)
	// Recurrent users could signup for an account and obtains an
	// API KEY that allows to programmatically invoke topups.
	TopupWithKey(context.Context, *TopupWithKeyRequest) (*TopupWithKeyReply, error)
	// A subset of supported stablecoins could be accepted as payment
	// for topups. Before calling the actual TopupWithAsset the user could
	// call the GetAssetEstimate rpc
	GetAssetEstimate(context.Context, *GetAssetEstimateRequest) (*GetAssetEstimateReply, error)
	TopupWithAsset(context.Context, *TopupWithAssetRequest) (*TopupWithAssetReply, error)
	mustEmbedUnimplementedTaxiServer()
}

// UnimplementedTaxiServer must be embedded to have forward compatible implementations.
type UnimplementedTaxiServer struct {
}

func (*UnimplementedTaxiServer) Topup(context.Context, *TopupRequest) (*TopupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Topup not implemented")
}
func (*UnimplementedTaxiServer) GetTopup(context.Context, *GetTopupRequest) (*GetTopupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopup not implemented")
}
func (*UnimplementedTaxiServer) TopupWithKey(context.Context, *TopupWithKeyRequest) (*TopupWithKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopupWithKey not implemented")
}
func (*UnimplementedTaxiServer) GetAssetEstimate(context.Context, *GetAssetEstimateRequest) (*GetAssetEstimateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetEstimate not implemented")
}
func (*UnimplementedTaxiServer) TopupWithAsset(context.Context, *TopupWithAssetRequest) (*TopupWithAssetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TopupWithAsset not implemented")
}
func (*UnimplementedTaxiServer) mustEmbedUnimplementedTaxiServer() {}

func RegisterTaxiServer(s *grpc.Server, srv TaxiServer) {
	s.RegisterService(&_Taxi_serviceDesc, srv)
}

func _Taxi_Topup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxiServer).Topup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Taxi/Topup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxiServer).Topup(ctx, req.(*TopupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taxi_GetTopup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxiServer).GetTopup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Taxi/GetTopup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxiServer).GetTopup(ctx, req.(*GetTopupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taxi_TopupWithKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopupWithKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxiServer).TopupWithKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Taxi/TopupWithKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxiServer).TopupWithKey(ctx, req.(*TopupWithKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taxi_GetAssetEstimate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssetEstimateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxiServer).GetAssetEstimate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Taxi/GetAssetEstimate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxiServer).GetAssetEstimate(ctx, req.(*GetAssetEstimateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Taxi_TopupWithAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopupWithAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxiServer).TopupWithAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Taxi/TopupWithAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxiServer).TopupWithAsset(ctx, req.(*TopupWithAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Taxi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Taxi",
	HandlerType: (*TaxiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Topup",
			Handler:    _Taxi_Topup_Handler,
		},
		{
			MethodName: "GetTopup",
			Handler:    _Taxi_GetTopup_Handler,
		},
		{
			MethodName: "TopupWithKey",
			Handler:    _Taxi_TopupWithKey_Handler,
		},
		{
			MethodName: "GetAssetEstimate",
			Handler:    _Taxi_GetAssetEstimate_Handler,
		},
		{
			MethodName: "TopupWithAsset",
			Handler:    _Taxi_TopupWithAsset_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "taxi.proto",
}
