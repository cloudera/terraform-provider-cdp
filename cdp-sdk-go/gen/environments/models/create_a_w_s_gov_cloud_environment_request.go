// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateAWSGovCloudEnvironmentRequest Request object for a create AWS GovCloud environment request.
//
// swagger:model CreateAWSGovCloudEnvironmentRequest
type CreateAWSGovCloudEnvironmentRequest struct {

	// SSH authentication information for accessing cluster node instances. Users with access to this authentication information have root level access to the Data Lake and Data Hub cluster instances.
	// Required: true
	Authentication *AuthenticationRequest `json:"authentication"`

	// Whether to create private subnets or not.
	CreatePrivateSubnets bool `json:"createPrivateSubnets,omitempty"`

	// Whether to create service endpoints or not.
	CreateServiceEndpoints bool `json:"createServiceEndpoints,omitempty"`

	// Name of the credential to use for the environment.
	// Required: true
	CredentialName *string `json:"credentialName"`

	// Configures the desired custom docker registry for data services.
	CustomDockerRegistry *CustomDockerRegistryRequest `json:"customDockerRegistry,omitempty"`

	// An description of the environment.
	Description string `json:"description,omitempty"`

	// Whether to enable SSH tunneling for the environment.
	EnableTunnel *bool `json:"enableTunnel,omitempty"`

	// ARN of the AWS KMS CMK to use for the server-side encryption of AWS storage resources.
	EncryptionKeyArn string `json:"encryptionKeyArn,omitempty"`

	// The scheme for the endpoint gateway. PUBLIC creates an external endpoint that can be accessed over the Internet. Defaults to PRIVATE which restricts the traffic to be internal to the VPC.
	// Enum: [PUBLIC PRIVATE]
	EndpointAccessGatewayScheme string `json:"endpointAccessGatewayScheme,omitempty"`

	// The subnets to use for endpoint access gateway.
	EndpointAccessGatewaySubnetIds []string `json:"endpointAccessGatewaySubnetIds"`

	// The name of the environment. Must contain only lowercase letters, numbers and hyphens.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// The FreeIPA creation request for the environment
	FreeIpa *AWSFreeIpaCreationRequest `json:"freeIpa,omitempty"`

	// The FreeIPA image request for the environment
	Image *FreeIpaImageRequest `json:"image,omitempty"`

	// AWS storage configuration for cluster and audit logs.
	// Required: true
	LogStorage *AwsLogStorageRequest `json:"logStorage"`

	// The network CIDR. This will create a VPC along with subnets in multiple Availability Zones.
	NetworkCidr string `json:"networkCidr,omitempty"`

	// Name of the proxy config to use for the environment.
	ProxyConfigName string `json:"proxyConfigName,omitempty"`

	// The region of the environment.
	// Required: true
	Region *string `json:"region"`

	// When true, this will report additional diagnostic information back to Cloudera.
	ReportDeploymentLogs bool `json:"reportDeploymentLogs,omitempty"`

	// The name for the DynamoDB table backing S3Guard.
	S3GuardTableName string `json:"s3GuardTableName,omitempty"`

	// Security control for FreeIPA and Data Lake deployment.
	// Required: true
	SecurityAccess *SecurityAccessRequest `json:"securityAccess"`

	// One or more subnet IDs within the VPC. Mutually exclusive with networkCidr.
	// Unique: true
	SubnetIds []string `json:"subnetIds"`

	// Tags associated with the resources.
	Tags []*TagRequest `json:"tags"`

	// The Amazon VPC ID. Mutually exclusive with networkCidr.
	VpcID string `json:"vpcId,omitempty"`

	// When this is enabled, diagnostic information about job and query execution is sent to Workload Manager for Data Hub clusters created within this environment.
	WorkloadAnalytics bool `json:"workloadAnalytics,omitempty"`
}

// Validate validates this create a w s gov cloud environment request
func (m *CreateAWSGovCloudEnvironmentRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCredentialName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomDockerRegistry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpointAccessGatewayScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeIpa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) validateAuthentication(formats strfmt.Registry) error {

	if err := validate.Required("authentication", "body", m.Authentication); err != nil {
		return err
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) validateCredentialName(formats strfmt.Registry) error {

	if err := validate.Required("credentialName", "body", m.CredentialName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) validateCustomDockerRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomDockerRegistry) { // not required
		return nil
	}

	if m.CustomDockerRegistry != nil {
		if err := m.CustomDockerRegistry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customDockerRegistry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customDockerRegistry")
			}
			return err
		}
	}

	return nil
}

var createAWSGovCloudEnvironmentRequestTypeEndpointAccessGatewaySchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PUBLIC","PRIVATE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createAWSGovCloudEnvironmentRequestTypeEndpointAccessGatewaySchemePropEnum = append(createAWSGovCloudEnvironmentRequestTypeEndpointAccessGatewaySchemePropEnum, v)
	}
}

const (

	// CreateAWSGovCloudEnvironmentRequestEndpointAccessGatewaySchemePUBLIC captures enum value "PUBLIC"
	CreateAWSGovCloudEnvironmentRequestEndpointAccessGatewaySchemePUBLIC string = "PUBLIC"

	// CreateAWSGovCloudEnvironmentRequestEndpointAccessGatewaySchemePRIVATE captures enum value "PRIVATE"
	CreateAWSGovCloudEnvironmentRequestEndpointAccessGatewaySchemePRIVATE string = "PRIVATE"
)

// prop value enum
func (m *CreateAWSGovCloudEnvironmentRequest) validateEndpointAccessGatewaySchemeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createAWSGovCloudEnvironmentRequestTypeEndpointAccessGatewaySchemePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) validateEndpointAccessGatewayScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.EndpointAccessGatewayScheme) { // not required
		return nil
	}

	// value enum
	if err := m.validateEndpointAccessGatewaySchemeEnum("endpointAccessGatewayScheme", "body", m.EndpointAccessGatewayScheme); err != nil {
		return err
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) validateFreeIpa(formats strfmt.Registry) error {
	if swag.IsZero(m.FreeIpa) { // not required
		return nil
	}

	if m.FreeIpa != nil {
		if err := m.FreeIpa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeIpa")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("freeIpa")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) validateImage(formats strfmt.Registry) error {
	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) validateLogStorage(formats strfmt.Registry) error {

	if err := validate.Required("logStorage", "body", m.LogStorage); err != nil {
		return err
	}

	if m.LogStorage != nil {
		if err := m.LogStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logStorage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logStorage")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) validateSecurityAccess(formats strfmt.Registry) error {

	if err := validate.Required("securityAccess", "body", m.SecurityAccess); err != nil {
		return err
	}

	if m.SecurityAccess != nil {
		if err := m.SecurityAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityAccess")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityAccess")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) validateSubnetIds(formats strfmt.Registry) error {
	if swag.IsZero(m.SubnetIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("subnetIds", "body", m.SubnetIds); err != nil {
		return err
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create a w s gov cloud environment request based on the context it is used
func (m *CreateAWSGovCloudEnvironmentRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomDockerRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFreeIpa(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecurityAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) contextValidateAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if m.Authentication != nil {

		if err := m.Authentication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) contextValidateCustomDockerRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomDockerRegistry != nil {

		if swag.IsZero(m.CustomDockerRegistry) { // not required
			return nil
		}

		if err := m.CustomDockerRegistry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customDockerRegistry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customDockerRegistry")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) contextValidateFreeIpa(ctx context.Context, formats strfmt.Registry) error {

	if m.FreeIpa != nil {

		if swag.IsZero(m.FreeIpa) { // not required
			return nil
		}

		if err := m.FreeIpa.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeIpa")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("freeIpa")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) contextValidateImage(ctx context.Context, formats strfmt.Registry) error {

	if m.Image != nil {

		if swag.IsZero(m.Image) { // not required
			return nil
		}

		if err := m.Image.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) contextValidateLogStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.LogStorage != nil {

		if err := m.LogStorage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logStorage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logStorage")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) contextValidateSecurityAccess(ctx context.Context, formats strfmt.Registry) error {

	if m.SecurityAccess != nil {

		if err := m.SecurityAccess.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityAccess")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securityAccess")
			}
			return err
		}
	}

	return nil
}

func (m *CreateAWSGovCloudEnvironmentRequest) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAWSGovCloudEnvironmentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAWSGovCloudEnvironmentRequest) UnmarshalBinary(b []byte) error {
	var res CreateAWSGovCloudEnvironmentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
