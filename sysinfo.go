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
	"unsafe"

	"github.com/winlabs/gowin32/wrappers"
)

type ProcessorArchitecture uint16

const (
	ProcessorArchitectureIntel   ProcessorArchitecture = wrappers.PROCESSOR_ARCHITECTURE_INTEL
	ProcessorArchitectureMIPS    ProcessorArchitecture = wrappers.PROCESSOR_ARCHITECTURE_MIPS
	ProcessorArchitectureAlpha   ProcessorArchitecture = wrappers.PROCESSOR_ARCHITECTURE_ALPHA
	ProcessorArchitecturePowerPC ProcessorArchitecture = wrappers.PROCESSOR_ARCHITECTURE_PPC
	ProcessorArchitectureARM     ProcessorArchitecture = wrappers.PROCESSOR_ARCHITECTURE_ARM
	ProcessorArchitectureIA64    ProcessorArchitecture = wrappers.PROCESSOR_ARCHITECTURE_IA64
	ProcessorArchitectureAMD64   ProcessorArchitecture = wrappers.PROCESSOR_ARCHITECTURE_AMD64
)

type DisplayDevice struct {
	DeviceName   string
	DeviceString string
	StateFlags   DisplayDeviceStateFlags
	DeviceID     string
	DeviceKey    string
}

type DisplayDeviceStateFlags uint32

const (
	DisplayDeviceActive          DisplayDeviceStateFlags = wrappers.DISPLAY_DEVICE_ACTIVE
	DisplayDevicePrimaryDevice   DisplayDeviceStateFlags = wrappers.DISPLAY_DEVICE_PRIMARY_DEVICE
	DisplayDeviceMirroringDriver DisplayDeviceStateFlags = wrappers.DISPLAY_DEVICE_MIRRORING_DRIVER
	DisplayDeviceVGACompatible   DisplayDeviceStateFlags = wrappers.DISPLAY_DEVICE_VGA_COMPATIBLE
	DisplayDeviceRemovable       DisplayDeviceStateFlags = wrappers.DISPLAY_DEVICE_REMOVABLE
	DisplayDeviceModeSpruned     DisplayDeviceStateFlags = wrappers.DISPLAY_DEVICE_MODESPRUNED
)

type DisplayMonitorInfo struct {
	Handle        syscall.Handle
	DeviceContext syscall.Handle
	Rectangle     Rectangle
}

type ProcessorInfo struct {
	ProcessorArchitecture ProcessorArchitecture
	NumberOfProcessors    uint
	ProcessorLevel        uint
	ProcessorRevision     uint
}

func GetAllDisplayDevices() []DisplayDevice {
	result := make([]DisplayDevice, 0)
	var device *uint16
	var dd wrappers.DISPLAY_DEVICE
	dd.Cb = uint32(unsafe.Sizeof(dd))

	var i uint32
	for i = 0; ; i++ {
		if wrappers.EnumDisplayDevices(device, i, &dd, 0) {
			result = append(result,
				DisplayDevice{DeviceName: syscall.UTF16ToString(dd.DeviceName[:]),
					DeviceString: syscall.UTF16ToString(dd.DeviceString[:]),
					StateFlags:   DisplayDeviceStateFlags(dd.StateFlags),
					DeviceID:     syscall.UTF16ToString(dd.DeviceID[:]),
					DeviceKey:    syscall.UTF16ToString(dd.DeviceKey[:]),
				})
		} else {
			break
		}
	}
	return result

}

func GetProcessorInfo() *ProcessorInfo {
	var si wrappers.SYSTEM_INFO
	wrappers.GetSystemInfo(&si)
	return &ProcessorInfo{
		ProcessorArchitecture: ProcessorArchitecture(si.ProcessorArchitecture),
		NumberOfProcessors:    uint(si.NumberOfProcessors),
		ProcessorLevel:        uint(si.ProcessorLevel),
		ProcessorRevision:     uint(si.ProcessorRevision),
	}
}

func GetAllDisplayMonitors() []DisplayMonitorInfo {

	result := []DisplayMonitorInfo{}
	wrappers.EnumDisplayMonitors(
		syscall.Handle(0),
		nil,
		func(hmonitor syscall.Handle, hdc syscall.Handle, rect *wrappers.RECT, lparam uintptr) int32 {
			result = append(result, DisplayMonitorInfo{Handle: hmonitor, DeviceContext: hdc, Rectangle: rectToRectangle(*rect)})
			return 1
		},
		uintptr(0),
	)

	return result
}
