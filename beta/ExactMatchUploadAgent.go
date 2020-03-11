// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// ExactMatchUploadAgent undocumented
type ExactMatchUploadAgent struct {
	// Entity is the base model of ExactMatchUploadAgent
	Entity
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// CreationDateTime undocumented
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
}

// ExactMatchUploadAgentRequestBuilder is request builder for ExactMatchUploadAgent
type ExactMatchUploadAgentRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExactMatchUploadAgentRequest
func (b *ExactMatchUploadAgentRequestBuilder) Request() *ExactMatchUploadAgentRequest {
	return &ExactMatchUploadAgentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExactMatchUploadAgentRequest is request for ExactMatchUploadAgent
type ExactMatchUploadAgentRequest struct{ BaseRequest }

// Get performs GET request for ExactMatchUploadAgent
func (r *ExactMatchUploadAgentRequest) Get(ctx context.Context) (resObj *ExactMatchUploadAgent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ExactMatchUploadAgent
func (r *ExactMatchUploadAgentRequest) Update(ctx context.Context, reqObj *ExactMatchUploadAgent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ExactMatchUploadAgent
func (r *ExactMatchUploadAgentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
