// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eightball/eightball/verse.proto

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

type Verse struct {
	Id          uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FromSpirit  string   `protobuf:"bytes,2,opt,name=fromSpirit,proto3" json:"fromSpirit,omitempty"`
	ToSpirit    []string `protobuf:"bytes,3,rep,name=toSpirit,proto3" json:"toSpirit,omitempty"`
	WireBytes   string   `protobuf:"bytes,4,opt,name=wireBytes,proto3" json:"wireBytes,omitempty"`
	Broadcast   bool     `protobuf:"varint,5,opt,name=broadcast,proto3" json:"broadcast,omitempty"`
	ToOld       bool     `protobuf:"varint,6,opt,name=toOld,proto3" json:"toOld,omitempty"`
	ToOldAndNew bool     `protobuf:"varint,7,opt,name=toOldAndNew,proto3" json:"toOldAndNew,omitempty"`
	MagicKeyId  uint64   `protobuf:"varint,8,opt,name=magicKeyId,proto3" json:"magicKeyId,omitempty"`
	SummoningId uint64   `protobuf:"varint,9,opt,name=summoningId,proto3" json:"summoningId,omitempty"`
	Creator     string   `protobuf:"bytes,10,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Verse) Reset()         { *m = Verse{} }
func (m *Verse) String() string { return proto.CompactTextString(m) }
func (*Verse) ProtoMessage()    {}
func (*Verse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e53e974f6bf3918c, []int{0}
}
func (m *Verse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Verse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Verse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Verse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Verse.Merge(m, src)
}
func (m *Verse) XXX_Size() int {
	return m.Size()
}
func (m *Verse) XXX_DiscardUnknown() {
	xxx_messageInfo_Verse.DiscardUnknown(m)
}

var xxx_messageInfo_Verse proto.InternalMessageInfo

func (m *Verse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Verse) GetFromSpirit() string {
	if m != nil {
		return m.FromSpirit
	}
	return ""
}

func (m *Verse) GetToSpirit() []string {
	if m != nil {
		return m.ToSpirit
	}
	return nil
}

func (m *Verse) GetWireBytes() string {
	if m != nil {
		return m.WireBytes
	}
	return ""
}

func (m *Verse) GetBroadcast() bool {
	if m != nil {
		return m.Broadcast
	}
	return false
}

func (m *Verse) GetToOld() bool {
	if m != nil {
		return m.ToOld
	}
	return false
}

func (m *Verse) GetToOldAndNew() bool {
	if m != nil {
		return m.ToOldAndNew
	}
	return false
}

func (m *Verse) GetMagicKeyId() uint64 {
	if m != nil {
		return m.MagicKeyId
	}
	return 0
}

func (m *Verse) GetSummoningId() uint64 {
	if m != nil {
		return m.SummoningId
	}
	return 0
}

func (m *Verse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Verse)(nil), "eightball.eightball.Verse")
}

func init() { proto.RegisterFile("eightball/eightball/verse.proto", fileDescriptor_e53e974f6bf3918c) }

var fileDescriptor_e53e974f6bf3918c = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4e, 0x84, 0x30,
	0x14, 0x86, 0x29, 0x33, 0xcc, 0xc0, 0x33, 0x71, 0x51, 0x5d, 0x34, 0x6a, 0x2a, 0x71, 0xc5, 0x6a,
	0x5c, 0x18, 0x0f, 0xe0, 0xec, 0x26, 0x26, 0x9a, 0x60, 0xe2, 0xc2, 0x1d, 0xd0, 0x8a, 0x4d, 0x80,
	0x92, 0x52, 0x1d, 0xb9, 0x85, 0x77, 0xf0, 0x32, 0x2e, 0x67, 0xe9, 0xd2, 0xc0, 0x45, 0x0c, 0xd5,
	0x81, 0xee, 0xfe, 0xff, 0xfb, 0x5f, 0x9b, 0xf7, 0x7e, 0x38, 0xe7, 0x22, 0x7f, 0xd1, 0x69, 0x52,
	0x14, 0x97, 0x93, 0x7a, 0xe3, 0xaa, 0xe1, 0xab, 0x5a, 0x49, 0x2d, 0xf1, 0xd1, 0x88, 0x57, 0xa3,
	0xba, 0xf8, 0x74, 0xc1, 0x7b, 0x1c, 0x86, 0xf0, 0x21, 0xb8, 0x82, 0x11, 0x14, 0xa2, 0x68, 0x1e,
	0xbb, 0x82, 0x61, 0x0a, 0xf0, 0xac, 0x64, 0xf9, 0x50, 0x0b, 0x25, 0x34, 0x71, 0x43, 0x14, 0x05,
	0xb1, 0x45, 0xf0, 0x09, 0xf8, 0x5a, 0xfe, 0xa7, 0xb3, 0x70, 0x16, 0x05, 0xf1, 0xe8, 0xf1, 0x19,
	0x04, 0x5b, 0xa1, 0xf8, 0xba, 0xd5, 0xbc, 0x21, 0x73, 0xf3, 0x74, 0x02, 0x43, 0x9a, 0x2a, 0x99,
	0xb0, 0x2c, 0x69, 0x34, 0xf1, 0x42, 0x14, 0xf9, 0xf1, 0x04, 0xf0, 0x31, 0x78, 0x5a, 0xde, 0x17,
	0x8c, 0x2c, 0x4c, 0xf2, 0x67, 0x70, 0x08, 0x07, 0x46, 0xdc, 0x54, 0xec, 0x8e, 0x6f, 0xc9, 0xd2,
	0x64, 0x36, 0x1a, 0xf6, 0x2d, 0x93, 0x5c, 0x64, 0xb7, 0xbc, 0xdd, 0x30, 0xe2, 0x9b, 0x3b, 0x2c,
	0x32, 0xfc, 0xd0, 0xbc, 0x96, 0xa5, 0xac, 0x44, 0x95, 0x6f, 0x18, 0x09, 0xcc, 0x80, 0x8d, 0x30,
	0x81, 0x65, 0xa6, 0x78, 0xa2, 0xa5, 0x22, 0x60, 0x76, 0xde, 0xdb, 0xf5, 0xf5, 0x57, 0x47, 0xd1,
	0xae, 0xa3, 0xe8, 0xa7, 0xa3, 0xe8, 0xa3, 0xa7, 0xce, 0xae, 0xa7, 0xce, 0x77, 0x4f, 0x9d, 0xa7,
	0xd3, 0xa9, 0xeb, 0x77, 0xab, 0x77, 0xdd, 0xd6, 0xbc, 0x49, 0x17, 0xa6, 0xf8, 0xab, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc4, 0x6b, 0x6c, 0x30, 0x9b, 0x01, 0x00, 0x00,
}

func (m *Verse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Verse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Verse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintVerse(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x52
	}
	if m.SummoningId != 0 {
		i = encodeVarintVerse(dAtA, i, uint64(m.SummoningId))
		i--
		dAtA[i] = 0x48
	}
	if m.MagicKeyId != 0 {
		i = encodeVarintVerse(dAtA, i, uint64(m.MagicKeyId))
		i--
		dAtA[i] = 0x40
	}
	if m.ToOldAndNew {
		i--
		if m.ToOldAndNew {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.ToOld {
		i--
		if m.ToOld {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Broadcast {
		i--
		if m.Broadcast {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if len(m.WireBytes) > 0 {
		i -= len(m.WireBytes)
		copy(dAtA[i:], m.WireBytes)
		i = encodeVarintVerse(dAtA, i, uint64(len(m.WireBytes)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ToSpirit) > 0 {
		for iNdEx := len(m.ToSpirit) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ToSpirit[iNdEx])
			copy(dAtA[i:], m.ToSpirit[iNdEx])
			i = encodeVarintVerse(dAtA, i, uint64(len(m.ToSpirit[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.FromSpirit) > 0 {
		i -= len(m.FromSpirit)
		copy(dAtA[i:], m.FromSpirit)
		i = encodeVarintVerse(dAtA, i, uint64(len(m.FromSpirit)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintVerse(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVerse(dAtA []byte, offset int, v uint64) int {
	offset -= sovVerse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Verse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovVerse(uint64(m.Id))
	}
	l = len(m.FromSpirit)
	if l > 0 {
		n += 1 + l + sovVerse(uint64(l))
	}
	if len(m.ToSpirit) > 0 {
		for _, s := range m.ToSpirit {
			l = len(s)
			n += 1 + l + sovVerse(uint64(l))
		}
	}
	l = len(m.WireBytes)
	if l > 0 {
		n += 1 + l + sovVerse(uint64(l))
	}
	if m.Broadcast {
		n += 2
	}
	if m.ToOld {
		n += 2
	}
	if m.ToOldAndNew {
		n += 2
	}
	if m.MagicKeyId != 0 {
		n += 1 + sovVerse(uint64(m.MagicKeyId))
	}
	if m.SummoningId != 0 {
		n += 1 + sovVerse(uint64(m.SummoningId))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovVerse(uint64(l))
	}
	return n
}

func sovVerse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVerse(x uint64) (n int) {
	return sovVerse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Verse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerse
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
			return fmt.Errorf("proto: Verse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Verse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerse
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromSpirit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerse
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
				return ErrInvalidLengthVerse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromSpirit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToSpirit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerse
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
				return ErrInvalidLengthVerse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToSpirit = append(m.ToSpirit, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WireBytes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerse
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
				return ErrInvalidLengthVerse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WireBytes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Broadcast", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Broadcast = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToOld", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ToOld = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToOldAndNew", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ToOldAndNew = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MagicKeyId", wireType)
			}
			m.MagicKeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerse
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
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SummoningId", wireType)
			}
			m.SummoningId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SummoningId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerse
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
				return ErrInvalidLengthVerse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVerse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVerse
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
func skipVerse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVerse
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
					return 0, ErrIntOverflowVerse
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
					return 0, ErrIntOverflowVerse
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
				return 0, ErrInvalidLengthVerse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVerse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVerse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVerse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVerse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVerse = fmt.Errorf("proto: unexpected end of group")
)