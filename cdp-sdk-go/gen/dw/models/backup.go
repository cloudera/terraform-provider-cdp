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
)

// Backup Backup entry
//
// swagger:model Backup
type Backup struct {

	// The time when the backup was created.
	BackupCreationTime string `json:"backupCreationTime,omitempty"`

	// The CRN of the backup.
	BackupCrn string `json:"backupCrn,omitempty"`

	// The backup job name.
	BackupJob string `json:"backupJob,omitempty"`

	// The current state of the backup job.
	BackupJobState string `json:"backupJobState,omitempty"`

	// The display name of the backup.
	BackupName string `json:"backupName,omitempty"`

	// The phase of the backup operation.
	BackupPhase string `json:"backupPhase,omitempty"`

	// The time when the backup was updated.
	BackupUpdatedTime string `json:"backupUpdatedTime,omitempty"`

	// The errors from backup job.
	Errors []*Message `json:"errors"`

	// The list of namespaces to be included in backup.
	IncludedNamespaces []string `json:"includedNamespaces"`

	// The warnings from backup job.
	Warnings []*Message `json:"warnings"`
}

// Validate validates this backup
func (m *Backup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarnings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Backup) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Backup) validateWarnings(formats strfmt.Registry) error {
	if swag.IsZero(m.Warnings) { // not required
		return nil
	}

	for i := 0; i < len(m.Warnings); i++ {
		if swag.IsZero(m.Warnings[i]) { // not required
			continue
		}

		if m.Warnings[i] != nil {
			if err := m.Warnings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("warnings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("warnings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this backup based on the context it is used
func (m *Backup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWarnings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Backup) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Backup) contextValidateWarnings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Warnings); i++ {

		if m.Warnings[i] != nil {
			if err := m.Warnings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("warnings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("warnings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Backup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Backup) UnmarshalBinary(b []byte) error {
	var res Backup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
