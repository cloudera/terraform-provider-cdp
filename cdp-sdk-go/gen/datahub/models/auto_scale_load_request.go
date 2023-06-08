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

// AutoScaleLoadRequest Configuration for Load Based Scaling
//
// swagger:model AutoScaleLoadRequest
type AutoScaleLoadRequest struct {

	// Configuration for Load Based Scaling
	// Required: true
	Configuration *AutoScaleLoadRequestConfiguration `json:"configuration"`

	// An optional description for the specific schedule.
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// An optional identifier for the rule. Generally useful for debugging. Will be auto-generated if none is provided.
	// Max Length: 200
	// Min Length: 5
	Identifier string `json:"identifier,omitempty"`
}

// Validate validates this auto scale load request
func (m *AutoScaleLoadRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoScaleLoadRequest) validateConfiguration(formats strfmt.Registry) error {

	if err := validate.Required("configuration", "body", m.Configuration); err != nil {
		return err
	}

	if m.Configuration != nil {
		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *AutoScaleLoadRequest) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", *m.Description, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *m.Description, 1000); err != nil {
		return err
	}

	return nil
}

func (m *AutoScaleLoadRequest) validateIdentifier(formats strfmt.Registry) error {
	if swag.IsZero(m.Identifier) { // not required
		return nil
	}

	if err := validate.MinLength("identifier", "body", m.Identifier, 5); err != nil {
		return err
	}

	if err := validate.MaxLength("identifier", "body", m.Identifier, 200); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this auto scale load request based on the context it is used
func (m *AutoScaleLoadRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoScaleLoadRequest) contextValidateConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.Configuration != nil {
		if err := m.Configuration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoScaleLoadRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoScaleLoadRequest) UnmarshalBinary(b []byte) error {
	var res AutoScaleLoadRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
