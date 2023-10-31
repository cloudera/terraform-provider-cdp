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

// DescribeVwRequest Request object for the describeVw method.
//
// swagger:model DescribeVwRequest
type DescribeVwRequest struct {

	// ID of the Virtual Warehouse's cluster.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// ID of the Virtual Warehouse.
	// Required: true
	VwID *string `json:"vwId"`
}

// Validate validates this describe vw request
func (m *DescribeVwRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVwID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeVwRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *DescribeVwRequest) validateVwID(formats strfmt.Registry) error {

	if err := validate.Required("vwId", "body", m.VwID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this describe vw request based on context it is used
func (m *DescribeVwRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeVwRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeVwRequest) UnmarshalBinary(b []byte) error {
	var res DescribeVwRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}