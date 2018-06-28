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
)

func MAKEINTRESOURCE(integer uint16) uintptr {
	return uintptr(integer)
}

var (
	RT_CURSOR       = MAKEINTRESOURCE(1)
	RT_BITMAP       = MAKEINTRESOURCE(2)
	RT_ICON         = MAKEINTRESOURCE(3)
	RT_MENU         = MAKEINTRESOURCE(4)
	RT_DIALOG       = MAKEINTRESOURCE(5)
	RT_STRING       = MAKEINTRESOURCE(6)
	RT_FONTDIR      = MAKEINTRESOURCE(7)
	RT_FONT         = MAKEINTRESOURCE(8)
	RT_ACCELERATOR  = MAKEINTRESOURCE(9)
	RT_RCDATA       = MAKEINTRESOURCE(10)
	RT_MESSAGETABLE = MAKEINTRESOURCE(11)
	RT_GROUP_CURSOR = MAKEINTRESOURCE(12)
	RT_GROUP_ICON   = MAKEINTRESOURCE(14)
	RT_VERSION      = MAKEINTRESOURCE(16)
	RT_DLGINCLUDE   = MAKEINTRESOURCE(17)
	RT_PLUGPLAY     = MAKEINTRESOURCE(19)
	RT_VXD          = MAKEINTRESOURCE(20)
	RT_ANICURSOR    = MAKEINTRESOURCE(21)
	RT_ANIICON      = MAKEINTRESOURCE(22)
	RT_HTML         = MAKEINTRESOURCE(23)
	RT_MANIFEST     = MAKEINTRESOURCE(24)
)

var (
	CREATEPROCESS_MANIFEST_RESOURCE_ID                 = MAKEINTRESOURCE(1)
	ISOLATIONAWARE_MANIFEST_RESOURCE_ID                = MAKEINTRESOURCE(2)
	ISOLATIONAWARE_NOSTATICIMPORT_MANIFEST_RESOURCE_ID = MAKEINTRESOURCE(3)
)

const (
	SW_HIDE            = 0
	SW_SHOWNORMAL      = 1
	SW_SHOWMINIMIZED   = 2
	SW_SHOWMAXIMIZED   = 3
	SW_MAXIMIZE        = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
)

const (
	WTS_CONSOLE_CONNECT        = 0x1
	WTS_CONSOLE_DISCONNECT     = 0x2
	WTS_REMOTE_CONNECT         = 0x3
	WTS_REMOTE_DISCONNECT      = 0x4
	WTS_SESSION_LOGON          = 0x5
	WTS_SESSION_LOGOFF         = 0x6
	WTS_SESSION_LOCK           = 0x7
	WTS_SESSION_UNLOCK         = 0x8
	WTS_SESSION_REMOTE_CONTROL = 0x9
	WTS_SESSION_CREATE         = 0xA
	WTS_SESSION_TERMINATE      = 0xB
)

type WTSSESSION_NOTIFICATION struct {
	Size      uint32
	SessionId uint32
}

const (
	DESKTOP_READOBJECTS     = 0x00000001
	DESKTOP_CREATEWINDOW    = 0x00000002
	DESKTOP_CREATEMENU      = 0x00000004
	DESKTOP_HOOKCONTROL     = 0x00000008
	DESKTOP_JOURNALRECORD   = 0x00000010
	DESKTOP_JOURNALPLAYBACK = 0x00000020
	DESKTOP_ENUMERATE       = 0x00000040
	DESKTOP_WRITEOBJECTS    = 0x00000080
	DESKTOP_SWITCHDESKTOP   = 0x00000100
)

const (
	DF_ALLOWOTHERACCOUNTHOOK = 0x00000001
)

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	procCloseDesktop     = moduser32.NewProc("CloseDesktop")
	procOpenInputDesktop = moduser32.NewProc("OpenInputDesktop")
	procSetThreadDesktop = moduser32.NewProc("SetThreadDesktop")
)

func CloseDesktop(desktop syscall.Handle) error {
	r1, _, e1 := syscall.Syscall(procCloseDesktop.Addr(), 1, uintptr(desktop), 0, 0)
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func OpenInputDesktop(flags uint32, inherit bool, desiredAccess uint32) (syscall.Handle, error) {
	var inheritRaw int32
	if inherit {
		inheritRaw = 1
	} else {
		inheritRaw = 0
	}
	r1, _, e1 := syscall.Syscall(
		procOpenInputDesktop.Addr(),
		3,
		uintptr(flags),
		uintptr(inheritRaw),
		uintptr(desiredAccess))

	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}
	return syscall.Handle(r1), nil
}

func SetThreadDesktop(desktop syscall.Handle) error {
	r1, _, e1 := syscall.Syscall(procSetThreadDesktop.Addr(), 1, uintptr(desktop), 0, 0)
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}
