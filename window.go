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

	"github.com/winlabs/gowin32/wrappers"
)

func GetWindowProcessId(hwnd syscall.Handle) syscall.Handle {
	var pid uint32
	wrappers.GetWindowThreadProcessId(hwnd, &pid)
	return syscall.Handle(pid)
}

func GetWindowText(hwnd syscall.Handle) (string, error) {
	buf := make([]uint16, 512)
	if _, err := wrappers.GetWindowText(hwnd, &buf[0], len(buf)); err != nil {
		return "", err
	}
	return syscall.UTF16ToString(buf), nil
}

func GetWindowThreadId(hwnd syscall.Handle) syscall.Handle {
	var pid uint32
	return syscall.Handle(wrappers.GetWindowThreadProcessId(hwnd, &pid))
}
