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

// GetEnvironmentUserSyncStateRequest Request object for retrieving the user sync state of an environment.
//
// swagger:model GetEnvironmentUserSyncStateRequest
type GetEnvironmentUserSyncStateRequest struct {

	// The name or CRN of the environment.
	// Required: true
	EnvironmentName *string `json:"environmentName"`
}

// Validate validates this get environment user sync state request
func (m *GetEnvironmentUserSyncStateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetEnvironmentUserSyncStateRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get environment user sync state request based on context it is used
func (m *GetEnvironmentUserSyncStateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetEnvironmentUserSyncStateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetEnvironmentUserSyncStateRequest) UnmarshalBinary(b []byte) error {
	var res GetEnvironmentUserSyncStateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
