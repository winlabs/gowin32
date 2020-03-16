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
	"syscall"

	"github.com/winlabs/gowin32/wrappers"
)

type InitiateShutdownFlag uint32

const (
	InitiateShutdownFlagForceOthers        InitiateShutdownFlag = wrappers.SHUTDOWN_FORCE_OTHERS
	InitiateShutdownFlagForceSelf          InitiateShutdownFlag = wrappers.SHUTDOWN_FORCE_SELF
	InitiateShutdownFlagRestart            InitiateShutdownFlag = wrappers.SHUTDOWN_RESTART
	InitiateShutdownFlagPowerOff           InitiateShutdownFlag = wrappers.SHUTDOWN_POWEROFF
	InitiateShutdownFlagNoReboot           InitiateShutdownFlag = wrappers.SHUTDOWN_NOREBOOT
	InitiateShutdownFlagGraceOverride      InitiateShutdownFlag = wrappers.SHUTDOWN_GRACE_OVERRIDE
	InitiateShutdownFlagInstallUpdates     InitiateShutdownFlag = wrappers.SHUTDOWN_INSTALL_UPDATES
	InitiateShutdownFlagRestartApps        InitiateShutdownFlag = wrappers.SHUTDOWN_RESTARTAPPS
	InitiateShutdownFlagSkipSvcPreshutdown InitiateShutdownFlag = wrappers.SHUTDOWN_SKIP_SVC_PRESHUTDOWN
	InitiateShutdownFlagHybrid             InitiateShutdownFlag = wrappers.SHUTDOWN_HYBRID
)

type InitiateShutdownReason uint32

const (
	InitiateShutdownReasonUserDefinied InitiateShutdownReason = wrappers.SHTDN_REASON_FLAG_USER_DEFINED
	InitiateShutdownReasonPlaned       InitiateShutdownReason = wrappers.SHTDN_REASON_FLAG_PLANNED
)

const (
	InitiateShutdownReasonMajorOther          InitiateShutdownReason = wrappers.SHTDN_REASON_MAJOR_OTHER
	InitiateShutdownReasonMajorHardware       InitiateShutdownReason = wrappers.SHTDN_REASON_MAJOR_HARDWARE
	InitiateShutdownReasonMajorperatingSystem InitiateShutdownReason = wrappers.SHTDN_REASON_MAJOR_OPERATINGSYSTEM
	InitiateShutdownReasonMajorSoftware       InitiateShutdownReason = wrappers.SHTDN_REASON_MAJOR_SOFTWARE
	InitiateShutdownReasonMajorApplication    InitiateShutdownReason = wrappers.SHTDN_REASON_MAJOR_APPLICATION
	InitiateShutdownReasonMajorSystem         InitiateShutdownReason = wrappers.SHTDN_REASON_MAJOR_SYSTEM
	InitiateShutdownReasonMajorPower          InitiateShutdownReason = wrappers.SHTDN_REASON_MAJOR_POWER
	InitiateShutdownReasonMajorLegacyApi      InitiateShutdownReason = wrappers.SHTDN_REASON_MAJOR_LEGACY_API
)

const (
	InitiateShutdownReasonMinorOther                InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_OTHER
	InitiateShutdownReasonMinorMaintenance          InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_MAINTENANCE
	InitiateShutdownReasonMinorInstallation         InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_INSTALLATION
	InitiateShutdownReasonMinorUpgrade              InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_UPGRADE
	InitiateShutdownReasonMinorReconfig             InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_RECONFIG
	InitiateShutdownReasonMinorHung                 InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_HUNG
	InitiateShutdownReasonMinorUnstable             InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_UNSTABLE
	InitiateShutdownReasonMinorDisk                 InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_DISK
	InitiateShutdownReasonMinorProcessor            InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_PROCESSOR
	InitiateShutdownReasonMinorNetwordCard          InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_NETWORKCARD
	InitiateShutdownReasonMinorPowerSupply          InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_POWER_SUPPLY
	InitiateShutdownReasonMinorCordunplugged        InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_CORDUNPLUGGED
	InitiateShutdownReasonMinorEnvironment          InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_ENVIRONMENT
	InitiateShutdownReasonMinorHardwareDriver       InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_HARDWARE_DRIVER
	InitiateShutdownReasonMinorOtherdriver          InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_OTHERDRIVER
	InitiateShutdownReasonMinorBluescreen           InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_BLUESCREEN
	InitiateShutdownReasonMinorServicepack          InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_SERVICEPACK
	InitiateShutdownReasonMinorHotfix               InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_HOTFIX
	InitiateShutdownReasonMinorSecurityfix          InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_SECURITYFIX
	InitiateShutdownReasonMinorSecurity             InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_SECURITY
	InitiateShutdownReasonMinorConnectivity         InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_NETWORK_CONNECTIVITY
	InitiateShutdownReasonMinormi                   InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_WMI
	InitiateShutdownReasonMinorServicepackUninstall InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_SERVICEPACK_UNINSTALL
	InitiateShutdownReasonMinorHotfixUninstall      InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_HOTFIX_UNINSTALL
	InitiateShutdownReasonMinorSecurityfixUninstall InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_SECURITYFIX_UNINSTALL
	InitiateShutdownReasonMinorMmc                  InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_MMC
	InitiateShutdownReasonMinorTermsrv              InitiateShutdownReason = wrappers.SHTDN_REASON_MINOR_TERMSRV
)

func InitiateShutdown(machineName string, message string, gracePeriod int, shutdownFlags InitiateShutdownFlag, reason InitiateShutdownReason) error {
	if err := wrappers.InitiateShutdown(
		syscall.StringToUTF16Ptr(machineName),
		syscall.StringToUTF16Ptr(message),
		uint32(gracePeriod),
		uint32(shutdownFlags),
		uint32(reason)); err != nil {
		return NewWindowsError("InitiateShutdown", err)
	}
	return nil
}

func InitiateSystemShutdown(machineName string, message string, timeout int, forceAppsClosed bool, rebootAfterShutdown bool) error {
	if err := wrappers.InitiateSystemShutdown(
		syscall.StringToUTF16Ptr(machineName),
		syscall.StringToUTF16Ptr(message),
		uint32(timeout),
		forceAppsClosed,
		rebootAfterShutdown); err != nil {
		return NewWindowsError("InitiateSystemShutdown", err)
	}
	return nil
}
