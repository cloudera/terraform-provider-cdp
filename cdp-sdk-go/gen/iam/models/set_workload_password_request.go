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

// SetWorkloadPasswordRequest Request object for a set workload password request.
//
// swagger:model SetWorkloadPasswordRequest
type SetWorkloadPasswordRequest struct {

	// The CRN of the user or machine user for whom the password will be set. If it is not included, it defaults to the user making the request.
	ActorCrn string `json:"actorCrn,omitempty"`

	// The password value to set
	// Required: true
	Password *string `json:"password"`
}

// Validate validates this set workload password request
func (m *SetWorkloadPasswordRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetWorkloadPasswordRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this set workload password request based on context it is used
func (m *SetWorkloadPasswordRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetWorkloadPasswordRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetWorkloadPasswordRequest) UnmarshalBinary(b []byte) error {
	var res SetWorkloadPasswordRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}