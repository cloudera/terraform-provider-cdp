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

// CreateGCPEnvironmentResponse Response object for a create GCP environment request.
//
// swagger:model CreateGCPEnvironmentResponse
type CreateGCPEnvironmentResponse struct {

	// Created environment object.
	// Required: true
	Environment *Environment `json:"environment"`
}

// Validate validates this create g c p environment response
func (m *CreateGCPEnvironmentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateGCPEnvironmentResponse) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create g c p environment response based on the context it is used
func (m *CreateGCPEnvironmentResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnvironment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateGCPEnvironmentResponse) contextValidateEnvironment(ctx context.Context, formats strfmt.Registry) error {

	if m.Environment != nil {
		if err := m.Environment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateGCPEnvironmentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateGCPEnvironmentResponse) UnmarshalBinary(b []byte) error {
	var res CreateGCPEnvironmentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
