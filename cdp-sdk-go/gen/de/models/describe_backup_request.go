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

// DescribeBackupRequest Request object for Describe Backup command.
//
// swagger:model DescribeBackupRequest
type DescribeBackupRequest struct {

	// The ID of the backup to describe.
	// Required: true
	BackupID *int64 `json:"backupID"`
}

// Validate validates this describe backup request
func (m *DescribeBackupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeBackupRequest) validateBackupID(formats strfmt.Registry) error {

	if err := validate.Required("backupID", "body", m.BackupID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this describe backup request based on context it is used
func (m *DescribeBackupRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DescribeBackupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeBackupRequest) UnmarshalBinary(b []byte) error {
	var res DescribeBackupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
