// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/icq/v1/packet.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/tendermint/tendermint/abci/types"
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

// Type defines a classification of message issued from a controller chain to its associated interchain query
// host
type Type int32

const (
	// Default zero value enumeration
	UNSPECIFIED Type = 0
	// Runs ABCI query request in host chain and sends back the results to the controller chain
	ABCI Type = 1
)

var Type_name = map[int32]string{
	0: "TYPE_UNSPECIFIED",
	1: "TYPE_ABCI",
}

var Type_value = map[string]int32{
	"TYPE_UNSPECIFIED": 0,
	"TYPE_ABCI":        1,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_841c9ad988725ce1, []int{0}
}

// InterchainQueryPacketData is comprised of an ABCI query request and type of query.
type InterchainQueryPacketData struct {
	Type    Type               `protobuf:"varint,1,opt,name=type,proto3,enum=ibc.applications.icq.v1.Type" json:"type,omitempty"`
	Request types.RequestQuery `protobuf:"bytes,2,opt,name=request,proto3" json:"request"`
}

func (m *InterchainQueryPacketData) Reset()         { *m = InterchainQueryPacketData{} }
func (m *InterchainQueryPacketData) String() string { return proto.CompactTextString(m) }
func (*InterchainQueryPacketData) ProtoMessage()    {}
func (*InterchainQueryPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_841c9ad988725ce1, []int{0}
}
func (m *InterchainQueryPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterchainQueryPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterchainQueryPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterchainQueryPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterchainQueryPacketData.Merge(m, src)
}
func (m *InterchainQueryPacketData) XXX_Size() int {
	return m.Size()
}
func (m *InterchainQueryPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_InterchainQueryPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_InterchainQueryPacketData proto.InternalMessageInfo

func (m *InterchainQueryPacketData) GetType() Type {
	if m != nil {
		return m.Type
	}
	return UNSPECIFIED
}

func (m *InterchainQueryPacketData) GetRequest() types.RequestQuery {
	if m != nil {
		return m.Request
	}
	return types.RequestQuery{}
}

// InterchainQueryPacketAck is comprised of an ABCI query response with non-deterministic fields left empty (e.g. Codespace, Log, Info and ...).
type InterchainQueryPacketAck struct {
	Response types.ResponseQuery `protobuf:"bytes,1,opt,name=response,proto3" json:"response"`
}

func (m *InterchainQueryPacketAck) Reset()         { *m = InterchainQueryPacketAck{} }
func (m *InterchainQueryPacketAck) String() string { return proto.CompactTextString(m) }
func (*InterchainQueryPacketAck) ProtoMessage()    {}
func (*InterchainQueryPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_841c9ad988725ce1, []int{1}
}
func (m *InterchainQueryPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterchainQueryPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterchainQueryPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterchainQueryPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterchainQueryPacketAck.Merge(m, src)
}
func (m *InterchainQueryPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *InterchainQueryPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_InterchainQueryPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_InterchainQueryPacketAck proto.InternalMessageInfo

func (m *InterchainQueryPacketAck) GetResponse() types.ResponseQuery {
	if m != nil {
		return m.Response
	}
	return types.ResponseQuery{}
}

func init() {
	proto.RegisterEnum("ibc.applications.icq.v1.Type", Type_name, Type_value)
	proto.RegisterType((*InterchainQueryPacketData)(nil), "ibc.applications.icq.v1.InterchainQueryPacketData")
	proto.RegisterType((*InterchainQueryPacketAck)(nil), "ibc.applications.icq.v1.InterchainQueryPacketAck")
}

func init() {
	proto.RegisterFile("ibc/applications/icq/v1/packet.proto", fileDescriptor_841c9ad988725ce1)
}

var fileDescriptor_841c9ad988725ce1 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdf, 0xca, 0xd3, 0x30,
	0x18, 0xc6, 0x1b, 0x29, 0xfa, 0x99, 0x81, 0x8e, 0x22, 0x7c, 0xfb, 0x2a, 0x8b, 0x65, 0x28, 0x0c,
	0xc1, 0x84, 0x75, 0xc7, 0x82, 0xfb, 0x27, 0x14, 0x86, 0xcc, 0x39, 0x0f, 0x14, 0x41, 0xd2, 0x2c,
	0x76, 0x61, 0x6b, 0xd3, 0x35, 0xe9, 0xa0, 0x77, 0x20, 0x03, 0xc1, 0x1b, 0xd8, 0x91, 0x37, 0xb3,
	0xc3, 0x1d, 0x7a, 0x24, 0xb2, 0xdd, 0x88, 0x34, 0xf5, 0xcf, 0x0e, 0xe6, 0xd9, 0x4b, 0xde, 0xdf,
	0xf3, 0x3e, 0x0f, 0x3c, 0x81, 0x8f, 0x45, 0xc8, 0x08, 0x4d, 0xd3, 0x95, 0x60, 0x54, 0x0b, 0x99,
	0x28, 0x22, 0xd8, 0x9a, 0x6c, 0x3a, 0x24, 0xa5, 0x6c, 0xc9, 0x35, 0x4e, 0x33, 0xa9, 0xa5, 0x73,
	0x2d, 0x42, 0x86, 0xcf, 0x29, 0x2c, 0xd8, 0x1a, 0x6f, 0x3a, 0xee, 0x4d, 0x24, 0x65, 0xb4, 0xe2,
	0xc4, 0x60, 0x61, 0xfe, 0x89, 0xd0, 0xa4, 0xa8, 0x34, 0xee, 0x83, 0x48, 0x46, 0xd2, 0x8c, 0xa4,
	0x9c, 0x7e, 0xbf, 0x3e, 0xd4, 0x3c, 0x99, 0xf3, 0x2c, 0x16, 0x89, 0x26, 0x34, 0x64, 0x82, 0xe8,
	0x22, 0xe5, 0xaa, 0x5a, 0xb6, 0xbe, 0x00, 0x78, 0x13, 0x24, 0x9a, 0x67, 0x6c, 0x41, 0x45, 0xf2,
	0x3a, 0xe7, 0x59, 0x31, 0x31, 0x31, 0x86, 0x54, 0x53, 0xa7, 0x03, 0xed, 0x12, 0x6e, 0x00, 0x0f,
	0xb4, 0xef, 0xf9, 0x4d, 0xfc, 0x9f, 0x4c, 0x78, 0x56, 0xa4, 0x7c, 0x6a, 0x50, 0xe7, 0x39, 0xbc,
	0x93, 0xf1, 0x75, 0xce, 0x95, 0x6e, 0xdc, 0xf2, 0x40, 0xbb, 0xe6, 0x37, 0xf1, 0x3f, 0x7f, 0x5c,
	0xfa, 0xe3, 0x69, 0xb5, 0x37, 0x66, 0x7d, 0x7b, 0xff, 0xe3, 0x91, 0x35, 0xfd, 0xa3, 0x69, 0x7d,
	0x80, 0x8d, 0x8b, 0x71, 0x7a, 0x6c, 0xe9, 0xbc, 0x80, 0x57, 0x19, 0x57, 0xa9, 0x4c, 0x54, 0x95,
	0xa8, 0xe6, 0xa3, 0x0b, 0xb7, 0x2b, 0xe0, 0xfc, 0xf8, 0x5f, 0xd5, 0xd3, 0x31, 0xb4, 0xcb, 0xa8,
	0xce, 0x13, 0x58, 0x9f, 0xbd, 0x9b, 0x8c, 0x3e, 0xbe, 0x7d, 0xf5, 0x66, 0x32, 0x1a, 0x04, 0x2f,
	0x83, 0xd1, 0xb0, 0x6e, 0xb9, 0xf7, 0xb7, 0x3b, 0xaf, 0x76, 0xf6, 0xe4, 0x5c, 0xc3, 0xbb, 0x06,
	0xeb, 0xf5, 0x07, 0x41, 0x1d, 0xb8, 0x57, 0xdb, 0x9d, 0x67, 0x97, 0xb3, 0x6b, 0x7f, 0xfe, 0x86,
	0xac, 0xfe, 0x78, 0x7f, 0x44, 0xe0, 0x70, 0x44, 0xe0, 0xe7, 0x11, 0x81, 0xaf, 0x27, 0x64, 0x1d,
	0x4e, 0xc8, 0xfa, 0x7e, 0x42, 0xd6, 0x7b, 0x3f, 0x12, 0x7a, 0x91, 0x87, 0x98, 0xc9, 0x98, 0x30,
	0xa9, 0x62, 0xa9, 0x88, 0x08, 0xd9, 0xb3, 0x48, 0x92, 0x4d, 0x97, 0xc4, 0x72, 0x9e, 0xaf, 0xb8,
	0x2a, 0xbf, 0x40, 0x55, 0xbd, 0xe9, 0x23, 0xbc, 0x6d, 0x0a, 0xe9, 0xfe, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0xef, 0x28, 0xeb, 0xe1, 0x1f, 0x02, 0x00, 0x00,
}

func (m *InterchainQueryPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterchainQueryPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterchainQueryPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Request.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPacket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Type != 0 {
		i = encodeVarintPacket(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *InterchainQueryPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterchainQueryPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterchainQueryPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Response.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPacket(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InterchainQueryPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovPacket(uint64(m.Type))
	}
	l = m.Request.Size()
	n += 1 + l + sovPacket(uint64(l))
	return n
}

func (m *InterchainQueryPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Response.Size()
	n += 1 + l + sovPacket(uint64(l))
	return n
}

func sovPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPacket(x uint64) (n int) {
	return sovPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InterchainQueryPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: InterchainQueryPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterchainQueryPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func (m *InterchainQueryPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPacket
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
			return fmt.Errorf("proto: InterchainQueryPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterchainQueryPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPacket
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
				return ErrInvalidLengthPacket
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Response.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPacket
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
func skipPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
					return 0, ErrIntOverflowPacket
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
				return 0, ErrInvalidLengthPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPacket = fmt.Errorf("proto: unexpected end of group")
)
