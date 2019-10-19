// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IOSLobAppProvisioningConfigurationAssignment A class containing the properties used for Group Assignment of an iOS LOB App Provisioning and Configuration.
type IOSLobAppProvisioningConfigurationAssignment struct {
	Entity
	// Target The target group assignment defined by the admin.
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}