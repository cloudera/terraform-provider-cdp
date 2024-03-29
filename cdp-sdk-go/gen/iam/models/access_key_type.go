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

// AccessKeyType The version of an access key. `V1` - Deprecated, use RSA as the request signing algorithm. `V2` - Use ED25519 as the request signing algorithm. `V3` - Use ECDSA as the request signing algorithm. `DEFAULT` - Use the system default signing algorithm (V3 in GovCloud, V2 in other regions).
//
// swagger:model AccessKeyType
type AccessKeyType string

func NewAccessKeyType(value AccessKeyType) *AccessKeyType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AccessKeyType.
func (m AccessKeyType) Pointer() *AccessKeyType {
	return &m
}

const (

	// AccessKeyTypeV1 captures enum value "V1"
	AccessKeyTypeV1 AccessKeyType = "V1"

	// AccessKeyTypeV2 captures enum value "V2"
	AccessKeyTypeV2 AccessKeyType = "V2"

	// AccessKeyTypeV3 captures enum value "V3"
	AccessKeyTypeV3 AccessKeyType = "V3"
)

// for schema
var accessKeyTypeEnum []interface{}

func init() {
	var res []AccessKeyType
	if err := json.Unmarshal([]byte(`["V1","V2","V3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessKeyTypeEnum = append(accessKeyTypeEnum, v)
	}
}

func (m AccessKeyType) validateAccessKeyTypeEnum(path, location string, value AccessKeyType) error {
	if err := validate.EnumCase(path, location, value, accessKeyTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this access key type
func (m AccessKeyType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccessKeyTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this access key type based on context it is used
func (m AccessKeyType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
