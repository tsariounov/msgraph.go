// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConnectionStatus undocumented
type ConnectionStatus string

const (
	// ConnectionStatusVUnknown undocumented
	ConnectionStatusVUnknown ConnectionStatus = "unknown"
	// ConnectionStatusVAttempted undocumented
	ConnectionStatusVAttempted ConnectionStatus = "attempted"
	// ConnectionStatusVSucceeded undocumented
	ConnectionStatusVSucceeded ConnectionStatus = "succeeded"
	// ConnectionStatusVBlocked undocumented
	ConnectionStatusVBlocked ConnectionStatus = "blocked"
	// ConnectionStatusVFailed undocumented
	ConnectionStatusVFailed ConnectionStatus = "failed"
	// ConnectionStatusVUnknownFutureValue undocumented
	ConnectionStatusVUnknownFutureValue ConnectionStatus = "unknownFutureValue"
)

var (
	// ConnectionStatusPUnknown is a pointer to ConnectionStatusVUnknown
	ConnectionStatusPUnknown = &_ConnectionStatusPUnknown
	// ConnectionStatusPAttempted is a pointer to ConnectionStatusVAttempted
	ConnectionStatusPAttempted = &_ConnectionStatusPAttempted
	// ConnectionStatusPSucceeded is a pointer to ConnectionStatusVSucceeded
	ConnectionStatusPSucceeded = &_ConnectionStatusPSucceeded
	// ConnectionStatusPBlocked is a pointer to ConnectionStatusVBlocked
	ConnectionStatusPBlocked = &_ConnectionStatusPBlocked
	// ConnectionStatusPFailed is a pointer to ConnectionStatusVFailed
	ConnectionStatusPFailed = &_ConnectionStatusPFailed
	// ConnectionStatusPUnknownFutureValue is a pointer to ConnectionStatusVUnknownFutureValue
	ConnectionStatusPUnknownFutureValue = &_ConnectionStatusPUnknownFutureValue
)

var (
	_ConnectionStatusPUnknown            = ConnectionStatusVUnknown
	_ConnectionStatusPAttempted          = ConnectionStatusVAttempted
	_ConnectionStatusPSucceeded          = ConnectionStatusVSucceeded
	_ConnectionStatusPBlocked            = ConnectionStatusVBlocked
	_ConnectionStatusPFailed             = ConnectionStatusVFailed
	_ConnectionStatusPUnknownFutureValue = ConnectionStatusVUnknownFutureValue
)
