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

// SetAzureAuditCredentialRequest Request object for a set Azure audit credential request.
//
// swagger:model SetAzureAuditCredentialRequest
type SetAzureAuditCredentialRequest struct {

	// app based
	// Required: true
	AppBased *SetAzureAuditCredentialRequestAppBased `json:"appBased"`

	// The Azure subscription ID.
	// Required: true
	SubscriptionID *string `json:"subscriptionId"`

	// The Azure AD tenant ID for the Azure subscription.
	// Required: true
	TenantID *string `json:"tenantId"`
}

// Validate validates this set azure audit credential request
func (m *SetAzureAuditCredentialRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppBased(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetAzureAuditCredentialRequest) validateAppBased(formats strfmt.Registry) error {

	if err := validate.Required("appBased", "body", m.AppBased); err != nil {
		return err
	}

	if m.AppBased != nil {
		if err := m.AppBased.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appBased")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appBased")
			}
			return err
		}
	}

	return nil
}

func (m *SetAzureAuditCredentialRequest) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionId", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}

func (m *SetAzureAuditCredentialRequest) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this set azure audit credential request based on the context it is used
func (m *SetAzureAuditCredentialRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppBased(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetAzureAuditCredentialRequest) contextValidateAppBased(ctx context.Context, formats strfmt.Registry) error {

	if m.AppBased != nil {

		if err := m.AppBased.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appBased")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appBased")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetAzureAuditCredentialRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetAzureAuditCredentialRequest) UnmarshalBinary(b []byte) error {
	var res SetAzureAuditCredentialRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// SetAzureAuditCredentialRequestAppBased Additional configurations needed for app-based authentication.
//
// swagger:model SetAzureAuditCredentialRequestAppBased
type SetAzureAuditCredentialRequestAppBased struct {

	// The id of the application registered in Azure.
	// Required: true
	ApplicationID *string `json:"applicationId"`

	// The client secret key (also referred to as application password) for the registered application.
	// Required: true
	SecretKey *string `json:"secretKey"`
}

// Validate validates this set azure audit credential request app based
func (m *SetAzureAuditCredentialRequestAppBased) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetAzureAuditCredentialRequestAppBased) validateApplicationID(formats strfmt.Registry) error {

	if err := validate.Required("appBased"+"."+"applicationId", "body", m.ApplicationID); err != nil {
		return err
	}

	return nil
}

func (m *SetAzureAuditCredentialRequestAppBased) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("appBased"+"."+"secretKey", "body", m.SecretKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this set azure audit credential request app based based on context it is used
func (m *SetAzureAuditCredentialRequestAppBased) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetAzureAuditCredentialRequestAppBased) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetAzureAuditCredentialRequestAppBased) UnmarshalBinary(b []byte) error {
	var res SetAzureAuditCredentialRequestAppBased
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
