/*
 * Copyright (c) 2014-2015 MongoDB, Inc.
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
	ProcessBasicInformation           = 0
	ProcessQuotaLimits                = 1
	ProcessIoCounters                 = 2
	ProcessVmCounters                 = 3
	ProcessTimes                      = 4
	ProcessBasePriority               = 5
	ProcessRaisePriority              = 6
	ProcessDebugPort                  = 7
	ProcessExceptionPort              = 8
	ProcessAccessToken                = 9
	ProcessLdtInformation             = 10
	ProcessLdtSize                    = 11
	ProcessDefaultHardErrorMode       = 12
	ProcessIoPortHandlers             = 13
	ProcessPooledUsageAndLimits       = 14
	ProcessWorkingSetWatch            = 15
	ProcessUserModeIOPL               = 16
	ProcessEnableAlignmentFaultFixup  = 17
	ProcessPriorityClass              = 18
	ProcessWx86Information            = 19
	ProcessHandleCount                = 20
	ProcessAffinityMask               = 21
	ProcessPriorityBoost              = 22
	ProcessDeviceMap                  = 23
	ProcessSessionInformation         = 24
	ProcessForegroundInformation      = 25
	ProcessWow64Information           = 26
	ProcessImageFileName              = 27
	ProcessLUIDDeviceMapsEnabled      = 28
	ProcessBreakOnTermination         = 29
	ProcessDebugObjectHandle          = 30
	ProcessDebugFlags                 = 31
	ProcessHandleTracing              = 32
	ProcessIoPriority                 = 33
	ProcessExecuteFlags               = 34
	ProcessTlsInformation             = 35
	ProcessCookie                     = 36
	ProcessImageInformation           = 37
	ProcessCycleTime                  = 38
	ProcessPagePriority               = 39
	ProcessInstrumentationCallback    = 40
	ProcessThreadStackAllocation      = 41
	ProcessWorkingSetWatchEx          = 42
	ProcessImageFileNameWin32         = 43
	ProcessImageFileMapping           = 44
	ProcessAffinityUpdateMode         = 45
	ProcessMemoryAllocationMode       = 46
	ProcessGroupInformation           = 47
	ProcessTokenVirtualizationEnabled = 48
	ProcessConsoleHostProcess         = 49
	ProcessWindowInformation          = 50
	MaxProcessInfoClass               = 51
)

type PROCESS_BASIC_INFORMATION struct {
	ExitStatus                   int32
	PebBaseAddress               uintptr
	AffinityMask                 uintptr
	BasePriority                 int32
	UniqueProcessId              uintptr
	InheritedFromUniqueProcessId uintptr
}

var (
	modntdll = syscall.NewLazyDLL("ntdll.dll")

	procNtQueryInformationProcess = modntdll.NewProc("NtQueryInformationProcess")
)

func NtQueryInformationProcess(processHandle syscall.Handle, processInformationClass int32, processInformation *byte, processInformationLength uint32, returnLength *uint32) int32 {
	r1, _, _ := procNtQueryInformationProcess.Call(
		uintptr(processHandle),
		uintptr(processInformationClass),
		uintptr(unsafe.Pointer(processInformation)),
		uintptr(processInformationLength),
		uintptr(unsafe.Pointer(returnLength)))
	return int32(r1)
}
