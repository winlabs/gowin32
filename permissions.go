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
	"github.com/winlabs/gowin32/wrappers"

	"syscall"
	"unsafe"
)

type Permission uint32

const (
	GenericRead            Permission = wrappers.GENERIC_READ
	GenericWrite           Permission = wrappers.GENERIC_WRITE
	GenericExecute         Permission = wrappers.GENERIC_EXECUTE
	GenericAll             Permission = wrappers.GENERIC_ALL
	FileReadData           Permission = wrappers.FILE_READ_DATA
	FileListDirectory      Permission = wrappers.FILE_LIST_DIRECTORY
	FileWriteData          Permission = wrappers.FILE_WRITE_DATA
	FileAddFile            Permission = wrappers.FILE_ADD_FILE
	FileAppendData         Permission = wrappers.FILE_APPEND_DATA
	FileAddSubdirectory    Permission = wrappers.FILE_ADD_SUBDIRECTORY
	FileCreatePipeInstance Permission = wrappers.FILE_CREATE_PIPE_INSTANCE
	FileReadEA             Permission = wrappers.FILE_READ_EA
	FileWriteEA            Permission = wrappers.FILE_WRITE_EA
	FileExecute            Permission = wrappers.FILE_EXECUTE
	FileTraverse           Permission = wrappers.FILE_TRAVERSE
	FileDeleteChild        Permission = wrappers.FILE_DELETE_CHILD
	FileReadAttributes     Permission = wrappers.FILE_READ_ATTRIBUTES
	FileWriteAttributes    Permission = wrappers.FILE_WRITE_ATTRIBUTES
	FileAllAccess          Permission = wrappers.FILE_ALL_ACCESS
	FileGenericRead        Permission = wrappers.FILE_GENERIC_READ
	FileGenericWrite       Permission = wrappers.FILE_GENERIC_WRITE
	FileGenericExecute     Permission = wrappers.FILE_GENERIC_EXECUTE
)

type AccessMode int32

const (
	AccessNotUsed      AccessMode = wrappers.NOT_USED_ACCESS
	AccessGrant        AccessMode = wrappers.GRANT_ACCESS
	AccessSet          AccessMode = wrappers.SET_ACCESS
	AccessDeny         AccessMode = wrappers.DENY_ACCESS
	AccessRevoke       AccessMode = wrappers.REVOKE_ACCESS
	AccessAuditSuccess AccessMode = wrappers.SET_AUDIT_SUCCESS
	AccessAuditFailure AccessMode = wrappers.SET_AUDIT_FAILURE
)

type TrusteeType int32

const (
	TrusteeUnknown        TrusteeType = wrappers.TRUSTEE_IS_UNKNOWN
	TrusteeUser           TrusteeType = wrappers.TRUSTEE_IS_USER
	TrusteeGroup          TrusteeType = wrappers.TRUSTEE_IS_GROUP
	TrusteeDomain         TrusteeType = wrappers.TRUSTEE_IS_DOMAIN
	TrusteeAlias          TrusteeType = wrappers.TRUSTEE_IS_ALIAS
	TrusteeWellKnownGroup TrusteeType = wrappers.TRUSTEE_IS_WELL_KNOWN_GROUP
	TrusteeDeleted        TrusteeType = wrappers.TRUSTEE_IS_DELETED
	TrusteeInvalid        TrusteeType = wrappers.TRUSTEE_IS_INVALID
	TrusteeComputer       TrusteeType = wrappers.TRUSTEE_IS_COMPUTER
)

type PermissionEntry struct {
	TrusteeType TrusteeType
	Trustee     SecurityID
	Permissions Permission
	AccessMode  AccessMode
}

func SetFilePermissions(fileName string, permissions []PermissionEntry) error {
	explicitAccess := []wrappers.EXPLICIT_ACCESS{}
	for _, entry := range permissions {
		explicitAccess = append(explicitAccess, wrappers.EXPLICIT_ACCESS{
			AccessPermissions: uint32(entry.Permissions),
			AccessMode:        int32(entry.AccessMode),
			Inheritance:       wrappers.NO_INHERITANCE,
			Trustee:           wrappers.TRUSTEE{
				MultipleTrustee:          nil,
				MultipleTrusteeOperation: wrappers.NO_MULTIPLE_TRUSTEE,
				TrusteeForm:              wrappers.TRUSTEE_IS_SID,
				TrusteeType:              int32(entry.TrusteeType),
				Name:                     (*uint16)(unsafe.Pointer(entry.Trustee.sid)),
			},
		})
	}
	
	var acl *wrappers.ACL
	if err := wrappers.SetEntriesInAcl(uint32(len(explicitAccess)), &explicitAccess[0], nil, &acl); err != nil {
		return NewWindowsError("SetEntriesInAcl", err)
	}
	defer wrappers.LocalFree(syscall.Handle(unsafe.Pointer(acl)))

	sd := make([]byte, wrappers.SECURITY_DESCRIPTOR_MIN_LENGTH)
	if err := wrappers.InitializeSecurityDescriptor(&sd[0], wrappers.SECURITY_DESCRIPTOR_REVISION); err != nil {
		return NewWindowsError("InitializeSecurityDescriptor", err)
	}
	if err := wrappers.SetSecurityDescriptorDacl(&sd[0], true, acl, false); err != nil {
		return NewWindowsError("SetSecurityDescriptorDacl", err)
	}

	err := wrappers.SetFileSecurity(
		syscall.StringToUTF16Ptr(fileName),
		wrappers.DACL_SECURITY_INFORMATION,
		&sd[0])
	if err != nil {
		return NewWindowsError("SetFileSecurity", err)
	}

	return nil
}
