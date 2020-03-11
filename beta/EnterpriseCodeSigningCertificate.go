// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"time"
)

// EnterpriseCodeSigningCertificate undocumented
type EnterpriseCodeSigningCertificate struct {
	// Entity is the base model of EnterpriseCodeSigningCertificate
	Entity
	// Content The Windows Enterprise Code-Signing Certificate in the raw data format.
	Content *Binary `json:"content,omitempty"`
	// Status The Certificate Status Provisioned or not Provisioned.
	Status *CertificateStatus `json:"status,omitempty"`
	// SubjectName The Subject Name for the cert.
	SubjectName *string `json:"subjectName,omitempty"`
	// Subject The Subject Value for the cert.
	Subject *string `json:"subject,omitempty"`
	// IssuerName The Issuer Name for the cert.
	IssuerName *string `json:"issuerName,omitempty"`
	// Issuer The Issuer value for the cert.
	Issuer *string `json:"issuer,omitempty"`
	// ExpirationDateTime The Cert Expiration Date.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// UploadDateTime The date time of CodeSigning Cert when it is uploaded.
	UploadDateTime *time.Time `json:"uploadDateTime,omitempty"`
}

// EnterpriseCodeSigningCertificateRequestBuilder is request builder for EnterpriseCodeSigningCertificate
type EnterpriseCodeSigningCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns EnterpriseCodeSigningCertificateRequest
func (b *EnterpriseCodeSigningCertificateRequestBuilder) Request() *EnterpriseCodeSigningCertificateRequest {
	return &EnterpriseCodeSigningCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EnterpriseCodeSigningCertificateRequest is request for EnterpriseCodeSigningCertificate
type EnterpriseCodeSigningCertificateRequest struct{ BaseRequest }

// Get performs GET request for EnterpriseCodeSigningCertificate
func (r *EnterpriseCodeSigningCertificateRequest) Get(ctx context.Context) (resObj *EnterpriseCodeSigningCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EnterpriseCodeSigningCertificate
func (r *EnterpriseCodeSigningCertificateRequest) Update(ctx context.Context, reqObj *EnterpriseCodeSigningCertificate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EnterpriseCodeSigningCertificate
func (r *EnterpriseCodeSigningCertificateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
