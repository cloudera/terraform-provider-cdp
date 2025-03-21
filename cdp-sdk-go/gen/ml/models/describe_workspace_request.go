// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DescribeWorkspaceRequest Request object for the DescribeWorkspace method.
//
// swagger:model DescribeWorkspaceRequest
type DescribeWorkspaceRequest struct {

	// The environment for the workbench to describe.
	EnvironmentName string `json:"environmentName,omitempty"`

	// The CRN of the workbench to describe. If CRN is specified only the CRN is used for identifying the workbench, environment and name arguments are ignored.
	WorkspaceCrn string `json:"workspaceCrn,omitempty"`

	// The name of the workbench to describe.
	WorkspaceName string `json:"workspaceName,omitempty"`
}

// Validate validates this describe workspace request
func (m *DescribeWorkspaceRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this describe workspace request based on context it is used
func (m *DescribeWorkspaceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeWorkspaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeWorkspaceRequest) UnmarshalBinary(b []byte) error {
	var res DescribeWorkspaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
