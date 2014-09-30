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

type ProcessInfo struct {
	ProcessID       uint32
	Threads         uint32
	ParentProcessID uint32
	BasePriority    int32
	ExeFile         string
}

func GetProcesses() ([]ProcessInfo, error) {
	hSnapshot, err := wrappers.CreateToolhelp32Snapshot(wrappers.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(hSnapshot)
	pe := wrappers.ProcessEntry32{}
	pe.Size = uint32(unsafe.Sizeof(pe))
	if err := wrappers.Process32First(hSnapshot, &pe); err != nil {
		return nil, err
	}
	pi := []ProcessInfo{}
	for {
		pi = append(pi, ProcessInfo{
			ProcessID:       pe.ProcessID,
			Threads:         pe.Threads,
			ParentProcessID: pe.ParentProcessID,
			BasePriority:    pe.PriClassBase,
			ExeFile:         syscall.UTF16ToString((&pe.ExeFile)[:]),
		})
		err := wrappers.Process32Next(hSnapshot, &pe)
		if err == syscall.ERROR_NO_MORE_FILES {
			return pi, nil
		} else if err != nil {
			return nil, err
		}
	}
}
