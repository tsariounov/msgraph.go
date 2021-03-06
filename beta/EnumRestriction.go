// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RestrictionAction undocumented
type RestrictionAction string

const (
	// RestrictionActionVWarn undocumented
	RestrictionActionVWarn RestrictionAction = "warn"
	// RestrictionActionVAudit undocumented
	RestrictionActionVAudit RestrictionAction = "audit"
	// RestrictionActionVBlock undocumented
	RestrictionActionVBlock RestrictionAction = "block"
)

var (
	// RestrictionActionPWarn is a pointer to RestrictionActionVWarn
	RestrictionActionPWarn = &_RestrictionActionPWarn
	// RestrictionActionPAudit is a pointer to RestrictionActionVAudit
	RestrictionActionPAudit = &_RestrictionActionPAudit
	// RestrictionActionPBlock is a pointer to RestrictionActionVBlock
	RestrictionActionPBlock = &_RestrictionActionPBlock
)

var (
	_RestrictionActionPWarn  = RestrictionActionVWarn
	_RestrictionActionPAudit = RestrictionActionVAudit
	_RestrictionActionPBlock = RestrictionActionVBlock
)

// RestrictionTrigger undocumented
type RestrictionTrigger string

const (
	// RestrictionTriggerVCopyPaste undocumented
	RestrictionTriggerVCopyPaste RestrictionTrigger = "copyPaste"
	// RestrictionTriggerVCopyToNetworkShare undocumented
	RestrictionTriggerVCopyToNetworkShare RestrictionTrigger = "copyToNetworkShare"
	// RestrictionTriggerVCopyToRemovableMedia undocumented
	RestrictionTriggerVCopyToRemovableMedia RestrictionTrigger = "copyToRemovableMedia"
	// RestrictionTriggerVScreenCapture undocumented
	RestrictionTriggerVScreenCapture RestrictionTrigger = "screenCapture"
	// RestrictionTriggerVPrint undocumented
	RestrictionTriggerVPrint RestrictionTrigger = "print"
	// RestrictionTriggerVCloudEgress undocumented
	RestrictionTriggerVCloudEgress RestrictionTrigger = "cloudEgress"
	// RestrictionTriggerVUnallowedApps undocumented
	RestrictionTriggerVUnallowedApps RestrictionTrigger = "unallowedApps"
)

var (
	// RestrictionTriggerPCopyPaste is a pointer to RestrictionTriggerVCopyPaste
	RestrictionTriggerPCopyPaste = &_RestrictionTriggerPCopyPaste
	// RestrictionTriggerPCopyToNetworkShare is a pointer to RestrictionTriggerVCopyToNetworkShare
	RestrictionTriggerPCopyToNetworkShare = &_RestrictionTriggerPCopyToNetworkShare
	// RestrictionTriggerPCopyToRemovableMedia is a pointer to RestrictionTriggerVCopyToRemovableMedia
	RestrictionTriggerPCopyToRemovableMedia = &_RestrictionTriggerPCopyToRemovableMedia
	// RestrictionTriggerPScreenCapture is a pointer to RestrictionTriggerVScreenCapture
	RestrictionTriggerPScreenCapture = &_RestrictionTriggerPScreenCapture
	// RestrictionTriggerPPrint is a pointer to RestrictionTriggerVPrint
	RestrictionTriggerPPrint = &_RestrictionTriggerPPrint
	// RestrictionTriggerPCloudEgress is a pointer to RestrictionTriggerVCloudEgress
	RestrictionTriggerPCloudEgress = &_RestrictionTriggerPCloudEgress
	// RestrictionTriggerPUnallowedApps is a pointer to RestrictionTriggerVUnallowedApps
	RestrictionTriggerPUnallowedApps = &_RestrictionTriggerPUnallowedApps
)

var (
	_RestrictionTriggerPCopyPaste            = RestrictionTriggerVCopyPaste
	_RestrictionTriggerPCopyToNetworkShare   = RestrictionTriggerVCopyToNetworkShare
	_RestrictionTriggerPCopyToRemovableMedia = RestrictionTriggerVCopyToRemovableMedia
	_RestrictionTriggerPScreenCapture        = RestrictionTriggerVScreenCapture
	_RestrictionTriggerPPrint                = RestrictionTriggerVPrint
	_RestrictionTriggerPCloudEgress          = RestrictionTriggerVCloudEgress
	_RestrictionTriggerPUnallowedApps        = RestrictionTriggerVUnallowedApps
)
