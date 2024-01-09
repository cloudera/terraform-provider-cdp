// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateAzureClusterResponse Response object for the createCluster method.
//
// swagger:model CreateAzureClusterResponse
type CreateAzureClusterResponse struct {

	// ID of new Azure cluster.
	ClusterID string `json:"clusterId,omitempty"`
}

// Validate validates this create azure cluster response
func (m *CreateAzureClusterResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create azure cluster response based on context it is used
func (m *CreateAzureClusterResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAzureClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAzureClusterResponse) UnmarshalBinary(b []byte) error {
	var res CreateAzureClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
