// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/entry.proto

package raftpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Change struct {
	// id of the file that was changed
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// new version of the file
	Version uint64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	// where the change starts (unused for now)
	Offset uint64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	// size of the change
	Size int64 `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	// raft id of the peer who has the latest file
	PeerId               uint64   `protobuf:"varint,5,opt,name=peer_id,json=peerId,proto3" json:"peer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Change) Reset()         { *m = Change{} }
func (m *Change) String() string { return proto.CompactTextString(m) }
func (*Change) ProtoMessage()    {}
func (*Change) Descriptor() ([]byte, []int) {
	return fileDescriptor_a245e8f22934927e, []int{0}
}

func (m *Change) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Change.Unmarshal(m, b)
}
func (m *Change) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Change.Marshal(b, m, deterministic)
}
func (m *Change) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Change.Merge(m, src)
}
func (m *Change) XXX_Size() int {
	return xxx_messageInfo_Change.Size(m)
}
func (m *Change) XXX_DiscardUnknown() {
	xxx_messageInfo_Change.DiscardUnknown(m)
}

var xxx_messageInfo_Change proto.InternalMessageInfo

func (m *Change) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Change) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Change) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Change) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Change) GetPeerId() uint64 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

type Rename struct {
	// id of the renamed file
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// the id of the old parent of the file
	OldParentId uint64 `protobuf:"varint,2,opt,name=old_parent_id,json=oldParentId,proto3" json:"old_parent_id,omitempty"`
	// the id of the new parent of the file
	NewParentId uint64 `protobuf:"varint,3,opt,name=new_parent_id,json=newParentId,proto3" json:"new_parent_id,omitempty"`
	// the new name of the file
	NewName string `protobuf:"bytes,4,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
	// old name
	OldName              string   `protobuf:"bytes,5,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rename) Reset()         { *m = Rename{} }
func (m *Rename) String() string { return proto.CompactTextString(m) }
func (*Rename) ProtoMessage()    {}
func (*Rename) Descriptor() ([]byte, []int) {
	return fileDescriptor_a245e8f22934927e, []int{1}
}

func (m *Rename) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rename.Unmarshal(m, b)
}
func (m *Rename) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rename.Marshal(b, m, deterministic)
}
func (m *Rename) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rename.Merge(m, src)
}
func (m *Rename) XXX_Size() int {
	return xxx_messageInfo_Rename.Size(m)
}
func (m *Rename) XXX_DiscardUnknown() {
	xxx_messageInfo_Rename.DiscardUnknown(m)
}

var xxx_messageInfo_Rename proto.InternalMessageInfo

func (m *Rename) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Rename) GetOldParentId() uint64 {
	if m != nil {
		return m.OldParentId
	}
	return 0
}

func (m *Rename) GetNewParentId() uint64 {
	if m != nil {
		return m.NewParentId
	}
	return 0
}

func (m *Rename) GetNewName() string {
	if m != nil {
		return m.NewName
	}
	return ""
}

func (m *Rename) GetOldName() string {
	if m != nil {
		return m.OldName
	}
	return ""
}

type Delete struct {
	// id of the file to delete
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// id of the file's parent
	ParentId uint64 `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// name of the file
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Delete) Reset()         { *m = Delete{} }
func (m *Delete) String() string { return proto.CompactTextString(m) }
func (*Delete) ProtoMessage()    {}
func (*Delete) Descriptor() ([]byte, []int) {
	return fileDescriptor_a245e8f22934927e, []int{2}
}

func (m *Delete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Delete.Unmarshal(m, b)
}
func (m *Delete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Delete.Marshal(b, m, deterministic)
}
func (m *Delete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delete.Merge(m, src)
}
func (m *Delete) XXX_Size() int {
	return xxx_messageInfo_Delete.Size(m)
}
func (m *Delete) XXX_DiscardUnknown() {
	xxx_messageInfo_Delete.DiscardUnknown(m)
}

var xxx_messageInfo_Delete proto.InternalMessageInfo

func (m *Delete) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Delete) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *Delete) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Add struct {
	// id of file to add
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// id of parent file
	ParentId uint64 `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// file name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// file mode (store.FileMode)
	Mode uint32 `protobuf:"varint,4,opt,name=mode,proto3" json:"mode,omitempty"`
	// if the add is a hard link or genuinely new file
	IsHardLink           bool     `protobuf:"varint,5,opt,name=is_hard_link,json=isHardLink,proto3" json:"is_hard_link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Add) Reset()         { *m = Add{} }
func (m *Add) String() string { return proto.CompactTextString(m) }
func (*Add) ProtoMessage()    {}
func (*Add) Descriptor() ([]byte, []int) {
	return fileDescriptor_a245e8f22934927e, []int{3}
}

func (m *Add) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Add.Unmarshal(m, b)
}
func (m *Add) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Add.Marshal(b, m, deterministic)
}
func (m *Add) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Add.Merge(m, src)
}
func (m *Add) XXX_Size() int {
	return xxx_messageInfo_Add.Size(m)
}
func (m *Add) XXX_DiscardUnknown() {
	xxx_messageInfo_Add.DiscardUnknown(m)
}

var xxx_messageInfo_Add proto.InternalMessageInfo

func (m *Add) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Add) GetParentId() uint64 {
	if m != nil {
		return m.ParentId
	}
	return 0
}

func (m *Add) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Add) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *Add) GetIsHardLink() bool {
	if m != nil {
		return m.IsHardLink
	}
	return false
}

type Entry struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Message:
	//	*Entry_Rename
	//	*Entry_Delete
	//	*Entry_Change
	//	*Entry_Add
	Message              isEntry_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_a245e8f22934927e, []int{4}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type isEntry_Message interface {
	isEntry_Message()
}

type Entry_Rename struct {
	Rename *Rename `protobuf:"bytes,2,opt,name=rename,proto3,oneof"`
}

type Entry_Delete struct {
	Delete *Delete `protobuf:"bytes,3,opt,name=delete,proto3,oneof"`
}

type Entry_Change struct {
	Change *Change `protobuf:"bytes,4,opt,name=change,proto3,oneof"`
}

type Entry_Add struct {
	Add *Add `protobuf:"bytes,5,opt,name=add,proto3,oneof"`
}

func (*Entry_Rename) isEntry_Message() {}

func (*Entry_Delete) isEntry_Message() {}

func (*Entry_Change) isEntry_Message() {}

func (*Entry_Add) isEntry_Message() {}

func (m *Entry) GetMessage() isEntry_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Entry) GetRename() *Rename {
	if x, ok := m.GetMessage().(*Entry_Rename); ok {
		return x.Rename
	}
	return nil
}

func (m *Entry) GetDelete() *Delete {
	if x, ok := m.GetMessage().(*Entry_Delete); ok {
		return x.Delete
	}
	return nil
}

func (m *Entry) GetChange() *Change {
	if x, ok := m.GetMessage().(*Entry_Change); ok {
		return x.Change
	}
	return nil
}

func (m *Entry) GetAdd() *Add {
	if x, ok := m.GetMessage().(*Entry_Add); ok {
		return x.Add
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Entry) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Entry_Rename)(nil),
		(*Entry_Delete)(nil),
		(*Entry_Change)(nil),
		(*Entry_Add)(nil),
	}
}

func init() {
	proto.RegisterType((*Change)(nil), "Change")
	proto.RegisterType((*Rename)(nil), "Rename")
	proto.RegisterType((*Delete)(nil), "Delete")
	proto.RegisterType((*Add)(nil), "Add")
	proto.RegisterType((*Entry)(nil), "Entry")
}

func init() { proto.RegisterFile("pb/entry.proto", fileDescriptor_a245e8f22934927e) }

var fileDescriptor_a245e8f22934927e = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xc1, 0x8e, 0xd3, 0x40,
	0x0c, 0xdd, 0x34, 0x69, 0x92, 0x3a, 0xec, 0x1e, 0xe6, 0x00, 0x59, 0x71, 0x29, 0x39, 0xed, 0xa9,
	0x48, 0xcb, 0x17, 0xec, 0x02, 0x52, 0x2b, 0x21, 0x84, 0xe6, 0xc8, 0x25, 0x9a, 0x62, 0x77, 0x3b,
	0x6a, 0x3a, 0x13, 0x4d, 0x22, 0x22, 0x10, 0xdf, 0xc1, 0x1f, 0xf0, 0x9f, 0xc8, 0x9e, 0x16, 0x21,
	0x7a, 0xe4, 0x66, 0xfb, 0xbd, 0xf1, 0x7b, 0x63, 0x1b, 0x6e, 0xfa, 0xed, 0x6b, 0x72, 0x63, 0xf8,
	0xb6, 0xea, 0x83, 0x1f, 0x7d, 0x33, 0x41, 0xfe, 0x76, 0x6f, 0xdc, 0x13, 0xa9, 0x1b, 0x98, 0x59,
	0xac, 0x93, 0x65, 0x72, 0x97, 0xe9, 0x99, 0x45, 0x55, 0x43, 0xf1, 0x95, 0xc2, 0x60, 0xbd, 0xab,
	0x67, 0x52, 0x3c, 0xa7, 0xea, 0x39, 0xe4, 0x7e, 0xb7, 0x1b, 0x68, 0xac, 0x53, 0x01, 0x4e, 0x99,
	0x52, 0x90, 0x0d, 0xf6, 0x3b, 0xd5, 0xd9, 0x32, 0xb9, 0x4b, 0xb5, 0xc4, 0xea, 0x05, 0x14, 0x3d,
	0x51, 0x68, 0x2d, 0xd6, 0xf3, 0x48, 0xe6, 0x74, 0x83, 0xcd, 0xcf, 0x04, 0x72, 0x4d, 0xce, 0x1c,
	0x2f, 0x95, 0x1b, 0xb8, 0xf6, 0x1d, 0xb6, 0xbd, 0x09, 0xe4, 0x46, 0x7e, 0x19, 0xf5, 0x2b, 0xdf,
	0xe1, 0x27, 0xa9, 0x6d, 0x84, 0xe3, 0x68, 0xfa, 0x8b, 0x13, 0xad, 0x54, 0x8e, 0xa6, 0x3f, 0x9c,
	0x5b, 0x28, 0x99, 0xc3, 0x1a, 0xe2, 0x69, 0xa1, 0x0b, 0x47, 0xd3, 0x47, 0x96, 0xbc, 0x85, 0x92,
	0x25, 0x04, 0x9a, 0x47, 0xc8, 0x77, 0xc8, 0x50, 0xb3, 0x81, 0xfc, 0x1d, 0x75, 0x34, 0x5e, 0xfa,
	0x7a, 0x09, 0x8b, 0x7f, 0x3d, 0x95, 0xfd, 0x59, 0x4c, 0x41, 0x26, 0xdd, 0x52, 0xe9, 0x26, 0x71,
	0xf3, 0x03, 0xd2, 0x07, 0xc4, 0xff, 0xee, 0xc3, 0xb5, 0xa3, 0xc7, 0xf8, 0x89, 0x6b, 0x2d, 0xb1,
	0x5a, 0xc2, 0x33, 0x3b, 0xb4, 0x7b, 0x13, 0xb0, 0xed, 0xac, 0x3b, 0xc8, 0x2f, 0x4a, 0x0d, 0x76,
	0x58, 0x9b, 0x80, 0x1f, 0xac, 0x3b, 0x34, 0xbf, 0x12, 0x98, 0xbf, 0xe7, 0x55, 0x5f, 0x18, 0x78,
	0x05, 0x79, 0x90, 0xd1, 0x8b, 0x7a, 0x75, 0x5f, 0xac, 0xe2, 0x26, 0xd6, 0x57, 0xfa, 0x04, 0x30,
	0x05, 0x65, 0x0a, 0x62, 0x84, 0x29, 0x71, 0x28, 0x4c, 0x89, 0x00, 0x53, 0xbe, 0xc8, 0xe9, 0x88,
	0x2f, 0xa6, 0xc4, 0x4b, 0x62, 0x4a, 0x04, 0x54, 0x0d, 0xa9, 0xc1, 0xb8, 0xf9, 0xea, 0x3e, 0x5b,
	0x3d, 0x20, 0xae, 0xaf, 0x34, 0x97, 0x1e, 0x17, 0x50, 0x1c, 0x69, 0x18, 0xcc, 0x13, 0x3d, 0x96,
	0x9f, 0xf3, 0x60, 0x76, 0x63, 0xbf, 0xdd, 0xe6, 0x72, 0x93, 0x6f, 0x7e, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x7d, 0xcc, 0x5c, 0x27, 0xa5, 0x02, 0x00, 0x00,
}
