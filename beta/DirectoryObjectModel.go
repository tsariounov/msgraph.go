// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// DirectoryObject Represents an Azure Active Directory object. The directoryObject type is the base type for many other directory entity types.
type DirectoryObject struct {
	Entity
	// DeletedDateTime undocumented
	DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
}