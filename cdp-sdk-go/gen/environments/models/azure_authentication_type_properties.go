// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AzureAuthenticationTypeProperties Authentication type of the credential
//
// swagger:model AzureAuthenticationTypeProperties
type AzureAuthenticationTypeProperties string

func NewAzureAuthenticationTypeProperties(value AzureAuthenticationTypeProperties) *AzureAuthenticationTypeProperties {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AzureAuthenticationTypeProperties.
func (m AzureAuthenticationTypeProperties) Pointer() *AzureAuthenticationTypeProperties {
	return &m
}

const (

	// AzureAuthenticationTypePropertiesCERTIFICATE captures enum value "CERTIFICATE"
	AzureAuthenticationTypePropertiesCERTIFICATE AzureAuthenticationTypeProperties = "CERTIFICATE"

	// AzureAuthenticationTypePropertiesSECRET captures enum value "SECRET"
	AzureAuthenticationTypePropertiesSECRET AzureAuthenticationTypeProperties = "SECRET"
)

// for schema
var azureAuthenticationTypePropertiesEnum []interface{}

func init() {
	var res []AzureAuthenticationTypeProperties
	if err := json.Unmarshal([]byte(`["CERTIFICATE","SECRET"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		azureAuthenticationTypePropertiesEnum = append(azureAuthenticationTypePropertiesEnum, v)
	}
}

func (m AzureAuthenticationTypeProperties) validateAzureAuthenticationTypePropertiesEnum(path, location string, value AzureAuthenticationTypeProperties) error {
	if err := validate.EnumCase(path, location, value, azureAuthenticationTypePropertiesEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this azure authentication type properties
func (m AzureAuthenticationTypeProperties) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAzureAuthenticationTypePropertiesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this azure authentication type properties based on context it is used
func (m AzureAuthenticationTypeProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
