// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1/user.proto

package mysql

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type GlobalPermission int32

const (
	GlobalPermission_GLOBAL_PERMISSION_UNSPECIFIED GlobalPermission = 0
	// Enables use of the SHOW MASTER STATUS, SHOW SLAVE STATUS, and SHOW BINARY LOGS statements.
	GlobalPermission_REPLICATION_CLIENT GlobalPermission = 1
	// Enables the account to request updates that have been made to databases on the master server,
	// using the SHOW SLAVE HOSTS, SHOW RELAYLOG EVENTS, and SHOW BINLOG EVENTS statements.
	GlobalPermission_REPLICATION_SLAVE GlobalPermission = 2
)

var GlobalPermission_name = map[int32]string{
	0: "GLOBAL_PERMISSION_UNSPECIFIED",
	1: "REPLICATION_CLIENT",
	2: "REPLICATION_SLAVE",
}

var GlobalPermission_value = map[string]int32{
	"GLOBAL_PERMISSION_UNSPECIFIED": 0,
	"REPLICATION_CLIENT":            1,
	"REPLICATION_SLAVE":             2,
}

func (x GlobalPermission) String() string {
	return proto.EnumName(GlobalPermission_name, int32(x))
}

func (GlobalPermission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5965312201724c83, []int{0}
}

type AuthPlugin int32

const (
	AuthPlugin_AUTH_PLUGIN_UNSPECIFIED AuthPlugin = 0
	// Use [Native Pluggable Authentication](https://dev.mysql.com/doc/refman/8.0/en/native-pluggable-authentication.html).
	AuthPlugin_MYSQL_NATIVE_PASSWORD AuthPlugin = 1
	// Use [Caching SHA-2 Pluggable Authentication](https://dev.mysql.com/doc/refman/8.0/en/caching-sha2-pluggable-authentication.html).
	AuthPlugin_CACHING_SHA2_PASSWORD AuthPlugin = 2
	// Use [SHA-256 Pluggable Authentication](https://dev.mysql.com/doc/refman/8.0/en/sha256-pluggable-authentication.html).
	AuthPlugin_SHA256_PASSWORD AuthPlugin = 3
)

var AuthPlugin_name = map[int32]string{
	0: "AUTH_PLUGIN_UNSPECIFIED",
	1: "MYSQL_NATIVE_PASSWORD",
	2: "CACHING_SHA2_PASSWORD",
	3: "SHA256_PASSWORD",
}

var AuthPlugin_value = map[string]int32{
	"AUTH_PLUGIN_UNSPECIFIED": 0,
	"MYSQL_NATIVE_PASSWORD":   1,
	"CACHING_SHA2_PASSWORD":   2,
	"SHA256_PASSWORD":         3,
}

func (x AuthPlugin) String() string {
	return proto.EnumName(AuthPlugin_name, int32(x))
}

func (AuthPlugin) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5965312201724c83, []int{1}
}

type Permission_Privilege int32

const (
	Permission_PRIVILEGE_UNSPECIFIED Permission_Privilege = 0
	// All privileges that can be made available to the user.
	Permission_ALL_PRIVILEGES Permission_Privilege = 1
	// Altering tables.
	Permission_ALTER Permission_Privilege = 2
	// Altering stored routines (stored procedures and functions).
	Permission_ALTER_ROUTINE Permission_Privilege = 3
	// Creating tables or indexes.
	Permission_CREATE Permission_Privilege = 4
	// Creating stored routines.
	Permission_CREATE_ROUTINE Permission_Privilege = 5
	// Creating temporary tables.
	Permission_CREATE_TEMPORARY_TABLES Permission_Privilege = 6
	// Creating views.
	Permission_CREATE_VIEW Permission_Privilege = 7
	// Deleting tables.
	Permission_DELETE Permission_Privilege = 8
	// Removing tables or views.
	Permission_DROP Permission_Privilege = 9
	// Creating, altering, dropping, or displaying events for the Event Scheduler.
	Permission_EVENT Permission_Privilege = 10
	// Executing stored routines.
	Permission_EXECUTE Permission_Privilege = 11
	// Creating and removing indexes.
	Permission_INDEX Permission_Privilege = 12
	// Inserting rows into the database.
	Permission_INSERT Permission_Privilege = 13
	// Using LOCK TABLES statement for tables available with SELECT privilege.
	Permission_LOCK_TABLES Permission_Privilege = 14
	// Selecting rows from tables.
	//
	// Some SELECT statements can be allowed without the SELECT privilege. All
	// statements that read column values require the SELECT privilege. See
	// details in [MySQL documentation](https://dev.mysql.com/doc/refman/5.7/en/privileges-provided.html#priv_select).
	Permission_SELECT Permission_Privilege = 15
	// Using the SHOW CREATE VIEW statement. Also needed for views used with EXPLAIN.
	Permission_SHOW_VIEW Permission_Privilege = 16
	// Creating, removing, executing, or displaying triggers for a table.
	Permission_TRIGGER Permission_Privilege = 17
	// Updating rows in the database.
	Permission_UPDATE Permission_Privilege = 18
)

var Permission_Privilege_name = map[int32]string{
	0:  "PRIVILEGE_UNSPECIFIED",
	1:  "ALL_PRIVILEGES",
	2:  "ALTER",
	3:  "ALTER_ROUTINE",
	4:  "CREATE",
	5:  "CREATE_ROUTINE",
	6:  "CREATE_TEMPORARY_TABLES",
	7:  "CREATE_VIEW",
	8:  "DELETE",
	9:  "DROP",
	10: "EVENT",
	11: "EXECUTE",
	12: "INDEX",
	13: "INSERT",
	14: "LOCK_TABLES",
	15: "SELECT",
	16: "SHOW_VIEW",
	17: "TRIGGER",
	18: "UPDATE",
}

var Permission_Privilege_value = map[string]int32{
	"PRIVILEGE_UNSPECIFIED":   0,
	"ALL_PRIVILEGES":          1,
	"ALTER":                   2,
	"ALTER_ROUTINE":           3,
	"CREATE":                  4,
	"CREATE_ROUTINE":          5,
	"CREATE_TEMPORARY_TABLES": 6,
	"CREATE_VIEW":             7,
	"DELETE":                  8,
	"DROP":                    9,
	"EVENT":                   10,
	"EXECUTE":                 11,
	"INDEX":                   12,
	"INSERT":                  13,
	"LOCK_TABLES":             14,
	"SELECT":                  15,
	"SHOW_VIEW":               16,
	"TRIGGER":                 17,
	"UPDATE":                  18,
}

func (x Permission_Privilege) String() string {
	return proto.EnumName(Permission_Privilege_name, int32(x))
}

func (Permission_Privilege) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5965312201724c83, []int{1, 0}
}

// A MySQL user. For more information, see
// the [documentation](/docs/managed-mysql/concepts).
type User struct {
	// Name of the MySQL user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the MySQL cluster the user belongs to.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Set of permissions granted to the user.
	Permissions          []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_5965312201724c83, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *User) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Permission struct {
	// Name of the database that the permission grants access to.
	DatabaseName string `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	// Roles granted to the user within the database.
	Roles                []Permission_Privilege `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=yandex.cloud.mdb.mysql.v1.Permission_Privilege" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_5965312201724c83, []int{1}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

func (m *Permission) GetRoles() []Permission_Privilege {
	if m != nil {
		return m.Roles
	}
	return nil
}

type ConnectionLimits struct {
	// The maximum permitted number of user questions per hour.
	MaxQuestions *wrappers.Int64Value `protobuf:"bytes,1,opt,name=max_questions,json=maxQuestions,proto3" json:"max_questions,omitempty"`
	// The maximum permitted number of user updates per hour.
	MaxUpdates *wrappers.Int64Value `protobuf:"bytes,2,opt,name=max_updates,json=maxUpdates,proto3" json:"max_updates,omitempty"`
	// The maximum permitted number of simultaneous client connections per hour.
	MaxConnections *wrappers.Int64Value `protobuf:"bytes,3,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// The maximum number of simultaneous connections permitted to any given MySQL user account.
	MaxUserConnections   *wrappers.Int64Value `protobuf:"bytes,4,opt,name=max_user_connections,json=maxUserConnections,proto3" json:"max_user_connections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ConnectionLimits) Reset()         { *m = ConnectionLimits{} }
func (m *ConnectionLimits) String() string { return proto.CompactTextString(m) }
func (*ConnectionLimits) ProtoMessage()    {}
func (*ConnectionLimits) Descriptor() ([]byte, []int) {
	return fileDescriptor_5965312201724c83, []int{2}
}

func (m *ConnectionLimits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionLimits.Unmarshal(m, b)
}
func (m *ConnectionLimits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionLimits.Marshal(b, m, deterministic)
}
func (m *ConnectionLimits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionLimits.Merge(m, src)
}
func (m *ConnectionLimits) XXX_Size() int {
	return xxx_messageInfo_ConnectionLimits.Size(m)
}
func (m *ConnectionLimits) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionLimits.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionLimits proto.InternalMessageInfo

func (m *ConnectionLimits) GetMaxQuestions() *wrappers.Int64Value {
	if m != nil {
		return m.MaxQuestions
	}
	return nil
}

func (m *ConnectionLimits) GetMaxUpdates() *wrappers.Int64Value {
	if m != nil {
		return m.MaxUpdates
	}
	return nil
}

func (m *ConnectionLimits) GetMaxConnections() *wrappers.Int64Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *ConnectionLimits) GetMaxUserConnections() *wrappers.Int64Value {
	if m != nil {
		return m.MaxUserConnections
	}
	return nil
}

type UserSpec struct {
	// Name of the MySQL user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Password of the MySQL user.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Set of permissions to grant to the user.
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// Set of global permissions to grant to the user.
	GlobalPermissions []GlobalPermission `protobuf:"varint,4,rep,packed,name=global_permissions,json=globalPermissions,proto3,enum=yandex.cloud.mdb.mysql.v1.GlobalPermission" json:"global_permissions,omitempty"`
	// Set of user connection limits.
	ConnectionLimits *ConnectionLimits `protobuf:"bytes,5,opt,name=connection_limits,json=connectionLimits,proto3" json:"connection_limits,omitempty"`
	// User authentication plugin.
	AuthenticationPlugin AuthPlugin `protobuf:"varint,6,opt,name=authentication_plugin,json=authenticationPlugin,proto3,enum=yandex.cloud.mdb.mysql.v1.AuthPlugin" json:"authentication_plugin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UserSpec) Reset()         { *m = UserSpec{} }
func (m *UserSpec) String() string { return proto.CompactTextString(m) }
func (*UserSpec) ProtoMessage()    {}
func (*UserSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_5965312201724c83, []int{3}
}

func (m *UserSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSpec.Unmarshal(m, b)
}
func (m *UserSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSpec.Marshal(b, m, deterministic)
}
func (m *UserSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSpec.Merge(m, src)
}
func (m *UserSpec) XXX_Size() int {
	return xxx_messageInfo_UserSpec.Size(m)
}
func (m *UserSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UserSpec proto.InternalMessageInfo

func (m *UserSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserSpec) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserSpec) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *UserSpec) GetGlobalPermissions() []GlobalPermission {
	if m != nil {
		return m.GlobalPermissions
	}
	return nil
}

func (m *UserSpec) GetConnectionLimits() *ConnectionLimits {
	if m != nil {
		return m.ConnectionLimits
	}
	return nil
}

func (m *UserSpec) GetAuthenticationPlugin() AuthPlugin {
	if m != nil {
		return m.AuthenticationPlugin
	}
	return AuthPlugin_AUTH_PLUGIN_UNSPECIFIED
}

func init() {
	proto.RegisterEnum("yandex.cloud.mdb.mysql.v1.GlobalPermission", GlobalPermission_name, GlobalPermission_value)
	proto.RegisterEnum("yandex.cloud.mdb.mysql.v1.AuthPlugin", AuthPlugin_name, AuthPlugin_value)
	proto.RegisterEnum("yandex.cloud.mdb.mysql.v1.Permission_Privilege", Permission_Privilege_name, Permission_Privilege_value)
	proto.RegisterType((*User)(nil), "yandex.cloud.mdb.mysql.v1.User")
	proto.RegisterType((*Permission)(nil), "yandex.cloud.mdb.mysql.v1.Permission")
	proto.RegisterType((*ConnectionLimits)(nil), "yandex.cloud.mdb.mysql.v1.ConnectionLimits")
	proto.RegisterType((*UserSpec)(nil), "yandex.cloud.mdb.mysql.v1.UserSpec")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1/user.proto", fileDescriptor_5965312201724c83)
}

var fileDescriptor_5965312201724c83 = []byte{
	// 938 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4b, 0x6f, 0xdb, 0x46,
	0x17, 0xfd, 0xf4, 0xf0, 0x43, 0x57, 0x96, 0x3d, 0x9e, 0x2f, 0x6e, 0xe5, 0x04, 0x2e, 0x5c, 0xb7,
	0x05, 0x5c, 0x17, 0xa2, 0x42, 0xa5, 0x0d, 0x52, 0xb4, 0x29, 0x40, 0x51, 0x53, 0x99, 0x28, 0x4d,
	0x31, 0x43, 0x4a, 0x76, 0x5c, 0x14, 0x04, 0x25, 0x4e, 0x65, 0x02, 0x7c, 0x28, 0x7c, 0x38, 0x4e,
	0x97, 0x59, 0xe6, 0xd7, 0x74, 0xd1, 0x7f, 0xd0, 0x85, 0x0d, 0xf4, 0x8f, 0x74, 0xdd, 0x65, 0x57,
	0xc5, 0x90, 0xb6, 0x24, 0x0b, 0x88, 0xe1, 0x02, 0xdd, 0x0d, 0xce, 0x39, 0xf7, 0xcc, 0x9d, 0x3b,
	0xf7, 0xce, 0xc0, 0xa7, 0x6f, 0xec, 0xc0, 0x61, 0x17, 0xcd, 0x91, 0x17, 0xa6, 0x4e, 0xd3, 0x77,
	0x86, 0x4d, 0xff, 0x4d, 0xfc, 0xca, 0x6b, 0x9e, 0x8b, 0xcd, 0x34, 0x66, 0x91, 0x30, 0x89, 0xc2,
	0x24, 0xc4, 0xdb, 0xb9, 0x4a, 0xc8, 0x54, 0x82, 0xef, 0x0c, 0x85, 0x4c, 0x25, 0x9c, 0x8b, 0x0f,
	0x3f, 0x1a, 0x87, 0xe1, 0xd8, 0x63, 0xcd, 0x4c, 0x38, 0x4c, 0x7f, 0x6e, 0xbe, 0x8e, 0xec, 0xc9,
	0x84, 0x45, 0x71, 0x1e, 0xfa, 0x70, 0xe7, 0xd6, 0x06, 0xe7, 0xb6, 0xe7, 0x3a, 0x76, 0xe2, 0x86,
	0x41, 0x4e, 0xef, 0xbd, 0x2d, 0x40, 0xb9, 0x1f, 0xb3, 0x08, 0x63, 0x28, 0x07, 0xb6, 0xcf, 0xea,
	0x85, 0xdd, 0xc2, 0x7e, 0x85, 0x66, 0x6b, 0xbc, 0x03, 0x30, 0xf2, 0xd2, 0x38, 0x61, 0x91, 0xe5,
	0x3a, 0xf5, 0x62, 0xc6, 0x54, 0xae, 0x11, 0xc5, 0xc1, 0x5d, 0xa8, 0x4e, 0x58, 0xe4, 0xbb, 0x71,
	0xec, 0x86, 0x41, 0x5c, 0x2f, 0xed, 0x96, 0xf6, 0xab, 0xad, 0xcf, 0x84, 0xf7, 0xe6, 0x2a, 0xe8,
	0x53, 0x35, 0x9d, 0x8f, 0xdc, 0xfb, 0xad, 0x04, 0x30, 0xe3, 0xf0, 0x27, 0x50, 0x73, 0xec, 0xc4,
	0x1e, 0xda, 0x31, 0xb3, 0xe6, 0x72, 0x5a, 0xbb, 0x01, 0x35, 0x9e, 0x9b, 0x06, 0x4b, 0x51, 0xe8,
	0xb1, 0xb8, 0x5e, 0xdc, 0x2d, 0xed, 0xaf, 0xb7, 0x9a, 0xf7, 0xda, 0x56, 0xd0, 0x23, 0xf7, 0xdc,
	0xf5, 0xd8, 0x98, 0xb5, 0x57, 0xde, 0x5e, 0x89, 0xa5, 0xef, 0x9e, 0x8b, 0x34, 0xb7, 0xd9, 0xfb,
	0xb5, 0x08, 0x95, 0x29, 0x8b, 0xb7, 0x61, 0x4b, 0xa7, 0xca, 0x40, 0x51, 0x49, 0x97, 0x58, 0x7d,
	0xcd, 0xd0, 0x89, 0xac, 0x7c, 0xaf, 0x90, 0x0e, 0xfa, 0x1f, 0xc6, 0xb0, 0x2e, 0xa9, 0xaa, 0x35,
	0xa5, 0x0d, 0x54, 0xc0, 0x15, 0x58, 0x92, 0x54, 0x93, 0x50, 0x54, 0xc4, 0x9b, 0x50, 0xcb, 0x96,
	0x16, 0xed, 0xf5, 0x4d, 0x45, 0x23, 0xa8, 0x84, 0x01, 0x96, 0x65, 0x4a, 0x24, 0x93, 0xa0, 0x32,
	0x8f, 0xce, 0xd7, 0x53, 0x7e, 0x09, 0x3f, 0x82, 0x0f, 0xaf, 0x31, 0x93, 0x1c, 0xe9, 0x3d, 0x2a,
	0xd1, 0x97, 0x96, 0x29, 0xb5, 0x55, 0x62, 0xa0, 0x65, 0xbc, 0x01, 0xd5, 0x6b, 0x72, 0xa0, 0x90,
	0x63, 0xb4, 0xc2, 0xdd, 0x3a, 0x44, 0x25, 0x26, 0x41, 0xab, 0x78, 0x15, 0xca, 0x1d, 0xda, 0xd3,
	0x51, 0x85, 0x67, 0x40, 0x06, 0x44, 0x33, 0x11, 0xe0, 0x2a, 0xac, 0x90, 0x13, 0x22, 0xf7, 0x4d,
	0x82, 0xaa, 0x1c, 0x57, 0xb4, 0x0e, 0x39, 0x41, 0x6b, 0x3c, 0x50, 0xd1, 0x0c, 0x42, 0x4d, 0x54,
	0xe3, 0xae, 0x6a, 0x4f, 0xfe, 0xe1, 0x66, 0x9b, 0x75, 0x4e, 0x1a, 0x44, 0x25, 0xb2, 0x89, 0x36,
	0x70, 0x0d, 0x2a, 0xc6, 0x61, 0xef, 0x38, 0xdf, 0x10, 0x71, 0x3f, 0x93, 0x2a, 0xdd, 0x2e, 0xa1,
	0x68, 0x93, 0xeb, 0xfa, 0x7a, 0x87, 0x9f, 0x05, 0xef, 0xfd, 0x51, 0x04, 0x24, 0x87, 0x41, 0xc0,
	0x46, 0xbc, 0xa1, 0x54, 0xd7, 0x77, 0x93, 0x18, 0x2b, 0x50, 0xf3, 0xed, 0x0b, 0xeb, 0x55, 0xca,
	0xe2, 0x24, 0x6b, 0x0b, 0x7e, 0x79, 0xd5, 0xd6, 0x23, 0x21, 0xef, 0x53, 0xe1, 0xa6, 0x4f, 0x05,
	0x25, 0x48, 0x9e, 0x7e, 0x39, 0xb0, 0xbd, 0x94, 0xb5, 0x57, 0xfe, 0xbe, 0xe4, 0x77, 0xf1, 0x98,
	0xae, 0xf9, 0xf6, 0xc5, 0x8b, 0x9b, 0x48, 0x4c, 0xa0, 0xca, 0xad, 0xd2, 0x89, 0x63, 0x27, 0xd9,
	0x45, 0xdf, 0xdf, 0x08, 0x7c, 0xfb, 0xa2, 0x9f, 0xc7, 0xe1, 0x23, 0xd8, 0xe0, 0x36, 0xa3, 0x69,
	0xa6, 0xbc, 0x55, 0xef, 0x6f, 0xb5, 0xee, 0xdb, 0x17, 0xb3, 0x53, 0xc6, 0x78, 0x00, 0x0f, 0xb2,
	0xac, 0x62, 0x16, 0xdd, 0xf2, 0x2c, 0xff, 0x0b, 0x4f, 0xcc, 0xd3, 0x8b, 0x59, 0x34, 0xe7, 0xbb,
	0xf7, 0x7b, 0x09, 0x56, 0x39, 0x66, 0x4c, 0xd8, 0x08, 0x8b, 0xf3, 0xd3, 0xd8, 0xde, 0xf9, 0xf3,
	0x52, 0x2c, 0xfc, 0x75, 0x29, 0xd6, 0x7e, 0xb4, 0x1b, 0xbf, 0x48, 0x8d, 0xd3, 0xc7, 0x8d, 0xaf,
	0xad, 0x9f, 0x0e, 0xde, 0x5d, 0x89, 0xe5, 0x6f, 0x9f, 0x3f, 0x69, 0x5d, 0x0f, 0xeb, 0xe7, 0xb0,
	0x3a, 0xb1, 0xe3, 0xf8, 0x75, 0x18, 0x5d, 0x8f, 0x6a, 0xbb, 0xc6, 0xc3, 0xde, 0x5d, 0x89, 0x4b,
	0xcf, 0x1a, 0x62, 0xeb, 0x19, 0x9d, 0xd2, 0xff, 0xd9, 0xe0, 0xe2, 0x53, 0xc0, 0x63, 0x2f, 0x1c,
	0xda, 0x9e, 0x35, 0xef, 0x57, 0xce, 0x26, 0xf2, 0x8b, 0x3b, 0xfc, 0xba, 0x59, 0xd0, 0x9c, 0xeb,
	0xe6, 0x78, 0x01, 0x89, 0xf1, 0x09, 0x6c, 0xce, 0xca, 0x6b, 0x79, 0x59, 0x77, 0xd5, 0x97, 0xb2,
	0x22, 0xdf, 0x65, 0xbd, 0xd8, 0x90, 0x14, 0x8d, 0x16, 0x5b, 0xf4, 0x14, 0xb6, 0xec, 0x34, 0x39,
	0x63, 0x41, 0xe2, 0x8e, 0xb2, 0xb7, 0xd0, 0x9a, 0x78, 0xe9, 0xd8, 0x0d, 0xea, 0xcb, 0xbb, 0x85,
	0xfd, 0xf5, 0x3b, 0x0b, 0x21, 0xa5, 0xc9, 0x99, 0x9e, 0x89, 0xe9, 0x83, 0xdb, 0x1e, 0x39, 0x7a,
	0xe0, 0x00, 0x5a, 0x3c, 0x1c, 0xfe, 0x18, 0x76, 0xba, 0x6a, 0xaf, 0x2d, 0xa9, 0x96, 0x4e, 0xe8,
	0x91, 0x62, 0x18, 0x4a, 0x4f, 0x5b, 0x78, 0x54, 0x3e, 0x00, 0x4c, 0x89, 0xae, 0x2a, 0xb2, 0x64,
	0x72, 0x52, 0x56, 0x15, 0x3e, 0xcb, 0x05, 0xbc, 0x05, 0x9b, 0xf3, 0xb8, 0xa1, 0x4a, 0x03, 0x82,
	0x8a, 0x07, 0x09, 0xc0, 0x2c, 0x13, 0xfe, 0x7e, 0x48, 0x7d, 0xf3, 0xd0, 0xd2, 0xd5, 0x7e, 0x57,
	0x59, 0x74, 0xde, 0x86, 0xad, 0xa3, 0x97, 0xc6, 0x0b, 0xd5, 0xd2, 0x24, 0x53, 0x19, 0x10, 0x4b,
	0x97, 0x0c, 0xe3, 0xb8, 0x47, 0x3b, 0xa8, 0xc0, 0x29, 0x59, 0x92, 0x0f, 0x15, 0xad, 0x6b, 0x19,
	0x87, 0x52, 0x6b, 0x46, 0x15, 0xf1, 0xff, 0x61, 0x83, 0x43, 0x5f, 0x3d, 0x9d, 0x81, 0xa5, 0xb6,
	0x03, 0x3b, 0xb7, 0x2a, 0x63, 0x4f, 0xdc, 0x5b, 0xd5, 0x39, 0x95, 0xc7, 0x6e, 0x72, 0x96, 0x0e,
	0x85, 0x51, 0xe8, 0x37, 0x73, 0x65, 0x23, 0xff, 0x76, 0xc6, 0x61, 0x63, 0xcc, 0x82, 0x6c, 0x24,
	0x9a, 0xef, 0xfd, 0xf0, 0xbe, 0xc9, 0x16, 0xc3, 0xe5, 0x4c, 0xf6, 0xe4, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0x1b, 0xde, 0x57, 0x1a, 0x07, 0x00, 0x00,
}
