// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateBackupRequest Request object for the create backup request.
//
// swagger:model CreateBackupRequest
type CreateBackupRequest struct {

	// Specified name for the backup. If not set, the name will be blank.
	BackupName string `json:"backupName,omitempty"`

	// DEPRECATED in favor of the namespaceNames. Namespace of the potential candidate for backup. If not set, all of the Data Warehouse namespaces will be backed up.
	NamespaceName string `json:"namespaceName,omitempty"`

	// If both namespaceName and namespaceNames are set, the namespaceName will be ignored! A list of namespace of the potential candidates for backup. If not set, all of the Data Warehouse namespaces will be backed up.
	NamespaceNames []string `json:"namespaceNames"`
}

// Validate validates this create backup request
func (m *CreateBackupRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create backup request based on context it is used
func (m *CreateBackupRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateBackupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBackupRequest) UnmarshalBinary(b []byte) error {
	var res CreateBackupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
