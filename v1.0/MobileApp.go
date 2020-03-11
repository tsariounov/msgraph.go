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

// MobileApp An abstract class containing the base properties for Intune mobile apps.
type MobileApp struct {
	// Entity is the base model of MobileApp
	Entity
	// DisplayName The admin provided or imported title of the app.
	DisplayName *string `json:"displayName,omitempty"`
	// Description The description of the app.
	Description *string `json:"description,omitempty"`
	// Publisher The publisher of the app.
	Publisher *string `json:"publisher,omitempty"`
	// LargeIcon The large icon, to be displayed in the app details and used for upload of the icon.
	LargeIcon *MimeContent `json:"largeIcon,omitempty"`
	// CreatedDateTime The date and time the app was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime The date and time the app was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// IsFeatured The value indicating whether the app is marked as featured by the admin.
	IsFeatured *bool `json:"isFeatured,omitempty"`
	// PrivacyInformationURL The privacy statement Url.
	PrivacyInformationURL *string `json:"privacyInformationUrl,omitempty"`
	// InformationURL The more information Url.
	InformationURL *string `json:"informationUrl,omitempty"`
	// Owner The owner of the app.
	Owner *string `json:"owner,omitempty"`
	// Developer The developer of the app.
	Developer *string `json:"developer,omitempty"`
	// Notes Notes for the app.
	Notes *string `json:"notes,omitempty"`
	// PublishingState The publishing state for the app. The app cannot be assigned unless the app is published.
	PublishingState *MobileAppPublishingState `json:"publishingState,omitempty"`
	// Categories undocumented
	Categories []MobileAppCategory `json:"categories,omitempty"`
	// Assignments undocumented
	Assignments []MobileAppAssignment `json:"assignments,omitempty"`
}

// MobileAppAssignRequestParameter undocumented
type MobileAppAssignRequestParameter struct {
	// MobileAppAssignments undocumented
	MobileAppAssignments []MobileAppAssignment `json:"mobileAppAssignments,omitempty"`
}

// MobileAppRequestBuilder is request builder for MobileApp
type MobileAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppRequest
func (b *MobileAppRequestBuilder) Request() *MobileAppRequest {
	return &MobileAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppRequest is request for MobileApp
type MobileAppRequest struct{ BaseRequest }

// Get performs GET request for MobileApp
func (r *MobileAppRequest) Get(ctx context.Context) (resObj *MobileApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileApp
func (r *MobileAppRequest) Update(ctx context.Context, reqObj *MobileApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileApp
func (r *MobileAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for MobileAppAssignment collection
func (b *MobileAppRequestBuilder) Assignments() *MobileAppAssignmentsCollectionRequestBuilder {
	bb := &MobileAppAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// MobileAppAssignmentsCollectionRequestBuilder is request builder for MobileAppAssignment collection
type MobileAppAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppAssignment collection
func (b *MobileAppAssignmentsCollectionRequestBuilder) Request() *MobileAppAssignmentsCollectionRequest {
	return &MobileAppAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppAssignment item
func (b *MobileAppAssignmentsCollectionRequestBuilder) ID(id string) *MobileAppAssignmentRequestBuilder {
	bb := &MobileAppAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppAssignmentsCollectionRequest is request for MobileAppAssignment collection
type MobileAppAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MobileAppAssignment, error) {
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
	var values []MobileAppAssignment
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
			value  []MobileAppAssignment
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

// Get performs GET request for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Get(ctx context.Context) ([]MobileAppAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MobileAppAssignment collection
func (r *MobileAppAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *MobileAppAssignment) (resObj *MobileAppAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Categories returns request builder for MobileAppCategory collection
func (b *MobileAppRequestBuilder) Categories() *MobileAppCategoriesCollectionRequestBuilder {
	bb := &MobileAppCategoriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/categories"
	return bb
}

// MobileAppCategoriesCollectionRequestBuilder is request builder for MobileAppCategory collection
type MobileAppCategoriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppCategory collection
func (b *MobileAppCategoriesCollectionRequestBuilder) Request() *MobileAppCategoriesCollectionRequest {
	return &MobileAppCategoriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppCategory item
func (b *MobileAppCategoriesCollectionRequestBuilder) ID(id string) *MobileAppCategoryRequestBuilder {
	bb := &MobileAppCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppCategoriesCollectionRequest is request for MobileAppCategory collection
type MobileAppCategoriesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]MobileAppCategory, error) {
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
	var values []MobileAppCategory
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
			value  []MobileAppCategory
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

// Get performs GET request for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Get(ctx context.Context) ([]MobileAppCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for MobileAppCategory collection
func (r *MobileAppCategoriesCollectionRequest) Add(ctx context.Context, reqObj *MobileAppCategory) (resObj *MobileAppCategory, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

//
type MobileAppAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *MobileAppRequestBuilder) Assign(reqObj *MobileAppAssignRequestParameter) *MobileAppAssignRequestBuilder {
	bb := &MobileAppAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MobileAppAssignRequest struct{ BaseRequest }

//
func (b *MobileAppAssignRequestBuilder) Request() *MobileAppAssignRequest {
	return &MobileAppAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MobileAppAssignRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
