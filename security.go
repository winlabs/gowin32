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

package gowin32

import (
	"github.com/winlabs/gowin32/wrappers"

	"syscall"
	"unsafe"
)

type SecurityID struct {
	sid *wrappers.SID
}

func (self SecurityID) GetLength() uint {
	return uint(wrappers.GetLengthSid(self.sid))
}

func (self SecurityID) Copy() (SecurityID, error) {
	length := self.GetLength()
	buf := make([]byte, length)
	sid := (*wrappers.SID)(unsafe.Pointer(&buf[0]))
	err := wrappers.CopySid(uint32(length), sid, self.sid)
	if err != nil {
		return SecurityID{}, NewWindowsError("CopySid", err)
	}
	return SecurityID{sid}, nil
}

func (self SecurityID) Equal(other SecurityID) bool {
	return wrappers.EqualSid(self.sid, other.sid)
}

func GetFileOwner(path string) (SecurityID, error) {
	var needed uint32
	wrappers.GetFileSecurity(
		syscall.StringToUTF16Ptr(path),
		wrappers.OWNER_SECURITY_INFORMATION,
		nil,
		0,
		&needed)
	buf := make([]byte, needed)
	err := wrappers.GetFileSecurity(
		syscall.StringToUTF16Ptr(path),
		wrappers.OWNER_SECURITY_INFORMATION,
		&buf[0],
		needed,
		&needed)
	if err != nil {
		return SecurityID{}, NewWindowsError("GetFileSecurity", err)
	}
	var ownerSid *wrappers.SID
	if err := wrappers.GetSecurityDescriptorOwner(&buf[0], &ownerSid, nil); err != nil {
		return SecurityID{}, NewWindowsError("GetSecurityDescriptorOwner", err)
	}
	return SecurityID{ownerSid}, nil
}

const (
	PrivilegeCreateToken                    = wrappers.SE_CREATE_TOKEN_NAME
	PrivilegeAssignPrimaryToken             = wrappers.SE_ASSIGNPRIMARYTOKEN_NAME
	PrivilegeLockMemory                     = wrappers.SE_LOCK_MEMORY_NAME
	PrivilegeIncreaseQuota                  = wrappers.SE_INCREASE_QUOTA_NAME
	PrivilegeUnsolicitedInput               = wrappers.SE_UNSOLICITED_INPUT_NAME
	PrivilegeMachineAccount                 = wrappers.SE_MACHINE_ACCOUNT_NAME
	PrivilegeTrustedComputerBase            = wrappers.SE_TCB_NAME
	PrivilegeSecurity                       = wrappers.SE_SECURITY_NAME
	PrivilegeTakeOwnership                  = wrappers.SE_TAKE_OWNERSHIP_NAME
	PrivilegeLoadDriver                     = wrappers.SE_LOAD_DRIVER_NAME
	PrivilegeSystemProfile                  = wrappers.SE_SYSTEM_PROFILE_NAME
	PrivilegeSystemTime                     = wrappers.SE_SYSTEMTIME_NAME
	PrivilegeProfileSingleProcess           = wrappers.SE_PROF_SINGLE_PROCESS_NAME
	PrivilegeIncreaseBasePriority           = wrappers.SE_INC_BASE_PRIORITY_NAME
	PrivilegeCreatePageFile                 = wrappers.SE_CREATE_PAGEFILE_NAME
	PrivilegeCreatePermanent                = wrappers.SE_CREATE_PERMANENT_NAME
	PrivilegeBackup                         = wrappers.SE_BACKUP_NAME
	PrivilegeRestore                        = wrappers.SE_RESTORE_NAME
	PrivilegeShutdown                       = wrappers.SE_SHUTDOWN_NAME
	PrivilegeDebug                          = wrappers.SE_DEBUG_NAME
	PrivilegeAudit                          = wrappers.SE_AUDIT_NAME
	PrivilegeSystemEnvironment              = wrappers.SE_SYSTEM_ENVIRONMENT_NAME
	PrivilegeChangeNotify                   = wrappers.SE_CHANGE_NOTIFY_NAME
	PrivilegeRemoteShutdown                 = wrappers.SE_REMOTE_SHUTDOWN_NAME
	PrivilegeUndock                         = wrappers.SE_UNDOCK_NAME
	PrivilegeSyncAgent                      = wrappers.SE_SYNC_AGENT_NAME
	PrivilegeEnableDelegation               = wrappers.SE_ENABLE_DELEGATION_NAME
	PrivilegeManageVolume                   = wrappers.SE_MANAGE_VOLUME_NAME
	PrivilegeImpersonate                    = wrappers.SE_IMPERSONATE_NAME
	PrivilegeCreateGlobal                   = wrappers.SE_CREATE_GLOBAL_NAME
	PrivilegeTrustedCredentialManagerAccess = wrappers.SE_TRUSTED_CREDMAN_ACCESS_NAME
	PrivilegeRelabel                        = wrappers.SE_RELABEL_NAME
	PrivilegeIncreaseWorkingSet             = wrappers.SE_INC_WORKING_SET_NAME
	PrivilegeTimeZone                       = wrappers.SE_TIME_ZONE_NAME
	PrivilegeCreateSymbolicLink             = wrappers.SE_CREATE_SYMBOLIC_LINK_NAME
)

type Privilege struct {
	luid wrappers.LUID
}

func GetPrivilegeByName(name string) (*Privilege, error) {
	var luid wrappers.LUID
	if err := wrappers.LookupPrivilegeValue(nil, syscall.StringToUTF16Ptr(name), &luid); err != nil {
		return nil, NewWindowsError("LookupPrivilegeValue", err)
	}
	return &Privilege{luid: luid}, nil
}

func (privilege *Privilege) EnableForCurrentProcess() error {
	hProcess := wrappers.GetCurrentProcess()
	var hToken syscall.Handle
	if err := wrappers.OpenProcessToken(hProcess, wrappers.TOKEN_ADJUST_PRIVILEGES, &hToken); err != nil {
		return NewWindowsError("OpenProcessToken", err)
	}
	defer wrappers.CloseHandle(hToken)
	state := wrappers.TOKEN_PRIVILEGES{
		PrivilegeCount: 1,
		Privileges:     [1]wrappers.LUID_AND_ATTRIBUTES{
			wrappers.LUID_AND_ATTRIBUTES{
				Luid:       privilege.luid,
				Attributes: wrappers.SE_PRIVILEGE_ENABLED,
			},
		},
	}
	if err := wrappers.AdjustTokenPrivileges(hToken, false, &state, 0, nil, nil); err != nil {
		return NewWindowsError("AdjustTokenPrivileges", err)
	}
	return nil
}

type Token struct {
	handle syscall.Handle
}

func OpenCurrentProcessToken() (*Token, error) {
	hProcess := wrappers.GetCurrentProcess()
	var hToken syscall.Handle
	if err := wrappers.OpenProcessToken(hProcess, wrappers.TOKEN_QUERY, &hToken); err != nil {
		return nil, NewWindowsError("OpenProcessToken", err)
	}
	return &Token{handle: hToken}, nil
}

func (self *Token) Close() error {
	if self.handle != 0 {
		if err := wrappers.CloseHandle(self.handle); err != nil {
			return NewWindowsError("CloseHandle", err)
		}
		self.handle = 0
	}
	return nil
}

func (self *Token) GetOwner() (SecurityID, error) {
	var needed uint32
	wrappers.GetTokenInformation(
		self.handle,
		wrappers.TokenOwner,
		nil,
		0,
		&needed)
	buf := make([]byte, needed)
	err := wrappers.GetTokenInformation(
		self.handle,
		wrappers.TokenOwner,
		&buf[0],
		needed,
		&needed)
	if err != nil {
		return SecurityID{}, NewWindowsError("GetTokenInformation", err)
	}
	ownerData := (*wrappers.TOKEN_OWNER)(unsafe.Pointer(&buf[0]))
	sid, err := SecurityID{ownerData.Owner}.Copy()
	if err != nil {
		return SecurityID{}, err
	}
	return sid, nil
}

func IsAdmin() (bool, error) {
	var sid *wrappers.SID
	err := wrappers.AllocateAndInitializeSid(
		&wrappers.SECURITY_NT_AUTHORITY,
		2,
		wrappers.SECURITY_BUILTIN_DOMAIN_RID,
		wrappers.DOMAIN_ALIAS_RID_ADMINS,
		0,
		0,
		0,
		0,
		0,
		0,
		&sid)
	if err != nil {
		return false, NewWindowsError("AllocateAndInitializeSid", err)
	}
	defer wrappers.FreeSid(sid)
	var isAdmin bool
	if err := wrappers.CheckTokenMembership(0, sid, &isAdmin); err != nil {
		return false, NewWindowsError("CheckTokenMembership", err)
	}
	return isAdmin, nil
}
