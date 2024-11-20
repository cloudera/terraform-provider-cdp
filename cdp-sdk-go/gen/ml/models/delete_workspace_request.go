// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeleteWorkspaceRequest Request object for the DeleteWorkspace method.
//
// swagger:model DeleteWorkspaceRequest
type DeleteWorkspaceRequest struct {

	// The environment for the workbench to delete.
	EnvironmentName string `json:"environmentName,omitempty"`

	// Force delete a workbench even if errors occur during deletion. Force delete removes the guarantee that resources in your cloud account will be cleaned up.
	// Required: true
	Force *bool `json:"force"`

	// The remove storage flag indicates weather to keep the backing workbench filesystem storage or remove it during delete.
	RemoveStorage bool `json:"removeStorage,omitempty"`

	// The CRN of the workbench to delete. If CRN is specified only the CRN is used for identifying the workbench, environment and name arguments are ignored.
	WorkspaceCrn string `json:"workspaceCrn,omitempty"`

	// The name of the workbench to delete.
	WorkspaceName string `json:"workspaceName,omitempty"`
}

// Validate validates this delete workspace request
func (m *DeleteWorkspaceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateForce(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteWorkspaceRequest) validateForce(formats strfmt.Registry) error {

	if err := validate.Required("force", "body", m.Force); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete workspace request based on context it is used
func (m *DeleteWorkspaceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteWorkspaceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteWorkspaceRequest) UnmarshalBinary(b []byte) error {
	var res DeleteWorkspaceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
