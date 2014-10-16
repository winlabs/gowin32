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

var (
	IID_INetFwRule  = syscall.GUID{0xAF230D27, 0xBABA, 0x4E42, [8]byte{0xAC, 0xED, 0xF5, 0x24, 0xF2, 0x2C, 0xFC, 0xE2}}
	IID_INetFwRules = syscall.GUID{0x9C4C6277, 0x5027, 0x441E, [8]byte{0xAF, 0xAE, 0xCA, 0x1F, 0x54, 0x2D, 0xA0, 0x09}}
)

type INetFwRuleVtbl struct {
	IDispatchVtbl
	Get_Name              uintptr
	Put_Name              uintptr
	Get_Description       uintptr
	Put_Description       uintptr
	Get_ApplicationName   uintptr
	Put_ApplicationName   uintptr
	Get_ServiceName       uintptr
	Put_ServiceName       uintptr
	Get_Protocol          uintptr
	Put_Protocol          uintptr
	Get_LocalPorts        uintptr
	Put_LocalPorts        uintptr
	Get_RemotePorts       uintptr
	Put_RemotePorts       uintptr
	Get_LocalAddresses    uintptr
	Put_LocalAddresses    uintptr
	Get_RemoteAddresses   uintptr
	Put_RemoteAddresses   uintptr
	Get_IcmpTypesAndCodes uintptr
	Put_IcmpTypesAndCodes uintptr
	Get_Direction         uintptr
	Put_Direction         uintptr
	Get_Interfaces        uintptr
	Put_Interfaces        uintptr
	Get_InterfaceTypes    uintptr
	Put_InterfaceTypes    uintptr
	Get_Enabled           uintptr
	Put_Enabled           uintptr
	Get_Grouping          uintptr
	Put_Grouping          uintptr
	Get_Profiles          uintptr
	Put_Profiles          uintptr
	Get_EdgeTraversal     uintptr
	Put_EdgeTraversal     uintptr
	Get_Action            uintptr
	Put_Action            uintptr
}

type INetFwRule struct {
	IDispatch
}

func (self *INetFwRule) Get_Name(name **uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_Name,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(name)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_Name(name *uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_Name,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(name)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_Description(desc **uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_Description,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(desc)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_Description(desc *uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_Description,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(desc)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_ApplicationName(imageFileName **uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_ApplicationName,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(imageFileName)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_ApplicationName(imageFileName *uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_ApplicationName,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(imageFileName)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_ServiceName(serviceName **uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_ServiceName,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(serviceName)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_ServiceName(serviceName *uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_ServiceName,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(serviceName)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_Protocol(protocol *int32) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_Protocol,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(protocol)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_Protocol(protocol int32) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_Protocol,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(protocol),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_LocalPorts(portNumbers **uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_LocalPorts,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(portNumbers)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_LocalPorts(portNumbers *uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_LocalPorts,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(portNumbers)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_RemotePorts(portNumbers **uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_RemotePorts,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(portNumbers)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_RemotePorts(portNumbers *uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_RemotePorts,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(portNumbers)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_LocalAddresses(localAddrs **uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_LocalAddresses,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(localAddrs)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_LocalAddresses(localAddrs *uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_LocalAddresses,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(localAddrs)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_RemoteAddresses(remoteAddrs **uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_RemoteAddresses,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(remoteAddrs)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_RemoteAddresses(remoteAddrs *uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_RemoteAddresses,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(remoteAddrs)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_IcmpTypesAndCodes(icmpTypesAndCodes **uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_IcmpTypesAndCodes,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(icmpTypesAndCodes)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_IcmpTypesAndCodes(icmpTypesAndCodes *uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_IcmpTypesAndCodes,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(icmpTypesAndCodes)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_Direction(dir *int32) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_Direction,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(dir)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_Direction(dir int32) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_Direction,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(dir),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_InterfaceTypes(interfaceTypes **uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_InterfaceTypes,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(interfaceTypes)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_InterfaceTypes(interfaceTypes *uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_InterfaceTypes,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(interfaceTypes)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_Enabled(enabled *bool) error {
	if enabled == nil {
		return E_POINTER
	}
	enabledRaw := int16(VARIANT_FALSE)
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_Enabled,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&enabledRaw)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	*enabled = (enabledRaw != VARIANT_FALSE)
	return nil
}

func (self *INetFwRule) Put_Enabled(enabled bool) error {
	var enabledRaw int16
	if enabled {
		enabledRaw = VARIANT_TRUE
	} else {
		enabledRaw = VARIANT_FALSE
	}
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_Enabled,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(enabledRaw),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_Grouping(context **uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_Grouping,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(context)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_Grouping(context *uint16) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_Grouping,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(context)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_Profiles(profileTypesBitmask *int32) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_Profiles,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(profileTypesBitmask)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_Profiles(profileTypesBitmask int32) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_Profiles,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(profileTypesBitmask),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_EdgeTraversal(enabled *bool) error {
	if enabled == nil {
		return E_POINTER
	}
	enabledRaw := int16(VARIANT_FALSE)
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_EdgeTraversal,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(&enabledRaw)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	*enabled = (enabledRaw != VARIANT_FALSE)
	return nil
}

func (self *INetFwRule) Put_EdgeTraversal(enabled bool) error {
	var enabledRaw int16
	if enabled {
		enabledRaw = VARIANT_TRUE
	} else {
		enabledRaw = VARIANT_FALSE
	}
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_EdgeTraversal,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(enabledRaw),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Get_Action(action *int32) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_Action,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(action)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Put_Action(action int32) error {
	vtbl := (*INetFwRuleVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Put_Action,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(action),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

type INetFwRulesVtbl struct {
	IDispatchVtbl
	Get_Count    uintptr
	Add          uintptr
	Remove       uintptr
	Item         uintptr
	Get__NewEnum uintptr
}

type INetFwRules struct {
	IDispatch
}

func (self *INetFwRules) Get_Count(count *int32) error {
	vtbl := (*INetFwRulesVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Get_Count,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(count)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRules) Add(rule *INetFwRule) error {
	vtbl := (*INetFwRulesVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Add,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(rule)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRules) Remove(name *uint16) error {
	vtbl := (*INetFwRulesVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Remove,
		2,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(name)),
		0)
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}

func (self *INetFwRule) Item(name *uint16, rule **INetFwRule) error {
	vtbl := (*INetFwRulesVtbl)(unsafe.Pointer(self.Vtbl))
	r1, _, _ := syscall.Syscall(
		vtbl.Item,
		3,
		uintptr(unsafe.Pointer(self)),
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(rule)))
	if int32(r1) < 0 {
		return syscall.Errno(r1)
	}
	return nil
}
