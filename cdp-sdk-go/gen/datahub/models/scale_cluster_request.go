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

// ScaleClusterRequest Request object for scale cluster request.
//
// swagger:model ScaleClusterRequest
type ScaleClusterRequest struct {

	// The name or CRN of the cluster to be scaled.
	// Required: true
	ClusterName *string `json:"clusterName"`

	// The desired number of instances in the instance group.
	// Required: true
	InstanceGroupDesiredCount *int32 `json:"instanceGroupDesiredCount"`

	// The name of the instance group which needs to be scaled.
	// Required: true
	InstanceGroupName *string `json:"instanceGroupName"`
}

// Validate validates this scale cluster request
func (m *ScaleClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroupDesiredCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroupName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScaleClusterRequest) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("clusterName", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

func (m *ScaleClusterRequest) validateInstanceGroupDesiredCount(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroupDesiredCount", "body", m.InstanceGroupDesiredCount); err != nil {
		return err
	}

	return nil
}

func (m *ScaleClusterRequest) validateInstanceGroupName(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroupName", "body", m.InstanceGroupName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this scale cluster request based on context it is used
func (m *ScaleClusterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScaleClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScaleClusterRequest) UnmarshalBinary(b []byte) error {
	var res ScaleClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
