// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// SynchronizationSchema undocumented
type SynchronizationSchema struct {
	// Entity is the base model of SynchronizationSchema
	Entity
	// ProvisioningTaskIdentifier undocumented
	ProvisioningTaskIdentifier *string `json:"provisioningTaskIdentifier,omitempty"`
	// SynchronizationRules undocumented
	SynchronizationRules []SynchronizationRule `json:"synchronizationRules,omitempty"`
	// Version undocumented
	Version *string `json:"version,omitempty"`
	// Directories undocumented
	Directories []DirectoryDefinition `json:"directories,omitempty"`
}

// SynchronizationSchemaParseExpressionRequestParameter undocumented
type SynchronizationSchemaParseExpressionRequestParameter struct {
	// Expression undocumented
	Expression *string `json:"expression,omitempty"`
	// TestInputObject undocumented
	TestInputObject *ExpressionInputObject `json:"testInputObject,omitempty"`
	// TargetAttributeDefinition undocumented
	TargetAttributeDefinition *AttributeDefinition `json:"targetAttributeDefinition,omitempty"`
}

// SynchronizationSchemaRequestBuilder is request builder for SynchronizationSchema
type SynchronizationSchemaRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationSchemaRequest
func (b *SynchronizationSchemaRequestBuilder) Request() *SynchronizationSchemaRequest {
	return &SynchronizationSchemaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationSchemaRequest is request for SynchronizationSchema
type SynchronizationSchemaRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationSchema
func (r *SynchronizationSchemaRequest) Get(ctx context.Context) (resObj *SynchronizationSchema, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationSchema
func (r *SynchronizationSchemaRequest) Update(ctx context.Context, reqObj *SynchronizationSchema) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationSchema
func (r *SynchronizationSchemaRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Directories returns request builder for DirectoryDefinition collection
func (b *SynchronizationSchemaRequestBuilder) Directories() *SynchronizationSchemaDirectoriesCollectionRequestBuilder {
	bb := &SynchronizationSchemaDirectoriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/directories"
	return bb
}

// SynchronizationSchemaDirectoriesCollectionRequestBuilder is request builder for DirectoryDefinition collection
type SynchronizationSchemaDirectoriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryDefinition collection
func (b *SynchronizationSchemaDirectoriesCollectionRequestBuilder) Request() *SynchronizationSchemaDirectoriesCollectionRequest {
	return &SynchronizationSchemaDirectoriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryDefinition item
func (b *SynchronizationSchemaDirectoriesCollectionRequestBuilder) ID(id string) *DirectoryDefinitionRequestBuilder {
	bb := &DirectoryDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SynchronizationSchemaDirectoriesCollectionRequest is request for DirectoryDefinition collection
type SynchronizationSchemaDirectoriesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryDefinition collection
func (r *SynchronizationSchemaDirectoriesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DirectoryDefinition, error) {
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
	var values []DirectoryDefinition
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
			value  []DirectoryDefinition
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

// Get performs GET request for DirectoryDefinition collection
func (r *SynchronizationSchemaDirectoriesCollectionRequest) Get(ctx context.Context) ([]DirectoryDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DirectoryDefinition collection
func (r *SynchronizationSchemaDirectoriesCollectionRequest) Add(ctx context.Context, reqObj *DirectoryDefinition) (resObj *DirectoryDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

//
type SynchronizationSchemaParseExpressionRequestBuilder struct{ BaseRequestBuilder }

// ParseExpression action undocumented
func (b *SynchronizationSchemaRequestBuilder) ParseExpression(reqObj *SynchronizationSchemaParseExpressionRequestParameter) *SynchronizationSchemaParseExpressionRequestBuilder {
	bb := &SynchronizationSchemaParseExpressionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/parseExpression"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SynchronizationSchemaParseExpressionRequest struct{ BaseRequest }

//
func (b *SynchronizationSchemaParseExpressionRequestBuilder) Request() *SynchronizationSchemaParseExpressionRequest {
	return &SynchronizationSchemaParseExpressionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SynchronizationSchemaParseExpressionRequest) Post(ctx context.Context) (resObj *ParseExpressionResponse, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
