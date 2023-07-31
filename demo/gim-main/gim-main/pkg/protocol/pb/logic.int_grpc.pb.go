// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: logic.int.proto

package pb

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

// LogicIntClient is the client API for LogicInt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogicIntClient interface {
	// 登录
	ConnSignIn(ctx context.Context, in *ConnSignInReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 消息同步
	Sync(ctx context.Context, in *SyncReq, opts ...grpc.CallOption) (*SyncResp, error)
	// 设备收到消息回执
	MessageACK(ctx context.Context, in *MessageACKReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 设备离线
	Offline(ctx context.Context, in *OfflineReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 订阅房间
	SubscribeRoom(ctx context.Context, in *SubscribeRoomReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 推送
	Push(ctx context.Context, in *PushReq, opts ...grpc.CallOption) (*PushResp, error)
	// 推送消息到房间
	PushRoom(ctx context.Context, in *PushRoomReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 全服推送
	PushAll(ctx context.Context, in *PushAllReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 获取设备信息
	GetDevice(ctx context.Context, in *GetDeviceReq, opts ...grpc.CallOption) (*GetDeviceResp, error)
	// 服务停止
	ServerStop(ctx context.Context, in *ServerStopReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type logicIntClient struct {
	cc grpc.ClientConnInterface
}

func NewLogicIntClient(cc grpc.ClientConnInterface) LogicIntClient {
	return &logicIntClient{cc}
}

func (c *logicIntClient) ConnSignIn(ctx context.Context, in *ConnSignInReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicInt/ConnSignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicIntClient) Sync(ctx context.Context, in *SyncReq, opts ...grpc.CallOption) (*SyncResp, error) {
	out := new(SyncResp)
	err := c.cc.Invoke(ctx, "/pb.LogicInt/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicIntClient) MessageACK(ctx context.Context, in *MessageACKReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicInt/MessageACK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicIntClient) Offline(ctx context.Context, in *OfflineReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicInt/Offline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicIntClient) SubscribeRoom(ctx context.Context, in *SubscribeRoomReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicInt/SubscribeRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicIntClient) Push(ctx context.Context, in *PushReq, opts ...grpc.CallOption) (*PushResp, error) {
	out := new(PushResp)
	err := c.cc.Invoke(ctx, "/pb.LogicInt/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicIntClient) PushRoom(ctx context.Context, in *PushRoomReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicInt/PushRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicIntClient) PushAll(ctx context.Context, in *PushAllReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicInt/PushAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicIntClient) GetDevice(ctx context.Context, in *GetDeviceReq, opts ...grpc.CallOption) (*GetDeviceResp, error) {
	out := new(GetDeviceResp)
	err := c.cc.Invoke(ctx, "/pb.LogicInt/GetDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicIntClient) ServerStop(ctx context.Context, in *ServerStopReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicInt/ServerStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicIntServer is the client API for LogicInt service.
// All implementations must embed UnimplementedLogicIntServer
// for forward compatibility
type LogicIntServer interface {
	// 登录
	ConnSignIn(context.Context, *ConnSignInReq) (*emptypb.Empty, error)
	// 消息同步
	Sync(context.Context, *SyncReq) (*SyncResp, error)
	// 设备收到消息回执
	MessageACK(context.Context, *MessageACKReq) (*emptypb.Empty, error)
	// 设备离线
	Offline(context.Context, *OfflineReq) (*emptypb.Empty, error)
	// 订阅房间
	SubscribeRoom(context.Context, *SubscribeRoomReq) (*emptypb.Empty, error)
	// 推送
	Push(context.Context, *PushReq) (*PushResp, error)
	// 推送消息到房间
	PushRoom(context.Context, *PushRoomReq) (*emptypb.Empty, error)
	// 全服推送
	PushAll(context.Context, *PushAllReq) (*emptypb.Empty, error)
	// 获取设备信息
	GetDevice(context.Context, *GetDeviceReq) (*GetDeviceResp, error)
	// 服务停止
	ServerStop(context.Context, *ServerStopReq) (*emptypb.Empty, error)
	mustEmbedUnimplementedLogicIntServer()
}

// UnimplementedLogicIntServer must be embedded to have forward compatible implementations.
type UnimplementedLogicIntServer struct {
}

func (UnimplementedLogicIntServer) ConnSignIn(context.Context, *ConnSignInReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnSignIn not implemented")
}
func (UnimplementedLogicIntServer) Sync(context.Context, *SyncReq) (*SyncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedLogicIntServer) MessageACK(context.Context, *MessageACKReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageACK not implemented")
}
func (UnimplementedLogicIntServer) Offline(context.Context, *OfflineReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Offline not implemented")
}
func (UnimplementedLogicIntServer) SubscribeRoom(context.Context, *SubscribeRoomReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeRoom not implemented")
}
func (UnimplementedLogicIntServer) Push(context.Context, *PushReq) (*PushResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedLogicIntServer) PushRoom(context.Context, *PushRoomReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushRoom not implemented")
}
func (UnimplementedLogicIntServer) PushAll(context.Context, *PushAllReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushAll not implemented")
}
func (UnimplementedLogicIntServer) GetDevice(context.Context, *GetDeviceReq) (*GetDeviceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevice not implemented")
}
func (UnimplementedLogicIntServer) ServerStop(context.Context, *ServerStopReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerStop not implemented")
}
func (UnimplementedLogicIntServer) mustEmbedUnimplementedLogicIntServer() {}

// UnsafeLogicIntServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogicIntServer will
// result in compilation errors.
type UnsafeLogicIntServer interface {
	mustEmbedUnimplementedLogicIntServer()
}

func RegisterLogicIntServer(s grpc.ServiceRegistrar, srv LogicIntServer) {
	s.RegisterService(&LogicInt_ServiceDesc, srv)
}

func _LogicInt_ConnSignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnSignInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicIntServer).ConnSignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicInt/ConnSignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicIntServer).ConnSignIn(ctx, req.(*ConnSignInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicInt_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicIntServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicInt/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicIntServer).Sync(ctx, req.(*SyncReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicInt_MessageACK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageACKReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicIntServer).MessageACK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicInt/MessageACK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicIntServer).MessageACK(ctx, req.(*MessageACKReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicInt_Offline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfflineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicIntServer).Offline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicInt/Offline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicIntServer).Offline(ctx, req.(*OfflineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicInt_SubscribeRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicIntServer).SubscribeRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicInt/SubscribeRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicIntServer).SubscribeRoom(ctx, req.(*SubscribeRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicInt_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicIntServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicInt/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicIntServer).Push(ctx, req.(*PushReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicInt_PushRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicIntServer).PushRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicInt/PushRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicIntServer).PushRoom(ctx, req.(*PushRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicInt_PushAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicIntServer).PushAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicInt/PushAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicIntServer).PushAll(ctx, req.(*PushAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicInt_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicIntServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicInt/GetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicIntServer).GetDevice(ctx, req.(*GetDeviceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicInt_ServerStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerStopReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicIntServer).ServerStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicInt/ServerStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicIntServer).ServerStop(ctx, req.(*ServerStopReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LogicInt_ServiceDesc is the grpc.ServiceDesc for LogicInt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogicInt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LogicInt",
	HandlerType: (*LogicIntServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConnSignIn",
			Handler:    _LogicInt_ConnSignIn_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _LogicInt_Sync_Handler,
		},
		{
			MethodName: "MessageACK",
			Handler:    _LogicInt_MessageACK_Handler,
		},
		{
			MethodName: "Offline",
			Handler:    _LogicInt_Offline_Handler,
		},
		{
			MethodName: "SubscribeRoom",
			Handler:    _LogicInt_SubscribeRoom_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _LogicInt_Push_Handler,
		},
		{
			MethodName: "PushRoom",
			Handler:    _LogicInt_PushRoom_Handler,
		},
		{
			MethodName: "PushAll",
			Handler:    _LogicInt_PushAll_Handler,
		},
		{
			MethodName: "GetDevice",
			Handler:    _LogicInt_GetDevice_Handler,
		},
		{
			MethodName: "ServerStop",
			Handler:    _LogicInt_ServerStop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic.int.proto",
}
