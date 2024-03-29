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

// RemoveCoprocessorRequest The request for removing a coprocessor of a table in a database.
//
// swagger:model RemoveCoprocessorRequest
type RemoveCoprocessorRequest struct {

	// The coprocessor canonical name. It is unique per database.
	// Required: true
	CoprocessorCanonicalName *string `json:"coprocessorCanonicalName"`

	// The name or CRN of the database.
	Database string `json:"database,omitempty"`

	// The name or CRN of the environment.
	Environment string `json:"environment,omitempty"`

	// Forcefully remove the coprocessor.
	Force bool `json:"force,omitempty"`

	// Fully qualified table name.
	// Required: true
	TableName *string `json:"tableName"`
}

// Validate validates this remove coprocessor request
func (m *RemoveCoprocessorRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoprocessorCanonicalName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTableName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoveCoprocessorRequest) validateCoprocessorCanonicalName(formats strfmt.Registry) error {

	if err := validate.Required("coprocessorCanonicalName", "body", m.CoprocessorCanonicalName); err != nil {
		return err
	}

	return nil
}

func (m *RemoveCoprocessorRequest) validateTableName(formats strfmt.Registry) error {

	if err := validate.Required("tableName", "body", m.TableName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this remove coprocessor request based on context it is used
func (m *RemoveCoprocessorRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RemoveCoprocessorRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoveCoprocessorRequest) UnmarshalBinary(b []byte) error {
	var res RemoveCoprocessorRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
