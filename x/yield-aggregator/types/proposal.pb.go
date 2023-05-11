// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: yield-aggregator/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// proposal to add new strategy.
type ProposalAddStrategy struct {
	Title           string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description     string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Denom           string `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	ContractAddress string `protobuf:"bytes,4,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Name            string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ProposalAddStrategy) Reset()         { *m = ProposalAddStrategy{} }
func (m *ProposalAddStrategy) String() string { return proto.CompactTextString(m) }
func (*ProposalAddStrategy) ProtoMessage()    {}
func (*ProposalAddStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_de48655730b76ef1, []int{0}
}
func (m *ProposalAddStrategy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalAddStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalAddStrategy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalAddStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalAddStrategy.Merge(m, src)
}
func (m *ProposalAddStrategy) XXX_Size() int {
	return m.Size()
}
func (m *ProposalAddStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalAddStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalAddStrategy proto.InternalMessageInfo

func (m *ProposalAddStrategy) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ProposalAddStrategy) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProposalAddStrategy) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *ProposalAddStrategy) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ProposalAddStrategy) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ProposalAddStrategy)(nil), "ununifi.yieldaggregator.ProposalAddStrategy")
}

func init() { proto.RegisterFile("yield-aggregator/proposal.proto", fileDescriptor_de48655730b76ef1) }

var fileDescriptor_de48655730b76ef1 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xb1, 0x4a, 0x34, 0x31,
	0x14, 0x85, 0x27, 0xff, 0xbf, 0x2b, 0x18, 0x0b, 0x25, 0x2e, 0x38, 0x2c, 0x18, 0x17, 0x2b, 0x2d,
	0x9c, 0x20, 0x3e, 0xc1, 0x5a, 0xd8, 0xd8, 0x88, 0xb2, 0x8d, 0x8d, 0x64, 0x27, 0x31, 0x1b, 0x98,
	0xc9, 0x1d, 0x92, 0x3b, 0xe0, 0xbc, 0x85, 0xcf, 0xe0, 0xd3, 0x58, 0x6e, 0x69, 0x29, 0x33, 0x2f,
	0x22, 0x9b, 0x28, 0x0a, 0xdb, 0xdd, 0x73, 0xce, 0x77, 0x12, 0x0e, 0x3d, 0xe9, 0xac, 0xae, 0xd4,
	0x85, 0x34, 0xc6, 0x6b, 0x23, 0x11, 0xbc, 0x68, 0x3c, 0x34, 0x10, 0x64, 0x55, 0x34, 0x1e, 0x10,
	0xd8, 0x51, 0xeb, 0x5a, 0x67, 0x9f, 0x6d, 0x11, 0xc1, 0x5f, 0x6e, 0x3a, 0x31, 0x60, 0x20, 0x32,
	0x62, 0x73, 0x25, 0x7c, 0x7a, 0xbc, 0xfd, 0x9e, 0xf4, 0xb2, 0x0e, 0x29, 0x3e, 0x7d, 0x23, 0xf4,
	0xf0, 0xee, 0xfb, 0x83, 0xb9, 0x52, 0x0f, 0xe8, 0x25, 0x6a, 0xd3, 0xb1, 0x09, 0x1d, 0xa3, 0xc5,
	0x4a, 0xe7, 0x64, 0x46, 0xce, 0x76, 0xef, 0x93, 0x60, 0x33, 0xba, 0xa7, 0x74, 0x28, 0xbd, 0x6d,
	0xd0, 0x82, 0xcb, 0xff, 0xc5, 0xec, 0xaf, 0xb5, 0xe9, 0x29, 0xed, 0xa0, 0xce, 0xff, 0xa7, 0x5e,
	0x14, 0xec, 0x9c, 0x1e, 0x94, 0xe0, 0xd0, 0xcb, 0x12, 0x9f, 0xa4, 0x52, 0x5e, 0x87, 0x90, 0x8f,
	0x22, 0xb0, 0xff, 0xe3, 0xcf, 0x93, 0xcd, 0x18, 0x1d, 0x39, 0x59, 0xeb, 0x7c, 0x1c, 0xe3, 0x78,
	0x5f, 0xdf, 0xbe, 0xf7, 0x9c, 0xac, 0x7b, 0x4e, 0x3e, 0x7b, 0x4e, 0x5e, 0x07, 0x9e, 0xad, 0x07,
	0x9e, 0x7d, 0x0c, 0x3c, 0x7b, 0xbc, 0x34, 0x16, 0x57, 0xed, 0xb2, 0x28, 0xa1, 0x16, 0x0b, 0xb7,
	0x70, 0xf6, 0xc6, 0x8a, 0x72, 0x25, 0xad, 0x13, 0x2f, 0x62, 0x6b, 0x38, 0x76, 0x8d, 0x0e, 0xcb,
	0x9d, 0x38, 0xfc, 0xea, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x67, 0x2f, 0xee, 0xd7, 0x69, 0x01, 0x00,
	0x00,
}

func (m *ProposalAddStrategy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalAddStrategy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalAddStrategy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposalAddStrategy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposalAddStrategy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: ProposalAddStrategy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalAddStrategy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
