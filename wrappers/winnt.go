/*
 * Copyright (c) 2014 MongoDB, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the license is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package wrappers

import (
	"unsafe"
)

const (
	VER_SUITE_SMALLBUSINESS            = 0x00000001
	VER_SUITE_ENTERPRISE               = 0x00000002
	VER_SUITE_BACKOFFICE               = 0x00000004
	VER_SUITE_COMMUNICATIONS           = 0x00000008
	VER_SUITE_TERMINAL                 = 0x00000010
	VER_SUITE_SMALLBUSINESS_RESTRICTED = 0x00000020
	VER_SUITE_EMBEDDEDNT               = 0x00000040
	VER_SUITE_DATACENTER               = 0x00000080
	VER_SUITE_SINGLEUSERTS             = 0x00000100
	VER_SUITE_PERSONAL                 = 0x00000200
	VER_SUITE_BLADE                    = 0x00000400
	VER_SUITE_EMBEDDED_RESTRICTED      = 0x00000800
	VER_SUITE_SECURITY_APPLIANCE       = 0x00001000
	VER_SUITE_STORAGE_SERVER           = 0x00002000
	VER_SUITE_COMPUTE_SERVER           = 0x00004000
	VER_SUITE_WH_SERVER                = 0x00008000
)

type SIDIdentifierAuthority struct {
	Value [6]byte
}

var (
	SECURITY_NULL_SID_AUTHORITY        = SIDIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 0}}
	SECURITY_WORLD_SID_AUTHORITY       = SIDIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 1}}
	SECURITY_LOCAL_SID_AUTHORITY       = SIDIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 2}}
	SECURITY_CREATOR_SID_AUTHORITY     = SIDIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 3}}
	SECURITY_NT_AUTHORITY              = SIDIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 5}}
	SECURITY_MANDATORY_LABEL_AUTHORITY = SIDIdentifierAuthority{[6]byte{0, 0, 0, 0, 0, 16}}
)

const (
	SECURITY_NULL_RID          = 0x00000000
	SECURITY_WORLD_RID         = 0x00000000
	SECURITY_LOCAL_RID         = 0x00000000
	SECURITY_LOCAL_LOGON_RID   = 0x00000001
	SECURITY_CREATOR_OWNER_RID = 0x00000000
	SECURITY_CREATOR_GROUP_RID = 0x00000001
)

const (
	SECURITY_DIALUP_RID                 = 0x00000001
	SECURITY_NETWORK_RID                = 0x00000002
	SECURITY_BATCH_RID                  = 0x00000003
	SECURITY_INTERACTIVE_RID            = 0x00000004
	SECURITY_LOGON_IDS_RID              = 0x00000005
	SECURITY_SERVICE_RID                = 0x00000006
	SECURITY_ANONYMOUS_LOGON_RID        = 0x00000007
	SECURITY_PROXY_RID                  = 0x00000008
	SECURITY_ENTERPRISE_CONTROLLERS_RID = 0x00000009
	SECURITY_PRINCIPAL_SELF_RID         = 0x0000000A
	SECURITY_AUTHENTICATED_USER_RID     = 0x0000000B
	SECURITY_RESTRICTED_CODE_RID        = 0x0000000C
	SECURITY_TERMINAL_SERVER_RID        = 0x0000000D
	SECURITY_LOCAL_SYSTEM_RID           = 0x00000012
	SECURITY_LOCAL_SERVICE_RID          = 0x00000013
	SECURITY_NETWORK_SERVICE_RID        = 0x00000014
	SECURITY_NT_NON_UNIQUE              = 0x00000015
	SECURITY_BUILTIN_DOMAIN_RID         = 0x00000020
)

const (
	DOMAIN_USER_RID_ADMIN = 0x000001F4
	DOMAIN_USER_RID_GUEST = 0x000001F5
)

const (
	DOMAIN_GROUP_RID_ADMINS               = 0x00000200
	DOMAIN_GROUP_RID_USERS                = 0x00000201
	DOMAIN_GROUP_RID_GUESTS               = 0x00000202
	DOMAIN_GROUP_RID_COMPUTERS            = 0x00000203
	DOMAIN_GROUP_RID_CONTROLLERS          = 0x00000204
	DOMAIN_GROUP_RID_CERT_ADMINS          = 0x00000205
	DOMAIN_GROUP_RID_SCHEMA_ADMINS        = 0x00000206
	DOMAIN_GROUP_RID_ENTERPRISE_ADMINS    = 0x00000207
	DOMAIN_GROUP_RID_POLICY_ADMINS        = 0x00000208
	DOMAIN_GROUP_RID_READONLY_CONTROLLERS = 0x00000209
)

const (
	DOMAIN_ALIAS_RID_ADMINS                         = 0x00000220
	DOMAIN_ALIAS_RID_USERS                          = 0x00000221
	DOMAIN_ALIAS_RID_GUESTS                         = 0x00000222
	DOMAIN_ALIAS_RID_POWER_USERS                    = 0x00000223
	DOMAIN_ALIAS_RID_ACCOUNT_OPS                    = 0x00000224
	DOMAIN_ALIAS_RID_SYSTEM_OPS                     = 0x00000225
	DOMAIN_ALIAS_RID_PRINT_OPS                      = 0x00000226
	DOMAIN_ALIAS_RID_BACKUP_OPS                     = 0x00000227
	DOMAIN_ALIAS_RID_REPLICATOR                     = 0x00000228
	DOMAIN_ALIAS_RID_RAS_SERVERS                    = 0x00000229
	DOMAIN_ALIAS_RID_PREW2KCOMPACCESS               = 0x0000022A
	DOMAIN_ALIAS_RID_REMOTE_DESKTOP_USERS           = 0x0000022B
	DOMAIN_ALIAS_RID_NETWORK_CONFIGURATION_OPS      = 0x0000022C
	DOMAIN_ALIAS_RID_INCOMING_FOREST_TRUST_BUILDERS = 0x0000022D
	DOMAIN_ALIAS_RID_MONITORING_USERS               = 0x0000022E
	DOMAIN_ALIAS_RID_LOGGING_USERS                  = 0x0000022F
	DOMAIN_ALIAS_RID_AUTHORIZATIONACCESS            = 0x00000230
	DOMAIN_ALIAS_RID_TS_LICENSE_SERVERS             = 0x00000231
	DOMAIN_ALIAS_RID_DCOM_USERS                     = 0x00000232
	DOMAIN_ALIAS_RID_IUSERS                         = 0x00000238
	DOMAIN_ALIAS_RID_CRYPTO_OPERATORS               = 0x00000239
	DOMAIN_ALIAS_RID_CACHEABLE_PRINCIPALS_GROUP     = 0x0000023B
	DOMAIN_ALIAS_RID_NON_CACHEABLE_PRINCIPALS_GROUP = 0x0000023C
	DOMAIN_ALIAS_RID_EVENT_LOG_READERS_GROUP        = 0x0000023D
	DOMAIN_ALIAS_RID_CERTSVC_DCOM_ACCESS_GROUP      = 0x0000023E
)

const (
	SECURITY_MANDATORY_UNTRUSTED_RID         = 0x00000000
	SECURITY_MANDATORY_LOW_RID               = 0x00001000
	SECURITY_MANDATORY_MEDIUM_RID            = 0x00002000
	SECURITY_MANDATORY_MEDIUM_PLUS_RID       = SECURITY_MANDATORY_MEDIUM_RID + 0x00000100
	SECURITY_MANDATORY_HIGH_RID              = 0x00003000
	SECURITY_MANDATORY_SYSTEM_RID            = 0x00004000
	SECURITY_MANDATORY_PROTECTED_PROCESS_RID = 0x00005000
)

const (
	OWNER_SECURITY_INFORMATION            = 0x00000001
	GROUP_SECURITY_INFORMATION            = 0x00000002
	DACL_SECURITY_INFORMATION             = 0x00000004
	SACL_SECURITY_INFORMATION             = 0x00000008
	LABEL_SECURITY_INFORMATION            = 0x00000010
	PROTECTED_DACL_SECURITY_INFORMATION   = 0x80000000
	PROTECTED_SACL_SECURITY_INFORMATION   = 0x40000000
	UNPROTECTED_DACL_SECURITY_INFORMATION = 0x20000000
	UNPROTECTED_SACL_SECURITY_INFORMATION = 0x10000000
)

type OSVERSIONINFO struct {
	OSVersionInfoSize uint32
	MajorVersion      uint32
	MinorVersion      uint32
	BuildNumber       uint32
	PlatformId        uint32
	CSDVersion        [128]uint16
}

type OSVERSIONINFOEX struct {
	OSVERSIONINFO
	ServicePackMajor uint16
	ServicePackMinor uint16
	SuiteMask        uint16
	ProductType      uint8
	Reserved         uint8
}

const (
	VER_EQUAL         = 1
	VER_GREATER       = 2
	VER_GREATER_EQUAL = 3
	VER_LESS          = 4
	VER_LESS_EQUAL    = 5
	VER_AND           = 6
	VER_OR            = 7
)

const (
	VER_MINORVERSION     = 0x00000001
	VER_MAJORVERSION     = 0x00000002
	VER_BUILDNUMBER      = 0x00000004
	VER_PLATFORMID       = 0x00000008
	VER_SERVICEPACKMINOR = 0x00000010
	VER_SERVICEPACKMAJOR = 0x00000020
	VER_SUITENAME        = 0x00000040
	VER_PRODUCT_TYPE     = 0x00000080
)

const (
	VER_NT_WORKSTATION       = 0x00000001
	VER_NT_DOMAIN_CONTROLLER = 0x00000002
	VER_NT_SERVER            = 0x00000003
)

const (
	VER_PLATFORM_WIN32s        = 0
	VER_PLATFORM_WIN32_WINDOWS = 1
	VER_PLATFORM_WIN32_NT      = 2
)

const (
	SERVICE_KERNEL_DRIVER       = 0x00000001
	SERVICE_FILE_SYSTEM_DRIVER  = 0x00000002
	SERVICE_ADAPTER             = 0x00000004
	SERVICE_RECOGNIZER_DRIVER   = 0x00000008
	SERVICE_WIN32_OWN_PROCESS   = 0x00000010
	SERVICE_WIN32_SHARE_PROCESS = 0x00000020
	SERVICE_INTERACTIVE_PROCESS = 0x00000100
)

const (
	SERVICE_BOOT_START   = 0x00000000
	SERVICE_SYSTEM_START = 0x00000001
	SERVICE_AUTO_START   = 0x00000002
	SERVICE_DEMAND_START = 0x00000003
	SERVICE_DISABLED     = 0x00000004
)

const (
	SERVICE_ERROR_IGNORE   = 0x00000000
	SERVICE_ERROR_NORMAL   = 0x00000001
	SERVICE_ERROR_SEVERE   = 0x00000002
	SERVICE_ERROR_CRITICAL = 0x00000003
)

var (
	procRtlMoveMemory       = modkernel32.NewProc("RtlMoveMemory")
	procVerSetConditionMask = modkernel32.NewProc("VerSetConditionMask")
)

func RtlMoveMemory(destination *byte, source *byte, length uintptr) {
	procRtlMoveMemory.Call(
		uintptr(unsafe.Pointer(destination)),
		uintptr(unsafe.Pointer(source)),
		length)
}

func VerSetConditionMask(conditionMask uint64, typeBitMask uint32, condition uint8) uint64 {
	r1, _, _ := procVerSetConditionMask.Call(
		uintptr(conditionMask),
		uintptr(typeBitMask),
		uintptr(condition))
	return uint64(r1)
}
