// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SimpleBackup A simple backup entry for listBackup usage
//
// swagger:model SimpleBackup
type SimpleBackup struct {

	// The time when the backup was created.
	BackupCreationTime string `json:"backupCreationTime,omitempty"`

	// The CRN of the backup.
	BackupCrn string `json:"backupCrn,omitempty"`

	// The display name of the backup.
	BackupName string `json:"backupName,omitempty"`
}

// Validate validates this simple backup
func (m *SimpleBackup) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this simple backup based on context it is used
func (m *SimpleBackup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SimpleBackup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimpleBackup) UnmarshalBinary(b []byte) error {
	var res SimpleBackup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
