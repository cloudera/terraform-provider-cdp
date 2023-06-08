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

// GenerateWorkloadAuthTokenRequest Request object for GenerateWorkloadAuthToken method.
//
// swagger:model GenerateWorkloadAuthTokenRequest
type GenerateWorkloadAuthTokenRequest struct {

	// The environment CRN, required by DF.
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// The workload name
	// Required: true
	WorkloadName *WorkloadName `json:"workloadName"`
}

// Validate validates this generate workload auth token request
func (m *GenerateWorkloadAuthTokenRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWorkloadName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenerateWorkloadAuthTokenRequest) validateWorkloadName(formats strfmt.Registry) error {

	if err := validate.Required("workloadName", "body", m.WorkloadName); err != nil {
		return err
	}

	if err := validate.Required("workloadName", "body", m.WorkloadName); err != nil {
		return err
	}

	if m.WorkloadName != nil {
		if err := m.WorkloadName.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workloadName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workloadName")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this generate workload auth token request based on the context it is used
func (m *GenerateWorkloadAuthTokenRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWorkloadName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenerateWorkloadAuthTokenRequest) contextValidateWorkloadName(ctx context.Context, formats strfmt.Registry) error {

	if m.WorkloadName != nil {
		if err := m.WorkloadName.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workloadName")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workloadName")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenerateWorkloadAuthTokenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenerateWorkloadAuthTokenRequest) UnmarshalBinary(b []byte) error {
	var res GenerateWorkloadAuthTokenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
