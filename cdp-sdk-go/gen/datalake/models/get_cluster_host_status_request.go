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

// GetClusterHostStatusRequest Request object to get host status.
//
// swagger:model GetClusterHostStatusRequest
type GetClusterHostStatusRequest struct {

	// The name or CRN of the cluster.
	// Required: true
	ClusterName *string `json:"clusterName"`
}

// Validate validates this get cluster host status request
func (m *GetClusterHostStatusRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetClusterHostStatusRequest) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("clusterName", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get cluster host status request based on context it is used
func (m *GetClusterHostStatusRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetClusterHostStatusRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetClusterHostStatusRequest) UnmarshalBinary(b []byte) error {
	var res GetClusterHostStatusRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
