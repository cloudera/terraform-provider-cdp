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

// DeleteClusterDiagnosticDataJobRequest Request object for the deleteClusterDiagnosticDataJob method.
//
// swagger:model DeleteClusterDiagnosticDataJobRequest
type DeleteClusterDiagnosticDataJobRequest struct {

	// ID of the Cluster.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// ID of the diagnostic job.
	// Required: true
	JobID *string `json:"jobId"`
}

// Validate validates this delete cluster diagnostic data job request
func (m *DeleteClusterDiagnosticDataJobRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteClusterDiagnosticDataJobRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *DeleteClusterDiagnosticDataJobRequest) validateJobID(formats strfmt.Registry) error {

	if err := validate.Required("jobId", "body", m.JobID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete cluster diagnostic data job request based on context it is used
func (m *DeleteClusterDiagnosticDataJobRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteClusterDiagnosticDataJobRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteClusterDiagnosticDataJobRequest) UnmarshalBinary(b []byte) error {
	var res DeleteClusterDiagnosticDataJobRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
