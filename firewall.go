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
	"unsafe"
)

type FirewallIPVersion int32

const (
	FirewallIPV4         FirewallIPVersion = wrappers.NET_FW_IP_VERSION_V4
	FirewallIPV6         FirewallIPVersion = wrappers.NET_FW_IP_VERSION_V6
	FirewallAnyIPVersion FirewallIPVersion = wrappers.NET_FW_IP_VERSION_ANY
)

type FirewallProtocol int32

const (
	FirewallTCP         FirewallProtocol = wrappers.NET_FW_PROTOCOL_TCP
	FirewallUDP         FirewallProtocol = wrappers.NET_FW_PROTOCOL_UDP
	FirewallAnyProtocol FirewallProtocol = wrappers.NET_FW_PROTOCOL_ANY
)

type FirewallManager struct {
	object *wrappers.INetFwMgr
}

func CreateFirewallManager() (*FirewallManager, error) {
	object := uintptr(0)
	err := wrappers.CoCreateInstance(
		&wrappers.CLSID_NetFwMgr,
		nil,
		wrappers.CLSCTX_INPROC_SERVER,
		&wrappers.IID_INetFwMgr,
		&object)
	if err != nil {
		return nil, err
	}
	return &FirewallManager{object: (*wrappers.INetFwMgr)(unsafe.Pointer(object))}, nil
}

func (self *FirewallManager) Close() error {
	if self.object != nil {
		self.object.Release()
		self.object = nil
	}
	return nil
}

func (self *FirewallManager) IsPortAllowed(imageFileName string, ipVersion FirewallIPVersion, portNumber int32, localAddress string, ipProtocol FirewallProtocol) (allowed bool, restricted bool, err error) {
	var imageFileNameRaw *uint16
	if imageFileName != "" {
		imageFileNameRaw = wrappers.SysAllocString(syscall.StringToUTF16Ptr(imageFileName))
		defer wrappers.SysFreeString(imageFileNameRaw)
	}
	var localAddressRaw *uint16
	if localAddress != "" {
		localAddressRaw = wrappers.SysAllocString(syscall.StringToUTF16Ptr(localAddress))
		defer wrappers.SysFreeString(localAddressRaw)
	}
	var allowedRaw wrappers.Variant
	wrappers.VariantInit(&allowedRaw)
	defer wrappers.VariantClear(&allowedRaw)
	var restrictedRaw wrappers.Variant
	wrappers.VariantInit(&restrictedRaw)
	defer wrappers.VariantClear(&restrictedRaw)
	err = self.object.IsPortAllowed(
		imageFileNameRaw,
		int32(ipVersion),
		portNumber,
		localAddressRaw,
		int32(ipProtocol),
		&allowedRaw,
		&restrictedRaw)
	if err == nil {
		allowed = allowedRaw.Vt == wrappers.VT_BOOL && int16(allowedRaw.Val) != wrappers.VARIANT_FALSE
		restricted = restrictedRaw.Vt == wrappers.VT_BOOL && int16(restrictedRaw.Val) != wrappers.VARIANT_FALSE
	}
	return
}
