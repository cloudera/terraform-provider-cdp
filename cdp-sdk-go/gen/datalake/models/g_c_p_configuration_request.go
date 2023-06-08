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

// GCPConfigurationRequest Request object for GCP configuration.
//
// swagger:model GCPConfigurationRequest
type GCPConfigurationRequest struct {

	// Email id of the service account to be associated with the datalake IdBroker instance. This service account should have "token.creator" role for one or more storage accounts that has access to storage.
	// Required: true
	ServiceAccountEmail *string `json:"serviceAccountEmail"`

	// The location of the GCS bucket to be used as storage. The location has to start with gs:// followed by the bucket name.
	// Required: true
	StorageLocation *string `json:"storageLocation"`
}

// Validate validates this g c p configuration request
func (m *GCPConfigurationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceAccountEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GCPConfigurationRequest) validateServiceAccountEmail(formats strfmt.Registry) error {

	if err := validate.Required("serviceAccountEmail", "body", m.ServiceAccountEmail); err != nil {
		return err
	}

	return nil
}

func (m *GCPConfigurationRequest) validateStorageLocation(formats strfmt.Registry) error {

	if err := validate.Required("storageLocation", "body", m.StorageLocation); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this g c p configuration request based on context it is used
func (m *GCPConfigurationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GCPConfigurationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GCPConfigurationRequest) UnmarshalBinary(b []byte) error {
	var res GCPConfigurationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
