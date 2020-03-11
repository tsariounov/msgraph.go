// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// ManagedDeviceMobileAppConfigurationDeviceSummary Contains properties, inherited properties and actions for an MDM mobile app configuration device status summary.
type ManagedDeviceMobileAppConfigurationDeviceSummary struct {
	// Entity is the base model of ManagedDeviceMobileAppConfigurationDeviceSummary
	Entity
	// PendingCount Number of pending devices
	PendingCount *int `json:"pendingCount,omitempty"`
	// NotApplicableCount Number of not applicable devices
	NotApplicableCount *int `json:"notApplicableCount,omitempty"`
	// NotApplicablePlatformCount Number of not applicable devices due to mismatch platform and policy
	NotApplicablePlatformCount *int `json:"notApplicablePlatformCount,omitempty"`
	// SuccessCount Number of succeeded devices
	SuccessCount *int `json:"successCount,omitempty"`
	// ErrorCount Number of error devices
	ErrorCount *int `json:"errorCount,omitempty"`
	// FailedCount Number of failed devices
	FailedCount *int `json:"failedCount,omitempty"`
	// ConflictCount Number of devices in conflict
	ConflictCount *int `json:"conflictCount,omitempty"`
	// LastUpdateDateTime Last update time
	LastUpdateDateTime *time.Time `json:"lastUpdateDateTime,omitempty"`
	// ConfigurationVersion Version of the policy for that overview
	ConfigurationVersion *int `json:"configurationVersion,omitempty"`
}

// ManagedDeviceMobileAppConfigurationDeviceSummaryRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationDeviceSummary
type ManagedDeviceMobileAppConfigurationDeviceSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceMobileAppConfigurationDeviceSummaryRequest
func (b *ManagedDeviceMobileAppConfigurationDeviceSummaryRequestBuilder) Request() *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest {
	return &ManagedDeviceMobileAppConfigurationDeviceSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceMobileAppConfigurationDeviceSummaryRequest is request for ManagedDeviceMobileAppConfigurationDeviceSummary
type ManagedDeviceMobileAppConfigurationDeviceSummaryRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceMobileAppConfigurationDeviceSummary
func (r *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest) Get(ctx context.Context) (resObj *ManagedDeviceMobileAppConfigurationDeviceSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceMobileAppConfigurationDeviceSummary
func (r *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest) Update(ctx context.Context, reqObj *ManagedDeviceMobileAppConfigurationDeviceSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceMobileAppConfigurationDeviceSummary
func (r *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
