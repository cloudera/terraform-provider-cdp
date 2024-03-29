// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListRestoresRequest Request object for the list restores request.
//
// swagger:model ListRestoresRequest
type ListRestoresRequest struct {

	// CRN of the backup.
	BackupCrn string `json:"backupCrn,omitempty"`

	// The job states to filter by.
	JobStates []string `json:"jobStates"`
}

// Validate validates this list restores request
func (m *ListRestoresRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list restores request based on context it is used
func (m *ListRestoresRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListRestoresRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListRestoresRequest) UnmarshalBinary(b []byte) error {
	var res ListRestoresRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
