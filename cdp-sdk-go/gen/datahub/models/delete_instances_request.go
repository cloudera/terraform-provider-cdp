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

// DeleteInstancesRequest Request object for deleting multiple instance from a cluster
//
// swagger:model DeleteInstancesRequest
type DeleteInstancesRequest struct {

	// The name or CRN of the cluster for which instances are to be deleted.
	// Required: true
	ClusterName *string `json:"clusterName"`

	// Whether the termination would be forced or not. If it is true, the termination would not be stopped by other - usually blocking - circumstances. Defaults to false.
	Force bool `json:"force,omitempty"`

	// The instanceIds to be deleted from the cluster.
	// Required: true
	InstanceIds []string `json:"instanceIds"`
}

// Validate validates this delete instances request
func (m *DeleteInstancesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteInstancesRequest) validateClusterName(formats strfmt.Registry) error {

	if err := validate.Required("clusterName", "body", m.ClusterName); err != nil {
		return err
	}

	return nil
}

func (m *DeleteInstancesRequest) validateInstanceIds(formats strfmt.Registry) error {

	if err := validate.Required("instanceIds", "body", m.InstanceIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete instances request based on context it is used
func (m *DeleteInstancesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteInstancesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteInstancesRequest) UnmarshalBinary(b []byte) error {
	var res DeleteInstancesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
