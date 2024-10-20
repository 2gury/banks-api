// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: banks.proto

package banks

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
	Banks_GetBanks_FullMethodName                     = "/worker.Banks/GetBanks"
	Banks_UpdateBank_FullMethodName                   = "/worker.Banks/UpdateBank"
	Banks_GetPossibleBanks_FullMethodName             = "/worker.Banks/GetPossibleBanks"
	Banks_RequestBankInformation_FullMethodName       = "/worker.Banks/RequestBankInformation"
	Banks_RequestTranslationText_FullMethodName       = "/worker.Banks/RequestTranslationText"
	Banks_GetTranslationText_FullMethodName           = "/worker.Banks/GetTranslationText"
	Banks_CreateTranslationText_FullMethodName        = "/worker.Banks/CreateTranslationText"
	Banks_GetTranslations_FullMethodName              = "/worker.Banks/GetTranslations"
	Banks_GetReviews_FullMethodName                   = "/worker.Banks/GetReviews"
	Banks_CreateReview_FullMethodName                 = "/worker.Banks/CreateReview"
	Banks_UpdateReview_FullMethodName                 = "/worker.Banks/UpdateReview"
	Banks_UpdateAutomoderationStrategy_FullMethodName = "/worker.Banks/UpdateAutomoderationStrategy"
)

// BanksClient is the client API for Banks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BanksClient interface {
	GetBanks(ctx context.Context, in *GetBanksRequest, opts ...grpc.CallOption) (*GetBanksResponse, error)
	UpdateBank(ctx context.Context, in *UpdateBankRequest, opts ...grpc.CallOption) (*UpdateBankResponse, error)
	GetPossibleBanks(ctx context.Context, in *GetPossibleBanksRequest, opts ...grpc.CallOption) (*GetPossibleBanksResponse, error)
	RequestBankInformation(ctx context.Context, in *RequestBankInformationRequest, opts ...grpc.CallOption) (*RequestBankInformationResponse, error)
	RequestTranslationText(ctx context.Context, in *RequestTranslationTextRequest, opts ...grpc.CallOption) (*RequestTranslationTextResponse, error)
	GetTranslationText(ctx context.Context, in *GetTranslationTextRequest, opts ...grpc.CallOption) (*GetTranslationTextResponse, error)
	CreateTranslationText(ctx context.Context, in *CreateTranslationTextRequest, opts ...grpc.CallOption) (*CreateTranslationTextResponse, error)
	GetTranslations(ctx context.Context, in *GetTranslationsRequest, opts ...grpc.CallOption) (*GetTranslationsResponse, error)
	GetReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetReviewsResponse, error)
	CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*CreateReviewResponse, error)
	UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewResponse, error)
	UpdateAutomoderationStrategy(ctx context.Context, in *UpdateAutomoderationStrategyRequest, opts ...grpc.CallOption) (*UpdateAutomoderationStrategyResponse, error)
}

type banksClient struct {
	cc grpc.ClientConnInterface
}

func NewBanksClient(cc grpc.ClientConnInterface) BanksClient {
	return &banksClient{cc}
}

func (c *banksClient) GetBanks(ctx context.Context, in *GetBanksRequest, opts ...grpc.CallOption) (*GetBanksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBanksResponse)
	err := c.cc.Invoke(ctx, Banks_GetBanks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) UpdateBank(ctx context.Context, in *UpdateBankRequest, opts ...grpc.CallOption) (*UpdateBankResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBankResponse)
	err := c.cc.Invoke(ctx, Banks_UpdateBank_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) GetPossibleBanks(ctx context.Context, in *GetPossibleBanksRequest, opts ...grpc.CallOption) (*GetPossibleBanksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPossibleBanksResponse)
	err := c.cc.Invoke(ctx, Banks_GetPossibleBanks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) RequestBankInformation(ctx context.Context, in *RequestBankInformationRequest, opts ...grpc.CallOption) (*RequestBankInformationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestBankInformationResponse)
	err := c.cc.Invoke(ctx, Banks_RequestBankInformation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) RequestTranslationText(ctx context.Context, in *RequestTranslationTextRequest, opts ...grpc.CallOption) (*RequestTranslationTextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RequestTranslationTextResponse)
	err := c.cc.Invoke(ctx, Banks_RequestTranslationText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) GetTranslationText(ctx context.Context, in *GetTranslationTextRequest, opts ...grpc.CallOption) (*GetTranslationTextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTranslationTextResponse)
	err := c.cc.Invoke(ctx, Banks_GetTranslationText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) CreateTranslationText(ctx context.Context, in *CreateTranslationTextRequest, opts ...grpc.CallOption) (*CreateTranslationTextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTranslationTextResponse)
	err := c.cc.Invoke(ctx, Banks_CreateTranslationText_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) GetTranslations(ctx context.Context, in *GetTranslationsRequest, opts ...grpc.CallOption) (*GetTranslationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTranslationsResponse)
	err := c.cc.Invoke(ctx, Banks_GetTranslations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) GetReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetReviewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReviewsResponse)
	err := c.cc.Invoke(ctx, Banks_GetReviews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) CreateReview(ctx context.Context, in *CreateReviewRequest, opts ...grpc.CallOption) (*CreateReviewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateReviewResponse)
	err := c.cc.Invoke(ctx, Banks_CreateReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) UpdateReview(ctx context.Context, in *UpdateReviewRequest, opts ...grpc.CallOption) (*UpdateReviewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateReviewResponse)
	err := c.cc.Invoke(ctx, Banks_UpdateReview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *banksClient) UpdateAutomoderationStrategy(ctx context.Context, in *UpdateAutomoderationStrategyRequest, opts ...grpc.CallOption) (*UpdateAutomoderationStrategyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAutomoderationStrategyResponse)
	err := c.cc.Invoke(ctx, Banks_UpdateAutomoderationStrategy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BanksServer is the server API for Banks service.
// All implementations must embed UnimplementedBanksServer
// for forward compatibility.
type BanksServer interface {
	GetBanks(context.Context, *GetBanksRequest) (*GetBanksResponse, error)
	UpdateBank(context.Context, *UpdateBankRequest) (*UpdateBankResponse, error)
	GetPossibleBanks(context.Context, *GetPossibleBanksRequest) (*GetPossibleBanksResponse, error)
	RequestBankInformation(context.Context, *RequestBankInformationRequest) (*RequestBankInformationResponse, error)
	RequestTranslationText(context.Context, *RequestTranslationTextRequest) (*RequestTranslationTextResponse, error)
	GetTranslationText(context.Context, *GetTranslationTextRequest) (*GetTranslationTextResponse, error)
	CreateTranslationText(context.Context, *CreateTranslationTextRequest) (*CreateTranslationTextResponse, error)
	GetTranslations(context.Context, *GetTranslationsRequest) (*GetTranslationsResponse, error)
	GetReviews(context.Context, *GetReviewsRequest) (*GetReviewsResponse, error)
	CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewResponse, error)
	UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewResponse, error)
	UpdateAutomoderationStrategy(context.Context, *UpdateAutomoderationStrategyRequest) (*UpdateAutomoderationStrategyResponse, error)
	mustEmbedUnimplementedBanksServer()
}

// UnimplementedBanksServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBanksServer struct{}

func (UnimplementedBanksServer) GetBanks(context.Context, *GetBanksRequest) (*GetBanksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBanks not implemented")
}
func (UnimplementedBanksServer) UpdateBank(context.Context, *UpdateBankRequest) (*UpdateBankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBank not implemented")
}
func (UnimplementedBanksServer) GetPossibleBanks(context.Context, *GetPossibleBanksRequest) (*GetPossibleBanksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPossibleBanks not implemented")
}
func (UnimplementedBanksServer) RequestBankInformation(context.Context, *RequestBankInformationRequest) (*RequestBankInformationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestBankInformation not implemented")
}
func (UnimplementedBanksServer) RequestTranslationText(context.Context, *RequestTranslationTextRequest) (*RequestTranslationTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestTranslationText not implemented")
}
func (UnimplementedBanksServer) GetTranslationText(context.Context, *GetTranslationTextRequest) (*GetTranslationTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTranslationText not implemented")
}
func (UnimplementedBanksServer) CreateTranslationText(context.Context, *CreateTranslationTextRequest) (*CreateTranslationTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTranslationText not implemented")
}
func (UnimplementedBanksServer) GetTranslations(context.Context, *GetTranslationsRequest) (*GetTranslationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTranslations not implemented")
}
func (UnimplementedBanksServer) GetReviews(context.Context, *GetReviewsRequest) (*GetReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReviews not implemented")
}
func (UnimplementedBanksServer) CreateReview(context.Context, *CreateReviewRequest) (*CreateReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReview not implemented")
}
func (UnimplementedBanksServer) UpdateReview(context.Context, *UpdateReviewRequest) (*UpdateReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReview not implemented")
}
func (UnimplementedBanksServer) UpdateAutomoderationStrategy(context.Context, *UpdateAutomoderationStrategyRequest) (*UpdateAutomoderationStrategyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAutomoderationStrategy not implemented")
}
func (UnimplementedBanksServer) mustEmbedUnimplementedBanksServer() {}
func (UnimplementedBanksServer) testEmbeddedByValue()               {}

// UnsafeBanksServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BanksServer will
// result in compilation errors.
type UnsafeBanksServer interface {
	mustEmbedUnimplementedBanksServer()
}

func RegisterBanksServer(s grpc.ServiceRegistrar, srv BanksServer) {
	// If the following call pancis, it indicates UnimplementedBanksServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Banks_ServiceDesc, srv)
}

func _Banks_GetBanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBanksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).GetBanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_GetBanks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).GetBanks(ctx, req.(*GetBanksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_UpdateBank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).UpdateBank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_UpdateBank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).UpdateBank(ctx, req.(*UpdateBankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_GetPossibleBanks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPossibleBanksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).GetPossibleBanks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_GetPossibleBanks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).GetPossibleBanks(ctx, req.(*GetPossibleBanksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_RequestBankInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestBankInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).RequestBankInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_RequestBankInformation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).RequestBankInformation(ctx, req.(*RequestBankInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_RequestTranslationText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestTranslationTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).RequestTranslationText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_RequestTranslationText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).RequestTranslationText(ctx, req.(*RequestTranslationTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_GetTranslationText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTranslationTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).GetTranslationText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_GetTranslationText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).GetTranslationText(ctx, req.(*GetTranslationTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_CreateTranslationText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTranslationTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).CreateTranslationText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_CreateTranslationText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).CreateTranslationText(ctx, req.(*CreateTranslationTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_GetTranslations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTranslationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).GetTranslations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_GetTranslations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).GetTranslations(ctx, req.(*GetTranslationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_GetReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).GetReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_GetReviews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).GetReviews(ctx, req.(*GetReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_CreateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).CreateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_CreateReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).CreateReview(ctx, req.(*CreateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_UpdateReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).UpdateReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_UpdateReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).UpdateReview(ctx, req.(*UpdateReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banks_UpdateAutomoderationStrategy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAutomoderationStrategyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BanksServer).UpdateAutomoderationStrategy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Banks_UpdateAutomoderationStrategy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BanksServer).UpdateAutomoderationStrategy(ctx, req.(*UpdateAutomoderationStrategyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Banks_ServiceDesc is the grpc.ServiceDesc for Banks service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Banks_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "worker.Banks",
	HandlerType: (*BanksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBanks",
			Handler:    _Banks_GetBanks_Handler,
		},
		{
			MethodName: "UpdateBank",
			Handler:    _Banks_UpdateBank_Handler,
		},
		{
			MethodName: "GetPossibleBanks",
			Handler:    _Banks_GetPossibleBanks_Handler,
		},
		{
			MethodName: "RequestBankInformation",
			Handler:    _Banks_RequestBankInformation_Handler,
		},
		{
			MethodName: "RequestTranslationText",
			Handler:    _Banks_RequestTranslationText_Handler,
		},
		{
			MethodName: "GetTranslationText",
			Handler:    _Banks_GetTranslationText_Handler,
		},
		{
			MethodName: "CreateTranslationText",
			Handler:    _Banks_CreateTranslationText_Handler,
		},
		{
			MethodName: "GetTranslations",
			Handler:    _Banks_GetTranslations_Handler,
		},
		{
			MethodName: "GetReviews",
			Handler:    _Banks_GetReviews_Handler,
		},
		{
			MethodName: "CreateReview",
			Handler:    _Banks_CreateReview_Handler,
		},
		{
			MethodName: "UpdateReview",
			Handler:    _Banks_UpdateReview_Handler,
		},
		{
			MethodName: "UpdateAutomoderationStrategy",
			Handler:    _Banks_UpdateAutomoderationStrategy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "banks.proto",
}
