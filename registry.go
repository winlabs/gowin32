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

package gowin32

import (
	"github.com/winlabs/gowin32/wrappers"

	"syscall"
	"unsafe"
)

type RegRoot syscall.Handle

const (
	RegRootHKCR RegRoot = syscall.HKEY_CLASSES_ROOT
	RegRootHKCU RegRoot = syscall.HKEY_CURRENT_USER
	RegRootHKLM RegRoot = syscall.HKEY_LOCAL_MACHINE
	RegRootHKU  RegRoot = syscall.HKEY_USERS
	RegRootHKPD RegRoot = syscall.HKEY_PERFORMANCE_DATA
	RegRootHKCC RegRoot = syscall.HKEY_CURRENT_CONFIG
	RegRootHKDD RegRoot = syscall.HKEY_DYN_DATA
)

func DeleteRegValue(root RegRoot, subKey string, valueName string) error {
	var hKey syscall.Handle
	err := syscall.RegOpenKeyEx(
		syscall.Handle(root),
		syscall.StringToUTF16Ptr(subKey),
		0,
		syscall.KEY_WRITE,
		&hKey)
	if err != nil {
		return err
	}
	defer syscall.RegCloseKey(hKey)
	return wrappers.RegDeleteValue(hKey, syscall.StringToUTF16Ptr(valueName))
}

func GetRegValueString(root RegRoot, subKey string, valueName string) (string, error) {
	var hKey syscall.Handle
	err := syscall.RegOpenKeyEx(
		syscall.Handle(root),
		syscall.StringToUTF16Ptr(subKey),
		0,
		syscall.KEY_READ,
		&hKey)
	if err != nil {
		return "", err
	}
	defer syscall.RegCloseKey(hKey)
	var valueType uint32
	var size uint32
	err = syscall.RegQueryValueEx(
		hKey,
		syscall.StringToUTF16Ptr(valueName),
		nil,
		&valueType,
		nil,
		&size)
	if err != nil {
		return "", err
	}
	if valueType != syscall.REG_SZ {
		// use the same error code as RegGetValue, although that function is not used here in order to maintain
		// compatibility with older versions of Windows
		return "", wrappers.ERROR_UNSUPPORTED_TYPE
	}
	buf := make([]uint16, size/2)
	err = syscall.RegQueryValueEx(
		hKey,
		syscall.StringToUTF16Ptr(valueName),
		nil,
		nil,
		(*byte)(unsafe.Pointer(&buf[0])),
		&size)
	if err != nil {
		return "", err
	}
	return syscall.UTF16ToString(buf), nil
}

func SetRegValueString(root RegRoot, subKey string, valueName string, data string) error {
	var hKey syscall.Handle
	err := wrappers.RegCreateKeyEx(
		syscall.Handle(root),
		syscall.StringToUTF16Ptr(subKey),
		0,
		nil,
		0,
		syscall.KEY_WRITE,
		nil,
		&hKey,
		nil)
	if err != nil {
		return err
	}
	defer syscall.RegCloseKey(hKey)
	return wrappers.RegSetValueEx(
		hKey,
		syscall.StringToUTF16Ptr(valueName),
		0,
		syscall.REG_SZ,
		(*byte)(unsafe.Pointer(syscall.StringToUTF16Ptr(data))),
		uint32(2*(len(data)+1)))
}
