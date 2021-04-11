// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: synchronizer/syncOperation.proto

package types

import (
	fmt "fmt"
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

type SyncOperation struct {
	Op       string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	Uuid     string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Key      string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value    []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	Bookmark uint64 `protobuf:"varint,5,opt,name=bookmark,proto3" json:"bookmark,omitempty"`
	Creator  string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *SyncOperation) Reset()         { *m = SyncOperation{} }
func (m *SyncOperation) String() string { return proto.CompactTextString(m) }
func (*SyncOperation) ProtoMessage()    {}
func (*SyncOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a9812b73830535, []int{0}
}
func (m *SyncOperation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncOperation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SyncOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncOperation.Merge(m, src)
}
func (m *SyncOperation) XXX_Size() int {
	return m.Size()
}
func (m *SyncOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncOperation.DiscardUnknown(m)
}

var xxx_messageInfo_SyncOperation proto.InternalMessageInfo

func (m *SyncOperation) GetOp() string {
	if m != nil {
		return m.Op
	}
	return ""
}

func (m *SyncOperation) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *SyncOperation) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SyncOperation) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SyncOperation) GetBookmark() uint64 {
	if m != nil {
		return m.Bookmark
	}
	return 0
}

func (m *SyncOperation) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*SyncOperation)(nil), "bluzelle.curium.synchronizer.SyncOperation")
}

func init() { proto.RegisterFile("synchronizer/syncOperation.proto", fileDescriptor_73a9812b73830535) }

var fileDescriptor_73a9812b73830535 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0xe3, 0x34, 0x2d, 0x60, 0x01, 0x42, 0x56, 0x07, 0xab, 0x42, 0x56, 0xc4, 0x94, 0x29,
	0x1e, 0x78, 0x03, 0x36, 0x26, 0xa4, 0xb0, 0xb1, 0x25, 0xc1, 0x4a, 0xad, 0xfc, 0x5c, 0xeb, 0xd6,
	0x46, 0xa4, 0x4f, 0x01, 0x6f, 0xc5, 0xd8, 0x91, 0x11, 0x25, 0x2f, 0x82, 0xe2, 0xa8, 0xa8, 0x6c,
	0xe7, 0x3b, 0xba, 0xe7, 0x48, 0xe7, 0xd2, 0x78, 0xd7, 0x77, 0xe5, 0x16, 0xa1, 0xd3, 0x7b, 0x85,
	0x72, 0x82, 0x27, 0xa3, 0x30, 0xb7, 0x1a, 0xba, 0xd4, 0x20, 0x58, 0x60, 0xb7, 0x45, 0xe3, 0xf6,
	0xaa, 0x69, 0x54, 0x5a, 0x3a, 0xd4, 0xae, 0x4d, 0x4f, 0x13, 0x9b, 0x75, 0x05, 0x15, 0xf8, 0x43,
	0x39, 0xa9, 0x39, 0x73, 0xf7, 0x49, 0xe8, 0xd5, 0xf3, 0x69, 0x17, 0xbb, 0xa6, 0x21, 0x18, 0x4e,
	0x62, 0x92, 0x5c, 0x64, 0x21, 0x18, 0xc6, 0x68, 0xe4, 0x9c, 0x7e, 0xe5, 0xa1, 0x77, 0xbc, 0x66,
	0x37, 0x74, 0x51, 0xab, 0x9e, 0x2f, 0xbc, 0x35, 0x49, 0xb6, 0xa6, 0xcb, 0xb7, 0xbc, 0x71, 0x8a,
	0x47, 0x31, 0x49, 0x2e, 0xb3, 0x19, 0xd8, 0x86, 0x9e, 0x17, 0x00, 0x75, 0x9b, 0x63, 0xcd, 0x97,
	0x31, 0x49, 0xa2, 0xec, 0x8f, 0x19, 0xa7, 0x67, 0x25, 0xaa, 0xdc, 0x02, 0xf2, 0x95, 0xef, 0x39,
	0xe2, 0xc3, 0xe3, 0xd7, 0x20, 0xc8, 0x61, 0x10, 0xe4, 0x67, 0x10, 0xe4, 0x63, 0x14, 0xc1, 0x61,
	0x14, 0xc1, 0xf7, 0x28, 0x82, 0x17, 0x59, 0x69, 0xbb, 0x75, 0x45, 0x5a, 0x42, 0x2b, 0x8f, 0x63,
	0xe5, 0x3c, 0x56, 0xbe, 0xcb, 0x7f, 0x0f, 0xb2, 0xbd, 0x51, 0xbb, 0x62, 0xe5, 0x57, 0xde, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xf1, 0x80, 0xf9, 0x3d, 0x01, 0x00, 0x00,
}

func (m *SyncOperation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncOperation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SyncOperation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSyncOperation(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if m.Bookmark != 0 {
		i = encodeVarintSyncOperation(dAtA, i, uint64(m.Bookmark))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintSyncOperation(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintSyncOperation(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Uuid) > 0 {
		i -= len(m.Uuid)
		copy(dAtA[i:], m.Uuid)
		i = encodeVarintSyncOperation(dAtA, i, uint64(len(m.Uuid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Op) > 0 {
		i -= len(m.Op)
		copy(dAtA[i:], m.Op)
		i = encodeVarintSyncOperation(dAtA, i, uint64(len(m.Op)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSyncOperation(dAtA []byte, offset int, v uint64) int {
	offset -= sovSyncOperation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SyncOperation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Op)
	if l > 0 {
		n += 1 + l + sovSyncOperation(uint64(l))
	}
	l = len(m.Uuid)
	if l > 0 {
		n += 1 + l + sovSyncOperation(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovSyncOperation(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovSyncOperation(uint64(l))
	}
	if m.Bookmark != 0 {
		n += 1 + sovSyncOperation(uint64(m.Bookmark))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSyncOperation(uint64(l))
	}
	return n
}

func sovSyncOperation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSyncOperation(x uint64) (n int) {
	return sovSyncOperation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SyncOperation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSyncOperation
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
			return fmt.Errorf("proto: SyncOperation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncOperation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncOperation
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
				return ErrInvalidLengthSyncOperation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSyncOperation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Op = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uuid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncOperation
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
				return ErrInvalidLengthSyncOperation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSyncOperation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uuid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncOperation
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
				return ErrInvalidLengthSyncOperation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSyncOperation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncOperation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSyncOperation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSyncOperation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bookmark", wireType)
			}
			m.Bookmark = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSyncOperation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Bookmark |= uint64(b&0x7F) << shift
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
					return ErrIntOverflowSyncOperation
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
				return ErrInvalidLengthSyncOperation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSyncOperation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSyncOperation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSyncOperation
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
func skipSyncOperation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSyncOperation
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
					return 0, ErrIntOverflowSyncOperation
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
					return 0, ErrIntOverflowSyncOperation
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
				return 0, ErrInvalidLengthSyncOperation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSyncOperation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSyncOperation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSyncOperation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSyncOperation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSyncOperation = fmt.Errorf("proto: unexpected end of group")
)
