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

// EnableServiceResponse Response object for EnableService request.
//
// swagger:model EnableServiceResponse
type EnableServiceResponse struct {

	// Service details of the created service.
	Service *ServiceDescription `json:"service,omitempty"`
}

// Validate validates this enable service response
func (m *EnableServiceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnableServiceResponse) validateService(formats strfmt.Registry) error {
	if swag.IsZero(m.Service) { // not required
		return nil
	}

	if m.Service != nil {
		if err := m.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this enable service response based on the context it is used
func (m *EnableServiceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnableServiceResponse) contextValidateService(ctx context.Context, formats strfmt.Registry) error {

	if m.Service != nil {

		if swag.IsZero(m.Service) { // not required
			return nil
		}

		if err := m.Service.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnableServiceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnableServiceResponse) UnmarshalBinary(b []byte) error {
	var res EnableServiceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
