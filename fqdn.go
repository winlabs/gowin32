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

func GetFQDN() (string, error) {
	fqdnLength := uint32(0)
	err := wrappers.GetComputerNameEx(wrappers.ComputerNameDnsFullyQualified, nil, &fqdnLength)
	if err != wrappers.ERROR_MORE_DATA {
		return "", NewWindowsError("GetComputerNameEx", err)
	}
	fqdnBuffer := make([]uint16, fqdnLength)
	err = wrappers.GetComputerNameEx(wrappers.ComputerNameDnsFullyQualified, &fqdnBuffer[0], &fqdnLength)
	if err != nil {
		return "", NewWindowsError("GetComputerNameEx", err)
	}
	return syscall.UTF16ToString(fqdnBuffer), nil
}
