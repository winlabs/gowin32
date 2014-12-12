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
	CSIDL_DESKTOP                 = 0x0000
	CSIDL_INTERNET                = 0x0001
	CSIDL_PROGRAMS                = 0x0002
	CSIDL_CONTROLS                = 0x0003
	CSIDL_PRINTERS                = 0x0004
	CSIDL_PERSONAL                = 0x0005
	CSIDL_FAVORITES               = 0x0006
	CSIDL_STARTUP                 = 0x0007
	CSIDL_RECENT                  = 0x0008
	CSIDL_SENDTO                  = 0x0009
	CSIDL_BITBUCKET               = 0x000a
	CSIDL_STARTMENU               = 0x000b
	CSIDL_MYDOCUMENTS             = CSIDL_PERSONAL
	CSIDL_MYMUSIC                 = 0x000d
	CSIDL_MYVIDEO                 = 0x000e
	CSIDL_DESKTOPDIRECTORY        = 0x0010
	CSIDL_DRIVES                  = 0x0011
	CSIDL_NETWORK                 = 0x0012
	CSIDL_NETHOOD                 = 0x0013
	CSIDL_FONTS                   = 0x0014
	CSIDL_TEMPLATES               = 0x0015
	CSIDL_COMMON_STARTMENU        = 0x0016
	CSIDL_COMMON_PROGRAMS         = 0x0017
	CSIDL_COMMON_STARTUP          = 0x0018
	CSIDL_COMMON_DESKTOPDIRECTORY = 0x0019
	CSIDL_APPDATA                 = 0x001a
	CSIDL_PRINTHOOD               = 0x001b
	CSIDL_LOCAL_APPDATA           = 0x001c
	CSIDL_ALTSTARTUP              = 0x001d
	CSIDL_COMMON_ALTSTARTUP       = 0x001e
	CSIDL_COMMON_FAVORITES        = 0x001f
	CSIDL_INTERNET_CACHE          = 0x0020
	CSIDL_COOKIES                 = 0x0021
	CSIDL_HISTORY                 = 0x0022
	CSIDL_COMMON_APPDATA          = 0x0023
	CSIDL_WINDOWS                 = 0x0024
	CSIDL_SYSTEM                  = 0x0025
	CSIDL_PROGRAM_FILES           = 0x0026
	CSIDL_MYPICTURES              = 0x0027
	CSIDL_PROFILE                 = 0x0028
	CSIDL_SYSTEMX86               = 0x0029
	CSIDL_PROGRAM_FILESX86        = 0x002a
	CSIDL_PROGRAM_FILES_COMMON    = 0x002b
	CSIDL_PROGRAM_FILES_COMMONX86 = 0x002c
	CSIDL_COMMON_TEMPLATES        = 0x002d
	CSIDL_COMMON_DOCUMENTS        = 0x002e
	CSIDL_COMMON_ADMINTOOLS       = 0x002f
	CSIDL_ADMINTOOLS              = 0x0030
	CSIDL_CONNECTIONS             = 0x0031
	CSIDL_COMMON_MUSIC            = 0x0035
	CSIDL_COMMON_PICTURES         = 0x0036
	CSIDL_COMMON_VIDEO            = 0x0037
	CSIDL_RESOURCES               = 0x0038
	CSIDL_RESOURCES_LOCALIZED     = 0x0039
	CSIDL_COMMON_OEM_LINKS        = 0x003a
	CSIDL_CDBURN_AREA             = 0x003b
	CSIDL_COMPUTERSNEARME         = 0x003d

	CSIDL_FLAG_CREATE        = 0x8000
	CSIDL_FLAG_DONT_VERIFY   = 0x4000
	CSIDL_FLAG_UNEXPAND      = 0x2000
	CSIDL_FLAG_NO_ALIAS      = 0x1000
	CSIDL_FLAG_PER_USER_INIT = 0x0800
)

var (
	modshell32 = syscall.NewLazyDLL("shell32.dll")

	procSHGetFolderPathW = modshell32.NewProc("SHGetFolderPathW")
)

func SHGetFolderPath(owner syscall.Handle, folder uint32, token syscall.Handle, flags uint32, path *uint16) uint32 {
	r1, _, _ := procSHGetFolderPathW.Call(
		uintptr(owner),
		uintptr(folder),
		uintptr(token),
		uintptr(flags),
		uintptr(unsafe.Pointer(path)))
	return uint32(r1)
}
