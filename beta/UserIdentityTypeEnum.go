// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserIdentityType undocumented
type UserIdentityType string

const (
	// UserIdentityTypeVAadUser undocumented
	UserIdentityTypeVAadUser UserIdentityType = "aadUser"
	// UserIdentityTypeVOnPremiseAadUser undocumented
	UserIdentityTypeVOnPremiseAadUser UserIdentityType = "onPremiseAadUser"
	// UserIdentityTypeVAnonymousGuest undocumented
	UserIdentityTypeVAnonymousGuest UserIdentityType = "anonymousGuest"
	// UserIdentityTypeVFederatedUser undocumented
	UserIdentityTypeVFederatedUser UserIdentityType = "federatedUser"
)

var (
	// UserIdentityTypePAadUser is a pointer to UserIdentityTypeVAadUser
	UserIdentityTypePAadUser = &_UserIdentityTypePAadUser
	// UserIdentityTypePOnPremiseAadUser is a pointer to UserIdentityTypeVOnPremiseAadUser
	UserIdentityTypePOnPremiseAadUser = &_UserIdentityTypePOnPremiseAadUser
	// UserIdentityTypePAnonymousGuest is a pointer to UserIdentityTypeVAnonymousGuest
	UserIdentityTypePAnonymousGuest = &_UserIdentityTypePAnonymousGuest
	// UserIdentityTypePFederatedUser is a pointer to UserIdentityTypeVFederatedUser
	UserIdentityTypePFederatedUser = &_UserIdentityTypePFederatedUser
)

var (
	_UserIdentityTypePAadUser          = UserIdentityTypeVAadUser
	_UserIdentityTypePOnPremiseAadUser = UserIdentityTypeVOnPremiseAadUser
	_UserIdentityTypePAnonymousGuest   = UserIdentityTypeVAnonymousGuest
	_UserIdentityTypePFederatedUser    = UserIdentityTypeVFederatedUser
)
