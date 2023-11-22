// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RestoreWorkspaceResponse Response object for the RestoreWorkspace method.
//
// swagger:model RestoreWorkspaceResponse
type RestoreWorkspaceResponse struct {

	// The CRN of the Cloudera Machine Learning workspace being provisioned.
	WorkspaceCrn string `json:"workspaceCrn,omitempty"`
}

// Validate validates this restore workspace response
func (m *RestoreWorkspaceResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this restore workspace response based on context it is used
func (m *RestoreWorkspaceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestoreWorkspaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreWorkspaceResponse) UnmarshalBinary(b []byte) error {
	var res RestoreWorkspaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}