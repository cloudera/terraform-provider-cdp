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

// ImageRequest The image request for the datalake. This must not be set if the runtime parameter is provided. The image ID parameter is required if this is present, but the image catalog name is optional, defaulting to 'cdp-default' if not present.
//
// swagger:model ImageRequest
type ImageRequest struct {

	// The name of the custom image catalog to use.
	CatalogName *string `json:"catalogName,omitempty"`

	// The image ID from the catalog. The corresponding image will be used for the created cluster machines.
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this image request
func (m *ImageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this image request based on context it is used
func (m *ImageRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageRequest) UnmarshalBinary(b []byte) error {
	var res ImageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
