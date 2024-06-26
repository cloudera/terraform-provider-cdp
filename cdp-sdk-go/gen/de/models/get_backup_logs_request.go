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

// GetBackupLogsRequest Request object for Get Backup Logs command.
//
// swagger:model GetBackupLogsRequest
type GetBackupLogsRequest struct {

	// The ID of the backup.
	// Required: true
	BackupID *int64 `json:"backupID"`

	// The size of each page.
	// Maximum: 1000
	// Minimum: 1
	PageSize int32 `json:"pageSize,omitempty"`

	// A token to specify where to start paginating. This is the nextToken from a previously truncated response.
	StartingToken string `json:"startingToken,omitempty"`
}

// Validate validates this get backup logs request
func (m *GetBackupLogsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePageSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBackupLogsRequest) validateBackupID(formats strfmt.Registry) error {

	if err := validate.Required("backupID", "body", m.BackupID); err != nil {
		return err
	}

	return nil
}

func (m *GetBackupLogsRequest) validatePageSize(formats strfmt.Registry) error {
	if swag.IsZero(m.PageSize) { // not required
		return nil
	}

	if err := validate.MinimumInt("pageSize", "body", int64(m.PageSize), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("pageSize", "body", int64(m.PageSize), 1000, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get backup logs request based on context it is used
func (m *GetBackupLogsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetBackupLogsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBackupLogsRequest) UnmarshalBinary(b []byte) error {
	var res GetBackupLogsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
