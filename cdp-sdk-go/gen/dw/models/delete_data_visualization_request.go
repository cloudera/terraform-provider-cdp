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

// DeleteDataVisualizationRequest Request object for the deleteDataVisualization method.
//
// swagger:model DeleteDataVisualizationRequest
type DeleteDataVisualizationRequest struct {

	// ID of the Cloudera Data Visualization's cluster.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// ID of the Cloudera Data Visualization to delete from the cluster.
	// Required: true
	DataVisualizationID *string `json:"dataVisualizationId"`
}

// Validate validates this delete data visualization request
func (m *DeleteDataVisualizationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataVisualizationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteDataVisualizationRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *DeleteDataVisualizationRequest) validateDataVisualizationID(formats strfmt.Registry) error {

	if err := validate.Required("dataVisualizationId", "body", m.DataVisualizationID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete data visualization request based on context it is used
func (m *DeleteDataVisualizationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteDataVisualizationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteDataVisualizationRequest) UnmarshalBinary(b []byte) error {
	var res DeleteDataVisualizationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
