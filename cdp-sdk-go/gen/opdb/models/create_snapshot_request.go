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

// CreateSnapshotRequest Create Snapshot Request.
//
// swagger:model CreateSnapshotRequest
type CreateSnapshotRequest struct {

	// The name of the database.
	// Required: true
	DatabaseName *string `json:"databaseName"`

	// The name of the environment.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// The snapshot location URL on object store.
	// Required: true
	SnapshotLocation *string `json:"snapshotLocation"`

	// Snapshot name unique per database.
	// Required: true
	SnapshotName *string `json:"snapshotName"`

	// The fully qualified table name.
	// Required: true
	TableName *string `json:"tableName"`
}

// Validate validates this create snapshot request
func (m *CreateSnapshotRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabaseName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTableName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSnapshotRequest) validateDatabaseName(formats strfmt.Registry) error {

	if err := validate.Required("databaseName", "body", m.DatabaseName); err != nil {
		return err
	}

	return nil
}

func (m *CreateSnapshotRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *CreateSnapshotRequest) validateSnapshotLocation(formats strfmt.Registry) error {

	if err := validate.Required("snapshotLocation", "body", m.SnapshotLocation); err != nil {
		return err
	}

	return nil
}

func (m *CreateSnapshotRequest) validateSnapshotName(formats strfmt.Registry) error {

	if err := validate.Required("snapshotName", "body", m.SnapshotName); err != nil {
		return err
	}

	return nil
}

func (m *CreateSnapshotRequest) validateTableName(formats strfmt.Registry) error {

	if err := validate.Required("tableName", "body", m.TableName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create snapshot request based on context it is used
func (m *CreateSnapshotRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateSnapshotRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSnapshotRequest) UnmarshalBinary(b []byte) error {
	var res CreateSnapshotRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
