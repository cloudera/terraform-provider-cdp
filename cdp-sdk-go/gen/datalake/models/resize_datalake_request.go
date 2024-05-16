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

// ResizeDatalakeRequest Datalake resize request.
//
// swagger:model ResizeDatalakeRequest
type ResizeDatalakeRequest struct {

	// Any custom database properties to override defaults.
	CustomDatabaseComputeStorage *CustomDatabaseComputeStorage `json:"customDatabaseComputeStorage,omitempty"`

	// Any custom instance disk size to override defaults.
	CustomInstanceDisks []*CustomInstanceDisk `json:"customInstanceDisks"`

	// Any custom instance type to override defaults.
	CustomInstanceTypes []*CustomInstanceType `json:"customInstanceTypes"`

	// The name or CRN of the datalake.
	// Required: true
	DatalakeName *string `json:"datalakeName"`

	// Whether to deploy a new datalake in a multi-availability zone way.
	MultiAz *bool `json:"multiAz,omitempty"`

	// The target size for the datalake. The resize target size can be MEDIUM_DUTY or ENTERPRISE. If the runtime version >= 7.2.17 target size is ENTERPRISE. If not, the target size is MEDIUM_DUTY.
	// Required: true
	// Enum: ["MEDIUM_DUTY_HA","ENTERPRISE"]
	TargetSize *string `json:"targetSize"`
}

// Validate validates this resize datalake request
func (m *ResizeDatalakeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomDatabaseComputeStorage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomInstanceDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomInstanceTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalakeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResizeDatalakeRequest) validateCustomDatabaseComputeStorage(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomDatabaseComputeStorage) { // not required
		return nil
	}

	if m.CustomDatabaseComputeStorage != nil {
		if err := m.CustomDatabaseComputeStorage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customDatabaseComputeStorage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customDatabaseComputeStorage")
			}
			return err
		}
	}

	return nil
}

func (m *ResizeDatalakeRequest) validateCustomInstanceDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomInstanceDisks) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomInstanceDisks); i++ {
		if swag.IsZero(m.CustomInstanceDisks[i]) { // not required
			continue
		}

		if m.CustomInstanceDisks[i] != nil {
			if err := m.CustomInstanceDisks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customInstanceDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customInstanceDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResizeDatalakeRequest) validateCustomInstanceTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomInstanceTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomInstanceTypes); i++ {
		if swag.IsZero(m.CustomInstanceTypes[i]) { // not required
			continue
		}

		if m.CustomInstanceTypes[i] != nil {
			if err := m.CustomInstanceTypes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customInstanceTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customInstanceTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResizeDatalakeRequest) validateDatalakeName(formats strfmt.Registry) error {

	if err := validate.Required("datalakeName", "body", m.DatalakeName); err != nil {
		return err
	}

	return nil
}

var resizeDatalakeRequestTypeTargetSizePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MEDIUM_DUTY_HA","ENTERPRISE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resizeDatalakeRequestTypeTargetSizePropEnum = append(resizeDatalakeRequestTypeTargetSizePropEnum, v)
	}
}

const (

	// ResizeDatalakeRequestTargetSizeMEDIUMDUTYHA captures enum value "MEDIUM_DUTY_HA"
	ResizeDatalakeRequestTargetSizeMEDIUMDUTYHA string = "MEDIUM_DUTY_HA"

	// ResizeDatalakeRequestTargetSizeENTERPRISE captures enum value "ENTERPRISE"
	ResizeDatalakeRequestTargetSizeENTERPRISE string = "ENTERPRISE"
)

// prop value enum
func (m *ResizeDatalakeRequest) validateTargetSizeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, resizeDatalakeRequestTypeTargetSizePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ResizeDatalakeRequest) validateTargetSize(formats strfmt.Registry) error {

	if err := validate.Required("targetSize", "body", m.TargetSize); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetSizeEnum("targetSize", "body", *m.TargetSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this resize datalake request based on the context it is used
func (m *ResizeDatalakeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomDatabaseComputeStorage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomInstanceDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomInstanceTypes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResizeDatalakeRequest) contextValidateCustomDatabaseComputeStorage(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomDatabaseComputeStorage != nil {

		if swag.IsZero(m.CustomDatabaseComputeStorage) { // not required
			return nil
		}

		if err := m.CustomDatabaseComputeStorage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customDatabaseComputeStorage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customDatabaseComputeStorage")
			}
			return err
		}
	}

	return nil
}

func (m *ResizeDatalakeRequest) contextValidateCustomInstanceDisks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomInstanceDisks); i++ {

		if m.CustomInstanceDisks[i] != nil {

			if swag.IsZero(m.CustomInstanceDisks[i]) { // not required
				return nil
			}

			if err := m.CustomInstanceDisks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customInstanceDisks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customInstanceDisks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResizeDatalakeRequest) contextValidateCustomInstanceTypes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomInstanceTypes); i++ {

		if m.CustomInstanceTypes[i] != nil {

			if swag.IsZero(m.CustomInstanceTypes[i]) { // not required
				return nil
			}

			if err := m.CustomInstanceTypes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customInstanceTypes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("customInstanceTypes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResizeDatalakeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResizeDatalakeRequest) UnmarshalBinary(b []byte) error {
	var res ResizeDatalakeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
