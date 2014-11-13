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

var (
	procRegCreateKeyExW = modadvapi32.NewProc("RegCreateKeyExW")
	procRegDeleteValueW = modadvapi32.NewProc("RegDeleteValueW")
	procRegSetValueExW  = modadvapi32.NewProc("RegSetValueExW")
)

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
	if r1 != 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func RegDeleteValue(key syscall.Handle, valueName *uint16) error {
	r1, _, _ := procRegDeleteValueW.Call(
		uintptr(key),
		uintptr(unsafe.Pointer(valueName)))
	if r1 != 0 {
		return syscall.Errno(r1)
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
	if r1 != 0 {
		return syscall.Errno(r1)
	}
	return nil
}
