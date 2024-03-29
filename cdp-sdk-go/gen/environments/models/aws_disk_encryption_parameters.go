// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AwsDiskEncryptionParameters Object containing details of encryption parameters for AWS cloud.
//
// swagger:model AwsDiskEncryptionParameters
type AwsDiskEncryptionParameters struct {

	// ARN of the CMK which is used to encrypt the AWS EBS volumes.
	EncryptionKeyArn string `json:"encryptionKeyArn,omitempty"`
}

// Validate validates this aws disk encryption parameters
func (m *AwsDiskEncryptionParameters) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aws disk encryption parameters based on context it is used
func (m *AwsDiskEncryptionParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AwsDiskEncryptionParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsDiskEncryptionParameters) UnmarshalBinary(b []byte) error {
	var res AwsDiskEncryptionParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
