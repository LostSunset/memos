// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: store/user_setting.proto

package store

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

type UserSettingKey int32

const (
	UserSettingKey_USER_SETTING_KEY_UNSPECIFIED UserSettingKey = 0
	// Access tokens for the user.
	UserSettingKey_ACCESS_TOKENS UserSettingKey = 1
	// The locale of the user.
	UserSettingKey_LOCALE UserSettingKey = 2
	// The appearance of the user.
	UserSettingKey_APPEARANCE UserSettingKey = 3
	// The visibility of the memo.
	UserSettingKey_MEMO_VISIBILITY UserSettingKey = 4
)

// Enum value maps for UserSettingKey.
var (
	UserSettingKey_name = map[int32]string{
		0: "USER_SETTING_KEY_UNSPECIFIED",
		1: "ACCESS_TOKENS",
		2: "LOCALE",
		3: "APPEARANCE",
		4: "MEMO_VISIBILITY",
	}
	UserSettingKey_value = map[string]int32{
		"USER_SETTING_KEY_UNSPECIFIED": 0,
		"ACCESS_TOKENS":                1,
		"LOCALE":                       2,
		"APPEARANCE":                   3,
		"MEMO_VISIBILITY":              4,
	}
)

func (x UserSettingKey) Enum() *UserSettingKey {
	p := new(UserSettingKey)
	*p = x
	return p
}

func (x UserSettingKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserSettingKey) Descriptor() protoreflect.EnumDescriptor {
	return file_store_user_setting_proto_enumTypes[0].Descriptor()
}

func (UserSettingKey) Type() protoreflect.EnumType {
	return &file_store_user_setting_proto_enumTypes[0]
}

func (x UserSettingKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserSettingKey.Descriptor instead.
func (UserSettingKey) EnumDescriptor() ([]byte, []int) {
	return file_store_user_setting_proto_rawDescGZIP(), []int{0}
}

type UserSetting struct {
	state  protoimpl.MessageState `protogen:"open.v1"`
	UserId int32                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Key    UserSettingKey         `protobuf:"varint,2,opt,name=key,proto3,enum=memos.store.UserSettingKey" json:"key,omitempty"`
	// Types that are valid to be assigned to Value:
	//
	//	*UserSetting_AccessTokens
	//	*UserSetting_Locale
	//	*UserSetting_Appearance
	//	*UserSetting_MemoVisibility
	Value         isUserSetting_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserSetting) Reset() {
	*x = UserSetting{}
	mi := &file_store_user_setting_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSetting) ProtoMessage() {}

func (x *UserSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_user_setting_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSetting.ProtoReflect.Descriptor instead.
func (*UserSetting) Descriptor() ([]byte, []int) {
	return file_store_user_setting_proto_rawDescGZIP(), []int{0}
}

func (x *UserSetting) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserSetting) GetKey() UserSettingKey {
	if x != nil {
		return x.Key
	}
	return UserSettingKey_USER_SETTING_KEY_UNSPECIFIED
}

func (x *UserSetting) GetValue() isUserSetting_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *UserSetting) GetAccessTokens() *AccessTokensUserSetting {
	if x != nil {
		if x, ok := x.Value.(*UserSetting_AccessTokens); ok {
			return x.AccessTokens
		}
	}
	return nil
}

func (x *UserSetting) GetLocale() string {
	if x != nil {
		if x, ok := x.Value.(*UserSetting_Locale); ok {
			return x.Locale
		}
	}
	return ""
}

func (x *UserSetting) GetAppearance() string {
	if x != nil {
		if x, ok := x.Value.(*UserSetting_Appearance); ok {
			return x.Appearance
		}
	}
	return ""
}

func (x *UserSetting) GetMemoVisibility() string {
	if x != nil {
		if x, ok := x.Value.(*UserSetting_MemoVisibility); ok {
			return x.MemoVisibility
		}
	}
	return ""
}

type isUserSetting_Value interface {
	isUserSetting_Value()
}

type UserSetting_AccessTokens struct {
	AccessTokens *AccessTokensUserSetting `protobuf:"bytes,3,opt,name=access_tokens,json=accessTokens,proto3,oneof"`
}

type UserSetting_Locale struct {
	Locale string `protobuf:"bytes,4,opt,name=locale,proto3,oneof"`
}

type UserSetting_Appearance struct {
	Appearance string `protobuf:"bytes,5,opt,name=appearance,proto3,oneof"`
}

type UserSetting_MemoVisibility struct {
	MemoVisibility string `protobuf:"bytes,6,opt,name=memo_visibility,json=memoVisibility,proto3,oneof"`
}

func (*UserSetting_AccessTokens) isUserSetting_Value() {}

func (*UserSetting_Locale) isUserSetting_Value() {}

func (*UserSetting_Appearance) isUserSetting_Value() {}

func (*UserSetting_MemoVisibility) isUserSetting_Value() {}

type AccessTokensUserSetting struct {
	state         protoimpl.MessageState                 `protogen:"open.v1"`
	AccessTokens  []*AccessTokensUserSetting_AccessToken `protobuf:"bytes,1,rep,name=access_tokens,json=accessTokens,proto3" json:"access_tokens,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccessTokensUserSetting) Reset() {
	*x = AccessTokensUserSetting{}
	mi := &file_store_user_setting_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessTokensUserSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessTokensUserSetting) ProtoMessage() {}

func (x *AccessTokensUserSetting) ProtoReflect() protoreflect.Message {
	mi := &file_store_user_setting_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessTokensUserSetting.ProtoReflect.Descriptor instead.
func (*AccessTokensUserSetting) Descriptor() ([]byte, []int) {
	return file_store_user_setting_proto_rawDescGZIP(), []int{1}
}

func (x *AccessTokensUserSetting) GetAccessTokens() []*AccessTokensUserSetting_AccessToken {
	if x != nil {
		return x.AccessTokens
	}
	return nil
}

type AccessTokensUserSetting_AccessToken struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The access token is a JWT token.
	// Including expiration time, issuer, etc.
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// A description for the access token.
	Description   string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccessTokensUserSetting_AccessToken) Reset() {
	*x = AccessTokensUserSetting_AccessToken{}
	mi := &file_store_user_setting_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessTokensUserSetting_AccessToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessTokensUserSetting_AccessToken) ProtoMessage() {}

func (x *AccessTokensUserSetting_AccessToken) ProtoReflect() protoreflect.Message {
	mi := &file_store_user_setting_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessTokensUserSetting_AccessToken.ProtoReflect.Descriptor instead.
func (*AccessTokensUserSetting_AccessToken) Descriptor() ([]byte, []int) {
	return file_store_user_setting_proto_rawDescGZIP(), []int{1, 0}
}

func (x *AccessTokensUserSetting_AccessToken) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AccessTokensUserSetting_AccessToken) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_store_user_setting_proto protoreflect.FileDescriptor

var file_store_user_setting_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x65, 0x6d, 0x6f,
	0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x92, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2d, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e,
	0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x4b, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x06,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x65, 0x61, 0x72,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x70,
	0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0f, 0x6d, 0x65, 0x6d, 0x6f,
	0x5f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0e, 0x6d, 0x65, 0x6d, 0x6f, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc4, 0x01, 0x0a,
	0x17, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x55, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x30, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x1a,
	0x52, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2a, 0x76, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x1c, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x53, 0x45,
	0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f,
	0x43, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x50, 0x50, 0x45, 0x41, 0x52,
	0x41, 0x4e, 0x43, 0x45, 0x10, 0x03, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x4d, 0x4f, 0x5f, 0x56,
	0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x04, 0x42, 0x9b, 0x01, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x42,
	0x10, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x75, 0x73, 0x65, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0xa2, 0x02,
	0x03, 0x4d, 0x53, 0x58, 0xaa, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0xca, 0x02, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0xe2, 0x02, 0x17, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x5c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x4d, 0x65, 0x6d,
	0x6f, 0x73, 0x3a, 0x3a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_store_user_setting_proto_rawDescOnce sync.Once
	file_store_user_setting_proto_rawDescData = file_store_user_setting_proto_rawDesc
)

func file_store_user_setting_proto_rawDescGZIP() []byte {
	file_store_user_setting_proto_rawDescOnce.Do(func() {
		file_store_user_setting_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_user_setting_proto_rawDescData)
	})
	return file_store_user_setting_proto_rawDescData
}

var file_store_user_setting_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_store_user_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_store_user_setting_proto_goTypes = []any{
	(UserSettingKey)(0),                         // 0: memos.store.UserSettingKey
	(*UserSetting)(nil),                         // 1: memos.store.UserSetting
	(*AccessTokensUserSetting)(nil),             // 2: memos.store.AccessTokensUserSetting
	(*AccessTokensUserSetting_AccessToken)(nil), // 3: memos.store.AccessTokensUserSetting.AccessToken
}
var file_store_user_setting_proto_depIdxs = []int32{
	0, // 0: memos.store.UserSetting.key:type_name -> memos.store.UserSettingKey
	2, // 1: memos.store.UserSetting.access_tokens:type_name -> memos.store.AccessTokensUserSetting
	3, // 2: memos.store.AccessTokensUserSetting.access_tokens:type_name -> memos.store.AccessTokensUserSetting.AccessToken
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_store_user_setting_proto_init() }
func file_store_user_setting_proto_init() {
	if File_store_user_setting_proto != nil {
		return
	}
	file_store_user_setting_proto_msgTypes[0].OneofWrappers = []any{
		(*UserSetting_AccessTokens)(nil),
		(*UserSetting_Locale)(nil),
		(*UserSetting_Appearance)(nil),
		(*UserSetting_MemoVisibility)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_user_setting_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_user_setting_proto_goTypes,
		DependencyIndexes: file_store_user_setting_proto_depIdxs,
		EnumInfos:         file_store_user_setting_proto_enumTypes,
		MessageInfos:      file_store_user_setting_proto_msgTypes,
	}.Build()
	File_store_user_setting_proto = out.File
	file_store_user_setting_proto_rawDesc = nil
	file_store_user_setting_proto_goTypes = nil
	file_store_user_setting_proto_depIdxs = nil
}
