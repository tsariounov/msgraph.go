// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// WindowsAutopilotSettings The windowsAutopilotSettings resource represents a Windows Autopilot Account to sync data with Windows device data sync service.
type WindowsAutopilotSettings struct {
	// Entity is the base model of WindowsAutopilotSettings
	Entity
	// LastSyncDateTime Last data sync date time with DDS service.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// LastManualSyncTriggerDateTime Last data sync date time with DDS service.
	LastManualSyncTriggerDateTime *time.Time `json:"lastManualSyncTriggerDateTime,omitempty"`
	// SyncStatus Indicates the status of sync with Device data sync (DDS) service.
	SyncStatus *WindowsAutopilotSyncStatus `json:"syncStatus,omitempty"`
}

// WindowsAutopilotSettingsSyncRequestParameter undocumented
type WindowsAutopilotSettingsSyncRequestParameter struct {
}

// WindowsAutopilotSettingsRequestBuilder is request builder for WindowsAutopilotSettings
type WindowsAutopilotSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsAutopilotSettingsRequest
func (b *WindowsAutopilotSettingsRequestBuilder) Request() *WindowsAutopilotSettingsRequest {
	return &WindowsAutopilotSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsAutopilotSettingsRequest is request for WindowsAutopilotSettings
type WindowsAutopilotSettingsRequest struct{ BaseRequest }

// Get performs GET request for WindowsAutopilotSettings
func (r *WindowsAutopilotSettingsRequest) Get(ctx context.Context) (resObj *WindowsAutopilotSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsAutopilotSettings
func (r *WindowsAutopilotSettingsRequest) Update(ctx context.Context, reqObj *WindowsAutopilotSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsAutopilotSettings
func (r *WindowsAutopilotSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type WindowsAutopilotSettingsSyncRequestBuilder struct{ BaseRequestBuilder }

// Sync action undocumented
func (b *WindowsAutopilotSettingsRequestBuilder) Sync(reqObj *WindowsAutopilotSettingsSyncRequestParameter) *WindowsAutopilotSettingsSyncRequestBuilder {
	bb := &WindowsAutopilotSettingsSyncRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/sync"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WindowsAutopilotSettingsSyncRequest struct{ BaseRequest }

//
func (b *WindowsAutopilotSettingsSyncRequestBuilder) Request() *WindowsAutopilotSettingsSyncRequest {
	return &WindowsAutopilotSettingsSyncRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WindowsAutopilotSettingsSyncRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
