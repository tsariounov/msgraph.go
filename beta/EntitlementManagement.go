// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// EntitlementManagement undocumented
type EntitlementManagement struct {
	// Entity is the base model of EntitlementManagement
	Entity
	// AccessPackageCatalogs undocumented
	AccessPackageCatalogs []AccessPackageCatalog `json:"accessPackageCatalogs,omitempty"`
	// AccessPackageResources undocumented
	AccessPackageResources []AccessPackageResource `json:"accessPackageResources,omitempty"`
	// AccessPackageResourceRequests undocumented
	AccessPackageResourceRequests []AccessPackageResourceRequestObject `json:"accessPackageResourceRequests,omitempty"`
	// AccessPackageResourceRoleScopes undocumented
	AccessPackageResourceRoleScopes []AccessPackageResourceRoleScope `json:"accessPackageResourceRoleScopes,omitempty"`
	// AccessPackages undocumented
	AccessPackages []AccessPackage `json:"accessPackages,omitempty"`
	// AccessPackageAssignmentPolicies undocumented
	AccessPackageAssignmentPolicies []AccessPackageAssignmentPolicy `json:"accessPackageAssignmentPolicies,omitempty"`
	// AccessPackageAssignments undocumented
	AccessPackageAssignments []AccessPackageAssignment `json:"accessPackageAssignments,omitempty"`
	// AccessPackageAssignmentRequests undocumented
	AccessPackageAssignmentRequests []AccessPackageAssignmentRequestObject `json:"accessPackageAssignmentRequests,omitempty"`
	// AccessPackageAssignmentResourceRoles undocumented
	AccessPackageAssignmentResourceRoles []AccessPackageAssignmentResourceRole `json:"accessPackageAssignmentResourceRoles,omitempty"`
}

// EntitlementManagementRequestBuilder is request builder for EntitlementManagement
type EntitlementManagementRequestBuilder struct{ BaseRequestBuilder }

// Request returns EntitlementManagementRequest
func (b *EntitlementManagementRequestBuilder) Request() *EntitlementManagementRequest {
	return &EntitlementManagementRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EntitlementManagementRequest is request for EntitlementManagement
type EntitlementManagementRequest struct{ BaseRequest }

// Get performs GET request for EntitlementManagement
func (r *EntitlementManagementRequest) Get(ctx context.Context) (resObj *EntitlementManagement, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EntitlementManagement
func (r *EntitlementManagementRequest) Update(ctx context.Context, reqObj *EntitlementManagement) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EntitlementManagement
func (r *EntitlementManagementRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AccessPackageAssignmentPolicies returns request builder for AccessPackageAssignmentPolicy collection
func (b *EntitlementManagementRequestBuilder) AccessPackageAssignmentPolicies() *EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequestBuilder {
	bb := &EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageAssignmentPolicies"
	return bb
}

// EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequestBuilder is request builder for AccessPackageAssignmentPolicy collection
type EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageAssignmentPolicy collection
func (b *EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequestBuilder) Request() *EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequest {
	return &EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageAssignmentPolicy item
func (b *EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequestBuilder) ID(id string) *AccessPackageAssignmentPolicyRequestBuilder {
	bb := &AccessPackageAssignmentPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequest is request for AccessPackageAssignmentPolicy collection
type EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageAssignmentPolicy collection
func (r *EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageAssignmentPolicy, error) {
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
	var values []AccessPackageAssignmentPolicy
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
			value  []AccessPackageAssignmentPolicy
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

// Get performs GET request for AccessPackageAssignmentPolicy collection
func (r *EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequest) Get(ctx context.Context) ([]AccessPackageAssignmentPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageAssignmentPolicy collection
func (r *EntitlementManagementAccessPackageAssignmentPoliciesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageAssignmentPolicy) (resObj *AccessPackageAssignmentPolicy, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackageAssignmentRequests returns request builder for AccessPackageAssignmentRequestObject collection
func (b *EntitlementManagementRequestBuilder) AccessPackageAssignmentRequests() *EntitlementManagementAccessPackageAssignmentRequestsCollectionRequestBuilder {
	bb := &EntitlementManagementAccessPackageAssignmentRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageAssignmentRequests"
	return bb
}

// EntitlementManagementAccessPackageAssignmentRequestsCollectionRequestBuilder is request builder for AccessPackageAssignmentRequestObject collection
type EntitlementManagementAccessPackageAssignmentRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageAssignmentRequestObject collection
func (b *EntitlementManagementAccessPackageAssignmentRequestsCollectionRequestBuilder) Request() *EntitlementManagementAccessPackageAssignmentRequestsCollectionRequest {
	return &EntitlementManagementAccessPackageAssignmentRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageAssignmentRequestObject item
func (b *EntitlementManagementAccessPackageAssignmentRequestsCollectionRequestBuilder) ID(id string) *AccessPackageAssignmentRequestObjectRequestBuilder {
	bb := &AccessPackageAssignmentRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAccessPackageAssignmentRequestsCollectionRequest is request for AccessPackageAssignmentRequestObject collection
type EntitlementManagementAccessPackageAssignmentRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageAssignmentRequestObject collection
func (r *EntitlementManagementAccessPackageAssignmentRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageAssignmentRequestObject, error) {
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
	var values []AccessPackageAssignmentRequestObject
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
			value  []AccessPackageAssignmentRequestObject
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

// Get performs GET request for AccessPackageAssignmentRequestObject collection
func (r *EntitlementManagementAccessPackageAssignmentRequestsCollectionRequest) Get(ctx context.Context) ([]AccessPackageAssignmentRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageAssignmentRequestObject collection
func (r *EntitlementManagementAccessPackageAssignmentRequestsCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageAssignmentRequestObject) (resObj *AccessPackageAssignmentRequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackageAssignmentResourceRoles returns request builder for AccessPackageAssignmentResourceRole collection
func (b *EntitlementManagementRequestBuilder) AccessPackageAssignmentResourceRoles() *EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequestBuilder {
	bb := &EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageAssignmentResourceRoles"
	return bb
}

// EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequestBuilder is request builder for AccessPackageAssignmentResourceRole collection
type EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageAssignmentResourceRole collection
func (b *EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequestBuilder) Request() *EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequest {
	return &EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageAssignmentResourceRole item
func (b *EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequestBuilder) ID(id string) *AccessPackageAssignmentResourceRoleRequestBuilder {
	bb := &AccessPackageAssignmentResourceRoleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequest is request for AccessPackageAssignmentResourceRole collection
type EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageAssignmentResourceRole collection
func (r *EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageAssignmentResourceRole, error) {
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
	var values []AccessPackageAssignmentResourceRole
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
			value  []AccessPackageAssignmentResourceRole
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

// Get performs GET request for AccessPackageAssignmentResourceRole collection
func (r *EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequest) Get(ctx context.Context) ([]AccessPackageAssignmentResourceRole, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageAssignmentResourceRole collection
func (r *EntitlementManagementAccessPackageAssignmentResourceRolesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageAssignmentResourceRole) (resObj *AccessPackageAssignmentResourceRole, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackageAssignments returns request builder for AccessPackageAssignment collection
func (b *EntitlementManagementRequestBuilder) AccessPackageAssignments() *EntitlementManagementAccessPackageAssignmentsCollectionRequestBuilder {
	bb := &EntitlementManagementAccessPackageAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageAssignments"
	return bb
}

// EntitlementManagementAccessPackageAssignmentsCollectionRequestBuilder is request builder for AccessPackageAssignment collection
type EntitlementManagementAccessPackageAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageAssignment collection
func (b *EntitlementManagementAccessPackageAssignmentsCollectionRequestBuilder) Request() *EntitlementManagementAccessPackageAssignmentsCollectionRequest {
	return &EntitlementManagementAccessPackageAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageAssignment item
func (b *EntitlementManagementAccessPackageAssignmentsCollectionRequestBuilder) ID(id string) *AccessPackageAssignmentRequestBuilder {
	bb := &AccessPackageAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAccessPackageAssignmentsCollectionRequest is request for AccessPackageAssignment collection
type EntitlementManagementAccessPackageAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageAssignment collection
func (r *EntitlementManagementAccessPackageAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageAssignment, error) {
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
	var values []AccessPackageAssignment
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
			value  []AccessPackageAssignment
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

// Get performs GET request for AccessPackageAssignment collection
func (r *EntitlementManagementAccessPackageAssignmentsCollectionRequest) Get(ctx context.Context) ([]AccessPackageAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageAssignment collection
func (r *EntitlementManagementAccessPackageAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageAssignment) (resObj *AccessPackageAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackageCatalogs returns request builder for AccessPackageCatalog collection
func (b *EntitlementManagementRequestBuilder) AccessPackageCatalogs() *EntitlementManagementAccessPackageCatalogsCollectionRequestBuilder {
	bb := &EntitlementManagementAccessPackageCatalogsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageCatalogs"
	return bb
}

// EntitlementManagementAccessPackageCatalogsCollectionRequestBuilder is request builder for AccessPackageCatalog collection
type EntitlementManagementAccessPackageCatalogsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageCatalog collection
func (b *EntitlementManagementAccessPackageCatalogsCollectionRequestBuilder) Request() *EntitlementManagementAccessPackageCatalogsCollectionRequest {
	return &EntitlementManagementAccessPackageCatalogsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageCatalog item
func (b *EntitlementManagementAccessPackageCatalogsCollectionRequestBuilder) ID(id string) *AccessPackageCatalogRequestBuilder {
	bb := &AccessPackageCatalogRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAccessPackageCatalogsCollectionRequest is request for AccessPackageCatalog collection
type EntitlementManagementAccessPackageCatalogsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageCatalog collection
func (r *EntitlementManagementAccessPackageCatalogsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageCatalog, error) {
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
	var values []AccessPackageCatalog
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
			value  []AccessPackageCatalog
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

// Get performs GET request for AccessPackageCatalog collection
func (r *EntitlementManagementAccessPackageCatalogsCollectionRequest) Get(ctx context.Context) ([]AccessPackageCatalog, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageCatalog collection
func (r *EntitlementManagementAccessPackageCatalogsCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageCatalog) (resObj *AccessPackageCatalog, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackageResourceRequests returns request builder for AccessPackageResourceRequestObject collection
func (b *EntitlementManagementRequestBuilder) AccessPackageResourceRequests() *EntitlementManagementAccessPackageResourceRequestsCollectionRequestBuilder {
	bb := &EntitlementManagementAccessPackageResourceRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResourceRequests"
	return bb
}

// EntitlementManagementAccessPackageResourceRequestsCollectionRequestBuilder is request builder for AccessPackageResourceRequestObject collection
type EntitlementManagementAccessPackageResourceRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageResourceRequestObject collection
func (b *EntitlementManagementAccessPackageResourceRequestsCollectionRequestBuilder) Request() *EntitlementManagementAccessPackageResourceRequestsCollectionRequest {
	return &EntitlementManagementAccessPackageResourceRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageResourceRequestObject item
func (b *EntitlementManagementAccessPackageResourceRequestsCollectionRequestBuilder) ID(id string) *AccessPackageResourceRequestObjectRequestBuilder {
	bb := &AccessPackageResourceRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAccessPackageResourceRequestsCollectionRequest is request for AccessPackageResourceRequestObject collection
type EntitlementManagementAccessPackageResourceRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageResourceRequestObject collection
func (r *EntitlementManagementAccessPackageResourceRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageResourceRequestObject, error) {
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
	var values []AccessPackageResourceRequestObject
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
			value  []AccessPackageResourceRequestObject
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

// Get performs GET request for AccessPackageResourceRequestObject collection
func (r *EntitlementManagementAccessPackageResourceRequestsCollectionRequest) Get(ctx context.Context) ([]AccessPackageResourceRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageResourceRequestObject collection
func (r *EntitlementManagementAccessPackageResourceRequestsCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageResourceRequestObject) (resObj *AccessPackageResourceRequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackageResourceRoleScopes returns request builder for AccessPackageResourceRoleScope collection
func (b *EntitlementManagementRequestBuilder) AccessPackageResourceRoleScopes() *EntitlementManagementAccessPackageResourceRoleScopesCollectionRequestBuilder {
	bb := &EntitlementManagementAccessPackageResourceRoleScopesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResourceRoleScopes"
	return bb
}

// EntitlementManagementAccessPackageResourceRoleScopesCollectionRequestBuilder is request builder for AccessPackageResourceRoleScope collection
type EntitlementManagementAccessPackageResourceRoleScopesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageResourceRoleScope collection
func (b *EntitlementManagementAccessPackageResourceRoleScopesCollectionRequestBuilder) Request() *EntitlementManagementAccessPackageResourceRoleScopesCollectionRequest {
	return &EntitlementManagementAccessPackageResourceRoleScopesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageResourceRoleScope item
func (b *EntitlementManagementAccessPackageResourceRoleScopesCollectionRequestBuilder) ID(id string) *AccessPackageResourceRoleScopeRequestBuilder {
	bb := &AccessPackageResourceRoleScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAccessPackageResourceRoleScopesCollectionRequest is request for AccessPackageResourceRoleScope collection
type EntitlementManagementAccessPackageResourceRoleScopesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageResourceRoleScope collection
func (r *EntitlementManagementAccessPackageResourceRoleScopesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageResourceRoleScope, error) {
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
	var values []AccessPackageResourceRoleScope
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
			value  []AccessPackageResourceRoleScope
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

// Get performs GET request for AccessPackageResourceRoleScope collection
func (r *EntitlementManagementAccessPackageResourceRoleScopesCollectionRequest) Get(ctx context.Context) ([]AccessPackageResourceRoleScope, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageResourceRoleScope collection
func (r *EntitlementManagementAccessPackageResourceRoleScopesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageResourceRoleScope) (resObj *AccessPackageResourceRoleScope, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackageResources returns request builder for AccessPackageResource collection
func (b *EntitlementManagementRequestBuilder) AccessPackageResources() *EntitlementManagementAccessPackageResourcesCollectionRequestBuilder {
	bb := &EntitlementManagementAccessPackageResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResources"
	return bb
}

// EntitlementManagementAccessPackageResourcesCollectionRequestBuilder is request builder for AccessPackageResource collection
type EntitlementManagementAccessPackageResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageResource collection
func (b *EntitlementManagementAccessPackageResourcesCollectionRequestBuilder) Request() *EntitlementManagementAccessPackageResourcesCollectionRequest {
	return &EntitlementManagementAccessPackageResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageResource item
func (b *EntitlementManagementAccessPackageResourcesCollectionRequestBuilder) ID(id string) *AccessPackageResourceRequestBuilder {
	bb := &AccessPackageResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAccessPackageResourcesCollectionRequest is request for AccessPackageResource collection
type EntitlementManagementAccessPackageResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageResource collection
func (r *EntitlementManagementAccessPackageResourcesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageResource, error) {
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
	var values []AccessPackageResource
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
			value  []AccessPackageResource
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

// Get performs GET request for AccessPackageResource collection
func (r *EntitlementManagementAccessPackageResourcesCollectionRequest) Get(ctx context.Context) ([]AccessPackageResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageResource collection
func (r *EntitlementManagementAccessPackageResourcesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageResource) (resObj *AccessPackageResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackages returns request builder for AccessPackage collection
func (b *EntitlementManagementRequestBuilder) AccessPackages() *EntitlementManagementAccessPackagesCollectionRequestBuilder {
	bb := &EntitlementManagementAccessPackagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackages"
	return bb
}

// EntitlementManagementAccessPackagesCollectionRequestBuilder is request builder for AccessPackage collection
type EntitlementManagementAccessPackagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackage collection
func (b *EntitlementManagementAccessPackagesCollectionRequestBuilder) Request() *EntitlementManagementAccessPackagesCollectionRequest {
	return &EntitlementManagementAccessPackagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackage item
func (b *EntitlementManagementAccessPackagesCollectionRequestBuilder) ID(id string) *AccessPackageRequestBuilder {
	bb := &AccessPackageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EntitlementManagementAccessPackagesCollectionRequest is request for AccessPackage collection
type EntitlementManagementAccessPackagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackage collection
func (r *EntitlementManagementAccessPackagesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackage, error) {
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
	var values []AccessPackage
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
			value  []AccessPackage
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

// Get performs GET request for AccessPackage collection
func (r *EntitlementManagementAccessPackagesCollectionRequest) Get(ctx context.Context) ([]AccessPackage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackage collection
func (r *EntitlementManagementAccessPackagesCollectionRequest) Add(ctx context.Context, reqObj *AccessPackage) (resObj *AccessPackage, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
