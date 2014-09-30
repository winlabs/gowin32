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

const (
	TH32CS_SNAPHEAPLIST = 0x00000001
	TH32CS_SNAPPROCESS  = 0x00000002
	TH32CS_SNAPTHREAD   = 0x00000004
	TH32CS_SNAPMODULE   = 0x00000008
	TH32CS_SNAPMODULE32 = 0x00000010
	TH32CS_SNAPALL      = TH32CS_SNAPHEAPLIST | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD | TH32CS_SNAPMODULE
	TH32CS_INHERIT      = 0x80000000
)

type ProcessEntry32 struct {
	Size            uint32
	Usage           uint32
	ProcessID       uint32
	DefaultHeapID   uintptr
	ModuleID        uint32
	Threads         uint32
	ParentProcessID uint32
	PriClassBase    int32
	Flags           uint32
	ExeFile         [syscall.MAX_PATH]uint16
}

var (
	procCreateToolhelp32Snapshot = modkernel32.NewProc("CreateToolhelp32Snapshot")
	procProcess32FirstW          = modkernel32.NewProc("Process32FirstW")
	procProcess32NextW           = modkernel32.NewProc("Process32NextW")
)

func CreateToolhelp32Snapshot(flags uint32, processID uint32) (syscall.Handle, error) {
	r1, _, e1 := procCreateToolhelp32Snapshot.Call(uintptr(flags), uintptr(processID))
	handle := syscall.Handle(r1)
	if handle == syscall.InvalidHandle {
		if e1.(syscall.Errno) != 0 {
			return syscall.InvalidHandle, e1
		} else {
			return syscall.InvalidHandle, syscall.EINVAL
		}
	}
	return handle, nil
}

func Process32First(snapshot syscall.Handle, pe *ProcessEntry32) error {
	r1, _, e1 := procProcess32FirstW.Call(uintptr(snapshot), uintptr(unsafe.Pointer(pe)))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func Process32Next(snapshot syscall.Handle, pe *ProcessEntry32) error {
	r1, _, e1 := procProcess32NextW.Call(uintptr(snapshot), uintptr(unsafe.Pointer(pe)))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}
