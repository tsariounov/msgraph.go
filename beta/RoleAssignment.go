// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// RoleAssignment The Role Assignment resource. Role assignments tie together a role definition with members and scopes. There can be one or more role assignments per role. This applies to custom and built-in roles.
type RoleAssignment struct {
	// Entity is the base model of RoleAssignment
	Entity
	// DisplayName The display or friendly name of the role Assignment.
	DisplayName *string `json:"displayName,omitempty"`
	// Description Description of the Role Assignment.
	Description *string `json:"description,omitempty"`
	// ScopeMembers List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
	ScopeMembers []string `json:"scopeMembers,omitempty"`
	// ScopeType Specifies the type of scope for a Role Assignment. Default type 'ResourceScope' allows assignment of ResourceScopes. For 'AllDevices', 'AllLicensedUsers', and 'AllDevicesAndLicensedUsers', the ResourceScopes property should be left empty.
	ScopeType *RoleAssignmentScopeType `json:"scopeType,omitempty"`
	// ResourceScopes List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
	ResourceScopes []string `json:"resourceScopes,omitempty"`
	// RoleDefinition undocumented
	RoleDefinition *RoleDefinition `json:"roleDefinition,omitempty"`
}

// RoleAssignmentRequestBuilder is request builder for RoleAssignment
type RoleAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoleAssignmentRequest
func (b *RoleAssignmentRequestBuilder) Request() *RoleAssignmentRequest {
	return &RoleAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RoleAssignmentRequest is request for RoleAssignment
type RoleAssignmentRequest struct{ BaseRequest }

// Get performs GET request for RoleAssignment
func (r *RoleAssignmentRequest) Get(ctx context.Context) (resObj *RoleAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RoleAssignment
func (r *RoleAssignmentRequest) Update(ctx context.Context, reqObj *RoleAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RoleAssignment
func (r *RoleAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RoleDefinition is navigation property
func (b *RoleAssignmentRequestBuilder) RoleDefinition() *RoleDefinitionRequestBuilder {
	bb := &RoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}
