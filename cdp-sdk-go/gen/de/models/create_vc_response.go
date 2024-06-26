// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateVcResponse Response object for CreateVc method.
//
// swagger:model CreateVcResponse
type CreateVcResponse struct {

	// Created Virtual Cluster
	Vc *VcDescription `json:"Vc,omitempty"`
}

// Validate validates this create vc response
func (m *CreateVcResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateVcResponse) validateVc(formats strfmt.Registry) error {
	if swag.IsZero(m.Vc) { // not required
		return nil
	}

	if m.Vc != nil {
		if err := m.Vc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Vc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Vc")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create vc response based on the context it is used
func (m *CreateVcResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateVcResponse) contextValidateVc(ctx context.Context, formats strfmt.Registry) error {

	if m.Vc != nil {

		if swag.IsZero(m.Vc) { // not required
			return nil
		}

		if err := m.Vc.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Vc")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Vc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateVcResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateVcResponse) UnmarshalBinary(b []byte) error {
	var res CreateVcResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
