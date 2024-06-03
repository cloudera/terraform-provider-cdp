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

// ServiceDescription Detailed description of a CDE service.
//
// swagger:model ServiceDescription
type ServiceDescription struct {

	// Base location for the service backup archives.
	BackupLocation string `json:"backupLocation,omitempty"`

	// Chart overrides for the Virtual Cluster.
	ChartValueOverrides []*ChartValueOverridesResponse `json:"chartValueOverrides"`

	// The cloud platform where the CDE service is enabled.
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// FQDN of the CDE service.
	ClusterFqdn string `json:"clusterFqdn,omitempty"`

	// Cluster Id of the CDE Service.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// CRN of the creator.
	CreatorCrn string `json:"creatorCrn,omitempty"`

	// Email address of the creator of the CDE service.
	CreatorEmail string `json:"creatorEmail,omitempty"`

	// Endpoint of Data Lake Atlas.
	DataLakeAtlasUIEndpoint string `json:"dataLakeAtlasUIEndpoint,omitempty"`

	// The Data lake file system.
	DataLakeFileSystems string `json:"dataLakeFileSystems,omitempty"`

	// Timestamp of service enabling.
	EnablingTime string `json:"enablingTime,omitempty"`

	// CRN of the environment.
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// CDP Environment Name.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// Comma-separated CIDRs that would be allowed to access the load balancer.
	LoadbalancerAllowlist string `json:"loadbalancerAllowlist,omitempty"`

	// Location for the log files of jobs.
	LogLocation string `json:"logLocation,omitempty"`

	// Name of the CDE Service.
	// Required: true
	Name *string `json:"name"`

	// Network outbound type. Currently 'udr' is the only supported.
	NetworkOutboundType string `json:"networkOutboundType,omitempty"`

	// The "true" value indicates that the previous version of the CDE service was requested to be deployed.
	PreviousVersionDeployed bool `json:"previousVersionDeployed,omitempty"`

	// If true, the CDE service was created with fully private Azure services (AKS, MySQL, etc.).
	PrivateClusterEnabled bool `json:"privateClusterEnabled,omitempty"`

	// If true, the CDE endpoint was created in a publicly accessible subnet.
	PublicEndpointEnabled bool `json:"publicEndpointEnabled,omitempty"`

	// Resources details of CDE Service.
	Resources *ServiceResources `json:"resources,omitempty"`

	// If true, SSD would have been be used for workload filesystem.
	SsdUsed bool `json:"ssdUsed,omitempty"`

	// Status of the CDE service.
	Status string `json:"status,omitempty"`

	// List of Subnet IDs of the CDP subnets used by the kubernetes worker node.
	Subnets string `json:"subnets,omitempty"`

	// CDP tenant ID.
	TenantID string `json:"tenantId,omitempty"`

	// List of CIDRs that would be allowed to access kubernetes master API server.
	WhitelistIps string `json:"whitelistIps,omitempty"`

	// If true, diagnostic information about job and query execution is sent to Cloudera Workload Manager.
	WorkloadAnalyticsEnabled bool `json:"workloadAnalyticsEnabled,omitempty"`
}

// Validate validates this service description
func (m *ServiceDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChartValueOverrides(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDescription) validateChartValueOverrides(formats strfmt.Registry) error {
	if swag.IsZero(m.ChartValueOverrides) { // not required
		return nil
	}

	for i := 0; i < len(m.ChartValueOverrides); i++ {
		if swag.IsZero(m.ChartValueOverrides[i]) { // not required
			continue
		}

		if m.ChartValueOverrides[i] != nil {
			if err := m.ChartValueOverrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chartValueOverrides" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chartValueOverrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDescription) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDescription) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDescription) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ServiceDescription) validateResources(formats strfmt.Registry) error {
	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this service description based on the context it is used
func (m *ServiceDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChartValueOverrides(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceDescription) contextValidateChartValueOverrides(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChartValueOverrides); i++ {

		if m.ChartValueOverrides[i] != nil {

			if swag.IsZero(m.ChartValueOverrides[i]) { // not required
				return nil
			}

			if err := m.ChartValueOverrides[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chartValueOverrides" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("chartValueOverrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ServiceDescription) contextValidateResources(ctx context.Context, formats strfmt.Registry) error {

	if m.Resources != nil {

		if swag.IsZero(m.Resources) { // not required
			return nil
		}

		if err := m.Resources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceDescription) UnmarshalBinary(b []byte) error {
	var res ServiceDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
