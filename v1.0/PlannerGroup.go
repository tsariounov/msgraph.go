// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// PlannerGroup undocumented
type PlannerGroup struct {
	// Entity is the base model of PlannerGroup
	Entity
	// Plans undocumented
	Plans []PlannerPlan `json:"plans,omitempty"`
}

// PlannerGroupRequestBuilder is request builder for PlannerGroup
type PlannerGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlannerGroupRequest
func (b *PlannerGroupRequestBuilder) Request() *PlannerGroupRequest {
	return &PlannerGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlannerGroupRequest is request for PlannerGroup
type PlannerGroupRequest struct{ BaseRequest }

// Get performs GET request for PlannerGroup
func (r *PlannerGroupRequest) Get(ctx context.Context) (resObj *PlannerGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PlannerGroup
func (r *PlannerGroupRequest) Update(ctx context.Context, reqObj *PlannerGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PlannerGroup
func (r *PlannerGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Plans returns request builder for PlannerPlan collection
func (b *PlannerGroupRequestBuilder) Plans() *PlannerGroupPlansCollectionRequestBuilder {
	bb := &PlannerGroupPlansCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/plans"
	return bb
}

// PlannerGroupPlansCollectionRequestBuilder is request builder for PlannerPlan collection
type PlannerGroupPlansCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerPlan collection
func (b *PlannerGroupPlansCollectionRequestBuilder) Request() *PlannerGroupPlansCollectionRequest {
	return &PlannerGroupPlansCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerPlan item
func (b *PlannerGroupPlansCollectionRequestBuilder) ID(id string) *PlannerPlanRequestBuilder {
	bb := &PlannerPlanRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerGroupPlansCollectionRequest is request for PlannerPlan collection
type PlannerGroupPlansCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PlannerPlan collection
func (r *PlannerGroupPlansCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]PlannerPlan, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PlannerPlan
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []PlannerPlan
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for PlannerPlan collection
func (r *PlannerGroupPlansCollectionRequest) Get(ctx context.Context) ([]PlannerPlan, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for PlannerPlan collection
func (r *PlannerGroupPlansCollectionRequest) Add(ctx context.Context, reqObj *PlannerPlan) (resObj *PlannerPlan, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
