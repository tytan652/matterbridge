// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IdentitySet undocumented
type IdentitySet struct {
	// Object is the base model of IdentitySet
	Object
	// Application undocumented
	Application *Identity `json:"application,omitempty"`
	// Device undocumented
	Device *Identity `json:"device,omitempty"`
	// User undocumented
	User *Identity `json:"user,omitempty"`
}