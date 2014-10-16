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

func BstrToString(bstr *uint16) string {
	if bstr == nil {
		return ""
	}
	len := int(wrappers.SysStringLen(bstr))
	buf := make([]uint16, len)
	ptr := uintptr(unsafe.Pointer(bstr))
	for i := 0; i < len; i++ {
		buf[i] = *(*uint16)(unsafe.Pointer(ptr))
		ptr += 2
	}
	return syscall.UTF16ToString(buf)
}
