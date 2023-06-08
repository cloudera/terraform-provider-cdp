// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RevokeWorkspaceAccessRequest Request object for the RevokeWorkspace method.
//
// swagger:model RevokeWorkspaceAccessRequest
type RevokeWorkspaceAccessRequest struct {

	// The aws user ARN to revoke access to the corresponding EKS cluster. (Deprecated: Use identifier instead).
	Arn string `json:"arn,omitempty"`

	// The environment that the workspace is a member of.
	EnvironmentName string `json:"environmentName,omitempty"`

	// The cloud provider user id which will be granted access to the workspace's Kubernetes cluster.
	Identifier string `json:"identifier,omitempty"`

	// The CRN of the workspace to revoke access to. If CRN is specified only the CRN is used for identifying the workspace, environment and name arguments are ignored.
	WorkspaceCrn string `json:"workspaceCrn,omitempty"`

	// The name of the workspace to revoke access to.
	WorkspaceName string `json:"workspaceName,omitempty"`
}

// Validate validates this revoke workspace access request
func (m *RevokeWorkspaceAccessRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this revoke workspace access request based on context it is used
func (m *RevokeWorkspaceAccessRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RevokeWorkspaceAccessRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RevokeWorkspaceAccessRequest) UnmarshalBinary(b []byte) error {
	var res RevokeWorkspaceAccessRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
