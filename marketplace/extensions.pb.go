// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.3
// source: extensions.proto

package marketplace

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WebsocketMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//	*WebsocketMessage_Install
	//	*WebsocketMessage_Simplemessage
	Content isWebsocketMessage_Content `protobuf_oneof:"content"`
}

func (x *WebsocketMessage) Reset() {
	*x = WebsocketMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebsocketMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebsocketMessage) ProtoMessage() {}

func (x *WebsocketMessage) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebsocketMessage.ProtoReflect.Descriptor instead.
func (*WebsocketMessage) Descriptor() ([]byte, []int) {
	return file_extensions_proto_rawDescGZIP(), []int{0}
}

func (m *WebsocketMessage) GetContent() isWebsocketMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *WebsocketMessage) GetInstall() *Install {
	if x, ok := x.GetContent().(*WebsocketMessage_Install); ok {
		return x.Install
	}
	return nil
}

func (x *WebsocketMessage) GetSimplemessage() *SimpleMessage {
	if x, ok := x.GetContent().(*WebsocketMessage_Simplemessage); ok {
		return x.Simplemessage
	}
	return nil
}

type isWebsocketMessage_Content interface {
	isWebsocketMessage_Content()
}

type WebsocketMessage_Install struct {
	Install *Install `protobuf:"bytes,1,opt,name=install,proto3,oneof"`
}

type WebsocketMessage_Simplemessage struct {
	Simplemessage *SimpleMessage `protobuf:"bytes,2,opt,name=simplemessage,proto3,oneof"`
}

func (*WebsocketMessage_Install) isWebsocketMessage_Content() {}

func (*WebsocketMessage_Simplemessage) isWebsocketMessage_Content() {}

type Install struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applications   *Install_Applications `protobuf:"bytes,1,opt,name=applications,proto3" json:"applications,omitempty"`
	DeploymentName string                `protobuf:"bytes,3,opt,name=DeploymentName,proto3" json:"DeploymentName,omitempty"`
}

func (x *Install) Reset() {
	*x = Install{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Install) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Install) ProtoMessage() {}

func (x *Install) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Install.ProtoReflect.Descriptor instead.
func (*Install) Descriptor() ([]byte, []int) {
	return file_extensions_proto_rawDescGZIP(), []int{1}
}

func (x *Install) GetApplications() *Install_Applications {
	if x != nil {
		return x.Applications
	}
	return nil
}

func (x *Install) GetDeploymentName() string {
	if x != nil {
		return x.DeploymentName
	}
	return ""
}

type SimpleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operation string `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	Payload   string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *SimpleMessage) Reset() {
	*x = SimpleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMessage) ProtoMessage() {}

func (x *SimpleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMessage.ProtoReflect.Descriptor instead.
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return file_extensions_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleMessage) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *SimpleMessage) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

type ActionMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetadataVersion string                 `protobuf:"bytes,1,opt,name=metadataVersion,proto3" json:"metadataVersion,omitempty"`
	Exectarget      string                 `protobuf:"bytes,2,opt,name=exectarget,proto3" json:"exectarget,omitempty"`
	DeploymentName  string                 `protobuf:"bytes,3,opt,name=deploymentName,proto3" json:"deploymentName,omitempty"`
	Services        []*ActionMeta_Services `protobuf:"bytes,4,rep,name=services,proto3" json:"services,omitempty"`
	Extensions      *ActionMeta_Extensions `protobuf:"bytes,5,opt,name=extensions,proto3" json:"extensions,omitempty"`
	DeploymentType  string                 `protobuf:"bytes,6,opt,name=deploymentType,proto3" json:"deploymentType,omitempty"`
}

func (x *ActionMeta) Reset() {
	*x = ActionMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionMeta) ProtoMessage() {}

func (x *ActionMeta) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionMeta.ProtoReflect.Descriptor instead.
func (*ActionMeta) Descriptor() ([]byte, []int) {
	return file_extensions_proto_rawDescGZIP(), []int{3}
}

func (x *ActionMeta) GetMetadataVersion() string {
	if x != nil {
		return x.MetadataVersion
	}
	return ""
}

func (x *ActionMeta) GetExectarget() string {
	if x != nil {
		return x.Exectarget
	}
	return ""
}

func (x *ActionMeta) GetDeploymentName() string {
	if x != nil {
		return x.DeploymentName
	}
	return ""
}

func (x *ActionMeta) GetServices() []*ActionMeta_Services {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *ActionMeta) GetExtensions() *ActionMeta_Extensions {
	if x != nil {
		return x.Extensions
	}
	return nil
}

func (x *ActionMeta) GetDeploymentType() string {
	if x != nil {
		return x.DeploymentType
	}
	return ""
}

type Install_Applications struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version   string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Variables map[string]string `protobuf:"bytes,3,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Install_Applications) Reset() {
	*x = Install_Applications{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Install_Applications) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Install_Applications) ProtoMessage() {}

func (x *Install_Applications) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Install_Applications.ProtoReflect.Descriptor instead.
func (*Install_Applications) Descriptor() ([]byte, []int) {
	return file_extensions_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Install_Applications) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Install_Applications) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Install_Applications) GetVariables() map[string]string {
	if x != nil {
		return x.Variables
	}
	return nil
}

type ActionMeta_Services struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Source  string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Branch  string `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
}

func (x *ActionMeta_Services) Reset() {
	*x = ActionMeta_Services{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionMeta_Services) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionMeta_Services) ProtoMessage() {}

func (x *ActionMeta_Services) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionMeta_Services.ProtoReflect.Descriptor instead.
func (*ActionMeta_Services) Descriptor() ([]byte, []int) {
	return file_extensions_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ActionMeta_Services) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ActionMeta_Services) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *ActionMeta_Services) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ActionMeta_Services) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

type ActionMeta_Extensions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eks  *ActionMeta_Extensions_Eks        `protobuf:"bytes,1,opt,name=eks,proto3" json:"eks,omitempty"`
	Apis *ActionMeta_Extensions_Apigateway `protobuf:"bytes,2,opt,name=apis,proto3" json:"apis,omitempty"`
}

func (x *ActionMeta_Extensions) Reset() {
	*x = ActionMeta_Extensions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionMeta_Extensions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionMeta_Extensions) ProtoMessage() {}

func (x *ActionMeta_Extensions) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionMeta_Extensions.ProtoReflect.Descriptor instead.
func (*ActionMeta_Extensions) Descriptor() ([]byte, []int) {
	return file_extensions_proto_rawDescGZIP(), []int{3, 1}
}

func (x *ActionMeta_Extensions) GetEks() *ActionMeta_Extensions_Eks {
	if x != nil {
		return x.Eks
	}
	return nil
}

func (x *ActionMeta_Extensions) GetApis() *ActionMeta_Extensions_Apigateway {
	if x != nil {
		return x.Apis
	}
	return nil
}

type ActionMeta_Extensions_Apis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ActionMeta_Extensions_Apis) Reset() {
	*x = ActionMeta_Extensions_Apis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionMeta_Extensions_Apis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionMeta_Extensions_Apis) ProtoMessage() {}

func (x *ActionMeta_Extensions_Apis) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionMeta_Extensions_Apis.ProtoReflect.Descriptor instead.
func (*ActionMeta_Extensions_Apis) Descriptor() ([]byte, []int) {
	return file_extensions_proto_rawDescGZIP(), []int{3, 1, 0}
}

func (x *ActionMeta_Extensions_Apis) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ActionMeta_Extensions_Apigateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apis []*ActionMeta_Extensions_Apis `protobuf:"bytes,1,rep,name=apis,proto3" json:"apis,omitempty"`
}

func (x *ActionMeta_Extensions_Apigateway) Reset() {
	*x = ActionMeta_Extensions_Apigateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionMeta_Extensions_Apigateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionMeta_Extensions_Apigateway) ProtoMessage() {}

func (x *ActionMeta_Extensions_Apigateway) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionMeta_Extensions_Apigateway.ProtoReflect.Descriptor instead.
func (*ActionMeta_Extensions_Apigateway) Descriptor() ([]byte, []int) {
	return file_extensions_proto_rawDescGZIP(), []int{3, 1, 1}
}

func (x *ActionMeta_Extensions_Apigateway) GetApis() []*ActionMeta_Extensions_Apis {
	if x != nil {
		return x.Apis
	}
	return nil
}

type ActionMeta_Extensions_Nodegroups struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Instancetype string `protobuf:"bytes,2,opt,name=instancetype,proto3" json:"instancetype,omitempty"`
	Nodecount    string `protobuf:"bytes,3,opt,name=nodecount,proto3" json:"nodecount,omitempty"`
}

func (x *ActionMeta_Extensions_Nodegroups) Reset() {
	*x = ActionMeta_Extensions_Nodegroups{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionMeta_Extensions_Nodegroups) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionMeta_Extensions_Nodegroups) ProtoMessage() {}

func (x *ActionMeta_Extensions_Nodegroups) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionMeta_Extensions_Nodegroups.ProtoReflect.Descriptor instead.
func (*ActionMeta_Extensions_Nodegroups) Descriptor() ([]byte, []int) {
	return file_extensions_proto_rawDescGZIP(), []int{3, 1, 2}
}

func (x *ActionMeta_Extensions_Nodegroups) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ActionMeta_Extensions_Nodegroups) GetInstancetype() string {
	if x != nil {
		return x.Instancetype
	}
	return ""
}

func (x *ActionMeta_Extensions_Nodegroups) GetNodecount() string {
	if x != nil {
		return x.Nodecount
	}
	return ""
}

type ActionMeta_Extensions_Eks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clustername string                              `protobuf:"bytes,1,opt,name=clustername,proto3" json:"clustername,omitempty"`
	Owner       string                              `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Projectname string                              `protobuf:"bytes,3,opt,name=projectname,proto3" json:"projectname,omitempty"`
	Nodegroups  []*ActionMeta_Extensions_Nodegroups `protobuf:"bytes,4,rep,name=nodegroups,proto3" json:"nodegroups,omitempty"`
}

func (x *ActionMeta_Extensions_Eks) Reset() {
	*x = ActionMeta_Extensions_Eks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extensions_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionMeta_Extensions_Eks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionMeta_Extensions_Eks) ProtoMessage() {}

func (x *ActionMeta_Extensions_Eks) ProtoReflect() protoreflect.Message {
	mi := &file_extensions_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionMeta_Extensions_Eks.ProtoReflect.Descriptor instead.
func (*ActionMeta_Extensions_Eks) Descriptor() ([]byte, []int) {
	return file_extensions_proto_rawDescGZIP(), []int{3, 1, 3}
}

func (x *ActionMeta_Extensions_Eks) GetClustername() string {
	if x != nil {
		return x.Clustername
	}
	return ""
}

func (x *ActionMeta_Extensions_Eks) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ActionMeta_Extensions_Eks) GetProjectname() string {
	if x != nil {
		return x.Projectname
	}
	return ""
}

func (x *ActionMeta_Extensions_Eks) GetNodegroups() []*ActionMeta_Extensions_Nodegroups {
	if x != nil {
		return x.Nodegroups
	}
	return nil
}

var File_extensions_proto protoreflect.FileDescriptor

var file_extensions_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x10, 0x57, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c,
	0x6c, 0x48, 0x00, 0x52, 0x07, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x36, 0x0a, 0x0d,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22,
	0xad, 0x02, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x12, 0x39, 0x0a, 0x0c, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0xbe,
	0x01, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a,
	0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x47, 0x0a, 0x0d, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xd2, 0x06, 0x0a, 0x0a, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x68, 0x0a, 0x08, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x1a, 0xd5, 0x03, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x03, 0x65, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6b, 0x73, 0x52, 0x03, 0x65,
	0x6b, 0x73, 0x12, 0x35, 0x0a, 0x04, 0x61, 0x70, 0x69, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x52, 0x04, 0x61, 0x70, 0x69, 0x73, 0x1a, 0x1a, 0x0a, 0x04, 0x41, 0x70, 0x69,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x3d, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x12, 0x2f, 0x0a, 0x04, 0x61, 0x70, 0x69, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x70, 0x69, 0x73, 0x52, 0x04,
	0x61, 0x70, 0x69, 0x73, 0x1a, 0x62, 0x0a, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x6f, 0x64, 0x65, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0xa2, 0x01, 0x0a, 0x03, 0x45, 0x6b, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x6e, 0x6f,
	0x64, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42, 0x1e, 0x5a,
	0x1c, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_extensions_proto_rawDescOnce sync.Once
	file_extensions_proto_rawDescData = file_extensions_proto_rawDesc
)

func file_extensions_proto_rawDescGZIP() []byte {
	file_extensions_proto_rawDescOnce.Do(func() {
		file_extensions_proto_rawDescData = protoimpl.X.CompressGZIP(file_extensions_proto_rawDescData)
	})
	return file_extensions_proto_rawDescData
}

var file_extensions_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_extensions_proto_goTypes = []interface{}{
	(*WebsocketMessage)(nil),                 // 0: WebsocketMessage
	(*Install)(nil),                          // 1: Install
	(*SimpleMessage)(nil),                    // 2: SimpleMessage
	(*ActionMeta)(nil),                       // 3: ActionMeta
	(*Install_Applications)(nil),             // 4: Install.Applications
	nil,                                      // 5: Install.Applications.VariablesEntry
	(*ActionMeta_Services)(nil),              // 6: ActionMeta.Services
	(*ActionMeta_Extensions)(nil),            // 7: ActionMeta.Extensions
	(*ActionMeta_Extensions_Apis)(nil),       // 8: ActionMeta.Extensions.Apis
	(*ActionMeta_Extensions_Apigateway)(nil), // 9: ActionMeta.Extensions.Apigateway
	(*ActionMeta_Extensions_Nodegroups)(nil), // 10: ActionMeta.Extensions.Nodegroups
	(*ActionMeta_Extensions_Eks)(nil),        // 11: ActionMeta.Extensions.Eks
}
var file_extensions_proto_depIdxs = []int32{
	1,  // 0: WebsocketMessage.install:type_name -> Install
	2,  // 1: WebsocketMessage.simplemessage:type_name -> SimpleMessage
	4,  // 2: Install.applications:type_name -> Install.Applications
	6,  // 3: ActionMeta.services:type_name -> ActionMeta.Services
	7,  // 4: ActionMeta.extensions:type_name -> ActionMeta.Extensions
	5,  // 5: Install.Applications.variables:type_name -> Install.Applications.VariablesEntry
	11, // 6: ActionMeta.Extensions.eks:type_name -> ActionMeta.Extensions.Eks
	9,  // 7: ActionMeta.Extensions.apis:type_name -> ActionMeta.Extensions.Apigateway
	8,  // 8: ActionMeta.Extensions.Apigateway.apis:type_name -> ActionMeta.Extensions.Apis
	10, // 9: ActionMeta.Extensions.Eks.nodegroups:type_name -> ActionMeta.Extensions.Nodegroups
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_extensions_proto_init() }
func file_extensions_proto_init() {
	if File_extensions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extensions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebsocketMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Install); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Install_Applications); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionMeta_Services); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionMeta_Extensions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionMeta_Extensions_Apis); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionMeta_Extensions_Apigateway); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionMeta_Extensions_Nodegroups); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_extensions_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionMeta_Extensions_Eks); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_extensions_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*WebsocketMessage_Install)(nil),
		(*WebsocketMessage_Simplemessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_extensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extensions_proto_goTypes,
		DependencyIndexes: file_extensions_proto_depIdxs,
		MessageInfos:      file_extensions_proto_msgTypes,
	}.Build()
	File_extensions_proto = out.File
	file_extensions_proto_rawDesc = nil
	file_extensions_proto_goTypes = nil
	file_extensions_proto_depIdxs = nil
}
