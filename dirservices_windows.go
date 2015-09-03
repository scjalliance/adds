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

func GetSiteName(computerName string) (siteName string, err error) {
	var cnp, snp *uint16
	if len(siteName) == 0 {
		cnp = nil
	} else {
		cnp, err = syscall.UTF16PtrFromString(computerName)
		if err != nil {
			return
		}
	}
	snp, err = DsGetSiteName(cnp)
	if err != nil {
		return
	}
	defer syscall.NetApiBufferFree((*byte)(unsafe.Pointer(snp)))
	siteName = syscall.UTF16ToString((*[256]uint16)(unsafe.Pointer(snp))[:]) // Assumes max site name length of 256 characters
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
