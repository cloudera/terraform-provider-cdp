// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Message The message object used to display warnings and errors during DRS workflows
//
// swagger:model Message
type Message struct {

	// The namespace that has the warning/error.
	Namespace string `json:"namespace,omitempty"`

	// The text message of the warning/error.
	Text string `json:"text,omitempty"`

	// The time when the warning/error is hit.
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this message
func (m *Message) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this message based on context it is used
func (m *Message) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Message) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Message) UnmarshalBinary(b []byte) error {
	var res Message
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
