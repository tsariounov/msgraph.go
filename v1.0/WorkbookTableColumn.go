// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"encoding/json"
)

// WorkbookTableColumn undocumented
type WorkbookTableColumn struct {
	// Entity is the base model of WorkbookTableColumn
	Entity
	// Index undocumented
	Index *int `json:"index,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Values undocumented
	Values json.RawMessage `json:"values,omitempty"`
	// Filter undocumented
	Filter *WorkbookFilter `json:"filter,omitempty"`
}

// WorkbookTableColumnCollectionAddRequestParameter undocumented
type WorkbookTableColumnCollectionAddRequestParameter struct {
	// Index undocumented
	Index *int `json:"index,omitempty"`
	// Values undocumented
	Values json.RawMessage `json:"values,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
}

// WorkbookTableColumnRequestBuilder is request builder for WorkbookTableColumn
type WorkbookTableColumnRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookTableColumnRequest
func (b *WorkbookTableColumnRequestBuilder) Request() *WorkbookTableColumnRequest {
	return &WorkbookTableColumnRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookTableColumnRequest is request for WorkbookTableColumn
type WorkbookTableColumnRequest struct{ BaseRequest }

// Get performs GET request for WorkbookTableColumn
func (r *WorkbookTableColumnRequest) Get(ctx context.Context) (resObj *WorkbookTableColumn, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookTableColumn
func (r *WorkbookTableColumnRequest) Update(ctx context.Context, reqObj *WorkbookTableColumn) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookTableColumn
func (r *WorkbookTableColumnRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Filter is navigation property
func (b *WorkbookTableColumnRequestBuilder) Filter() *WorkbookFilterRequestBuilder {
	bb := &WorkbookFilterRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/filter"
	return bb
}

//
type WorkbookTableColumnCollectionAddRequestBuilder struct{ BaseRequestBuilder }

// Add action undocumented
func (b *WorkbookTableColumnsCollectionRequestBuilder) Add(reqObj *WorkbookTableColumnCollectionAddRequestParameter) *WorkbookTableColumnCollectionAddRequestBuilder {
	bb := &WorkbookTableColumnCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookTableColumnCollectionAddRequest struct{ BaseRequest }

//
func (b *WorkbookTableColumnCollectionAddRequestBuilder) Request() *WorkbookTableColumnCollectionAddRequest {
	return &WorkbookTableColumnCollectionAddRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookTableColumnCollectionAddRequest) Post(ctx context.Context) (resObj *WorkbookTableColumn, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
