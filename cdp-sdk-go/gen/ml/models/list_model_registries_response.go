// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListModelRegistriesResponse List of all available model registries.
//
// swagger:model ListModelRegistriesResponse
type ListModelRegistriesResponse struct {

	// The list of model registry.
	ModelRegistries []*ModelRegistry `json:"modelRegistries"`
}

// Validate validates this list model registries response
func (m *ListModelRegistriesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModelRegistries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListModelRegistriesResponse) validateModelRegistries(formats strfmt.Registry) error {
	if swag.IsZero(m.ModelRegistries) { // not required
		return nil
	}

	for i := 0; i < len(m.ModelRegistries); i++ {
		if swag.IsZero(m.ModelRegistries[i]) { // not required
			continue
		}

		if m.ModelRegistries[i] != nil {
			if err := m.ModelRegistries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modelRegistries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modelRegistries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list model registries response based on the context it is used
func (m *ListModelRegistriesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModelRegistries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListModelRegistriesResponse) contextValidateModelRegistries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ModelRegistries); i++ {

		if m.ModelRegistries[i] != nil {

			if swag.IsZero(m.ModelRegistries[i]) { // not required
				return nil
			}

			if err := m.ModelRegistries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("modelRegistries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("modelRegistries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListModelRegistriesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListModelRegistriesResponse) UnmarshalBinary(b []byte) error {
	var res ListModelRegistriesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
