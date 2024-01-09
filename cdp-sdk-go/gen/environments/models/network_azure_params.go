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

// NetworkAzureParams Azure network parameters.
//
// swagger:model NetworkAzureParams
type NetworkAzureParams struct {

	// The full Azure resource ID of an existing Private DNS zone used for the AKS.
	AksPrivateDNSZoneID string `json:"aksPrivateDnsZoneId,omitempty"`

	// The full Azure resource ID of the existing Private DNS Zone used for Flexible Server and Single Server Databases.
	DatabasePrivateDNSZoneID string `json:"databasePrivateDnsZoneId,omitempty"`

	// Whether the outbound load balancer was created for this environment.
	EnableOutboundLoadBalancer bool `json:"enableOutboundLoadBalancer,omitempty"`

	// The subnets delegated for Flexible Server database. Accepts either the name or the full resource id.
	FlexibleServerSubnetIds []string `json:"flexibleServerSubnetIds"`

	// The id of the Azure VNet.
	// Required: true
	NetworkID *string `json:"networkId"`

	// The name of the resource group associated with the VNet.
	// Required: true
	ResourceGroupName *string `json:"resourceGroupName"`

	// Whether to associate public ip's to the resources within the network.
	// Required: true
	UsePublicIP *bool `json:"usePublicIp"`
}

// Validate validates this network azure params
func (m *NetworkAzureParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroupName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsePublicIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkAzureParams) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("networkId", "body", m.NetworkID); err != nil {
		return err
	}

	return nil
}

func (m *NetworkAzureParams) validateResourceGroupName(formats strfmt.Registry) error {

	if err := validate.Required("resourceGroupName", "body", m.ResourceGroupName); err != nil {
		return err
	}

	return nil
}

func (m *NetworkAzureParams) validateUsePublicIP(formats strfmt.Registry) error {

	if err := validate.Required("usePublicIp", "body", m.UsePublicIP); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network azure params based on context it is used
func (m *NetworkAzureParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkAzureParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkAzureParams) UnmarshalBinary(b []byte) error {
	var res NetworkAzureParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
