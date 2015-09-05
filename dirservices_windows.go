package adds

import (
	"syscall"
	"unsafe"
)

var (
	modnetapi32 = syscall.NewLazyDLL("netapi32.dll")

	procDsAddressToSiteNamesW             = modnetapi32.NewProc("DsAddressToSiteNamesW")
	procDsAddressToSiteNamesExW           = modnetapi32.NewProc("DsAddressToSiteNamesExW")
	procDsDeregisterDnsHostRecordsW       = modnetapi32.NewProc("DsDeregisterDnsHostRecordsW")
	procDsEnumerateDomainTrustsW          = modnetapi32.NewProc("DsEnumerateDomainTrustsW")
	procDsGetDcCloseW                     = modnetapi32.NewProc("DsGetDcCloseW")
	procDsGetDcNameW                      = modnetapi32.NewProc("DsGetDcNameW")
	procDsGetDcNextW                      = modnetapi32.NewProc("DsGetDcNextW")
	procDsGetDcOpenW                      = modnetapi32.NewProc("DsGetDcOpenW")
	procDsGetDcSiteCoverageW              = modnetapi32.NewProc("DsGetDcSiteCoverageW")
	procDsGetForestTrustInformationW      = modnetapi32.NewProc("DsGetForestTrustInformationW")
	procDsGetSiteNameW                    = modnetapi32.NewProc("DsGetSiteNameW")
	procDsMergeForestTrustInformationW    = modnetapi32.NewProc("DsMergeForestTrustInformationW")
	procDsRoleFreeMemory                  = modnetapi32.NewProc("DsRoleFreeMemory")
	procDsRoleGetPrimaryDomainInformation = modnetapi32.NewProc("DsRoleGetPrimaryDomainInformation")
	procDsValidateSubnetNameW             = modnetapi32.NewProc("DsValidateSubnetNameW")
)

const (
	DS_AVOID_SELF                   = 0x00004000
	DS_BACKGROUND_ONLY              = 0x00000100
	DS_DIRECTORY_SERVICE_PREFERRED  = 0x00000020
	DS_DIRECTORY_SERVICE_REQUIRED   = 0x00000010
	DS_DIRECTORY_SERVICE_6_REQUIRED = 0x00080000
	DS_DIRECTORY_SERVICE_8_REQUIRED = 0x00200000
	DS_FORCE_REDISCOVERY            = 0x00000001
	DS_GC_SERVER_REQUIRED           = 0x00000040
	DS_GOOD_TIMESERV_PREFERRED      = 0x00002000
	DS_IP_REQUIRED                  = 0x00000200
	DS_IS_DNS_NAME                  = 0x00020000
	DS_IS_FLAT_NAME                 = 0x00010000
	DS_KDC_REQUIRED                 = 0x00000400
	DS_ONLY_LDAP_NEEDED             = 0x00008000
	DS_PDC_REQUIRED                 = 0x00000080
	DS_RETURN_DNS_NAME              = 0x40000000
	DS_RETURN_FLAT_NAME             = 0x80000000
	DS_TIMESERV_REQUIRED            = 0x00000800
	DS_TRY_NEXTCLOSEST_SITE         = 0x00040000
	DS_WRITABLE_REQUIRED            = 0x00001000
	DS_WEB_SERVICE_REQUIRED         = 0x00100000
)

type DOMAIN_CONTROLLER_INFO struct {
	DomainControllerName        *uint16
	DomainControllerAddress     *uint16
	DomainControllerAddressType uint32
	DomainGuid                  [16]byte
	DomainName                  *uint16
	DnsForestName               *uint16
	Flags                       uint32
	DcSiteName                  *uint16
	ClientSiteName              *uint16
}

type DomainControllerInfo struct {
	DomainControllerName        string
	DomainControllerAddress     string
	DomainControllerAddressType uint32
	DomainGuid                  [16]byte
	DomainName                  string
	DnsForestName               string
	Flags                       uint32
	DcSiteName                  string
	ClientSiteName              string
}

func RefreshDC(computerName, domainName, siteName string) (info *DomainControllerInfo, err error) {
	return GetDcName(computerName, domainName, siteName, DS_FORCE_REDISCOVERY)
}

func GetDcName(computerName, domainName, siteName string, flags uint32) (info *DomainControllerInfo, err error) {
	var cnp, dnp, snp *uint16
	cnp, err = utf16PtrFromString(computerName)
	if err != nil {
		return
	}
	dnp, err = utf16PtrFromString(domainName)
	if err != nil {
		return
	}
	snp, err = utf16PtrFromString(siteName)
	if err != nil {
		return
	}
	var dci *DOMAIN_CONTROLLER_INFO
	err = DsGetDcName(cnp, dnp, snp, flags, &dci)
	if err != nil {
		return
	}
	defer syscall.NetApiBufferFree((*byte)(unsafe.Pointer(dci)))
	info = &DomainControllerInfo{
		DomainControllerName:        utf16PtrToString(dci.DomainControllerName),
		DomainControllerAddress:     utf16PtrToString(dci.DomainControllerAddress),
		DomainControllerAddressType: dci.DomainControllerAddressType,
		DomainGuid:                  dci.DomainGuid,
		DomainName:                  utf16PtrToString(dci.DomainName),
		DnsForestName:               utf16PtrToString(dci.DnsForestName),
		Flags:                       dci.Flags,
		DcSiteName:                  utf16PtrToString(dci.DcSiteName),
		ClientSiteName:              utf16PtrToString(dci.ClientSiteName),
	}
	return
}

func DsGetDcName(computerName, domainName, siteName *uint16, flags uint32, info **DOMAIN_CONTROLLER_INFO) (err error) {
	// See https://msdn.microsoft.com/en-us/library/ms675983
	r0, _, _ := syscall.Syscall6(procDsGetDcNameW.Addr(), 6, uintptr(unsafe.Pointer(computerName)), uintptr(unsafe.Pointer(domainName)), uintptr(0), uintptr(unsafe.Pointer(siteName)), uintptr(flags), uintptr(unsafe.Pointer(info)))
	if r0 != 0 {
		err = syscall.Errno(r0)
	}
	return
}

func GetSiteName(computerName string) (siteName string, err error) {
	var cnp, snp *uint16
	cnp, err = utf16PtrFromString(computerName)
	if err != nil {
		return
	}
	snp, err = DsGetSiteName(cnp)
	if err != nil {
		return
	}
	defer syscall.NetApiBufferFree((*byte)(unsafe.Pointer(snp)))
	siteName = utf16PtrToString(snp)
	return
}

func DsGetSiteName(computerName *uint16) (siteName *uint16, err error) {
	// See https://msdn.microsoft.com/en-us/library/ms675992
	r0, _, _ := syscall.Syscall(procDsGetSiteNameW.Addr(), 2, uintptr(unsafe.Pointer(computerName)), uintptr(unsafe.Pointer(&siteName)), 0)
	if r0 != 0 {
		err = syscall.Errno(r0)
	}
	return
}

func utf16PtrToString(str *uint16) string {
	if str == nil {
		return ""
	} else {
		return syscall.UTF16ToString((*[256]uint16)(unsafe.Pointer(str))[:]) // Assumes max site name length of 256 characters
	}
}

func utf16PtrFromString(str string) (*uint16, error) {
	if len(str) == 0 {
		return nil, nil
	} else {
		return syscall.UTF16PtrFromString(str)
	}
}
