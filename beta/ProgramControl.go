// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// ProgramControl undocumented
type ProgramControl struct {
	// Entity is the base model of ProgramControl
	Entity
	// ControlID undocumented
	ControlID *string `json:"controlId,omitempty"`
	// ProgramID undocumented
	ProgramID *string `json:"programId,omitempty"`
	// ControlTypeID undocumented
	ControlTypeID *string `json:"controlTypeId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// Owner undocumented
	Owner *UserIdentity `json:"owner,omitempty"`
	// Resource undocumented
	Resource *ProgramResource `json:"resource,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Program undocumented
	Program *Program `json:"program,omitempty"`
}

// ProgramControlRequestBuilder is request builder for ProgramControl
type ProgramControlRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProgramControlRequest
func (b *ProgramControlRequestBuilder) Request() *ProgramControlRequest {
	return &ProgramControlRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProgramControlRequest is request for ProgramControl
type ProgramControlRequest struct{ BaseRequest }

// Get performs GET request for ProgramControl
func (r *ProgramControlRequest) Get(ctx context.Context) (resObj *ProgramControl, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ProgramControl
func (r *ProgramControlRequest) Update(ctx context.Context, reqObj *ProgramControl) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ProgramControl
func (r *ProgramControlRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Program is navigation property
func (b *ProgramControlRequestBuilder) Program() *ProgramRequestBuilder {
	bb := &ProgramRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/program"
	return bb
}
