// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DisableS3GuardResponse Response object for disabling S3Guard in an environment.
//
// swagger:model DisableS3GuardResponse
type DisableS3GuardResponse struct {

	// Response status for disabling S3Guard in an environment.
	S3GuardResponse string `json:"s3GuardResponse,omitempty"`
}

// Validate validates this disable s3 guard response
func (m *DisableS3GuardResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this disable s3 guard response based on context it is used
func (m *DisableS3GuardResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DisableS3GuardResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DisableS3GuardResponse) UnmarshalBinary(b []byte) error {
	var res DisableS3GuardResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}