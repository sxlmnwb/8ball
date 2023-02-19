// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eightball/eightball/spirit_conjuring_poems.proto

package types

import (
	fmt "fmt"
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

type SpiritConjuringPoems struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MagicKeyId  uint64 `protobuf:"varint,2,opt,name=magicKeyId,proto3" json:"magicKeyId,omitempty"`
	BlockHeight uint64 `protobuf:"varint,3,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	Poem        string `protobuf:"bytes,4,opt,name=poem,proto3" json:"poem,omitempty"`
	ConjuringId uint64 `protobuf:"varint,5,opt,name=conjuringId,proto3" json:"conjuringId,omitempty"`
	UeblPower   uint64 `protobuf:"varint,6,opt,name=ueblPower,proto3" json:"ueblPower,omitempty"`
	Creator     string `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *SpiritConjuringPoems) Reset()         { *m = SpiritConjuringPoems{} }
func (m *SpiritConjuringPoems) String() string { return proto.CompactTextString(m) }
func (*SpiritConjuringPoems) ProtoMessage()    {}
func (*SpiritConjuringPoems) Descriptor() ([]byte, []int) {
	return fileDescriptor_00a8a10c1447f0f5, []int{0}
}
func (m *SpiritConjuringPoems) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SpiritConjuringPoems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SpiritConjuringPoems.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SpiritConjuringPoems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpiritConjuringPoems.Merge(m, src)
}
func (m *SpiritConjuringPoems) XXX_Size() int {
	return m.Size()
}
func (m *SpiritConjuringPoems) XXX_DiscardUnknown() {
	xxx_messageInfo_SpiritConjuringPoems.DiscardUnknown(m)
}

var xxx_messageInfo_SpiritConjuringPoems proto.InternalMessageInfo

func (m *SpiritConjuringPoems) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SpiritConjuringPoems) GetMagicKeyId() uint64 {
	if m != nil {
		return m.MagicKeyId
	}
	return 0
}

func (m *SpiritConjuringPoems) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *SpiritConjuringPoems) GetPoem() string {
	if m != nil {
		return m.Poem
	}
	return ""
}

func (m *SpiritConjuringPoems) GetConjuringId() uint64 {
	if m != nil {
		return m.ConjuringId
	}
	return 0
}

func (m *SpiritConjuringPoems) GetUeblPower() uint64 {
	if m != nil {
		return m.UeblPower
	}
	return 0
}

func (m *SpiritConjuringPoems) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*SpiritConjuringPoems)(nil), "eightball.eightball.SpiritConjuringPoems")
}

func init() {
	proto.RegisterFile("eightball/eightball/spirit_conjuring_poems.proto", fileDescriptor_00a8a10c1447f0f5)
}

var fileDescriptor_00a8a10c1447f0f5 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcd, 0x4c, 0xcf,
	0x28, 0x49, 0x4a, 0xcc, 0xc9, 0xd1, 0x47, 0xb0, 0x8a, 0x0b, 0x32, 0x8b, 0x32, 0x4b, 0xe2, 0x93,
	0xf3, 0xf3, 0xb2, 0x4a, 0x8b, 0x32, 0xf3, 0xd2, 0xe3, 0x0b, 0xf2, 0x53, 0x73, 0x8b, 0xf5, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0xe1, 0xea, 0xf4, 0xe0, 0x2c, 0xa5, 0x6b, 0x8c, 0x5c, 0x22,
	0xc1, 0x60, 0x5d, 0xce, 0x30, 0x4d, 0x01, 0x20, 0x3d, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99, 0x29, 0x42, 0x72, 0x5c, 0x5c, 0xb9, 0x89, 0xe9,
	0x99, 0xc9, 0xde, 0xa9, 0x95, 0x9e, 0x29, 0x12, 0x4c, 0x60, 0x71, 0x24, 0x11, 0x21, 0x05, 0x2e,
	0xee, 0xa4, 0x9c, 0xfc, 0xe4, 0x6c, 0x0f, 0xb0, 0xd9, 0x12, 0xcc, 0x60, 0x05, 0xc8, 0x42, 0x42,
	0x42, 0x5c, 0x2c, 0x20, 0xe7, 0x48, 0xb0, 0x28, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x20, 0x5d,
	0x70, 0xc7, 0x7a, 0xa6, 0x48, 0xb0, 0x42, 0x74, 0x21, 0x09, 0x09, 0xc9, 0x70, 0x71, 0x96, 0xa6,
	0x26, 0xe5, 0x04, 0xe4, 0x97, 0xa7, 0x16, 0x49, 0xb0, 0x81, 0xe5, 0x11, 0x02, 0x42, 0x12, 0x5c,
	0xec, 0xc9, 0x45, 0xa9, 0x89, 0x25, 0xf9, 0x45, 0x12, 0xec, 0x60, 0x63, 0x61, 0x5c, 0x27, 0xd3,
	0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39,
	0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x92, 0x46, 0x84, 0x57, 0x05, 0x52,
	0xd8, 0x95, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xc3, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x7a, 0x70, 0x4d, 0xe7, 0x5f, 0x01, 0x00, 0x00,
}

func (m *SpiritConjuringPoems) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpiritConjuringPoems) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SpiritConjuringPoems) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSpiritConjuringPoems(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x3a
	}
	if m.UeblPower != 0 {
		i = encodeVarintSpiritConjuringPoems(dAtA, i, uint64(m.UeblPower))
		i--
		dAtA[i] = 0x30
	}
	if m.ConjuringId != 0 {
		i = encodeVarintSpiritConjuringPoems(dAtA, i, uint64(m.ConjuringId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Poem) > 0 {
		i -= len(m.Poem)
		copy(dAtA[i:], m.Poem)
		i = encodeVarintSpiritConjuringPoems(dAtA, i, uint64(len(m.Poem)))
		i--
		dAtA[i] = 0x22
	}
	if m.BlockHeight != 0 {
		i = encodeVarintSpiritConjuringPoems(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if m.MagicKeyId != 0 {
		i = encodeVarintSpiritConjuringPoems(dAtA, i, uint64(m.MagicKeyId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintSpiritConjuringPoems(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSpiritConjuringPoems(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpiritConjuringPoems(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SpiritConjuringPoems) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSpiritConjuringPoems(uint64(m.Id))
	}
	if m.MagicKeyId != 0 {
		n += 1 + sovSpiritConjuringPoems(uint64(m.MagicKeyId))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovSpiritConjuringPoems(uint64(m.BlockHeight))
	}
	l = len(m.Poem)
	if l > 0 {
		n += 1 + l + sovSpiritConjuringPoems(uint64(l))
	}
	if m.ConjuringId != 0 {
		n += 1 + sovSpiritConjuringPoems(uint64(m.ConjuringId))
	}
	if m.UeblPower != 0 {
		n += 1 + sovSpiritConjuringPoems(uint64(m.UeblPower))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSpiritConjuringPoems(uint64(l))
	}
	return n
}

func sovSpiritConjuringPoems(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpiritConjuringPoems(x uint64) (n int) {
	return sovSpiritConjuringPoems(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SpiritConjuringPoems) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpiritConjuringPoems
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
			return fmt.Errorf("proto: SpiritConjuringPoems: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpiritConjuringPoems: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpiritConjuringPoems
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MagicKeyId", wireType)
			}
			m.MagicKeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpiritConjuringPoems
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MagicKeyId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpiritConjuringPoems
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Poem", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpiritConjuringPoems
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
				return ErrInvalidLengthSpiritConjuringPoems
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpiritConjuringPoems
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Poem = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConjuringId", wireType)
			}
			m.ConjuringId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpiritConjuringPoems
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConjuringId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UeblPower", wireType)
			}
			m.UeblPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpiritConjuringPoems
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UeblPower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpiritConjuringPoems
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
				return ErrInvalidLengthSpiritConjuringPoems
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpiritConjuringPoems
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpiritConjuringPoems(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpiritConjuringPoems
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
func skipSpiritConjuringPoems(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpiritConjuringPoems
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
					return 0, ErrIntOverflowSpiritConjuringPoems
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
					return 0, ErrIntOverflowSpiritConjuringPoems
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
				return 0, ErrInvalidLengthSpiritConjuringPoems
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpiritConjuringPoems
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpiritConjuringPoems
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpiritConjuringPoems        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpiritConjuringPoems          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpiritConjuringPoems = fmt.Errorf("proto: unexpected end of group")
)