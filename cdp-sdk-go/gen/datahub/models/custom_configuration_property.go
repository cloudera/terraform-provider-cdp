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

// CustomConfigurationProperty Information about Custom Configuration Property.
//
// swagger:model CustomConfigurationProperty
type CustomConfigurationProperty struct {

	// The name of the custom configuration property.
	// Required: true
	ConfigName *string `json:"configName"`

	// The value of the custom configuration property.
	// Required: true
	ConfigValue *string `json:"configValue"`

	// The role within the service type.
	RoleType string `json:"roleType,omitempty"`

	// The service under which the custom configuration property belongs.
	// Required: true
	ServiceType *string `json:"serviceType"`
}

// Validate validates this custom configuration property
func (m *CustomConfigurationProperty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfigValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomConfigurationProperty) validateConfigName(formats strfmt.Registry) error {

	if err := validate.Required("configName", "body", m.ConfigName); err != nil {
		return err
	}

	return nil
}

func (m *CustomConfigurationProperty) validateConfigValue(formats strfmt.Registry) error {

	if err := validate.Required("configValue", "body", m.ConfigValue); err != nil {
		return err
	}

	return nil
}

func (m *CustomConfigurationProperty) validateServiceType(formats strfmt.Registry) error {

	if err := validate.Required("serviceType", "body", m.ServiceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this custom configuration property based on context it is used
func (m *CustomConfigurationProperty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomConfigurationProperty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomConfigurationProperty) UnmarshalBinary(b []byte) error {
	var res CustomConfigurationProperty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
