// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: store/data_source.proto

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

type DataSourceExternalSecret_SecretType int32

const (
	DataSourceExternalSecret_SAECRET_TYPE_UNSPECIFIED DataSourceExternalSecret_SecretType = 0
	// ref: https://developer.hashicorp.com/vault/api-docs/secret/kv/kv-v2
	DataSourceExternalSecret_VAULT_KV_V2 DataSourceExternalSecret_SecretType = 1
	// ref: https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html
	DataSourceExternalSecret_AWS_SECRETS_MANAGER DataSourceExternalSecret_SecretType = 2
)

// Enum value maps for DataSourceExternalSecret_SecretType.
var (
	DataSourceExternalSecret_SecretType_name = map[int32]string{
		0: "SAECRET_TYPE_UNSPECIFIED",
		1: "VAULT_KV_V2",
		2: "AWS_SECRETS_MANAGER",
	}
	DataSourceExternalSecret_SecretType_value = map[string]int32{
		"SAECRET_TYPE_UNSPECIFIED": 0,
		"VAULT_KV_V2":              1,
		"AWS_SECRETS_MANAGER":      2,
	}
)

func (x DataSourceExternalSecret_SecretType) Enum() *DataSourceExternalSecret_SecretType {
	p := new(DataSourceExternalSecret_SecretType)
	*p = x
	return p
}

func (x DataSourceExternalSecret_SecretType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSourceExternalSecret_SecretType) Descriptor() protoreflect.EnumDescriptor {
	return file_store_data_source_proto_enumTypes[0].Descriptor()
}

func (DataSourceExternalSecret_SecretType) Type() protoreflect.EnumType {
	return &file_store_data_source_proto_enumTypes[0]
}

func (x DataSourceExternalSecret_SecretType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSourceExternalSecret_SecretType.Descriptor instead.
func (DataSourceExternalSecret_SecretType) EnumDescriptor() ([]byte, []int) {
	return file_store_data_source_proto_rawDescGZIP(), []int{0, 0}
}

type DataSourceExternalSecret_AuthType int32

const (
	DataSourceExternalSecret_AUTH_TYPE_UNSPECIFIED DataSourceExternalSecret_AuthType = 0
	// ref: https://developer.hashicorp.com/vault/docs/auth/token
	DataSourceExternalSecret_TOKEN DataSourceExternalSecret_AuthType = 1
	// ref: https://developer.hashicorp.com/vault/docs/auth/approle
	DataSourceExternalSecret_APP_ROLE DataSourceExternalSecret_AuthType = 2
	// ref: https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/#specifying-credentials
	DataSourceExternalSecret_AWS_ENVIRONMENT DataSourceExternalSecret_AuthType = 3
)

// Enum value maps for DataSourceExternalSecret_AuthType.
var (
	DataSourceExternalSecret_AuthType_name = map[int32]string{
		0: "AUTH_TYPE_UNSPECIFIED",
		1: "TOKEN",
		2: "APP_ROLE",
		3: "AWS_ENVIRONMENT",
	}
	DataSourceExternalSecret_AuthType_value = map[string]int32{
		"AUTH_TYPE_UNSPECIFIED": 0,
		"TOKEN":                 1,
		"APP_ROLE":              2,
		"AWS_ENVIRONMENT":       3,
	}
)

func (x DataSourceExternalSecret_AuthType) Enum() *DataSourceExternalSecret_AuthType {
	p := new(DataSourceExternalSecret_AuthType)
	*p = x
	return p
}

func (x DataSourceExternalSecret_AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSourceExternalSecret_AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_store_data_source_proto_enumTypes[1].Descriptor()
}

func (DataSourceExternalSecret_AuthType) Type() protoreflect.EnumType {
	return &file_store_data_source_proto_enumTypes[1]
}

func (x DataSourceExternalSecret_AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSourceExternalSecret_AuthType.Descriptor instead.
func (DataSourceExternalSecret_AuthType) EnumDescriptor() ([]byte, []int) {
	return file_store_data_source_proto_rawDescGZIP(), []int{0, 1}
}

type DataSourceExternalSecret_AppRoleAuthOption_SecretType int32

const (
	DataSourceExternalSecret_AppRoleAuthOption_SECRET_TYPE_UNSPECIFIED DataSourceExternalSecret_AppRoleAuthOption_SecretType = 0
	DataSourceExternalSecret_AppRoleAuthOption_PLAIN                   DataSourceExternalSecret_AppRoleAuthOption_SecretType = 1
	DataSourceExternalSecret_AppRoleAuthOption_ENVIRONMENT             DataSourceExternalSecret_AppRoleAuthOption_SecretType = 2
)

// Enum value maps for DataSourceExternalSecret_AppRoleAuthOption_SecretType.
var (
	DataSourceExternalSecret_AppRoleAuthOption_SecretType_name = map[int32]string{
		0: "SECRET_TYPE_UNSPECIFIED",
		1: "PLAIN",
		2: "ENVIRONMENT",
	}
	DataSourceExternalSecret_AppRoleAuthOption_SecretType_value = map[string]int32{
		"SECRET_TYPE_UNSPECIFIED": 0,
		"PLAIN":                   1,
		"ENVIRONMENT":             2,
	}
)

func (x DataSourceExternalSecret_AppRoleAuthOption_SecretType) Enum() *DataSourceExternalSecret_AppRoleAuthOption_SecretType {
	p := new(DataSourceExternalSecret_AppRoleAuthOption_SecretType)
	*p = x
	return p
}

func (x DataSourceExternalSecret_AppRoleAuthOption_SecretType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSourceExternalSecret_AppRoleAuthOption_SecretType) Descriptor() protoreflect.EnumDescriptor {
	return file_store_data_source_proto_enumTypes[2].Descriptor()
}

func (DataSourceExternalSecret_AppRoleAuthOption_SecretType) Type() protoreflect.EnumType {
	return &file_store_data_source_proto_enumTypes[2]
}

func (x DataSourceExternalSecret_AppRoleAuthOption_SecretType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSourceExternalSecret_AppRoleAuthOption_SecretType.Descriptor instead.
func (DataSourceExternalSecret_AppRoleAuthOption_SecretType) EnumDescriptor() ([]byte, []int) {
	return file_store_data_source_proto_rawDescGZIP(), []int{0, 0, 0}
}

type DataSourceExternalSecret struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretType DataSourceExternalSecret_SecretType `protobuf:"varint,1,opt,name=secret_type,json=secretType,proto3,enum=bytebase.store.DataSourceExternalSecret_SecretType" json:"secret_type,omitempty"`
	Url        string                              `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	AuthType   DataSourceExternalSecret_AuthType   `protobuf:"varint,3,opt,name=auth_type,json=authType,proto3,enum=bytebase.store.DataSourceExternalSecret_AuthType" json:"auth_type,omitempty"`
	// Types that are assignable to AuthOption:
	//
	//	*DataSourceExternalSecret_AppRole
	//	*DataSourceExternalSecret_Token
	AuthOption isDataSourceExternalSecret_AuthOption `protobuf_oneof:"auth_option"`
	// engine name is the name for secret engine.
	EngineName string `protobuf:"bytes,6,opt,name=engine_name,json=engineName,proto3" json:"engine_name,omitempty"`
	// the secret name in the engine to store the password.
	SecretName string `protobuf:"bytes,7,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	// the key name for the password.
	PasswordKeyName string `protobuf:"bytes,8,opt,name=password_key_name,json=passwordKeyName,proto3" json:"password_key_name,omitempty"`
}

func (x *DataSourceExternalSecret) Reset() {
	*x = DataSourceExternalSecret{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_data_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceExternalSecret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceExternalSecret) ProtoMessage() {}

func (x *DataSourceExternalSecret) ProtoReflect() protoreflect.Message {
	mi := &file_store_data_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceExternalSecret.ProtoReflect.Descriptor instead.
func (*DataSourceExternalSecret) Descriptor() ([]byte, []int) {
	return file_store_data_source_proto_rawDescGZIP(), []int{0}
}

func (x *DataSourceExternalSecret) GetSecretType() DataSourceExternalSecret_SecretType {
	if x != nil {
		return x.SecretType
	}
	return DataSourceExternalSecret_SAECRET_TYPE_UNSPECIFIED
}

func (x *DataSourceExternalSecret) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *DataSourceExternalSecret) GetAuthType() DataSourceExternalSecret_AuthType {
	if x != nil {
		return x.AuthType
	}
	return DataSourceExternalSecret_AUTH_TYPE_UNSPECIFIED
}

func (m *DataSourceExternalSecret) GetAuthOption() isDataSourceExternalSecret_AuthOption {
	if m != nil {
		return m.AuthOption
	}
	return nil
}

func (x *DataSourceExternalSecret) GetAppRole() *DataSourceExternalSecret_AppRoleAuthOption {
	if x, ok := x.GetAuthOption().(*DataSourceExternalSecret_AppRole); ok {
		return x.AppRole
	}
	return nil
}

func (x *DataSourceExternalSecret) GetToken() string {
	if x, ok := x.GetAuthOption().(*DataSourceExternalSecret_Token); ok {
		return x.Token
	}
	return ""
}

func (x *DataSourceExternalSecret) GetEngineName() string {
	if x != nil {
		return x.EngineName
	}
	return ""
}

func (x *DataSourceExternalSecret) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

func (x *DataSourceExternalSecret) GetPasswordKeyName() string {
	if x != nil {
		return x.PasswordKeyName
	}
	return ""
}

type isDataSourceExternalSecret_AuthOption interface {
	isDataSourceExternalSecret_AuthOption()
}

type DataSourceExternalSecret_AppRole struct {
	AppRole *DataSourceExternalSecret_AppRoleAuthOption `protobuf:"bytes,4,opt,name=app_role,json=appRole,proto3,oneof"`
}

type DataSourceExternalSecret_Token struct {
	Token string `protobuf:"bytes,5,opt,name=token,proto3,oneof"`
}

func (*DataSourceExternalSecret_AppRole) isDataSourceExternalSecret_AuthOption() {}

func (*DataSourceExternalSecret_Token) isDataSourceExternalSecret_AuthOption() {}

type DataSourceOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// srv is a boolean flag that indicates whether the host is a DNS SRV record.
	Srv bool `protobuf:"varint,1,opt,name=srv,proto3" json:"srv,omitempty"`
	// authentication_database is the database name to authenticate against, which stores the user credentials.
	AuthenticationDatabase string `protobuf:"bytes,2,opt,name=authentication_database,json=authenticationDatabase,proto3" json:"authentication_database,omitempty"`
	// sid and service_name are used for Oracle.
	Sid         string `protobuf:"bytes,3,opt,name=sid,proto3" json:"sid,omitempty"`
	ServiceName string `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// SSH related
	// The hostname of the SSH server agent.
	SshHost string `protobuf:"bytes,5,opt,name=ssh_host,json=sshHost,proto3" json:"ssh_host,omitempty"`
	// The port of the SSH server agent. It's 22 typically.
	SshPort string `protobuf:"bytes,6,opt,name=ssh_port,json=sshPort,proto3" json:"ssh_port,omitempty"`
	// The user to login the server.
	SshUser string `protobuf:"bytes,7,opt,name=ssh_user,json=sshUser,proto3" json:"ssh_user,omitempty"`
	// The password to login the server. If it's empty string, no password is required.
	SshObfuscatedPassword string `protobuf:"bytes,8,opt,name=ssh_obfuscated_password,json=sshObfuscatedPassword,proto3" json:"ssh_obfuscated_password,omitempty"`
	// The private key to login the server. If it's empty string, we will use the system default private key from os.Getenv("SSH_AUTH_SOCK").
	SshObfuscatedPrivateKey string `protobuf:"bytes,9,opt,name=ssh_obfuscated_private_key,json=sshObfuscatedPrivateKey,proto3" json:"ssh_obfuscated_private_key,omitempty"`
	// PKCS#8 private key in PEM format. If it's empty string, no private key is required.
	// Used for authentication when connecting to the data source.
	AuthenticationPrivateKeyObfuscated string                    `protobuf:"bytes,10,opt,name=authentication_private_key_obfuscated,json=authenticationPrivateKeyObfuscated,proto3" json:"authentication_private_key_obfuscated,omitempty"`
	ExternalSecret                     *DataSourceExternalSecret `protobuf:"bytes,11,opt,name=external_secret,json=externalSecret,proto3" json:"external_secret,omitempty"`
}

func (x *DataSourceOptions) Reset() {
	*x = DataSourceOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_data_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceOptions) ProtoMessage() {}

func (x *DataSourceOptions) ProtoReflect() protoreflect.Message {
	mi := &file_store_data_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceOptions.ProtoReflect.Descriptor instead.
func (*DataSourceOptions) Descriptor() ([]byte, []int) {
	return file_store_data_source_proto_rawDescGZIP(), []int{1}
}

func (x *DataSourceOptions) GetSrv() bool {
	if x != nil {
		return x.Srv
	}
	return false
}

func (x *DataSourceOptions) GetAuthenticationDatabase() string {
	if x != nil {
		return x.AuthenticationDatabase
	}
	return ""
}

func (x *DataSourceOptions) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *DataSourceOptions) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *DataSourceOptions) GetSshHost() string {
	if x != nil {
		return x.SshHost
	}
	return ""
}

func (x *DataSourceOptions) GetSshPort() string {
	if x != nil {
		return x.SshPort
	}
	return ""
}

func (x *DataSourceOptions) GetSshUser() string {
	if x != nil {
		return x.SshUser
	}
	return ""
}

func (x *DataSourceOptions) GetSshObfuscatedPassword() string {
	if x != nil {
		return x.SshObfuscatedPassword
	}
	return ""
}

func (x *DataSourceOptions) GetSshObfuscatedPrivateKey() string {
	if x != nil {
		return x.SshObfuscatedPrivateKey
	}
	return ""
}

func (x *DataSourceOptions) GetAuthenticationPrivateKeyObfuscated() string {
	if x != nil {
		return x.AuthenticationPrivateKeyObfuscated
	}
	return ""
}

func (x *DataSourceOptions) GetExternalSecret() *DataSourceExternalSecret {
	if x != nil {
		return x.ExternalSecret
	}
	return nil
}

type DataSourceExternalSecret_AppRoleAuthOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId string `protobuf:"bytes,1,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	// the secret id for the role without ttl.
	SecretId string                                                `protobuf:"bytes,2,opt,name=secret_id,json=secretId,proto3" json:"secret_id,omitempty"`
	Type     DataSourceExternalSecret_AppRoleAuthOption_SecretType `protobuf:"varint,3,opt,name=type,proto3,enum=bytebase.store.DataSourceExternalSecret_AppRoleAuthOption_SecretType" json:"type,omitempty"`
	// The path where the approle auth method is mounted.
	MountPath string `protobuf:"bytes,4,opt,name=mount_path,json=mountPath,proto3" json:"mount_path,omitempty"`
}

func (x *DataSourceExternalSecret_AppRoleAuthOption) Reset() {
	*x = DataSourceExternalSecret_AppRoleAuthOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_data_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceExternalSecret_AppRoleAuthOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceExternalSecret_AppRoleAuthOption) ProtoMessage() {}

func (x *DataSourceExternalSecret_AppRoleAuthOption) ProtoReflect() protoreflect.Message {
	mi := &file_store_data_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceExternalSecret_AppRoleAuthOption.ProtoReflect.Descriptor instead.
func (*DataSourceExternalSecret_AppRoleAuthOption) Descriptor() ([]byte, []int) {
	return file_store_data_source_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DataSourceExternalSecret_AppRoleAuthOption) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *DataSourceExternalSecret_AppRoleAuthOption) GetSecretId() string {
	if x != nil {
		return x.SecretId
	}
	return ""
}

func (x *DataSourceExternalSecret_AppRoleAuthOption) GetType() DataSourceExternalSecret_AppRoleAuthOption_SecretType {
	if x != nil {
		return x.Type
	}
	return DataSourceExternalSecret_AppRoleAuthOption_SECRET_TYPE_UNSPECIFIED
}

func (x *DataSourceExternalSecret_AppRoleAuthOption) GetMountPath() string {
	if x != nil {
		return x.MountPath
	}
	return ""
}

var File_store_data_source_proto protoreflect.FileDescriptor

var file_store_data_source_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0xf8, 0x06, 0x0a, 0x18, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x54, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x62, 0x79,
	0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x4e,
	0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x31, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x57,
	0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x6f,
	0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x07,
	0x61, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x8a, 0x02,
	0x0a, 0x11, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x49, 0x64, 0x12, 0x59, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x22, 0x45, 0x0a, 0x0a, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x50, 0x4c, 0x41, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x56,
	0x49, 0x52, 0x4f, 0x4e, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x22, 0x54, 0x0a, 0x0a, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x41, 0x45, 0x43,
	0x52, 0x45, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x41, 0x55, 0x4c, 0x54, 0x5f,
	0x4b, 0x56, 0x5f, 0x56, 0x32, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x57, 0x53, 0x5f, 0x53,
	0x45, 0x43, 0x52, 0x45, 0x54, 0x53, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x02,
	0x22, 0x53, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15,
	0x41, 0x55, 0x54, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x4f, 0x4b, 0x45, 0x4e,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x10, 0x02,
	0x12, 0x13, 0x0a, 0x0f, 0x41, 0x57, 0x53, 0x5f, 0x45, 0x4e, 0x56, 0x49, 0x52, 0x4f, 0x4e, 0x4d,
	0x45, 0x4e, 0x54, 0x10, 0x03, 0x42, 0x0d, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xff, 0x03, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x72,
	0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x72, 0x76, 0x12, 0x37, 0x0a, 0x17,
	0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x73,
	0x68, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x73,
	0x68, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x73, 0x68, 0x5f, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x73, 0x68, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x73, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x17, 0x73,
	0x73, 0x68, 0x5f, 0x6f, 0x62, 0x66, 0x75, 0x73, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x73,
	0x68, 0x4f, 0x62, 0x66, 0x75, 0x73, 0x63, 0x61, 0x74, 0x65, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x3b, 0x0a, 0x1a, 0x73, 0x73, 0x68, 0x5f, 0x6f, 0x62, 0x66, 0x75, 0x73,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x73, 0x73, 0x68, 0x4f, 0x62, 0x66, 0x75,
	0x73, 0x63, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x51, 0x0a, 0x25, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x6f,
	0x62, 0x66, 0x75, 0x73, 0x63, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x22, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x4f, 0x62, 0x66, 0x75, 0x73, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x51, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62,
	0x79, 0x74, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42, 0x14, 0x5a, 0x12, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_store_data_source_proto_rawDescOnce sync.Once
	file_store_data_source_proto_rawDescData = file_store_data_source_proto_rawDesc
)

func file_store_data_source_proto_rawDescGZIP() []byte {
	file_store_data_source_proto_rawDescOnce.Do(func() {
		file_store_data_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_data_source_proto_rawDescData)
	})
	return file_store_data_source_proto_rawDescData
}

var file_store_data_source_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_store_data_source_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_store_data_source_proto_goTypes = []interface{}{
	(DataSourceExternalSecret_SecretType)(0),                   // 0: bytebase.store.DataSourceExternalSecret.SecretType
	(DataSourceExternalSecret_AuthType)(0),                     // 1: bytebase.store.DataSourceExternalSecret.AuthType
	(DataSourceExternalSecret_AppRoleAuthOption_SecretType)(0), // 2: bytebase.store.DataSourceExternalSecret.AppRoleAuthOption.SecretType
	(*DataSourceExternalSecret)(nil),                           // 3: bytebase.store.DataSourceExternalSecret
	(*DataSourceOptions)(nil),                                  // 4: bytebase.store.DataSourceOptions
	(*DataSourceExternalSecret_AppRoleAuthOption)(nil),         // 5: bytebase.store.DataSourceExternalSecret.AppRoleAuthOption
}
var file_store_data_source_proto_depIdxs = []int32{
	0, // 0: bytebase.store.DataSourceExternalSecret.secret_type:type_name -> bytebase.store.DataSourceExternalSecret.SecretType
	1, // 1: bytebase.store.DataSourceExternalSecret.auth_type:type_name -> bytebase.store.DataSourceExternalSecret.AuthType
	5, // 2: bytebase.store.DataSourceExternalSecret.app_role:type_name -> bytebase.store.DataSourceExternalSecret.AppRoleAuthOption
	3, // 3: bytebase.store.DataSourceOptions.external_secret:type_name -> bytebase.store.DataSourceExternalSecret
	2, // 4: bytebase.store.DataSourceExternalSecret.AppRoleAuthOption.type:type_name -> bytebase.store.DataSourceExternalSecret.AppRoleAuthOption.SecretType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_store_data_source_proto_init() }
func file_store_data_source_proto_init() {
	if File_store_data_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_data_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceExternalSecret); i {
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
		file_store_data_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceOptions); i {
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
		file_store_data_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceExternalSecret_AppRoleAuthOption); i {
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
	file_store_data_source_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DataSourceExternalSecret_AppRole)(nil),
		(*DataSourceExternalSecret_Token)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_store_data_source_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_data_source_proto_goTypes,
		DependencyIndexes: file_store_data_source_proto_depIdxs,
		EnumInfos:         file_store_data_source_proto_enumTypes,
		MessageInfos:      file_store_data_source_proto_msgTypes,
	}.Build()
	File_store_data_source_proto = out.File
	file_store_data_source_proto_rawDesc = nil
	file_store_data_source_proto_goTypes = nil
	file_store_data_source_proto_depIdxs = nil
}
