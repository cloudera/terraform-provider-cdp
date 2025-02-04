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

// RecipeSummary Information about a recipe.
//
// swagger:model RecipeSummary
type RecipeSummary struct {

	// The CRN of the recipe.
	// Required: true
	Crn *string `json:"crn"`

	// The description of the recipe.
	Description string `json:"description,omitempty"`

	// The name of the recipe.
	// Required: true
	RecipeName *string `json:"recipeName"`

	// The type of recipe. Supported values are : PRE_SERVICE_DEPLOYMENT, PRE_TERMINATION, POST_SERVICE_DEPLOYMENT, POST_CLOUDERA_MANAGER_START.
	Type string `json:"type,omitempty"`
}

// Validate validates this recipe summary
func (m *RecipeSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipeName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecipeSummary) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *RecipeSummary) validateRecipeName(formats strfmt.Registry) error {

	if err := validate.Required("recipeName", "body", m.RecipeName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this recipe summary based on context it is used
func (m *RecipeSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecipeSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecipeSummary) UnmarshalBinary(b []byte) error {
	var res RecipeSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
