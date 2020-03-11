// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// VppToken You purchase multiple licenses for iOS apps through the Apple Volume Purchase Program for Business or Education. This involves setting up an Apple VPP account from the Apple website and uploading the Apple VPP Business or Education token to Intune. You can then synchronize your volume purchase information with Intune and track your volume-purchased app use. You can upload multiple Apple VPP Business or Education tokens.
type VppToken struct {
	// Entity is the base model of VppToken
	Entity
	// OrganizationName The organization associated with the Apple Volume Purchase Program Token
	OrganizationName *string `json:"organizationName,omitempty"`
	// VppTokenAccountType The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: `business`, `education`.
	VppTokenAccountType *VppTokenAccountType `json:"vppTokenAccountType,omitempty"`
	// AppleID The apple Id associated with the given Apple Volume Purchase Program Token.
	AppleID *string `json:"appleId,omitempty"`
	// ExpirationDateTime The expiration date time of the Apple Volume Purchase Program Token.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// LastSyncDateTime The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// Token The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
	Token *string `json:"token,omitempty"`
	// LastModifiedDateTime Last modification date time associated with the Apple Volume Purchase Program Token.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// State Current state of the Apple Volume Purchase Program Token. Possible values are: `unknown`, `valid`, `expired`, `invalid`, `assignedToExternalMDM`.
	State *VppTokenState `json:"state,omitempty"`
	// TokenActionResults The collection of statuses of the actions performed on the Apple Volume Purchase Program Token.
	TokenActionResults []VppTokenActionResult `json:"tokenActionResults,omitempty"`
	// LastSyncStatus Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: `none`, `inProgress`, `completed`, `failed`.
	LastSyncStatus *VppTokenSyncStatus `json:"lastSyncStatus,omitempty"`
	// AutomaticallyUpdateApps Whether or not apps for the VPP token will be automatically updated.
	AutomaticallyUpdateApps *bool `json:"automaticallyUpdateApps,omitempty"`
	// CountryOrRegion Whether or not apps for the VPP token will be automatically updated.
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// DataSharingConsentGranted Consent granted for data sharing with the Apple Volume Purchase Program.
	DataSharingConsentGranted *bool `json:"dataSharingConsentGranted,omitempty"`
	// DisplayName An admin specified token friendly name.
	DisplayName *string `json:"displayName,omitempty"`
	// LocationName Token location returned from Apple VPP.
	LocationName *string `json:"locationName,omitempty"`
	// ClaimTokenManagementFromExternalMdm Admin consent to allow claiming token management from external MDM.
	ClaimTokenManagementFromExternalMdm *bool `json:"claimTokenManagementFromExternalMdm,omitempty"`
	// RoleScopeTagIDs Role Scope Tags IDs assigned to this entity.
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
}

// VppTokenSyncLicensesRequestParameter undocumented
type VppTokenSyncLicensesRequestParameter struct {
}

// VppTokenRevokeLicensesRequestParameter undocumented
type VppTokenRevokeLicensesRequestParameter struct {
	// NotifyManagedDevices undocumented
	NotifyManagedDevices *bool `json:"notifyManagedDevices,omitempty"`
}

// VppTokenRequestBuilder is request builder for VppToken
type VppTokenRequestBuilder struct{ BaseRequestBuilder }

// Request returns VppTokenRequest
func (b *VppTokenRequestBuilder) Request() *VppTokenRequest {
	return &VppTokenRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// VppTokenRequest is request for VppToken
type VppTokenRequest struct{ BaseRequest }

// Get performs GET request for VppToken
func (r *VppTokenRequest) Get(ctx context.Context) (resObj *VppToken, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for VppToken
func (r *VppTokenRequest) Update(ctx context.Context, reqObj *VppToken) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for VppToken
func (r *VppTokenRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type VppTokenSyncLicensesRequestBuilder struct{ BaseRequestBuilder }

// SyncLicenses action undocumented
func (b *VppTokenRequestBuilder) SyncLicenses(reqObj *VppTokenSyncLicensesRequestParameter) *VppTokenSyncLicensesRequestBuilder {
	bb := &VppTokenSyncLicensesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/syncLicenses"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type VppTokenSyncLicensesRequest struct{ BaseRequest }

//
func (b *VppTokenSyncLicensesRequestBuilder) Request() *VppTokenSyncLicensesRequest {
	return &VppTokenSyncLicensesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *VppTokenSyncLicensesRequest) Post(ctx context.Context) (resObj *VppToken, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type VppTokenRevokeLicensesRequestBuilder struct{ BaseRequestBuilder }

// RevokeLicenses action undocumented
func (b *VppTokenRequestBuilder) RevokeLicenses(reqObj *VppTokenRevokeLicensesRequestParameter) *VppTokenRevokeLicensesRequestBuilder {
	bb := &VppTokenRevokeLicensesRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/revokeLicenses"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type VppTokenRevokeLicensesRequest struct{ BaseRequest }

//
func (b *VppTokenRevokeLicensesRequestBuilder) Request() *VppTokenRevokeLicensesRequest {
	return &VppTokenRevokeLicensesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *VppTokenRevokeLicensesRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
