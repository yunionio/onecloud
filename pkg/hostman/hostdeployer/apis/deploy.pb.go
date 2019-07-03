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

type Routes struct {
	Route                []*Route `protobuf:"bytes,1,rep,name=route,proto3" json:"route,omitempty"`
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

func (m *Routes) GetRoute() []*Route {
	if m != nil {
		return m.Route
	}
	return nil
}

type Route struct {
	Route                []string `protobuf:"bytes,1,rep,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{4}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetRoute() []string {
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
	return fileDescriptor_05f09e103004e384, []int{5}
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
	return fileDescriptor_05f09e103004e384, []int{6}
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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeployContent) Reset()         { *m = DeployContent{} }
func (m *DeployContent) String() string { return proto.CompactTextString(m) }
func (*DeployContent) ProtoMessage()    {}
func (*DeployContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_05f09e103004e384, []int{7}
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
	return fileDescriptor_05f09e103004e384, []int{8}
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
	return fileDescriptor_05f09e103004e384, []int{9}
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
	return fileDescriptor_05f09e103004e384, []int{10}
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
	proto.RegisterType((*Route)(nil), "apis.Route")
	proto.RegisterType((*DeployInfo)(nil), "apis.DeployInfo")
	proto.RegisterType((*SSHKeys)(nil), "apis.SSHKeys")
	proto.RegisterType((*DeployContent)(nil), "apis.DeployContent")
	proto.RegisterType((*Error)(nil), "apis.Error")
	proto.RegisterType((*DeployGuestFsResponse)(nil), "apis.DeployGuestFsResponse")
	proto.RegisterType((*DeployParams)(nil), "apis.DeployParams")
}

func init() { proto.RegisterFile("deploy.proto", fileDescriptor_05f09e103004e384) }

var fileDescriptor_05f09e103004e384 = []byte{
	// 1201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0xdb, 0x6e, 0x1b, 0x37,
	0x13, 0x86, 0xce, 0xda, 0x91, 0x4f, 0x61, 0x4e, 0xfb, 0x27, 0xbf, 0x1b, 0x57, 0x68, 0x01, 0x23,
	0x4d, 0x0d, 0xd4, 0xbd, 0x2c, 0x0a, 0x24, 0x88, 0x7b, 0x10, 0x82, 0xa6, 0xc1, 0x3a, 0x45, 0x2f,
	0x17, 0xd4, 0xee, 0x48, 0x66, 0xad, 0x25, 0x17, 0x24, 0x65, 0x45, 0x7d, 0x88, 0x3e, 0x49, 0x5f,
	0xa0, 0x7d, 0x8c, 0xa2, 0x37, 0x7d, 0x9b, 0x62, 0x86, 0x5c, 0x79, 0x13, 0xa0, 0x57, 0x9a, 0xf9,
	0xbe, 0x21, 0x77, 0x38, 0xdf, 0x70, 0x28, 0xd8, 0x2b, 0xb1, 0x5e, 0x99, 0xed, 0x59, 0x6d, 0x8d,
	0x37, 0xa2, 0x2f, 0x6b, 0xe5, 0xa6, 0x7f, 0x76, 0x20, 0xf9, 0x6e, 0x8d, 0xce, 0x5f, 0xa0, 0x2b,
	0x84, 0x80, 0xbe, 0x96, 0x15, 0xa6, 0x9d, 0x93, 0xce, 0x69, 0x92, 0xb1, 0x4d, 0xd8, 0x7a, 0xad,
	0xca, 0xb4, 0x1b, 0x30, 0xb2, 0xc5, 0x03, 0x18, 0x96, 0xa6, 0x92, 0x4a, 0xa7, 0x3d, 0x46, 0xa3,
	0x27, 0x8e, 0xa1, 0xaf, 0x55, 0xe1, 0xd2, 0xfe, 0x49, 0xef, 0x74, 0x72, 0x9e, 0x9c, 0xd1, 0x27,
	0xce, 0x5e, 0xab, 0x22, 0x63, 0x58, 0x3c, 0x83, 0x3d, 0xfa, 0xcd, 0x9d, 0x97, 0xba, 0x9c, 0x6f,
	0xd3, 0xc1, 0x87, 0x61, 0x13, 0xa2, 0x2f, 0x03, 0x2b, 0x4e, 0x60, 0x50, 0x2a, 0x77, 0xed, 0xd2,
	0x21, 0x87, 0x41, 0x08, 0xbb, 0x50, 0xee, 0x3a, 0x0b, 0xc4, 0xf4, 0x9f, 0x1e, 0xf4, 0xc9, 0x17,
	0x0f, 0x61, 0x44, 0x48, 0xae, 0xca, 0x98, 0xfa, 0x90, 0xdc, 0x59, 0x48, 0xd4, 0xaa, 0x1b, 0xb4,
	0x31, 0xfd, 0xe8, 0x89, 0x63, 0x80, 0x42, 0x16, 0x57, 0x98, 0x57, 0xa6, 0xc4, 0x78, 0x88, 0x84,
	0x91, 0x1f, 0x4c, 0x89, 0xe2, 0x7f, 0x30, 0x96, 0xca, 0x04, 0xb2, 0xcf, 0xe4, 0x48, 0x2a, 0xc3,
	0x94, 0x80, 0xbe, 0x53, 0xbf, 0x62, 0x3a, 0x38, 0xe9, 0x9c, 0xf6, 0x32, 0xb6, 0xc5, 0x13, 0x98,
	0x78, 0xac, 0xea, 0x95, 0xf4, 0x48, 0x29, 0x0c, 0x79, 0x05, 0x34, 0xd0, 0xac, 0xa4, 0xcf, 0xa9,
	0x4a, 0x2e, 0x31, 0xaf, 0xa5, 0xbf, 0x4a, 0x47, 0xe1, 0x73, 0x8c, 0xbc, 0x91, 0xfe, 0x8a, 0x68,
	0xe7, 0x8d, 0xa5, 0x00, 0x55, 0xa6, 0xe3, 0x40, 0x47, 0x64, 0x56, 0x8a, 0xff, 0x43, 0x52, 0xa9,
	0xa5, 0x95, 0x5e, 0xe9, 0x65, 0x9a, 0x9c, 0x74, 0x4e, 0xc7, 0xd9, 0x2d, 0x20, 0x9e, 0xc2, 0x1d,
	0x2f, 0xed, 0x12, 0x7d, 0xde, 0xda, 0x03, 0x78, 0x8f, 0xc3, 0x40, 0x5c, 0xee, 0x76, 0x12, 0xd0,
	0xe7, 0x0c, 0x26, 0x41, 0x4b, 0xb2, 0xa9, 0x44, 0x0b, 0x63, 0x2b, 0xe9, 0xd3, 0xbd, 0x50, 0xa2,
	0xe0, 0x89, 0x7b, 0x30, 0x50, 0xba, 0xc4, 0x77, 0xe9, 0xfe, 0x49, 0xe7, 0x74, 0x90, 0x05, 0x47,
	0x7c, 0x0a, 0x07, 0x15, 0xda, 0x25, 0xe6, 0x4e, 0xcb, 0xda, 0x5d, 0x19, 0x9f, 0x1e, 0x70, 0x42,
	0xfb, 0x8c, 0x5e, 0x46, 0x50, 0x1c, 0x40, 0x77, 0xe1, 0xd2, 0x43, 0xde, 0xb0, 0xbb, 0x70, 0xe2,
	0x23, 0x80, 0xca, 0xac, 0xb5, 0xaf, 0x8d, 0xd2, 0x3e, 0x3d, 0x0a, 0x05, 0xba, 0x45, 0xc4, 0x11,
	0xf4, 0x4a, 0xbc, 0x49, 0xef, 0x30, 0x41, 0xe6, 0xf4, 0x8f, 0x3e, 0xf4, 0x5e, 0xab, 0x82, 0x98,
	0x4a, 0x16, 0x51, 0x56, 0x32, 0x69, 0x6f, 0x55, 0x47, 0x3d, 0xbb, 0xaa, 0xa6, 0x08, 0x8d, 0x3e,
	0x8a, 0x48, 0xa6, 0xb8, 0x0f, 0x43, 0x8d, 0x9e, 0xea, 0x10, 0xc4, 0x1b, 0x68, 0xf4, 0xb3, 0x52,
	0xa4, 0x30, 0xba, 0x51, 0xd6, 0xaf, 0xe5, 0x8a, 0xd5, 0x1b, 0x67, 0x8d, 0x4b, 0xcc, 0x52, 0x7a,
	0xdc, 0xc8, 0x6d, 0x14, 0xaf, 0x71, 0x39, 0x31, 0xed, 0xa2, 0x64, 0x64, 0xb6, 0x7a, 0x7f, 0xfc,
	0x5e, 0xef, 0x7f, 0x02, 0x43, 0x6b, 0xd6, 0x1e, 0x5d, 0x9a, 0x70, 0xbf, 0xee, 0x85, 0x7e, 0xcd,
	0x18, 0xcb, 0x22, 0x47, 0xab, 0xd5, 0x82, 0xef, 0x58, 0x90, 0x28, 0x7a, 0x94, 0x41, 0x25, 0xdd,
	0xf5, 0x0a, 0x35, 0x8b, 0x33, 0xc8, 0x1a, 0xb7, 0xd5, 0xc2, 0x7b, 0xef, 0xb5, 0xf0, 0x03, 0x18,
	0xce, 0xad, 0x2a, 0x97, 0xc8, 0x02, 0x25, 0x59, 0xf4, 0xe8, 0x2e, 0x6c, 0x94, 0xe5, 0x2e, 0x38,
	0x08, 0x04, 0xb9, 0x41, 0xfc, 0x9b, 0x95, 0xd4, 0xac, 0xca, 0x20, 0x63, 0x9b, 0x5a, 0x4b, 0x69,
	0x8f, 0x76, 0x21, 0x0b, 0x8c, 0xb2, 0xdc, 0x02, 0x54, 0xe9, 0xf9, 0x86, 0x45, 0x19, 0x64, 0xdd,
	0xf9, 0xe6, 0xb6, 0x25, 0x44, 0xbb, 0x25, 0x9e, 0xc0, 0x24, 0xd6, 0x31, 0x57, 0xb5, 0x4b, 0xef,
	0x9e, 0xf4, 0x48, 0xdc, 0x08, 0xcd, 0x6a, 0x47, 0x01, 0xf8, 0xce, 0xa3, 0xd5, 0xb8, 0xa2, 0xac,
	0xee, 0x05, 0xf5, 0x1b, 0x68, 0x56, 0x8a, 0xc7, 0x90, 0x78, 0x94, 0x55, 0xbe, 0x51, 0xfe, 0x2a,
	0xbd, 0xcf, 0xf4, 0x98, 0x80, 0x9f, 0x55, 0xe8, 0xcf, 0x4a, 0x6a, 0x12, 0xed, 0x01, 0x8b, 0x16,
	0x3d, 0xba, 0xa3, 0x5a, 0x15, 0xb9, 0xdf, 0xd6, 0x98, 0x3e, 0x0c, 0xa2, 0x69, 0x55, 0xbc, 0xdd,
	0xd6, 0x38, 0xfd, 0x0c, 0x86, 0xa1, 0xec, 0xe2, 0x63, 0x18, 0x70, 0xe1, 0xd3, 0x0e, 0x6b, 0x32,
	0x69, 0x69, 0x92, 0x05, 0x66, 0x7a, 0x0c, 0x03, 0xf6, 0xe9, 0x74, 0xb7, 0xb1, 0x49, 0x43, 0xff,
	0xd5, 0x05, 0xb8, 0xe0, 0xb9, 0x39, 0xd3, 0x0b, 0x23, 0x9e, 0x01, 0xd4, 0xeb, 0xf9, 0x4a, 0x15,
	0xf9, 0x35, 0x6e, 0xb9, 0x2b, 0x27, 0xe7, 0xfb, 0x61, 0xd7, 0xcb, 0xcb, 0xef, 0x5f, 0xe1, 0xd6,
	0x65, 0x49, 0x08, 0x78, 0x85, 0x5b, 0xf1, 0x39, 0x8c, 0xc2, 0xcc, 0x75, 0x69, 0x97, 0x13, 0xb8,
	0x1b, 0x87, 0x18, 0x83, 0x2f, 0x8d, 0xf6, 0xa8, 0x7d, 0xd6, 0xc4, 0x88, 0x47, 0x30, 0xae, 0xa5,
	0x73, 0x1b, 0x63, 0xcb, 0xd8, 0xce, 0x3b, 0x9f, 0x64, 0x55, 0x2e, 0x57, 0x5a, 0x79, 0x6e, 0xea,
	0x71, 0x36, 0x54, 0x6e, 0xa6, 0x95, 0xa7, 0xe1, 0x81, 0x5a, 0xce, 0x57, 0x98, 0x7b, 0xbf, 0x8d,
	0x8d, 0x9d, 0x04, 0xe4, 0xad, 0xdf, 0xd2, 0x78, 0x28, 0x71, 0x21, 0xd7, 0x2b, 0x9f, 0x5b, 0x63,
	0x7c, 0xbe, 0x76, 0x68, 0xb9, 0xc9, 0xc7, 0xd9, 0x61, 0x24, 0x32, 0x63, 0xfc, 0x4f, 0x0e, 0xad,
	0xf8, 0x0a, 0x1e, 0x6d, 0x94, 0x2e, 0xcd, 0xc6, 0xe5, 0xcd, 0x1a, 0x59, 0x56, 0x4a, 0x87, 0x45,
	0x23, 0x5e, 0xf4, 0x30, 0x46, 0x5c, 0x84, 0x80, 0x17, 0xc4, 0xf3, 0xe2, 0xa7, 0x70, 0x27, 0xe6,
	0x51, 0xac, 0xcc, 0xba, 0x0c, 0xa9, 0x8e, 0xc3, 0x87, 0x02, 0xf1, 0x92, 0x70, 0xca, 0x79, 0xfa,
	0x7b, 0x07, 0x46, 0xb1, 0x5c, 0x94, 0xff, 0x07, 0x15, 0x4d, 0xda, 0x25, 0xe4, 0xfc, 0x57, 0xe8,
	0x31, 0x6f, 0x45, 0x85, 0xcb, 0x7f, 0x18, 0x88, 0x37, 0xbb, 0xd8, 0x53, 0x38, 0x0a, 0xf9, 0xb6,
	0x42, 0x43, 0x1d, 0x0f, 0x18, 0xbf, 0x8d, 0x7c, 0x06, 0xa2, 0xb6, 0xe6, 0x17, 0x2c, 0x7c, 0x3b,
	0x36, 0x4c, 0x8b, 0xa3, 0xc8, 0xec, 0xa2, 0xa7, 0x5f, 0xc3, 0xfe, 0x7b, 0x8a, 0xed, 0xe6, 0x68,
	0xa7, 0x35, 0x47, 0x53, 0x18, 0x15, 0x81, 0x8e, 0xe9, 0x35, 0xee, 0xf4, 0x39, 0x0c, 0xbe, 0xb1,
	0xd6, 0xf0, 0xab, 0x83, 0x64, 0xe4, 0x05, 0x3d, 0x2c, 0x1d, 0xbe, 0x44, 0x09, 0x23, 0x2f, 0xe9,
	0x69, 0xf9, 0xef, 0x1d, 0xfe, 0xee, 0xc0, 0xfd, 0x90, 0x01, 0xbf, 0xd5, 0xdf, 0xba, 0x0c, 0x5d,
	0x6d, 0xb4, 0x43, 0x9e, 0x0e, 0xca, 0x79, 0x6b, 0x5a, 0x0f, 0x9f, 0xb7, 0x86, 0x67, 0x1d, 0x5a,
	0xa7, 0x8c, 0x6e, 0xf6, 0x8a, 0x2e, 0xe5, 0x2e, 0x6d, 0x71, 0x15, 0x0b, 0xc3, 0x36, 0x35, 0xde,
	0x4a, 0xea, 0xe5, 0x5a, 0x2e, 0x9b, 0xf7, 0x6e, 0xe7, 0xd3, 0x10, 0x30, 0x8e, 0xfb, 0x2a, 0xc9,
	0xba, 0xc6, 0xd1, 0xce, 0xb2, 0x28, 0x68, 0x74, 0x37, 0xb3, 0x32, 0xba, 0x34, 0x2b, 0xa9, 0x8a,
	0x71, 0x56, 0x5e, 0xe3, 0x56, 0x1c, 0x43, 0x0f, 0xad, 0xe5, 0x2e, 0xd8, 0x5d, 0x3e, 0x2e, 0x45,
	0x46, 0xf8, 0xf4, 0xb7, 0x0e, 0xec, 0x85, 0x63, 0xbd, 0x91, 0x56, 0x56, 0x8e, 0x06, 0x01, 0xbf,
	0xe3, 0xad, 0xe2, 0x8e, 0x09, 0xe0, 0x57, 0xf2, 0x0c, 0x60, 0x49, 0xa7, 0xcf, 0x4b, 0x74, 0x05,
	0x9f, 0x6a, 0x72, 0x7e, 0x18, 0xf6, 0xdc, 0xfd, 0x83, 0xc9, 0x92, 0xe5, 0xee, 0xcf, 0xcc, 0x17,
	0x30, 0x09, 0x17, 0x2b, 0x57, 0x7a, 0x61, 0xf8, 0xbc, 0x93, 0xf3, 0xa3, 0xf6, 0x05, 0xa4, 0x1b,
	0x9d, 0x41, 0xb9, 0xb3, 0xcf, 0x7f, 0x84, 0x49, 0x60, 0x5e, 0x2c, 0x49, 0xe6, 0xe7, 0x8d, 0xee,
	0xb1, 0xea, 0x42, 0xb4, 0x57, 0x87, 0x9c, 0x1f, 0x3d, 0x6e, 0x63, 0x1f, 0xc8, 0x33, 0x1f, 0xf2,
	0x7f, 0xad, 0x2f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xee, 0x40, 0xb5, 0xab, 0x7b, 0x09, 0x00,
	0x00,
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
