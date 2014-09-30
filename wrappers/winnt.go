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

var (
	procVerSetConditionMask = modkernel32.NewProc("VerSetConditionMask")
)

func VerSetConditionMask(conditionMask uint64, typeBitMask uint32, condition uint8) uint64 {
	r1, _, _ := procVerSetConditionMask.Call(
		uintptr(conditionMask),
		uintptr(typeBitMask),
		uintptr(condition))
	return uint64(r1)
}
