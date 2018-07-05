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
	"github.com/winlabs/gowin32/wrappers"
)

type AddressFamily uint32

const (
	AddressFamilyUnspecified AddressFamily = wrappers.AF_UNSPEC
	AddressFamilyIP          AddressFamily = wrappers.AF_INET
	AddressFamilyIPX         AddressFamily = wrappers.AF_IPX
	AddressFamilyAppleTalk   AddressFamily = wrappers.AF_APPLETALK
	AddressFamilyNetBIOS     AddressFamily = wrappers.AF_NETBIOS
	AddressFamilyIPv6        AddressFamily = wrappers.AF_INET6
	AddressFamilyIrDA        AddressFamily = wrappers.AF_IRDA
	AddressFamilyBluetooth   AddressFamily = wrappers.AF_BTH
)
