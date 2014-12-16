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
)

type SpecialFolder uint32

const (
	FolderDesktop                SpecialFolder = wrappers.CSIDL_DESKTOP
	FolderInternet               SpecialFolder = wrappers.CSIDL_INTERNET
	FolderPrograms               SpecialFolder = wrappers.CSIDL_PROGRAMS
	FolderControls               SpecialFolder = wrappers.CSIDL_CONTROLS
	FolderPrinters               SpecialFolder = wrappers.CSIDL_PRINTERS
	FolderPersonal               SpecialFolder = wrappers.CSIDL_PERSONAL
	FolderFavorites              SpecialFolder = wrappers.CSIDL_FAVORITES
	FolderStartup                SpecialFolder = wrappers.CSIDL_STARTUP
	FolderRecent                 SpecialFolder = wrappers.CSIDL_RECENT
	FolderSendTo                 SpecialFolder = wrappers.CSIDL_SENDTO
	FolderBitBucket              SpecialFolder = wrappers.CSIDL_BITBUCKET
	FolderStartMenu              SpecialFolder = wrappers.CSIDL_STARTMENU
	FolderMyDocuments            SpecialFolder = wrappers.CSIDL_MYDOCUMENTS
	FolderMyMusic                SpecialFolder = wrappers.CSIDL_MYMUSIC
	FolderMyVideo                SpecialFolder = wrappers.CSIDL_MYVIDEO
	FolderDesktopDirectory       SpecialFolder = wrappers.CSIDL_DESKTOPDIRECTORY
	FolderDrives                 SpecialFolder = wrappers.CSIDL_DRIVES
	FolderNetwork                SpecialFolder = wrappers.CSIDL_NETWORK
	FolderNetHood                SpecialFolder = wrappers.CSIDL_NETHOOD
	FolderFonts                  SpecialFolder = wrappers.CSIDL_FONTS
	FolderTemplates              SpecialFolder = wrappers.CSIDL_TEMPLATES
	FolderCommonStartMenu        SpecialFolder = wrappers.CSIDL_COMMON_STARTMENU
	FolderCommonPrograms         SpecialFolder = wrappers.CSIDL_COMMON_PROGRAMS
	FolderCommonStartup          SpecialFolder = wrappers.CSIDL_COMMON_STARTUP
	FolderCommonDesktopDirectory SpecialFolder = wrappers.CSIDL_COMMON_DESKTOPDIRECTORY
	FolderAppData                SpecialFolder = wrappers.CSIDL_APPDATA
	FolderPrintHood              SpecialFolder = wrappers.CSIDL_PRINTHOOD
	FolderLocalAppData           SpecialFolder = wrappers.CSIDL_LOCAL_APPDATA
	FolderAltStartup             SpecialFolder = wrappers.CSIDL_ALTSTARTUP
	FolderCommonAltStartup       SpecialFolder = wrappers.CSIDL_COMMON_ALTSTARTUP
	FolderCommonFavorites        SpecialFolder = wrappers.CSIDL_COMMON_FAVORITES
	FolderInternetCache          SpecialFolder = wrappers.CSIDL_INTERNET_CACHE
	FolderCookies                SpecialFolder = wrappers.CSIDL_COOKIES
	FolderHistory                SpecialFolder = wrappers.CSIDL_HISTORY
	FolderCommonAppData          SpecialFolder = wrappers.CSIDL_COMMON_APPDATA
	FolderWindows                SpecialFolder = wrappers.CSIDL_WINDOWS
	FolderSystem                 SpecialFolder = wrappers.CSIDL_SYSTEM
	FolderProgramFiles           SpecialFolder = wrappers.CSIDL_PROGRAM_FILES
	FolderMyPictures             SpecialFolder = wrappers.CSIDL_MYPICTURES
	FolderProfile                SpecialFolder = wrappers.CSIDL_PROFILE
	FolderSystemX86              SpecialFolder = wrappers.CSIDL_SYSTEMX86
	FolderProgramFilesX86        SpecialFolder = wrappers.CSIDL_PROGRAM_FILESX86
	FolderProgramFilesCommon     SpecialFolder = wrappers.CSIDL_PROGRAM_FILES_COMMON
	FolderProgramFilesCommonX86  SpecialFolder = wrappers.CSIDL_PROGRAM_FILES_COMMONX86
	FolderCommonTemplates        SpecialFolder = wrappers.CSIDL_COMMON_TEMPLATES
	FolderCommonDocuments        SpecialFolder = wrappers.CSIDL_COMMON_DOCUMENTS
	FolderCommonAdminTools       SpecialFolder = wrappers.CSIDL_COMMON_ADMINTOOLS
	FolderAdminTools             SpecialFolder = wrappers.CSIDL_ADMINTOOLS
	FolderConnections            SpecialFolder = wrappers.CSIDL_CONNECTIONS
	FolderCommonMusic            SpecialFolder = wrappers.CSIDL_COMMON_MUSIC
	FolderCommonPictures         SpecialFolder = wrappers.CSIDL_COMMON_PICTURES
	FolderCommonVideo            SpecialFolder = wrappers.CSIDL_COMMON_VIDEO
	FolderResources              SpecialFolder = wrappers.CSIDL_RESOURCES
	FolderResourcesLocalized     SpecialFolder = wrappers.CSIDL_RESOURCES_LOCALIZED
	FolderCommonOEMLinks         SpecialFolder = wrappers.CSIDL_COMMON_OEM_LINKS
	FolderCDBurnArea             SpecialFolder = wrappers.CSIDL_CDBURN_AREA
	FolderComputersNearMe        SpecialFolder = wrappers.CSIDL_COMPUTERSNEARME
)

func GetSpecialFolderPath(folder SpecialFolder) (string, error) {
	buf := [wrappers.MAX_PATH]uint16{}
	if hr := wrappers.SHGetFolderPath(0, uint32(folder), 0, 0, &buf[0]); wrappers.FAILED(hr) {
		return "", COMError(hr)
	}
	return syscall.UTF16ToString((&buf)[:]), nil
}
