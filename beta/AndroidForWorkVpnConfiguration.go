// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidForWorkVpnConfiguration By providing the configurations in this profile you can instruct the Android device to connect to desired VPN endpoint. By specifying the authentication method and security types expected by VPN endpoint you can make the VPN connection seamless for end user.
type AndroidForWorkVpnConfiguration struct {
	// DeviceConfiguration is the base model of AndroidForWorkVpnConfiguration
	DeviceConfiguration
	// ConnectionName Connection name displayed to the user.
	ConnectionName *string `json:"connectionName,omitempty"`
	// ConnectionType Connection type.
	ConnectionType *AndroidForWorkVpnConnectionType `json:"connectionType,omitempty"`
	// Role Role when connection type is set to Pulse Secure.
	Role *string `json:"role,omitempty"`
	// Realm Realm when connection type is set to Pulse Secure.
	Realm *string `json:"realm,omitempty"`
	// Servers List of VPN Servers on the network. Make sure end users can access these network locations. This collection can contain a maximum of 500 elements.
	Servers []VpnServer `json:"servers,omitempty"`
	// Fingerprint Fingerprint is a string that will be used to verify the VPN server can be trusted, which is only applicable when connection type is Check Point Capsule VPN.
	Fingerprint *string `json:"fingerprint,omitempty"`
	// CustomData Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
	CustomData []KeyValue `json:"customData,omitempty"`
	// CustomKeyValueData Custom data when connection type is set to Citrix. This collection can contain a maximum of 25 elements.
	CustomKeyValueData []KeyValuePair `json:"customKeyValueData,omitempty"`
	// AuthenticationMethod Authentication method.
	AuthenticationMethod *VpnAuthenticationMethod `json:"authenticationMethod,omitempty"`
	// IdentityCertificate undocumented
	IdentityCertificate *AndroidForWorkCertificateProfileBase `json:"identityCertificate,omitempty"`
}

// AndroidForWorkVpnConfigurationRequestBuilder is request builder for AndroidForWorkVpnConfiguration
type AndroidForWorkVpnConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidForWorkVpnConfigurationRequest
func (b *AndroidForWorkVpnConfigurationRequestBuilder) Request() *AndroidForWorkVpnConfigurationRequest {
	return &AndroidForWorkVpnConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidForWorkVpnConfigurationRequest is request for AndroidForWorkVpnConfiguration
type AndroidForWorkVpnConfigurationRequest struct{ BaseRequest }

// Get performs GET request for AndroidForWorkVpnConfiguration
func (r *AndroidForWorkVpnConfigurationRequest) Get(ctx context.Context) (resObj *AndroidForWorkVpnConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidForWorkVpnConfiguration
func (r *AndroidForWorkVpnConfigurationRequest) Update(ctx context.Context, reqObj *AndroidForWorkVpnConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidForWorkVpnConfiguration
func (r *AndroidForWorkVpnConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityCertificate is navigation property
func (b *AndroidForWorkVpnConfigurationRequestBuilder) IdentityCertificate() *AndroidForWorkCertificateProfileBaseRequestBuilder {
	bb := &AndroidForWorkCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificate"
	return bb
}
