// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListClusterTemplatesResponse Response object for list cluster templates request.
//
// swagger:model ListClusterTemplatesResponse
type ListClusterTemplatesResponse struct {

	// The cluster templates.
	// Required: true
	ClusterTemplates []*ClusterTemplateSummary `json:"clusterTemplates"`
}

// Validate validates this list cluster templates response
func (m *ListClusterTemplatesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListClusterTemplatesResponse) validateClusterTemplates(formats strfmt.Registry) error {

	if err := validate.Required("clusterTemplates", "body", m.ClusterTemplates); err != nil {
		return err
	}

	for i := 0; i < len(m.ClusterTemplates); i++ {
		if swag.IsZero(m.ClusterTemplates[i]) { // not required
			continue
		}

		if m.ClusterTemplates[i] != nil {
			if err := m.ClusterTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterTemplates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list cluster templates response based on the context it is used
func (m *ListClusterTemplatesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListClusterTemplatesResponse) contextValidateClusterTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterTemplates); i++ {

		if m.ClusterTemplates[i] != nil {
			if err := m.ClusterTemplates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clusterTemplates" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clusterTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListClusterTemplatesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListClusterTemplatesResponse) UnmarshalBinary(b []byte) error {
	var res ListClusterTemplatesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
