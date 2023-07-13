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

// ClusterSummaryResponse A Cloudera Data Warehouse cluster.
//
// swagger:model ClusterSummaryResponse
type ClusterSummaryResponse struct {

	// Additional (fallback) instance types listed in their priority order. They are used instead of the primary compute instance type in case it is unavailable. Since additional instance types are not supported for Azure, this is always empty for it.
	AdditionalInstanceTypes []string `json:"additionalInstanceTypes"`

	// Response object of AWS related cluster options.
	AwsOptions *AwsOptionsResponse `json:"awsOptions,omitempty"`

	// Response object of Azure related cluster options.
	AzureOptions *AzureOptionsResponse `json:"azureOptions,omitempty"`

	// The cloud platform of the environment that was used to create this cluster.
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// Compute instance types that the environment is restricted to use. This affects the creation of the virtual warehouses where this restriction will apply.
	ComputeInstanceTypes []string `json:"computeInstanceTypes"`

	// Creation date of cluster.
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// The creator of the cluster.
	Creator *ActorResponse `json:"creator,omitempty"`

	// The CRN of the cluster.
	Crn string `json:"crn,omitempty"`

	// Denotes whether the spot instances have been enabled for the cluster. This value is only available for AWS and Azure clusters.
	EnableSpotInstances bool `json:"enableSpotInstances,omitempty"`

	// Enable Storage Roles checkbox was checked when creating/activating this cluster.
	EnableStorageRoles bool `json:"enableStorageRoles,omitempty"`

	// The CRN of the environment where the cluster is located.
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// The ID of the cluster.
	ID string `json:"id,omitempty"`

	// Name of the cluster (same as the name of the environment).
	Name string `json:"name,omitempty"`

	// Status of the cluster. Possible values are: Creating, Created, Accepted, Starting, Running, Stopping, Stopped, Updating, PreUpdate, Upgrading, PreUpgrade, Restarting, Deleting, Waiting, Failed, Error.
	Status string `json:"status,omitempty"`
}

// Validate validates this cluster summary response
func (m *ClusterSummaryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSummaryResponse) validateAwsOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsOptions) { // not required
		return nil
	}

	if m.AwsOptions != nil {
		if err := m.AwsOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSummaryResponse) validateAzureOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureOptions) { // not required
		return nil
	}

	if m.AzureOptions != nil {
		if err := m.AzureOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSummaryResponse) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClusterSummaryResponse) validateCreator(formats strfmt.Registry) error {
	if swag.IsZero(m.Creator) { // not required
		return nil
	}

	if m.Creator != nil {
		if err := m.Creator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creator")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster summary response based on the context it is used
func (m *ClusterSummaryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSummaryResponse) contextValidateAwsOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsOptions != nil {

		if swag.IsZero(m.AwsOptions) { // not required
			return nil
		}

		if err := m.AwsOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSummaryResponse) contextValidateAzureOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureOptions != nil {

		if swag.IsZero(m.AzureOptions) { // not required
			return nil
		}

		if err := m.AzureOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureOptions")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSummaryResponse) contextValidateCreator(ctx context.Context, formats strfmt.Registry) error {

	if m.Creator != nil {

		if swag.IsZero(m.Creator) { // not required
			return nil
		}

		if err := m.Creator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creator")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creator")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSummaryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSummaryResponse) UnmarshalBinary(b []byte) error {
	var res ClusterSummaryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
