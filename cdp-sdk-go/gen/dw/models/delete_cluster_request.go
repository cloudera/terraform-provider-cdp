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

// DeleteClusterRequest Request object for the deleteCluster method.
//
// swagger:model DeleteClusterRequest
type DeleteClusterRequest struct {

	// The ID of the cluster to delete.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// Force delete means CDW will delete the cluster even if there are attached DB Catalogs and Virtual Warehouses. All managed data will be lost and will not be recoverable. Force delete attempts all steps of the deletion even if previous steps have failed. NOTICE: Be aware that a Force delete may not remove all resources due to cloud provider constraints. Should this happen, it is responsibility of the user to ensure the impacted resources are deleted on the cloud provider.
	Force *bool `json:"force,omitempty"`
}

// Validate validates this delete cluster request
func (m *DeleteClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteClusterRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete cluster request based on context it is used
func (m *DeleteClusterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteClusterRequest) UnmarshalBinary(b []byte) error {
	var res DeleteClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
