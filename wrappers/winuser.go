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
	"unsafe"
)

func MakeIntResource(integer uint16) *uint16 {
	return (*uint16)(unsafe.Pointer(uintptr(integer)))
}

var (
	RT_CURSOR       = MakeIntResource(1)
	RT_BITMAP       = MakeIntResource(2)
	RT_ICON         = MakeIntResource(3)
	RT_MENU         = MakeIntResource(4)
	RT_DIALOG       = MakeIntResource(5)
	RT_STRING       = MakeIntResource(6)
	RT_FONTDIR      = MakeIntResource(7)
	RT_FONT         = MakeIntResource(8)
	RT_ACCELERATOR  = MakeIntResource(9)
	RT_RCDATA       = MakeIntResource(10)
	RT_MESSAGETABLE = MakeIntResource(11)
	RT_GROUP_CURSOR = MakeIntResource(12)
	RT_GROUP_ICON   = MakeIntResource(14)
	RT_VERSION      = MakeIntResource(16)
	RT_DLGINCLUDE   = MakeIntResource(17)
	RT_PLUGPLAY     = MakeIntResource(19)
	RT_VXD          = MakeIntResource(20)
	RT_ANICURSOR    = MakeIntResource(21)
	RT_ANIICON      = MakeIntResource(22)
	RT_HTML         = MakeIntResource(23)
	RT_MANIFEST     = MakeIntResource(24)
)
