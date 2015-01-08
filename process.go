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
	ProcessID       uint
	Threads         uint
	ParentProcessID uint
	BasePriority    int
	ExeFile         string
}

type ModuleInfo struct {
	ProcessID         uint
	ModuleBaseAddress *byte
	ModuleBaseSize    uint
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
		return nil, NewWindowsError("CreateToolhelp32Snapshot", err)
	}
	defer wrappers.CloseHandle(hSnapshot)
	pe := wrappers.PROCESSENTRY32{}
	pe.Size = uint32(unsafe.Sizeof(pe))
	if err := wrappers.Process32First(hSnapshot, &pe); err != nil {
		return nil, NewWindowsError("Process32First", err)
	}
	pi := []ProcessInfo{}
	for {
		pi = append(pi, ProcessInfo{
			ProcessID:       uint(pe.ProcessID),
			Threads:         uint(pe.Threads),
			ParentProcessID: uint(pe.ParentProcessID),
			BasePriority:    int(pe.PriClassBase),
			ExeFile:         syscall.UTF16ToString((&pe.ExeFile)[:]),
		})
		err := wrappers.Process32Next(hSnapshot, &pe)
		if err == wrappers.ERROR_NO_MORE_FILES {
			return pi, nil
		} else if err != nil {
			return nil, NewWindowsError("Process32Next", err)
		}
	}
}

func GetProcessModules(pid uint32) ([]ModuleInfo, error) {
	hSnapshot, err := wrappers.CreateToolhelp32Snapshot(wrappers.TH32CS_SNAPMODULE, pid)
	if err != nil {
		return nil, NewWindowsError("CreateToolhelp32Snapshot", err)
	}
	defer wrappers.CloseHandle(hSnapshot)
	me := wrappers.MODULEENTRY32{}
	me.Size = uint32(unsafe.Sizeof(me))
	if err := wrappers.Module32First(hSnapshot, &me); err != nil {
		return nil, NewWindowsError("Module32First", err)
	}
	mi := []ModuleInfo{}
	for {
		mi = append(mi, ModuleInfo{
			ProcessID:         uint(me.ProcessID),
			ModuleBaseAddress: me.ModBaseAddr,
			ModuleBaseSize:    uint(me.ModBaseSize),
			ModuleHandle:      me.Module,
			ModuleName:        syscall.UTF16ToString((&me.ModuleName)[:]),
			ExePath:           syscall.UTF16ToString((&me.ExePath)[:]),
		})
		err := wrappers.Module32Next(hSnapshot, &me)
		if err == wrappers.ERROR_NO_MORE_FILES {
			return mi, nil
		} else if err != nil {
			return nil, NewWindowsError("Module32Next", err)
		}
	}
}

func SignalProcessAndWait(pid uint, timeout time.Duration) error {
	milliseconds := uint32(timeout / time.Millisecond)
	if timeout < 0 {
		milliseconds = wrappers.INFINITE
	}
	hProcess, err := wrappers.OpenProcess(wrappers.SYNCHRONIZE, false, uint32(pid))
	if err != nil {
		return NewWindowsError("OpenProcess", err)
	}
	defer wrappers.CloseHandle(hProcess)
	if err := wrappers.GenerateConsoleCtrlEvent(wrappers.CTRL_BREAK_EVENT, uint32(pid)); err != nil {
		return NewWindowsError("GenerateConsoleCtrlEvent", err)
	}
	if _, err := wrappers.WaitForSingleObject(hProcess, milliseconds); err != nil {
		return NewWindowsError("WaitForSingleObject", err)
	}
	return nil
}

func KillProcess(pid uint, exitCode uint) error {
	hProcess, err := wrappers.OpenProcess(wrappers.PROCESS_TERMINATE, false, uint32(pid))
	if err != nil {
		return NewWindowsError("OpenProcess", err)
	}
	defer wrappers.CloseHandle(hProcess)
	if err := wrappers.TerminateProcess(hProcess, uint32(exitCode)); err != nil {
		return NewWindowsError("TerminateProcess", err)
	}
	return nil
}

func IsProcessRunning(pid uint) (bool, error) {
	hProcess, err := wrappers.OpenProcess(wrappers.SYNCHRONIZE, false, uint32(pid))
	if err == wrappers.ERROR_INVALID_PARAMETER {
		// the process no longer exists
		return false, nil
	} else if err != nil {
		return false, NewWindowsError("OpenProcess", err)
	}
	defer wrappers.CloseHandle(hProcess)

	// wait with a timeout of 0 to check the process's status and make sure it's not a zombie
	event, err := wrappers.WaitForSingleObject(hProcess, 0)
	if err != nil {
		return false, NewWindowsError("WaitForSingleObject", err)
	}
	return event != wrappers.WAIT_OBJECT_0, nil
}

func GetProcessFullPathName(pid uint, flags ProcessNameFlags) (string, error) {
	hProcess, err := wrappers.OpenProcess(wrappers.PROCESS_QUERY_LIMITED_INFORMATION, false, uint32(pid))
	if err != nil {
		return "", NewWindowsError("OpenProcess", err)
	}
	defer wrappers.CloseHandle(hProcess)

	buf := make([]uint16, wrappers.MAX_PATH)
	size := uint32(wrappers.MAX_PATH)
	if err := wrappers.QueryFullProcessImageName(hProcess, uint32(flags), &buf[0], &size); err != nil {
		if err == wrappers.ERROR_INSUFFICIENT_BUFFER {
			buf = make([]uint16, syscall.MAX_LONG_PATH)
			size = syscall.MAX_LONG_PATH
			if err := wrappers.QueryFullProcessImageName(hProcess, uint32(flags), &buf[0], &size); err != nil {
				return "", NewWindowsError("QueryFullProcessImageName", err)
			}
		} else {
			return "", NewWindowsError("QueryFullProcessImageName", err)
		}
	}
	return syscall.UTF16ToString(buf[0:size]), nil
}
