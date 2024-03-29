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

// DeleteProxyConfigRequest Request object for a delete proxy config request.
//
// swagger:model DeleteProxyConfigRequest
type DeleteProxyConfigRequest struct {

	// The name or CRN of the proxy config.
	// Required: true
	ProxyConfigName *string `json:"proxyConfigName"`
}

// Validate validates this delete proxy config request
func (m *DeleteProxyConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProxyConfigName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteProxyConfigRequest) validateProxyConfigName(formats strfmt.Registry) error {

	if err := validate.Required("proxyConfigName", "body", m.ProxyConfigName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete proxy config request based on context it is used
func (m *DeleteProxyConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteProxyConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteProxyConfigRequest) UnmarshalBinary(b []byte) error {
	var res DeleteProxyConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
