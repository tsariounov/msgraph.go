// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

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

// MobileAppAssignment A class containing the properties used for Group Assignment of a Mobile App.
type MobileAppAssignment struct {
	// Entity is the base model of MobileAppAssignment
	Entity
	// Intent The install intent defined by the admin.
	Intent *InstallIntent `json:"intent,omitempty"`
	// Target The target group assignment defined by the admin.
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
	// Settings The settings for target assignment defined by the admin.
	Settings *MobileAppAssignmentSettings `json:"settings,omitempty"`
}

// MobileAppAssignmentSettings undocumented
type MobileAppAssignmentSettings struct {
	// Object is the base model of MobileAppAssignmentSettings
	Object
}

// MobileAppCategory Contains properties for a single Intune app category.
type MobileAppCategory struct {
	// Entity is the base model of MobileAppCategory
	Entity
	// DisplayName The name of the app category.
	DisplayName *string `json:"displayName,omitempty"`
	// LastModifiedDateTime The date and time the mobileAppCategory was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

// MobileAppContent Contains content properties for a specific app version. Each mobileAppContent can have multiple mobileAppContentFile.
type MobileAppContent struct {
	// Entity is the base model of MobileAppContent
	Entity
	// Files undocumented
	Files []MobileAppContentFile `json:"files,omitempty"`
}

// MobileAppContentFile Contains properties for a single installer file that is associated with a given mobileAppContent version.
type MobileAppContentFile struct {
	// Entity is the base model of MobileAppContentFile
	Entity
	// AzureStorageURI The Azure Storage URI.
	AzureStorageURI *string `json:"azureStorageUri,omitempty"`
	// IsCommitted A value indicating whether the file is committed.
	IsCommitted *bool `json:"isCommitted,omitempty"`
	// CreatedDateTime The time the file was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Name the file name.
	Name *string `json:"name,omitempty"`
	// Size The size of the file prior to encryption.
	Size *int `json:"size,omitempty"`
	// SizeEncrypted The size of the file after encryption.
	SizeEncrypted *int `json:"sizeEncrypted,omitempty"`
	// AzureStorageURIExpirationDateTime The time the Azure storage Uri expires.
	AzureStorageURIExpirationDateTime *time.Time `json:"azureStorageUriExpirationDateTime,omitempty"`
	// Manifest The manifest information.
	Manifest *Binary `json:"manifest,omitempty"`
	// UploadState The state of the current upload request.
	UploadState *MobileAppContentFileUploadState `json:"uploadState,omitempty"`
}

// MobileAppIdentifier undocumented
type MobileAppIdentifier struct {
	// Object is the base model of MobileAppIdentifier
	Object
}

// MobileLobApp An abstract base class containing properties for all mobile line of business apps.
type MobileLobApp struct {
	// MobileApp is the base model of MobileLobApp
	MobileApp
	// CommittedContentVersion The internal committed content version.
	CommittedContentVersion *string `json:"committedContentVersion,omitempty"`
	// FileName The name of the main Lob application file.
	FileName *string `json:"fileName,omitempty"`
	// Size The total size, including all uploaded files.
	Size *int `json:"size,omitempty"`
	// ContentVersions undocumented
	ContentVersions []MobileAppContent `json:"contentVersions,omitempty"`
}

// MobileThreatDefenseConnector Entity which represents a connection to Mobile threat defense partner.
type MobileThreatDefenseConnector struct {
	// Entity is the base model of MobileThreatDefenseConnector
	Entity
	// LastHeartbeatDateTime DateTime of last Heartbeat recieved from the Data Sync Partner
	LastHeartbeatDateTime *time.Time `json:"lastHeartbeatDateTime,omitempty"`
	// PartnerState Data Sync Partner state for this account
	PartnerState *MobileThreatPartnerTenantState `json:"partnerState,omitempty"`
	// AndroidEnabled For Android, set whether data from the data sync partner should be used during compliance evaluations
	AndroidEnabled *bool `json:"androidEnabled,omitempty"`
	// IOSEnabled For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
	IOSEnabled *bool `json:"iosEnabled,omitempty"`
	// AndroidDeviceBlockedOnMissingPartnerData For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
	AndroidDeviceBlockedOnMissingPartnerData *bool `json:"androidDeviceBlockedOnMissingPartnerData,omitempty"`
	// IOSDeviceBlockedOnMissingPartnerData For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
	IOSDeviceBlockedOnMissingPartnerData *bool `json:"iosDeviceBlockedOnMissingPartnerData,omitempty"`
	// PartnerUnsupportedOsVersionBlocked Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
	PartnerUnsupportedOsVersionBlocked *bool `json:"partnerUnsupportedOsVersionBlocked,omitempty"`
	// PartnerUnresponsivenessThresholdInDays Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
	PartnerUnresponsivenessThresholdInDays *int `json:"partnerUnresponsivenessThresholdInDays,omitempty"`
}
