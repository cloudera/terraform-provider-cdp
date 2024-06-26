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

// SMTPConfigRequest SMTP config request object.
//
// swagger:model SmtpConfigRequest
type SMTPConfigRequest struct {

	// Sender's email address.
	// Required: true
	Email *string `json:"email"`

	// SMTP host.
	// Required: true
	Host *string `json:"host"`

	// SMTP password.
	Password string `json:"password,omitempty"`

	// SMTP port.
	Port *int32 `json:"port,omitempty"`

	// Use SSL to secure the connection to the email server.
	Ssl *bool `json:"ssl,omitempty"`

	// Use SMTP STARTTLS command to encrypt the mail.
	StartTLS *bool `json:"startTls,omitempty"`

	// SMTP username.
	Username string `json:"username,omitempty"`
}

// Validate validates this Smtp config request
func (m *SMTPConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SMTPConfigRequest) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *SMTPConfigRequest) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Smtp config request based on context it is used
func (m *SMTPConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SMTPConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SMTPConfigRequest) UnmarshalBinary(b []byte) error {
	var res SMTPConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
