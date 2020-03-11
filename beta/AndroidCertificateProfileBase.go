// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidCertificateProfileBase Android certificate profile base.
type AndroidCertificateProfileBase struct {
	// DeviceConfiguration is the base model of AndroidCertificateProfileBase
	DeviceConfiguration
	// RenewalThresholdPercentage Certificate renewal threshold percentage. Valid values 1 to 99
	RenewalThresholdPercentage *int `json:"renewalThresholdPercentage,omitempty"`
	// SubjectNameFormat Certificate Subject Name Format.
	SubjectNameFormat *SubjectNameFormat `json:"subjectNameFormat,omitempty"`
	// SubjectAlternativeNameType Certificate Subject Alternative Name Type.
	SubjectAlternativeNameType *SubjectAlternativeNameType `json:"subjectAlternativeNameType,omitempty"`
	// CertificateValidityPeriodValue Value for the Certificate Validity Period.
	CertificateValidityPeriodValue *int `json:"certificateValidityPeriodValue,omitempty"`
	// CertificateValidityPeriodScale Scale for the Certificate Validity Period.
	CertificateValidityPeriodScale *CertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`
	// ExtendedKeyUsages Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
	ExtendedKeyUsages []ExtendedKeyUsage `json:"extendedKeyUsages,omitempty"`
	// RootCertificate undocumented
	RootCertificate *AndroidTrustedRootCertificate `json:"rootCertificate,omitempty"`
}

// AndroidCertificateProfileBaseRequestBuilder is request builder for AndroidCertificateProfileBase
type AndroidCertificateProfileBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidCertificateProfileBaseRequest
func (b *AndroidCertificateProfileBaseRequestBuilder) Request() *AndroidCertificateProfileBaseRequest {
	return &AndroidCertificateProfileBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidCertificateProfileBaseRequest is request for AndroidCertificateProfileBase
type AndroidCertificateProfileBaseRequest struct{ BaseRequest }

// Get performs GET request for AndroidCertificateProfileBase
func (r *AndroidCertificateProfileBaseRequest) Get(ctx context.Context) (resObj *AndroidCertificateProfileBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidCertificateProfileBase
func (r *AndroidCertificateProfileBaseRequest) Update(ctx context.Context, reqObj *AndroidCertificateProfileBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidCertificateProfileBase
func (r *AndroidCertificateProfileBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// RootCertificate is navigation property
func (b *AndroidCertificateProfileBaseRequestBuilder) RootCertificate() *AndroidTrustedRootCertificateRequestBuilder {
	bb := &AndroidTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificate"
	return bb
}
