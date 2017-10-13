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

package wrappers

const (
	SE_UNKNOWN_OBJECT_TYPE     = 0
	SE_FILE_OBJECT             = 1
	SE_SERVICE                 = 2
	SE_PRINTER                 = 3
	SE_REGISTRY_KEY            = 4
	SE_LMSHARE                 = 5
	SE_KERNEL_OBJECT           = 6
	SE_WINDOW_OBJECT           = 7
	SE_DS_OBJECT               = 8
	SE_DS_OBJECT_ALL           = 9
	SE_PROVIDER_DEFINED_OBJECT = 10
	SE_WMIGUID_OBJECT          = 11
	SE_REGISTRY_WOW64_32KEY    = 12
)

const (
	TRUSTEE_IS_UNKNOWN          = 0
	TRUSTEE_IS_USER             = 1
	TRUSTEE_IS_GROUP            = 2
	TRUSTEE_IS_DOMAIN           = 3
	TRUSTEE_IS_ALIAS            = 4
	TRUSTEE_IS_WELL_KNOWN_GROUP = 5
	TRUSTEE_IS_DELETED          = 6
	TRUSTEE_IS_INVALID          = 7
	TRUSTEE_IS_COMPUTER         = 8
)

const (
	TRUSTEE_IS_SID              = 0
	TRUSTEE_IS_NAME             = 1
	TRUSTEE_BAD_FORM            = 2
	TRUSTEE_IS_OBJECTS_AND_SID  = 3
	TRUSTEE_IS_OBJECTS_AND_NAME = 4
)

const (
	NO_MULTIPLE_TRUSTEE    = 0
	TRUSTEE_IS_IMPERSONATE = 1
)

type OBJECTS_AND_SID struct {
	ObjectsPresent          uint32
	ObjectTypeGuid          GUID
	InheritedObjectTypeGuid GUID
	Sid                     *SID
}

type OBJECTS_AND_NAME struct {
	ObjectsPresent          uint32
	ObjectType              int32
	ObjectTypeName          *uint16
	InheritedObjectTypeName *uint16
	Name                    *uint16
}

type TRUSTEE struct {
	MultipleTrustee          *TRUSTEE
	MultipleTrusteeOperation int32
	TrusteeForm              int32
	TrusteeType              int32
	Name                     *uint16
}

const (
	NOT_USED_ACCESS   = 0
	GRANT_ACCESS      = 1
	SET_ACCESS        = 2
	DENY_ACCESS       = 3
	REVOKE_ACCESS     = 4
	SET_AUDIT_SUCCESS = 5
	SET_AUDIT_FAILURE = 6
)

const (
	NO_INHERITANCE                     = 0x0
	SUB_OBJECTS_ONLY_INHERIT           = 0x1
	SUB_CONTAINERS_ONLY_INHERIT        = 0x2
	SUB_CONTAINERS_AND_OBJECTS_INHERIT = 0x3
	INHERIT_NO_PROPAGATE               = 0x4
	INHERIT_ONLY                       = 0x8
)

type EXPLICIT_ACCESS struct {
	AccessPermissions uint32
	AccessMode        int32
	Inheritance       uint32
	Trustee           TRUSTEE
}
