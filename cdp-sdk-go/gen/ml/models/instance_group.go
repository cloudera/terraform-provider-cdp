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

// InstanceGroup Contains the necessary info for an instance group.
//
// swagger:model InstanceGroup
type InstanceGroup struct {

	// The auto scaling configuration.
	Autoscaling *Autoscaling `json:"autoscaling,omitempty"`

	// The networking rules for the ingress.
	IngressRules []string `json:"ingressRules"`

	// The initial number of instance node.
	InstanceCount int32 `json:"instanceCount,omitempty"`

	// The tier of the instance i.e. on-demand/spot.
	InstanceTier string `json:"instanceTier,omitempty"`

	// The cloud provider instance type for the node instance.
	// Required: true
	InstanceType *string `json:"instanceType"`

	// The unique name for the instance or resource group of the workbench.
	Name string `json:"name,omitempty"`

	// The root volume of the instance.
	RootVolume *RootVolume `json:"rootVolume,omitempty"`
}

// Validate validates this instance group
func (m *InstanceGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoscaling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRootVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroup) validateAutoscaling(formats strfmt.Registry) error {
	if swag.IsZero(m.Autoscaling) { // not required
		return nil
	}

	if m.Autoscaling != nil {
		if err := m.Autoscaling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoscaling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoscaling")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroup) validateInstanceType(formats strfmt.Registry) error {

	if err := validate.Required("instanceType", "body", m.InstanceType); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroup) validateRootVolume(formats strfmt.Registry) error {
	if swag.IsZero(m.RootVolume) { // not required
		return nil
	}

	if m.RootVolume != nil {
		if err := m.RootVolume.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rootVolume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rootVolume")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this instance group based on the context it is used
func (m *InstanceGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoscaling(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRootVolume(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroup) contextValidateAutoscaling(ctx context.Context, formats strfmt.Registry) error {

	if m.Autoscaling != nil {

		if swag.IsZero(m.Autoscaling) { // not required
			return nil
		}

		if err := m.Autoscaling.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoscaling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoscaling")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroup) contextValidateRootVolume(ctx context.Context, formats strfmt.Registry) error {

	if m.RootVolume != nil {

		if swag.IsZero(m.RootVolume) { // not required
			return nil
		}

		if err := m.RootVolume.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rootVolume")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rootVolume")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroup) UnmarshalBinary(b []byte) error {
	var res InstanceGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
