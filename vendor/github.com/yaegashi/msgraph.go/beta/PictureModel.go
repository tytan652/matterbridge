// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Picture undocumented
type Picture struct {
	// Entity is the base model of Picture
	Entity
	// Width undocumented
	Width *int `json:"width,omitempty"`
	// Height undocumented
	Height *int `json:"height,omitempty"`
	// ContentType undocumented
	ContentType *string `json:"contentType,omitempty"`
	// Content undocumented
	Content *Stream `json:"content,omitempty"`
}