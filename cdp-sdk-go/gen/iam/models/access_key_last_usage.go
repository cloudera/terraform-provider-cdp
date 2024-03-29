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

// AccessKeyLastUsage Information on the last time an access key was used.
//
// swagger:model AccessKeyLastUsage
type AccessKeyLastUsage struct {

	// The date when the access key was last used.
	// Format: date-time
	LastUsageDate strfmt.DateTime `json:"lastUsageDate,omitempty"`

	// The name of the service with which this access key was most recently used.
	ServiceName string `json:"serviceName,omitempty"`
}

// Validate validates this access key last usage
func (m *AccessKeyLastUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastUsageDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessKeyLastUsage) validateLastUsageDate(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUsageDate) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUsageDate", "body", "date-time", m.LastUsageDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this access key last usage based on context it is used
func (m *AccessKeyLastUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccessKeyLastUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessKeyLastUsage) UnmarshalBinary(b []byte) error {
	var res AccessKeyLastUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
