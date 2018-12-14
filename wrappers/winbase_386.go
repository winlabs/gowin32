/*
 * Copyright (c) 2014-2017 MongoDB, Inc.
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

// +build 386

package wrappers

import (
	"syscall"
	"unsafe"
)

func splitUint64To32(value uint64) (uint32, uint32) {
	return uint32(value & 0xffffffff), uint32(value >> 32)
}

func joinUint32To64(lo, hi uint32) uint64 {
	return uint64(lo) | (uint64(hi) << 32)
}

func VerifyVersionInfo(versionInfo *OSVERSIONINFOEX, typeMask uint32, conditionMask uint64) error {
	conditionMaskLo, conditionMaskHi := splitUint64To32(conditionMask)
	r1, _, e1 := syscall.Syscall6(
		procVerifyVersionInfoW.Addr(),
		6,
		uintptr(unsafe.Pointer(versionInfo)),
		uintptr(typeMask),
		uintptr(conditionMaskLo),
		uintptr(conditionMaskHi),
		0,
		0)
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}
