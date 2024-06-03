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

// SMTPConfigResponse SMTP config response object
//
// swagger:model SmtpConfigResponse
type SMTPConfigResponse struct {

	// Sender's email address.
	// Required: true
	Email *string `json:"email"`

	// SMTP host.
	// Required: true
	Host *string `json:"host"`

	// SMTP port.
	Port *int32 `json:"port,omitempty"`

	// Use SSL to secure the connection to the email server.
	Ssl *bool `json:"ssl,omitempty"`

	// Use SMTP STARTTLS command to encrypt the mail.
	StartTLS *bool `json:"startTls,omitempty"`

	// SMTP username.
	Username string `json:"username,omitempty"`
}

// Validate validates this Smtp config response
func (m *SMTPConfigResponse) Validate(formats strfmt.Registry) error {
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

func (m *SMTPConfigResponse) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *SMTPConfigResponse) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Smtp config response based on context it is used
func (m *SMTPConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SMTPConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SMTPConfigResponse) UnmarshalBinary(b []byte) error {
	var res SMTPConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
