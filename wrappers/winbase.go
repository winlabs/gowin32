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
	INVALID_HANDLE_VALUE    = ^syscall.Handle(0)
	INVALID_FILE_SIZE       = 0xFFFFFFFF
	INVALID_FILE_ATTRIBUTES = 0xFFFFFFFF
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

type OVERLAPPED struct {
	Internal     uintptr
	InternalHigh uintptr
	Offset       uint32
	OffsetHigh   uint32
	Event        syscall.Handle
}

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

type CRITICAL_SECTION RTL_CRITICAL_SECTION

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

type WIN32_FIND_DATA struct {
	FileAttributes    uint32
	CreationTime      FILETIME
	LastAccessTime    FILETIME
	LastWriteTime     FILETIME
	FileSizeHigh      uint32
	FileSizeLow       uint32
	Reserved0         uint32
	Reserved1         uint32
	FileName          [MAX_PATH]uint16
	AlternateFileName [14]uint16
}

const (
	PROCESS_NAME_NATIVE = 0x00000001
)

const (
	MOVEFILE_REPLACE_EXISTING      = 0x00000001
	MOVEFILE_COPY_ALLOWED          = 0x00000002
	MOVEFILE_DELAY_UNTIL_REBOOT    = 0x00000004
	MOVEFILE_WRITE_THROUGH         = 0x00000008
	MOVEFILE_CREATE_HARDLINK       = 0x00000010
	MOVEFILE_FAIL_IF_NOT_TRACKABLE = 0x00000020
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

const (
	SYMBOLIC_LINK_FLAG_DIRECTORY = 0x00000001
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	modadvapi32 = syscall.NewLazyDLL("advapi32.dll")

	procAssignProcessToJobObject          = modkernel32.NewProc("AssignProcessToJobObject")
	procBeginUpdateResourceW              = modkernel32.NewProc("BeginUpdateResourceW")
	procCloseHandle                       = modkernel32.NewProc("CloseHandle")
	procCopyFileW                         = modkernel32.NewProc("CopyFileW")
	procCreateFileW                       = modkernel32.NewProc("CreateFileW")
	procCreateJobObjectW                  = modkernel32.NewProc("CreateJobObjectW")
	procCreateProcessW                    = modkernel32.NewProc("CreateProcessW")
	procCreateSymbolicLinkW               = modkernel32.NewProc("CreateSymbolicLinkW")
	procDeleteCriticalSection             = modkernel32.NewProc("DeleteCriticalSection")
	procDeleteFileW                       = modkernel32.NewProc("DeleteFileW")
	procDeviceIoControl                   = modkernel32.NewProc("DeviceIoControl")
	procEndUpdateResourceW                = modkernel32.NewProc("EndUpdateResourceW")
	procEnterCriticalSection              = modkernel32.NewProc("EnterCriticalSection")
	procExpandEnvironmentStringsW         = modkernel32.NewProc("ExpandEnvironmentStringsW")
	procFindClose                         = modkernel32.NewProc("FindClose")
	procFindFirstFileW                    = modkernel32.NewProc("FindFirstFileW")
	procFindNextFileW                     = modkernel32.NewProc("FindNextFileW")
	procFormatMessageW                    = modkernel32.NewProc("FormatMessageW")
	procFreeEnvironmentStringsW           = modkernel32.NewProc("FreeEnvironmentStringsW")
	procGetCompressedFileSizeW            = modkernel32.NewProc("GetCompressedFileSizeW")
	procGetComputerNameExW                = modkernel32.NewProc("GetComputerNameExW")
	procGetCurrentProcess                 = modkernel32.NewProc("GetCurrentProcess")
	procGetDriveTypeW                     = modkernel32.NewProc("GetDriveTypeW")
	procGetDiskFreeSpaceExW               = modkernel32.NewProc("GetDiskFreeSpaceExW")
	procGetDiskFreeSpaceW                 = modkernel32.NewProc("GetDiskFreeSpaceW")
	procGetEnvironmentStringsW            = modkernel32.NewProc("GetEnvironmentStringsW")
	procGetEnvironmentVariableW           = modkernel32.NewProc("GetEnvironmentVariableW")
	procGetFileAttributesW                = modkernel32.NewProc("GetFileAttributesW")
	procGetFileSize                       = modkernel32.NewProc("GetFileSize")
	procGetModuleFileNameW                = modkernel32.NewProc("GetModuleFileNameW")
	procGetProcessTimes                   = modkernel32.NewProc("GetProcessTimes")
	procGetStdHandle                      = modkernel32.NewProc("GetStdHandle")
	procGetSystemDirectoryW               = modkernel32.NewProc("GetSystemDirectoryW")
	procGetSystemInfo                     = modkernel32.NewProc("GetSystemInfo")
	procGetSystemTimeAsFileTime           = modkernel32.NewProc("GetSystemTimeAsFileTime")
	procGetSystemTimes                    = modkernel32.NewProc("GetSystemTimes")
	procGetSystemWindowsDirectoryW        = modkernel32.NewProc("GetSystemWindowsDirectoryW")
	procGetSystemWow64DirectoryW          = modkernel32.NewProc("GetSystemWow64DirectoryW")
	procGetTempFileNameW                  = modkernel32.NewProc("GetTempFileNameW")
	procGetTempPathW                      = modkernel32.NewProc("GetTempPathW")
	procGetVersionExW                     = modkernel32.NewProc("GetVersionExW")
	procGetVolumeInformationW             = modkernel32.NewProc("GetVolumeInformationW")
	procGetVolumeNameForVolumeMountPointW = modkernel32.NewProc("GetVolumeNameForVolumeMountPointW")
	procGetVolumePathNameW                = modkernel32.NewProc("GetVolumePathNameW")
	procGetWindowsDirectoryW              = modkernel32.NewProc("GetWindowsDirectoryW")
	procInitializeCriticalSection         = modkernel32.NewProc("InitializeCriticalSection")
	procIsProcessInJob                    = modkernel32.NewProc("IsProcessInJob")
	procLeaveCriticalSection              = modkernel32.NewProc("LeaveCriticalSection")
	procLocalFree                         = modkernel32.NewProc("LocalFree")
	procMoveFileExW                       = modkernel32.NewProc("MoveFileExW")
	procMoveFileW                         = modkernel32.NewProc("MoveFileW")
	procOpenJobObjectW                    = modkernel32.NewProc("OpenJobObjectW")
	procOpenProcess                       = modkernel32.NewProc("OpenProcess")
	procQueryFullProcessImageNameW        = modkernel32.NewProc("QueryFullProcessImageNameW")
	procQueryInformationJobObject         = modkernel32.NewProc("QueryInformationJobObject")
	procReadFile                          = modkernel32.NewProc("ReadFile")
	procSetEnvironmentVariableW           = modkernel32.NewProc("SetEnvironmentVariableW")
	procSetFileAttributesW                = modkernel32.NewProc("SetFileAttributesW")
	procSetFileTime                       = modkernel32.NewProc("SetFileTime")
	procSetInformationJobObject           = modkernel32.NewProc("SetInformationJobObject")
	procSetStdHandle                      = modkernel32.NewProc("SetStdHandle")
	procTerminateJobObject                = modkernel32.NewProc("TerminateJobObject")
	procTerminateProcess                  = modkernel32.NewProc("TerminateProcess")
	procTryEnterCriticalSection           = modkernel32.NewProc("TryEnterCriticalSection")
	procUpdateResourceW                   = modkernel32.NewProc("UpdateResourceW")
	procVerifyVersionInfoW                = modkernel32.NewProc("VerifyVersionInfoW")
	procWaitForSingleObject               = modkernel32.NewProc("WaitForSingleObject")
	proclstrlenW                          = modkernel32.NewProc("lstrlenW")

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

func CopyFile(existingFileName *uint16, newFileName *uint16, failIfExists bool) error {
	var failIfExistsRaw int32
	if failIfExists {
		failIfExistsRaw = 1
	} else {
		failIfExistsRaw = 0
	}
	r1, _, e1 := procCopyFileW.Call(
		uintptr(unsafe.Pointer(existingFileName)),
		uintptr(unsafe.Pointer(newFileName)),
		uintptr(failIfExistsRaw))
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

func CreateJobObject(jobAttributes *SECURITY_ATTRIBUTES, name *uint16) (syscall.Handle, error) {
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

func CreateSymbolicLink(symlinkFileName *uint16, targetFileName *uint16, flags uint32) error {
	r1, _, e1 := procCreateSymbolicLinkW.Call(
		uintptr(unsafe.Pointer(symlinkFileName)),
		uintptr(unsafe.Pointer(targetFileName)),
		uintptr(flags))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func DeleteCriticalSection(criticalSection *CRITICAL_SECTION) {
	procDeleteCriticalSection.Call(uintptr(unsafe.Pointer(criticalSection)))
}

func DeleteFile(fileName *uint16) error {
	r1, _, e1 := procDeleteFileW.Call(uintptr(unsafe.Pointer(fileName)))
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

func EnterCriticalSection(criticalSection *CRITICAL_SECTION) {
	procEnterCriticalSection.Call(uintptr(unsafe.Pointer(criticalSection)))
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

func FindClose(findFile syscall.Handle) error {
	r1, _, e1 := procFindClose.Call(uintptr(findFile))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func FindFirstFile(fileName *uint16, findFileData *WIN32_FIND_DATA) (syscall.Handle, error) {
	r1, _, e1 := procFindFirstFileW.Call(
		uintptr(unsafe.Pointer(fileName)),
		uintptr(unsafe.Pointer(findFileData)))
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

func FindNextFile(findFile syscall.Handle, findFileData *WIN32_FIND_DATA) error {
	r1, _, e1 := procFindNextFileW.Call(uintptr(findFile), uintptr(unsafe.Pointer(findFileData)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
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

func FreeEnvironmentStrings(environmentBlock *uint16) error {
	r1, _, e1 := procFreeEnvironmentStringsW.Call(uintptr(unsafe.Pointer(environmentBlock)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func GetCompressedFileSize(fileName *uint16, fileSizeHigh *uint32) (uint32, error) {
	r1, _, e1 := procGetCompressedFileSizeW.Call(
		uintptr(unsafe.Pointer(fileName)),
		uintptr(unsafe.Pointer(fileSizeHigh)))
	if r1 == INVALID_FILE_SIZE {
		if e1 != ERROR_SUCCESS {
			return uint32(r1), e1
		} else {
			return uint32(r1), syscall.EINVAL
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

func GetDiskFreeSpace(rootPathName *uint16, sectorsPerCluster *uint32, bytesPerSector *uint32, numberOfFreeClusters *uint32, totalNumberOfClusters *uint32) error {
	r1, _, e1 := procGetDiskFreeSpaceW.Call(
		uintptr(unsafe.Pointer(rootPathName)),
		uintptr(unsafe.Pointer(sectorsPerCluster)),
		uintptr(unsafe.Pointer(bytesPerSector)),
		uintptr(unsafe.Pointer(numberOfFreeClusters)),
		uintptr(unsafe.Pointer(totalNumberOfClusters)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
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

func GetEnvironmentStrings() (*uint16, error) {
	r1, _, e1 := procGetEnvironmentStringsW.Call()
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return nil, e1
		} else {
			return nil, syscall.EINVAL
		}
	}
	return (*uint16)(unsafe.Pointer(r1)), nil
}

func GetEnvironmentVariable(name *uint16, buffer *uint16, size uint32) (uint32, error) {
	r1, _, e1 := procGetEnvironmentVariableW.Call(
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(buffer)),
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

func GetFileAttributes(fileName *uint16) (uint32, error) {
	r1, _, e1 := procGetFileAttributesW.Call(uintptr(unsafe.Pointer(fileName)))
	if r1 == INVALID_FILE_ATTRIBUTES {
		if e1 != ERROR_SUCCESS {
			return uint32(r1), e1
		} else {
			return uint32(r1), syscall.EINVAL
		}
	}
	return uint32(r1), nil
}

func GetFileSize(file syscall.Handle, fileSizeHigh *uint32) (uint32, error) {
	r1, _, e1 := procGetFileSize.Call(
		uintptr(file),
		uintptr(unsafe.Pointer(fileSizeHigh)))
	if r1 == INVALID_FILE_SIZE {
		if e1 != ERROR_SUCCESS {
			return uint32(r1), e1
		} else {
			return uint32(r1), syscall.EINVAL
		}
	}
	return uint32(r1), nil
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

func GetProcessTimes(hProcess syscall.Handle, creationTime, exitTime, kernelTime, userTime *FILETIME) error {
	r1, _, e1 := procGetProcessTimes.Call(
		uintptr(hProcess),
		uintptr(unsafe.Pointer(creationTime)),
		uintptr(unsafe.Pointer(exitTime)),
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
			return 0, e1
		} else {
			return 0, syscall.EINVAL
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

func GetSystemTimes(idleTime, kernelTime, userTime *FILETIME) error {
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

func GetSystemWindowsDirectory(buffer *uint16, size uint32) (uint32, error) {
	r1, _, e1 := procGetSystemWindowsDirectoryW.Call(
		uintptr(unsafe.Pointer(buffer)),
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

func GetSystemWow64Directory(buffer *uint16, size uint32) (uint32, error) {
	r1, _, e1 := procGetSystemWow64DirectoryW.Call(
		uintptr(unsafe.Pointer(buffer)),
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

func GetTempFileName(pathName *uint16, prefixString *uint16, unique uint32, tempFileName *uint16) (uint32, error) {
	r1, _, e1 := procGetTempFileNameW.Call(
		uintptr(unsafe.Pointer(pathName)),
		uintptr(unsafe.Pointer(prefixString)),
		uintptr(unique),
		uintptr(unsafe.Pointer(tempFileName)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}
	return uint32(r1), nil
}

func GetTempPath(bufferLength uint32, buffer *uint16) (uint32, error) {
	r1, _, e1 := procGetTempPathW.Call(
		uintptr(bufferLength),
		uintptr(unsafe.Pointer(buffer)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return 0, e1
		} else {
			return 0, syscall.EINVAL
		}
	}
	return uint32(r1), nil
}

func GetVersionEx(osvi *OSVERSIONINFOEX) error {
	r1, _, e1 := procGetVersionExW.Call(uintptr(unsafe.Pointer(osvi)))
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

func GetVolumeNameForVolumeMountPoint(volumeMountPoint *uint16, volumeName *uint16, bufferLength uint32) error {
	r1, _, e1 := procGetVolumeNameForVolumeMountPointW.Call(
		uintptr(unsafe.Pointer(volumeMountPoint)),
		uintptr(unsafe.Pointer(volumeName)),
		uintptr(bufferLength))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func GetVolumePathName(fileName *uint16, volumePathName *uint16, bufferLength uint32) error {
	r1, _, e1 := procGetVolumePathNameW.Call(
		uintptr(unsafe.Pointer(fileName)),
		uintptr(unsafe.Pointer(volumePathName)),
		uintptr(bufferLength))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func GetWindowsDirectory(buffer *uint16, size uint32) (uint32, error) {
	r1, _, e1 := procGetWindowsDirectoryW.Call(
		uintptr(unsafe.Pointer(buffer)),
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

func InitializeCriticalSection(criticalSection *CRITICAL_SECTION) {
	procInitializeCriticalSection.Call(uintptr(unsafe.Pointer(criticalSection)))
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

func LeaveCriticalSection(criticalSection *CRITICAL_SECTION) {
	procLeaveCriticalSection.Call(uintptr(unsafe.Pointer(criticalSection)))
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

func MoveFile(existingFileName *uint16, newFileName *uint16) error {
	r1, _, e1 := procMoveFileW.Call(
		uintptr(unsafe.Pointer(existingFileName)),
		uintptr(unsafe.Pointer(newFileName)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func MoveFileEx(existingFileName *uint16, newFileName *uint16, flags uint32) error {
	r1, _, e1 := procMoveFileExW.Call(
		uintptr(unsafe.Pointer(existingFileName)),
		uintptr(unsafe.Pointer(newFileName)),
		uintptr(flags))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func OpenJobObject(desiredAccess uint32, inheritHandle bool, name *uint16) (syscall.Handle, error) {
	var inheritHandleRaw int32
	if inheritHandle {
		inheritHandleRaw = 1
	} else {
		inheritHandleRaw = 0
	}
	r1, _, e1 := procOpenJobObjectW.Call(
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

func ReadFile(file syscall.Handle, buffer *byte, numberOfBytesToRead uint32, numberOfBytesRead *uint32, overlapped *OVERLAPPED) error {
	r1, _, e1 := procReadFile.Call(
		uintptr(file),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(numberOfBytesToRead),
		uintptr(unsafe.Pointer(numberOfBytesRead)),
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

func SetEnvironmentVariable(name *uint16, value *uint16) error {
	r1, _, e1 := procSetEnvironmentVariableW.Call(
		uintptr(unsafe.Pointer(name)),
		uintptr(unsafe.Pointer(value)))
	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func SetFileAttributes(fileName *uint16, fileAttributes uint32) error {
	r1, _, e1 := procSetFileAttributesW.Call(
		uintptr(unsafe.Pointer(fileName)),
		uintptr(fileAttributes))
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

func TryEnterCriticalSection(criticalSection *CRITICAL_SECTION) bool {
	r1, _, _ := procTryEnterCriticalSection.Call(uintptr(unsafe.Pointer(criticalSection)))
	return r1 != 0
}

func UpdateResource(update syscall.Handle, resourceType uintptr, name uintptr, language uint16, data *byte, cbData uint32) error {
	r1, _, e1 := procUpdateResourceW.Call(
		uintptr(update),
		resourceType,
		name,
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

func GetFileSecurity(fileName *uint16, requestedInformation uint32, securityDescriptor *byte, length uint32, lengthNeeded *uint32) error {
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

func GetSecurityDescriptorOwner(securityDescriptor *byte, owner **SID, ownerDefaulted *bool) error {
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
