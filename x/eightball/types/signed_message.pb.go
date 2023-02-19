// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eightball/eightball/signed_message.proto

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

type SignedMessage struct {
	Id             uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Body           string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Signature      string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	BitcoinAddress string `protobuf:"bytes,4,opt,name=bitcoinAddress,proto3" json:"bitcoinAddress,omitempty"`
	MessageId      uint64 `protobuf:"varint,5,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Creator        string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	MagicKeyId     uint64 `protobuf:"varint,7,opt,name=magicKeyId,proto3" json:"magicKeyId,omitempty"`
}

func (m *SignedMessage) Reset()         { *m = SignedMessage{} }
func (m *SignedMessage) String() string { return proto.CompactTextString(m) }
func (*SignedMessage) ProtoMessage()    {}
func (*SignedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_26619026f17c52ac, []int{0}
}
func (m *SignedMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignedMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedMessage.Merge(m, src)
}
func (m *SignedMessage) XXX_Size() int {
	return m.Size()
}
func (m *SignedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SignedMessage proto.InternalMessageInfo

func (m *SignedMessage) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SignedMessage) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *SignedMessage) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *SignedMessage) GetBitcoinAddress() string {
	if m != nil {
		return m.BitcoinAddress
	}
	return ""
}

func (m *SignedMessage) GetMessageId() uint64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *SignedMessage) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SignedMessage) GetMagicKeyId() uint64 {
	if m != nil {
		return m.MagicKeyId
	}
	return 0
}

func init() {
	proto.RegisterType((*SignedMessage)(nil), "eightball.eightball.SignedMessage")
}

func init() {
	proto.RegisterFile("eightball/eightball/signed_message.proto", fileDescriptor_26619026f17c52ac)
}

var fileDescriptor_26619026f17c52ac = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xcd, 0x4c, 0xcf,
	0x28, 0x49, 0x4a, 0xcc, 0xc9, 0xd1, 0x47, 0xb0, 0x8a, 0x33, 0xd3, 0xf3, 0x52, 0x53, 0xe2, 0x73,
	0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0xe1, 0xf2,
	0x7a, 0x70, 0x96, 0xd2, 0x45, 0x46, 0x2e, 0xde, 0x60, 0xb0, 0x6a, 0x5f, 0x88, 0x62, 0x21, 0x3e,
	0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x21,
	0x2e, 0x96, 0xa4, 0xfc, 0x94, 0x4a, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48,
	0x86, 0x8b, 0x13, 0x64, 0x45, 0x62, 0x49, 0x69, 0x51, 0xaa, 0x04, 0x33, 0x58, 0x02, 0x21, 0x20,
	0xa4, 0xc6, 0xc5, 0x97, 0x94, 0x59, 0x92, 0x9c, 0x9f, 0x99, 0xe7, 0x98, 0x92, 0x52, 0x94, 0x5a,
	0x5c, 0x2c, 0xc1, 0x02, 0x56, 0x82, 0x26, 0x0a, 0x32, 0x05, 0xea, 0x42, 0xcf, 0x14, 0x09, 0x56,
	0xb0, 0x85, 0x08, 0x01, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4, 0x92, 0xfc, 0x22, 0x09,
	0x36, 0xb0, 0x76, 0x18, 0x57, 0x48, 0x8e, 0x8b, 0x2b, 0x37, 0x31, 0x3d, 0x33, 0xd9, 0x3b, 0xb5,
	0xd2, 0x33, 0x45, 0x82, 0x1d, 0xac, 0x11, 0x49, 0xc4, 0xc9, 0xf4, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xa4, 0x11, 0x41, 0x54, 0x81, 0x14, 0x5c, 0x25, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0xe0, 0x60, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xed, 0x3a, 0x18, 0x7c,
	0x52, 0x01, 0x00, 0x00,
}

func (m *SignedMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignedMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignedMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MagicKeyId != 0 {
		i = encodeVarintSignedMessage(dAtA, i, uint64(m.MagicKeyId))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSignedMessage(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if m.MessageId != 0 {
		i = encodeVarintSignedMessage(dAtA, i, uint64(m.MessageId))
		i--
		dAtA[i] = 0x28
	}
	if len(m.BitcoinAddress) > 0 {
		i -= len(m.BitcoinAddress)
		copy(dAtA[i:], m.BitcoinAddress)
		i = encodeVarintSignedMessage(dAtA, i, uint64(len(m.BitcoinAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintSignedMessage(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintSignedMessage(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSignedMessage(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSignedMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovSignedMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignedMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSignedMessage(uint64(m.Id))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovSignedMessage(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovSignedMessage(uint64(l))
	}
	l = len(m.BitcoinAddress)
	if l > 0 {
		n += 1 + l + sovSignedMessage(uint64(l))
	}
	if m.MessageId != 0 {
		n += 1 + sovSignedMessage(uint64(m.MessageId))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSignedMessage(uint64(l))
	}
	if m.MagicKeyId != 0 {
		n += 1 + sovSignedMessage(uint64(m.MagicKeyId))
	}
	return n
}

func sovSignedMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSignedMessage(x uint64) (n int) {
	return sovSignedMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignedMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignedMessage
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
			return fmt.Errorf("proto: SignedMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignedMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignedMessage
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
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignedMessage
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
				return ErrInvalidLengthSignedMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignedMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignedMessage
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
				return ErrInvalidLengthSignedMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignedMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BitcoinAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignedMessage
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
				return ErrInvalidLengthSignedMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignedMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BitcoinAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageId", wireType)
			}
			m.MessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignedMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignedMessage
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
				return ErrInvalidLengthSignedMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignedMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MagicKeyId", wireType)
			}
			m.MagicKeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignedMessage
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
		default:
			iNdEx = preIndex
			skippy, err := skipSignedMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignedMessage
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
func skipSignedMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSignedMessage
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
					return 0, ErrIntOverflowSignedMessage
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
					return 0, ErrIntOverflowSignedMessage
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
				return 0, ErrInvalidLengthSignedMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSignedMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSignedMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSignedMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSignedMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSignedMessage = fmt.Errorf("proto: unexpected end of group")
)