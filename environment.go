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
)

func ExpandEnvironment(text string) (string, error) {
	size, err := wrappers.ExpandEnvironmentStrings(syscall.StringToUTF16Ptr(text), nil, 0)
	if err != nil {
		return "", err
	}
	buf := make([]uint16, size)
	if _, err := wrappers.ExpandEnvironmentStrings(syscall.StringToUTF16Ptr(text), &buf[0], size); err != nil {
		return "", err
	}
	return syscall.UTF16ToString(buf), nil
}
