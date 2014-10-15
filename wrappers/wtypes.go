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

const (
	CLSCTX_INPROC_SERVER          = 0x00000001
	CLSCTX_INPROC_HANDLER         = 0x00000002
	CLSCTX_LOCAL_SERVER           = 0x00000004
	CLSCTX_INPROC_SERVER16        = 0x00000008
	CLSCTX_REMOTE_SERVER          = 0x00000010
	CLSCTX_INPROC_HANDLER16       = 0x00000020
	CLSCTX_RESERVED1              = 0x00000040
	CLSCTX_RESERVED2              = 0x00000080
	CLSCTX_RESERVED3              = 0x00000100
	CLSCTX_RESERVED4              = 0x00000200
	CLSCTX_NO_CODE_DOWNLOAD       = 0x00000400
	CLSCTX_RESERVED5              = 0x00000800
	CLSCTX_NO_CUSTOM_MARSHAL      = 0x00001000
	CLSCTX_ENABLE_CODE_DOWNLOAD   = 0x00002000
	CLSCTX_NO_FAILURE_LOG         = 0x00004000
	CLSCTX_DISABLE_AAA            = 0x00008000
	CLSCTX_ENABLE_AAA             = 0x00010000
	CLSCTX_FROM_DEFAULT_CONTEXT   = 0x00020000
	CLSCTX_ACTIVATE_32_BIT_SERVER = 0x00040000
	CLSCTX_ACTIVATE_64_BIT_SERVER = 0x00080000
	CLSCTX_ENABLE_CLOAKING        = 0x00100000
	CLSCTX_PS_DLL                 = 0x80000000
)
