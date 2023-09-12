// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SetAccountMessagesRequest Request object for set account messages for account.
//
// swagger:model SetAccountMessagesRequest
type SetAccountMessagesRequest struct {

	// Message shown to user when user does not have sufficient rights. Length of message cannot be more than 512 characters. If string is empty, default message is displayed.
	ContactYourAdministratorMessage string `json:"contactYourAdministratorMessage,omitempty"`
}

// Validate validates this set account messages request
func (m *SetAccountMessagesRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set account messages request based on context it is used
func (m *SetAccountMessagesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetAccountMessagesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetAccountMessagesRequest) UnmarshalBinary(b []byte) error {
	var res SetAccountMessagesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}