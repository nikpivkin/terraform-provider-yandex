// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/containerregistry/v1/repository_service.proto

package containerregistry

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListRepositoriesRequest struct {
	// ID of the registry to list repositories in.
	//
	// To get the registry ID use a [RegistryService.List] request.
	RegistryId string `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	// ID of the folder to list registries in.
	//
	// [folder_id] is ignored if a [ListImagesRequest.registry_id] is specified in the request.
	//
	// To get the folder ID use a [yandex.cloud.resourcemanager.v1.FolderService.List] request.
	FolderId string `protobuf:"bytes,6,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListRepositoriesResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the
	// [ListRepositoriesResponse.next_page_token] returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter expression that filters resources listed in the response.
	// The expression must specify:
	// 1. The field name. Currently you can use filtering only on [Repository.name] field.
	// 2. An operator. Can be either `=` or `!=` for single values, `IN` or `NOT IN` for lists of values.
	// 3. Value or a list of values to compare against the values of the field.
	Filter               string   `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	OrderBy              string   `protobuf:"bytes,5,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRepositoriesRequest) Reset()         { *m = ListRepositoriesRequest{} }
func (m *ListRepositoriesRequest) String() string { return proto.CompactTextString(m) }
func (*ListRepositoriesRequest) ProtoMessage()    {}
func (*ListRepositoriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3ad489b83930188, []int{0}
}

func (m *ListRepositoriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRepositoriesRequest.Unmarshal(m, b)
}
func (m *ListRepositoriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRepositoriesRequest.Marshal(b, m, deterministic)
}
func (m *ListRepositoriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRepositoriesRequest.Merge(m, src)
}
func (m *ListRepositoriesRequest) XXX_Size() int {
	return xxx_messageInfo_ListRepositoriesRequest.Size(m)
}
func (m *ListRepositoriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRepositoriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRepositoriesRequest proto.InternalMessageInfo

func (m *ListRepositoriesRequest) GetRegistryId() string {
	if m != nil {
		return m.RegistryId
	}
	return ""
}

func (m *ListRepositoriesRequest) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *ListRepositoriesRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListRepositoriesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListRepositoriesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListRepositoriesRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

type ListRepositoriesResponse struct {
	// List of Repository resources.
	Repositories []*Repository `protobuf:"bytes,1,rep,name=repositories,proto3" json:"repositories,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListRepositoriesRequest.page_size], use
	// the [next_page_token] as the value
	// for the [ListRepositoriesRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRepositoriesResponse) Reset()         { *m = ListRepositoriesResponse{} }
func (m *ListRepositoriesResponse) String() string { return proto.CompactTextString(m) }
func (*ListRepositoriesResponse) ProtoMessage()    {}
func (*ListRepositoriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3ad489b83930188, []int{1}
}

func (m *ListRepositoriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRepositoriesResponse.Unmarshal(m, b)
}
func (m *ListRepositoriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRepositoriesResponse.Marshal(b, m, deterministic)
}
func (m *ListRepositoriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRepositoriesResponse.Merge(m, src)
}
func (m *ListRepositoriesResponse) XXX_Size() int {
	return xxx_messageInfo_ListRepositoriesResponse.Size(m)
}
func (m *ListRepositoriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRepositoriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRepositoriesResponse proto.InternalMessageInfo

func (m *ListRepositoriesResponse) GetRepositories() []*Repository {
	if m != nil {
		return m.Repositories
	}
	return nil
}

func (m *ListRepositoriesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*ListRepositoriesRequest)(nil), "yandex.cloud.containerregistry.v1.ListRepositoriesRequest")
	proto.RegisterType((*ListRepositoriesResponse)(nil), "yandex.cloud.containerregistry.v1.ListRepositoriesResponse")
}

func init() {
	proto.RegisterFile("yandex/cloud/containerregistry/v1/repository_service.proto", fileDescriptor_b3ad489b83930188)
}

var fileDescriptor_b3ad489b83930188 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0x8e, 0xd3, 0x40,
	0x14, 0x85, 0xe5, 0x24, 0x1b, 0xe2, 0xbb, 0x20, 0xc4, 0x34, 0x58, 0x11, 0x48, 0x21, 0xb0, 0x10,
	0x84, 0x62, 0xc7, 0x41, 0x34, 0xfb, 0xd3, 0xa4, 0x5b, 0x89, 0x02, 0xbc, 0x48, 0x08, 0x9a, 0xc8,
	0x89, 0xef, 0x9a, 0x11, 0x66, 0xae, 0x99, 0x99, 0x44, 0xeb, 0x2d, 0xa9, 0x50, 0x5a, 0xc4, 0x5b,
	0xf0, 0x14, 0x34, 0xa4, 0xe7, 0x15, 0x78, 0x0a, 0x2a, 0xe4, 0x99, 0xf5, 0xfe, 0x05, 0x14, 0xa0,
	0x9d, 0x73, 0xce, 0x37, 0x77, 0xee, 0x1c, 0xd8, 0x2e, 0x62, 0x91, 0xe0, 0x51, 0x30, 0xcd, 0x68,
	0x96, 0x04, 0x53, 0x12, 0x3a, 0xe6, 0x02, 0xa5, 0xc4, 0x94, 0x2b, 0x2d, 0x8b, 0x60, 0x1e, 0x06,
	0x12, 0x73, 0x52, 0x5c, 0x93, 0x2c, 0xc6, 0x0a, 0xe5, 0x9c, 0x4f, 0xd1, 0xcf, 0x25, 0x69, 0x62,
	0x77, 0x6c, 0xd6, 0x37, 0x59, 0x7f, 0x25, 0xeb, 0xcf, 0xc3, 0xf6, 0xf0, 0x5f, 0xf0, 0x16, 0xdb,
	0xbe, 0x7d, 0x21, 0x33, 0x8f, 0x33, 0x9e, 0xc4, 0x9a, 0x93, 0x38, 0x91, 0x6f, 0xa5, 0x44, 0x69,
	0x86, 0x41, 0x9c, 0xf3, 0x20, 0x16, 0x82, 0xb4, 0x11, 0x95, 0x55, 0xbb, 0x1f, 0x6b, 0x70, 0xf3,
	0x29, 0x57, 0x3a, 0xaa, 0xa8, 0x1c, 0x55, 0x84, 0xef, 0x67, 0xa8, 0x34, 0x7b, 0x08, 0x9b, 0xd5,
	0xc5, 0x63, 0x9e, 0x78, 0x4e, 0xc7, 0xe9, 0xb9, 0xa3, 0xd6, 0x62, 0x19, 0x36, 0x76, 0xf7, 0x9e,
	0x0c, 0x22, 0xa8, 0xc4, 0xfd, 0x84, 0x6d, 0x81, 0x7b, 0x48, 0x59, 0x82, 0xb2, 0x34, 0x36, 0x2f,
	0x19, 0x5b, 0x56, 0xda, 0x4f, 0xd8, 0x03, 0x70, 0xf3, 0x38, 0xc5, 0xb1, 0xe2, 0xc7, 0xe8, 0xd5,
	0x3a, 0x4e, 0xaf, 0x3e, 0x82, 0x9f, 0xdf, 0xc2, 0xe6, 0xee, 0x5e, 0x38, 0x18, 0x0c, 0xa2, 0x56,
	0x29, 0x1e, 0xf0, 0x63, 0x64, 0x3d, 0x00, 0x63, 0xd4, 0xf4, 0x16, 0x85, 0x57, 0x37, 0x40, 0x77,
	0xb1, 0x0c, 0x37, 0x8c, 0x33, 0x32, 0x94, 0x17, 0xa5, 0xc6, 0xba, 0xd0, 0x3c, 0xe4, 0x99, 0x46,
	0xe9, 0x35, 0x8c, 0x0b, 0x16, 0xcb, 0x53, 0xde, 0x89, 0xc2, 0xee, 0x41, 0x8b, 0x64, 0x39, 0xdc,
	0xa4, 0xf0, 0x36, 0x2e, 0xb3, 0xae, 0x18, 0x69, 0x54, 0x74, 0x3f, 0x3b, 0xe0, 0xad, 0xae, 0x42,
	0xe5, 0x24, 0x14, 0xb2, 0xe7, 0x70, 0x55, 0x9e, 0x3b, 0xf7, 0x9c, 0x4e, 0xbd, 0xb7, 0x39, 0xec,
	0xfb, 0x6b, 0xbf, 0xd4, 0x3f, 0xc5, 0x15, 0xd1, 0x05, 0x04, 0xbb, 0x0f, 0xd7, 0x05, 0x1e, 0xe9,
	0xf1, 0xb9, 0x87, 0x96, 0x2b, 0x71, 0xa3, 0x6b, 0xe5, 0xf1, 0xb3, 0xea, 0x85, 0xc3, 0xaf, 0x0e,
	0xdc, 0x38, 0x83, 0x1c, 0xd8, 0x4a, 0xb1, 0x2f, 0x0e, 0x34, 0xca, 0x69, 0xd9, 0xf6, 0x5f, 0xcc,
	0xf0, 0x87, 0x1f, 0x6e, 0xef, 0xfc, 0x57, 0xd6, 0xae, 0xa4, 0xfb, 0xe8, 0xc3, 0xf7, 0x1f, 0x9f,
	0x6a, 0x5b, 0xec, 0xee, 0x59, 0x4f, 0xfb, 0xbf, 0x2d, 0x2a, 0x47, 0x35, 0x7a, 0xf5, 0xfa, 0x65,
	0xca, 0xf5, 0x9b, 0xd9, 0xc4, 0x9f, 0xd2, 0xbb, 0xc0, 0xde, 0xda, 0xb7, 0x8d, 0x4d, 0xa9, 0x9f,
	0xa2, 0x30, 0x75, 0x0c, 0xd6, 0xd6, 0x7f, 0x67, 0xe5, 0x70, 0xd2, 0x34, 0xd1, 0xc7, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xde, 0x4e, 0x2c, 0x46, 0x9b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RepositoryServiceClient is the client API for RepositoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepositoryServiceClient interface {
	// Retrieves the list of Repository resources in the specified registry.
	List(ctx context.Context, in *ListRepositoriesRequest, opts ...grpc.CallOption) (*ListRepositoriesResponse, error)
}

type repositoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryServiceClient(cc *grpc.ClientConn) RepositoryServiceClient {
	return &repositoryServiceClient{cc}
}

func (c *repositoryServiceClient) List(ctx context.Context, in *ListRepositoriesRequest, opts ...grpc.CallOption) (*ListRepositoriesResponse, error) {
	out := new(ListRepositoriesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.containerregistry.v1.RepositoryService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServiceServer is the server API for RepositoryService service.
type RepositoryServiceServer interface {
	// Retrieves the list of Repository resources in the specified registry.
	List(context.Context, *ListRepositoriesRequest) (*ListRepositoriesResponse, error)
}

// UnimplementedRepositoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRepositoryServiceServer struct {
}

func (*UnimplementedRepositoryServiceServer) List(ctx context.Context, req *ListRepositoriesRequest) (*ListRepositoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterRepositoryServiceServer(s *grpc.Server, srv RepositoryServiceServer) {
	s.RegisterService(&_RepositoryService_serviceDesc, srv)
}

func _RepositoryService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepositoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.containerregistry.v1.RepositoryService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServiceServer).List(ctx, req.(*ListRepositoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RepositoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.containerregistry.v1.RepositoryService",
	HandlerType: (*RepositoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _RepositoryService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/containerregistry/v1/repository_service.proto",
}