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

// GetAuditCredentialPrerequisitesResponse The audit credential prerequisites.
//
// swagger:model GetAuditCredentialPrerequisitesResponse
type GetAuditCredentialPrerequisitesResponse struct {

	// The provider specific identifier of the account/subscription/project that is used by Cloudbreak.
	AccountID string `json:"accountId,omitempty"`

	// Provides the external id and policy JSON (this one encoded in base64) for AWS credential creation.
	Aws *AwsCredentialPrerequisitesResponse `json:"aws,omitempty"`

	// Provides the app creation command and role definition Json for Azure credential creation.
	Azure *AzureCredentialPrerequisitesResponse `json:"azure,omitempty"`

	// The name of the given cloud platform
	// Required: true
	CloudPlatform *string `json:"cloudPlatform"`

	// Provides the service account creation command for GCP credential creation.
	Gcp *GcpCredentialPrerequisitesResponse `json:"gcp,omitempty"`
}

// Validate validates this get audit credential prerequisites response
func (m *GetAuditCredentialPrerequisitesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAuditCredentialPrerequisitesResponse) validateAws(formats strfmt.Registry) error {
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

func (m *GetAuditCredentialPrerequisitesResponse) validateAzure(formats strfmt.Registry) error {
	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *GetAuditCredentialPrerequisitesResponse) validateCloudPlatform(formats strfmt.Registry) error {

	if err := validate.Required("cloudPlatform", "body", m.CloudPlatform); err != nil {
		return err
	}

	return nil
}

func (m *GetAuditCredentialPrerequisitesResponse) validateGcp(formats strfmt.Registry) error {
	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get audit credential prerequisites response based on the context it is used
func (m *GetAuditCredentialPrerequisitesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAuditCredentialPrerequisitesResponse) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetAuditCredentialPrerequisitesResponse) contextValidateAzure(ctx context.Context, formats strfmt.Registry) error {

	if m.Azure != nil {

		if swag.IsZero(m.Azure) { // not required
			return nil
		}

		if err := m.Azure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *GetAuditCredentialPrerequisitesResponse) contextValidateGcp(ctx context.Context, formats strfmt.Registry) error {

	if m.Gcp != nil {

		if swag.IsZero(m.Gcp) { // not required
			return nil
		}

		if err := m.Gcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetAuditCredentialPrerequisitesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAuditCredentialPrerequisitesResponse) UnmarshalBinary(b []byte) error {
	var res GetAuditCredentialPrerequisitesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
