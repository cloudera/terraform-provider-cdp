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
	"github.com/go-openapi/validate"
)

// ServicePrincipalCloudIdentities Cloud identity mappings for a service principal.
//
// swagger:model ServicePrincipalCloudIdentities
type ServicePrincipalCloudIdentities struct {

	// The list of Azure cloud identities assigned to the service principal.
	AzureCloudIdentities []*AzureCloudIdentity `json:"azureCloudIdentities"`

	// The name of the service principal that the cloud identities are assigned to.
	// Required: true
	ServicePrincipal *string `json:"servicePrincipal"`
}

// Validate validates this service principal cloud identities
func (m *ServicePrincipalCloudIdentities) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzureCloudIdentities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicePrincipal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServicePrincipalCloudIdentities) validateAzureCloudIdentities(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureCloudIdentities) { // not required
		return nil
	}

	for i := 0; i < len(m.AzureCloudIdentities); i++ {
		if swag.IsZero(m.AzureCloudIdentities[i]) { // not required
			continue
		}

		if m.AzureCloudIdentities[i] != nil {
			if err := m.AzureCloudIdentities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azureCloudIdentities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azureCloudIdentities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServicePrincipalCloudIdentities) validateServicePrincipal(formats strfmt.Registry) error {

	if err := validate.Required("servicePrincipal", "body", m.ServicePrincipal); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this service principal cloud identities based on the context it is used
func (m *ServicePrincipalCloudIdentities) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAzureCloudIdentities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServicePrincipalCloudIdentities) contextValidateAzureCloudIdentities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AzureCloudIdentities); i++ {

		if m.AzureCloudIdentities[i] != nil {

			if swag.IsZero(m.AzureCloudIdentities[i]) { // not required
				return nil
			}

			if err := m.AzureCloudIdentities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azureCloudIdentities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azureCloudIdentities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServicePrincipalCloudIdentities) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServicePrincipalCloudIdentities) UnmarshalBinary(b []byte) error {
	var res ServicePrincipalCloudIdentities
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
