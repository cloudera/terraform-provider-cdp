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

// DescribeClusterDefinitionRequest Request object for describe cluster definition request.
//
// swagger:model DescribeClusterDefinitionRequest
type DescribeClusterDefinitionRequest struct {

	// The name or CRN of the cluster definition.
	// Required: true
	ClusterDefinitionName *string `json:"clusterDefinitionName"`
}

// Validate validates this describe cluster definition request
func (m *DescribeClusterDefinitionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterDefinitionName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeClusterDefinitionRequest) validateClusterDefinitionName(formats strfmt.Registry) error {

	if err := validate.Required("clusterDefinitionName", "body", m.ClusterDefinitionName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this describe cluster definition request based on context it is used
func (m *DescribeClusterDefinitionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeClusterDefinitionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeClusterDefinitionRequest) UnmarshalBinary(b []byte) error {
	var res DescribeClusterDefinitionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
