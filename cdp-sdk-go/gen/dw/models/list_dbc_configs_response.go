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

// ListDbcConfigsResponse Response object for the listDbcConfigs method.
//
// swagger:model ListDbcConfigsResponse
type ListDbcConfigsResponse struct {

	// Configuration history of a service.
	ConfigHistory []*ConfigHistoryItem `json:"configHistory"`
}

// Validate validates this list dbc configs response
func (m *ListDbcConfigsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigHistory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDbcConfigsResponse) validateConfigHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.ConfigHistory) { // not required
		return nil
	}

	for i := 0; i < len(m.ConfigHistory); i++ {
		if swag.IsZero(m.ConfigHistory[i]) { // not required
			continue
		}

		if m.ConfigHistory[i] != nil {
			if err := m.ConfigHistory[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configHistory" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configHistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list dbc configs response based on the context it is used
func (m *ListDbcConfigsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListDbcConfigsResponse) contextValidateConfigHistory(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ConfigHistory); i++ {

		if m.ConfigHistory[i] != nil {

			if swag.IsZero(m.ConfigHistory[i]) { // not required
				return nil
			}

			if err := m.ConfigHistory[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configHistory" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configHistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListDbcConfigsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListDbcConfigsResponse) UnmarshalBinary(b []byte) error {
	var res ListDbcConfigsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}