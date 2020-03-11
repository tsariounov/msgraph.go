// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MacOSTrustedRootCertificate OS X Trusted Root Certificate configuration profile.
type MacOSTrustedRootCertificate struct {
	// DeviceConfiguration is the base model of MacOSTrustedRootCertificate
	DeviceConfiguration
	// TrustedRootCertificate Trusted Root Certificate.
	TrustedRootCertificate *Binary `json:"trustedRootCertificate,omitempty"`
	// CertFileName File name to display in UI.
	CertFileName *string `json:"certFileName,omitempty"`
}

// MacOSTrustedRootCertificateRequestBuilder is request builder for MacOSTrustedRootCertificate
type MacOSTrustedRootCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSTrustedRootCertificateRequest
func (b *MacOSTrustedRootCertificateRequestBuilder) Request() *MacOSTrustedRootCertificateRequest {
	return &MacOSTrustedRootCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSTrustedRootCertificateRequest is request for MacOSTrustedRootCertificate
type MacOSTrustedRootCertificateRequest struct{ BaseRequest }

// Get performs GET request for MacOSTrustedRootCertificate
func (r *MacOSTrustedRootCertificateRequest) Get(ctx context.Context) (resObj *MacOSTrustedRootCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSTrustedRootCertificate
func (r *MacOSTrustedRootCertificateRequest) Update(ctx context.Context, reqObj *MacOSTrustedRootCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSTrustedRootCertificate
func (r *MacOSTrustedRootCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
