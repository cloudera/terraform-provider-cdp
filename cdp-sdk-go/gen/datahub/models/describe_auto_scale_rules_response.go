// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DescribeAutoScaleRulesResponse The response object which describes the AutoScale rules for a DataHub cluster.
//
// swagger:model DescribeAutoScaleRulesResponse
type DescribeAutoScaleRulesResponse struct {

	// The autoscale rules.
	AutoScaleRules *AutoScaleRulesResponse `json:"autoScaleRules,omitempty"`
}

// Validate validates this describe auto scale rules response
func (m *DescribeAutoScaleRulesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoScaleRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeAutoScaleRulesResponse) validateAutoScaleRules(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoScaleRules) { // not required
		return nil
	}

	if m.AutoScaleRules != nil {
		if err := m.AutoScaleRules.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoScaleRules")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoScaleRules")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe auto scale rules response based on the context it is used
func (m *DescribeAutoScaleRulesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoScaleRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeAutoScaleRulesResponse) contextValidateAutoScaleRules(ctx context.Context, formats strfmt.Registry) error {

	if m.AutoScaleRules != nil {

		if swag.IsZero(m.AutoScaleRules) { // not required
			return nil
		}

		if err := m.AutoScaleRules.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoScaleRules")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoScaleRules")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeAutoScaleRulesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeAutoScaleRulesResponse) UnmarshalBinary(b []byte) error {
	var res DescribeAutoScaleRulesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
