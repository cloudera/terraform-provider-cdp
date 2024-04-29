// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateAzureClusterRequest Request object for the createAzureCluster method.
//
// swagger:model CreateAzureClusterRequest
type CreateAzureClusterRequest struct {

	// Enable AKS VNet Azure Virtual Network (VNet) integration by specifying the delegated subnet name. An Azure Kubernetes Service (AKS) cluster configured with API Server VNet Integration projects the API server endpoint directly into a delegated subnet in the VNet where AKS is deployed. API Server VNet Integration enables network communication between the API server and the cluster nodes without requiring a private link or tunnel.
	AksVNETIntegrationSubnetName string `json:"aksVNETIntegrationSubnetName,omitempty"`

	// Azure compute instance types that the environment is restricted to use. This affects the creation of virtual warehouses where this restriction will apply. Select an instance type that meets your computing, memory, networking, or storage needs. As of now, only a single instance type can be listed. Use describe-allowed-instance-types to see currently possible values and the default value used for the case it is not provided.
	ComputeInstanceTypes []string `json:"computeInstanceTypes"`

	// Options for custom ACR/ECR/Docker registries.
	CustomRegistryOptions *CustomRegistryOptions `json:"customRegistryOptions,omitempty"`

	// Custom environment subdomain. Overrides the environment subdomain using a customized domain either in the old subdomain format like ENV_ID.dw or the new format like dw-ENV_NAME.
	CustomSubdomain string `json:"customSubdomain,omitempty"`

	// PostgreSQL server backup retention days.
	DatabaseBackupRetentionPeriod *int32 `json:"databaseBackupRetentionPeriod,omitempty"`

	// Enables Azure Availability Zones for the cluster deployment.
	EnableAZ bool `json:"enableAZ,omitempty"`

	// Enable Azure Private AKS mode.
	EnablePrivateAks *bool `json:"enablePrivateAks,omitempty"`

	// Enables private SQL for the cluster deployment.
	EnablePrivateSQL *bool `json:"enablePrivateSQL,omitempty"`

	// Whether to enable spot instances for Virtual warehouses. It cannot be updated later. Defaults to false.
	EnableSpotInstances *bool `json:"enableSpotInstances,omitempty"`

	// The CRN of the environment for the cluster to create.
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// Enable monitoring of Azure Kubernetes Service (AKS) cluster. Workspace ID for Azure log analytics.
	LogAnalyticsWorkspaceID string `json:"logAnalyticsWorkspaceId,omitempty"`

	// Network outbound type. This setting controls the egress traffic for cluster nodes in Azure Kubernetes Service. Please refer to the following AKS documentation on the Azure portal. https://learn.microsoft.com/en-us/azure/aks/egress-outboundtype, https://learn.microsoft.com/en-us/azure/aks/nat-gateway
	// Enum: [LoadBalancer UserAssignedNATGateway UserDefinedRouting]
	OutboundType string `json:"outboundType,omitempty"`

	// Private DNS zone AKS resource ID.
	PrivateDNSZoneAKS string `json:"privateDNSZoneAKS,omitempty"`

	// Set additional number of nodes to reserve for executors and coordinators to use during autoscaling. Adding more reserved nodes increases your cloud costs.
	ReservedComputeNodes int32 `json:"reservedComputeNodes,omitempty"`

	// Set additional number of nodes to reserve for other services in the cluster. Adding more reserved nodes increases your cloud costs.
	ReservedSharedServicesNodes int32 `json:"reservedSharedServicesNodes,omitempty"`

	// Name of Azure subnet where the cluster should be deployed. It is a mandatory parameter for Azure cluster creation.
	// Required: true
	SubnetName *string `json:"subnetName"`

	// Set up load balancer with private IP address. An internal load balancer gets created. Make sure there is connectivity between your client network and the network VNet where CDW environment is deployed.
	UseInternalLoadBalancer bool `json:"useInternalLoadBalancer,omitempty"`

	// With overlay network nodes get an IP address from the Azure virtual network subnet. Pods receive an IP address from a logically different address space to the Azure virtual network subnet of the nodes.
	UseOverlayNetworking bool `json:"useOverlayNetworking,omitempty"`

	// Resource ID of the managed identity used by AKS. It is a mandatory parameter for Azure cluster creation.
	// Required: true
	UserAssignedManagedIdentity *string `json:"userAssignedManagedIdentity"`

	// List of IP address CIDRs to whitelist for kubernetes cluster access.
	WhitelistK8sClusterAccessIPCIDRs []string `json:"whitelistK8sClusterAccessIpCIDRs"`

	// List of IP address CIDRs to whitelist for workload access.
	WhitelistWorkloadAccessIPCIDRs []string `json:"whitelistWorkloadAccessIpCIDRs"`
}

// Validate validates this create azure cluster request
func (m *CreateAzureClusterRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomRegistryOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutboundType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserAssignedManagedIdentity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureClusterRequest) validateCustomRegistryOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomRegistryOptions) { // not required
		return nil
	}

	if m.CustomRegistryOptions != nil {
		if err := m.CustomRegistryOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customRegistryOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customRegistryOptions")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAzureClusterRequest) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

var createAzureClusterRequestTypeOutboundTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LoadBalancer","UserAssignedNATGateway","UserDefinedRouting"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createAzureClusterRequestTypeOutboundTypePropEnum = append(createAzureClusterRequestTypeOutboundTypePropEnum, v)
	}
}

const (

	// CreateAzureClusterRequestOutboundTypeLoadBalancer captures enum value "LoadBalancer"
	CreateAzureClusterRequestOutboundTypeLoadBalancer string = "LoadBalancer"

	// CreateAzureClusterRequestOutboundTypeUserAssignedNATGateway captures enum value "UserAssignedNATGateway"
	CreateAzureClusterRequestOutboundTypeUserAssignedNATGateway string = "UserAssignedNATGateway"

	// CreateAzureClusterRequestOutboundTypeUserDefinedRouting captures enum value "UserDefinedRouting"
	CreateAzureClusterRequestOutboundTypeUserDefinedRouting string = "UserDefinedRouting"
)

// prop value enum
func (m *CreateAzureClusterRequest) validateOutboundTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createAzureClusterRequestTypeOutboundTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateAzureClusterRequest) validateOutboundType(formats strfmt.Registry) error {
	if swag.IsZero(m.OutboundType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOutboundTypeEnum("outboundType", "body", m.OutboundType); err != nil {
		return err
	}

	return nil
}

func (m *CreateAzureClusterRequest) validateSubnetName(formats strfmt.Registry) error {

	if err := validate.Required("subnetName", "body", m.SubnetName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAzureClusterRequest) validateUserAssignedManagedIdentity(formats strfmt.Registry) error {

	if err := validate.Required("userAssignedManagedIdentity", "body", m.UserAssignedManagedIdentity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create azure cluster request based on the context it is used
func (m *CreateAzureClusterRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomRegistryOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAzureClusterRequest) contextValidateCustomRegistryOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomRegistryOptions != nil {

		if swag.IsZero(m.CustomRegistryOptions) { // not required
			return nil
		}

		if err := m.CustomRegistryOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customRegistryOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customRegistryOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAzureClusterRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAzureClusterRequest) UnmarshalBinary(b []byte) error {
	var res CreateAzureClusterRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
