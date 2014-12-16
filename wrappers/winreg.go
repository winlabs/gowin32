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
	HKEY_CLASSES_ROOT                = 0x80000000
	HKEY_CURRENT_USER                = 0x80000001
	HKEY_LOCAL_MACHINE               = 0x80000002
	HKEY_USERS                       = 0x80000003
	HKEY_PERFORMANCE_DATA            = 0x80000004
	HKEY_PERFORMANCE_TEXT            = 0x80000050
	HKEY_PERFORMANCE_NLSTEXT         = 0x80000060
	HKEY_CURRENT_CONFIG              = 0x80000005
	HKEY_DYN_DATA                    = 0x80000006
	HKEY_CURRENT_USER_LOCAL_SETTINGS = 0x80000007
)

var (
	procRegCloseKey      = modadvapi32.NewProc("RegCloseKey")
	procRegCreateKeyExW  = modadvapi32.NewProc("RegCreateKeyExW")
	procRegDeleteValueW  = modadvapi32.NewProc("RegDeleteValueW")
	procRegOpenKeyExW    = modadvapi32.NewProc("RegOpenKeyExW")
	procRegQueryValueExW = modadvapi32.NewProc("RegQueryValueExW")
	procRegSetValueExW   = modadvapi32.NewProc("RegSetValueExW")
)

func RegCloseKey(key syscall.Handle) error {
	r1, _, _ := procRegCloseKey.Call(uintptr(key))
	if err := syscall.Errno(r1); err != ERROR_SUCCESS {
		return err
	}
	return nil
}

func RegCreateKeyEx(key syscall.Handle, subKey *uint16, reserved uint32, class *uint16, options uint32, samDesired uint32, securityAttributes *syscall.SecurityAttributes, result *syscall.Handle, disposition *uint32) error {
	r1, _, _ := procRegCreateKeyExW.Call(
		uintptr(key),
		uintptr(unsafe.Pointer(subKey)),
		uintptr(reserved),
		uintptr(unsafe.Pointer(class)),
		uintptr(options),
		uintptr(samDesired),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(unsafe.Pointer(result)),
		uintptr(unsafe.Pointer(disposition)))
	if err := syscall.Errno(r1); err != ERROR_SUCCESS {
		return err
	}
	return nil
}

func RegDeleteValue(key syscall.Handle, valueName *uint16) error {
	r1, _, _ := procRegDeleteValueW.Call(
		uintptr(key),
		uintptr(unsafe.Pointer(valueName)))
	if err := syscall.Errno(r1); err != ERROR_SUCCESS {
		return err
	}
	return nil
}

func RegOpenKeyEx(key syscall.Handle, subKey *uint16, options uint32, samDesired uint32, result *syscall.Handle) error {
	r1, _, _ := procRegOpenKeyExW.Call(
		uintptr(key),
		uintptr(unsafe.Pointer(subKey)),
		uintptr(options),
		uintptr(samDesired),
		uintptr(unsafe.Pointer(result)))
	if err := syscall.Errno(r1); err != ERROR_SUCCESS {
		return err
	}
	return nil
}

func RegQueryValueEx(key syscall.Handle, valueName *uint16, reserved *uint32, valueType *uint32, data *byte, cbData *uint32) error {
	r1, _, _ := procRegQueryValueExW.Call(
		uintptr(key),
		uintptr(unsafe.Pointer(valueName)),
		uintptr(unsafe.Pointer(reserved)),
		uintptr(unsafe.Pointer(valueType)),
		uintptr(unsafe.Pointer(data)),
		uintptr(unsafe.Pointer(cbData)))
	if err := syscall.Errno(r1); err != ERROR_SUCCESS {
		return err
	}
	return nil
}

func RegSetValueEx(key syscall.Handle, valueName *uint16, reserved uint32, valueType uint32, data *byte, cbData uint32) error {
	r1, _, _ := procRegSetValueExW.Call(
		uintptr(key),
		uintptr(unsafe.Pointer(valueName)),
		uintptr(reserved),
		uintptr(valueType),
		uintptr(unsafe.Pointer(data)),
		uintptr(cbData))
	if err := syscall.Errno(r1); err != ERROR_SUCCESS {
		return err
	}
	return nil
}
