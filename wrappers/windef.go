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

package wrappers

const (
	MAX_PATH = 260
)

type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

func MAKELONG(low uint16, high uint16) uint32 {
	return uint32(low) | (uint32(high) << 16)
}

func LOWORD(value uint32) uint16 {
	return uint16(value & 0x0000FFFF)
}

func HIWORD(value uint32) uint16 {
	return uint16((value >> 16) & 0x0000FFFF)
}

func boolToUintptr(value bool) uintptr {
	var valueRaw int32
	if value {
		valueRaw = 1
	} else {
		valueRaw = 0
	}

	return uintptr(valueRaw)
}

func makeUint64(lo, hi uint32) uint64 {
	return uint64(lo) | (uint64(hi) << 32)
}

func loUint32(value uint64) uint32 {
	return uint32(value & 0xffffffff)
}

func hiUint32(value uint64) uint32 {
	return uint32(value >> 32)
}
