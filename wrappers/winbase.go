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
	INVALID_HANDLE_VALUE = ^syscall.Handle(0)
)

const (
	WAIT_FAILED        = 0xFFFFFFFF
	WAIT_OBJECT_0      = STATUS_WAIT_0
	WAIT_ABANDONED     = STATUS_ABANDONED_WAIT_0
	WAIT_ABANDONED_0   = STATUS_ABANDONED_WAIT_0
	WAIT_IO_COMPLETION = STATUS_USER_APC
)

const (
	FILE_FLAG_WRITE_THROUGH       = 0x80000000
	FILE_FLAG_OVERLAPPED          = 0x40000000
	FILE_FLAG_NO_BUFFERING        = 0x20000000
	FILE_FLAG_RANDOM_ACCESS       = 0x10000000
	FILE_FLAG_SEQUENTIAL_SCAN     = 0x08000000
	FILE_FLAG_DELETE_ON_CLOSE     = 0x04000000
	FILE_FLAG_BACKUP_SEMANTICS    = 0x02000000
	FILE_FLAG_POSIX_SEMANTICS     = 0x01000000
	FILE_FLAG_OPEN_REPARSE_POINT  = 0x00200000
	FILE_FLAG_OPEN_NO_RECALL      = 0x00100000
	FILE_FLAG_FIRST_PIPE_INSTANCE = 0x00080000
)

const (
	CREATE_NEW        = 1
	CREATE_ALWAYS     = 2
	OPEN_EXISTING     = 3
	OPEN_ALWAYS       = 4
	TRUNCATE_EXISTING = 5
)

const (
	SECURITY_ANONYMOUS        = SecurityAnonymous << 16
	SECURITY_IDENTIFICATION   = SecurityIdentification << 16
	SECURITY_IMPERSONATION    = SecurityImpersonation << 16
	SECURITY_DELEGATION       = SecurityDelegation << 16
	SECURITY_CONTEXT_TRACKING = 0x00040000
	SECURITY_EFFECTIVE_ONLY   = 0x00080000
)

type SECURITY_ATTRIBUTES struct {
	Length             uint32
	SecurityDescriptor *byte
	InheritHandle      int32
}

type PROCESS_INFORMATION struct {
	Process   syscall.Handle
	Thread    syscall.Handle
	ProcessId uint32
	ThreadId  uint32
}

type FILETIME struct {
	LowDateTime  uint32
	HighDateTime uint32
}

type SYSTEM_INFO struct {
	ProcessorArchitecture     uint16
	Reserved                  uint16
	PageSize                  uint32
	MinimumApplicationAddress *byte
	MaximumApplicationAddress *byte
	ActiveProcessorMask       uintptr
	NumberOfProcessors        uint32
	ProcessorType             uint32
	AllocationGranularity     uint32
	ProcessorLevel            uint16
	ProcessorRevision         uint16
}

const (
	DEBUG_PROCESS                    = 0x00000001
	DEBUG_ONLY_THIS_PROCESS          = 0x00000002
	CREATE_SUSPENDED                 = 0x00000004
	DETACHED_PROCESS                 = 0x00000008
	CREATE_NEW_CONSOLE               = 0x00000010
	NORMAL_PRIORITY_CLASS            = 0x00000020
	IDLE_PRIORITY_CLASS              = 0x00000040
	HIGH_PRIORITY_CLASS              = 0x00000080
	REALTIME_PRIORITY_CLASS          = 0x00000100
	CREATE_NEW_PROCESS_GROUP         = 0x00000200
	CREATE_UNICODE_ENVIRONMENT       = 0x00000400
	CREATE_SEPARATE_WOW_VDM          = 0x00000800
	CREATE_SHARED_WOW_VDM            = 0x00001000
	BELOW_NORMAL_PRIORITY_CLASS      = 0x00004000
	ABOVE_NORMAL_PRIORITY_CLASS      = 0x00008000
	INHERIT_PARENT_AFFINITY          = 0x00010000
	CREATE_PROTECTED_PROCESS         = 0x00040000
	EXTENDED_STARTUPINFO_PRESENT     = 0x00080000
	PROCESS_MODE_BACKGROUND_BEGIN    = 0x00100000
	PROCESS_MODE_BACKGROUND_END      = 0x00200000
	CREATE_BREAKAWAY_FROM_JOB        = 0x01000000
	CREATE_PRESERVE_CODE_AUTHZ_LEVEL = 0x02000000
	CREATE_DEFAULT_ERROR_MODE        = 0x04000000
	CREATE_NO_WINDOW                 = 0x08000000
)

const (
	DRIVE_UNKNOWN     = 0
	DRIVE_NO_ROOT_DIR = 1
	DRIVE_REMOVABLE   = 2
	DRIVE_FIXED       = 3
	DRIVE_REMOTE      = 4
	DRIVE_CDROM       = 5
	DRIVE_RAMDISK     = 6
)

const (
	STD_INPUT_HANDLE  = ^uint32(10) + 1
	STD_OUTPUT_HANDLE = ^uint32(11) + 1
	STD_ERROR_HANDLE  = ^uint32(12) + 1
)

const (
	INFINITE = 0xFFFFFFFF
)

const (
	FORMAT_MESSAGE_ALLOCATE_BUFFER = 0x00000100
	FORMAT_MESSAGE_IGNORE_INSERTS  = 0x00000200
	FORMAT_MESSAGE_FROM_STRING     = 0x00000400
	FORMAT_MESSAGE_FROM_HMODULE    = 0x00000800
	FORMAT_MESSAGE_FROM_SYSTEM     = 0x00001000
	FORMAT_MESSAGE_ARGUMENT_ARRAY  = 0x00002000
	FORMAT_MESSAGE_MAX_WIDTH_MASK  = 0x000000FF
)

type STARTUPINFO struct {
	Cb            uint32
	Reserved      *uint16
	Desktop       *uint16
	Title         *uint16
	X             uint32
	Y             uint32
	XSize         uint32
	YSize         uint32
	XCountChars   uint32
	YCountChars   uint32
	FillAttribute uint32
	Flags         uint32
	ShowWindow    uint16
	CbReserved2   uint16
	Reserved2     *byte
	StdInput      syscall.Handle
	StdOutput     syscall.Handle
	StdError      syscall.Handle
}

const (
	STARTF_USESHOWWINDOW    = 0x00000001
	STARTF_USESIZE          = 0x00000002
	STARTF_USEPOSITION      = 0x00000004
	STARTF_USECOUNTCHARS    = 0x00000008
	STARTF_USEFILLATTRIBUTE = 0x00000010
	STARTF_RUNFULLSCREEN    = 0x00000020
	STARTF_FORCEONFEEDBACK  = 0x00000040
	STARTF_FORCEOFFFEEDBACK = 0x00000080
	STARTF_USESTDHANDLES    = 0x00000100
	STARTF_USEHOTKEY        = 0x00000200
	STARTF_TITLEISLINKNAME  = 0x00000800
	STARTF_TITLEISAPPID     = 0x00001000
	STARTF_PREVENTPINNING   = 0x00002000
)

const (
	PROCESS_NAME_NATIVE = 0x00000001
)

const (
	ComputerNameNetBIOS                   = 0
	ComputerNameDnsHostname               = 1
	ComputerNameDnsDomain                 = 2
	ComputerNameDnsFullyQualified         = 3
	ComputerNamePhysicalNetBIOS           = 4
	ComputerNamePhysicalDnsHostname       = 5
	ComputerNamePhysicalDnsDomain         = 6
	ComputerNamePhysicalDnsFullyQualified = 7
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	modadvapi32 = syscall.NewLazyDLL("advapi32.dll")

	procAssignProcessToJobObject   = modkernel32.NewProc("AssignProcessToJobObject")
	procBeginUpdateResourceW       = modkernel32.NewProc("BeginUpdateResourceW")
	procCloseHandle                = modkernel32.NewProc("CloseHandle")
	procCreateFileW                = modkernel32.NewProc("CreateFileW")
	procCreateJobObjectW           = modkernel32.NewProc("CreateJobObjectW")
	procCreateProcessW             = modkernel32.NewProc("CreateProcessW")
	procDeviceIoControl            = modkernel32.NewProc("DeviceIoControl")
	procEndUpdateResourceW         = modkernel32.NewProc("EndUpdateResourceW")
	procExpandEnvironmentStringsW  = modkernel32.NewProc("ExpandEnvironmentStringsW")
	procFormatMessageW             = modkernel32.NewProc("FormatMessageW")
	procGetComputerNameExW         = modkernel32.NewProc("GetComputerNameExW")
	procGetCurrentProcess          = modkernel32.NewProc("GetCurrentProcess")
	procGetDriveTypeW              = modkernel32.NewProc("GetDriveTypeW")
	procGetDiskFreeSpaceExW        = modkernel32.NewProc("GetDiskFreeSpaceExW")
	procGetModuleFileNameW         = modkernel32.NewProc("GetModuleFileNameW")
	procGetStdHandle               = modkernel32.NewProc("GetStdHandle")
	procGetSystemDirectoryW        = modkernel32.NewProc("GetSystemDirectoryW")
	procGetSystemInfo              = modkernel32.NewProc("GetSystemInfo")
	procGetSystemTimeAsFileTime    = modkernel32.NewProc("GetSystemTimeAsFileTime")
	procGetSystemTimes             = modkernel32.NewProc("GetSystemTimes")
	procGetVolumeInformationW      = modkernel32.NewProc("GetVolumeInformationW")
	procIsProcessInJob             = modkernel32.NewProc("IsProcessInJob")
	procLocalFree                  = modkernel32.NewProc("LocalFree")
	procOpenJobObjectW             = modkernel32.NewProc("OpenJobObjectW")
	procOpenProcess                = modkernel32.NewProc("OpenProcess")
	procQueryFullProcessImageNameW = modkernel32.NewProc("QueryFullProcessImageNameW")
	procQueryInformationJobObject  = modkernel32.NewProc("QueryInformationJobObject")
	procSetFileTime                = modkernel32.NewProc("SetFileTime")
	procSetInformationJobObject    = modkernel32.NewProc("SetInformationJobObject")
	procSetStdHandle               = modkernel32.NewProc("SetStdHandle")
	procTerminateJobObject         = modkernel32.NewProc("TerminateJobObject")
	procTerminateProcess           = modkernel32.NewProc("TerminateProcess")
	procUpdateResourceW            = modkernel32.NewProc("UpdateResourceW")
	procVerifyVersionInfoW         = modkernel32.NewProc("VerifyVersionInfoW")
	procWaitForSingleObject        = modkernel32.NewProc("WaitForSingleObject")
	proclstrlenW                   = modkernel32.NewProc("lstrlenW")

	procAllocateAndInitializeSid   = modadvapi32.NewProc("AllocateAndInitializeSid")
	procCheckTokenMembership       = modadvapi32.NewProc("CheckTokenMembership")
	procCopySid                    = modadvapi32.NewProc("CopySid")
	procDeregisterEventSource      = modadvapi32.NewProc("DeregisterEventSource")
	procEqualSid                   = modadvapi32.NewProc("EqualSid")
	procFreeSid                    = modadvapi32.NewProc("FreeSid")
	procGetFileSecurityW           = modadvapi32.NewProc("GetFileSecurityW")
	procGetLengthSid               = modadvapi32.NewProc("GetLengthSid")
	procGetSecurityDescriptorOwner = modadvapi32.NewProc("GetSecurityDescriptorOwner")
	procGetTokenInformation        = modadvapi32.NewProc("GetTokenInformation")
	procOpenProcessToken           = modadvapi32.NewProc("OpenProcessToken")
	procRegisterEventSourceW       = modadvapi32.NewProc("RegisterEventSourceW")
	procReportEventW               = modadvapi32.NewProc("ReportEventW")
)

func AssignProcessToJobObject(job syscall.Handle, process syscall.Handle) error {
	r1, _, e1 := procAssignProcessToJobObject.Call(uintptr(job), uintptr(process))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func BeginUpdateResource(fileName *uint16, deleteExistingResources bool) (syscall.Handle, error) {
	var deleteExistingResourcesRaw int32
	if deleteExistingResources {
		deleteExistingResourcesRaw = 1
	} else {
		deleteExistingResourcesRaw = 0
	}
	r1, _, e1 := procBeginUpdateResourceW.Call(
		uintptr(unsafe.Pointer(fileName)),
		uintptr(deleteExistingResourcesRaw))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}
	return syscall.Handle(r1), nil
}

func CloseHandle(object syscall.Handle) error {
	r1, _, e1 := procCloseHandle.Call(uintptr(object))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func CreateFile(fileName *uint16, desiredAccess uint32, shareMode uint32, securityAttributes *SECURITY_ATTRIBUTES, creationDisposition uint32, flagsAndAttributes uint32, templateFile syscall.Handle) (syscall.Handle, error) {
	r1, _, e1 := procCreateFileW.Call(
		uintptr(unsafe.Pointer(fileName)),
		uintptr(desiredAccess),
		uintptr(shareMode),
		uintptr(unsafe.Pointer(securityAttributes)),
		uintptr(creationDisposition),
		uintptr(flagsAndAttributes),
		uintptr(templateFile))
	handle := syscall.Handle(r1)
	if handle == INVALID_HANDLE_VALUE {
		if e1 != ERROR_SUCCESS {
			return handle, e1
		} else {
			return handle, syscall.EINVAL
		}
	}
	return handle, nil
}

func CreateJobObject(jobAttributes* SECURITY_ATTRIBUTES, name *uint16) (syscall.Handle, error) {
	r1, _, e1 := procCreateJobObjectW.Call(
		uintptr(unsafe.Pointer(jobAttributes)),
		uintptr(unsafe.Pointer(name)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}
	return syscall.Handle(r1), nil
}

func CreateProcess(applicationName *uint16, commandLine *uint16, processAttributes *SECURITY_ATTRIBUTES, threadAttributes *SECURITY_ATTRIBUTES, inheritHandles bool, creationFlags uint32, environment *byte, currentDirectory *uint16, startupInfo *STARTUPINFO, processInformation *PROCESS_INFORMATION) error {
	var inheritHandlesRaw int32
	if inheritHandles {
		inheritHandlesRaw = 1
	} else {
		inheritHandlesRaw = 0
	}
	r1, _, e1 := procCreateProcessW.Call(
		uintptr(unsafe.Pointer(applicationName)),
		uintptr(unsafe.Pointer(commandLine)),
		uintptr(unsafe.Pointer(processAttributes)),
		uintptr(unsafe.Pointer(threadAttributes)),
		uintptr(inheritHandlesRaw),
		uintptr(creationFlags),
		uintptr(unsafe.Pointer(environment)),
		uintptr(unsafe.Pointer(currentDirectory)),
		uintptr(unsafe.Pointer(startupInfo)),
		uintptr(unsafe.Pointer(processInformation)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func DeviceIoControl(device syscall.Handle, ioControlCode uint32, inBuffer *byte, inBufferSize uint32, outBuffer *byte, outBufferSize uint32, bytesReturned *uint32, overlapped *syscall.Overlapped) error {
	r1, _, e1 := procDeviceIoControl.Call(
		uintptr(device),
		uintptr(ioControlCode),
		uintptr(unsafe.Pointer(inBuffer)),
		uintptr(inBufferSize),
		uintptr(unsafe.Pointer(outBuffer)),
		uintptr(outBufferSize),
		uintptr(unsafe.Pointer(bytesReturned)),
		uintptr(unsafe.Pointer(overlapped)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func EndUpdateResource(update syscall.Handle, discard bool) error {
	var discardRaw int32
	if discard {
		discardRaw = 1
	} else {
		discardRaw = 0
	}
	r1, _, e1 := procEndUpdateResourceW.Call(
		uintptr(update),
		uintptr(discardRaw))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func ExpandEnvironmentStrings(src *uint16, dst *uint16, size uint32) (uint32, error) {
	r1, _, e1 := procExpandEnvironmentStringsW.Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(size))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}
	return uint32(r1), nil
}

func FormatMessage(flags uint32, source uintptr, messageId uint32, languageId uint32, buffer *uint16, size uint32, arguments *byte) (uint32, error) {
	r1, _, e1 := procFormatMessageW.Call(
		uintptr(flags),
		source,
		uintptr(messageId),
		uintptr(languageId),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(size),
		uintptr(unsafe.Pointer(arguments)))
	if r1 == 0 {
		if e1.(syscall.Errno) != 0 {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}
	return uint32(r1), nil
}

func GetComputerNameEx(nameType uint32, buffer *uint16, size *uint32) error {
	r1, _, e1 := procGetComputerNameExW.Call(
		uintptr(nameType),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(size)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func GetCurrentProcess() syscall.Handle {
	r1, _, _ := procGetCurrentProcess.Call()
	return syscall.Handle(r1)
}

func GetDiskFreeSpaceEx(directoryName *uint16, freeBytesAvailable *uint64, totalNumberOfBytes *uint64, totalNumberOfFreeBytes *uint64) error {
	r1, _, e1 := procGetDiskFreeSpaceExW.Call(
		uintptr(unsafe.Pointer(directoryName)),
		uintptr(unsafe.Pointer(freeBytesAvailable)),
		uintptr(unsafe.Pointer(totalNumberOfBytes)),
		uintptr(unsafe.Pointer(totalNumberOfFreeBytes)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func GetDriveType(rootPathName *uint16) uint32 {
	r1, _, _ := procGetDriveTypeW.Call(uintptr(unsafe.Pointer(rootPathName)))
	return uint32(r1)
}

func GetModuleFileName(module syscall.Handle, filename *uint16, size uint32) (uint32, error) {
	r1, _, e1 := procGetModuleFileNameW.Call(
		uintptr(module),
		uintptr(unsafe.Pointer(filename)),
		uintptr(size))
	if r1 == 0 || r1 == uintptr(size) {
		if e1 != ERROR_SUCCESS {
			return uint32(r1), e1
		} else if r1 == uintptr(size) {
			return uint32(r1), ERROR_INSUFFICIENT_BUFFER
		} else {
			return uint32(r1), syscall.EINVAL
		}
	}
	return uint32(r1), nil
}

func GetStdHandle(stdHandle uint32) (syscall.Handle, error) {
	r1, _, e1 := procGetStdHandle.Call(uintptr(stdHandle))
	handle := (syscall.Handle)(r1)
	if handle == INVALID_HANDLE_VALUE {
		if e1 != ERROR_SUCCESS {
			return handle, e1
		} else {
			return handle, syscall.EINVAL
		}
	}
	return handle, nil
}

func GetSystemDirectory(buffer *uint16, size uint32) (uint32, error) {
	r1, _, e1 := procGetSystemDirectoryW.Call(
		uintptr(unsafe.Pointer(buffer)),
		uintptr(size))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return uint32(r1), e1
		} else {
			return uint32(r1), syscall.EINVAL
		}
	}
	return uint32(r1), nil
}

func GetSystemInfo(systemInfo *SYSTEM_INFO) {
	procGetSystemInfo.Call(uintptr(unsafe.Pointer(systemInfo)))
}

func GetSystemTimeAsFileTime(systemTimeAsFileTime *FILETIME) {
	procGetSystemTimeAsFileTime.Call(uintptr(unsafe.Pointer(systemTimeAsFileTime)))
}

func GetSystemTimes(idleTime *int64, kernelTime *int64, userTime *int64) error {
	r1, _, e1 := procGetSystemTimes.Call(
		uintptr(unsafe.Pointer(idleTime)),
		uintptr(unsafe.Pointer(kernelTime)),
		uintptr(unsafe.Pointer(userTime)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func GetVolumeInformation(rootPathName *uint16, volumeNameBuffer *uint16, volumeNameSize uint32, volumeSerialNumber *uint32, maximumComponentLength *uint32, fileSystemFlags *uint32, fileSystemNameBuffer *uint16, fileSystemNameSize uint32) error {
	r1, _, e1 := procGetVolumeInformationW.Call(
		uintptr(unsafe.Pointer(rootPathName)),
		uintptr(unsafe.Pointer(volumeNameBuffer)),
		uintptr(volumeNameSize),
		uintptr(unsafe.Pointer(volumeSerialNumber)),
		uintptr(unsafe.Pointer(maximumComponentLength)),
		uintptr(unsafe.Pointer(fileSystemFlags)),
		uintptr(unsafe.Pointer(fileSystemNameBuffer)),
		uintptr(fileSystemNameSize))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func IsProcessInJob(processHandle syscall.Handle, jobHandle syscall.Handle, result *bool) error {
	var resultRaw int32
	r1, _, e1 := procIsProcessInJob.Call(
		uintptr(processHandle),
		uintptr(jobHandle),
		uintptr(unsafe.Pointer(&resultRaw)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	if result != nil {
		*result = (resultRaw != 0)
	}
	return nil
}

func LocalFree(mem syscall.Handle) (syscall.Handle, error) {
	// LocalFree returns NULL to indicate success!
	r1, _, e1 := procLocalFree.Call(uintptr(mem))
	if r1 != 0 {
		if e1.(syscall.Errno) != 0 {
			return syscall.Handle(r1), e1
		} else {
			return syscall.Handle(r1), syscall.EINVAL
		}
	}
	return 0, nil
}

func OpenJobObject(desiredAccess uint32, inheritHandle bool, name *uint16) (syscall.Handle, error) {
	var inheritHandleRaw int32
	if inheritHandle {
		inheritHandleRaw = 1
	} else {
		inheritHandleRaw = 0
	}
	r1, _, e1 :=  procOpenJobObjectW.Call(
		uintptr(desiredAccess),
		uintptr(inheritHandleRaw),
		uintptr(unsafe.Pointer(name)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}
	return syscall.Handle(r1), nil
}

func OpenProcess(desiredAccess uint32, inheritHandle bool, processId uint32) (syscall.Handle, error) {
	var inheritHandleRaw int32
	if inheritHandle {
		inheritHandleRaw = 1
	} else {
		inheritHandleRaw = 0
	}
	r1, _, e1 := procOpenProcess.Call(
		uintptr(desiredAccess),
		uintptr(inheritHandleRaw),
		uintptr(processId))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}
	return syscall.Handle(r1), nil
}

func QueryFullProcessImageName(process syscall.Handle, flags uint32, exeName *uint16, size *uint32) error {
	r1, _, e1 := procQueryFullProcessImageNameW.Call(
		uintptr(process),
		uintptr(flags),
		uintptr(unsafe.Pointer(exeName)),
		uintptr(unsafe.Pointer(size)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func QueryInformationJobObject(job syscall.Handle, jobObjectInfoClass int32, jobObjectInfo *byte, jobObjectInfoLength uint32, returnLength *uint32) error {
	r1, _, e1 := procQueryInformationJobObject.Call(
		uintptr(job),
		uintptr(jobObjectInfoClass),
		uintptr(unsafe.Pointer(jobObjectInfo)),
		uintptr(jobObjectInfoLength),
		uintptr(unsafe.Pointer(returnLength)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func SetFileTime(file syscall.Handle, creationTime *FILETIME, lastAccessTime *FILETIME, lastWriteTime *FILETIME) error {
	r1, _, e1 := procSetFileTime.Call(
		uintptr(file),
		uintptr(unsafe.Pointer(creationTime)),
		uintptr(unsafe.Pointer(lastAccessTime)),
		uintptr(unsafe.Pointer(lastWriteTime)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func SetInformationJobObject(job syscall.Handle, jobObjectInfoClass int32, jobObjectInfo *byte, jobObjectInfoLength uint32) error {
	r1, _, e1 := procSetInformationJobObject.Call(
		uintptr(job),
		uintptr(jobObjectInfoClass),
		uintptr(unsafe.Pointer(jobObjectInfo)),
		uintptr(jobObjectInfoLength))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func SetStdHandle(stdHandle uint32, handle syscall.Handle) error {
	r1, _, e1 := procSetStdHandle.Call(uintptr(stdHandle), uintptr(handle))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func TerminateJobObject(job syscall.Handle, exitCode uint32) error {
	r1, _, e1 := procTerminateJobObject.Call(uintptr(job), uintptr(exitCode))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func TerminateProcess(process syscall.Handle, exitCode uint32) error {
	r1, _, e1 := procTerminateProcess.Call(uintptr(process), uintptr(exitCode))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func UpdateResource(update syscall.Handle, resourceType *uint16, name *uint16, language uint16, data *byte, cbData uint32) error {
	r1, _, e1 := procUpdateResourceW.Call(
		uintptr(update),
		uintptr(unsafe.Pointer(resourceType)),
		uintptr(unsafe.Pointer(name)),
		uintptr(language),
		uintptr(unsafe.Pointer(data)),
		uintptr(cbData))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func VerifyVersionInfo(versionInfo *OSVERSIONINFOEX, typeMask uint32, conditionMask uint64) error {
	r1, _, e1 := procVerifyVersionInfoW.Call(
		uintptr(unsafe.Pointer(versionInfo)),
		uintptr(typeMask),
		uintptr(conditionMask))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func WaitForSingleObject(handle syscall.Handle, milliseconds uint32) (uint32, error) {
	r1, _, e1 := procWaitForSingleObject.Call(uintptr(handle), uintptr(milliseconds))
	if r1 == WAIT_FAILED {
		if e1 != ERROR_SUCCESS {
			return uint32(r1), e1
		} else {
			return uint32(r1), syscall.EINVAL
		}
	}
	return uint32(r1), nil
}

func Lstrlen(string *uint16) int32 {
	r1, _, _ := proclstrlenW.Call(uintptr(unsafe.Pointer(string)))
	return int32(r1)
}

func AllocateAndInitializeSid(identifierAuthority *SID_IDENTIFIER_AUTHORITY, subAuthorityCount byte, subAuthority0 uint32, subAuthority1 uint32, subAuthority2 uint32, subAuthority3 uint32, subAuthority4 uint32, subAuthority5 uint32, subAuthority6 uint32, subAuthority7 uint32, sid **SID) error {
	r1, _, e1 := procAllocateAndInitializeSid.Call(
		uintptr(unsafe.Pointer(identifierAuthority)),
		uintptr(subAuthorityCount),
		uintptr(subAuthority0),
		uintptr(subAuthority1),
		uintptr(subAuthority2),
		uintptr(subAuthority3),
		uintptr(subAuthority4),
		uintptr(subAuthority5),
		uintptr(subAuthority6),
		uintptr(subAuthority7),
		uintptr(unsafe.Pointer(sid)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func CheckTokenMembership(tokenHandle syscall.Handle, sidToCheck *SID, isMember *bool) error {
	var isMemberRaw int32
	r1, _, e1 := procCheckTokenMembership.Call(
		uintptr(tokenHandle),
		uintptr(unsafe.Pointer(sidToCheck)),
		uintptr(unsafe.Pointer(&isMemberRaw)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	if isMember != nil {
		*isMember = (isMemberRaw != 0)
	}
	return nil
}

func CopySid(destinationSidLength uint32, destinationSid *SID, sourceSid *SID) error {
	r1, _, e1 := procCopySid.Call(
		uintptr(destinationSidLength),
		uintptr(unsafe.Pointer(destinationSid)),
		uintptr(unsafe.Pointer(sourceSid)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func DeregisterEventSource(eventLog syscall.Handle) error {
	r1, _, e1 := procDeregisterEventSource.Call(uintptr(eventLog))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func EqualSid(sid1 *SID, sid2 *SID) bool {
	r1, _, _ := procEqualSid.Call(
		uintptr(unsafe.Pointer(sid1)),
		uintptr(unsafe.Pointer(sid2)))
	return r1 != 0
}

func FreeSid(sid *SID) {
	procFreeSid.Call(uintptr(unsafe.Pointer(sid)))
}

func GetFileSecurity(fileName *uint16, requestedInformation uint32, securityDescriptor *uint8, length uint32, lengthNeeded *uint32) error {
	r1, _, e1 := procGetFileSecurityW.Call(
		uintptr(unsafe.Pointer(fileName)),
		uintptr(requestedInformation),
		uintptr(unsafe.Pointer(securityDescriptor)),
		uintptr(length),
		uintptr(unsafe.Pointer(lengthNeeded)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func GetLengthSid(sid *SID) uint32 {
	r1, _, _ := procGetLengthSid.Call(uintptr(unsafe.Pointer(sid)))
	return uint32(r1)
}

func GetSecurityDescriptorOwner(securityDescriptor *uint8, owner **SID, ownerDefaulted *bool) error {
	var ownerDefaultedRaw int32
	r1, _, e1 := procGetSecurityDescriptorOwner.Call(
		uintptr(unsafe.Pointer(securityDescriptor)),
		uintptr(unsafe.Pointer(owner)),
		uintptr(unsafe.Pointer(&ownerDefaultedRaw)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	if ownerDefaulted != nil {
		*ownerDefaulted = (ownerDefaultedRaw != 0)
	}
	return nil
}

func GetTokenInformation(tokenHandle syscall.Handle, tokenInformationClass int32, tokenInformation *byte, tokenInformationLength uint32, returnLength *uint32) error {
	r1, _, e1 := procGetTokenInformation.Call(
		uintptr(tokenHandle),
		uintptr(tokenInformationClass),
		uintptr(unsafe.Pointer(tokenInformation)),
		uintptr(tokenInformationLength),
		uintptr(unsafe.Pointer(returnLength)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func OpenProcessToken(processHandle syscall.Handle, desiredAccess uint32, tokenHandle *syscall.Handle) error {
	r1, _, e1 := procOpenProcessToken.Call(
		uintptr(processHandle),
		uintptr(desiredAccess),
		uintptr(unsafe.Pointer(tokenHandle)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func RegisterEventSource(uncServerName *uint16, sourceName *uint16) (syscall.Handle, error) {
	r1, _, e1 := procRegisterEventSourceW.Call(
		uintptr(unsafe.Pointer(uncServerName)),
		uintptr(unsafe.Pointer(sourceName)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}
	return syscall.Handle(r1), nil
}

func ReportEvent(eventLog syscall.Handle, eventType uint16, category uint16, eventID uint32, userSid *SID, numStrings uint16, dataSize uint32, strings **uint16, rawData *byte) error {
	r1, _, e1 := procReportEventW.Call(
		uintptr(eventLog),
		uintptr(eventType),
		uintptr(category),
		uintptr(eventID),
		uintptr(unsafe.Pointer(userSid)),
		uintptr(numStrings),
		uintptr(dataSize),
		uintptr(unsafe.Pointer(strings)),
		uintptr(unsafe.Pointer(rawData)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}
