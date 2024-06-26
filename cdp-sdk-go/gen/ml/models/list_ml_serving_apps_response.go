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
)

// ListMlServingAppsResponse Response object for the ListMlServingApps method.
//
// swagger:model ListMlServingAppsResponse
type ListMlServingAppsResponse struct {

	// The list of Apps.
	Apps []*MlServingApp `json:"apps"`
}

// Validate validates this list ml serving apps response
func (m *ListMlServingAppsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListMlServingAppsResponse) validateApps(formats strfmt.Registry) error {
	if swag.IsZero(m.Apps) { // not required
		return nil
	}

	for i := 0; i < len(m.Apps); i++ {
		if swag.IsZero(m.Apps[i]) { // not required
			continue
		}

		if m.Apps[i] != nil {
			if err := m.Apps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list ml serving apps response based on the context it is used
func (m *ListMlServingAppsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListMlServingAppsResponse) contextValidateApps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Apps); i++ {

		if m.Apps[i] != nil {

			if swag.IsZero(m.Apps[i]) { // not required
				return nil
			}

			if err := m.Apps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("apps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("apps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListMlServingAppsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListMlServingAppsResponse) UnmarshalBinary(b []byte) error {
	var res ListMlServingAppsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
