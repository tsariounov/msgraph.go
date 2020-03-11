// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// ExactMatchSession undocumented
type ExactMatchSession struct {
	// ExactMatchJobBase is the base model of ExactMatchSession
	ExactMatchJobBase
	// DatastoreID undocumented
	DatastoreID *string `json:"datastoreId,omitempty"`
	// UploadAgentID undocumented
	UploadAgentID *string `json:"uploadAgentId,omitempty"`
	// Fields undocumented
	Fields []string `json:"fields,omitempty"`
	// FileName undocumented
	FileName *string `json:"fileName,omitempty"`
	// Checksum undocumented
	Checksum *string `json:"checksum,omitempty"`
	// DataUploadURI undocumented
	DataUploadURI *string `json:"dataUploadURI,omitempty"`
	// RemainingBlockCount undocumented
	RemainingBlockCount *int `json:"remainingBlockCount,omitempty"`
	// TotalBlockCount undocumented
	TotalBlockCount *int `json:"totalBlockCount,omitempty"`
	// State undocumented
	State *string `json:"state,omitempty"`
	// UploadCompletionDateTime undocumented
	UploadCompletionDateTime *time.Time `json:"uploadCompletionDateTime,omitempty"`
	// ProcessingCompletionDateTime undocumented
	ProcessingCompletionDateTime *time.Time `json:"processingCompletionDateTime,omitempty"`
	// RowsPerBlock undocumented
	RowsPerBlock *int `json:"rowsPerBlock,omitempty"`
	// TotalJobCount undocumented
	TotalJobCount *int `json:"totalJobCount,omitempty"`
	// RemainingJobCount undocumented
	RemainingJobCount *int `json:"remainingJobCount,omitempty"`
	// UploadAgent undocumented
	UploadAgent *ExactMatchUploadAgent `json:"uploadAgent,omitempty"`
}

// ExactMatchSessionCancelRequestParameter undocumented
type ExactMatchSessionCancelRequestParameter struct {
}

// ExactMatchSessionCommitRequestParameter undocumented
type ExactMatchSessionCommitRequestParameter struct {
}

// ExactMatchSessionRenewRequestParameter undocumented
type ExactMatchSessionRenewRequestParameter struct {
}

// ExactMatchSessionRequestBuilder is request builder for ExactMatchSession
type ExactMatchSessionRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExactMatchSessionRequest
func (b *ExactMatchSessionRequestBuilder) Request() *ExactMatchSessionRequest {
	return &ExactMatchSessionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExactMatchSessionRequest is request for ExactMatchSession
type ExactMatchSessionRequest struct{ BaseRequest }

// Get performs GET request for ExactMatchSession
func (r *ExactMatchSessionRequest) Get(ctx context.Context) (resObj *ExactMatchSession, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExactMatchSession
func (r *ExactMatchSessionRequest) Update(ctx context.Context, reqObj *ExactMatchSession) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExactMatchSession
func (r *ExactMatchSessionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UploadAgent is navigation property
func (b *ExactMatchSessionRequestBuilder) UploadAgent() *ExactMatchUploadAgentRequestBuilder {
	bb := &ExactMatchUploadAgentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/uploadAgent"
	return bb
}

//
type ExactMatchSessionCancelRequestBuilder struct{ BaseRequestBuilder }

// Cancel action undocumented
func (b *ExactMatchSessionRequestBuilder) Cancel(reqObj *ExactMatchSessionCancelRequestParameter) *ExactMatchSessionCancelRequestBuilder {
	bb := &ExactMatchSessionCancelRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/cancel"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ExactMatchSessionCancelRequest struct{ BaseRequest }

//
func (b *ExactMatchSessionCancelRequestBuilder) Request() *ExactMatchSessionCancelRequest {
	return &ExactMatchSessionCancelRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ExactMatchSessionCancelRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ExactMatchSessionCommitRequestBuilder struct{ BaseRequestBuilder }

// Commit action undocumented
func (b *ExactMatchSessionRequestBuilder) Commit(reqObj *ExactMatchSessionCommitRequestParameter) *ExactMatchSessionCommitRequestBuilder {
	bb := &ExactMatchSessionCommitRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/commit"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ExactMatchSessionCommitRequest struct{ BaseRequest }

//
func (b *ExactMatchSessionCommitRequestBuilder) Request() *ExactMatchSessionCommitRequest {
	return &ExactMatchSessionCommitRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ExactMatchSessionCommitRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ExactMatchSessionRenewRequestBuilder struct{ BaseRequestBuilder }

// Renew action undocumented
func (b *ExactMatchSessionRequestBuilder) Renew(reqObj *ExactMatchSessionRenewRequestParameter) *ExactMatchSessionRenewRequestBuilder {
	bb := &ExactMatchSessionRenewRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/renew"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ExactMatchSessionRenewRequest struct{ BaseRequest }

//
func (b *ExactMatchSessionRenewRequestBuilder) Request() *ExactMatchSessionRenewRequest {
	return &ExactMatchSessionRenewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ExactMatchSessionRenewRequest) Post(ctx context.Context) (resObj *ExactMatchSession, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
