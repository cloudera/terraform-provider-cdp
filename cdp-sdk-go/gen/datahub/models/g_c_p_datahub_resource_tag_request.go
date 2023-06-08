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

// GCPDatahubResourceTagRequest A label that can be attached to GCP Data Hub resources. Please refer to Google documentation for the rules https://cloud.google.com/compute/docs/labeling-resources#label_format.
//
// swagger:model GCPDatahubResourceTagRequest
type GCPDatahubResourceTagRequest struct {

	// The key of tag.
	// Required: true
	Key *string `json:"key"`

	// The value of the tag.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this g c p datahub resource tag request
func (m *GCPDatahubResourceTagRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GCPDatahubResourceTagRequest) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *GCPDatahubResourceTagRequest) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this g c p datahub resource tag request based on context it is used
func (m *GCPDatahubResourceTagRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GCPDatahubResourceTagRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GCPDatahubResourceTagRequest) UnmarshalBinary(b []byte) error {
	var res GCPDatahubResourceTagRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
