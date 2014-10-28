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
	"syscall"
	"unsafe"
)

const (
	PROCESS_NAME_NATIVE = 0x00000001
)

const (
	ComputerNameNetBIOS                   = 0
	ComputerNameDnsHostname               = 1
	ComputerNameDnsDomain                 = 2
	ComputerNameDnsFullyQualified         = 3
	ComputerNamePhysicalNetBIOS           = 4
	ComputerNamePhysicalDnsHostname       = 5
	ComputerNamePhysicalDnsDomain         = 6
	ComputerNamePhysicalDnsFullyQualified = 7
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	modadvapi32 = syscall.NewLazyDLL("advapi32.dll")

	procGetComputerNameExW         = modkernel32.NewProc("GetComputerNameExW")
	procGetDiskFreeSpaceExW        = modkernel32.NewProc("GetDiskFreeSpaceExW")
	procGetModuleFileNameW         = modkernel32.NewProc("GetModuleFileNameW")
	procGetSystemDirectoryW        = modkernel32.NewProc("GetSystemDirectoryW")
	procQueryFullProcessImageNameW = modkernel32.NewProc("QueryFullProcessImageNameW")
	procSetStdHandle               = modkernel32.NewProc("SetStdHandle")
	procVerifyVersionInfoW         = modkernel32.NewProc("VerifyVersionInfoW")
	proclstrlenW                   = modkernel32.NewProc("lstrlenW")

	procAllocateAndInitializeSid   = modadvapi32.NewProc("AllocateAndInitializeSid")
	procCheckTokenMembership       = modadvapi32.NewProc("CheckTokenMembership")
	procFreeSid                    = modadvapi32.NewProc("FreeSid")
	procGetFileSecurityW           = modadvapi32.NewProc("GetFileSecurityW")
	procGetSecurityDescriptorOwner = modadvapi32.NewProc("GetSecurityDescriptorOwner")
)

func GetComputerNameEx(nameType uint32, buffer *uint16, size *uint32) error {
	r1, _, e1 := procGetComputerNameExW.Call(
		uintptr(nameType),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(size)))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func GetDiskFreeSpaceEx(directoryName *uint16, freeBytesAvailable *uint64, totalNumberOfBytes *uint64, totalNumberOfFreeBytes *uint64) error {
	r1, _, e1 := procGetDiskFreeSpaceExW.Call(
		uintptr(unsafe.Pointer(directoryName)),
		uintptr(unsafe.Pointer(freeBytesAvailable)),
		uintptr(unsafe.Pointer(totalNumberOfBytes)),
		uintptr(unsafe.Pointer(totalNumberOfFreeBytes)))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func GetModuleFileName(module syscall.Handle, filename *uint16, size uint32) (uint32, error) {
	r1, _, e1 := procGetModuleFileNameW.Call(
		uintptr(module),
		uintptr(unsafe.Pointer(filename)),
		uintptr(size))
	if r1 == 0 || r1 == uintptr(size) {
		if e1.(syscall.Errno) != 0 {
			return uint32(r1), e1
		} else if r1 == uintptr(size) {
			return uint32(r1), syscall.ERROR_INSUFFICIENT_BUFFER
		} else {
			return uint32(r1), syscall.EINVAL
		}
	}
	return uint32(r1), nil
}

func GetSystemDirectory(buffer *uint16, size uint32) (uint32, error) {
	r1, _, e1 := procGetSystemDirectoryW.Call(
		uintptr(unsafe.Pointer(buffer)),
		uintptr(size))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return uint32(r1), e1
		} else {
			return uint32(r1), syscall.EINVAL
		}
	}
	return uint32(r1), nil
}

func QueryFullProcessImageName(process syscall.Handle, flags uint32, exeName *uint16, size *uint32) error {
	r1, _, e1 := procQueryFullProcessImageNameW.Call(
		uintptr(process),
		uintptr(flags),
		uintptr(unsafe.Pointer(exeName)),
		uintptr(unsafe.Pointer(size)))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func SetStdHandle(stdHandle int, handle syscall.Handle) error {
	r1, _, e1 := procSetStdHandle.Call(uintptr(stdHandle), uintptr(handle))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func VerifyVersionInfo(versionInfo *OSVERSIONINFOEX, typeMask uint32, conditionMask uint64) error {
	r1, _, e1 := procVerifyVersionInfoW.Call(
		uintptr(unsafe.Pointer(versionInfo)),
		uintptr(typeMask),
		uintptr(conditionMask))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func Lstrlen(string *uint16) int32 {
	r1, _, _ := proclstrlenW.Call(uintptr(unsafe.Pointer(string)))
	return int32(r1)
}

func AllocateAndInitializeSid(identifierAuthority *SIDIdentifierAuthority, subAuthorityCount byte, subAuthority0 uint32, subAuthority1 uint32, subAuthority2 uint32, subAuthority3 uint32, subAuthority4 uint32, subAuthority5 uint32, subAuthority6 uint32, subAuthority7 uint32, sid **syscall.SID) error {
	r1, _, e1 := procAllocateAndInitializeSid.Call(
		uintptr(unsafe.Pointer(identifierAuthority)),
		uintptr(subAuthorityCount),
		uintptr(subAuthority0),
		uintptr(subAuthority1),
		uintptr(subAuthority2),
		uintptr(subAuthority3),
		uintptr(subAuthority4),
		uintptr(subAuthority5),
		uintptr(subAuthority6),
		uintptr(subAuthority7),
		uintptr(unsafe.Pointer(sid)))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func CheckTokenMembership(tokenHandle syscall.Handle, sidToCheck *syscall.SID, isMember *bool) error {
	var isMemberRaw int32
	r1, _, e1 := procCheckTokenMembership.Call(
		uintptr(tokenHandle),
		uintptr(unsafe.Pointer(sidToCheck)),
		uintptr(unsafe.Pointer(&isMemberRaw)))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	if isMember != nil {
		*isMember = (isMemberRaw != 0)
	}
	return nil
}

func FreeSid(sid *syscall.SID) {
	procFreeSid.Call(uintptr(unsafe.Pointer(sid)))
}

func GetFileSecurity(fileName *uint16, requestedInformation uint32, securityDescriptor *uint8, length uint32, lengthNeeded *uint32) error {
	r1, _, e1 := procGetFileSecurityW.Call(
		uintptr(unsafe.Pointer(fileName)),
		uintptr(requestedInformation),
		uintptr(unsafe.Pointer(securityDescriptor)),
		uintptr(length),
		uintptr(unsafe.Pointer(lengthNeeded)))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func GetSecurityDescriptorOwner(securityDescriptor *uint8, owner **syscall.SID, ownerDefaulted *bool) error {
	var ownerDefaultedRaw int32
	r1, _, e1 := procGetSecurityDescriptorOwner.Call(
		uintptr(unsafe.Pointer(securityDescriptor)),
		uintptr(unsafe.Pointer(owner)),
		uintptr(unsafe.Pointer(&ownerDefaultedRaw)))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	if ownerDefaulted != nil {
		*ownerDefaulted = (ownerDefaultedRaw != 0)
	}
	return nil
}
