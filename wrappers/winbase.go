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

	procGetComputerNameExW  = modkernel32.NewProc("GetComputerNameExW")
	procGetDiskFreeSpaceExW = modkernel32.NewProc("GetDiskFreeSpaceExW")
	procVerifyVersionInfoW  = modkernel32.NewProc("VerifyVersionInfoW")

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
	rawOwnerDefaulted := int32(0)
	r1, _, e1 := procGetSecurityDescriptorOwner.Call(
		uintptr(unsafe.Pointer(securityDescriptor)),
		uintptr(unsafe.Pointer(owner)),
		uintptr(unsafe.Pointer(&rawOwnerDefaulted)))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	if ownerDefaulted != nil {
		*ownerDefaulted = (rawOwnerDefaulted != 0)
	}
	return nil
}
