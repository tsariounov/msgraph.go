// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// DeviceComplianceUserOverview undocumented
type DeviceComplianceUserOverview struct {
	// Entity is the base model of DeviceComplianceUserOverview
	Entity
	// PendingCount Number of pending Users
	PendingCount *int `json:"pendingCount,omitempty"`
	// NotApplicableCount Number of not applicable users
	NotApplicableCount *int `json:"notApplicableCount,omitempty"`
	// SuccessCount Number of succeeded Users
	SuccessCount *int `json:"successCount,omitempty"`
	// ErrorCount Number of error Users
	ErrorCount *int `json:"errorCount,omitempty"`
	// FailedCount Number of failed Users
	FailedCount *int `json:"failedCount,omitempty"`
	// ConflictCount Number of users in conflict
	ConflictCount *int `json:"conflictCount,omitempty"`
	// LastUpdateDateTime Last update time
	LastUpdateDateTime *time.Time `json:"lastUpdateDateTime,omitempty"`
	// ConfigurationVersion Version of the policy for that overview
	ConfigurationVersion *int `json:"configurationVersion,omitempty"`
}

// DeviceComplianceUserOverviewRequestBuilder is request builder for DeviceComplianceUserOverview
type DeviceComplianceUserOverviewRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceComplianceUserOverviewRequest
func (b *DeviceComplianceUserOverviewRequestBuilder) Request() *DeviceComplianceUserOverviewRequest {
	return &DeviceComplianceUserOverviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceComplianceUserOverviewRequest is request for DeviceComplianceUserOverview
type DeviceComplianceUserOverviewRequest struct{ BaseRequest }

// Get performs GET request for DeviceComplianceUserOverview
func (r *DeviceComplianceUserOverviewRequest) Get(ctx context.Context) (resObj *DeviceComplianceUserOverview, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceComplianceUserOverview
func (r *DeviceComplianceUserOverviewRequest) Update(ctx context.Context, reqObj *DeviceComplianceUserOverview) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceComplianceUserOverview
func (r *DeviceComplianceUserOverviewRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
