// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: store/idp.proto

package store

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IdentityProvider_Type int32

const (
	IdentityProvider_TYPE_UNSPECIFIED IdentityProvider_Type = 0
	IdentityProvider_OAUTH2           IdentityProvider_Type = 1
)

// Enum value maps for IdentityProvider_Type.
var (
	IdentityProvider_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "OAUTH2",
	}
	IdentityProvider_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"OAUTH2":           1,
	}
)

func (x IdentityProvider_Type) Enum() *IdentityProvider_Type {
	p := new(IdentityProvider_Type)
	*p = x
	return p
}

func (x IdentityProvider_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IdentityProvider_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_store_idp_proto_enumTypes[0].Descriptor()
}

func (IdentityProvider_Type) Type() protoreflect.EnumType {
	return &file_store_idp_proto_enumTypes[0]
}

func (x IdentityProvider_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IdentityProvider_Type.Descriptor instead.
func (IdentityProvider_Type) EnumDescriptor() ([]byte, []int) {
	return file_store_idp_proto_rawDescGZIP(), []int{0, 0}
}

type IdentityProvider struct {
	state            protoimpl.MessageState  `protogen:"open.v1"`
	Id               int32                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type             IdentityProvider_Type   `protobuf:"varint,3,opt,name=type,proto3,enum=memos.store.IdentityProvider_Type" json:"type,omitempty"`
	IdentifierFilter string                  `protobuf:"bytes,4,opt,name=identifier_filter,json=identifierFilter,proto3" json:"identifier_filter,omitempty"`
	Config           *IdentityProviderConfig `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *IdentityProvider) Reset() {
	*x = IdentityProvider{}
	mi := &file_store_idp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdentityProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProvider) ProtoMessage() {}

func (x *IdentityProvider) ProtoReflect() protoreflect.Message {
	mi := &file_store_idp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProvider.ProtoReflect.Descriptor instead.
func (*IdentityProvider) Descriptor() ([]byte, []int) {
	return file_store_idp_proto_rawDescGZIP(), []int{0}
}

func (x *IdentityProvider) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IdentityProvider) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *IdentityProvider) GetType() IdentityProvider_Type {
	if x != nil {
		return x.Type
	}
	return IdentityProvider_TYPE_UNSPECIFIED
}

func (x *IdentityProvider) GetIdentifierFilter() string {
	if x != nil {
		return x.IdentifierFilter
	}
	return ""
}

func (x *IdentityProvider) GetConfig() *IdentityProviderConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type IdentityProviderConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Config:
	//
	//	*IdentityProviderConfig_Oauth2Config
	Config        isIdentityProviderConfig_Config `protobuf_oneof:"config"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IdentityProviderConfig) Reset() {
	*x = IdentityProviderConfig{}
	mi := &file_store_idp_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdentityProviderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentityProviderConfig) ProtoMessage() {}

func (x *IdentityProviderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_store_idp_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentityProviderConfig.ProtoReflect.Descriptor instead.
func (*IdentityProviderConfig) Descriptor() ([]byte, []int) {
	return file_store_idp_proto_rawDescGZIP(), []int{1}
}

func (x *IdentityProviderConfig) GetConfig() isIdentityProviderConfig_Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *IdentityProviderConfig) GetOauth2Config() *OAuth2Config {
	if x != nil {
		if x, ok := x.Config.(*IdentityProviderConfig_Oauth2Config); ok {
			return x.Oauth2Config
		}
	}
	return nil
}

type isIdentityProviderConfig_Config interface {
	isIdentityProviderConfig_Config()
}

type IdentityProviderConfig_Oauth2Config struct {
	Oauth2Config *OAuth2Config `protobuf:"bytes,1,opt,name=oauth2_config,json=oauth2Config,proto3,oneof"`
}

func (*IdentityProviderConfig_Oauth2Config) isIdentityProviderConfig_Config() {}

type FieldMapping struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Identifier    string                 `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	DisplayName   string                 `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FieldMapping) Reset() {
	*x = FieldMapping{}
	mi := &file_store_idp_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMapping) ProtoMessage() {}

func (x *FieldMapping) ProtoReflect() protoreflect.Message {
	mi := &file_store_idp_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMapping.ProtoReflect.Descriptor instead.
func (*FieldMapping) Descriptor() ([]byte, []int) {
	return file_store_idp_proto_rawDescGZIP(), []int{2}
}

func (x *FieldMapping) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *FieldMapping) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *FieldMapping) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type OAuth2Config struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret  string                 `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	AuthUrl       string                 `protobuf:"bytes,3,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	TokenUrl      string                 `protobuf:"bytes,4,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	UserInfoUrl   string                 `protobuf:"bytes,5,opt,name=user_info_url,json=userInfoUrl,proto3" json:"user_info_url,omitempty"`
	Scopes        []string               `protobuf:"bytes,6,rep,name=scopes,proto3" json:"scopes,omitempty"`
	FieldMapping  *FieldMapping          `protobuf:"bytes,7,opt,name=field_mapping,json=fieldMapping,proto3" json:"field_mapping,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OAuth2Config) Reset() {
	*x = OAuth2Config{}
	mi := &file_store_idp_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OAuth2Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2Config) ProtoMessage() {}

func (x *OAuth2Config) ProtoReflect() protoreflect.Message {
	mi := &file_store_idp_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2Config.ProtoReflect.Descriptor instead.
func (*OAuth2Config) Descriptor() ([]byte, []int) {
	return file_store_idp_proto_rawDescGZIP(), []int{3}
}

func (x *OAuth2Config) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *OAuth2Config) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *OAuth2Config) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

func (x *OAuth2Config) GetTokenUrl() string {
	if x != nil {
		return x.TokenUrl
	}
	return ""
}

func (x *OAuth2Config) GetUserInfoUrl() string {
	if x != nil {
		return x.UserInfoUrl
	}
	return ""
}

func (x *OAuth2Config) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *OAuth2Config) GetFieldMapping() *FieldMapping {
	if x != nil {
		return x.FieldMapping
	}
	return nil
}

var File_store_idp_proto protoreflect.FileDescriptor

var file_store_idp_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x64, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x82,
	0x02, 0x0a, 0x10, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d,
	0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x28, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x41, 0x55, 0x54, 0x48,
	0x32, 0x10, 0x01, 0x22, 0x64, 0x0a, 0x16, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x40, 0x0a,
	0x0d, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48,
	0x00, 0x52, 0x0c, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x67, 0x0a, 0x0c, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x22, 0x84, 0x02, 0x0a, 0x0c, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x55, 0x72,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x42, 0x93, 0x01, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x08, 0x49,
	0x64, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x6d,
	0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0xa2, 0x02, 0x03, 0x4d, 0x53, 0x58, 0xaa, 0x02, 0x0b, 0x4d, 0x65, 0x6d,
	0x6f, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xca, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73,
	0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0xe2, 0x02, 0x17, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0c, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_store_idp_proto_rawDescOnce sync.Once
	file_store_idp_proto_rawDescData []byte
)

func file_store_idp_proto_rawDescGZIP() []byte {
	file_store_idp_proto_rawDescOnce.Do(func() {
		file_store_idp_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_store_idp_proto_rawDesc), len(file_store_idp_proto_rawDesc)))
	})
	return file_store_idp_proto_rawDescData
}

var file_store_idp_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_idp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_store_idp_proto_goTypes = []any{
	(IdentityProvider_Type)(0),     // 0: memos.store.IdentityProvider.Type
	(*IdentityProvider)(nil),       // 1: memos.store.IdentityProvider
	(*IdentityProviderConfig)(nil), // 2: memos.store.IdentityProviderConfig
	(*FieldMapping)(nil),           // 3: memos.store.FieldMapping
	(*OAuth2Config)(nil),           // 4: memos.store.OAuth2Config
}
var file_store_idp_proto_depIdxs = []int32{
	0, // 0: memos.store.IdentityProvider.type:type_name -> memos.store.IdentityProvider.Type
	2, // 1: memos.store.IdentityProvider.config:type_name -> memos.store.IdentityProviderConfig
	4, // 2: memos.store.IdentityProviderConfig.oauth2_config:type_name -> memos.store.OAuth2Config
	3, // 3: memos.store.OAuth2Config.field_mapping:type_name -> memos.store.FieldMapping
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_store_idp_proto_init() }
func file_store_idp_proto_init() {
	if File_store_idp_proto != nil {
		return
	}
	file_store_idp_proto_msgTypes[1].OneofWrappers = []any{
		(*IdentityProviderConfig_Oauth2Config)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_store_idp_proto_rawDesc), len(file_store_idp_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_idp_proto_goTypes,
		DependencyIndexes: file_store_idp_proto_depIdxs,
		EnumInfos:         file_store_idp_proto_enumTypes,
		MessageInfos:      file_store_idp_proto_msgTypes,
	}.Build()
	File_store_idp_proto = out.File
	file_store_idp_proto_goTypes = nil
	file_store_idp_proto_depIdxs = nil
}
