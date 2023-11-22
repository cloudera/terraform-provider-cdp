// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConfigBlockReq A piece of configuration stored in the same place (e.g. same file or environment variables).
//
// swagger:model ConfigBlockReq
type ConfigBlockReq struct {

	// Contents of a ConfigBlock.
	// Required: true
	Content *ConfigContentReq `json:"content"`

	// Format of ConfigBlock.
	// Required: true
	// Enum: [HADOOP_XML PROPERTIES TEXT JSON BINARY ENV FLAGFILE]
	Format *string `json:"format"`

	// ID of the ConfigBlock. Unique within an ApplicationConfig.
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this config block req
func (m *ConfigBlockReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigBlockReq) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

var configBlockReqTypeFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HADOOP_XML","PROPERTIES","TEXT","JSON","BINARY","ENV","FLAGFILE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configBlockReqTypeFormatPropEnum = append(configBlockReqTypeFormatPropEnum, v)
	}
}

const (

	// ConfigBlockReqFormatHADOOPXML captures enum value "HADOOP_XML"
	ConfigBlockReqFormatHADOOPXML string = "HADOOP_XML"

	// ConfigBlockReqFormatPROPERTIES captures enum value "PROPERTIES"
	ConfigBlockReqFormatPROPERTIES string = "PROPERTIES"

	// ConfigBlockReqFormatTEXT captures enum value "TEXT"
	ConfigBlockReqFormatTEXT string = "TEXT"

	// ConfigBlockReqFormatJSON captures enum value "JSON"
	ConfigBlockReqFormatJSON string = "JSON"

	// ConfigBlockReqFormatBINARY captures enum value "BINARY"
	ConfigBlockReqFormatBINARY string = "BINARY"

	// ConfigBlockReqFormatENV captures enum value "ENV"
	ConfigBlockReqFormatENV string = "ENV"

	// ConfigBlockReqFormatFLAGFILE captures enum value "FLAGFILE"
	ConfigBlockReqFormatFLAGFILE string = "FLAGFILE"
)

// prop value enum
func (m *ConfigBlockReq) validateFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, configBlockReqTypeFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConfigBlockReq) validateFormat(formats strfmt.Registry) error {

	if err := validate.Required("format", "body", m.Format); err != nil {
		return err
	}

	// value enum
	if err := m.validateFormatEnum("format", "body", *m.Format); err != nil {
		return err
	}

	return nil
}

func (m *ConfigBlockReq) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this config block req based on the context it is used
func (m *ConfigBlockReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigBlockReq) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	if m.Content != nil {

		if err := m.Content.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigBlockReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigBlockReq) UnmarshalBinary(b []byte) error {
	var res ConfigBlockReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}