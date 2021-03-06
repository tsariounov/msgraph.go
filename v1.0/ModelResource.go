// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ResourceAccess undocumented
type ResourceAccess struct {
	// Object is the base model of ResourceAccess
	Object
	// ID undocumented
	ID *UUID `json:"id,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

// ResourceAction undocumented
type ResourceAction struct {
	// Object is the base model of ResourceAction
	Object
	// AllowedResourceActions Allowed Actions
	AllowedResourceActions []string `json:"allowedResourceActions,omitempty"`
	// NotAllowedResourceActions Not Allowed Actions
	NotAllowedResourceActions []string `json:"notAllowedResourceActions,omitempty"`
}

// ResourceOperation This defines an operation or action that can be performed on an Intune resource (or entity).  Common operations are Read, Delete, Update or Create.  These operations provide basic management of the underlying Intune resource itself.  In some cases, an Intune resource can have operations that are used by the resource to take action in combination with other resources.  For example, the Assign operation is used to assign a MobileApp resource to an AAD security group.  Resource operations cannot be modified for built-in roles.This defines an operation or action that can be performed on an Intune resource (or entity).  Common operations are Get, List, Delete, Update or Create.  These operations provide basic management of the underlying Intune resource itself.  In some cases, an Intune resource can have operations that are used by the resource to take action in combination with other resources.  For example, the "Assign" operation is used to assign a MobileApp resource to an AAD security group.  Resource operations cannot be modified for built-in roles.
type ResourceOperation struct {
	// Entity is the base model of ResourceOperation
	Entity
	// ResourceName Name of the Resource this operation is performed on.
	ResourceName *string `json:"resourceName,omitempty"`
	// ActionName Type of action this operation is going to perform. The actionName should be concise and limited to as few words as possible.
	ActionName *string `json:"actionName,omitempty"`
	// Description Description of the resource operation. The description is used in mouse-over text for the operation when shown in the Azure Portal.
	Description *string `json:"description,omitempty"`
}

// ResourceReference undocumented
type ResourceReference struct {
	// Object is the base model of ResourceReference
	Object
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

// ResourceVisualization undocumented
type ResourceVisualization struct {
	// Object is the base model of ResourceVisualization
	Object
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// MediaType undocumented
	MediaType *string `json:"mediaType,omitempty"`
	// PreviewImageURL undocumented
	PreviewImageURL *string `json:"previewImageUrl,omitempty"`
	// PreviewText undocumented
	PreviewText *string `json:"previewText,omitempty"`
	// ContainerWebURL undocumented
	ContainerWebURL *string `json:"containerWebUrl,omitempty"`
	// ContainerDisplayName undocumented
	ContainerDisplayName *string `json:"containerDisplayName,omitempty"`
	// ContainerType undocumented
	ContainerType *string `json:"containerType,omitempty"`
}
