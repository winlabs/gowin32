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
	modoleaut32 = syscall.NewLazyDLL("oleaut32.dll")

	procSysAllocString = modoleaut32.NewProc("SysAllocString")
	procSysFreeString  = modoleaut32.NewProc("SysFreeString")
	procSysStringLen   = modoleaut32.NewProc("SysStringLen")
	procVariantClear   = modoleaut32.NewProc("VariantClear")
	procVariantInit    = modoleaut32.NewProc("VariantInit")
)

func SysAllocString(psz *uint16) *uint16 {
	r1, _, _ := procSysAllocString.Call(uintptr(unsafe.Pointer(psz)))
	return (*uint16)(unsafe.Pointer(r1))
}

func SysFreeString(bstrString *uint16) {
	procSysFreeString.Call(uintptr(unsafe.Pointer(bstrString)))
}

func SysStringLen(bstr *uint16) uint32 {
	r1, _, _ := procSysStringLen.Call(uintptr(unsafe.Pointer(bstr)))
	return uint32(r1)
}

func VariantClear(variant *Variant) error {
	r1, _, _ := procVariantClear.Call(uintptr(unsafe.Pointer(variant)))
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func VariantInit(variant *Variant) {
	procVariantInit.Call(uintptr(unsafe.Pointer(variant)))
}
