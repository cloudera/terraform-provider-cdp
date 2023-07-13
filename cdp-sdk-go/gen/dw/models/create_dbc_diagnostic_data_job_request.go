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

// CreateDbcDiagnosticDataJobRequest Request object for the createDbcDiagnosticDataJob method.
//
// swagger:model CreateDbcDiagnosticDataJobRequest
type CreateDbcDiagnosticDataJobRequest struct {

	// Additional user-defined metadata information which is attached to resulting bundle-info.json when posting the bundle.
	BundleMetadata map[string]string `json:"bundleMetadata,omitempty"`

	// Optional support case number in case of UPLOAD_TO_CLOUDERA destination, otherwise only act as additional data.
	CaseNumber string `json:"caseNumber,omitempty"`

	// ID of cluster.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// ID of the Database Catalog.
	// Required: true
	DbcID *string `json:"dbcId"`

	// Destination of the diagnostics collection.
	// Required: true
	// Enum: [UPLOAD_TO_CLOUDERA DOWNLOAD]
	Destination *string `json:"destination"`

	// Database Catalog diagnostic options. If not provided, everything will be included in the Diagnostic Data.
	DownloadOptions *DBCCreateDiagnosticDataDownloadOptions `json:"downloadOptions,omitempty"`

	// The resulting bundle will contain logs/metrics before the specified end time. If not indicated, then the current time is taken as the end time.
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// Forced recreation of the diagnostic job.
	Force *bool `json:"force,omitempty"`

	// The resulting bundle will contain logs/metrics after the specified start time. If not indicated, then 30 minutes ago from now is taken as the start time.
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`
}

// Validate validates this create dbc diagnostic data job request
func (m *CreateDbcDiagnosticDataJobRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDbcID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownloadOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDbcDiagnosticDataJobRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *CreateDbcDiagnosticDataJobRequest) validateDbcID(formats strfmt.Registry) error {

	if err := validate.Required("dbcId", "body", m.DbcID); err != nil {
		return err
	}

	return nil
}

var createDbcDiagnosticDataJobRequestTypeDestinationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UPLOAD_TO_CLOUDERA","DOWNLOAD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createDbcDiagnosticDataJobRequestTypeDestinationPropEnum = append(createDbcDiagnosticDataJobRequestTypeDestinationPropEnum, v)
	}
}

const (

	// CreateDbcDiagnosticDataJobRequestDestinationUPLOADTOCLOUDERA captures enum value "UPLOAD_TO_CLOUDERA"
	CreateDbcDiagnosticDataJobRequestDestinationUPLOADTOCLOUDERA string = "UPLOAD_TO_CLOUDERA"

	// CreateDbcDiagnosticDataJobRequestDestinationDOWNLOAD captures enum value "DOWNLOAD"
	CreateDbcDiagnosticDataJobRequestDestinationDOWNLOAD string = "DOWNLOAD"
)

// prop value enum
func (m *CreateDbcDiagnosticDataJobRequest) validateDestinationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createDbcDiagnosticDataJobRequestTypeDestinationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateDbcDiagnosticDataJobRequest) validateDestination(formats strfmt.Registry) error {

	if err := validate.Required("destination", "body", m.Destination); err != nil {
		return err
	}

	// value enum
	if err := m.validateDestinationEnum("destination", "body", *m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *CreateDbcDiagnosticDataJobRequest) validateDownloadOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.DownloadOptions) { // not required
		return nil
	}

	if m.DownloadOptions != nil {
		if err := m.DownloadOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downloadOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downloadOptions")
			}
			return err
		}
	}

	return nil
}

func (m *CreateDbcDiagnosticDataJobRequest) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateDbcDiagnosticDataJobRequest) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create dbc diagnostic data job request based on the context it is used
func (m *CreateDbcDiagnosticDataJobRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDownloadOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateDbcDiagnosticDataJobRequest) contextValidateDownloadOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.DownloadOptions != nil {

		if swag.IsZero(m.DownloadOptions) { // not required
			return nil
		}

		if err := m.DownloadOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("downloadOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("downloadOptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateDbcDiagnosticDataJobRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateDbcDiagnosticDataJobRequest) UnmarshalBinary(b []byte) error {
	var res CreateDbcDiagnosticDataJobRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
