// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eightball/eightball/meditation_summoning.proto

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

type MeditationSummoning struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MagicKeyId  uint64 `protobuf:"varint,2,opt,name=magicKeyId,proto3" json:"magicKeyId,omitempty"`
	Creator     string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	BlockHeight uint64 `protobuf:"varint,4,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
}

func (m *MeditationSummoning) Reset()         { *m = MeditationSummoning{} }
func (m *MeditationSummoning) String() string { return proto.CompactTextString(m) }
func (*MeditationSummoning) ProtoMessage()    {}
func (*MeditationSummoning) Descriptor() ([]byte, []int) {
	return fileDescriptor_3472eb6362af71b2, []int{0}
}
func (m *MeditationSummoning) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MeditationSummoning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MeditationSummoning.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MeditationSummoning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeditationSummoning.Merge(m, src)
}
func (m *MeditationSummoning) XXX_Size() int {
	return m.Size()
}
func (m *MeditationSummoning) XXX_DiscardUnknown() {
	xxx_messageInfo_MeditationSummoning.DiscardUnknown(m)
}

var xxx_messageInfo_MeditationSummoning proto.InternalMessageInfo

func (m *MeditationSummoning) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MeditationSummoning) GetMagicKeyId() uint64 {
	if m != nil {
		return m.MagicKeyId
	}
	return 0
}

func (m *MeditationSummoning) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MeditationSummoning) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*MeditationSummoning)(nil), "eightball.eightball.MeditationSummoning")
}

func init() {
	proto.RegisterFile("eightball/eightball/meditation_summoning.proto", fileDescriptor_3472eb6362af71b2)
}

var fileDescriptor_3472eb6362af71b2 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcd, 0x4c, 0xcf,
	0x28, 0x49, 0x4a, 0xcc, 0xc9, 0xd1, 0x47, 0xb0, 0x72, 0x53, 0x53, 0x32, 0x4b, 0x12, 0x4b, 0x32,
	0xf3, 0xf3, 0xe2, 0x8b, 0x4b, 0x73, 0x73, 0xf3, 0xf3, 0x32, 0xf3, 0xd2, 0xf5, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0x84, 0xe1, 0xaa, 0x10, 0x3a, 0x95, 0x1a, 0x19, 0xb9, 0x84, 0x7d, 0xe1, 0x7a,
	0x82, 0x61, 0x5a, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82,
	0x98, 0x32, 0x53, 0x84, 0xe4, 0xb8, 0xb8, 0x72, 0x13, 0xd3, 0x33, 0x93, 0xbd, 0x53, 0x2b, 0x3d,
	0x53, 0x24, 0x98, 0xc0, 0xe2, 0x48, 0x22, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0x45, 0xa9, 0x89, 0x25,
	0xf9, 0x45, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30, 0xae, 0x90, 0x02, 0x17, 0x77, 0x52,
	0x4e, 0x7e, 0x72, 0xb6, 0x07, 0xd8, 0x52, 0x09, 0x16, 0xb0, 0x56, 0x64, 0x21, 0x27, 0xd3, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86,
	0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x92, 0x46, 0x78, 0xac, 0x02, 0xc9, 0x93,
	0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x6f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x0a, 0x5b, 0xf5, 0x2d, 0x08, 0x01, 0x00, 0x00,
}

func (m *MeditationSummoning) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MeditationSummoning) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MeditationSummoning) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintMeditationSummoning(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMeditationSummoning(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.MagicKeyId != 0 {
		i = encodeVarintMeditationSummoning(dAtA, i, uint64(m.MagicKeyId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintMeditationSummoning(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMeditationSummoning(dAtA []byte, offset int, v uint64) int {
	offset -= sovMeditationSummoning(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MeditationSummoning) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMeditationSummoning(uint64(m.Id))
	}
	if m.MagicKeyId != 0 {
		n += 1 + sovMeditationSummoning(uint64(m.MagicKeyId))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMeditationSummoning(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovMeditationSummoning(uint64(m.BlockHeight))
	}
	return n
}

func sovMeditationSummoning(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMeditationSummoning(x uint64) (n int) {
	return sovMeditationSummoning(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MeditationSummoning) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMeditationSummoning
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
			return fmt.Errorf("proto: MeditationSummoning: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MeditationSummoning: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeditationSummoning
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
					return ErrIntOverflowMeditationSummoning
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeditationSummoning
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
				return ErrInvalidLengthMeditationSummoning
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMeditationSummoning
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMeditationSummoning
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
		default:
			iNdEx = preIndex
			skippy, err := skipMeditationSummoning(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMeditationSummoning
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
func skipMeditationSummoning(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMeditationSummoning
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
					return 0, ErrIntOverflowMeditationSummoning
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
					return 0, ErrIntOverflowMeditationSummoning
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
				return 0, ErrInvalidLengthMeditationSummoning
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMeditationSummoning
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMeditationSummoning
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMeditationSummoning        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMeditationSummoning          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMeditationSummoning = fmt.Errorf("proto: unexpected end of group")
)