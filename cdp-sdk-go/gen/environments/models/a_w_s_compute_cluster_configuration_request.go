// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AWSComputeClusterConfigurationRequest Request object for creating Externalized compute cluster for the environment.
//
// swagger:model AWSComputeClusterConfigurationRequest
type AWSComputeClusterConfigurationRequest struct {

	// Kubernetes API authorized IP ranges in CIDR notation. Mutually exclusive with privateCluster.
	KubeAPIAuthorizedIPRanges []string `json:"kubeApiAuthorizedIpRanges"`

	// If true, creates private cluster.
	PrivateCluster bool `json:"privateCluster,omitempty"`

	// Specify subnets for Kubernetes Worker Nodes
	WorkerNodeSubnets []string `json:"workerNodeSubnets"`
}

// Validate validates this a w s compute cluster configuration request
func (m *AWSComputeClusterConfigurationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this a w s compute cluster configuration request based on context it is used
func (m *AWSComputeClusterConfigurationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AWSComputeClusterConfigurationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AWSComputeClusterConfigurationRequest) UnmarshalBinary(b []byte) error {
	var res AWSComputeClusterConfigurationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
