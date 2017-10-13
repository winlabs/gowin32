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

import (
	"syscall"
	"unsafe"
)

var (
	procSetEntriesInAclW = modadvapi32.NewProc("SetEntriesInAclW")
)

func SetEntriesInAcl(countOfExplicitEntries uint32, listOfExplicitEntries *EXPLICIT_ACCESS, oldAcl *ACL, newAcl **ACL) error {
	r1, _, _ := syscall.Syscall6(
		procSetEntriesInAclW.Addr(),
		4,
		uintptr(countOfExplicitEntries),
		uintptr(unsafe.Pointer(listOfExplicitEntries)),
		uintptr(unsafe.Pointer(oldAcl)),
		uintptr(unsafe.Pointer(newAcl)),
		0,
		0)
	if err := syscall.Errno(r1); err != ERROR_SUCCESS {
		return err
	}
	return nil
}
