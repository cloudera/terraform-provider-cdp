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

// BackupDatalakeRequest Request object to perform a backup of datalake.
//
// swagger:model BackupDatalakeRequest
type BackupDatalakeRequest struct {

	// Location where the back-up has to be stored. For example s3a://Location/of/the/backup.
	BackupLocation string `json:"backupLocation,omitempty"`

	// The name of the backup.
	BackupName string `json:"backupName,omitempty"`

	// Close the database connections while performing backup. Default is true.
	CloseDbConnections *bool `json:"closeDbConnections,omitempty"`

	// The name of the datalake.
	// Required: true
	DatalakeName *string `json:"datalakeName"`

	// Skips the backup of the Atlas indexes. If this option or --skipAtlasMetadata is not provided, Atlas indexes are backed up by default. Redundant if --skipAtlasMetadata is included.
	SkipAtlasIndexes bool `json:"skipAtlasIndexes,omitempty"`

	// Skips the backup of the Atlas metadata. If this option is not provided, the Atlas metadata is backed up by default.
	SkipAtlasMetadata bool `json:"skipAtlasMetadata,omitempty"`

	// Skips the backup of the Ranger audits. If this option is not provided, Ranger audits are backed up by default.
	SkipRangerAudits bool `json:"skipRangerAudits,omitempty"`

	// Skips the backup of the databases backing HMS/Ranger services. If this option is not provided, the HMS/Ranger services are backed up by default.
	SkipRangerHmsMetadata bool `json:"skipRangerHmsMetadata,omitempty"`

	// Skips the validation steps that run prior to the backup. If this option is not provided, the validations are performed by default.
	SkipValidation bool `json:"skipValidation,omitempty"`

	// Runs only the validation steps and then returns. If this option is not provided, the backup is performed as normal by default.
	ValidationOnly bool `json:"validationOnly,omitempty"`
}

// Validate validates this backup datalake request
func (m *BackupDatalakeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalakeName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupDatalakeRequest) validateDatalakeName(formats strfmt.Registry) error {

	if err := validate.Required("datalakeName", "body", m.DatalakeName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this backup datalake request based on context it is used
func (m *BackupDatalakeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BackupDatalakeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupDatalakeRequest) UnmarshalBinary(b []byte) error {
	var res BackupDatalakeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}