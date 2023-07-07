// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coco/classicdice/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgDiceBetting struct {
	Creator       string      `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Option        string      `protobuf:"bytes,2,opt,name=option,proto3" json:"option,omitempty"`
	NumberBetting uint32      `protobuf:"varint,3,opt,name=number_betting,json=numberBetting,proto3" json:"number_betting,omitempty"`
	Coin          *types.Coin `protobuf:"bytes,4,opt,name=coin,proto3" json:"coin,omitempty"`
}

func (m *MsgDiceBetting) Reset()         { *m = MsgDiceBetting{} }
func (m *MsgDiceBetting) String() string { return proto.CompactTextString(m) }
func (*MsgDiceBetting) ProtoMessage()    {}
func (*MsgDiceBetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4c9b1f9abafa840, []int{0}
}
func (m *MsgDiceBetting) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDiceBetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDiceBetting.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDiceBetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDiceBetting.Merge(m, src)
}
func (m *MsgDiceBetting) XXX_Size() int {
	return m.Size()
}
func (m *MsgDiceBetting) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDiceBetting.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDiceBetting proto.InternalMessageInfo

func (m *MsgDiceBetting) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDiceBetting) GetOption() string {
	if m != nil {
		return m.Option
	}
	return ""
}

func (m *MsgDiceBetting) GetNumberBetting() uint32 {
	if m != nil {
		return m.NumberBetting
	}
	return 0
}

func (m *MsgDiceBetting) GetCoin() *types.Coin {
	if m != nil {
		return m.Coin
	}
	return nil
}

type MsgDiceBettingResponse struct {
}

func (m *MsgDiceBettingResponse) Reset()         { *m = MsgDiceBettingResponse{} }
func (m *MsgDiceBettingResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDiceBettingResponse) ProtoMessage()    {}
func (*MsgDiceBettingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4c9b1f9abafa840, []int{1}
}
func (m *MsgDiceBettingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDiceBettingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDiceBettingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDiceBettingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDiceBettingResponse.Merge(m, src)
}
func (m *MsgDiceBettingResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDiceBettingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDiceBettingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDiceBettingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDiceBetting)(nil), "coco.classicdice.MsgDiceBetting")
	proto.RegisterType((*MsgDiceBettingResponse)(nil), "coco.classicdice.MsgDiceBettingResponse")
}

func init() { proto.RegisterFile("coco/classicdice/tx.proto", fileDescriptor_b4c9b1f9abafa840) }

var fileDescriptor_b4c9b1f9abafa840 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x6b, 0x5a, 0x15, 0xe1, 0xaa, 0x15, 0xca, 0x50, 0xb9, 0x1d, 0xac, 0xa8, 0x12, 0x52,
	0x16, 0x1c, 0x35, 0xbc, 0x41, 0x61, 0xed, 0x92, 0x0d, 0x16, 0x48, 0xcc, 0x29, 0xb2, 0x44, 0x7d,
	0x51, 0xce, 0xa0, 0xf2, 0x16, 0x8c, 0x3c, 0x12, 0x63, 0x47, 0x46, 0x94, 0xbc, 0x08, 0x4a, 0xd2,
	0x48, 0x0d, 0x0b, 0xe3, 0xf9, 0xbf, 0xfb, 0xff, 0xcf, 0x3f, 0x5f, 0x68, 0xd4, 0x18, 0xea, 0x97,
	0x84, 0xc8, 0xe8, 0x67, 0xa3, 0x21, 0x74, 0x7b, 0x95, 0x17, 0xe8, 0xd0, 0xbb, 0xac, 0x25, 0x75,
	0x22, 0x2d, 0xa5, 0x46, 0xda, 0x21, 0x85, 0x69, 0x42, 0x10, 0xbe, 0xad, 0x53, 0x70, 0xc9, 0x3a,
	0xd4, 0x68, 0x6c, 0x7b, 0xb1, 0xfa, 0x64, 0x7c, 0xb6, 0xa5, 0xec, 0xce, 0x68, 0xd8, 0x80, 0x73,
	0xc6, 0x66, 0x9e, 0xe0, 0xe7, 0xba, 0x80, 0xc4, 0x61, 0x21, 0x98, 0xcf, 0x82, 0x8b, 0xb8, 0x1b,
	0xbd, 0x39, 0x1f, 0x63, 0xee, 0x0c, 0x5a, 0x71, 0xd6, 0x08, 0xc7, 0xc9, 0xbb, 0xe2, 0x33, 0xfb,
	0xba, 0x4b, 0xa1, 0x78, 0x4c, 0x5b, 0x0f, 0x31, 0xf4, 0x59, 0x30, 0x8d, 0xa7, 0xed, 0x6b, 0x67,
	0x7c, 0xcd, 0x47, 0x75, 0xb2, 0x18, 0xf9, 0x2c, 0x98, 0x44, 0x0b, 0xd5, 0xa2, 0xa9, 0x1a, 0x4d,
	0x1d, 0xd1, 0xd4, 0x2d, 0x1a, 0x1b, 0x37, 0x6b, 0x2b, 0xc1, 0xe7, 0x7d, 0xb2, 0x18, 0x28, 0x47,
	0x4b, 0x10, 0x3d, 0xf1, 0xe1, 0x96, 0x32, 0xef, 0x9e, 0x4f, 0x4e, 0xb9, 0x7d, 0xf5, 0xf7, 0xf7,
	0xaa, 0x7f, 0xbf, 0x0c, 0xfe, 0xdb, 0xe8, 0x12, 0x36, 0xd1, 0x57, 0x29, 0xd9, 0xa1, 0x94, 0xec,
	0xa7, 0x94, 0xec, 0xa3, 0x92, 0x83, 0x43, 0x25, 0x07, 0xdf, 0x95, 0x1c, 0x3c, 0x88, 0xa6, 0xfd,
	0x7d, 0xbf, 0xff, 0xf7, 0x1c, 0x28, 0x1d, 0x37, 0x8d, 0xde, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x95, 0x99, 0x99, 0xec, 0xa0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	DiceBetting(ctx context.Context, in *MsgDiceBetting, opts ...grpc.CallOption) (*MsgDiceBettingResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) DiceBetting(ctx context.Context, in *MsgDiceBetting, opts ...grpc.CallOption) (*MsgDiceBettingResponse, error) {
	out := new(MsgDiceBettingResponse)
	err := c.cc.Invoke(ctx, "/coco.classicdice.Msg/DiceBetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	DiceBetting(context.Context, *MsgDiceBetting) (*MsgDiceBettingResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) DiceBetting(ctx context.Context, req *MsgDiceBetting) (*MsgDiceBettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiceBetting not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_DiceBetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDiceBetting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DiceBetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coco.classicdice.Msg/DiceBetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DiceBetting(ctx, req.(*MsgDiceBetting))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "coco.classicdice.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DiceBetting",
			Handler:    _Msg_DiceBetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coco/classicdice/tx.proto",
}

func (m *MsgDiceBetting) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDiceBetting) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDiceBetting) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Coin != nil {
		{
			size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.NumberBetting != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.NumberBetting))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Option) > 0 {
		i -= len(m.Option)
		copy(dAtA[i:], m.Option)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Option)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDiceBettingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDiceBettingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDiceBettingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgDiceBetting) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Option)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.NumberBetting != 0 {
		n += 1 + sovTx(uint64(m.NumberBetting))
	}
	if m.Coin != nil {
		l = m.Coin.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDiceBettingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgDiceBetting) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDiceBetting: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDiceBetting: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Option", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Option = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumberBetting", wireType)
			}
			m.NumberBetting = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumberBetting |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Coin == nil {
				m.Coin = &types.Coin{}
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDiceBettingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDiceBettingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDiceBettingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
