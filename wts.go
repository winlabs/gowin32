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

package gowin32

import (
	"strconv"

	"github.com/winlabs/gowin32/wrappers"

	"fmt"
	"net"
	"syscall"
	"time"
	"unsafe"
)

// ConnectState enum type - Go version of WTS_CONNECTSTATE_CLASS
type ConnectState uint32

const (
	Active       ConnectState = wrappers.WTSActive
	Connected    ConnectState = wrappers.WTSConnected
	ConnectQuery ConnectState = wrappers.WTSConnectQuery
	Shadow       ConnectState = wrappers.WTSShadow
	Disconnected ConnectState = wrappers.WTSDisconnected
	Idle         ConnectState = wrappers.WTSIdle
	Listen       ConnectState = wrappers.WTSListen
	Reset        ConnectState = wrappers.WTSReset
	Down         ConnectState = wrappers.WTSDown
	Init         ConnectState = wrappers.WTSInit
)

// ClientProtocolType enum type go version of WTSClientProtocolType
// https://msdn.microsoft.com/en-us/library/aa383861%28v=vs.85%29.aspx
type ClientProtocolType uint32

const (
	ConsoleSession ClientProtocolType = 0
	Internal       ClientProtocolType = 1
	RDP            ClientProtocolType = 2
)

// SessionInfo - go version of WTS_SESSION_INFO structure
type SessionInfo struct {
	SessionID      uint32
	WinStationName string
	State          ConnectState
}

// ClientInfo - go version of WTSCLIENT structure
type ClientInfo struct {
	ClientName          string
	Domain              string
	UserName            string
	WorkDirectory       string
	InitialProgram      string
	EncryptionLevel     byte
	ClientAddressFamily uint32
	ClientAddress       [wrappers.CLIENTADDRESS_LENGTH + 1]uint16 // convert to net.IP ?
	HRes                uint16
	VRes                uint16
	ColorDepth          uint16
	ClientDirectory     string
	ClientBuildNumber   uint32
	ClientHardwareId    uint32
	ClientProductId     uint16
	OutBufCountHost     uint16
	OutBufCountClient   uint16
	OutBufLength        uint16
	DeviceId            string
}

// Info - go version of WTSINFO structure
type Info struct {
	State                   ConnectState
	SessionId               uint32
	IncomingBytes           uint32
	OutgoingBytes           uint32
	IncomingFrames          uint32
	OutgoingFrames          uint32
	IncomingCompressedBytes uint32
	OutgoingCompressedBytes uint32
	WinStationName          string
	Domain                  string
	UserName                string
	ConnectTime             time.Time
	DisconnectTime          time.Time
	LastInputTime           time.Time
	LogonTime               time.Time
	CurrentTime             time.Time
}

const connectStateName = "ActiveConnectedConnectQueryShadowDisconnectedIdleListenResetDownInit"

var connectStateNameIndex = [...]uint8{0, 6, 15, 27, 33, 45, 49, 55, 60, 64, 68}

func (c ConnectState) String() string {
	if c >= ConnectState(len(connectStateNameIndex)-1) {
		return "ConnectState(" + strconv.FormatInt(int64(c), 10) + ")"
	}
	return connectStateName[connectStateNameIndex[c]:connectStateNameIndex[c+1]]
}

type WTS struct {
	handle syscall.Handle
}

func NewWTS(serverName string) *WTS {
	result := WTS{}
	if serverName != "" {
		result.OpenServer(serverName)
	}
	return &result
}

func (wts *WTS) OpenServer(serverName string) {
	wts.CloseServer()
	wts.handle = wrappers.WTSOpenServer(syscall.StringToUTF16Ptr(serverName))
}

func (wts *WTS) CloseServer() {
	if wts.handle != 0 {
		wrappers.WTSCloseServer(wts.handle)
		wts.handle = 0
	}
}

func (wts *WTS) EnumerateSessions() ([]SessionInfo, error) {
	var sessionInfo *wrappers.WTS_SESSION_INFO
	var count uint32

	if err := wrappers.WTSEnumerateSessions(wts.handle, 0, 1, &sessionInfo, &count); err != nil {
		return nil, err
	}
	defer wrappers.WTSFreeMemory(unsafe.Pointer(sessionInfo))

	si := sessionInfo
	result := make([]SessionInfo, count)
	for i := uint32(0); i < count; i++ {
		result[i] = SessionInfo{SessionID: si.SessionId,
			WinStationName: utf16PtrToString(si.WinStationName),
			State:          ConnectState(si.State)}
		si = (*wrappers.WTS_SESSION_INFO)(unsafe.Pointer(uintptr(unsafe.Pointer(si)) + unsafe.Sizeof(*si)))
	}
	return result, nil
}

func (wts *WTS) QuerySessionInitialProgram(sessionID uint32) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSInitialProgram)
}

func (wts *WTS) QuerySessionApplicationName(sessionID uint32) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSApplicationName)
}
func (wts *WTS) QuerySessionWorkingDirectory(sessionID uint32) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSWorkingDirectory)
}

func (wts *WTS) QuerySessionID(sessionID uint32) (uint32, error) {
	return wts.querySessionInformationAsUint32(sessionID, wrappers.WTSSessionId)
}

func (wts *WTS) QuerySessionUserName(sessionID uint32) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSUserName)
}

func (wts *WTS) QuerySessionWinStationName(sessionID uint32) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSWinStationName)
}

func (wts *WTS) QuerySessionDomainName(sessionID uint32) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSDomainName)
}

func (wts *WTS) QuerySessionConnectState(sessionID uint32) (ConnectState, error) {
	r1, err := wts.querySessionInformationAsUint32(sessionID, wrappers.WTSConnectState)
	return ConnectState(r1), err
}

func (wts *WTS) QuerySessionClientBuildNumber(sessionID uint32) (uint32, error) {
	return wts.querySessionInformationAsUint32(sessionID, wrappers.WTSClientBuildNumber)
}

func (wts *WTS) QuerySessionClientName(sessionID uint32) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSClientName)
}

func (wts *WTS) QuerySessionClientDirectory(sessionID uint32) (string, error) {
	return wts.querySessionInformationAsString(sessionID, wrappers.WTSClientDirectory)
}

func (wts *WTS) QuerySessionClientProductId(sessionID uint32) (uint16, error) {
	return wts.querySessionInformationAsUint16(sessionID, wrappers.WTSClientProductId)
}

func (wts *WTS) QuerySessionClientHardwareId(sessionID uint32) (uint32, error) {
	return wts.querySessionInformationAsUint32(sessionID, wrappers.WTSClientHardwareId)
}

func (wts *WTS) QuerySessionClientAddress(sessionID uint32) (net.IP, error) {
	var buffer uintptr
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, sessionID, wrappers.WTSClientAddress, &buffer, &bytesReturned); err != nil {
		return net.IP{}, err
	}
	defer wrappers.WTSFreeMemory(unsafe.Pointer(buffer))

	if bytesReturned != 24 {
		return net.IP{}, buferSizeError(24, bytesReturned)
	}
	// MS doc: The IP address is offset by two bytes from the start of the Address member of the WTS_CLIENT_ADDRESS structure.
	// https://msdn.microsoft.com/en-us/library/aa383861%28v=vs.85%29.aspx
	a := *(*wrappers.WTS_CLIENT_ADDRESS)(unsafe.Pointer(buffer))
	return clientAddressToIP(a.AddressFamily, a.Address[2:])
}

func (wts *WTS) QuerySessionClientDisplay(sessionID uint32) (wrappers.WTS_CLIENT_DISPLAY, error) {
	var buffer uintptr
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, sessionID, wrappers.WTSClientDisplay, &buffer, &bytesReturned); err != nil {
		return wrappers.WTS_CLIENT_DISPLAY{}, err
	}
	defer wrappers.WTSFreeMemory(unsafe.Pointer(buffer))

	if bytesReturned != 12 {
		return wrappers.WTS_CLIENT_DISPLAY{}, buferSizeError(12, bytesReturned)
	}
	return *(*wrappers.WTS_CLIENT_DISPLAY)(unsafe.Pointer(buffer)), nil
}

func (wts *WTS) QuerySessionClientProtocolType(sessionID uint32) (ClientProtocolType, error) {
	r1, err := wts.querySessionInformationAsUint16(sessionID, wrappers.WTSClientProtocolType)
	return ClientProtocolType(r1), err
}

func (wts *WTS) QuerySessionClientInfo(sessionID uint32) (ClientInfo, error) {
	var buffer uintptr
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, sessionID, wrappers.WTSClientInfo, &buffer, &bytesReturned); err != nil {
		return ClientInfo{}, err
	}
	defer wrappers.WTSFreeMemory(unsafe.Pointer(buffer))

	if bytesReturned != 2304 {
		return ClientInfo{}, buferSizeError(2304, bytesReturned)
	}
	c := *(*wrappers.WTSCLIENT)(unsafe.Pointer(buffer))
	return ClientInfo{
		ClientName:          syscall.UTF16ToString(c.ClientName[:]),
		Domain:              syscall.UTF16ToString(c.Domain[:]),
		UserName:            syscall.UTF16ToString(c.UserName[:]),
		WorkDirectory:       syscall.UTF16ToString(c.WorkDirectory[:]),
		InitialProgram:      syscall.UTF16ToString(c.InitialProgram[:]),
		EncryptionLevel:     c.EncryptionLevel,
		ClientAddressFamily: c.ClientAddressFamily,
		ClientAddress:       c.ClientAddress,
		HRes:                c.HRes,
		VRes:                c.VRes,
		ColorDepth:          c.ColorDepth,
		ClientDirectory:     syscall.UTF16ToString(c.ClientDirectory[:]),
		ClientBuildNumber:   c.ClientBuildNumber,
		ClientHardwareId:    c.ClientHardwareId,
		ClientProductId:     c.ClientProductId,
		OutBufCountHost:     c.OutBufCountHost,
		OutBufCountClient:   c.OutBufCountClient,
		OutBufLength:        c.OutBufLength,
		DeviceId:            syscall.UTF16ToString(c.DeviceId[:]),
	}, nil
}

func (wts *WTS) QuerySessionSesionInfo(sessionID uint32) (Info, error) {
	var buffer uintptr
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, sessionID, wrappers.WTSSessionInfo, &buffer, &bytesReturned); err != nil {
		return Info{}, err
	}
	defer wrappers.WTSFreeMemory(unsafe.Pointer(buffer))

	if bytesReturned != 216 {
		return Info{}, buferSizeError(216, bytesReturned)
	}
	i := *(*wrappers.WTSINFO)(unsafe.Pointer(buffer))
	return Info{
		State:                   ConnectState(i.State),
		SessionId:               i.SessionId,
		IncomingBytes:           i.IncomingBytes,
		OutgoingBytes:           i.OutgoingBytes,
		IncomingFrames:          i.IncomingFrames,
		OutgoingFrames:          i.OutgoingFrames,
		IncomingCompressedBytes: i.IncomingCompressedBytes,
		OutgoingCompressedBytes: i.OutgoingCompressedBytes,
		WinStationName:          syscall.UTF16ToString(i.WinStationName[:]),
		Domain:                  syscall.UTF16ToString(i.Domain[:]),
		UserName:                syscall.UTF16ToString(i.UserName[:]),
		ConnectTime:             windowsFileTimeToTime(i.ConnectTime),
		DisconnectTime:          windowsFileTimeToTime(i.DisconnectTime),
		LastInputTime:           windowsFileTimeToTime(i.LastInputTime),
		LogonTime:               windowsFileTimeToTime(i.LogonTime),
		CurrentTime:             windowsFileTimeToTime(i.CurrentTime)}, nil
}

func (wts *WTS) QuerySessionAddressV4(sessionID uint32) (wrappers.WTS_CLIENT_ADDRESS, error) {
	var buffer uintptr
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, sessionID, wrappers.WTSSessionAddressV4, &buffer, &bytesReturned); err != nil {
		return wrappers.WTS_CLIENT_ADDRESS{}, err
	}
	defer wrappers.WTSFreeMemory(unsafe.Pointer(buffer))

	if bytesReturned != 24 {
		return wrappers.WTS_CLIENT_ADDRESS{}, buferSizeError(24, bytesReturned)
	}
	return *(*wrappers.WTS_CLIENT_ADDRESS)(unsafe.Pointer(buffer)), nil
}

func (wts *WTS) QuerySessionIsRemoteSession(sessionID uint32) (bool, error) {
	return wts.querySessionInformationAsBool(sessionID, wrappers.WTSIsRemoteSession)
}

func (wts *WTS) QueryUserToken(sessionID uint32) (*Token, error) {
	var handle syscall.Handle
	if err := wrappers.WTSQueryUserToken(sessionID, &handle); err != nil {
		return nil, err
	}
	return &Token{handle: handle}, nil
}

func clientAddressToIP(addressFamily uint32, address []byte) (net.IP, error) {
	switch addressFamily {
	case syscall.AF_INET:
		if len(address) >= 4 {
			return net.IPv4(address[0], address[1], address[2], address[3]), nil
		}
	case syscall.AF_INET6:
		if len(address) >= 16 {
			return net.IP(address[:16]), nil
		}
	}
	return nil, fmt.Errorf("Unknown addressFamily: %v", addressFamily)
}

func buferSizeError(excpected, returned uint32) error {
	return fmt.Errorf("Invalid buffer size. Expected: %d returned: %d", excpected, returned)
}

func utf16BufferToString(buffer *uintptr, len uint32) string {
	return syscall.UTF16ToString((*[256]uint16)(unsafe.Pointer(*buffer))[:len])
}

func utf16PtrToString(s *uint16) string {
	return syscall.UTF16ToString((*[256]uint16)(unsafe.Pointer(s))[:])
}

func windowsFileTimeToTime(fileTime int64) time.Time {
	const TicksPerSecond = 10000000
	const EpochDifference = 11644473600
	// we also can use win32 api FileTimeToSystemTime
	return time.Unix((fileTime/TicksPerSecond)-EpochDifference, 0)
}

func (wts *WTS) querySessionInformationAsBool(sessionID uint32, infoClass uint32) (bool, error) {
	var buffer uintptr
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, sessionID, infoClass, &buffer, &bytesReturned); err != nil {
		return false, err
	}
	defer wrappers.WTSFreeMemory(unsafe.Pointer(buffer))

	if bytesReturned != 1 {
		return false, buferSizeError(1, bytesReturned)
	}
	return *(*bool)(unsafe.Pointer(buffer)), nil
}

func (wts *WTS) querySessionInformationAsString(sessionID uint32, infoClass uint32) (string, error) {
	var buffer uintptr
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, sessionID, infoClass, &buffer, &bytesReturned); err != nil {
		return "", err
	}
	defer wrappers.WTSFreeMemory(unsafe.Pointer(buffer))

	return utf16BufferToString(&buffer, bytesReturned), nil
}

func (wts *WTS) querySessionInformationAsUint32(sessionID uint32, infoClass uint32) (uint32, error) {
	var buffer uintptr
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, sessionID, infoClass, &buffer, &bytesReturned); err != nil {
		return 0, err
	}
	defer wrappers.WTSFreeMemory(unsafe.Pointer(buffer))

	if bytesReturned != 4 {
		return 0, buferSizeError(4, bytesReturned)
	}
	return *(*uint32)(unsafe.Pointer(buffer)), nil
}

func (wts *WTS) querySessionInformationAsUint16(sessionID uint32, infoClass uint32) (uint16, error) {
	var buffer uintptr
	var bytesReturned uint32

	if err := wrappers.WTSQuerySessionInformation(wts.handle, sessionID, infoClass, &buffer, &bytesReturned); err != nil {
		return 0, err
	}
	defer wrappers.WTSFreeMemory(unsafe.Pointer(buffer))

	if bytesReturned != 2 {
		return 0, buferSizeError(2, bytesReturned)
	}
	return *(*uint16)(unsafe.Pointer(buffer)), nil
}
