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

func ParseCommandLine(commandLine string) ([]string, error) {
	var numArgs int32
	rawArgs, err := wrappers.CommandLineToArgvW(
		syscall.StringToUTF16Ptr(commandLine),
		&numArgs)
	if err != nil {
		return nil, NewWindowsError("CommandLineToArgvW", err)
	}
	args := make([]string, numArgs)
	for i, _ := range(args) {
		args[i] = LpstrToString(*rawArgs)
		rawArgs = (**uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(rawArgs)) + unsafe.Sizeof(*rawArgs)))
	}
	return args, nil
}
