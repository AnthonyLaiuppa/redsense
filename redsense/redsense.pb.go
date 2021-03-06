// Code generated by protoc-gen-go. DO NOT EDIT.
// source: redsense.proto

package redsense

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MediaPage struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Datetime             uint64   `protobuf:"varint,5,opt,name=datetime,proto3" json:"datetime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MediaPage) Reset()         { *m = MediaPage{} }
func (m *MediaPage) String() string { return proto.CompactTextString(m) }
func (*MediaPage) ProtoMessage()    {}
func (*MediaPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_922a1f652597f049, []int{0}
}

func (m *MediaPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaPage.Unmarshal(m, b)
}
func (m *MediaPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaPage.Marshal(b, m, deterministic)
}
func (m *MediaPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaPage.Merge(m, src)
}
func (m *MediaPage) XXX_Size() int {
	return xxx_messageInfo_MediaPage.Size(m)
}
func (m *MediaPage) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaPage.DiscardUnknown(m)
}

var xxx_messageInfo_MediaPage proto.InternalMessageInfo

func (m *MediaPage) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *MediaPage) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *MediaPage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MediaPage) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MediaPage) GetDatetime() uint64 {
	if m != nil {
		return m.Datetime
	}
	return 0
}

type MediaAnalysis struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Datetime             uint64   `protobuf:"varint,5,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Sentiment            string   `protobuf:"bytes,6,opt,name=sentiment,proto3" json:"sentiment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MediaAnalysis) Reset()         { *m = MediaAnalysis{} }
func (m *MediaAnalysis) String() string { return proto.CompactTextString(m) }
func (*MediaAnalysis) ProtoMessage()    {}
func (*MediaAnalysis) Descriptor() ([]byte, []int) {
	return fileDescriptor_922a1f652597f049, []int{1}
}

func (m *MediaAnalysis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MediaAnalysis.Unmarshal(m, b)
}
func (m *MediaAnalysis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MediaAnalysis.Marshal(b, m, deterministic)
}
func (m *MediaAnalysis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MediaAnalysis.Merge(m, src)
}
func (m *MediaAnalysis) XXX_Size() int {
	return xxx_messageInfo_MediaAnalysis.Size(m)
}
func (m *MediaAnalysis) XXX_DiscardUnknown() {
	xxx_messageInfo_MediaAnalysis.DiscardUnknown(m)
}

var xxx_messageInfo_MediaAnalysis proto.InternalMessageInfo

func (m *MediaAnalysis) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *MediaAnalysis) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *MediaAnalysis) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MediaAnalysis) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *MediaAnalysis) GetDatetime() uint64 {
	if m != nil {
		return m.Datetime
	}
	return 0
}

func (m *MediaAnalysis) GetSentiment() string {
	if m != nil {
		return m.Sentiment
	}
	return ""
}

func init() {
	proto.RegisterType((*MediaPage)(nil), "redsense.MediaPage")
	proto.RegisterType((*MediaAnalysis)(nil), "redsense.MediaAnalysis")
}

func init() { proto.RegisterFile("redsense.proto", fileDescriptor_922a1f652597f049) }

var fileDescriptor_922a1f652597f049 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x46, 0x89, 0xfd, 0xa1, 0xbd, 0xe0, 0x0f, 0x57, 0xd1, 0x50, 0x5c, 0x94, 0xae, 0xba, 0xea,
	0x42, 0x37, 0x6e, 0x5d, 0xb9, 0x10, 0x41, 0xea, 0x13, 0xa4, 0xcd, 0x45, 0x02, 0x35, 0x91, 0x24,
	0x15, 0x8a, 0xcf, 0x32, 0xef, 0x3a, 0x34, 0xd3, 0x76, 0x60, 0x1e, 0x60, 0x76, 0xdf, 0x39, 0x10,
	0x92, 0x1c, 0xb8, 0xb2, 0x24, 0x1d, 0x69, 0x47, 0xcd, 0xaf, 0x35, 0xde, 0x60, 0xb6, 0x72, 0xf5,
	0x0f, 0xf9, 0x07, 0x49, 0x25, 0x3e, 0xc5, 0x37, 0xe1, 0x3d, 0xa4, 0xce, 0x8c, 0xb6, 0x27, 0xce,
	0x4a, 0x56, 0xe7, 0xed, 0x42, 0x78, 0x03, 0xd1, 0x68, 0x07, 0x7e, 0x11, 0xe4, 0x3c, 0xf1, 0x0e,
	0x12, 0xaf, 0xfc, 0x40, 0x3c, 0x0a, 0xee, 0x00, 0x88, 0x10, 0x77, 0x46, 0x4e, 0x3c, 0x0e, 0x32,
	0x6c, 0x2c, 0x20, 0x93, 0xc2, 0x93, 0x57, 0x3f, 0xc4, 0x93, 0x92, 0xd5, 0x71, 0xbb, 0x71, 0xb5,
	0x63, 0x70, 0x19, 0x6e, 0x7f, 0xd5, 0x62, 0x98, 0x9c, 0x72, 0xe7, 0x7e, 0x01, 0x3e, 0x42, 0xee,
	0x48, 0xcf, 0x53, 0x7b, 0x9e, 0x86, 0x43, 0x47, 0xf1, 0xf4, 0x0e, 0xd7, 0xed, 0x12, 0xea, 0x8b,
	0xec, 0x9f, 0xea, 0x09, 0x5f, 0x20, 0x7b, 0x23, 0x4d, 0x56, 0x78, 0xc2, 0xdb, 0x66, 0xcb, 0xba,
	0x35, 0x2c, 0x1e, 0x4e, 0xe4, 0xfa, 0xb5, 0x2e, 0x0d, 0xe9, 0x9f, 0xf7, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x06, 0x67, 0x4f, 0xef, 0x8c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RedsenseServiceClient is the client API for RedsenseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RedsenseServiceClient interface {
	Generate(ctx context.Context, in *MediaPage, opts ...grpc.CallOption) (*MediaAnalysis, error)
}

type redsenseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRedsenseServiceClient(cc grpc.ClientConnInterface) RedsenseServiceClient {
	return &redsenseServiceClient{cc}
}

func (c *redsenseServiceClient) Generate(ctx context.Context, in *MediaPage, opts ...grpc.CallOption) (*MediaAnalysis, error) {
	out := new(MediaAnalysis)
	err := c.cc.Invoke(ctx, "/redsense.RedsenseService/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedsenseServiceServer is the server API for RedsenseService service.
type RedsenseServiceServer interface {
	Generate(context.Context, *MediaPage) (*MediaAnalysis, error)
}

// UnimplementedRedsenseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRedsenseServiceServer struct {
}

func (*UnimplementedRedsenseServiceServer) Generate(ctx context.Context, req *MediaPage) (*MediaAnalysis, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}

func RegisterRedsenseServiceServer(s *grpc.Server, srv RedsenseServiceServer) {
	s.RegisterService(&_RedsenseService_serviceDesc, srv)
}

func _RedsenseService_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MediaPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedsenseServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/redsense.RedsenseService/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedsenseServiceServer).Generate(ctx, req.(*MediaPage))
	}
	return interceptor(ctx, in, info, handler)
}

var _RedsenseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "redsense.RedsenseService",
	HandlerType: (*RedsenseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _RedsenseService_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "redsense.proto",
}
