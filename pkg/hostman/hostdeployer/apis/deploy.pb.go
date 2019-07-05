// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deploy.proto

package apis

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GuestDesc struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Domain               string   `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	Nics                 []*Nic   `protobuf:"bytes,4,rep,name=nics,proto3" json:"nics,omitempty"`
	NicsStandby          []*Nic   `protobuf:"bytes,5,rep,name=nics_standby,json=nicsStandby,proto3" json:"nics_standby,omitempty"`
	Disks                []*Disk  `protobuf:"bytes,6,rep,name=disks,proto3" json:"disks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GuestDesc) Reset()         { *m = GuestDesc{} }
func (m *GuestDesc) String() string { return proto.CompactTextString(m) }
func (*GuestDesc) ProtoMessage()    {}
func (*GuestDesc) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{0}
}

func (m *GuestDesc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuestDesc.Unmarshal(m, b)
}
func (m *GuestDesc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuestDesc.Marshal(b, m, deterministic)
}
func (m *GuestDesc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuestDesc.Merge(m, src)
}
func (m *GuestDesc) XXX_Size() int {
	return xxx_messageInfo_GuestDesc.Size(m)
}
func (m *GuestDesc) XXX_DiscardUnknown() {
	xxx_messageInfo_GuestDesc.DiscardUnknown(m)
}

var xxx_messageInfo_GuestDesc proto.InternalMessageInfo

func (m *GuestDesc) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GuestDesc) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *GuestDesc) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *GuestDesc) GetNics() []*Nic {
	if m != nil {
		return m.Nics
	}
	return nil
}

func (m *GuestDesc) GetNicsStandby() []*Nic {
	if m != nil {
		return m.NicsStandby
	}
	return nil
}

func (m *GuestDesc) GetDisks() []*Disk {
	if m != nil {
		return m.Disks
	}
	return nil
}

type Disk struct {
	DiskId               string   `protobuf:"bytes,1,opt,name=disk_id,json=diskId,proto3" json:"disk_id,omitempty"`
	Driver               string   `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	CacheMode            string   `protobuf:"bytes,3,opt,name=cache_mode,json=cacheMode,proto3" json:"cache_mode,omitempty"`
	AioMode              string   `protobuf:"bytes,4,opt,name=aio_mode,json=aioMode,proto3" json:"aio_mode,omitempty"`
	Size                 int64    `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	TemplateId           string   `protobuf:"bytes,6,opt,name=template_id,json=templateId,proto3" json:"template_id,omitempty"`
	ImagePath            string   `protobuf:"bytes,7,opt,name=image_path,json=imagePath,proto3" json:"image_path,omitempty"`
	StorageId            string   `protobuf:"bytes,8,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
	Migrating            bool     `protobuf:"varint,9,opt,name=migrating,proto3" json:"migrating,omitempty"`
	TargetStorageId      string   `protobuf:"bytes,10,opt,name=target_storage_id,json=targetStorageId,proto3" json:"target_storage_id,omitempty"`
	Path                 string   `protobuf:"bytes,11,opt,name=path,proto3" json:"path,omitempty"`
	Format               string   `protobuf:"bytes,12,opt,name=format,proto3" json:"format,omitempty"`
	Index                int32    `protobuf:"varint,13,opt,name=index,proto3" json:"index,omitempty"`
	MergeSnapshot        bool     `protobuf:"varint,14,opt,name=merge_snapshot,json=mergeSnapshot,proto3" json:"merge_snapshot,omitempty"`
	Fs                   string   `protobuf:"bytes,15,opt,name=fs,proto3" json:"fs,omitempty"`
	Mountpoint           string   `protobuf:"bytes,16,opt,name=mountpoint,proto3" json:"mountpoint,omitempty"`
	Dev                  string   `protobuf:"bytes,17,opt,name=dev,proto3" json:"dev,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Disk) Reset()         { *m = Disk{} }
func (m *Disk) String() string { return proto.CompactTextString(m) }
func (*Disk) ProtoMessage()    {}
func (*Disk) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{1}
}

func (m *Disk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Disk.Unmarshal(m, b)
}
func (m *Disk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Disk.Marshal(b, m, deterministic)
}
func (m *Disk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Disk.Merge(m, src)
}
func (m *Disk) XXX_Size() int {
	return xxx_messageInfo_Disk.Size(m)
}
func (m *Disk) XXX_DiscardUnknown() {
	xxx_messageInfo_Disk.DiscardUnknown(m)
}

var xxx_messageInfo_Disk proto.InternalMessageInfo

func (m *Disk) GetDiskId() string {
	if m != nil {
		return m.DiskId
	}
	return ""
}

func (m *Disk) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Disk) GetCacheMode() string {
	if m != nil {
		return m.CacheMode
	}
	return ""
}

func (m *Disk) GetAioMode() string {
	if m != nil {
		return m.AioMode
	}
	return ""
}

func (m *Disk) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Disk) GetTemplateId() string {
	if m != nil {
		return m.TemplateId
	}
	return ""
}

func (m *Disk) GetImagePath() string {
	if m != nil {
		return m.ImagePath
	}
	return ""
}

func (m *Disk) GetStorageId() string {
	if m != nil {
		return m.StorageId
	}
	return ""
}

func (m *Disk) GetMigrating() bool {
	if m != nil {
		return m.Migrating
	}
	return false
}

func (m *Disk) GetTargetStorageId() string {
	if m != nil {
		return m.TargetStorageId
	}
	return ""
}

func (m *Disk) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Disk) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *Disk) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Disk) GetMergeSnapshot() bool {
	if m != nil {
		return m.MergeSnapshot
	}
	return false
}

func (m *Disk) GetFs() string {
	if m != nil {
		return m.Fs
	}
	return ""
}

func (m *Disk) GetMountpoint() string {
	if m != nil {
		return m.Mountpoint
	}
	return ""
}

func (m *Disk) GetDev() string {
	if m != nil {
		return m.Dev
	}
	return ""
}

type Nic struct {
	Mac                  string    `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	Ip                   string    `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Net                  string    `protobuf:"bytes,3,opt,name=net,proto3" json:"net,omitempty"`
	NetId                string    `protobuf:"bytes,4,opt,name=net_id,json=netId,proto3" json:"net_id,omitempty"`
	Virtual              bool      `protobuf:"varint,5,opt,name=virtual,proto3" json:"virtual,omitempty"`
	Gateway              string    `protobuf:"bytes,6,opt,name=gateway,proto3" json:"gateway,omitempty"`
	Dns                  string    `protobuf:"bytes,7,opt,name=dns,proto3" json:"dns,omitempty"`
	Domain               string    `protobuf:"bytes,8,opt,name=domain,proto3" json:"domain,omitempty"`
	Routes               []*Routes `protobuf:"bytes,9,rep,name=routes,proto3" json:"routes,omitempty"`
	Ifname               string    `protobuf:"bytes,10,opt,name=ifname,proto3" json:"ifname,omitempty"`
	Masklen              int32     `protobuf:"varint,11,opt,name=masklen,proto3" json:"masklen,omitempty"`
	Driver               string    `protobuf:"bytes,12,opt,name=driver,proto3" json:"driver,omitempty"`
	Bridge               string    `protobuf:"bytes,13,opt,name=bridge,proto3" json:"bridge,omitempty"`
	WireId               string    `protobuf:"bytes,14,opt,name=wire_id,json=wireId,proto3" json:"wire_id,omitempty"`
	Vlan                 int32     `protobuf:"varint,15,opt,name=vlan,proto3" json:"vlan,omitempty"`
	Interface            string    `protobuf:"bytes,16,opt,name=interface,proto3" json:"interface,omitempty"`
	Bw                   int32     `protobuf:"varint,17,opt,name=bw,proto3" json:"bw,omitempty"`
	Index                int32     `protobuf:"varint,18,opt,name=index,proto3" json:"index,omitempty"`
	VirtualIps           []string  `protobuf:"bytes,19,rep,name=virtual_ips,json=virtualIps,proto3" json:"virtual_ips,omitempty"`
	ExternelId           string    `protobuf:"bytes,20,opt,name=externel_id,json=externelId,proto3" json:"externel_id,omitempty"`
	TeamWith             string    `protobuf:"bytes,21,opt,name=team_with,json=teamWith,proto3" json:"team_with,omitempty"`
	Manual               bool      `protobuf:"varint,22,opt,name=manual,proto3" json:"manual,omitempty"`
	NicType              string    `protobuf:"bytes,23,opt,name=nic_type,json=nicType,proto3" json:"nic_type,omitempty"`
	LinkUp               bool      `protobuf:"varint,24,opt,name=link_up,json=linkUp,proto3" json:"link_up,omitempty"`
	Mtu                  int64     `protobuf:"varint,25,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Name                 string    `protobuf:"bytes,26,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Nic) Reset()         { *m = Nic{} }
func (m *Nic) String() string { return proto.CompactTextString(m) }
func (*Nic) ProtoMessage()    {}
func (*Nic) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{2}
}

func (m *Nic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nic.Unmarshal(m, b)
}
func (m *Nic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nic.Marshal(b, m, deterministic)
}
func (m *Nic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nic.Merge(m, src)
}
func (m *Nic) XXX_Size() int {
	return xxx_messageInfo_Nic.Size(m)
}
func (m *Nic) XXX_DiscardUnknown() {
	xxx_messageInfo_Nic.DiscardUnknown(m)
}

var xxx_messageInfo_Nic proto.InternalMessageInfo

func (m *Nic) GetMac() string {
	if m != nil {
		return m.Mac
	}
	return ""
}

func (m *Nic) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Nic) GetNet() string {
	if m != nil {
		return m.Net
	}
	return ""
}

func (m *Nic) GetNetId() string {
	if m != nil {
		return m.NetId
	}
	return ""
}

func (m *Nic) GetVirtual() bool {
	if m != nil {
		return m.Virtual
	}
	return false
}

func (m *Nic) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Nic) GetDns() string {
	if m != nil {
		return m.Dns
	}
	return ""
}

func (m *Nic) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Nic) GetRoutes() []*Routes {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *Nic) GetIfname() string {
	if m != nil {
		return m.Ifname
	}
	return ""
}

func (m *Nic) GetMasklen() int32 {
	if m != nil {
		return m.Masklen
	}
	return 0
}

func (m *Nic) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *Nic) GetBridge() string {
	if m != nil {
		return m.Bridge
	}
	return ""
}

func (m *Nic) GetWireId() string {
	if m != nil {
		return m.WireId
	}
	return ""
}

func (m *Nic) GetVlan() int32 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *Nic) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *Nic) GetBw() int32 {
	if m != nil {
		return m.Bw
	}
	return 0
}

func (m *Nic) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Nic) GetVirtualIps() []string {
	if m != nil {
		return m.VirtualIps
	}
	return nil
}

func (m *Nic) GetExternelId() string {
	if m != nil {
		return m.ExternelId
	}
	return ""
}

func (m *Nic) GetTeamWith() string {
	if m != nil {
		return m.TeamWith
	}
	return ""
}

func (m *Nic) GetManual() bool {
	if m != nil {
		return m.Manual
	}
	return false
}

func (m *Nic) GetNicType() string {
	if m != nil {
		return m.NicType
	}
	return ""
}

func (m *Nic) GetLinkUp() bool {
	if m != nil {
		return m.LinkUp
	}
	return false
}

func (m *Nic) GetMtu() int64 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *Nic) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Routes struct {
	Route                []string `protobuf:"bytes,1,rep,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Routes) Reset()         { *m = Routes{} }
func (m *Routes) String() string { return proto.CompactTextString(m) }
func (*Routes) ProtoMessage()    {}
func (*Routes) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{3}
}

func (m *Routes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Routes.Unmarshal(m, b)
}
func (m *Routes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Routes.Marshal(b, m, deterministic)
}
func (m *Routes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Routes.Merge(m, src)
}
func (m *Routes) XXX_Size() int {
	return xxx_messageInfo_Routes.Size(m)
}
func (m *Routes) XXX_DiscardUnknown() {
	xxx_messageInfo_Routes.DiscardUnknown(m)
}

var xxx_messageInfo_Routes proto.InternalMessageInfo

func (m *Routes) GetRoute() []string {
	if m != nil {
		return m.Route
	}
	return nil
}

type DeployInfo struct {
	PublicKey               *SSHKeys         `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Deploys                 []*DeployContent `protobuf:"bytes,2,rep,name=deploys,proto3" json:"deploys,omitempty"`
	Password                string           `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	IsInit                  bool             `protobuf:"varint,4,opt,name=is_init,json=isInit,proto3" json:"is_init,omitempty"`
	EnableTty               bool             `protobuf:"varint,5,opt,name=enable_tty,json=enableTty,proto3" json:"enable_tty,omitempty"`
	DefaultRootUser         bool             `protobuf:"varint,6,opt,name=default_root_user,json=defaultRootUser,proto3" json:"default_root_user,omitempty"`
	WindowsDefaultAdminUser bool             `protobuf:"varint,7,opt,name=windows_default_admin_user,json=windowsDefaultAdminUser,proto3" json:"windows_default_admin_user,omitempty"`
	EnableCloudInit         bool             `protobuf:"varint,8,opt,name=enable_cloud_init,json=enableCloudInit,proto3" json:"enable_cloud_init,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}         `json:"-"`
	XXX_unrecognized        []byte           `json:"-"`
	XXX_sizecache           int32            `json:"-"`
}

func (m *DeployInfo) Reset()         { *m = DeployInfo{} }
func (m *DeployInfo) String() string { return proto.CompactTextString(m) }
func (*DeployInfo) ProtoMessage()    {}
func (*DeployInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{4}
}

func (m *DeployInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployInfo.Unmarshal(m, b)
}
func (m *DeployInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployInfo.Marshal(b, m, deterministic)
}
func (m *DeployInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployInfo.Merge(m, src)
}
func (m *DeployInfo) XXX_Size() int {
	return xxx_messageInfo_DeployInfo.Size(m)
}
func (m *DeployInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeployInfo proto.InternalMessageInfo

func (m *DeployInfo) GetPublicKey() *SSHKeys {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *DeployInfo) GetDeploys() []*DeployContent {
	if m != nil {
		return m.Deploys
	}
	return nil
}

func (m *DeployInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *DeployInfo) GetIsInit() bool {
	if m != nil {
		return m.IsInit
	}
	return false
}

func (m *DeployInfo) GetEnableTty() bool {
	if m != nil {
		return m.EnableTty
	}
	return false
}

func (m *DeployInfo) GetDefaultRootUser() bool {
	if m != nil {
		return m.DefaultRootUser
	}
	return false
}

func (m *DeployInfo) GetWindowsDefaultAdminUser() bool {
	if m != nil {
		return m.WindowsDefaultAdminUser
	}
	return false
}

func (m *DeployInfo) GetEnableCloudInit() bool {
	if m != nil {
		return m.EnableCloudInit
	}
	return false
}

type SSHKeys struct {
	PublicKey            string   `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	DeletePublicKey      string   `protobuf:"bytes,2,opt,name=delete_public_key,json=deletePublicKey,proto3" json:"delete_public_key,omitempty"`
	AdminPublicKey       string   `protobuf:"bytes,3,opt,name=admin_public_key,json=adminPublicKey,proto3" json:"admin_public_key,omitempty"`
	ProjectPublicKey     string   `protobuf:"bytes,4,opt,name=project_public_key,json=projectPublicKey,proto3" json:"project_public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSHKeys) Reset()         { *m = SSHKeys{} }
func (m *SSHKeys) String() string { return proto.CompactTextString(m) }
func (*SSHKeys) ProtoMessage()    {}
func (*SSHKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{5}
}

func (m *SSHKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSHKeys.Unmarshal(m, b)
}
func (m *SSHKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSHKeys.Marshal(b, m, deterministic)
}
func (m *SSHKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSHKeys.Merge(m, src)
}
func (m *SSHKeys) XXX_Size() int {
	return xxx_messageInfo_SSHKeys.Size(m)
}
func (m *SSHKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_SSHKeys.DiscardUnknown(m)
}

var xxx_messageInfo_SSHKeys proto.InternalMessageInfo

func (m *SSHKeys) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *SSHKeys) GetDeletePublicKey() string {
	if m != nil {
		return m.DeletePublicKey
	}
	return ""
}

func (m *SSHKeys) GetAdminPublicKey() string {
	if m != nil {
		return m.AdminPublicKey
	}
	return ""
}

func (m *SSHKeys) GetProjectPublicKey() string {
	if m != nil {
		return m.ProjectPublicKey
	}
	return ""
}

type DeployContent struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeployContent) Reset()         { *m = DeployContent{} }
func (m *DeployContent) String() string { return proto.CompactTextString(m) }
func (*DeployContent) ProtoMessage()    {}
func (*DeployContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{6}
}

func (m *DeployContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployContent.Unmarshal(m, b)
}
func (m *DeployContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployContent.Marshal(b, m, deterministic)
}
func (m *DeployContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployContent.Merge(m, src)
}
func (m *DeployContent) XXX_Size() int {
	return xxx_messageInfo_DeployContent.Size(m)
}
func (m *DeployContent) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployContent.DiscardUnknown(m)
}

var xxx_messageInfo_DeployContent proto.InternalMessageInfo

func (m *DeployContent) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *DeployContent) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *DeployContent) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

type Error struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{7}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *Error) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type DeployGuestFsResponse struct {
	Distro               string   `protobuf:"bytes,1,opt,name=distro,proto3" json:"distro,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Arch                 string   `protobuf:"bytes,3,opt,name=arch,proto3" json:"arch,omitempty"`
	Language             string   `protobuf:"bytes,4,opt,name=language,proto3" json:"language,omitempty"`
	Os                   string   `protobuf:"bytes,5,opt,name=os,proto3" json:"os,omitempty"`
	Account              string   `protobuf:"bytes,6,opt,name=account,proto3" json:"account,omitempty"`
	Key                  string   `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
	Err                  *Error   `protobuf:"bytes,8,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeployGuestFsResponse) Reset()         { *m = DeployGuestFsResponse{} }
func (m *DeployGuestFsResponse) String() string { return proto.CompactTextString(m) }
func (*DeployGuestFsResponse) ProtoMessage()    {}
func (*DeployGuestFsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{8}
}

func (m *DeployGuestFsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployGuestFsResponse.Unmarshal(m, b)
}
func (m *DeployGuestFsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployGuestFsResponse.Marshal(b, m, deterministic)
}
func (m *DeployGuestFsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployGuestFsResponse.Merge(m, src)
}
func (m *DeployGuestFsResponse) XXX_Size() int {
	return xxx_messageInfo_DeployGuestFsResponse.Size(m)
}
func (m *DeployGuestFsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployGuestFsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeployGuestFsResponse proto.InternalMessageInfo

func (m *DeployGuestFsResponse) GetDistro() string {
	if m != nil {
		return m.Distro
	}
	return ""
}

func (m *DeployGuestFsResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DeployGuestFsResponse) GetArch() string {
	if m != nil {
		return m.Arch
	}
	return ""
}

func (m *DeployGuestFsResponse) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *DeployGuestFsResponse) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *DeployGuestFsResponse) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *DeployGuestFsResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DeployGuestFsResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type DeployParams struct {
	DiskPath             string      `protobuf:"bytes,1,opt,name=disk_path,json=diskPath,proto3" json:"disk_path,omitempty"`
	GuestDesc            *GuestDesc  `protobuf:"bytes,2,opt,name=guest_desc,json=guestDesc,proto3" json:"guest_desc,omitempty"`
	DeployInfo           *DeployInfo `protobuf:"bytes,3,opt,name=deploy_info,json=deployInfo,proto3" json:"deploy_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeployParams) Reset()         { *m = DeployParams{} }
func (m *DeployParams) String() string { return proto.CompactTextString(m) }
func (*DeployParams) ProtoMessage()    {}
func (*DeployParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{9}
}

func (m *DeployParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeployParams.Unmarshal(m, b)
}
func (m *DeployParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeployParams.Marshal(b, m, deterministic)
}
func (m *DeployParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployParams.Merge(m, src)
}
func (m *DeployParams) XXX_Size() int {
	return xxx_messageInfo_DeployParams.Size(m)
}
func (m *DeployParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployParams.DiscardUnknown(m)
}

var xxx_messageInfo_DeployParams proto.InternalMessageInfo

func (m *DeployParams) GetDiskPath() string {
	if m != nil {
		return m.DiskPath
	}
	return ""
}

func (m *DeployParams) GetGuestDesc() *GuestDesc {
	if m != nil {
		return m.GuestDesc
	}
	return nil
}

func (m *DeployParams) GetDeployInfo() *DeployInfo {
	if m != nil {
		return m.DeployInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*GuestDesc)(nil), "apis.GuestDesc")
	proto.RegisterType((*Disk)(nil), "apis.Disk")
	proto.RegisterType((*Nic)(nil), "apis.Nic")
	proto.RegisterType((*Routes)(nil), "apis.Routes")
	proto.RegisterType((*DeployInfo)(nil), "apis.DeployInfo")
	proto.RegisterType((*SSHKeys)(nil), "apis.SSHKeys")
	proto.RegisterType((*DeployContent)(nil), "apis.DeployContent")
	proto.RegisterType((*Error)(nil), "apis.Error")
	proto.RegisterType((*DeployGuestFsResponse)(nil), "apis.DeployGuestFsResponse")
	proto.RegisterType((*DeployParams)(nil), "apis.DeployParams")
}

func init() { proto.RegisterFile("deploy.proto", fileDescriptor_05f09e103004e384) }

var fileDescriptor_05f09e103004e384 = []byte{
	// 1226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0xdb, 0x8e, 0xdb, 0x36,
	0x13, 0x86, 0xcf, 0xd6, 0x78, 0x4f, 0x61, 0x4e, 0xcc, 0xe6, 0x4f, 0x62, 0x18, 0x7f, 0x81, 0x45,
	0x90, 0x2e, 0xd0, 0xed, 0x65, 0x6f, 0x12, 0x64, 0x7b, 0x30, 0x82, 0xa6, 0x81, 0x36, 0x41, 0x2f,
	0x05, 0x5a, 0xa2, 0xbd, 0xec, 0x4a, 0xa4, 0x40, 0x52, 0xeb, 0xb8, 0x0f, 0xd1, 0x27, 0xe9, 0x13,
	0xf4, 0x31, 0x8a, 0xde, 0xf4, 0x4d, 0x7a, 0x59, 0xcc, 0x90, 0xb6, 0x95, 0x00, 0xbd, 0xf2, 0xcc,
	0x37, 0x43, 0x6a, 0x38, 0xdf, 0xc7, 0xa1, 0xe1, 0xa0, 0x90, 0x75, 0x69, 0x36, 0xe7, 0xb5, 0x35,
	0xde, 0xb0, 0xbe, 0xa8, 0x95, 0x9b, 0xfd, 0xd1, 0x81, 0xe4, 0xfb, 0x46, 0x3a, 0x7f, 0x29, 0x5d,
	0xce, 0x18, 0xf4, 0xb5, 0xa8, 0x24, 0xef, 0x4c, 0x3b, 0x67, 0x49, 0x4a, 0x36, 0x62, 0x4d, 0xa3,
	0x0a, 0xde, 0x0d, 0x18, 0xda, 0xec, 0x01, 0x0c, 0x0b, 0x53, 0x09, 0xa5, 0x79, 0x8f, 0xd0, 0xe8,
	0xb1, 0x27, 0xd0, 0xd7, 0x2a, 0x77, 0xbc, 0x3f, 0xed, 0x9d, 0x4d, 0x2e, 0x92, 0x73, 0xfc, 0xc4,
	0xf9, 0x5b, 0x95, 0xa7, 0x04, 0xb3, 0x17, 0x70, 0x80, 0xbf, 0x99, 0xf3, 0x42, 0x17, 0x8b, 0x0d,
	0x1f, 0x7c, 0x9e, 0x36, 0xc1, 0xf0, 0x55, 0x88, 0xb2, 0x29, 0x0c, 0x0a, 0xe5, 0x6e, 0x1c, 0x1f,
	0x52, 0x1a, 0x84, 0xb4, 0x4b, 0xe5, 0x6e, 0xd2, 0x10, 0x98, 0xfd, 0xdd, 0x83, 0x3e, 0xfa, 0xec,
	0x21, 0x8c, 0x10, 0xc9, 0x54, 0x11, 0x4b, 0x1f, 0xa2, 0x3b, 0x0f, 0x85, 0x5a, 0x75, 0x2b, 0x6d,
	0x2c, 0x3f, 0x7a, 0xec, 0x09, 0x40, 0x2e, 0xf2, 0x6b, 0x99, 0x55, 0xa6, 0x90, 0xf1, 0x10, 0x09,
	0x21, 0x3f, 0x9a, 0x42, 0xb2, 0x47, 0x30, 0x16, 0xca, 0x84, 0x60, 0x9f, 0x82, 0x23, 0xa1, 0x0c,
	0x85, 0x18, 0xf4, 0x9d, 0xfa, 0x55, 0xf2, 0xc1, 0xb4, 0x73, 0xd6, 0x4b, 0xc9, 0x66, 0xcf, 0x60,
	0xe2, 0x65, 0x55, 0x97, 0xc2, 0x4b, 0x2c, 0x61, 0x48, 0x2b, 0x60, 0x0b, 0xcd, 0x0b, 0xfc, 0x9c,
	0xaa, 0xc4, 0x4a, 0x66, 0xb5, 0xf0, 0xd7, 0x7c, 0x14, 0x3e, 0x47, 0xc8, 0x3b, 0xe1, 0xaf, 0x31,
	0xec, 0xbc, 0xb1, 0x98, 0xa0, 0x0a, 0x3e, 0x0e, 0xe1, 0x88, 0xcc, 0x0b, 0xf6, 0x3f, 0x48, 0x2a,
	0xb5, 0xb2, 0xc2, 0x2b, 0xbd, 0xe2, 0xc9, 0xb4, 0x73, 0x36, 0x4e, 0xf7, 0x00, 0x7b, 0x0e, 0x77,
	0xbc, 0xb0, 0x2b, 0xe9, 0xb3, 0xd6, 0x1e, 0x40, 0x7b, 0x1c, 0x87, 0xc0, 0xd5, 0x6e, 0x27, 0x06,
	0x7d, 0xaa, 0x60, 0x12, 0xb8, 0x44, 0x1b, 0x5b, 0xb4, 0x34, 0xb6, 0x12, 0x9e, 0x1f, 0x84, 0x16,
	0x05, 0x8f, 0xdd, 0x83, 0x81, 0xd2, 0x85, 0xfc, 0xc8, 0x0f, 0xa7, 0x9d, 0xb3, 0x41, 0x1a, 0x1c,
	0xf6, 0x05, 0x1c, 0x55, 0xd2, 0xae, 0x64, 0xe6, 0xb4, 0xa8, 0xdd, 0xb5, 0xf1, 0xfc, 0x88, 0x0a,
	0x3a, 0x24, 0xf4, 0x2a, 0x82, 0xec, 0x08, 0xba, 0x4b, 0xc7, 0x8f, 0x69, 0xc3, 0xee, 0xd2, 0xb1,
	0xa7, 0x00, 0x95, 0x69, 0xb4, 0xaf, 0x8d, 0xd2, 0x9e, 0x9f, 0x84, 0x06, 0xed, 0x11, 0x76, 0x02,
	0xbd, 0x42, 0xde, 0xf2, 0x3b, 0x14, 0x40, 0x73, 0xf6, 0x4f, 0x1f, 0x7a, 0x6f, 0x55, 0x8e, 0x91,
	0x4a, 0xe4, 0x91, 0x56, 0x34, 0x71, 0x6f, 0x55, 0x47, 0x3e, 0xbb, 0xaa, 0xc6, 0x0c, 0x2d, 0x7d,
	0x24, 0x11, 0x4d, 0x76, 0x1f, 0x86, 0x5a, 0x7a, 0xec, 0x43, 0x20, 0x6f, 0xa0, 0xa5, 0x9f, 0x17,
	0x8c, 0xc3, 0xe8, 0x56, 0x59, 0xdf, 0x88, 0x92, 0xd8, 0x1b, 0xa7, 0x5b, 0x17, 0x23, 0x2b, 0xe1,
	0xe5, 0x5a, 0x6c, 0x22, 0x79, 0x5b, 0x97, 0x0a, 0xd3, 0x2e, 0x52, 0x86, 0x66, 0x4b, 0xfb, 0xe3,
	0x4f, 0xb4, 0xff, 0x7f, 0x18, 0x5a, 0xd3, 0x78, 0xe9, 0x78, 0x42, 0x7a, 0x3d, 0x08, 0x7a, 0x4d,
	0x09, 0x4b, 0x63, 0x0c, 0x57, 0xab, 0x25, 0xdd, 0xb1, 0x40, 0x51, 0xf4, 0xb0, 0x82, 0x4a, 0xb8,
	0x9b, 0x52, 0x6a, 0x22, 0x67, 0x90, 0x6e, 0xdd, 0x96, 0x84, 0x0f, 0x3e, 0x91, 0xf0, 0x03, 0x18,
	0x2e, 0xac, 0x2a, 0x56, 0x92, 0x08, 0x4a, 0xd2, 0xe8, 0xe1, 0x5d, 0x58, 0x2b, 0x4b, 0x2a, 0x38,
	0x0a, 0x01, 0x74, 0x03, 0xf9, 0xb7, 0xa5, 0xd0, 0xc4, 0xca, 0x20, 0x25, 0x1b, 0xa5, 0xa5, 0xb4,
	0x97, 0x76, 0x29, 0x72, 0x19, 0x69, 0xd9, 0x03, 0xd8, 0xe9, 0xc5, 0x9a, 0x48, 0x19, 0xa4, 0xdd,
	0xc5, 0x7a, 0x2f, 0x09, 0xd6, 0x96, 0xc4, 0x33, 0x98, 0xc4, 0x3e, 0x66, 0xaa, 0x76, 0xfc, 0xee,
	0xb4, 0x87, 0xe4, 0x46, 0x68, 0x5e, 0x3b, 0x4c, 0x90, 0x1f, 0xbd, 0xb4, 0x5a, 0x96, 0x58, 0xd5,
	0xbd, 0xc0, 0xfe, 0x16, 0x9a, 0x17, 0xec, 0x31, 0x24, 0x5e, 0x8a, 0x2a, 0x5b, 0x2b, 0x7f, 0xcd,
	0xef, 0x53, 0x78, 0x8c, 0xc0, 0xcf, 0x2a, 0xe8, 0xb3, 0x12, 0x1a, 0x49, 0x7b, 0x40, 0xa4, 0x45,
	0x0f, 0xef, 0xa8, 0x56, 0x79, 0xe6, 0x37, 0xb5, 0xe4, 0x0f, 0x03, 0x69, 0x5a, 0xe5, 0xef, 0x37,
	0x35, 0xb5, 0xa0, 0x54, 0xfa, 0x26, 0x6b, 0x6a, 0xce, 0xc3, 0x1a, 0x74, 0x3f, 0x90, 0x54, 0x2a,
	0xdf, 0xf0, 0x47, 0x74, 0x77, 0xd1, 0xdc, 0x4d, 0xbc, 0xd3, 0xfd, 0xc4, 0x9b, 0x3d, 0x85, 0x61,
	0x60, 0x0d, 0x0f, 0x4c, 0xbc, 0xf1, 0x0e, 0x1d, 0x2a, 0x38, 0xb3, 0x3f, 0xbb, 0x00, 0x97, 0x34,
	0x4a, 0xe7, 0x7a, 0x69, 0xd8, 0x0b, 0x80, 0xba, 0x59, 0x94, 0x2a, 0xcf, 0x6e, 0xe4, 0x86, 0x84,
	0x3a, 0xb9, 0x38, 0x0c, 0xe4, 0x5f, 0x5d, 0xfd, 0xf0, 0x46, 0x6e, 0x5c, 0x9a, 0x84, 0x84, 0x37,
	0x72, 0xc3, 0xbe, 0x84, 0x51, 0x18, 0xc3, 0x8e, 0x77, 0x49, 0x27, 0x77, 0xe3, 0x5c, 0x23, 0xf0,
	0xb5, 0xd1, 0x5e, 0x6a, 0x9f, 0x6e, 0x73, 0xd8, 0x29, 0x8c, 0x6b, 0xe1, 0xdc, 0xda, 0xd8, 0x22,
	0x2a, 0x7c, 0xe7, 0xe3, 0x31, 0x95, 0xcb, 0x94, 0x56, 0x9e, 0x74, 0x3e, 0x4e, 0x87, 0xca, 0xcd,
	0xb5, 0xf2, 0x38, 0x4f, 0xa4, 0x16, 0x8b, 0x52, 0x66, 0xde, 0x6f, 0xa2, 0xd6, 0x93, 0x80, 0xbc,
	0xf7, 0x1b, 0x9c, 0x18, 0x85, 0x5c, 0x8a, 0xa6, 0xf4, 0x99, 0x35, 0xc6, 0x67, 0x8d, 0x93, 0x96,
	0x74, 0x3f, 0x4e, 0x8f, 0x63, 0x20, 0x35, 0xc6, 0x7f, 0x70, 0xd2, 0xb2, 0x6f, 0xe0, 0x74, 0xad,
	0x74, 0x61, 0xd6, 0x2e, 0xdb, 0xae, 0x11, 0x45, 0xa5, 0x74, 0x58, 0x34, 0xa2, 0x45, 0x0f, 0x63,
	0xc6, 0x65, 0x48, 0x78, 0x85, 0x71, 0x5a, 0xfc, 0x1c, 0xee, 0xc4, 0x3a, 0xf2, 0xd2, 0x34, 0x45,
	0x28, 0x75, 0x1c, 0x3e, 0x14, 0x02, 0xaf, 0x11, 0xc7, 0x9a, 0x67, 0xbf, 0x77, 0x60, 0x14, 0xdb,
	0x85, 0xf5, 0x7f, 0xd6, 0xd1, 0xa4, 0xdd, 0x42, 0xaa, 0xbf, 0x94, 0x5e, 0x66, 0xad, 0xac, 0x30,
	0x0f, 0x8e, 0x43, 0xe0, 0xdd, 0x2e, 0xf7, 0x0c, 0x4e, 0x42, 0xbd, 0xad, 0xd4, 0xd0, 0xc7, 0x23,
	0xc2, 0xf7, 0x99, 0x2f, 0x80, 0xd5, 0xd6, 0xfc, 0x22, 0x73, 0xdf, 0xce, 0x0d, 0x03, 0xe4, 0x24,
	0x46, 0x76, 0xd9, 0xb3, 0x0f, 0x70, 0xf8, 0x09, 0x63, 0xbb, 0xd1, 0xda, 0x69, 0x8d, 0x56, 0x0e,
	0xa3, 0x3c, 0x84, 0x63, 0x79, 0x5b, 0x17, 0x45, 0x2d, 0x72, 0xaf, 0xcc, 0xee, 0x01, 0x0d, 0xde,
	0xec, 0x25, 0x0c, 0xbe, 0xb5, 0xd6, 0xd0, 0x03, 0x25, 0xd1, 0xc8, 0x72, 0x7c, 0x83, 0x3a, 0x74,
	0xdf, 0x12, 0x42, 0x5e, 0xe3, 0x2b, 0xf4, 0x9f, 0x3b, 0xcf, 0xfe, 0xea, 0xc0, 0xfd, 0x50, 0x19,
	0x3d, 0xeb, 0xdf, 0xb9, 0x54, 0xba, 0xda, 0x68, 0x27, 0x69, 0x90, 0x28, 0xe7, 0xad, 0x69, 0xbd,
	0x91, 0xde, 0x1a, 0x1a, 0x8b, 0xd2, 0x3a, 0x2c, 0x26, 0xee, 0x15, 0x5d, 0x3c, 0x93, 0xb0, 0xf9,
	0x75, 0xac, 0x91, 0x6c, 0x14, 0x64, 0x29, 0xf4, 0xaa, 0x11, 0xab, 0xed, 0xd3, 0xb8, 0xf3, 0x71,
	0x5e, 0x18, 0x47, 0x7a, 0x4b, 0xd2, 0xae, 0x71, 0xb8, 0xb3, 0xc8, 0x73, 0x9c, 0xf2, 0xdb, 0xb1,
	0x1a, 0x5d, 0xbc, 0x88, 0xd8, 0xdd, 0x38, 0x56, 0x6f, 0xe4, 0x86, 0x3d, 0x81, 0x9e, 0xb4, 0x96,
	0xd4, 0x31, 0xb9, 0x98, 0x84, 0x3b, 0x41, 0xad, 0x48, 0x11, 0x9f, 0xfd, 0xd6, 0x81, 0x83, 0x70,
	0xac, 0x77, 0xc2, 0x8a, 0xca, 0xe1, 0xcc, 0xa0, 0x27, 0xbf, 0xd5, 0xf4, 0x31, 0x02, 0xf4, 0xa0,
	0x9e, 0x03, 0xac, 0xf0, 0xf4, 0x59, 0x21, 0x5d, 0x4e, 0xa7, 0x9a, 0x5c, 0x1c, 0x87, 0x3d, 0x77,
	0x7f, 0x76, 0xd2, 0x64, 0xb5, 0xfb, 0xdf, 0xf3, 0x15, 0x4c, 0xc2, 0x85, 0xcb, 0x94, 0x5e, 0x1a,
	0x3a, 0xef, 0xe4, 0xe2, 0xa4, 0x7d, 0x31, 0xf1, 0xa6, 0xa7, 0x50, 0xec, 0xec, 0x8b, 0x9f, 0x60,
	0x12, 0x22, 0xaf, 0x56, 0x48, 0xe8, 0xcb, 0xad, 0x1e, 0x62, 0xd7, 0x19, 0x6b, 0xaf, 0x0e, 0x35,
	0x9f, 0x3e, 0x6e, 0x63, 0x9f, 0xd1, 0xb3, 0x18, 0xd2, 0xdf, 0xb2, 0xaf, 0xff, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x07, 0x0b, 0x54, 0xc0, 0xa6, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeployAgentClient is the client API for DeployAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeployAgentClient interface {
	DeployGuestFs(ctx context.Context, in *DeployParams, opts ...grpc.CallOption) (*DeployGuestFsResponse, error)
}

type deployAgentClient struct {
	cc *grpc.ClientConn
}

func NewDeployAgentClient(cc *grpc.ClientConn) DeployAgentClient {
	return &deployAgentClient{cc}
}

func (c *deployAgentClient) DeployGuestFs(ctx context.Context, in *DeployParams, opts ...grpc.CallOption) (*DeployGuestFsResponse, error) {
	out := new(DeployGuestFsResponse)
	err := c.cc.Invoke(ctx, "/apis.DeployAgent/DeployGuestFs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeployAgentServer is the server API for DeployAgent service.
type DeployAgentServer interface {
	DeployGuestFs(context.Context, *DeployParams) (*DeployGuestFsResponse, error)
}

// UnimplementedDeployAgentServer can be embedded to have forward compatible implementations.
type UnimplementedDeployAgentServer struct {
}

func (*UnimplementedDeployAgentServer) DeployGuestFs(ctx context.Context, req *DeployParams) (*DeployGuestFsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployGuestFs not implemented")
}

func RegisterDeployAgentServer(s *grpc.Server, srv DeployAgentServer) {
	s.RegisterService(&_DeployAgent_serviceDesc, srv)
}

func _DeployAgent_DeployGuestFs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeployAgentServer).DeployGuestFs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.DeployAgent/DeployGuestFs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeployAgentServer).DeployGuestFs(ctx, req.(*DeployParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeployAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apis.DeployAgent",
	HandlerType: (*DeployAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeployGuestFs",
			Handler:    _DeployAgent_DeployGuestFs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deploy.proto",
}
