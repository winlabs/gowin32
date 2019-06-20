/*
 * Copyright (c) 2014-2019 MongoDB, Inc.
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
 )

 func EnableVTSequences() error {
	err := addStdConsoleModeFlag(
		wrappers.STD_OUTPUT_HANDLE,
		wrappers.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	if err != nil {
		return err
	}
	return addStdConsoleModeFlag(
		wrappers.STD_ERROR_HANDLE,
		wrappers.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}

func addStdConsoleModeFlag(stdHandle uint32, modeFlag uint32) error {
	hConsole, err := wrappers.GetStdHandle(stdHandle)
	if err != nil {
		return NewWindowsError("GetStdHandle", err)
	}
	var mode uint32
	if err := wrappers.GetConsoleMode(hConsole, &mode); err != nil {
		return NewWindowsError("GetConsoleMode", err)
	}
	mode |= modeFlag
	if err := wrappers.SetConsoleMode(hConsole, mode); err != nil {
		return NewWindowsError("SetConsoleMode", err)
	}
	return nil
 }
