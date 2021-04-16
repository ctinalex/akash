// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/cert/v1beta1/query.proto

package typesv1beta1

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CertificateResponse struct {
	Certificate Certificate `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate" yaml:"certificate"`
	Serial      string      `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial" yaml:"serial"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_253641229681779f, []int{0}
}
func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}
func (m *CertificateResponse) XXX_Size() int {
	return m.Size()
}
func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCertificate() Certificate {
	if m != nil {
		return m.Certificate
	}
	return Certificate{}
}

func (m *CertificateResponse) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

// QueryDeploymentsRequest is request type for the Query/Deployments RPC method
type QueryCertificatesRequest struct {
	Filter     CertificateFilter  `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter"`
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCertificatesRequest) Reset()         { *m = QueryCertificatesRequest{} }
func (m *QueryCertificatesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCertificatesRequest) ProtoMessage()    {}
func (*QueryCertificatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_253641229681779f, []int{1}
}
func (m *QueryCertificatesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCertificatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCertificatesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCertificatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCertificatesRequest.Merge(m, src)
}
func (m *QueryCertificatesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCertificatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCertificatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCertificatesRequest proto.InternalMessageInfo

func (m *QueryCertificatesRequest) GetFilter() CertificateFilter {
	if m != nil {
		return m.Filter
	}
	return CertificateFilter{}
}

func (m *QueryCertificatesRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryCertificatesResponse is response type for the Query/Certificates RPC method
type QueryCertificatesResponse struct {
	Certificates CertificatesResponse `protobuf:"bytes,1,rep,name=certificates,proto3,castrepeated=CertificatesResponse" json:"certificates"`
	Pagination   *query.PageResponse  `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryCertificatesResponse) Reset()         { *m = QueryCertificatesResponse{} }
func (m *QueryCertificatesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCertificatesResponse) ProtoMessage()    {}
func (*QueryCertificatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_253641229681779f, []int{2}
}
func (m *QueryCertificatesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCertificatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCertificatesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCertificatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCertificatesResponse.Merge(m, src)
}
func (m *QueryCertificatesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCertificatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCertificatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCertificatesResponse proto.InternalMessageInfo

func (m *QueryCertificatesResponse) GetCertificates() CertificatesResponse {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *QueryCertificatesResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateResponse)(nil), "akash.cert.v1beta1.CertificateResponse")
	proto.RegisterType((*QueryCertificatesRequest)(nil), "akash.cert.v1beta1.QueryCertificatesRequest")
	proto.RegisterType((*QueryCertificatesResponse)(nil), "akash.cert.v1beta1.QueryCertificatesResponse")
}

func init() { proto.RegisterFile("akash/cert/v1beta1/query.proto", fileDescriptor_253641229681779f) }

var fileDescriptor_253641229681779f = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x3f, 0x6f, 0xd4, 0x30,
	0x18, 0xc6, 0xcf, 0x05, 0x4e, 0xc2, 0x57, 0x16, 0xd3, 0xe1, 0x38, 0x4a, 0x72, 0x8a, 0x54, 0xee,
	0xf8, 0x53, 0x5b, 0x4d, 0x37, 0xc6, 0x54, 0x94, 0x15, 0x32, 0xb2, 0x39, 0x91, 0x9b, 0x5a, 0xcd,
	0xc5, 0x69, 0xec, 0xab, 0xb8, 0x95, 0x4f, 0x80, 0xc4, 0xc6, 0x8c, 0x84, 0xc4, 0x27, 0xe0, 0x23,
	0x74, 0xa3, 0x12, 0x0b, 0x53, 0x40, 0x77, 0x4c, 0x1d, 0xfb, 0x09, 0x50, 0x6c, 0x1f, 0x35, 0x22,
	0xa7, 0x63, 0x4b, 0xfc, 0xbc, 0xcf, 0xf3, 0xfe, 0xde, 0xd7, 0x09, 0xf4, 0xe8, 0x09, 0x95, 0xc7,
	0x24, 0x65, 0x95, 0x22, 0x67, 0x7b, 0x09, 0x53, 0x74, 0x8f, 0x9c, 0x4e, 0x59, 0x35, 0xc3, 0x65,
	0x25, 0x94, 0x40, 0x48, 0xeb, 0xb8, 0xd1, 0xb1, 0xd5, 0x07, 0x5b, 0x99, 0xc8, 0x84, 0x96, 0x49,
	0xf3, 0x64, 0x2a, 0x07, 0xdb, 0x99, 0x10, 0x59, 0xce, 0x08, 0x2d, 0x39, 0xa1, 0x45, 0x21, 0x14,
	0x55, 0x5c, 0x14, 0xd2, 0xaa, 0x8f, 0x53, 0x21, 0x27, 0x42, 0x92, 0x84, 0x4a, 0x66, 0x1a, 0xfc,
	0x69, 0x57, 0xd2, 0x8c, 0x17, 0xba, 0xd8, 0xd6, 0x3e, 0x68, 0x61, 0xd2, 0x00, 0x5a, 0x0e, 0xbe,
	0x00, 0x78, 0xf7, 0x80, 0x55, 0x8a, 0x1f, 0xf1, 0x94, 0x2a, 0x16, 0x33, 0x59, 0x8a, 0x42, 0x32,
	0x94, 0xc3, 0x5e, 0x7a, 0x7d, 0xdc, 0x07, 0x43, 0x30, 0xee, 0x85, 0x3e, 0xfe, 0x77, 0x00, 0xec,
	0xb8, 0xa3, 0x47, 0xe7, 0xb5, 0xdf, 0xb9, 0xac, 0x7d, 0xd7, 0x7b, 0x55, 0xfb, 0x68, 0x46, 0x27,
	0xf9, 0xb3, 0xc0, 0x39, 0x0c, 0x62, 0xb7, 0x04, 0xed, 0xc3, 0xae, 0x64, 0x15, 0xa7, 0x79, 0x7f,
	0x63, 0x08, 0xc6, 0xb7, 0xa3, 0xfb, 0x97, 0xb5, 0x6f, 0x4f, 0xae, 0x6a, 0xff, 0x8e, 0xb1, 0x9b,
	0xf7, 0x20, 0xb6, 0x42, 0xf0, 0x09, 0xc0, 0xfe, 0xab, 0x66, 0x78, 0x87, 0x40, 0xc6, 0xec, 0x74,
	0xca, 0xa4, 0x42, 0x07, 0xb0, 0x7b, 0xc4, 0x73, 0xc5, 0x2a, 0x8b, 0xbe, 0xb3, 0x06, 0xfd, 0x50,
	0x17, 0x47, 0x37, 0x9b, 0x01, 0x62, 0x6b, 0x45, 0x87, 0x10, 0x5e, 0xef, 0x53, 0xa3, 0xf5, 0xc2,
	0x87, 0xd8, 0x2c, 0x1f, 0x37, 0xcb, 0xc7, 0xe6, 0x76, 0x97, 0x79, 0x2f, 0x69, 0xc6, 0x2c, 0x40,
	0xec, 0x38, 0x83, 0xaf, 0x00, 0xde, 0x6b, 0x21, 0xb5, 0xab, 0xe6, 0x70, 0xd3, 0xd9, 0x85, 0xec,
	0x83, 0xe1, 0x8d, 0x71, 0x2f, 0x1c, 0xad, 0x01, 0x5e, 0xda, 0xa3, 0xed, 0x06, 0xf9, 0xf3, 0x0f,
	0x7f, 0xab, 0x2d, 0x3c, 0xfe, 0x2b, 0x1a, 0xbd, 0x68, 0x19, 0x68, 0xb4, 0x76, 0x20, 0x1b, 0xe5,
	0x58, 0xc3, 0x8f, 0x00, 0xde, 0xd2, 0x13, 0xa1, 0x0f, 0x00, 0x6e, 0xba, 0x9d, 0xd1, 0xd3, 0x36,
	0xf0, 0x55, 0xf7, 0x34, 0xd8, 0xfd, 0xcf, 0x6a, 0xc3, 0x10, 0xec, 0xbe, 0xfd, 0xf6, 0xeb, 0xfd,
	0xc6, 0x08, 0xed, 0x90, 0x15, 0x9f, 0xf5, 0xd2, 0x41, 0x72, 0x2e, 0x55, 0xf4, 0xfc, 0x7c, 0xee,
	0x81, 0x8b, 0xb9, 0x07, 0x7e, 0xce, 0x3d, 0xf0, 0x6e, 0xe1, 0x75, 0x2e, 0x16, 0x5e, 0xe7, 0xfb,
	0xc2, 0xeb, 0xbc, 0x7e, 0x92, 0x71, 0x75, 0x3c, 0x4d, 0x70, 0x2a, 0x26, 0x44, 0x9c, 0x55, 0x69,
	0x7e, 0x62, 0x13, 0xdf, 0x98, 0x4c, 0x35, 0x2b, 0x99, 0xb4, 0xc1, 0x49, 0x57, 0xff, 0x2b, 0xfb,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xc0, 0x54, 0x82, 0xe0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Certificates queries certificates
	Certificates(ctx context.Context, in *QueryCertificatesRequest, opts ...grpc.CallOption) (*QueryCertificatesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Certificates(ctx context.Context, in *QueryCertificatesRequest, opts ...grpc.CallOption) (*QueryCertificatesResponse, error) {
	out := new(QueryCertificatesResponse)
	err := c.cc.Invoke(ctx, "/akash.cert.v1beta1.Query/Certificates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Certificates queries certificates
	Certificates(context.Context, *QueryCertificatesRequest) (*QueryCertificatesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Certificates(ctx context.Context, req *QueryCertificatesRequest) (*QueryCertificatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Certificates not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Certificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Certificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/akash.cert.v1beta1.Query/Certificates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Certificates(ctx, req.(*QueryCertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "akash.cert.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Certificates",
			Handler:    _Query_Certificates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "akash/cert/v1beta1/query.proto",
}

func (m *CertificateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertificateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CertificateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Serial) > 0 {
		i -= len(m.Serial)
		copy(dAtA[i:], m.Serial)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Serial)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Certificate.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryCertificatesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCertificatesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCertificatesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Filter.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryCertificatesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCertificatesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCertificatesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Certificates) > 0 {
		for iNdEx := len(m.Certificates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Certificates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CertificateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Certificate.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = len(m.Serial)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCertificatesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Filter.Size()
	n += 1 + l + sovQuery(uint64(l))
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCertificatesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Certificates) > 0 {
		for _, e := range m.Certificates {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CertificateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertificateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertificateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Certificate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Serial", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Serial = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCertificatesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCertificatesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCertificatesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Filter.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCertificatesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCertificatesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCertificatesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificates = append(m.Certificates, CertificateResponse{})
			if err := m.Certificates[len(m.Certificates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)