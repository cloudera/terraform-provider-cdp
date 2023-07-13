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

// ListSamlProvidersResponse Response object for a list SAML providers request.
//
// swagger:model ListSamlProvidersResponse
type ListSamlProvidersResponse struct {

	// The token to use when requesting the next set of results. If not present, there are no additional results.
	NextToken string `json:"nextToken,omitempty"`

	// The SAML providers.
	// Required: true
	SamlProviders []*SamlProvider `json:"samlProviders"`
}

// Validate validates this list saml providers response
func (m *ListSamlProvidersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSamlProviders(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListSamlProvidersResponse) validateSamlProviders(formats strfmt.Registry) error {

	if err := validate.Required("samlProviders", "body", m.SamlProviders); err != nil {
		return err
	}

	for i := 0; i < len(m.SamlProviders); i++ {
		if swag.IsZero(m.SamlProviders[i]) { // not required
			continue
		}

		if m.SamlProviders[i] != nil {
			if err := m.SamlProviders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("samlProviders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("samlProviders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list saml providers response based on the context it is used
func (m *ListSamlProvidersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSamlProviders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListSamlProvidersResponse) contextValidateSamlProviders(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SamlProviders); i++ {

		if m.SamlProviders[i] != nil {

			if swag.IsZero(m.SamlProviders[i]) { // not required
				return nil
			}

			if err := m.SamlProviders[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("samlProviders" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("samlProviders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListSamlProvidersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListSamlProvidersResponse) UnmarshalBinary(b []byte) error {
	var res ListSamlProvidersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
