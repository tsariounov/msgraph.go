// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// SchemaExtension undocumented
type SchemaExtension struct {
	// Entity is the base model of SchemaExtension
	Entity
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// TargetTypes undocumented
	TargetTypes []string `json:"targetTypes,omitempty"`
	// Properties undocumented
	Properties []ExtensionSchemaProperty `json:"properties,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// Owner undocumented
	Owner *string `json:"owner,omitempty"`
}

// SchemaExtensionRequestBuilder is request builder for SchemaExtension
type SchemaExtensionRequestBuilder struct{ BaseRequestBuilder }

// Request returns SchemaExtensionRequest
func (b *SchemaExtensionRequestBuilder) Request() *SchemaExtensionRequest {
	return &SchemaExtensionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SchemaExtensionRequest is request for SchemaExtension
type SchemaExtensionRequest struct{ BaseRequest }

// Get performs GET request for SchemaExtension
func (r *SchemaExtensionRequest) Get(ctx context.Context) (resObj *SchemaExtension, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SchemaExtension
func (r *SchemaExtensionRequest) Update(ctx context.Context, reqObj *SchemaExtension) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SchemaExtension
func (r *SchemaExtensionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
