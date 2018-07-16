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
	modiphlpapi = syscall.NewLazyDLL("iphlpapi.dll")

	procGetTcpTable = modiphlpapi.NewProc("GetTcpTable")
	procSendARP     = modiphlpapi.NewProc("SendARP")
)

func GetTcpTable(tcpTable *MIB_TCPTABLE, size *uint32, order bool) error {
	r1, _, _ := syscall.Syscall(
		procGetTcpTable.Addr(),
		3,
		uintptr(unsafe.Pointer(tcpTable)),
		uintptr(unsafe.Pointer(size)),
		boolToUintptr(order))
	if err := syscall.Errno(r1); err != ERROR_SUCCESS {
		return err
	}
	return nil
}

func SendARP(destIP, srcIP uint32, macAddr, macAddrLen *uint32) error {
	r1, _, _ := syscall.Syscall6(
		procSendARP.Addr(),
		4,
		uintptr(destIP),
		uintptr(srcIP),
		uintptr(unsafe.Pointer(macAddr)),
		uintptr(unsafe.Pointer(macAddrLen)),
		0,
		0)
	if err := syscall.Errno(r1); err != ERROR_SUCCESS {
		return err
	}
	return nil
}
