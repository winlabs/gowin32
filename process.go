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
	"time"
	"unsafe"
)

type ProcessInfo struct {
	ProcessID       uint32
	Threads         uint32
	ParentProcessID uint32
	BasePriority    int32
	ExeFile         string
}

type ModuleInfo struct {
	ProcessID         uint32
	ModuleBaseAddress *uint8
	ModuleBaseSize    uint32
	ModuleHandle      syscall.Handle
	ModuleName        string
	ExePath           string
}

type ProcessNameFlags uint32

const (
	ProcessNameNative ProcessNameFlags = wrappers.PROCESS_NAME_NATIVE
)

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

func GetProcessModules(pid uint32) ([]ModuleInfo, error) {
	hSnapshot, err := wrappers.CreateToolhelp32Snapshot(wrappers.TH32CS_SNAPMODULE, pid)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(hSnapshot)
	me := wrappers.ModuleEntry32{}
	me.Size = uint32(unsafe.Sizeof(me))
	if err := wrappers.Module32First(hSnapshot, &me); err != nil {
		return nil, err
	}
	mi := []ModuleInfo{}
	for {
		mi = append(mi, ModuleInfo{
			ProcessID:         me.ProcessID,
			ModuleBaseAddress: me.ModBaseAddr,
			ModuleBaseSize:    me.ModBaseSize,
			ModuleHandle:      me.Module,
			ModuleName:        syscall.UTF16ToString((&me.ModuleName)[:]),
			ExePath:           syscall.UTF16ToString((&me.ExePath)[:]),
		})
		err := wrappers.Module32Next(hSnapshot, &me)
		if err == syscall.ERROR_NO_MORE_FILES {
			return mi, nil
		} else if err != nil {
			return nil, err
		}
	}
}

func SignalProcessAndWait(pid uint32, timeout time.Duration) error {
	milliseconds := uint32(timeout / time.Millisecond)
	if timeout < 0 {
		milliseconds = syscall.INFINITE
	}
	hProcess, err := syscall.OpenProcess(syscall.SYNCHRONIZE, false, pid)
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(hProcess)
	if err := wrappers.GenerateConsoleCtrlEvent(wrappers.CTRL_BREAK_EVENT, pid); err != nil {
		return err
	}
	if _, err := syscall.WaitForSingleObject(hProcess, milliseconds); err != nil {
		return err
	}
	return nil
}

func KillProcess(pid uint32, exitCode uint32) error {
	hProcess, err := syscall.OpenProcess(wrappers.PROCESS_TERMINATE, false, pid)
	if err != nil {
		return err
	}
	defer syscall.CloseHandle(hProcess)
	return syscall.TerminateProcess(hProcess, exitCode)
}

func IsProcessRunning(pid uint32) (bool, error) {
	hProcess, err := syscall.OpenProcess(syscall.SYNCHRONIZE, false, pid)
	if err == wrappers.ERROR_INVALID_PARAMETER {
		// the process no longer exists
		return false, nil
	} else if err != nil {
		return false, err
	}
	defer syscall.CloseHandle(hProcess)

	// wait with a timeout of 0 to check the process's status and make sure it's not a zombie
	event, err := syscall.WaitForSingleObject(hProcess, 0)
	if err != nil {
		return false, err
	}
	return event != syscall.WAIT_OBJECT_0, nil
}

func GetProcessFullPathName(pid uint32, flags ProcessNameFlags) (string, error) {
	hProcess, err := syscall.OpenProcess(wrappers.PROCESS_QUERY_LIMITED_INFORMATION, false, pid)
	if err != nil {
		return "", err
	}
	defer syscall.CloseHandle(hProcess)

	buf := make([]uint16, syscall.MAX_PATH)
	size := uint32(syscall.MAX_PATH)
	if err := wrappers.QueryFullProcessImageName(hProcess, uint32(flags), &buf[0], &size); err != nil {
		if err == syscall.ERROR_INSUFFICIENT_BUFFER {
			buf = make([]uint16, syscall.MAX_LONG_PATH)
			size = syscall.MAX_LONG_PATH
			if err := wrappers.QueryFullProcessImageName(hProcess, uint32(flags), &buf[0], &size); err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	return syscall.UTF16ToString(buf[0:size]), nil
}
