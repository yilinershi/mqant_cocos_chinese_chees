// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: http/login.proto

package pb_login

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	pb_enum "server/pb/common/pb_enum"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//登录
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID     string  `protobuf:"bytes,1,opt,name=AppID,proto3" json:"AppID,omitempty"`         //用户来自于哪一个应用
	ChannelID string  `protobuf:"bytes,2,opt,name=ChannelID,proto3" json:"ChannelID,omitempty"` //用户来自于哪一个渠道
	IMEI      string  `protobuf:"bytes,3,opt,name=IMEI,proto3" json:"IMEI,omitempty"`
	Device    *Device `protobuf:"bytes,4,opt,name=Device,proto3" json:"Device,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_http_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_http_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *LoginReq) GetChannelID() string {
	if x != nil {
		return x.ChannelID
	}
	return ""
}

func (x *LoginReq) GetIMEI() string {
	if x != nil {
		return x.IMEI
	}
	return ""
}

func (x *LoginReq) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OS       string `protobuf:"bytes,1,opt,name=OS,proto3" json:"OS,omitempty"`             // os版本号
	Model    string `protobuf:"bytes,2,opt,name=Model,proto3" json:"Model,omitempty"`       //硬件型号
	IP       string `protobuf:"bytes,3,opt,name=IP,proto3" json:"IP,omitempty"`             //内网IP
	RemoteIP string `protobuf:"bytes,4,opt,name=RemoteIP,proto3" json:"RemoteIP,omitempty"` //外网IP
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_http_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_http_login_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetOS() string {
	if x != nil {
		return x.OS
	}
	return ""
}

func (x *Device) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Device) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *Device) GetRemoteIP() string {
	if x != nil {
		return x.RemoteIP
	}
	return ""
}

type LoginResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    pb_enum.ErrorCode `protobuf:"varint,1,opt,name=Code,proto3,enum=pb_enum.ErrorCode" json:"Code,omitempty"`
	Name    string            `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Uid     int64             `protobuf:"varint,3,opt,name=Uid,proto3" json:"Uid,omitempty"`
	HeadUrl string            `protobuf:"bytes,4,opt,name=HeadUrl,proto3" json:"HeadUrl,omitempty"`
	Score   int64             `protobuf:"varint,5,opt,name=Score,proto3" json:"Score,omitempty"`
	Sex     pb_enum.Sex       `protobuf:"varint,6,opt,name=Sex,proto3,enum=pb_enum.Sex" json:"Sex,omitempty"`
	IP      string            `protobuf:"bytes,7,opt,name=IP,proto3" json:"IP,omitempty"`
	Port    int64             `protobuf:"varint,8,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_http_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_http_login_proto_rawDescGZIP(), []int{2}
}

func (x *LoginResp) GetCode() pb_enum.ErrorCode {
	if x != nil {
		return x.Code
	}
	return pb_enum.ErrorCode_Default
}

func (x *LoginResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginResp) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LoginResp) GetHeadUrl() string {
	if x != nil {
		return x.HeadUrl
	}
	return ""
}

func (x *LoginResp) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *LoginResp) GetSex() pb_enum.Sex {
	if x != nil {
		return x.Sex
	}
	return pb_enum.Sex_Unknow
}

func (x *LoginResp) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *LoginResp) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_http_login_proto protoreflect.FileDescriptor

var file_http_login_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x70, 0x62, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x11, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7c, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x49, 0x4d, 0x45, 0x49, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49,
	0x4d, 0x45, 0x49, 0x12, 0x28, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x5a, 0x0a,
	0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x53, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x4f, 0x53, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x50, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x1a, 0x0a,
	0x08, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x50, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x50, 0x22, 0xcd, 0x01, 0x0a, 0x09, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x26, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x55, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x65, 0x61, 0x64, 0x55, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x48, 0x65, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x53, 0x65, 0x78,
	0x52, 0x03, 0x53, 0x65, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x32, 0x41, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x38, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x70,
	0x62, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_http_login_proto_rawDescOnce sync.Once
	file_http_login_proto_rawDescData = file_http_login_proto_rawDesc
)

func file_http_login_proto_rawDescGZIP() []byte {
	file_http_login_proto_rawDescOnce.Do(func() {
		file_http_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_login_proto_rawDescData)
	})
	return file_http_login_proto_rawDescData
}

var file_http_login_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_http_login_proto_goTypes = []interface{}{
	(*LoginReq)(nil),       // 0: pb_login.LoginReq
	(*Device)(nil),         // 1: pb_login.Device
	(*LoginResp)(nil),      // 2: pb_login.LoginResp
	(pb_enum.ErrorCode)(0), // 3: pb_enum.ErrorCode
	(pb_enum.Sex)(0),       // 4: pb_enum.Sex
}
var file_http_login_proto_depIdxs = []int32{
	1, // 0: pb_login.LoginReq.Device:type_name -> pb_login.Device
	3, // 1: pb_login.LoginResp.Code:type_name -> pb_enum.ErrorCode
	4, // 2: pb_login.LoginResp.Sex:type_name -> pb_enum.Sex
	0, // 3: pb_login.Login.LoginHandle:input_type -> pb_login.LoginReq
	2, // 4: pb_login.Login.LoginHandle:output_type -> pb_login.LoginResp
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_http_login_proto_init() }
func file_http_login_proto_init() {
	if File_http_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_http_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_http_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_http_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_http_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_http_login_proto_goTypes,
		DependencyIndexes: file_http_login_proto_depIdxs,
		MessageInfos:      file_http_login_proto_msgTypes,
	}.Build()
	File_http_login_proto = out.File
	file_http_login_proto_rawDesc = nil
	file_http_login_proto_goTypes = nil
	file_http_login_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginClient interface {
	LoginHandle(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) LoginHandle(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/pb_login.Login/LoginHandle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
type LoginServer interface {
	LoginHandle(context.Context, *LoginReq) (*LoginResp, error)
}

// UnimplementedLoginServer can be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (*UnimplementedLoginServer) LoginHandle(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginHandle not implemented")
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_LoginHandle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).LoginHandle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_login.Login/LoginHandle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).LoginHandle(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb_login.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoginHandle",
			Handler:    _Login_LoginHandle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "http/login.proto",
}