// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateDbcDiagnosticDataJobResponse Response object for the createDbcDiagnosticDataJob method.
//
// swagger:model CreateDbcDiagnosticDataJobResponse
type CreateDbcDiagnosticDataJobResponse struct {

	// Identifier for each bundle collection.
	ID string `json:"id,omitempty"`

	// Additional key-value pair attributes associated with the Diagnostic Data Job.
	Labels map[string]string `json:"labels,omitempty"`

	// Status of the diagnostics collection request.
	Status string `json:"status,omitempty"`

	// This URL points to a download location if the destination is DOWNLOAD.
	URL string `json:"url,omitempty"`
}

// Validate validates this create dbc diagnostic data job response
func (m *CreateDbcDiagnosticDataJobResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create dbc diagnostic data job response based on context it is used
func (m *CreateDbcDiagnosticDataJobResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateDbcDiagnosticDataJobResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDbcDiagnosticDataJobResponse) UnmarshalBinary(b []byte) error {
	var res CreateDbcDiagnosticDataJobResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
