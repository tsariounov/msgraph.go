// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceCompliancePolicyGroupAssignmentRequestBuilder is request builder for DeviceCompliancePolicyGroupAssignment
type DeviceCompliancePolicyGroupAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceCompliancePolicyGroupAssignmentRequest
func (b *DeviceCompliancePolicyGroupAssignmentRequestBuilder) Request() *DeviceCompliancePolicyGroupAssignmentRequest {
	return &DeviceCompliancePolicyGroupAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceCompliancePolicyGroupAssignmentRequest is request for DeviceCompliancePolicyGroupAssignment
type DeviceCompliancePolicyGroupAssignmentRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceCompliancePolicyGroupAssignment
func (r *DeviceCompliancePolicyGroupAssignmentRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceCompliancePolicyGroupAssignment, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceCompliancePolicyGroupAssignment
func (r *DeviceCompliancePolicyGroupAssignmentRequest) Get() (*DeviceCompliancePolicyGroupAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceCompliancePolicyGroupAssignment
func (r *DeviceCompliancePolicyGroupAssignmentRequest) Update(reqObj *DeviceCompliancePolicyGroupAssignment) (*DeviceCompliancePolicyGroupAssignment, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceCompliancePolicyGroupAssignment
func (r *DeviceCompliancePolicyGroupAssignmentRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// DeviceCompliancePolicy is navigation property
func (b *DeviceCompliancePolicyGroupAssignmentRequestBuilder) DeviceCompliancePolicy() *DeviceCompliancePolicyRequestBuilder {
	bb := &DeviceCompliancePolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceCompliancePolicy"
	return bb
}