// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActorResponse A CDP actor (user or machine user).
//
// swagger:model ActorResponse
type ActorResponse struct {

	// Actor CRN.
	Crn string `json:"crn,omitempty"`

	// Email address for users.
	Email string `json:"email,omitempty"`

	// Username for machine users.
	MachineUsername string `json:"machineUsername,omitempty"`

	// Username for users.
	WorkloadUsername string `json:"workloadUsername,omitempty"`
}

// Validate validates this actor response
func (m *ActorResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this actor response based on context it is used
func (m *ActorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActorResponse) UnmarshalBinary(b []byte) error {
	var res ActorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
