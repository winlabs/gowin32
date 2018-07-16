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

type WTSSESSION_NOTIFICATION struct {
	Size      uint32
	SessionId uint32
}

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

const (
	SM_CXSCREEN                    = 0
	SM_CYSCREEN                    = 1
	SM_CXVSCROLL                   = 2
	SM_CYHSCROLL                   = 3
	SM_CYCAPTION                   = 4
	SM_CXBORDER                    = 5
	SM_CYBORDER                    = 6
	SM_CXDLGFRAME                  = 7
	SM_CYDLGFRAME                  = 8
	SM_CYVTHUMB                    = 9
	SM_CXHTHUMB                    = 10
	SM_CXICON                      = 11
	SM_CYICON                      = 12
	SM_CXCURSOR                    = 13
	SM_CYCURSOR                    = 14
	SM_CYMENU                      = 15
	SM_CXFULLSCREEN                = 16
	SM_CYFULLSCREEN                = 17
	SM_CYKANJIWINDOW               = 18
	SM_MOUSEPRESENT                = 19
	SM_CYVSCROLL                   = 20
	SM_CXHSCROLL                   = 21
	SM_DEBUG                       = 22
	SM_SWAPBUTTON                  = 23
	SM_CXMIN                       = 28
	SM_CYMIN                       = 29
	SM_CXSIZE                      = 30
	SM_CYSIZE                      = 31
	SM_CXFRAME                     = 32
	SM_CYFRAME                     = 33
	SM_CXMINTRACK                  = 34
	SM_CYMINTRACK                  = 35
	SM_CXDOUBLECLK                 = 36
	SM_CYDOUBLECLK                 = 37
	SM_CXICONSPACING               = 38
	SM_CYICONSPACING               = 39
	SM_MENUDROPALIGNMENT           = 40
	SM_PENWINDOWS                  = 41
	SM_DBCSENABLED                 = 42
	SM_CMOUSEBUTTONS               = 43
	SM_CXFIXEDFRAME                = SM_CXDLGFRAME
	SM_CYFIXEDFRAME                = SM_CYDLGFRAME
	SM_CXSIZEFRAME                 = SM_CXFRAME
	SM_CYSIZEFRAME                 = SM_CYFRAME
	SM_SECURE                      = 44
	SM_CXEDGE                      = 45
	SM_CYEDGE                      = 46
	SM_CXMINSPACING                = 47
	SM_CYMINSPACING                = 48
	SM_CXSMICON                    = 49
	SM_CYSMICON                    = 50
	SM_CYSMCAPTION                 = 51
	SM_CXSMSIZE                    = 52
	SM_CYSMSIZE                    = 53
	SM_CXMENUSIZE                  = 54
	SM_CYMENUSIZE                  = 55
	SM_ARRANGE                     = 56
	SM_CXMINIMIZED                 = 57
	SM_CYMINIMIZED                 = 58
	SM_CXMAXTRACK                  = 59
	SM_CYMAXTRACK                  = 60
	SM_CXMAXIMIZED                 = 61
	SM_CYMAXIMIZED                 = 62
	SM_NETWORK                     = 63
	SM_CLEANBOOT                   = 67
	SM_CXDRAG                      = 68
	SM_CYDRAG                      = 69
	SM_SHOWSOUNDS                  = 70
	SM_CXMENUCHECK                 = 71
	SM_CYMENUCHECK                 = 72
	SM_SLOWMACHINE                 = 73
	SM_MIDEASTENABLED              = 74
	SM_MOUSEWHEELPRESENT           = 75
	SM_XVIRTUALSCREEN              = 76
	SM_YVIRTUALSCREEN              = 77
	SM_CXVIRTUALSCREEN             = 78
	SM_CYVIRTUALSCREEN             = 79
	SM_CMONITORS                   = 80
	SM_SAMEDISPLAYFORMAT           = 81
	SM_IMMENABLED                  = 82
	SM_CXFOCUSBORDER               = 83
	SM_CYFOCUSBORDER               = 84
	SM_TABLETPC                    = 86
	SM_MEDIACENTER                 = 87
	SM_STARTER                     = 88
	SM_SERVERR2                    = 89
	SM_MOUSEHORIZONTALWHEELPRESENT = 91
	SM_CXPADDEDBORDER              = 92
	SM_DIGITIZER                   = 94
	SM_MAXIMUMTOUCHES              = 95
	SM_REMOTESESSION               = 0x1000
	SM_SHUTTINGDOWN                = 0x2000
	SM_REMOTECONTROL               = 0x2001
)

const (
	EDD_GET_DEVICE_INTERFACE_NAME = 0x00000001
)

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	procCloseDesktop       = moduser32.NewProc("CloseDesktop")
	procEnumDisplayDevices = moduser32.NewProc("EnumDisplayDevicesW")
	procGetSystemMetrics   = moduser32.NewProc("GetSystemMetrics")
	procOpenInputDesktop   = moduser32.NewProc("OpenInputDesktop")
	procSetThreadDesktop   = moduser32.NewProc("SetThreadDesktop")
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

func EnumDisplayDevices(device *uint16, devNum uint32, displayDevice *DISPLAY_DEVICE, flags uint32) bool {
	r1, _, _ := syscall.Syscall6(
		procEnumDisplayDevices.Addr(),
		4,
		uintptr(unsafe.Pointer(device)),
		uintptr(devNum),
		uintptr(unsafe.Pointer(displayDevice)),
		uintptr(flags),
		0,
		0)
	return r1 != 0
}

func GetSystemMetrics(index int32) int {
	ret, _, _ := syscall.Syscall(procGetSystemMetrics.Addr(), 1, uintptr(index), 0, 0)
	return int(ret)
}

func OpenInputDesktop(flags uint32, inherit bool, desiredAccess uint32) (syscall.Handle, error) {
	r1, _, e1 := syscall.Syscall(
		procOpenInputDesktop.Addr(),
		3,
		uintptr(flags),
		boolToUintptr(inherit),
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
