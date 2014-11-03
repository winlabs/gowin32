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

import "syscall"

const (
	ERROR_GEN_FAILURE            syscall.Errno = 31
	ERROR_INVALID_PARAMETER      syscall.Errno = 87
	ERROR_MORE_DATA              syscall.Errno = 234
	ERROR_SERVICE_DOES_NOT_EXIST syscall.Errno = 1060
	ERROR_OLD_WIN_VERSION        syscall.Errno = 1150
	ERROR_UNSUPPORTED_TYPE       syscall.Errno = 1630
)

const (
	E_POINTER syscall.Errno = 0x80004003
)
