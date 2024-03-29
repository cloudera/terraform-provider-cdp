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

// NetworkAwsParams AWS network parameters.
//
// swagger:model NetworkAwsParams
type NetworkAwsParams struct {

	// VPC ids of the specified networks.
	// Required: true
	VpcID *string `json:"vpcId"`
}

// Validate validates this network aws params
func (m *NetworkAwsParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVpcID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkAwsParams) validateVpcID(formats strfmt.Registry) error {

	if err := validate.Required("vpcId", "body", m.VpcID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network aws params based on context it is used
func (m *NetworkAwsParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkAwsParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkAwsParams) UnmarshalBinary(b []byte) error {
	var res NetworkAwsParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
