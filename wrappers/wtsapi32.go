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

import (
	"syscall"
	"unsafe"
)

// Misc consts from WtsApi32.h
const (
	CLIENTNAME_LENGTH     = 20
	DOMAIN_LENGTH         = 17
	USERNAME_LENGTH       = 20
	CLIENTADDRESS_LENGTH  = 30
	WINSTATIONNAME_LENGTH = 32
)

// WTS_CONNECTSTATE_CLASS enumeration
const (
	WTSActive       = 0
	WTSConnected    = 1
	WTSConnectQuery = 2
	WTSShadow       = 3
	WTSDisconnected = 4
	WTSIdle         = 5
	WTSListen       = 6
	WTSReset        = 7
	WTSDown         = 8
	WTSInit         = 9
)

type WTS_SESSION_INFO struct {
	SessionId      uint32
	WinStationName *uint16
	State          uint32
}

// WTS_INFO_CLASS enumeration
const (
	WTSInitialProgram     = 0
	WTSApplicationName    = 1
	WTSWorkingDirectory   = 2
	WTSOEMId              = 3
	WTSSessionId          = 4
	WTSUserName           = 5
	WTSWinStationName     = 6
	WTSDomainName         = 7
	WTSConnectState       = 8
	WTSClientBuildNumber  = 9
	WTSClientName         = 10
	WTSClientDirectory    = 11
	WTSClientProductId    = 12
	WTSClientHardwareId   = 13
	WTSClientAddress      = 14
	WTSClientDisplay      = 15
	WTSClientProtocolType = 16
	WTSIdleTime           = 17
	WTSLogonTime          = 18
	WTSIncomingBytes      = 19
	WTSOutgoingBytes      = 20
	WTSIncomingFrames     = 21
	WTSOutgoingFrames     = 22
	WTSClientInfo         = 23
	WTSSessionInfo        = 24
	WTSSessionInfoEx      = 25
	WTSConfigInfo         = 26
	WTSValidationInfo     = 27
	WTSSessionAddressV4   = 28
	WTSIsRemoteSession    = 29
)

type WTSINFO struct {
	State                   uint32
	SessionId               uint32
	IncomingBytes           uint32
	OutgoingBytes           uint32
	IncomingFrames          uint32
	OutgoingFrames          uint32
	IncomingCompressedBytes uint32
	OutgoingCompressedBytes uint32
	WinStationName          [WINSTATIONNAME_LENGTH]uint16
	Domain                  [DOMAIN_LENGTH]uint16
	UserName                [USERNAME_LENGTH + 1]uint16
	ConnectTime             int64
	DisconnectTime          int64
	LastInputTime           int64
	LogonTime               int64
	CurrentTime             int64
}

type WTSCLIENT struct {
	ClientName          [CLIENTNAME_LENGTH + 1]uint16
	Domain              [DOMAIN_LENGTH + 1]uint16
	UserName            [USERNAME_LENGTH + 1]uint16
	WorkDirectory       [MAX_PATH + 1]uint16
	InitialProgram      [MAX_PATH + 1]uint16
	EncryptionLevel     byte
	ClientAddressFamily uint32
	ClientAddress       [CLIENTADDRESS_LENGTH + 1]uint16
	HRes                uint16
	VRes                uint16
	ColorDepth          uint16
	ClientDirectory     [MAX_PATH + 1]uint16
	ClientBuildNumber   uint32
	ClientHardwareId    uint32
	ClientProductId     uint16
	OutBufCountHost     uint16
	OutBufCountClient   uint16
	OutBufLength        uint16
	DeviceId            [MAX_PATH + 1]uint16
}

type WTS_CLIENT_ADDRESS struct {
	AddressFamily uint32
	Address       [20]byte
}

type WTS_CLIENT_DISPLAY struct {
	HorizontalResolution uint32
	VerticalResolution   uint32
	ColorDepth           uint32
}

var (
	modwtsapi32 = syscall.NewLazyDLL("wtsapi32.dll")

	procWTSCloseServer             = modwtsapi32.NewProc("WTSCloseServer")
	procWTSEnumerateSessions       = modwtsapi32.NewProc("WTSEnumerateSessionsW")
	procWTSFreeMemory              = modwtsapi32.NewProc("WTSFreeMemory")
	procWTSOpenServer              = modwtsapi32.NewProc("WTSOpenServerW")
	procWTSLogoffSession           = modwtsapi32.NewProc("WTSLogoffSession")
	procWTSQuerySessionInformation = modwtsapi32.NewProc("WTSQuerySessionInformationW")
	procWTSQueryUserToken          = modwtsapi32.NewProc("WTSQueryUserToken")
)

func WTSCloseServer(handle syscall.Handle) {
	syscall.Syscall6(procWTSCloseServer.Addr(), 1, uintptr(handle), 0, 0, 0, 0, 0)
}

func WTSEnumerateSessions(server syscall.Handle, reserved uint32, version uint32, pSessionInfo **WTS_SESSION_INFO, count *uint32) error {
	r1, _, e1 := syscall.Syscall6(
		procWTSEnumerateSessions.Addr(),
		5,
		uintptr(server),
		uintptr(reserved),
		uintptr(version),
		uintptr(unsafe.Pointer(pSessionInfo)),
		uintptr(unsafe.Pointer(count)),
		0)

	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func WTSFreeMemory(memory *byte) {
	syscall.Syscall(procWTSFreeMemory.Addr(), 1, uintptr(unsafe.Pointer(memory)), 0, 0)
}

func WTSOpenServer(serverName *uint16) syscall.Handle {
	r1, _, _ := syscall.Syscall(procWTSOpenServer.Addr(), 1, uintptr(unsafe.Pointer(serverName)), 0, 0)
	return syscall.Handle(r1)
}

func WTSQuerySessionInformation(handle syscall.Handle, sessionId uint32, infoClass uint32, buffer **uint16, bytesReturned *uint32) error {
	r1, _, e1 := syscall.Syscall6(
		procWTSQuerySessionInformation.Addr(),
		5,
		uintptr(handle),
		uintptr(sessionId),
		uintptr(infoClass),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(unsafe.Pointer(bytesReturned)), 0)

	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func WTSLogoffSession(handle syscall.Handle, sessionId uint32, wait bool) error {
	r1, _, e1 := syscall.Syscall(
		procWTSLogoffSession.Addr(),
		3,
		uintptr(handle),
		uintptr(sessionId),
		boolToUintptr(wait))

	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}

func WTSQueryUserToken(sessionId uint32, handle *syscall.Handle) error {
	r1, _, e1 := syscall.Syscall(
		procWTSQueryUserToken.Addr(),
		2,
		uintptr(sessionId),
		uintptr(unsafe.Pointer(handle)),
		0)

	if r1 == 0 {
		if e1 != ERROR_SUCCESS {
			return e1
		} else {
			return syscall.EINVAL
		}
	}
	return nil
}
