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

// RotateSaltPasswordRequest Request object for rotating SaltStack user password on Data Hub instances (Deprecated).
//
// swagger:model RotateSaltPasswordRequest
type RotateSaltPasswordRequest struct {

	// The name or CRN of the cluster.
	// Required: true
	Cluster *string `json:"cluster"`
}

// Validate validates this rotate salt password request
func (m *RotateSaltPasswordRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RotateSaltPasswordRequest) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rotate salt password request based on context it is used
func (m *RotateSaltPasswordRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RotateSaltPasswordRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RotateSaltPasswordRequest) UnmarshalBinary(b []byte) error {
	var res RotateSaltPasswordRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
