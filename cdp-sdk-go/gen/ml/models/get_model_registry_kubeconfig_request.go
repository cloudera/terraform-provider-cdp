// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetModelRegistryKubeconfigRequest Request object for GetModelRegistryKubeconfig.
//
// swagger:model GetModelRegistryKubeconfigRequest
type GetModelRegistryKubeconfigRequest struct {

	// CRN of the Model Registry
	ModelRegistryCrn string `json:"modelRegistryCrn,omitempty"`
}

// Validate validates this get model registry kubeconfig request
func (m *GetModelRegistryKubeconfigRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get model registry kubeconfig request based on context it is used
func (m *GetModelRegistryKubeconfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetModelRegistryKubeconfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetModelRegistryKubeconfigRequest) UnmarshalBinary(b []byte) error {
	var res GetModelRegistryKubeconfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
