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

// GetAccessKeyResponse Response object for a get access key request.
//
// swagger:model GetAccessKeyResponse
type GetAccessKeyResponse struct {

	// Information about the access key.
	// Required: true
	AccessKey *AccessKey `json:"accessKey"`
}

// Validate validates this get access key response
func (m *GetAccessKeyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAccessKeyResponse) validateAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("accessKey", "body", m.AccessKey); err != nil {
		return err
	}

	if m.AccessKey != nil {
		if err := m.AccessKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessKey")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get access key response based on the context it is used
func (m *GetAccessKeyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAccessKeyResponse) contextValidateAccessKey(ctx context.Context, formats strfmt.Registry) error {

	if m.AccessKey != nil {

		if err := m.AccessKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accessKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("accessKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetAccessKeyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAccessKeyResponse) UnmarshalBinary(b []byte) error {
	var res GetAccessKeyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}