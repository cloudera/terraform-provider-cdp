// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KerberosConfiguration Configuration information to enable Kerberos authentication
//
// swagger:model KerberosConfiguration
type KerberosConfiguration struct {

	// The hostname of the KDC
	KdcHost string `json:"kdcHost,omitempty"`

	// A base64-encoded krb5.conf file
	Krb5Conf string `json:"krb5Conf,omitempty"`

	// The Kerberos realm
	Realm string `json:"realm,omitempty"`
}

// Validate validates this kerberos configuration
func (m *KerberosConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kerberos configuration based on context it is used
func (m *KerberosConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KerberosConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KerberosConfiguration) UnmarshalBinary(b []byte) error {
	var res KerberosConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
