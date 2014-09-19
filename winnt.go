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

const (
	OWNER_SECURITY_INFORMATION = 0x00000001
	GROUP_SECURITY_INFORMATION = 0x00000002
	DACL_SECURITY_INFORMATION = 0x00000004
	SACL_SECURITY_INFORMATION = 0x00000008
	LABEL_SECURITY_INFORMATION = 0x00000010
	PROTECTED_DACL_SECURITY_INFORMATION = 0x80000000
	PROTECTED_SACL_SECURITY_INFORMATION = 0x40000000
	UNPROTECTED_DACL_SECURITY_INFORMATION = 0x20000000
	UNPROTECTED_SACL_SECURITY_INFORMATION = 0x10000000
)
