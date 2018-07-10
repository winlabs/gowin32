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

type DISPLAY_DEVICE struct {
	Cb           uint32
	DeviceName   [32]uint16
	DeviceString [128]uint16
	StateFlags   uint32
	DeviceID     [128]uint16
	DeviceKey    [128]uint16
}

const (
	DISPLAY_DEVICE_ACTIVE           = 0x00000001
	DISPLAY_DEVICE_PRIMARY_DEVICE   = 0x00000004
	DISPLAY_DEVICE_MIRRORING_DRIVER = 0x00000008
	DISPLAY_DEVICE_VGA_COMPATIBLE   = 0x00000010
	DISPLAY_DEVICE_REMOVABLE        = 0x00000020
	DISPLAY_DEVICE_MODESPRUNED      = 0x08000000
)
