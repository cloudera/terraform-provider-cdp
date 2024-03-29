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

// DescribeScalingActivityRequest Request object for describing a particular scaling activity using clusterCrn or clusterName and operationId.
//
// swagger:model DescribeScalingActivityRequest
type DescribeScalingActivityRequest struct {

	// The name or CRN of the cluster.
	// Required: true
	Cluster *string `json:"cluster"`

	// Operation ID of the scaling activity.
	// Required: true
	OperationID *string `json:"operationId"`
}

// Validate validates this describe scaling activity request
func (m *DescribeScalingActivityRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeScalingActivityRequest) validateCluster(formats strfmt.Registry) error {

	if err := validate.Required("cluster", "body", m.Cluster); err != nil {
		return err
	}

	return nil
}

func (m *DescribeScalingActivityRequest) validateOperationID(formats strfmt.Registry) error {

	if err := validate.Required("operationId", "body", m.OperationID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this describe scaling activity request based on context it is used
func (m *DescribeScalingActivityRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeScalingActivityRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeScalingActivityRequest) UnmarshalBinary(b []byte) error {
	var res DescribeScalingActivityRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
