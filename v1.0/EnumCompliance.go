// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ComplianceState undocumented
type ComplianceState string

const (
	// ComplianceStateVUnknown undocumented
	ComplianceStateVUnknown ComplianceState = "unknown"
	// ComplianceStateVCompliant undocumented
	ComplianceStateVCompliant ComplianceState = "compliant"
	// ComplianceStateVNoncompliant undocumented
	ComplianceStateVNoncompliant ComplianceState = "noncompliant"
	// ComplianceStateVConflict undocumented
	ComplianceStateVConflict ComplianceState = "conflict"
	// ComplianceStateVError undocumented
	ComplianceStateVError ComplianceState = "error"
	// ComplianceStateVInGracePeriod undocumented
	ComplianceStateVInGracePeriod ComplianceState = "inGracePeriod"
	// ComplianceStateVConfigManager undocumented
	ComplianceStateVConfigManager ComplianceState = "configManager"
)

var (
	// ComplianceStatePUnknown is a pointer to ComplianceStateVUnknown
	ComplianceStatePUnknown = &_ComplianceStatePUnknown
	// ComplianceStatePCompliant is a pointer to ComplianceStateVCompliant
	ComplianceStatePCompliant = &_ComplianceStatePCompliant
	// ComplianceStatePNoncompliant is a pointer to ComplianceStateVNoncompliant
	ComplianceStatePNoncompliant = &_ComplianceStatePNoncompliant
	// ComplianceStatePConflict is a pointer to ComplianceStateVConflict
	ComplianceStatePConflict = &_ComplianceStatePConflict
	// ComplianceStatePError is a pointer to ComplianceStateVError
	ComplianceStatePError = &_ComplianceStatePError
	// ComplianceStatePInGracePeriod is a pointer to ComplianceStateVInGracePeriod
	ComplianceStatePInGracePeriod = &_ComplianceStatePInGracePeriod
	// ComplianceStatePConfigManager is a pointer to ComplianceStateVConfigManager
	ComplianceStatePConfigManager = &_ComplianceStatePConfigManager
)

var (
	_ComplianceStatePUnknown       = ComplianceStateVUnknown
	_ComplianceStatePCompliant     = ComplianceStateVCompliant
	_ComplianceStatePNoncompliant  = ComplianceStateVNoncompliant
	_ComplianceStatePConflict      = ComplianceStateVConflict
	_ComplianceStatePError         = ComplianceStateVError
	_ComplianceStatePInGracePeriod = ComplianceStateVInGracePeriod
	_ComplianceStatePConfigManager = ComplianceStateVConfigManager
)

// ComplianceStatus undocumented
type ComplianceStatus string

const (
	// ComplianceStatusVUnknown undocumented
	ComplianceStatusVUnknown ComplianceStatus = "unknown"
	// ComplianceStatusVNotApplicable undocumented
	ComplianceStatusVNotApplicable ComplianceStatus = "notApplicable"
	// ComplianceStatusVCompliant undocumented
	ComplianceStatusVCompliant ComplianceStatus = "compliant"
	// ComplianceStatusVRemediated undocumented
	ComplianceStatusVRemediated ComplianceStatus = "remediated"
	// ComplianceStatusVNonCompliant undocumented
	ComplianceStatusVNonCompliant ComplianceStatus = "nonCompliant"
	// ComplianceStatusVError undocumented
	ComplianceStatusVError ComplianceStatus = "error"
	// ComplianceStatusVConflict undocumented
	ComplianceStatusVConflict ComplianceStatus = "conflict"
	// ComplianceStatusVNotAssigned undocumented
	ComplianceStatusVNotAssigned ComplianceStatus = "notAssigned"
)

var (
	// ComplianceStatusPUnknown is a pointer to ComplianceStatusVUnknown
	ComplianceStatusPUnknown = &_ComplianceStatusPUnknown
	// ComplianceStatusPNotApplicable is a pointer to ComplianceStatusVNotApplicable
	ComplianceStatusPNotApplicable = &_ComplianceStatusPNotApplicable
	// ComplianceStatusPCompliant is a pointer to ComplianceStatusVCompliant
	ComplianceStatusPCompliant = &_ComplianceStatusPCompliant
	// ComplianceStatusPRemediated is a pointer to ComplianceStatusVRemediated
	ComplianceStatusPRemediated = &_ComplianceStatusPRemediated
	// ComplianceStatusPNonCompliant is a pointer to ComplianceStatusVNonCompliant
	ComplianceStatusPNonCompliant = &_ComplianceStatusPNonCompliant
	// ComplianceStatusPError is a pointer to ComplianceStatusVError
	ComplianceStatusPError = &_ComplianceStatusPError
	// ComplianceStatusPConflict is a pointer to ComplianceStatusVConflict
	ComplianceStatusPConflict = &_ComplianceStatusPConflict
	// ComplianceStatusPNotAssigned is a pointer to ComplianceStatusVNotAssigned
	ComplianceStatusPNotAssigned = &_ComplianceStatusPNotAssigned
)

var (
	_ComplianceStatusPUnknown       = ComplianceStatusVUnknown
	_ComplianceStatusPNotApplicable = ComplianceStatusVNotApplicable
	_ComplianceStatusPCompliant     = ComplianceStatusVCompliant
	_ComplianceStatusPRemediated    = ComplianceStatusVRemediated
	_ComplianceStatusPNonCompliant  = ComplianceStatusVNonCompliant
	_ComplianceStatusPError         = ComplianceStatusVError
	_ComplianceStatusPConflict      = ComplianceStatusVConflict
	_ComplianceStatusPNotAssigned   = ComplianceStatusVNotAssigned
)