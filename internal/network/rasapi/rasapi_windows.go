package rasapi

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	RAS_MaxAreaCode              = 10  //ok
	RAS_MaxPhoneNumber           = 128 //ok
	MAX_PATH                     = 260 //ok
	RAS_MaxDeviceType            = 16  //ok
	RAS_MaxDeviceName            = 128 //ok
	RAS_MaxPadType               = 32  //ok
	RAS_MaxX25Address            = 200 //ok
	RAS_MaxFacilities            = 200 //ok
	RAS_MaxUserData              = 200 //ok
	RAS_MaxDnsSuffix             = 256 //ok
	RAS_MaxEntryName             = 256 //ok
	RASEO_UseCountryAndAreaCodes = 0x00000001
	RASEO_SpecificIpAddr         = 0x00000002
	RASEO_SpecificNameServers    = 0x00000004
	RASEO_IpHeaderCompression    = 0x00000008
	RASEO_RemoteDefaultGateway   = 0x00000010
	RASEO_DisableLcpExtensions   = 0x00000020
	RASEO_TerminalBeforeDial     = 0x00000040
	RASEO_TerminalAfterDial      = 0x00000080
	RASEO_ModemLights            = 0x00000100
	RASEO_SwCompression          = 0x00000200
	RASEO_RequireEncryptedPw     = 0x00000400
	RASEO_RequireMsEncryptedPw   = 0x00000800
	RASEO_RequireDataEncryption  = 0x00001000
	RASEO_NetworkLogon           = 0x00002000
	RASEO_UseLogonCredentials    = 0x00004000
	RASEO_PromoteAlternates      = 0x00008000
	RASEO_SecureLocalFiles       = 0x00010000
	RASEO_RequireEAP             = 0x00020000
	RASEO_RequirePAP             = 0x00040000
	RASEO_RequireSPAP            = 0x00080000
	RASEO_Custom                 = 0x00100000
	RASEO_PreviewPhoneNumber     = 0x00200000
	RASEO_SharedPhoneNumbers     = 0x00800000
	RASEO_PreviewUserPw          = 0x01000000
	RASEO_PreviewDomain          = 0x02000000
	RASEO_ShowDialingProgress    = 0x04000000
	RASEO_RequireCHAP            = 0x08000000
	RASEO_RequireMsCHAP          = 0x10000000
	RASEO_RequireMsCHAP2         = 0x20000000
	RASEO_RequireW95MSCHAP       = 0x40000000
	RASEO_CustomScript           = 0x80000000
	RASNP_Ip                     = 0x00000004
	RASFP_Ppp                    = 0x00000001
	RASDT_PPPoE                  = "PPPoE"
)

type RASIPADDR struct {
	a byte
	b byte
	c byte
	d byte
}
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}
type RASIPV6ADDR struct {
	UCHAR [16]byte
	UWord [8]uint16
}
type RASDIALPARAMSW struct {
	DwSize           uint32
	SzEntryName      [256 + 1]uint16
	SzPhoneNumber    [128 + 1]uint16
	SzCallbackNumber [128 + 1]uint16
	SzUserName       [257]uint16
	SzPassword       [257]uint16
	SzDomain         [16]uint16
	DwSubEntry       uint32
	DwCallbackId     uint64
	//DwIfIndex uint32
}
type RASENTRY struct {
	DwSize                     uint32                         //DWORD dwSize;
	DwfOptions                 uint32                         //DWORD dwfOptions;
	DwCountryID                uint32                         //DWORD dwCountryID;
	DwCountryCode              uint32                         //DWORD dwCountryCode;
	SzAreaCode                 [RAS_MaxAreaCode + 1]uint16    //TCHAR szAreaCode[RAS_MaxAreaCode + 1];
	SzLocalPhoneNumber         [RAS_MaxPhoneNumber + 1]uint16 //TCHAR szLocalPhoneNumber[RAS_MaxPhoneNumber + 1];
	DwAlternateOffset          uint32                         // DWORD dwAlternateOffset;
	Ipaddr                     RASIPADDR                      //RASIPADDR ipaddr
	IpaddrDns                  RASIPADDR                      //RASIPADDR ipaddrDns
	IpaddrDnsAlt               RASIPADDR                      //RASIPADDR ipaddrDnsAlt
	IpaddrWins                 RASIPADDR                      // RASIPADDR ipaddrWins
	IpaddrWinsAlt              RASIPADDR                      //RASIPADDR ipaddrWinsAlt
	DwFrameSize                uint32                         //DWORD dwFrameSize;
	DwfNetProtocols            uint32                         // DWORD dwfNetProtocols;
	DwFramingProtocol          uint32                         // DWORD dwFramingProtocol;
	SzScript                   [MAX_PATH]uint16               //TCHAR szScript[MAX_PATH];
	SzAutodialDll              [MAX_PATH]uint16               // TCHAR szAutodialDll[MAX_PATH];
	SzAutodialFunc             [MAX_PATH]uint16               // TCHAR szAutodialFunc[MAX_PATH];
	SzDeviceType               [RAS_MaxDeviceType + 1]uint16  // TCHAR szDeviceType[RAS_MaxDeviceType + 1];
	SzDeviceName               [RAS_MaxDeviceName + 1]uint16  // TCHAR szDeviceName[RAS_MaxDeviceName + 1];
	SzX25PadType               [RAS_MaxPadType + 1]uint16     // TCHAR szX25PadType[RAS_MaxPadType + 1];
	SzX25Address               [RAS_MaxX25Address + 1]uint16  // TCHAR szX25Address[RAS_MaxX25Address + 1];
	SzX25Facilities            [RAS_MaxFacilities + 1]uint16  // TCHAR szX25Facilities[RAS_MaxFacilities + 1];
	SzX25UserData              [RAS_MaxUserData + 1]uint16    //TCHAR szX25UserData[RAS_MaxUserData + 1];
	DwChannels                 uint32                         //DWORD dwChannels;
	DwReserved1                uint32                         //DWORD dwReserved1;
	DwReserved2                uint32                         // DWORD dwReserved2;
	DwSubEntries               uint32                         // DWORD dwSubEntries;
	DwDialMode                 uint32                         // DWORD dwDialMode;
	DwDialExtraPercent         uint32                         // DWORD dwDialExtraPercent;
	DwDialExtraSampleSeconds   uint32                         // DWORD dwDialExtraSampleSeconds;
	DwHangUpExtraPercent       uint32                         // DWORD dwHangUpExtraPercent;
	DwHangUpExtraSampleSeconds uint32                         // DWORD dwHangUpExtraSampleSeconds;
	DwIdleDisconnectSeconds    uint32                         // DWORD dwIdleDisconnectSeconds;
	DwType                     uint32                         // DWORD dwType;
	DwEncryptionType           uint32                         // DWORD dwEncryptionType;
	DwCustomAuthKey            uint32                         // DWORD dwCustomAuthKey;
	GuidId                     GUID                           //GUID guidId
	SzCustomDialDll            [MAX_PATH]uint16               // TCHAR szCustomDialDll[MAX_PATH];
	DwVpnStrategy              uint32                         // DWORD dwVpnStrategy;
	DwfOptions2                uint32                         // DWORD dwfOptions2;
	DwfOptions3                uint32                         // DWORD dwfOptions3;
	SzDnsSuffix                [RAS_MaxDnsSuffix]uint16       // TCHAR szDnsSuffix[RAS_MaxDnsSuffix];
	DwTcpWindowSize            uint32                         // DWORD dwTcpWindowSize;
	SzPrerequisitePbk          [MAX_PATH]uint16               // TCHAR szPrerequisitePbk[MAX_PATH];
	SzPrerequisiteEntry        [RAS_MaxEntryName + 1]uint16   // TCHAR szPrerequisiteEntry[RAS_MaxEntryName + 1];
	DwRedialCount              uint32                         // DWORD dwRedialCount;
	DwRedialPause              uint32                         // DWORD dwRedialPause;
	Ipv6addrDns                RASIPV6ADDR                    //RASIPV6ADDR ipv6addrDns
	Ipv6addrDnsAlt             RASIPV6ADDR                    //RASIPV6ADDR ipv6addrDnsAlt
	DwIPv4InterfaceMetric      uint32                         // DWORD dwIPv4InterfaceMetric;
	DwIPv6InterfaceMetric      uint32                         // DWORD dwIPv6InterfaceMetric;
	Ipv6addr                   RASIPV6ADDR                    //RASIPV6ADDR ipv6addr
	DwIPv6PrefixLength         uint32                         // DWORD dwIPv6PrefixLength;
	DwNetworkOutageTime        uint32                         // DWORD dwNetworkOutageTime;
}
type RASCONNSTATUSW struct {
	DwSize        uint32
	Rasconnstate  uint32
	DwError       uint32
	SzDeviceType  [17]uint16
	SzDeviceName  [129]uint16
	SzPhoneNumber [129]uint16
}

var (
	rasapi32               *syscall.DLL
	rasSetEntryPropertiesW *syscall.Proc
	rasGetEntryPropertiesW *syscall.Proc
	rasDialW               *syscall.Proc
	rasGetConnectStatus    *syscall.Proc
	rasHangUpW             *syscall.Proc
)

func init() {
	rasapi32 = syscall.MustLoadDLL("Rasapi32.dll")
	rasSetEntryPropertiesW = rasapi32.MustFindProc("RasSetEntryPropertiesW")
	rasGetEntryPropertiesW = rasapi32.MustFindProc("RasGetEntryPropertiesW")
	rasDialW = rasapi32.MustFindProc("RasDialW")
	rasGetConnectStatus = rasapi32.MustFindProc("RasGetConnectStatusW")
	rasHangUpW = rasapi32.MustFindProc("RasHangUpW")
}

// 挂断拨号连接
func RasHangUpW(hconn uint32) bool {
	r1, r2, lastErr := rasDialW.Call(uintptr(hconn))
	fmt.Println(r1, r2, lastErr)
	return r1 == 0
}

// 取连接是否成功
func RasGetConnectStatusW(hconn uint32) int {
	r := RASCONNSTATUSW{}
	r.DwSize = uint32(unsafe.Sizeof(r))
	rasGetConnectStatus.Call(uintptr(hconn), uintptr(unsafe.Pointer(&r)))
	return int(r.Rasconnstate)
}

// 连接到VPN
func RasDialW(server, connName, name, pass string) (r1, r2 uintptr, hconn uint32, lastErr error) {
	p := RASDIALPARAMSW{}
	namesz := syscall.StringToUTF16(name)
	copy(p.SzUserName[0:len(namesz)], namesz)
	passsz := syscall.StringToUTF16(pass)
	copy(p.SzPassword[0:len(passsz)], passsz)
	connNamesz := syscall.StringToUTF16(connName)
	copy(p.SzEntryName[0:len(connNamesz)], connNamesz)
	serversz := syscall.StringToUTF16(server)
	copy(p.SzPhoneNumber[0:len(serversz)], serversz)
	p.DwSize = uint32(unsafe.Sizeof(p))
	r1, r2, lastErr = rasDialW.Call(0, 0, uintptr(unsafe.Pointer(&p)), uintptr(0), uintptr(0), uintptr(unsafe.Pointer(&hconn)))
	return
}

// 连接到VPN,指定电话薄路径
func RasDialW2(bookpath, server, connName, name, pass string) (r1, r2 uintptr, hconn uint32, lastErr error) {
	p := RASDIALPARAMSW{}
	namesz := syscall.StringToUTF16(name)
	copy(p.SzUserName[0:len(namesz)], namesz)
	passsz := syscall.StringToUTF16(pass)
	copy(p.SzPassword[0:len(passsz)], passsz)
	connNamesz := syscall.StringToUTF16(connName)
	copy(p.SzEntryName[0:len(connNamesz)], connNamesz)
	serversz := syscall.StringToUTF16(server)
	copy(p.SzPhoneNumber[0:len(serversz)], serversz)
	p.DwSize = uint32(unsafe.Sizeof(p))
	r1, r2, lastErr = rasDialW.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(bookpath))), uintptr(unsafe.Pointer(&p)), uintptr(0), uintptr(0), uintptr(unsafe.Pointer(&hconn)))
	return
}

// 取连接的信息
func RasGetEntryPropertiesW(connName string, r uintptr, dwBufferSize *uint32) (r1, r2 uintptr, lastErr error) {
	return rasGetEntryPropertiesW.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(connName))), r, uintptr(unsafe.Pointer(dwBufferSize)), 0, 0)
}

// 建立宽带连接
func RasSetEntryPropertiesW(connName string, deviceName string) (r1, r2 uintptr, lastErr error) {
	var dwBufferSize uint32 = 0
	r := RASENTRY{}
	RasGetEntryPropertiesW("", 0, &dwBufferSize)
	r.DwSize = dwBufferSize
	r.DwfOptions = 755302672
	r.DwCountryID = 0
	r.DwCountryCode = 0
	r.DwAlternateOffset = 0
	r.DwFrameSize = 0
	r.DwfNetProtocols = 12
	r.DwFramingProtocol = 1
	r.DwChannels = 0
	r.DwReserved1 = 0
	r.DwReserved2 = 0
	r.DwSubEntries = 1
	r.DwDialMode = 0
	r.DwDialExtraPercent = 0
	r.DwDialExtraSampleSeconds = 0
	r.DwHangUpExtraPercent = 0
	r.DwHangUpExtraSampleSeconds = 0
	r.DwIdleDisconnectSeconds = 0
	r.DwType = 5
	r.DwEncryptionType = 3
	r.DwCustomAuthKey = 0
	r.DwVpnStrategy = 0
	r.DwfOptions2 = 8559
	r.DwfOptions3 = 0
	r.DwTcpWindowSize = 0
	r.DwRedialCount = 3
	r.DwRedialPause = 60
	r.DwIPv4InterfaceMetric = 0
	r.DwIPv6InterfaceMetric = 0
	r.DwIPv6PrefixLength = 0
	r.DwNetworkOutageTime = 0
	dn := syscall.StringToUTF16(deviceName)
	copy(r.SzDeviceName[0:len(dn)], dn)
	dt := syscall.StringToUTF16(RASDT_PPPoE)
	copy(r.SzDeviceType[0:len(dt)], dt)
	return rasSetEntryPropertiesW.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(connName))), uintptr(unsafe.Pointer(&r)), uintptr(unsafe.Pointer(&dwBufferSize)), 0, 0)
}

// 建立VPN连接
func RasSetVPN(connName string) (r1, r2 uintptr, lastErr error) {
	var dwBufferSize uint32 = 0
	r := RASENTRY{}
	RasGetEntryPropertiesW("", 0, &dwBufferSize)
	r.DwSize = dwBufferSize
	r.DwfOptions = 654314512
	r.DwCountryID = 0
	r.DwCountryCode = 0
	r.DwAlternateOffset = 0
	r.DwFrameSize = 0
	r.DwfNetProtocols = 12
	r.DwFramingProtocol = 1
	r.DwChannels = 0
	r.DwReserved1 = 0
	r.DwReserved2 = 0
	r.DwSubEntries = 1
	r.DwDialMode = 0
	r.DwDialExtraPercent = 0
	r.DwDialExtraSampleSeconds = 0
	r.DwHangUpExtraPercent = 0
	r.DwHangUpExtraSampleSeconds = 0
	r.DwIdleDisconnectSeconds = 0
	r.DwType = 2
	r.DwEncryptionType = 3
	r.DwCustomAuthKey = 0
	r.DwVpnStrategy = 2
	r.DwfOptions2 = 33562884
	r.DwfOptions3 = 0
	r.DwTcpWindowSize = 0
	r.DwRedialCount = 3
	r.DwRedialPause = 60
	r.DwIPv4InterfaceMetric = 0
	r.DwIPv6InterfaceMetric = 0
	r.DwIPv6PrefixLength = 0
	r.DwNetworkOutageTime = 0
	dn := syscall.StringToUTF16("WAN Miniport (PPTP)")
	copy(r.SzDeviceName[0:len(dn)], dn)
	dt := syscall.StringToUTF16(RASDT_PPPoE)
	copy(r.SzDeviceType[0:len(dt)], dt)
	return rasSetEntryPropertiesW.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(connName))), uintptr(unsafe.Pointer(&r)), uintptr(unsafe.Pointer(&dwBufferSize)), 0, 0)
}

// 建立VPN连接,指定电话薄路径
func RasSetVPN2(connName, bookpath string) (r1, r2 uintptr, lastErr error) {
	var dwBufferSize uint32 = 0
	r := RASENTRY{}
	RasGetEntryPropertiesW("", 0, &dwBufferSize)
	r.DwSize = dwBufferSize
	r.DwfOptions = 654314512
	r.DwCountryID = 0
	r.DwCountryCode = 0
	r.DwAlternateOffset = 0
	r.DwFrameSize = 0
	r.DwfNetProtocols = 12
	r.DwFramingProtocol = 1
	r.DwChannels = 0
	r.DwReserved1 = 0
	r.DwReserved2 = 0
	r.DwSubEntries = 1
	r.DwDialMode = 0
	r.DwDialExtraPercent = 0
	r.DwDialExtraSampleSeconds = 0
	r.DwHangUpExtraPercent = 0
	r.DwHangUpExtraSampleSeconds = 0
	r.DwIdleDisconnectSeconds = 0
	r.DwType = 2
	r.DwEncryptionType = 3
	r.DwCustomAuthKey = 0
	r.DwVpnStrategy = 2
	r.DwfOptions2 = 33562884
	r.DwfOptions3 = 0
	r.DwTcpWindowSize = 0
	r.DwRedialCount = 3
	r.DwRedialPause = 60
	r.DwIPv4InterfaceMetric = 0
	r.DwIPv6InterfaceMetric = 0
	r.DwIPv6PrefixLength = 0
	r.DwNetworkOutageTime = 0
	dn := syscall.StringToUTF16("WAN Miniport (PPTP)")
	copy(r.SzDeviceName[0:len(dn)], dn)
	dt := syscall.StringToUTF16(RASDT_PPPoE)
	copy(r.SzDeviceType[0:len(dt)], dt)
	return rasSetEntryPropertiesW.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(bookpath))), uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(connName))), uintptr(unsafe.Pointer(&r)), uintptr(unsafe.Pointer(&dwBufferSize)), 0, 0)
}
