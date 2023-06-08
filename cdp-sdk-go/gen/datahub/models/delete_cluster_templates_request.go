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

// DeleteClusterTemplatesRequest Request object for delete cluster templates request.
//
// swagger:model DeleteClusterTemplatesRequest
type DeleteClusterTemplatesRequest struct {

	// The names or CRNs of the cluster templates to be deleted.
	// Required: true
	ClusterTemplateNames []string `json:"clusterTemplateNames"`
}

// Validate validates this delete cluster templates request
func (m *DeleteClusterTemplatesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterTemplateNames(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeleteClusterTemplatesRequest) validateClusterTemplateNames(formats strfmt.Registry) error {

	if err := validate.Required("clusterTemplateNames", "body", m.ClusterTemplateNames); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this delete cluster templates request based on context it is used
func (m *DeleteClusterTemplatesRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteClusterTemplatesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteClusterTemplatesRequest) UnmarshalBinary(b []byte) error {
	var res DeleteClusterTemplatesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
