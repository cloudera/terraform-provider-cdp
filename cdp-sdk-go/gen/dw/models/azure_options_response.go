// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AzureOptionsResponse Response object of the cluster Azure settings.
//
// swagger:model AzureOptionsResponse
type AzureOptionsResponse struct {

	// Pod CIDR setting for Azure CNI networking.
	AksPodCIDR string `json:"aksPodCIDR,omitempty"`

	// AKS VNet integration subnet name. If it's an empty string, then VNet integration is disabled.
	AksVNetIntegrationSubnetName string `json:"aksVNetIntegrationSubnetName"`

	// Denotes whther the Azure Availability Zones for the cluster is enabled or not.
	EnableAZ *bool `json:"enableAZ,omitempty"`

	// Denotes whether the AKS cluster is in private mode.
	EnablePrivateAKS *bool `json:"enablePrivateAKS,omitempty"`

	// Denotes whether the private SQL is enabled for the cluster.
	EnablePrivateSQL *bool `json:"enablePrivateSQL,omitempty"`

	// Workspace ID for Azure log analytics.
	LogAnalyticsWorkspaceID string `json:"logAnalyticsWorkspaceId,omitempty"`

	// The current outbound type setting.
	OutboundType string `json:"outboundType,omitempty"`

	// The resource ID of the private DNS zone for the AKS cluster.
	PrivateDNSZoneAKS string `json:"privateDNSZoneAKS,omitempty"`

	// Private DNS zone ID for the PostgreSQL server.
	PrivateDNSZoneSQL string `json:"privateDNSZoneSQL,omitempty"`

	// Name of the delegated subnet where the private SQL should be deployed.
	PrivateSQLSubnetName string `json:"privateSQLSubnetName,omitempty"`

	// ID of Azure subnet where the cluster is deployed.
	SubnetID string `json:"subnetId,omitempty"`

	// The resource ID of the managed identity used by the AKS cluster.
	UserAssignedManagedIdentity string `json:"userAssignedManagedIdentity,omitempty"`
}

// Validate validates this azure options response
func (m *AzureOptionsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this azure options response based on context it is used
func (m *AzureOptionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureOptionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureOptionsResponse) UnmarshalBinary(b []byte) error {
	var res AzureOptionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
