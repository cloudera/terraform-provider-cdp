// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateResourceTemplateResponse The response object for the createResourceTemplate method.
//
// swagger:model CreateResourceTemplateResponse
type CreateResourceTemplateResponse struct {

	// The id of the new template.
	TemplateID string `json:"templateId,omitempty"`
}

// Validate validates this create resource template response
func (m *CreateResourceTemplateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create resource template response based on context it is used
func (m *CreateResourceTemplateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateResourceTemplateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateResourceTemplateResponse) UnmarshalBinary(b []byte) error {
	var res CreateResourceTemplateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
