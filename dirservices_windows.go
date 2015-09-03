package adds

import "syscall"

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
