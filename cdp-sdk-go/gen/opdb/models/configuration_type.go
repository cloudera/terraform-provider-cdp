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

// ConfigurationType Type of an HBase configuration.
//
//	`SERVICE` - Service level configuration. `MASTER` - Configuration for the Master nodes. `REGIONSERVER` - Configuration for the RegionServer nodes. `STRONGMETA` - Configuration for the StrongMeta RegionServer nodes. `GATEWAY` - Configuration for the Gateway nodes.
//
// swagger:model ConfigurationType
type ConfigurationType string

func NewConfigurationType(value ConfigurationType) *ConfigurationType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConfigurationType.
func (m ConfigurationType) Pointer() *ConfigurationType {
	return &m
}

const (

	// ConfigurationTypeSERVICE captures enum value "SERVICE"
	ConfigurationTypeSERVICE ConfigurationType = "SERVICE"

	// ConfigurationTypeMASTER captures enum value "MASTER"
	ConfigurationTypeMASTER ConfigurationType = "MASTER"

	// ConfigurationTypeREGIONSERVER captures enum value "REGIONSERVER"
	ConfigurationTypeREGIONSERVER ConfigurationType = "REGIONSERVER"

	// ConfigurationTypeSTRONGMETA captures enum value "STRONGMETA"
	ConfigurationTypeSTRONGMETA ConfigurationType = "STRONGMETA"

	// ConfigurationTypeGATEWAY captures enum value "GATEWAY"
	ConfigurationTypeGATEWAY ConfigurationType = "GATEWAY"
)

// for schema
var configurationTypeEnum []interface{}

func init() {
	var res []ConfigurationType
	if err := json.Unmarshal([]byte(`["SERVICE","MASTER","REGIONSERVER","STRONGMETA","GATEWAY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configurationTypeEnum = append(configurationTypeEnum, v)
	}
}

func (m ConfigurationType) validateConfigurationTypeEnum(path, location string, value ConfigurationType) error {
	if err := validate.EnumCase(path, location, value, configurationTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this configuration type
func (m ConfigurationType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigurationTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this configuration type based on context it is used
func (m ConfigurationType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
