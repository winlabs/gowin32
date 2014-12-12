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
	FACILITY_WINRM                   = 51
	FACILITY_WINDOWSUPDATE           = 36
	FACILITY_WINDOWS_DEFENDER        = 80
	FACILITY_WINDOWS_CE              = 24
	FACILITY_WINDOWS                 = 8
	FACILITY_USERMODE_VOLMGR         = 56
	FACILITY_USERMODE_VIRTUALIZATION = 55
	FACILITY_USERMODE_VHD            = 58
	FACILITY_URT                     = 19
	FACILITY_UMI                     = 22
	FACILITY_TPM_SOFTWARE            = 41
	FACILITY_TPM_SERVICES            = 40
	FACILITY_SXS                     = 23
	FACILITY_STORAGE                 = 3
	FACILITY_STATE_MANAGEMENT        = 34
	FACILITY_SSPI                    = 9
	FACILITY_SCARD                   = 16
	FACILITY_SHELL                   = 39
	FACILITY_SETUPAPI                = 15
	FACILITY_SECURITY                = 9
	FACILITY_SDIAG                   = 60
	FACILITY_RPC                     = 1
	FACILITY_PLA                     = 48
	FACILITY_OPC                     = 81
	FACILITY_WIN32                   = 7
	FACILITY_CONTROL                 = 10
	FACILITY_WEBSERVICES             = 61
	FACILITY_NULL                    = 0
	FACILITY_NDIS                    = 52
	FACILITY_METADIRECTORY           = 35
	FACILITY_MSMQ                    = 14
	FACILITY_MEDIASERVER             = 13
	FACILITY_INTERNET                = 12
	FACILITY_ITF                     = 4
	FACILITY_USERMODE_HYPERVISOR     = 53
	FACILITY_HTTP                    = 25
	FACILITY_GRAPHICS                = 38
	FACILITY_FWP                     = 50
	FACILITY_FVE                     = 49
	FACILITY_USERMODE_FILTER_MANAGER = 31
	FACILITY_DPLAY                   = 21
	FACILITY_DISPATCH                = 2
	FACILITY_DIRECTORYSERVICE        = 37
	FACILITY_CONFIGURATION           = 33
	FACILITY_COMPLUS                 = 17
	FACILITY_USERMODE_COMMONLOG      = 26
	FACILITY_CMI                     = 54
	FACILITY_CERT                    = 11
	FACILITY_BCD                     = 57
	FACILITY_BACKGROUNDCOPY          = 32
	FACILITY_ACS                     = 20
	FACILITY_AAF                     = 18
)

const (
	ERROR_FILE_NOT_FOUND          syscall.Errno = 2
	ERROR_ACCESS_DENIED           syscall.Errno = 5
	ERROR_NO_MORE_FILES           syscall.Errno = 18
	ERROR_GEN_FAILURE             syscall.Errno = 31
	ERROR_INVALID_PARAMETER       syscall.Errno = 87
	ERROR_BROKEN_PIPE             syscall.Errno = 109
	ERROR_INSUFFICIENT_BUFFER     syscall.Errno = 122
	ERROR_MORE_DATA               syscall.Errno = 234
	WAIT_TIMEOUT                  syscall.Errno = 258
	ERROR_NO_MORE_ITEMS           syscall.Errno = 259
	ERROR_SERVICE_DOES_NOT_EXIST  syscall.Errno = 1060
	ERROR_OLD_WIN_VERSION         syscall.Errno = 1150
	ERROR_INSTALL_PACKAGE_INVALID syscall.Errno = 1620
	ERROR_UNSUPPORTED_TYPE        syscall.Errno = 1630
)

const (
	SEVERITY_SUCCESS = 0
	SEVERITY_ERROR   = 1
)

func SUCCEEDED(hr uint32) bool {
	return int32(hr) >= 0
}

func FAILED(hr uint32) bool {
	return int32(hr) < 0
}

func HRESULT_CODE(hr uint32) uint16 {
	return uint16(hr)
}

func HRESULT_FACILITY(hr uint32) uint16 {
	return uint16((hr >> 16) & 0x1FFF)
}

func HRESULT_SEVERITY(hr uint32) uint32 {
	return hr >> 31
}

const (
	E_UNEXPECTED                                  = 0x8000FFFF
	E_NOTIMPL                                     = 0x80004001
	E_OUTOFMEMORY                                 = 0x8007000E
	E_INVALIDARG                                  = 0x80070057
	E_NOINTERFACE                                 = 0x80004002
	E_POINTER                                     = 0x80004003
	E_HANDLE                                      = 0x80070006
	E_ABORT                                       = 0x80004004
	E_FAIL                                        = 0x80004005
	E_ACCESSDENIED                                = 0x80070005
	E_PENDING                                     = 0x8000000A
	CO_E_INIT_TLS                                 = 0x80004006
	CO_E_INIT_SHARED_ALLOCATOR                    = 0x80004007
	CO_E_INIT_MEMORY_ALLOCATOR                    = 0x80004008
	CO_E_INIT_CLASS_CACHE                         = 0x80004009
	CO_E_INIT_RPC_CHANNEL                         = 0x8000400A
	CO_E_INIT_TLS_SET_CHANNEL_CONTROL             = 0x8000400B
	CO_E_INIT_TLS_CHANNEL_CONTROL                 = 0x8000400C
	CO_E_INIT_UNACCEPTED_USER_ALLOCATOR           = 0x8000400D
	CO_E_INIT_SCM_MUTEX_EXISTS                    = 0x8000400E
	CO_E_INIT_SCM_FILE_MAPPING_EXISTS             = 0x8000400F
	CO_E_INIT_SCM_MAP_VIEW_OF_FILE                = 0x80004010
	CO_E_INIT_SCM_EXEC_FAILURE                    = 0x80004011
	CO_E_INIT_ONLY_SINGLE_THREADED                = 0x80004012
	CO_E_INIT_CANT_REMOTE                         = 0x80004013
	CO_E_BAD_SERVER_NAME                          = 0x80004014
	CO_E_WRONG_SERVER_IDENTITY                    = 0x80004015
	CO_E_OLE1DDE_DISABLED                         = 0x80004016
	CO_E_RUNAS_SYNTAX                             = 0x80004017
	CO_E_CREATEPROCESS_FAILURE                    = 0x80004018
	CO_E_RUNAS_CREATEPROCESS_FAILURE              = 0x80004019
	CO_E_RUNAS_LOGON_FAILURE                      = 0x8000401A
	CO_E_LAUNCH_PERMISSION_DENIED                 = 0x8000401B
	CO_E_START_SERVICE_FAILURE                    = 0x8000401C
	CO_E_REMOTE_COMMUNICATION_FAILURE             = 0x8000401D
	CO_E_SERVER_START_TIMEOUT                     = 0x8000401E
	CO_E_CLSREG_INCONSISTENT                      = 0x8000401F
	CO_E_IIDREG_INCONSISTENT                      = 0x80004020
	CO_E_NOT_SUPPORTED                            = 0x80004021
	CO_E_RELOAD_DLL                               = 0x80004022
	CO_E_MSI_ERROR                                = 0x80004023
	CO_E_ATTEMPT_TO_CREATE_OUTSIDE_CLIENT_CONTEXT = 0x80004024
	CO_E_SERVER_PAUSED                            = 0x80004025
	CO_E_SERVER_NOT_PAUSED                        = 0x80004026
	CO_E_CLASS_DISABLED                           = 0x80004027
	CO_E_CLRNOTAVAILABLE                          = 0x80004028
	CO_E_ASYNC_WORK_REJECTED                      = 0x80004029
	CO_E_SERVER_INIT_TIMEOUT                      = 0x8000402A
	CO_E_NO_SECCTX_IN_ACTIVATE                    = 0x8000402B
	CO_E_TRACKER_CONFIG                           = 0x80004030
	CO_E_THREADPOOL_CONFIG                        = 0x80004031
	CO_E_SXS_CONFIG                               = 0x80004032
	CO_E_MALFORMED_SPN                            = 0x80004033
)

const (
	S_OK    = 0
	S_FALSE = 1
)

const (
	OLE_E_OLEVERB                                     = 0x80040000
	OLE_E_ADVF                                        = 0x80040001
	OLE_E_ENUM_NOMORE                                 = 0x80040002
	OLE_E_ADVISENOTSUPPORTED                          = 0x80040003
	OLE_E_NOCONNECTION                                = 0x80040004
	OLE_E_NOTRUNNING                                  = 0x80040005
	OLE_E_NOCACHE                                     = 0x80040006
	OLE_E_BLANK                                       = 0x80040007
	OLE_E_CLASSDIFF                                   = 0x80040008
	OLE_E_CANT_GETMONIKER                             = 0x80040009
	OLE_E_CANT_BINDTOSOURCE                           = 0x8004000A
	OLE_E_STATIC                                      = 0x8004000B
	OLE_E_PROMPTSAVECANCELLED                         = 0x8004000C
	OLE_E_INVALIDRECT                                 = 0x8004000D
	OLE_E_WRONGCOMPOBJ                                = 0x8004000E
	OLE_E_INVALIDHWND                                 = 0x8004000F
	OLE_E_NOT_INPLACEACTIVE                           = 0x80040010
	OLE_E_CANTCONVERT                                 = 0x80040011
	OLE_E_NOSTORAGE                                   = 0x80040012
	DV_E_FORMATETC                                    = 0x80040064
	DV_E_DVTARGETDEVICE                               = 0x80040065
	DV_E_STGMEDIUM                                    = 0x80040066
	DV_E_STATDATA                                     = 0x80040067
	DV_E_LINDEX                                       = 0x80040068
	DV_E_TYMED                                        = 0x80040069
	DV_E_CLIPFORMAT                                   = 0x8004006A
	DV_E_DVASPECT                                     = 0x8004006B
	DV_E_DVTARGETDEVICE_SIZE                          = 0x8004006C
	DV_E_NOIVIEWOBJECT                                = 0x8004006D
	DRAGDROP_E_NOTREGISTERED                          = 0x80040100
	DRAGDROP_E_ALREADYREGISTERED                      = 0x80040101
	DRAGDROP_E_INVALIDHWND                            = 0x80040102
	CLASS_E_NOAGGREGATION                             = 0x80040110
	CLASS_E_CLASSNOTAVAILABLE                         = 0x80040111
	CLASS_E_NOTLICENSED                               = 0x80040112
	VIEW_E_DRAW                                       = 0x80040140
	REGDB_E_READREGDB                                 = 0x80040150
	REGDB_E_WRITEREGDB                                = 0x80040151
	REGDB_E_KEYMISSING                                = 0x80040152
	REGDB_E_INVALIDVALUE                              = 0x80040153
	REGDB_E_CLASSNOTREG                               = 0x80040154
	REGDB_E_IIDNOTREG                                 = 0x80040155
	REGDB_E_BADTHREADINGMODEL                         = 0x80040156
	CAT_E_CATIDNOEXIST                                = 0x80040160
	CAT_E_NODESCRIPTION                               = 0x80040161
	CS_E_PACKAGE_NOTFOUND                             = 0x80040164
	CS_E_NOT_DELETABLE                                = 0x80040165
	CS_E_CLASS_NOTFOUND                               = 0x80040166
	CS_E_INVALID_VERSION                              = 0x80040167
	CS_E_NO_CLASSSTORE                                = 0x80040168
	CS_E_OBJECT_NOTFOUND                              = 0x80040169
	CS_E_OBJECT_ALREADY_EXISTS                        = 0x8004016A
	CS_E_INVALID_PATH                                 = 0x8004016B
	CS_E_NETWORK_ERROR                                = 0x8004016C
	CS_E_ADMIN_LIMIT_EXCEEDED                         = 0x8004016D
	CS_E_SCHEMA_MISMATCH                              = 0x8004016E
	CS_E_INTERNAL_ERROR                               = 0x8004016F
	CACHE_E_NOCACHE_UPDATED                           = 0x80040170
	OLEOBJ_E_NOVERBS                                  = 0x80040180
	OLEOBJ_E_INVALIDVERB                              = 0x80040181
	INPLACE_E_NOTUNDOABLE                             = 0x800401A0
	INPLACE_E_NOTOOLSPACE                             = 0x800401A1
	CONVERT10_E_OLESTREAM_GET                         = 0x800401C0
	CONVERT10_E_OLESTREAM_PUT                         = 0x800401C1
	CONVERT10_E_OLESTREAM_FMT                         = 0x800401C2
	CONVERT10_E_OLESTREAM_BITMAP_TO_DIB               = 0x800401C3
	CONVERT10_E_STG_FMT                               = 0x800401C4
	CONVERT10_E_STG_NO_STD_STREAM                     = 0x800401C5
	CONVERT10_E_STG_DIB_TO_BITMAP                     = 0x800401C6
	CLIPBRD_E_CANT_OPEN                               = 0x800401D0
	CLIPBRD_E_CANT_EMPTY                              = 0x800401D1
	CLIPBRD_E_CANT_SET                                = 0x800401D2
	CLIPBRD_E_BAD_DATA                                = 0x800401D3
	CLIPBRD_E_CANT_CLOSE                              = 0x800401D4
	MK_E_CONNECTMANUALLY                              = 0x800401E0
	MK_E_EXCEEDEDDEADLINE                             = 0x800401E1
	MK_E_NEEDGENERIC                                  = 0x800401E2
	MK_E_UNAVAILABLE                                  = 0x800401E3
	MK_E_SYNTAX                                       = 0x800401E4
	MK_E_NOOBJECT                                     = 0x800401E5
	MK_E_INVALIDEXTENSION                             = 0x800401E6
	MK_E_INTERMEDIATEINTERFACENOTSUPPORTED            = 0x800401E7
	MK_E_NOTBINDABLE                                  = 0x800401E8
	MK_E_NOTBOUND                                     = 0x800401E9
	MK_E_CANTOPENFILE                                 = 0x800401EA
	MK_E_MUSTBOTHERUSER                               = 0x800401EB
	MK_E_NOINVERSE                                    = 0x800401EC
	MK_E_NOSTORAGE                                    = 0x800401ED
	MK_E_NOPREFIX                                     = 0x800401EE
	MK_E_ENUMERATION_FAILED                           = 0x800401EF
	CO_E_NOTINITIALIZED                               = 0x800401F0
	CO_E_ALREADYINITIALIZED                           = 0x800401F1
	CO_E_CANTDETERMINECLASS                           = 0x800401F2
	CO_E_CLASSSTRING                                  = 0x800401F3
	CO_E_IIDSTRING                                    = 0x800401F4
	CO_E_APPNOTFOUND                                  = 0x800401F5
	CO_E_APPSINGLEUSE                                 = 0x800401F6
	CO_E_ERRORINAPP                                   = 0x800401F7
	CO_E_DLLNOTFOUND                                  = 0x800401F8
	CO_E_ERRORINDLL                                   = 0x800401F9
	CO_E_WRONGOSFORAPP                                = 0x800401FA
	CO_E_OBJNOTREG                                    = 0x800401FB
	CO_E_OBJISREG                                     = 0x800401FC
	CO_E_OBJNOTCONNECTED                              = 0x800401FD
	CO_E_APPDIDNTREG                                  = 0x800401FE
	CO_E_RELEASED                                     = 0x800401FF
	EVENT_S_SOME_SUBSCRIBERS_FAILED                   = 0x00040200
	EVENT_E_ALL_SUBSCRIBERS_FAILED                    = 0x80040201
	EVENT_S_NOSUBSCRIBERS                             = 0x00040202
	EVENT_E_QUERYSYNTAX                               = 0x80040203
	EVENT_E_QUERYFIELD                                = 0x80040204
	EVENT_E_INTERNALEXCEPTION                         = 0x80040205
	EVENT_E_INTERNALERROR                             = 0x80040206
	EVENT_E_INVALID_PER_USER_SID                      = 0x80040207
	EVENT_E_USER_EXCEPTION                            = 0x80040208
	EVENT_E_TOO_MANY_METHODS                          = 0x80040209
	EVENT_E_MISSING_EVENTCLASS                        = 0x8004020A
	EVENT_E_NOT_ALL_REMOVED                           = 0x8004020B
	EVENT_E_COMPLUS_NOT_INSTALLED                     = 0x8004020C
	EVENT_E_CANT_MODIFY_OR_DELETE_UNCONFIGURED_OBJECT = 0x8004020D
	EVENT_E_CANT_MODIFY_OR_DELETE_CONFIGURED_OBJECT   = 0x8004020E
	EVENT_E_INVALID_EVENT_CLASS_PARTITION             = 0x8004020F
	EVENT_E_PER_USER_SID_NOT_LOGGED_ON                = 0x80040210
	XACT_E_ALREADYOTHERSINGLEPHASE                    = 0x8004D000
	XACT_E_CANTRETAIN                                 = 0x8004D001
	XACT_E_COMMITFAILED                               = 0x8004D002
	XACT_E_COMMITPREVENTED                            = 0x8004D003
	XACT_E_HEURISTICABORT                             = 0x8004D004
	XACT_E_HEURISTICCOMMIT                            = 0x8004D005
	XACT_E_HEURISTICDAMAGE                            = 0x8004D006
	XACT_E_HEURISTICDANGER                            = 0x8004D007
	XACT_E_ISOLATIONLEVEL                             = 0x8004D008
	XACT_E_NOASYNC                                    = 0x8004D009
	XACT_E_NOENLIST                                   = 0x8004D00A
	XACT_E_NOISORETAIN                                = 0x8004D00B
	XACT_E_NORESOURCE                                 = 0x8004D00C
	XACT_E_NOTCURRENT                                 = 0x8004D00D
	XACT_E_NOTRANSACTION                              = 0x8004D00E
	XACT_E_NOTSUPPORTED                               = 0x8004D00F
	XACT_E_UNKNOWNRMGRID                              = 0x8004D010
	XACT_E_WRONGSTATE                                 = 0x8004D011
	XACT_E_WRONGUOW                                   = 0x8004D012
	XACT_E_XTIONEXISTS                                = 0x8004D013
	XACT_E_NOIMPORTOBJECT                             = 0x8004D014
	XACT_E_INVALIDCOOKIE                              = 0x8004D015
	XACT_E_INDOUBT                                    = 0x8004D016
	XACT_E_NOTIMEOUT                                  = 0x8004D017
	XACT_E_ALREADYINPROGRESS                          = 0x8004D018
	XACT_E_ABORTED                                    = 0x8004D019
	XACT_E_LOGFULL                                    = 0x8004D01A
	XACT_E_TMNOTAVAILABLE                             = 0x8004D01B
	XACT_E_CONNECTION_DOWN                            = 0x8004D01C
	XACT_E_CONNECTION_DENIED                          = 0x8004D01D
	XACT_E_REENLISTTIMEOUT                            = 0x8004D01E
	XACT_E_TIP_CONNECT_FAILED                         = 0x8004D01F
	XACT_E_TIP_PROTOCOL_ERROR                         = 0x8004D020
	XACT_E_TIP_PULL_FAILED                            = 0x8004D021
	XACT_E_DEST_TMNOTAVAILABLE                        = 0x8004D022
	XACT_E_TIP_DISABLED                               = 0x8004D023
	XACT_E_NETWORK_TX_DISABLED                        = 0x8004D024
	XACT_E_PARTNER_NETWORK_TX_DISABLED                = 0x8004D025
	XACT_E_XA_TX_DISABLED                             = 0x8004D026
	XACT_E_UNABLE_TO_READ_DTC_CONFIG                  = 0x8004D027
	XACT_E_UNABLE_TO_LOAD_DTC_PROXY                   = 0x8004D028
	XACT_E_ABORTING                                   = 0x8004D029
	XACT_E_PUSH_COMM_FAILURE                          = 0x8004D02A
	XACT_E_PULL_COMM_FAILURE                          = 0x8004D02B
	XACT_E_LU_TX_DISABLED                             = 0x8004D02C
	XACT_E_CLERKNOTFOUND                              = 0x8004D080
	XACT_E_CLERKEXISTS                                = 0x8004D081
	XACT_E_RECOVERYINPROGRESS                         = 0x8004D082
	XACT_E_TRANSACTIONCLOSED                          = 0x8004D083
	XACT_E_INVALIDLSN                                 = 0x8004D084
	XACT_E_REPLAYREQUEST                              = 0x8004D085
	XACT_S_ASYNC                                      = 0x0004D000
	XACT_S_DEFECT                                     = 0x0004D001
	XACT_S_READONLY                                   = 0x0004D002
	XACT_S_SOMENORETAIN                               = 0x0004D003
	XACT_S_OKINFORM                                   = 0x0004D004
	XACT_S_MADECHANGESCONTENT                         = 0x0004D005
	XACT_S_MADECHANGESINFORM                          = 0x0004D006
	XACT_S_ALLNORETAIN                                = 0x0004D007
	XACT_S_ABORTING                                   = 0x0004D008
	XACT_S_SINGLEPHASE                                = 0x0004D009
	XACT_S_LOCALLY_OK                                 = 0x0004D00A
	XACT_S_LASTRESOURCEMANAGER                        = 0x0004D010
	CONTEXT_E_ABORTED                                 = 0x8004E002
	CONTEXT_E_ABORTING                                = 0x8004E003
	CONTEXT_E_NOCONTEXT                               = 0x8004E004
	CONTEXT_E_WOULD_DEADLOCK                          = 0x8004E005
	CONTEXT_E_SYNCH_TIMEOUT                           = 0x8004E006
	CONTEXT_E_OLDREF                                  = 0x8004E007
	CONTEXT_E_ROLENOTFOUND                            = 0x8004E00C
	CONTEXT_E_TMNOTAVAILABLE                          = 0x8004E00F
	CO_E_ACTIVATIONFAILED                             = 0x8004E021
	CO_E_ACTIVATIONFAILED_EVENTLOGGED                 = 0x8004E022
	CO_E_ACTIVATIONFAILED_CATALOGERROR                = 0x8004E023
	CO_E_ACTIVATIONFAILED_TIMEOUT                     = 0x8004E024
	CO_E_INITIALIZATIONFAILED                         = 0x8004E025
	CONTEXT_E_NOJIT                                   = 0x8004E026
	CONTEXT_E_NOTRANSACTION                           = 0x8004E027
	CO_E_THREADINGMODEL_CHANGED                       = 0x8004E028
	CO_E_NOIISINTRINSICS                              = 0x8004E029
	CO_E_NOCOOKIES                                    = 0x8004E02A
	CO_E_DBERROR                                      = 0x8004E02B
	CO_E_NOTPOOLED                                    = 0x8004E02C
	CO_E_NOTCONSTRUCTED                               = 0x8004E02D
	CO_E_NOSYNCHRONIZATION                            = 0x8004E02E
	CO_E_ISOLEVELMISMATCH                             = 0x8004E02F
	CO_E_CALL_OUT_OF_TX_SCOPE_NOT_ALLOWED             = 0x8004E030
	CO_E_EXIT_TRANSACTION_SCOPE_NOT_CALLED            = 0x8004E031
	OLE_S_USEREG                                      = 0x00040000
	OLE_S_STATIC                                      = 0x00040001
	OLE_S_MAC_CLIPFORMAT                              = 0x00040002
	DRAGDROP_S_DROP                                   = 0x00040100
	DRAGDROP_S_CANCEL                                 = 0x00040101
	DRAGDROP_S_USEDEFAULTCURSORS                      = 0x00040102
	DATA_S_SAMEFORMATETC                              = 0x00040130
	VIEW_S_ALREADY_FROZEN                             = 0x00040140
	CACHE_S_FORMATETC_NOTSUPPORTED                    = 0x00040170
	CACHE_S_SAMECACHE                                 = 0x00040171
	CACHE_S_SAMECACHES_NOTUPDATED                     = 0x00040172
	OLEOBJ_S_INVALIDVERB                              = 0x00040180
	OLEOBJ_S_CANNOT_DOVERB_NOW                        = 0x00040181
	OLEOBJ_S_INVALIDHWND                              = 0x00040182
	INPLACE_S_TRUNCATED                               = 0x000401A0
	CONVERT10_S_NO_PRESENTATION                       = 0x000401C0
	MK_S_REDUCED_TO_SELF                              = 0x000401E2
	MK_S_ME                                           = 0x000401E4
	MK_S_HIM                                          = 0x000401E5
	MK_S_US                                           = 0x000401E6
	MK_S_MONIKERALREADYREGISTERED                     = 0x000401E7
	SCHED_S_TASK_READY                                = 0x00041300
	SCHED_S_TASK_RUNNING                              = 0x00041301
	SCHED_S_TASK_DISABLED                             = 0x00041302
	SCHED_S_TASK_HAS_NOT_RUN                          = 0x00041303
	SCHED_S_TASK_NO_MORE_RUNS                         = 0x00041304
	SCHED_S_TASK_NOT_SCHEDULED                        = 0x00041305
	SCHED_S_TASK_TERMINATED                           = 0x00041306
	SCHED_S_TASK_NO_VALID_TRIGGERS                    = 0x00041307
	SCHED_S_EVENT_TRIGGER                             = 0x00041308
	SCHED_E_TRIGGER_NOT_FOUND                         = 0x80041309
	SCHED_E_TASK_NOT_READY                            = 0x8004130A
	SCHED_E_TASK_NOT_RUNNING                          = 0x8004130B
	SCHED_E_SERVICE_NOT_INSTALLED                     = 0x8004130C
	SCHED_E_CANNOT_OPEN_TASK                          = 0x8004130D
	SCHED_E_INVALID_TASK                              = 0x8004130E
	SCHED_E_ACCOUNT_INFORMATION_SET_NOT               = 0x8004130F
	SCHED_E_ACCOUNT_NAME_NOT_FOUND                    = 0x80041310
	SCHED_E_ACCOUNT_DBASE_CORRUPT                     = 0x80041311
	SCHED_E_NO_SECURITY_SERVICES                      = 0x80041312
	SCHED_E_UNKNOWN_OBJECT_VERSION                    = 0x80041313
	SCHED_E_UNSUPPORTED_ACCOUNT_OPTION                = 0x80041314
	SCHED_E_SERVICE_NOT_RUNNING                       = 0x80041315
	SCHED_E_UNEXPECTEDNODE                            = 0x80041316
	SCHED_E_NAMESPACE                                 = 0x80041317
	SCHED_E_INVALIDVALUE                              = 0x80041318
	SCHED_E_MISSINGNODE                               = 0x80041319
	SCHED_E_MALFORMEDXML                              = 0x8004131A
	SCHED_S_SOME_TRIGGERS_FAILED                      = 0x0004131B
	SCHED_S_BATCH_LOGON_PROBLEM                       = 0x0004131C
	SCHED_E_TOO_MANY_NODES                            = 0x8004131D
	SCHED_E_PAST_END_BOUNDARY                         = 0x8004131E
	SCHED_E_ALREADY_RUNNING                           = 0x8004131F
	SCHED_E_USER_NOT_LOGGED_ON                        = 0x80041320
	SCHED_E_INVALID_TASK_HASH                         = 0x80041321
	SCHED_E_SERVICE_NOT_AVAILABLE                     = 0x80041322
	SCHED_E_SERVICE_TOO_BUSY                          = 0x80041323
	SCHED_E_TASK_ATTEMPTED                            = 0x80041324
	SCHED_S_TASK_QUEUED                               = 0x00041325
	SCHED_E_TASK_DISABLED                             = 0x80041326
	SCHED_E_TASK_NOT_V1_COMPAT                        = 0x80041327
	SCHED_E_START_ON_DEMAND                           = 0x80041328
)

const (
	CO_E_CLASS_CREATE_FAILED     = 0x80080001
	CO_E_SCM_ERROR               = 0x80080002
	CO_E_SCM_RPC_FAILURE         = 0x80080003
	CO_E_BAD_PATH                = 0x80080004
	CO_E_SERVER_EXEC_FAILURE     = 0x80080005
	CO_E_OBJSRV_RPC_FAILURE      = 0x80080006
	MK_E_NO_NORMALIZED           = 0x80080007
	CO_E_SERVER_STOPPING         = 0x80080008
	MEM_E_INVALID_ROOT           = 0x80080009
	MEM_E_INVALID_LINK           = 0x80080010
	MEM_E_INVALID_SIZE           = 0x80080011
	CO_S_NOTALLINTERFACES        = 0x00080012
	CO_S_MACHINENAMENOTFOUND     = 0x00080013
	CO_E_MISSING_DISPLAYNAME     = 0x80080015
	CO_E_RUNAS_VALUE_MUST_BE_AAA = 0x80080016
	CO_E_ELEVATION_DISABLED      = 0x80080017
)

const (
	DISP_E_UNKNOWNINTERFACE        = 0x80020001
	DISP_E_MEMBERNOTFOUND          = 0x80020003
	DISP_E_PARAMNOTFOUND           = 0x80020004
	DISP_E_TYPEMISMATCH            = 0x80020005
	DISP_E_UNKNOWNNAME             = 0x80020006
	DISP_E_NONAMEDARGS             = 0x80020007
	DISP_E_BADVARTYPE              = 0x80020008
	DISP_E_EXCEPTION               = 0x80020009
	DISP_E_OVERFLOW                = 0x8002000A
	DISP_E_BADINDEX                = 0x8002000B
	DISP_E_UNKNOWNLCID             = 0x8002000C
	DISP_E_ARRAYISLOCKED           = 0x8002000D
	DISP_E_BADPARAMCOUNT           = 0x8002000E
	DISP_E_PARAMNOTOPTIONAL        = 0x8002000F
	DISP_E_BADCALLEE               = 0x80020010
	DISP_E_NOTACOLLECTION          = 0x80020011
	DISP_E_DIVBYZERO               = 0x80020012
	DISP_E_BUFFERTOOSMALL          = 0x80020013
	TYPE_E_BUFFERTOOSMALL          = 0x80028016
	TYPE_E_FIELDNOTFOUND           = 0x80028017
	TYPE_E_INVDATAREAD             = 0x80028018
	TYPE_E_UNSUPFORMAT             = 0x80028019
	TYPE_E_REGISTRYACCESS          = 0x8002801C
	TYPE_E_LIBNOTREGISTERED        = 0x8002801D
	TYPE_E_UNDEFINEDTYPE           = 0x80028027
	TYPE_E_QUALIFIEDNAMEDISALLOWED = 0x80028028
	TYPE_E_INVALIDSTATE            = 0x80028029
	TYPE_E_WRONGTYPEKIND           = 0x8002802A
	TYPE_E_ELEMENTNOTFOUND         = 0x8002802B
	TYPE_E_AMBIGUOUSNAME           = 0x8002802C
	TYPE_E_NAMECONFLICT            = 0x8002802D
	TYPE_E_UNKNOWNLCID             = 0x8002802E
	TYPE_E_DLLFUNCTIONNOTFOUND     = 0x8002802F
	TYPE_E_BADMODULEKIND           = 0x800288BD
	TYPE_E_SIZETOOBIG              = 0x800288C5
	TYPE_E_DUPLICATEID             = 0x800288C6
	TYPE_E_INVALIDID               = 0x800288CF
	TYPE_E_TYPEMISMATCH            = 0x80028CA0
	TYPE_E_OUTOFBOUNDS             = 0x80028CA1
	TYPE_E_IOERROR                 = 0x80028CA2
	TYPE_E_CANTCREATETMPFILE       = 0x80028CA3
	TYPE_E_CANTLOADLIBRARY         = 0x80028C4A
	TYPE_E_INCONSISTENTPROPFUNCS   = 0x80028C83
	TYPE_E_CIRCULARTYPE            = 0x80028C84
)

const (
	RPC_E_CALL_REJECTED               = 0x80010001
	RPC_E_CALL_CANCELED               = 0x80010002
	RPC_E_CANTPOST_INSENDCALL         = 0x80010003
	RPC_E_CANTCALLOUT_INASYNCCALL     = 0x80010004
	RPC_E_CANTCALLOUT_INEXTERNALCALL  = 0x80010005
	RPC_E_CONNECTION_TERMINATED       = 0x80010006
	RPC_E_SERVER_DIED                 = 0x80010007
	RPC_E_CLIENT_DIED                 = 0x80010008
	RPC_E_INVALID_DATAPACKET          = 0x80010009
	RPC_E_CANTTRANSMIT_CALL           = 0x8001000A
	RPC_E_CLIENT_CANTMARSHAL_DATA     = 0x8001000B
	RPC_E_CLIENT_CANTUNMARSHAL_DATA   = 0x8001000C
	RPC_E_SERVER_CANTMARSHAL_DATA     = 0x8001000D
	RPC_E_SERVER_CANTUNMARSHAL_DATA   = 0x8001000E
	RPC_E_INVALID_DATA                = 0x8001000F
	RPC_E_INVALID_PARAMETER           = 0x80010010
	RPC_E_CANTCALLOUT_AGAIN           = 0x80010011
	RPC_E_SERVER_DIED_DNE             = 0x80010012
	RPC_E_SYS_CALL_FAILED             = 0x80010100
	RPC_E_OUT_OF_RESOURCES            = 0x80010101
	RPC_E_ATTEMPTED_MULTITHREAD       = 0x80010102
	RPC_E_NOT_REGISTERED              = 0x80010103
	RPC_E_FAULT                       = 0x80010104
	RPC_E_SERVERFAULT                 = 0x80010105
	RPC_E_CHANGED_MODE                = 0x80010106
	RPC_E_INVALIDMETHOD               = 0x80010107
	RPC_E_DISCONNECTED                = 0x80010108
	RPC_E_RETRY                       = 0x80010109
	RPC_E_SERVERCALL_RETRYLATER       = 0x8001010A
	RPC_E_SERVERCALL_REJECTED         = 0x8001010B
	RPC_E_INVALID_CALLDATA            = 0x8001010C
	RPC_E_CANTCALLOUT_ININPUTSYNCCALL = 0x8001010D
	RPC_E_WRONG_THREAD                = 0x8001010E
	RPC_E_THREAD_NOT_INIT             = 0x8001010F
	RPC_E_VERSION_MISMATCH            = 0x80010110
	RPC_E_INVALID_HEADER              = 0x80010111
	RPC_E_INVALID_EXTENSION           = 0x80010112
	RPC_E_INVALID_IPID                = 0x80010113
	RPC_E_INVALID_OBJECT              = 0x80010114
	RPC_S_CALLPENDING                 = 0x80010115
	RPC_S_WAITONTIMER                 = 0x80010116
	RPC_E_CALL_COMPLETE               = 0x80010117
	RPC_E_UNSECURE_CALL               = 0x80010118
	RPC_E_TOO_LATE                    = 0x80010119
	RPC_E_NO_GOOD_SECURITY_PACKAGES   = 0x8001011A
	RPC_E_ACCESS_DENIED               = 0x8001011B
	RPC_E_REMOTE_DISABLED             = 0x8001011C
	RPC_E_INVALID_OBJREF              = 0x8001011D
	RPC_E_NO_CONTEXT                  = 0x8001011E
	RPC_E_TIMEOUT                     = 0x8001011F
	RPC_E_NO_SYNC                     = 0x80010120
	RPC_E_FULLSIC_REQUIRED            = 0x80010121
	RPC_E_INVALID_STD_NAME            = 0x80010122
	CO_E_FAILEDTOIMPERSONATE          = 0x80010123
	CO_E_FAILEDTOGETSECCTX            = 0x80010124
	CO_E_FAILEDTOOPENTHREADTOKEN      = 0x80010125
	CO_E_FAILEDTOGETTOKENINFO         = 0x80010126
	CO_E_TRUSTEEDOESNTMATCHCLIENT     = 0x80010127
	CO_E_FAILEDTOQUERYCLIENTBLANKET   = 0x80010128
	CO_E_FAILEDTOSETDACL              = 0x80010129
	CO_E_ACCESSCHECKFAILED            = 0x8001012A
	CO_E_NETACCESSAPIFAILED           = 0x8001012B
	CO_E_WRONGTRUSTEENAMESYNTAX       = 0x8001012C
	CO_E_INVALIDSID                   = 0x8001012D
	CO_E_CONVERSIONFAILED             = 0x8001012E
	CO_E_NOMATCHINGSIDFOUND           = 0x8001012F
	CO_E_LOOKUPACCSIDFAILED           = 0x80010130
	CO_E_NOMATCHINGNAMEFOUND          = 0x80010131
	CO_E_LOOKUPACCNAMEFAILED          = 0x80010132
	CO_E_SETSERLHNDLFAILED            = 0x80010133
	CO_E_FAILEDTOGETWINDIR            = 0x80010134
	CO_E_PATHTOOLONG                  = 0x80010135
	CO_E_FAILEDTOGENUUID              = 0x80010136
	CO_E_FAILEDTOCREATEFILE           = 0x80010137
	CO_E_FAILEDTOCLOSEHANDLE          = 0x80010138
	CO_E_EXCEEDSYSACLLIMIT            = 0x80010139
	CO_E_ACESINWRONGORDER             = 0x8001013A
	CO_E_INCOMPATIBLESTREAMVERSION    = 0x8001013B
	CO_E_FAILEDTOOPENPROCESSTOKEN     = 0x8001013C
	CO_E_DECODEFAILED                 = 0x8001013D
	CO_E_ACNOTINITIALIZED             = 0x8001013E
	CO_E_CANCEL_DISABLED              = 0x8001013F
	RPC_E_UNEXPECTED                  = 0x8001FFFF
)
