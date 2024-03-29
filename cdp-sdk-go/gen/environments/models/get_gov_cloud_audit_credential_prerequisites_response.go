// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetGovCloudAuditCredentialPrerequisitesResponse The audit credential prerequisites for GovCloud for the enabled providers.
//
// swagger:model GetGovCloudAuditCredentialPrerequisitesResponse
type GetGovCloudAuditCredentialPrerequisitesResponse struct {

	// The provider specific identifier of the account/subscription/project.
	AccountID string `json:"accountId,omitempty"`

	// Provides the external id and policy JSON (this one encoded in base64) for AWS credential creation.
	Aws *AwsCredentialPrerequisitesResponse `json:"aws,omitempty"`
}

// Validate validates this get gov cloud audit credential prerequisites response
func (m *GetGovCloudAuditCredentialPrerequisitesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetGovCloudAuditCredentialPrerequisitesResponse) validateAws(formats strfmt.Registry) error {
	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get gov cloud audit credential prerequisites response based on the context it is used
func (m *GetGovCloudAuditCredentialPrerequisitesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetGovCloudAuditCredentialPrerequisitesResponse) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	if m.Aws != nil {

		if swag.IsZero(m.Aws) { // not required
			return nil
		}

		if err := m.Aws.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetGovCloudAuditCredentialPrerequisitesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetGovCloudAuditCredentialPrerequisitesResponse) UnmarshalBinary(b []byte) error {
	var res GetGovCloudAuditCredentialPrerequisitesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
