// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AppleUserInitiatedEnrollmentProfile The enrollmentProfile resource represents a collection of configurations which must be provided pre-enrollment to enable enrolling certain devices whose identities have been pre-staged. Pre-staged device identities are assigned to this type of profile to apply the profile's configurations at enrollment of the corresponding device.
type AppleUserInitiatedEnrollmentProfile struct {
	// Entity is the base model of AppleUserInitiatedEnrollmentProfile
	Entity
	// DefaultEnrollmentType The default profile enrollment type.
	DefaultEnrollmentType *AppleUserInitiatedEnrollmentType `json:"defaultEnrollmentType,omitempty"`
	// AvailableEnrollmentTypeOptions List of available enrollment type options
	AvailableEnrollmentTypeOptions []AppleOwnerTypeEnrollmentType `json:"availableEnrollmentTypeOptions,omitempty"`
	// DisplayName Name of the profile
	DisplayName *string `json:"displayName,omitempty"`
	// Description Description of the profile
	Description *string `json:"description,omitempty"`
	// Priority Priority, 0 is highest
	Priority *int `json:"priority,omitempty"`
	// Platform The platform of the Device.
	Platform *DevicePlatformType `json:"platform,omitempty"`
	// CreatedDateTime Profile creation time
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime Profile last modified time
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Assignments undocumented
	Assignments []AppleEnrollmentProfileAssignment `json:"assignments,omitempty"`
}

// AppleUserInitiatedEnrollmentProfileSetPriorityRequestParameter undocumented
type AppleUserInitiatedEnrollmentProfileSetPriorityRequestParameter struct {
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
}

// AppleUserInitiatedEnrollmentProfileRequestBuilder is request builder for AppleUserInitiatedEnrollmentProfile
type AppleUserInitiatedEnrollmentProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppleUserInitiatedEnrollmentProfileRequest
func (b *AppleUserInitiatedEnrollmentProfileRequestBuilder) Request() *AppleUserInitiatedEnrollmentProfileRequest {
	return &AppleUserInitiatedEnrollmentProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppleUserInitiatedEnrollmentProfileRequest is request for AppleUserInitiatedEnrollmentProfile
type AppleUserInitiatedEnrollmentProfileRequest struct{ BaseRequest }

// Get performs GET request for AppleUserInitiatedEnrollmentProfile
func (r *AppleUserInitiatedEnrollmentProfileRequest) Get(ctx context.Context) (resObj *AppleUserInitiatedEnrollmentProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppleUserInitiatedEnrollmentProfile
func (r *AppleUserInitiatedEnrollmentProfileRequest) Update(ctx context.Context, reqObj *AppleUserInitiatedEnrollmentProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppleUserInitiatedEnrollmentProfile
func (r *AppleUserInitiatedEnrollmentProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for AppleEnrollmentProfileAssignment collection
func (b *AppleUserInitiatedEnrollmentProfileRequestBuilder) Assignments() *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequestBuilder {
	bb := &AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequestBuilder is request builder for AppleEnrollmentProfileAssignment collection
type AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AppleEnrollmentProfileAssignment collection
func (b *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequestBuilder) Request() *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest {
	return &AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AppleEnrollmentProfileAssignment item
func (b *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequestBuilder) ID(id string) *AppleEnrollmentProfileAssignmentRequestBuilder {
	bb := &AppleEnrollmentProfileAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest is request for AppleEnrollmentProfileAssignment collection
type AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AppleEnrollmentProfileAssignment collection
func (r *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AppleEnrollmentProfileAssignment, error) {
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
	var values []AppleEnrollmentProfileAssignment
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
			value  []AppleEnrollmentProfileAssignment
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

// Get performs GET request for AppleEnrollmentProfileAssignment collection
func (r *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest) Get(ctx context.Context) ([]AppleEnrollmentProfileAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AppleEnrollmentProfileAssignment collection
func (r *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *AppleEnrollmentProfileAssignment) (resObj *AppleEnrollmentProfileAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

//
type AppleUserInitiatedEnrollmentProfileSetPriorityRequestBuilder struct{ BaseRequestBuilder }

// SetPriority action undocumented
func (b *AppleUserInitiatedEnrollmentProfileRequestBuilder) SetPriority(reqObj *AppleUserInitiatedEnrollmentProfileSetPriorityRequestParameter) *AppleUserInitiatedEnrollmentProfileSetPriorityRequestBuilder {
	bb := &AppleUserInitiatedEnrollmentProfileSetPriorityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/setPriority"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AppleUserInitiatedEnrollmentProfileSetPriorityRequest struct{ BaseRequest }

//
func (b *AppleUserInitiatedEnrollmentProfileSetPriorityRequestBuilder) Request() *AppleUserInitiatedEnrollmentProfileSetPriorityRequest {
	return &AppleUserInitiatedEnrollmentProfileSetPriorityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AppleUserInitiatedEnrollmentProfileSetPriorityRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
