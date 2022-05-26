// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/applications/icq/v1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v3/modules/apps/icq/host/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the interchain query genesis state
type GenesisState struct {
	HostGenesisState HostGenesisState `protobuf:"bytes,2,opt,name=host_genesis_state,json=hostGenesisState,proto3" json:"host_genesis_state" yaml:"host_genesis_state"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_44cd15e54e20f195, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetHostGenesisState() HostGenesisState {
	if m != nil {
		return m.HostGenesisState
	}
	return HostGenesisState{}
}

// HostGenesisState defines the interchain accounts host genesis state
type HostGenesisState struct {
	Port   string       `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	Params types.Params `protobuf:"bytes,4,opt,name=params,proto3" json:"params"`
}

func (m *HostGenesisState) Reset()         { *m = HostGenesisState{} }
func (m *HostGenesisState) String() string { return proto.CompactTextString(m) }
func (*HostGenesisState) ProtoMessage()    {}
func (*HostGenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_44cd15e54e20f195, []int{1}
}
func (m *HostGenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostGenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostGenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostGenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostGenesisState.Merge(m, src)
}
func (m *HostGenesisState) XXX_Size() int {
	return m.Size()
}
func (m *HostGenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_HostGenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_HostGenesisState proto.InternalMessageInfo

func (m *HostGenesisState) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *HostGenesisState) GetParams() types.Params {
	if m != nil {
		return m.Params
	}
	return types.Params{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ibc.applications.icq.v1.GenesisState")
	proto.RegisterType((*HostGenesisState)(nil), "ibc.applications.icq.v1.HostGenesisState")
}

func init() {
	proto.RegisterFile("ibc/applications/icq/v1/genesis.proto", fileDescriptor_44cd15e54e20f195)
}

var fileDescriptor_44cd15e54e20f195 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4f, 0x32, 0x31,
	0x18, 0xc7, 0xaf, 0x6f, 0x08, 0xc9, 0x5b, 0x1d, 0x48, 0x63, 0x22, 0x32, 0x14, 0xbc, 0x68, 0xc4,
	0xc1, 0x36, 0xc0, 0xe6, 0x78, 0x8b, 0x0e, 0x0e, 0x06, 0x37, 0x17, 0xd2, 0xab, 0x97, 0xa3, 0x86,
	0xe3, 0x29, 0xb4, 0x5c, 0x82, 0x1f, 0xc1, 0xc9, 0x8f, 0xc5, 0xc8, 0xe8, 0x44, 0xcc, 0xdd, 0x37,
	0xf0, 0x13, 0x98, 0xf6, 0x18, 0x08, 0xe2, 0xd4, 0x27, 0xed, 0xef, 0xf9, 0x3d, 0xff, 0x3e, 0xf8,
	0x52, 0xc5, 0x92, 0x0b, 0xad, 0x27, 0x4a, 0x0a, 0xab, 0x60, 0x6a, 0xb8, 0x92, 0x33, 0x9e, 0xf7,
	0x78, 0x9a, 0x4c, 0x13, 0xa3, 0x0c, 0xd3, 0x73, 0xb0, 0x40, 0x4e, 0x55, 0x2c, 0xd9, 0x2e, 0xc6,
	0x94, 0x9c, 0xb1, 0xbc, 0xd7, 0x3a, 0x49, 0x21, 0x05, 0xcf, 0x70, 0x57, 0x55, 0x78, 0xeb, 0xea,
	0xa0, 0x75, 0x0c, 0xc6, 0x3a, 0xb5, 0x3b, 0x2b, 0x30, 0x7c, 0x47, 0xf8, 0xf8, 0xae, 0x9a, 0xf4,
	0x64, 0x85, 0x4d, 0xc8, 0x1b, 0x26, 0xee, 0x79, 0xb4, 0x1d, 0x3f, 0x32, 0xee, 0xb6, 0xf9, 0xaf,
	0x83, 0xba, 0x47, 0xfd, 0x6b, 0xf6, 0x47, 0x0a, 0x76, 0x0f, 0xc6, 0xee, 0x6a, 0xa2, 0xf3, 0xd5,
	0xa6, 0x1d, 0x7c, 0x6f, 0xda, 0x67, 0x4b, 0x91, 0x4d, 0x6e, 0xc3, 0xdf, 0xca, 0x70, 0xd8, 0x18,
	0xef, 0x35, 0x85, 0xaf, 0xb8, 0xb1, 0x2f, 0x22, 0x04, 0xd7, 0x34, 0xcc, 0x6d, 0x13, 0x75, 0x50,
	0xf7, 0xff, 0xd0, 0xd7, 0x24, 0xc2, 0x75, 0x2d, 0xe6, 0x22, 0x33, 0xcd, 0x9a, 0xcf, 0x75, 0x71,
	0x38, 0x97, 0xff, 0x66, 0xde, 0x63, 0x8f, 0x9e, 0x8d, 0x6a, 0x2e, 0xd2, 0x70, 0xdb, 0x19, 0x3d,
	0xac, 0x0a, 0x8a, 0xd6, 0x05, 0x45, 0x5f, 0x05, 0x45, 0x1f, 0x25, 0x0d, 0xd6, 0x25, 0x0d, 0x3e,
	0x4b, 0x1a, 0x3c, 0xf7, 0x53, 0x65, 0xc7, 0x8b, 0x98, 0x49, 0xc8, 0xb8, 0x04, 0x93, 0x81, 0xe1,
	0x2a, 0x96, 0x37, 0x29, 0xf0, 0x7c, 0xc0, 0x33, 0x78, 0x59, 0x4c, 0x12, 0xe3, 0x76, 0x5b, 0xed,
	0xd4, 0x2e, 0x75, 0x62, 0xe2, 0xba, 0xdf, 0xe6, 0xe0, 0x27, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xcf,
	0xc9, 0x75, 0xce, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.HostGenesisState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func (m *HostGenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostGenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HostGenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.HostGenesisState.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *HostGenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostGenesisState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HostGenesisState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *HostGenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: HostGenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostGenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)