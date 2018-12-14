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

import "syscall"

func VerSetConditionMask(conditionMask uint64, typeBitMask uint32, condition uint8) uint64 {
	conditionMaskLo, conditionMaskHi := splitUint64To32(conditionMask)
	r1, r2, _ := syscall.Syscall6(
		procVerSetConditionMask.Addr(),
		4,
		uintptr(conditionMaskLo),
		uintptr(conditionMaskHi),
		uintptr(typeBitMask),
		uintptr(condition),
		0,
		0)
	return joinUint32To64(uint32(r1), uint32(r2))
}
