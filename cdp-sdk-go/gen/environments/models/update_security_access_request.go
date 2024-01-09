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

// UpdateSecurityAccessRequest The request object for updating security access of the given environment.
//
// swagger:model UpdateSecurityAccessRequest
type UpdateSecurityAccessRequest struct {

	// Security group ID for non-gateway nodes.
	// Required: true
	DefaultSecurityGroupID *string `json:"defaultSecurityGroupId"`

	// The name or the CRN of the environment.
	// Required: true
	Environment *string `json:"environment"`

	// Security group ID where Knox-enabled hosts are placed.
	// Required: true
	GatewayNodeSecurityGroupID *string `json:"gatewayNodeSecurityGroupId"`
}

// Validate validates this update security access request
func (m *UpdateSecurityAccessRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultSecurityGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayNodeSecurityGroupID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateSecurityAccessRequest) validateDefaultSecurityGroupID(formats strfmt.Registry) error {

	if err := validate.Required("defaultSecurityGroupId", "body", m.DefaultSecurityGroupID); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSecurityAccessRequest) validateEnvironment(formats strfmt.Registry) error {

	if err := validate.Required("environment", "body", m.Environment); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSecurityAccessRequest) validateGatewayNodeSecurityGroupID(formats strfmt.Registry) error {

	if err := validate.Required("gatewayNodeSecurityGroupId", "body", m.GatewayNodeSecurityGroupID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update security access request based on context it is used
func (m *UpdateSecurityAccessRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateSecurityAccessRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateSecurityAccessRequest) UnmarshalBinary(b []byte) error {
	var res UpdateSecurityAccessRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
