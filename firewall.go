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

type FirewallDirection int32

const (
	FirewallInbound  FirewallDirection = wrappers.NET_FW_RULE_DIR_IN
	FirewallOutbound FirewallDirection = wrappers.NET_FW_RULE_DIR_OUT
)

type FirewallAction int32

const (
	FirewallBlock FirewallAction = wrappers.NET_FW_ACTION_BLOCK
	FirewallAllow FirewallAction = wrappers.NET_FW_ACTION_ALLOW
)

type FirewallRule struct {
	object *wrappers.INetFwRule
}

func NewFirewallRule() (*FirewallRule, error) {
	var object uintptr
	err := wrappers.CoCreateInstance(
		&wrappers.CLSID_NetFwRule,
		nil,
		wrappers.CLSCTX_INPROC_SERVER,
		&wrappers.IID_INetFwRule,
		&object)
	if err != nil {
		return nil, err
	}
	return &FirewallRule{object: (*wrappers.INetFwRule)(unsafe.Pointer(object))}, nil
}

func (self *FirewallRule) Close() error {
	if self.object != nil {
		self.object.Release()
		self.object = nil
	}
	return nil
}

func (self *FirewallRule) GetName() (string, error) {
	if self.object == nil {
		return "", wrappers.E_POINTER
	}
	var nameRaw *uint16
	if err := self.object.Get_Name(&nameRaw); err != nil {
		return "", err
	}
	return BstrToString(nameRaw), nil
}

func (self *FirewallRule) SetName(name string) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	nameRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(name))
	defer wrappers.SysFreeString(nameRaw)
	return self.object.Put_Name(nameRaw)
}

func (self *FirewallRule) GetDescription() (string, error) {
	if self.object == nil {
		return "", wrappers.E_POINTER
	}
	var descRaw *uint16
	if err := self.object.Get_Description(&descRaw); err != nil {
		return "", err
	}
	return BstrToString(descRaw), nil
}

func (self *FirewallRule) SetDescription(desc string) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	descRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(desc))
	defer wrappers.SysFreeString(descRaw)
	return self.object.Put_Description(descRaw)
}

func (self *FirewallRule) GetApplicationName() (string, error) {
	if self.object == nil {
		return "", wrappers.E_POINTER
	}
	var imageFileNameRaw *uint16
	if err := self.object.Get_ApplicationName(&imageFileNameRaw); err != nil {
		return "", err
	}
	return BstrToString(imageFileNameRaw), nil
}

func (self *FirewallRule) SetApplicationName(imageFileName string) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	imageFileNameRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(imageFileName))
	defer wrappers.SysFreeString(imageFileNameRaw)
	return self.object.Put_ApplicationName(imageFileNameRaw)
}

func (self *FirewallRule) GetServiceName() (string, error) {
	if self.object == nil {
		return "", wrappers.E_POINTER
	}
	var serviceNameRaw *uint16
	if err := self.object.Get_ServiceName(&serviceNameRaw); err != nil {
		return "", err
	}
	return BstrToString(serviceNameRaw), nil
}

func (self *FirewallRule) SetServiceName(serviceName string) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	serviceNameRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(serviceName))
	defer wrappers.SysFreeString(serviceNameRaw)
	return self.object.Put_ServiceName(serviceNameRaw)
}

func (self *FirewallRule) GetProtocol() (FirewallProtocol, error) {
	if self.object == nil {
		return 0, wrappers.E_POINTER
	}
	var protocolRaw int32
	if err := self.object.Get_Protocol(&protocolRaw); err != nil {
		return 0, err
	}
	return FirewallProtocol(protocolRaw), nil
}

func (self *FirewallRule) SetProtocol(protocol FirewallProtocol) error{
	if self.object == nil {
		return wrappers.E_POINTER
	}
	return self.object.Put_Protocol(int32(protocol))
}

func (self *FirewallRule) GetLocalPorts() (string, error) {
	if self.object == nil {
		return "", wrappers.E_POINTER
	}
	var portNumbersRaw *uint16
	if err := self.object.Get_LocalPorts(&portNumbersRaw); err != nil {
		return "", err
	}
	return BstrToString(portNumbersRaw), nil
}

func (self *FirewallRule) SetLocalPorts(portNumbers string) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	portNumbersRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(portNumbers))
	defer wrappers.SysFreeString(portNumbersRaw)
	return self.object.Put_LocalPorts(portNumbersRaw)
}

func (self *FirewallRule) GetRemotePorts() (string, error) {
	if self.object == nil {
		return "", wrappers.E_POINTER
	}
	var portNumbersRaw *uint16
	if err := self.object.Get_RemotePorts(&portNumbersRaw); err != nil {
		return "", err
	}
	return BstrToString(portNumbersRaw), nil
}

func (self *FirewallRule) SetRemotePorts(portNumbers string) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	portNumbersRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(portNumbers))
	defer wrappers.SysFreeString(portNumbersRaw)
	return self.object.Put_RemotePorts(portNumbersRaw)
}

func (self *FirewallRule) GetLocalAddresses() (string, error) {
	if self.object == nil {
		return "", wrappers.E_POINTER
	}
	var localAddrsRaw *uint16
	if err := self.object.Get_LocalAddresses(&localAddrsRaw); err != nil {
		return "", err
	}
	return BstrToString(localAddrsRaw), nil
}

func (self *FirewallRule) SetLocalAddresses(localAddrs string) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	localAddrsRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(localAddrs))
	defer wrappers.SysFreeString(localAddrsRaw)
	return self.object.Put_LocalAddresses(localAddrsRaw)
}

func (self *FirewallRule) GetRemoteAddresses() (string, error) {
	if self.object == nil {
		return "", wrappers.E_POINTER
	}
	var remoteAddrsRaw *uint16
	if err := self.object.Get_RemoteAddresses(&remoteAddrsRaw); err != nil {
		return "", err
	}
	return BstrToString(remoteAddrsRaw), nil
}

func (self *FirewallRule) SetRemoteAddresses(remoteAddrs string) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	remoteAddrsRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(remoteAddrs))
	defer wrappers.SysFreeString(remoteAddrsRaw)
	return self.object.Put_RemoteAddresses(remoteAddrsRaw)
}

func (self *FirewallRule) GetIcmpTypesAndCodes() (string, error) {
	if self.object == nil {
		return "", wrappers.E_POINTER
	}
	var icmpTypesAndCodesRaw *uint16
	if err := self.object.Get_IcmpTypesAndCodes(&icmpTypesAndCodesRaw); err != nil {
		return "", err
	}
	return BstrToString(icmpTypesAndCodesRaw), nil
}

func (self *FirewallRule) SetIcmpTypesAndCodes(icmpTypesAndCodes string) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	icmpTypesAndCodesRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(icmpTypesAndCodes))
	defer wrappers.SysFreeString(icmpTypesAndCodesRaw)
	return self.object.Put_IcmpTypesAndCodes(icmpTypesAndCodesRaw)
}

func (self *FirewallRule) GetDirection() (FirewallDirection, error) {
	if self.object == nil {
		return 0, wrappers.E_POINTER
	}
	var dirRaw int32
	if err := self.object.Get_Direction(&dirRaw); err != nil {
		return 0, err
	}
	return FirewallDirection(dirRaw), nil
}

func (self *FirewallRule) SetDirection(dir FirewallDirection) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	return self.object.Put_Direction(int32(dir))
}

func (self *FirewallRule) GetInterfaceTypes() (string, error) {
	if self.object == nil {
		return "", wrappers.E_POINTER
	}
	var interfaceTypesRaw *uint16
	if err := self.object.Get_InterfaceTypes(&interfaceTypesRaw); err != nil {
		return "", err
	}
	return BstrToString(interfaceTypesRaw), nil
}

func (self *FirewallRule) SetInterfaceTypes(interfaceTypes string) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	interfaceTypesRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(interfaceTypes))
	defer wrappers.SysFreeString(interfaceTypesRaw)
	return self.object.Put_InterfaceTypes(interfaceTypesRaw)
}

func (self *FirewallRule) GetEnabled() (bool, error) {
	if self.object == nil {
		return false, wrappers.E_POINTER
	}
	var enabled bool
	if err := self.object.Get_Enabled(&enabled); err != nil {
		return false, err
	}
	return enabled, nil
}

func (self *FirewallRule) SetEnabled(enabled bool) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	return self.object.Put_Enabled(enabled)
}

func (self *FirewallRule) GetGrouping() (string, error) {
	if self.object == nil {
		return "", wrappers.E_POINTER
	}
	var contextRaw *uint16
	if err := self.object.Get_Grouping(&contextRaw); err != nil {
		return "", err
	}
	return BstrToString(contextRaw), nil
}

func (self *FirewallRule) SetGrouping(context string) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	contextRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(context))
	defer wrappers.SysFreeString(contextRaw)
	return self.object.Put_Grouping(contextRaw)
}

func (self *FirewallRule) GetEdgeTraversal() (bool, error) {
	if self.object == nil {
		return false, wrappers.E_POINTER
	}
	var enabled bool
	if err := self.object.Get_EdgeTraversal(&enabled); err != nil {
		return false, err
	}
	return enabled, nil
}

func (self *FirewallRule) SetEdgeTraversal(enabled bool) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	return self.object.Put_EdgeTraversal(enabled)
}

func (self *FirewallRule) GetAction() (FirewallAction, error) {
	if self.object == nil {
		return 0, wrappers.E_POINTER
	}
	var actionRaw int32
	if err := self.object.Get_Action(&actionRaw); err != nil {
		return 0, err
	}
	return FirewallAction(actionRaw), nil
}

func (self *FirewallRule) SetAction(action FirewallAction) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	return self.object.Put_Action(int32(action))
}

type FirewallRuleCollection struct {
	object *wrappers.INetFwRules
}

func (self *FirewallRuleCollection) Close() error {
	if self.object != nil {
		self.object.Release()
		self.object = nil
	}
	return nil
}

func (self *FirewallRuleCollection) GetCount() (int32, error) {
	if self.object == nil {
		return 0, wrappers.E_POINTER
	}
	var count int32
	if err := self.object.Get_Count(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (self *FirewallRuleCollection) Add(rule *FirewallRule) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	return self.object.Add(rule.object)
}

func (self *FirewallRuleCollection) Remove(name string) error {
	if self.object == nil {
		return wrappers.E_POINTER
	}
	nameRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(name))
	defer wrappers.SysFreeString(nameRaw)
	return self.object.Remove(nameRaw)
}

func (self *FirewallRuleCollection) Item(name string) (*FirewallRule, error) {
	if self.object == nil {
		return nil, wrappers.E_POINTER
	}
	nameRaw := wrappers.SysAllocString(syscall.StringToUTF16Ptr(name))
	defer wrappers.SysFreeString(nameRaw)
	var rule *wrappers.INetFwRule
	if err := self.object.Item(nameRaw, &rule); err != nil {
		return nil, err
	}
	return &FirewallRule{object: rule}, nil
}

type FirewallPolicy struct {
	object *wrappers.INetFwPolicy2
}

func NewFirewallPolicy() (*FirewallPolicy, error) {
	var object uintptr
	err := wrappers.CoCreateInstance(
		&wrappers.CLSID_NetFwPolicy2,
		nil,
		wrappers.CLSCTX_INPROC_SERVER,
		&wrappers.IID_INetFwPolicy2,
		&object)
	if err != nil {
		return nil, err
	}
	return &FirewallPolicy{object: (*wrappers.INetFwPolicy2)(unsafe.Pointer(object))}, nil
}

func (self *FirewallPolicy) Close() error {
	if self.object != nil {
		self.object.Release()
		self.object = nil
	}
	return nil
}

func (self *FirewallPolicy) GetRules() (*FirewallRuleCollection, error) {
	if self.object == nil {
		return nil, wrappers.E_POINTER
	}
	var rules *wrappers.INetFwRules
	if err := self.object.Get_Rules(&rules); err != nil {
		return nil, err
	}
	return &FirewallRuleCollection{object: rules}, nil
}

type FirewallManager struct {
	object *wrappers.INetFwMgr
}

func NewFirewallManager() (*FirewallManager, error) {
	var object uintptr
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
	if self.object == nil {
		err = wrappers.E_POINTER
		return
	}
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
