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

func GetFileOwner(path string) (*syscall.SID, error) {
	needed := uint32(0)
	wrappers.GetFileSecurity(
		syscall.StringToUTF16Ptr(path),
		wrappers.OWNER_SECURITY_INFORMATION,
		nil,
		0,
		&needed)
	buf := make([]uint8, needed)
	err := wrappers.GetFileSecurity(
		syscall.StringToUTF16Ptr(path),
		wrappers.OWNER_SECURITY_INFORMATION,
		&buf[0],
		needed,
		&needed)
	if err != nil {
		return nil, err
	}
	ownerSid := (*syscall.SID)(nil)
	if err := wrappers.GetSecurityDescriptorOwner(&buf[0], &ownerSid, nil); err != nil {
		return nil, err
	}
	return ownerSid, nil
}

func GetTokenOwner(token syscall.Token) (*syscall.SID, error) {
	needed := uint32(0)
	syscall.GetTokenInformation(
		token,
		syscall.TokenOwner,
		nil,
		0,
		&needed)
	buf := make([]byte, needed)
	err := syscall.GetTokenInformation(
		token,
		syscall.TokenOwner,
		&buf[0],
		needed,
		&needed)
	if err != nil {
		return nil, err
	}
	ownerData := (*wrappers.TOKEN_OWNER)(unsafe.Pointer(&buf[0]))
	sid, err := ownerData.Owner.Copy()
	if err != nil {
		return nil, err
	}
	return sid, nil
}

func IsAdmin() (bool, error) {
	var sid *syscall.SID
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
		return false, err
	}
	defer wrappers.FreeSid(sid)
	var isAdmin bool
	if err := wrappers.CheckTokenMembership(0, sid, &isAdmin); err != nil {
		return false, err
	}
	return isAdmin, nil
}
