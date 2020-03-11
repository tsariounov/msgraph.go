// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceManagementScriptGroupAssignment Contains properties used to assign a device management script to a group.
type DeviceManagementScriptGroupAssignment struct {
	// Entity is the base model of DeviceManagementScriptGroupAssignment
	Entity
	// TargetGroupID The Id of the Azure Active Directory group we are targeting the script to.
	TargetGroupID *string `json:"targetGroupId,omitempty"`
}

// DeviceManagementScriptGroupAssignmentRequestBuilder is request builder for DeviceManagementScriptGroupAssignment
type DeviceManagementScriptGroupAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementScriptGroupAssignmentRequest
func (b *DeviceManagementScriptGroupAssignmentRequestBuilder) Request() *DeviceManagementScriptGroupAssignmentRequest {
	return &DeviceManagementScriptGroupAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementScriptGroupAssignmentRequest is request for DeviceManagementScriptGroupAssignment
type DeviceManagementScriptGroupAssignmentRequest struct{ BaseRequest }

// Get performs GET request for DeviceManagementScriptGroupAssignment
func (r *DeviceManagementScriptGroupAssignmentRequest) Get(ctx context.Context) (resObj *DeviceManagementScriptGroupAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceManagementScriptGroupAssignment
func (r *DeviceManagementScriptGroupAssignmentRequest) Update(ctx context.Context, reqObj *DeviceManagementScriptGroupAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceManagementScriptGroupAssignment
func (r *DeviceManagementScriptGroupAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
