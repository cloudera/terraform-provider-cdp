// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReplaceRecipesRequest The request for replacing recipes.
//
// swagger:model ReplaceRecipesRequest
type ReplaceRecipesRequest struct {

	// The name or CRN of the datahub.
	// Required: true
	Datahub *string `json:"datahub"`

	// The list of instance group and recipe name pairs.
	// Required: true
	InstanceGroupRecipes []*InstanceGroupRecipeRequest `json:"instanceGroupRecipes"`
}

// Validate validates this replace recipes request
func (m *ReplaceRecipesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatahub(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroupRecipes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplaceRecipesRequest) validateDatahub(formats strfmt.Registry) error {

	if err := validate.Required("datahub", "body", m.Datahub); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceRecipesRequest) validateInstanceGroupRecipes(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroupRecipes", "body", m.InstanceGroupRecipes); err != nil {
		return err
	}

	for i := 0; i < len(m.InstanceGroupRecipes); i++ {
		if swag.IsZero(m.InstanceGroupRecipes[i]) { // not required
			continue
		}

		if m.InstanceGroupRecipes[i] != nil {
			if err := m.InstanceGroupRecipes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroupRecipes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceGroupRecipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this replace recipes request based on the context it is used
func (m *ReplaceRecipesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstanceGroupRecipes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplaceRecipesRequest) contextValidateInstanceGroupRecipes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstanceGroupRecipes); i++ {

		if m.InstanceGroupRecipes[i] != nil {

			if swag.IsZero(m.InstanceGroupRecipes[i]) { // not required
				return nil
			}

			if err := m.InstanceGroupRecipes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroupRecipes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceGroupRecipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplaceRecipesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplaceRecipesRequest) UnmarshalBinary(b []byte) error {
	var res ReplaceRecipesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
