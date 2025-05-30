// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClusterCreateDiagnosticDataDownloadOptions DEPRECATED: Included by default, no need to specify
//
// swagger:model ClusterCreateDiagnosticDataDownloadOptions
type ClusterCreateDiagnosticDataDownloadOptions struct {

	// DEPRECATED: Included by default, no need to specify
	IncludeClusterInfo *bool `json:"includeClusterInfo,omitempty"`

	// DEPRECATED: Included by default, no need to specify
	IncludeIstioSystem *bool `json:"includeIstioSystem,omitempty"`

	// DEPRECATED: Included by default, no need to specify
	IncludeKubeSystem *bool `json:"includeKubeSystem,omitempty"`

	// DEPRECATED: Included by default, no need to specify
	IncludeSharedServices *bool `json:"includeSharedServices,omitempty"`
}

// Validate validates this cluster create diagnostic data download options
func (m *ClusterCreateDiagnosticDataDownloadOptions) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cluster create diagnostic data download options based on context it is used
func (m *ClusterCreateDiagnosticDataDownloadOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCreateDiagnosticDataDownloadOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCreateDiagnosticDataDownloadOptions) UnmarshalBinary(b []byte) error {
	var res ClusterCreateDiagnosticDataDownloadOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
