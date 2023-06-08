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

// CreateCustomConfigurationsResponse The response object for create custom configurations request.
//
// swagger:model CreateCustomConfigurationsResponse
type CreateCustomConfigurationsResponse struct {

	// The custom configurations.
	// Required: true
	CustomConfigurations *CustomConfigurations `json:"customConfigurations"`
}

// Validate validates this create custom configurations response
func (m *CreateCustomConfigurationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomConfigurations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCustomConfigurationsResponse) validateCustomConfigurations(formats strfmt.Registry) error {

	if err := validate.Required("customConfigurations", "body", m.CustomConfigurations); err != nil {
		return err
	}

	if m.CustomConfigurations != nil {
		if err := m.CustomConfigurations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customConfigurations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customConfigurations")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create custom configurations response based on the context it is used
func (m *CreateCustomConfigurationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomConfigurations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCustomConfigurationsResponse) contextValidateCustomConfigurations(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomConfigurations != nil {
		if err := m.CustomConfigurations.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customConfigurations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customConfigurations")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateCustomConfigurationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCustomConfigurationsResponse) UnmarshalBinary(b []byte) error {
	var res CreateCustomConfigurationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
