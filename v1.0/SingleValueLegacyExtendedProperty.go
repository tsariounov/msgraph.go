// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// SingleValueLegacyExtendedProperty undocumented
type SingleValueLegacyExtendedProperty struct {
	// Entity is the base model of SingleValueLegacyExtendedProperty
	Entity
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

// SingleValueLegacyExtendedPropertyRequestBuilder is request builder for SingleValueLegacyExtendedProperty
type SingleValueLegacyExtendedPropertyRequestBuilder struct{ BaseRequestBuilder }

// Request returns SingleValueLegacyExtendedPropertyRequest
func (b *SingleValueLegacyExtendedPropertyRequestBuilder) Request() *SingleValueLegacyExtendedPropertyRequest {
	return &SingleValueLegacyExtendedPropertyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SingleValueLegacyExtendedPropertyRequest is request for SingleValueLegacyExtendedProperty
type SingleValueLegacyExtendedPropertyRequest struct{ BaseRequest }

// Get performs GET request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Get(ctx context.Context) (resObj *SingleValueLegacyExtendedProperty, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Update(ctx context.Context, reqObj *SingleValueLegacyExtendedProperty) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
