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

package gowin32

import (
	"github.com/winlabs/gowin32/wrappers"

	"unsafe"
)

type MemoryStatus struct {
	MemoryLoad                    uint
	TotalPhysicalBytes            uint64
	AvailablePhysicalBytes        uint64
	TotalPageFileBytes            uint64
	AvailablePageFileBytes        uint64
	TotalVirtualBytes             uint64
	AvailableVirtualBytes         uint64
	AvailableExtendedVirtualBytes uint64
}

func GetMemoryStatus() (*MemoryStatus, error) {
	var status wrappers.MEMORYSTATUSEX
	status.Length = uint32(unsafe.Sizeof(status))
	if err := wrappers.GlobalMemoryStatusEx(&status); err != nil {
		return nil, NewWindowsError("GlobalMemoryStatusEx", err)
	}
	return &MemoryStatus{
		MemoryLoad:                    uint(status.MemoryLoad),
		TotalPhysicalBytes:            status.TotalPhys,
		AvailablePhysicalBytes:        status.AvailPhys,
		TotalPageFileBytes:            status.TotalPageFile,
		AvailablePageFileBytes:        status.AvailPageFile,
		TotalVirtualBytes:             status.TotalVirtual,
		AvailableVirtualBytes:         status.AvailVirtual,
		AvailableExtendedVirtualBytes: status.AvailExtendedVirtual,
	}, nil
}
