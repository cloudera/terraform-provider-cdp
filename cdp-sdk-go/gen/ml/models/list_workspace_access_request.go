// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListWorkspaceAccessRequest Request object for the ListWorkspace method.
//
// swagger:model ListWorkspaceAccessRequest
type ListWorkspaceAccessRequest struct {

	// The environment that the workspace is a member of.
	EnvironmentName string `json:"environmentName,omitempty"`

	// The CRN of the workspace to list access. If CRN is specified only the CRN is used for identifying the workspace, environment and name arguments are ignored.
	WorkspaceCrn string `json:"workspaceCrn,omitempty"`

	// The name of the workspace to list access.
	WorkspaceName string `json:"workspaceName,omitempty"`
}

// Validate validates this list workspace access request
func (m *ListWorkspaceAccessRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list workspace access request based on context it is used
func (m *ListWorkspaceAccessRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListWorkspaceAccessRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListWorkspaceAccessRequest) UnmarshalBinary(b []byte) error {
	var res ListWorkspaceAccessRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}