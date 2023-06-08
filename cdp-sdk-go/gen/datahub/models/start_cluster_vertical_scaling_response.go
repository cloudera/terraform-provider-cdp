// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StartClusterVerticalScalingResponse The response object for Data Hub cluster vertical scaling.
//
// swagger:model StartClusterVerticalScalingResponse
type StartClusterVerticalScalingResponse struct {

	// The result of the operation.
	Result string `json:"result,omitempty"`
}

// Validate validates this start cluster vertical scaling response
func (m *StartClusterVerticalScalingResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start cluster vertical scaling response based on context it is used
func (m *StartClusterVerticalScalingResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StartClusterVerticalScalingResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StartClusterVerticalScalingResponse) UnmarshalBinary(b []byte) error {
	var res StartClusterVerticalScalingResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
