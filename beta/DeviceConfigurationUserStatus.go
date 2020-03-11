// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// DeviceConfigurationUserStatus undocumented
type DeviceConfigurationUserStatus struct {
	// Entity is the base model of DeviceConfigurationUserStatus
	Entity
	// UserDisplayName User name of the DevicePolicyStatus.
	UserDisplayName *string `json:"userDisplayName,omitempty"`
	// DevicesCount Devices count for that user.
	DevicesCount *int `json:"devicesCount,omitempty"`
	// Status Compliance status of the policy report.
	Status *ComplianceStatus `json:"status,omitempty"`
	// LastReportedDateTime Last modified date time of the policy report.
	LastReportedDateTime *time.Time `json:"lastReportedDateTime,omitempty"`
	// UserPrincipalName UserPrincipalName.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// DeviceConfigurationUserStatusRequestBuilder is request builder for DeviceConfigurationUserStatus
type DeviceConfigurationUserStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceConfigurationUserStatusRequest
func (b *DeviceConfigurationUserStatusRequestBuilder) Request() *DeviceConfigurationUserStatusRequest {
	return &DeviceConfigurationUserStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceConfigurationUserStatusRequest is request for DeviceConfigurationUserStatus
type DeviceConfigurationUserStatusRequest struct{ BaseRequest }

// Get performs GET request for DeviceConfigurationUserStatus
func (r *DeviceConfigurationUserStatusRequest) Get(ctx context.Context) (resObj *DeviceConfigurationUserStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceConfigurationUserStatus
func (r *DeviceConfigurationUserStatusRequest) Update(ctx context.Context, reqObj *DeviceConfigurationUserStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceConfigurationUserStatus
func (r *DeviceConfigurationUserStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
