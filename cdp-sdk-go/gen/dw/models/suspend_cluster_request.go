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

// SuspendClusterRequest Request object for suspend cluster method.
//
// swagger:model SuspendClusterRequest
type SuspendClusterRequest struct {

	// The ID of the cluster to suspend.
	// Required: true
	ClusterID *string `json:"clusterId"`
}

// Validate validates this suspend cluster request
func (m *SuspendClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SuspendClusterRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this suspend cluster request based on context it is used
func (m *SuspendClusterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SuspendClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuspendClusterRequest) UnmarshalBinary(b []byte) error {
	var res SuspendClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
