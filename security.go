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

func (self SecurityID) String() (string, error) {
	var stringSid *uint16
	if err := wrappers.ConvertSidToStringSid(self.sid, &stringSid); err != nil {
		return "", NewWindowsError("ConvertSidToStringSid", err)
	}
	defer wrappers.LocalFree(syscall.Handle(unsafe.Pointer(stringSid)))
	return LpstrToString(stringSid), nil
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

func OpenOtherProcessToken(pid uint) (*Token, error) {
	hProcess, err := wrappers.OpenProcess(wrappers.PROCESS_QUERY_INFORMATION, false, uint32(pid))
	if err != nil {
		return nil, NewWindowsError("OpenProcess", err)
	}
	defer syscall.CloseHandle(hProcess)
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
