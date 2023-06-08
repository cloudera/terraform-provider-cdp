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

// RestartDbcRequest Request object for the restartDbc method.
//
// swagger:model RestartDbcRequest
type RestartDbcRequest struct {

	// ID of the Database Catalog's cluster.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// The id of the Database Catalog to restart.
	// Required: true
	DbcID *string `json:"dbcId"`
}

// Validate validates this restart dbc request
func (m *RestartDbcRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDbcID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestartDbcRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *RestartDbcRequest) validateDbcID(formats strfmt.Registry) error {

	if err := validate.Required("dbcId", "body", m.DbcID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this restart dbc request based on context it is used
func (m *RestartDbcRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestartDbcRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestartDbcRequest) UnmarshalBinary(b []byte) error {
	var res RestartDbcRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
