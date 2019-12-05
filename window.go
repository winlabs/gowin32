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
	"syscall"
	"unsafe"

	"github.com/winlabs/gowin32/wrappers"
)

func EnumDesktops(winsta syscall.Handle) ([]string, error) {

	result := make([]string, 0)
	if err := wrappers.EnumDesktops(winsta,
		func(name *uint16, lparam uintptr) int32 {
			result = append(result, syscall.UTF16ToString((*[1024]uint16)(unsafe.Pointer(name))[:]))
			return 1
		},
		0); err != nil {
		return nil, NewWindowsError("EnumDesktops", err)
	}
	return result, nil
}

func GetWindowProcessID(hwnd syscall.Handle) (uint, error) {
	var pid uint32
	_, err := wrappers.GetWindowThreadProcessId(hwnd, &pid)
	if err != nil {
		return 0, NewWindowsError("GetWindowThreadProcessId", err)
	}
	return uint(pid), nil
}

func GetWindowText(hwnd syscall.Handle) (string, error) {
	l, err := wrappers.GetWindowTextLength(hwnd)
	if err != nil {
		return "", NewWindowsError("GetWindowTextLength", err)
	}
	if l == 0 {
		return "", nil
	}
	buf := make([]uint16, l+1)
	if _, err := wrappers.GetWindowText(hwnd, &buf[0], l+1); err != nil {
		return "", NewWindowsError("GetWindowText", err)
	}
	return syscall.UTF16ToString(buf), nil
}

func GetWindowThreadID(hwnd syscall.Handle) (uint, error) {
	var pid uint32
	r, err := wrappers.GetWindowThreadProcessId(hwnd, &pid)
	if err != nil {
		return 0, NewWindowsError("GetWindowThreadProcessId", err)
	}
	return uint(r), nil
}
