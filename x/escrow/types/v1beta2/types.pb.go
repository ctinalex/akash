// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/escrow/v1beta2/types.proto

package v1beta2

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// State stores state for an escrow account
type Account_State int32

const (
	// AccountStateInvalid is an invalid state
	AccountStateInvalid Account_State = 0
	// AccountOpen is the state when an account is open
	AccountOpen Account_State = 1
	// AccountClosed is the state when an account is closed
	AccountClosed Account_State = 2
	// AccountOverdrawn is the state when an account is overdrawn
	AccountOverdrawn Account_State = 3
)

var Account_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "closed",
	3: "overdrawn",
}

var Account_State_value = map[string]int32{
	"invalid":   0,
	"open":      1,
	"closed":    2,
	"overdrawn": 3,
}

func (x Account_State) String() string {
	return proto.EnumName(Account_State_name, int32(x))
}

func (Account_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5b25ee303c78038d, []int{1, 0}
}

// Payment State
type FractionalPayment_State int32

const (
	// PaymentStateInvalid is the state when the payment is invalid
	PaymentStateInvalid FractionalPayment_State = 0
	// PaymentStateOpen is the state when the payment is open
	PaymentOpen FractionalPayment_State = 1
	// PaymentStateClosed is the state when the payment is closed
	PaymentClosed FractionalPayment_State = 2
	// PaymentStateOverdrawn is the state when the payment is overdrawn
	PaymentOverdrawn FractionalPayment_State = 3
)

var FractionalPayment_State_name = map[int32]string{
	0: "invalid",
	1: "open",
	2: "closed",
	3: "overdrawn",
}

var FractionalPayment_State_value = map[string]int32{
	"invalid":   0,
	"open":      1,
	"closed":    2,
	"overdrawn": 3,
}

func (x FractionalPayment_State) String() string {
	return proto.EnumName(FractionalPayment_State_name, int32(x))
}

func (FractionalPayment_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5b25ee303c78038d, []int{2, 0}
}

// AccountID is the account identifier
type AccountID struct {
	Scope string `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope" yaml:"scope"`
	XID   string `protobuf:"bytes,2,opt,name=xid,proto3" json:"xid" yaml:"xid"`
}

func (m *AccountID) Reset()         { *m = AccountID{} }
func (m *AccountID) String() string { return proto.CompactTextString(m) }
func (*AccountID) ProtoMessage()    {}
func (*AccountID) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b25ee303c78038d, []int{0}
}
func (m *AccountID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountID.Merge(m, src)
}
func (m *AccountID) XXX_Size() int {
	return m.Size()
}
func (m *AccountID) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountID.DiscardUnknown(m)
}

var xxx_messageInfo_AccountID proto.InternalMessageInfo

func (m *AccountID) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *AccountID) GetXID() string {
	if m != nil {
		return m.XID
	}
	return ""
}

// Account stores state for an escrow account
type Account struct {
	// unique identifier for this escrow account
	ID AccountID `protobuf:"bytes,1,opt,name=id,proto3" json:"id" yaml:"id"`
	// bech32 encoded account address of the owner of this escrow account
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	// current state of this escrow account
	State Account_State `protobuf:"varint,3,opt,name=state,proto3,enum=akash.escrow.v1beta2.Account_State" json:"state" yaml:"state"`
	// unspent coins received from the owner's wallet
	Balance types.DecCoin `protobuf:"bytes,4,opt,name=balance,proto3" json:"balance" yaml:"balance"`
	// total coins spent by this account
	Transferred types.DecCoin `protobuf:"bytes,5,opt,name=transferred,proto3" json:"transferred" yaml:"transferred"`
	// block height at which this account was last settled
	SettledAt int64 `protobuf:"varint,6,opt,name=settled_at,json=settledAt,proto3" json:"settledAt" yaml:"settledAt"`
	// bech32 encoded account address of the depositor.
	// If depositor is same as the owner, then any incoming coins are added to the Balance.
	// If depositor isn't same as the owner, then any incoming coins are added to the Funds.
	Depositor string `protobuf:"bytes,7,opt,name=depositor,proto3" json:"depositor" yaml:"depositor"`
	// Funds are unspent coins received from the (non-Owner) Depositor's wallet.
	// If there are any funds, they should be spent before spending the Balance.
	Funds types.DecCoin `protobuf:"bytes,8,opt,name=funds,proto3" json:"funds" yaml:"funds"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b25ee303c78038d, []int{1}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Account.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return m.Size()
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetID() AccountID {
	if m != nil {
		return m.ID
	}
	return AccountID{}
}

func (m *Account) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Account) GetState() Account_State {
	if m != nil {
		return m.State
	}
	return AccountStateInvalid
}

func (m *Account) GetBalance() types.DecCoin {
	if m != nil {
		return m.Balance
	}
	return types.DecCoin{}
}

func (m *Account) GetTransferred() types.DecCoin {
	if m != nil {
		return m.Transferred
	}
	return types.DecCoin{}
}

func (m *Account) GetSettledAt() int64 {
	if m != nil {
		return m.SettledAt
	}
	return 0
}

func (m *Account) GetDepositor() string {
	if m != nil {
		return m.Depositor
	}
	return ""
}

func (m *Account) GetFunds() types.DecCoin {
	if m != nil {
		return m.Funds
	}
	return types.DecCoin{}
}

// Payment stores state for a payment
type FractionalPayment struct {
	AccountID AccountID               `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"accountID" yaml:"accountID"`
	PaymentID string                  `protobuf:"bytes,2,opt,name=payment_id,json=paymentId,proto3" json:"paymentID" yaml:"paymentID"`
	Owner     string                  `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner" yaml:"owner"`
	State     FractionalPayment_State `protobuf:"varint,4,opt,name=state,proto3,enum=akash.escrow.v1beta2.FractionalPayment_State" json:"state" yaml:"state"`
	Rate      types.DecCoin           `protobuf:"bytes,5,opt,name=rate,proto3" json:"rate" yaml:"rate"`
	Balance   types.DecCoin           `protobuf:"bytes,6,opt,name=balance,proto3" json:"balance" yaml:"balance"`
	Withdrawn types.Coin              `protobuf:"bytes,7,opt,name=withdrawn,proto3" json:"withdrawn" yaml:"withdrawn"`
}

func (m *FractionalPayment) Reset()         { *m = FractionalPayment{} }
func (m *FractionalPayment) String() string { return proto.CompactTextString(m) }
func (*FractionalPayment) ProtoMessage()    {}
func (*FractionalPayment) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b25ee303c78038d, []int{2}
}
func (m *FractionalPayment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FractionalPayment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FractionalPayment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FractionalPayment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FractionalPayment.Merge(m, src)
}
func (m *FractionalPayment) XXX_Size() int {
	return m.Size()
}
func (m *FractionalPayment) XXX_DiscardUnknown() {
	xxx_messageInfo_FractionalPayment.DiscardUnknown(m)
}

var xxx_messageInfo_FractionalPayment proto.InternalMessageInfo

func (m *FractionalPayment) GetAccountID() AccountID {
	if m != nil {
		return m.AccountID
	}
	return AccountID{}
}

func (m *FractionalPayment) GetPaymentID() string {
	if m != nil {
		return m.PaymentID
	}
	return ""
}

func (m *FractionalPayment) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *FractionalPayment) GetState() FractionalPayment_State {
	if m != nil {
		return m.State
	}
	return PaymentStateInvalid
}

func (m *FractionalPayment) GetRate() types.DecCoin {
	if m != nil {
		return m.Rate
	}
	return types.DecCoin{}
}

func (m *FractionalPayment) GetBalance() types.DecCoin {
	if m != nil {
		return m.Balance
	}
	return types.DecCoin{}
}

func (m *FractionalPayment) GetWithdrawn() types.Coin {
	if m != nil {
		return m.Withdrawn
	}
	return types.Coin{}
}

func init() {
	proto.RegisterEnum("akash.escrow.v1beta2.Account_State", Account_State_name, Account_State_value)
	proto.RegisterEnum("akash.escrow.v1beta2.FractionalPayment_State", FractionalPayment_State_name, FractionalPayment_State_value)
	proto.RegisterType((*AccountID)(nil), "akash.escrow.v1beta2.AccountID")
	proto.RegisterType((*Account)(nil), "akash.escrow.v1beta2.Account")
	proto.RegisterType((*FractionalPayment)(nil), "akash.escrow.v1beta2.FractionalPayment")
}

func init() { proto.RegisterFile("akash/escrow/v1beta2/types.proto", fileDescriptor_5b25ee303c78038d) }

var fileDescriptor_5b25ee303c78038d = []byte{
	// 805 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xe3, 0x44,
	0x18, 0x8e, 0xf3, 0x89, 0x27, 0xb0, 0x64, 0x87, 0x4a, 0xb8, 0x61, 0xeb, 0xf1, 0xce, 0x82, 0x54,
	0x0e, 0xd8, 0x6a, 0xb8, 0xed, 0x05, 0x6d, 0xb6, 0x42, 0xf4, 0xb0, 0x7c, 0xb8, 0x1c, 0x80, 0x43,
	0xab, 0x89, 0x3d, 0x6d, 0xad, 0x3a, 0x1e, 0xcb, 0x9e, 0x26, 0xed, 0x3f, 0x40, 0x11, 0x07, 0xc4,
	0x89, 0x4b, 0x10, 0x12, 0x7f, 0xa6, 0xc7, 0x1e, 0x39, 0x59, 0x28, 0xbd, 0xe5, 0x98, 0x5f, 0x80,
	0x3c, 0x33, 0xb6, 0x83, 0x88, 0xa2, 0x1c, 0xf6, 0x64, 0xbf, 0xcf, 0xfb, 0xbe, 0xcf, 0xbc, 0x1f,
	0xf3, 0x68, 0x80, 0x45, 0xae, 0x49, 0x7a, 0xe5, 0xd0, 0xd4, 0x4b, 0xd8, 0xd4, 0x99, 0x1c, 0x8d,
	0x28, 0x27, 0x03, 0x87, 0xdf, 0xc5, 0x34, 0xb5, 0xe3, 0x84, 0x71, 0x06, 0xf7, 0x44, 0x84, 0x2d,
	0x23, 0x6c, 0x15, 0xd1, 0xdf, 0xbb, 0x64, 0x97, 0x4c, 0x04, 0x38, 0xf9, 0x9f, 0x8c, 0xed, 0x9b,
	0x1e, 0x4b, 0xc7, 0x2c, 0x75, 0x46, 0x24, 0xa5, 0x8a, 0xec, 0xc8, 0xf1, 0x58, 0x10, 0x49, 0x3f,
	0x0e, 0x81, 0xfe, 0xca, 0xf3, 0xd8, 0x4d, 0xc4, 0x4f, 0x8e, 0xa1, 0x03, 0x5a, 0xa9, 0xc7, 0x62,
	0x6a, 0x68, 0x96, 0x76, 0xa8, 0x0f, 0xf7, 0x97, 0x19, 0x92, 0xc0, 0x2a, 0x43, 0xef, 0xde, 0x91,
	0x71, 0xf8, 0x12, 0x0b, 0x13, 0xbb, 0x12, 0x86, 0x36, 0x68, 0xdc, 0x06, 0xbe, 0x51, 0x17, 0xe1,
	0xcf, 0x16, 0x19, 0x6a, 0xfc, 0x70, 0x72, 0xbc, 0xcc, 0x50, 0x8e, 0xae, 0x32, 0x04, 0x64, 0xce,
	0x6d, 0xe0, 0x63, 0x37, 0x87, 0xf0, 0x1f, 0x6d, 0xd0, 0x51, 0xc7, 0xc1, 0xaf, 0x41, 0x3d, 0xf0,
	0xc5, 0x49, 0xdd, 0x01, 0xb2, 0x37, 0xb5, 0x64, 0x97, 0x95, 0x0d, 0x0f, 0xee, 0x33, 0x54, 0x5b,
	0x64, 0xa8, 0x2e, 0xe8, 0xeb, 0x82, 0x5d, 0x97, 0xec, 0x39, 0x79, 0x3d, 0xf0, 0xf3, 0xe2, 0xd9,
	0x34, 0xa2, 0x89, 0xaa, 0x46, 0x14, 0x2f, 0x80, 0xaa, 0x78, 0x61, 0x62, 0x57, 0xc2, 0xf0, 0x7b,
	0xd0, 0x4a, 0x39, 0xe1, 0xd4, 0x68, 0x58, 0xda, 0xe1, 0x93, 0xc1, 0x8b, 0xad, 0x35, 0xd8, 0xa7,
	0x79, 0xa8, 0x1a, 0x49, 0xfe, 0xbb, 0x36, 0x92, 0xdc, 0xcc, 0x47, 0x92, 0x7f, 0xe1, 0x8f, 0xa0,
	0x33, 0x22, 0x21, 0x89, 0x3c, 0x6a, 0x34, 0x45, 0x6f, 0xcf, 0x6c, 0xb9, 0x02, 0x3b, 0x5f, 0x81,
	0xa2, 0x3d, 0xb2, 0x8f, 0xa9, 0xf7, 0x9a, 0x05, 0xd1, 0xf0, 0x79, 0xde, 0xd8, 0x32, 0x43, 0x45,
	0xd2, 0x2a, 0x43, 0x4f, 0x24, 0xad, 0x02, 0xb0, 0x5b, 0xb8, 0x60, 0x00, 0xba, 0x3c, 0x21, 0x51,
	0x7a, 0x41, 0x93, 0x84, 0xfa, 0x46, 0x6b, 0x07, 0xfa, 0x4f, 0x15, 0xfd, 0x7a, 0xe2, 0x2a, 0x43,
	0x50, 0x1e, 0xb1, 0x06, 0x62, 0x77, 0x3d, 0x04, 0xbe, 0x01, 0x20, 0xa5, 0x9c, 0x87, 0xd4, 0x3f,
	0x27, 0xdc, 0x68, 0x5b, 0xda, 0x61, 0x63, 0x68, 0x2f, 0x32, 0xa4, 0x9f, 0x4a, 0xf4, 0x15, 0x5f,
	0x66, 0x48, 0x4f, 0x0b, 0x63, 0x95, 0xa1, 0x9e, 0x1a, 0x46, 0x01, 0x61, 0xb7, 0x72, 0xc3, 0x2f,
	0x80, 0xee, 0xd3, 0x98, 0xa5, 0x01, 0x67, 0x89, 0xd1, 0x11, 0xfb, 0x79, 0x9e, 0x13, 0x94, 0x60,
	0x45, 0x50, 0x42, 0xd8, 0xad, 0xdc, 0xf0, 0x3b, 0xd0, 0xba, 0xb8, 0x89, 0xfc, 0xd4, 0x78, 0x67,
	0x87, 0xa6, 0x0f, 0x54, 0xd3, 0x32, 0xa5, 0x5a, 0x94, 0x30, 0xb1, 0x2b, 0x61, 0xfc, 0x9b, 0x06,
	0x5a, 0x62, 0xa9, 0xf0, 0x63, 0xd0, 0x09, 0xa2, 0x09, 0x09, 0x03, 0xbf, 0x57, 0xeb, 0x7f, 0x38,
	0x9b, 0x5b, 0x1f, 0xa8, 0xa5, 0x0b, 0xf7, 0x89, 0x74, 0xc1, 0x7d, 0xd0, 0x64, 0x31, 0x8d, 0x7a,
	0x5a, 0xff, 0xfd, 0xd9, 0xdc, 0xea, 0xaa, 0x90, 0x6f, 0x62, 0x1a, 0xc1, 0x03, 0xd0, 0xf6, 0x42,
	0x96, 0x52, 0xbf, 0x57, 0xef, 0x3f, 0x9d, 0xcd, 0xad, 0xf7, 0x94, 0xf3, 0xb5, 0x00, 0xe1, 0x0b,
	0xa0, 0xb3, 0x09, 0x4d, 0xfc, 0x84, 0x4c, 0xa3, 0x5e, 0xa3, 0xbf, 0x37, 0x9b, 0x5b, 0xbd, 0x22,
	0xbd, 0xc0, 0xfb, 0xcd, 0x9f, 0xff, 0x32, 0x6b, 0xf8, 0x97, 0x36, 0x78, 0xfa, 0x65, 0x42, 0x3c,
	0x1e, 0xb0, 0x88, 0x84, 0xdf, 0x92, 0xbb, 0x31, 0x8d, 0x38, 0x4c, 0x00, 0x20, 0x32, 0xfe, 0x7c,
	0x77, 0xc9, 0x0c, 0x94, 0x64, 0x2a, 0x7d, 0xe7, 0x13, 0x27, 0x85, 0x51, 0x4d, 0xbc, 0x84, 0xb0,
	0x5b, 0xba, 0xc5, 0x0d, 0x88, 0xe5, 0xf1, 0xe7, 0xa5, 0xc2, 0xc5, 0x0d, 0x50, 0x45, 0x49, 0xba,
	0xb8, 0x30, 0x2a, 0xba, 0x12, 0xc2, 0x6e, 0xe9, 0x5e, 0x53, 0x67, 0x63, 0x47, 0x75, 0x9e, 0x15,
	0xea, 0x6c, 0x0a, 0x75, 0x7e, 0xb6, 0xb9, 0xdd, 0xff, 0xcd, 0x6a, 0x67, 0x9d, 0xbe, 0x01, 0xcd,
	0x24, 0xa7, 0xdf, 0x45, 0x45, 0x1f, 0xa9, 0x0b, 0x25, 0x32, 0x56, 0x19, 0xea, 0x4a, 0xc2, 0x44,
	0xf0, 0x09, 0x70, 0x5d, 0xf6, 0xed, 0xb7, 0x2c, 0xfb, 0x33, 0xa0, 0x4f, 0x03, 0x7e, 0x25, 0xae,
	0x89, 0x10, 0x4f, 0x77, 0xb0, 0xbf, 0x91, 0x5c, 0x30, 0x7f, 0xa2, 0x98, 0xab, 0x9c, 0x6a, 0x35,
	0x25, 0x84, 0xdd, 0xca, 0xbd, 0x55, 0x08, 0x6a, 0xaa, 0xdb, 0x84, 0xa0, 0x42, 0x36, 0x0b, 0x41,
	0x39, 0xb7, 0x08, 0xa1, 0x48, 0xff, 0xaf, 0x10, 0x5e, 0x36, 0x7f, 0xff, 0x13, 0x69, 0xc3, 0xaf,
	0xee, 0x17, 0xa6, 0xf6, 0xb0, 0x30, 0xb5, 0x7f, 0x16, 0xa6, 0xf6, 0xeb, 0xa3, 0x59, 0x7b, 0x78,
	0x34, 0x6b, 0x7f, 0x3f, 0x9a, 0xb5, 0x9f, 0xec, 0xcb, 0x80, 0x5f, 0xdd, 0x8c, 0x6c, 0x8f, 0x8d,
	0x1d, 0x36, 0x49, 0xbc, 0xf0, 0xda, 0x91, 0xef, 0xe6, 0x6d, 0xf1, 0x72, 0x8a, 0x17, 0xb3, 0x78,
	0x3f, 0x47, 0x6d, 0xf1, 0xdc, 0x7d, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0x1a, 0x5d,
	0xc3, 0x5e, 0x07, 0x00, 0x00,
}

func (m *AccountID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.XID) > 0 {
		i -= len(m.XID)
		copy(dAtA[i:], m.XID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.XID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Scope) > 0 {
		i -= len(m.Scope)
		copy(dAtA[i:], m.Scope)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Scope)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Account) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Account) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Account) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Funds.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x3a
	}
	if m.SettledAt != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.SettledAt))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.Transferred.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.State != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.ID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *FractionalPayment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FractionalPayment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FractionalPayment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Withdrawn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Rate.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.State != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PaymentID) > 0 {
		i -= len(m.PaymentID)
		copy(dAtA[i:], m.PaymentID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.PaymentID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.AccountID.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Scope)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.XID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Account) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ID.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovTypes(uint64(m.State))
	}
	l = m.Balance.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Transferred.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.SettledAt != 0 {
		n += 1 + sovTypes(uint64(m.SettledAt))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Funds.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *FractionalPayment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AccountID.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.PaymentID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovTypes(uint64(m.State))
	}
	l = m.Rate.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Balance.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Withdrawn.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: AccountID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Scope", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Scope = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Account) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Account: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Account: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Account_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transferred", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Transferred.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SettledAt", wireType)
			}
			m.SettledAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SettledAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Funds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Funds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *FractionalPayment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: FractionalPayment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FractionalPayment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AccountID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PaymentID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= FractionalPayment_State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdrawn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Withdrawn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)