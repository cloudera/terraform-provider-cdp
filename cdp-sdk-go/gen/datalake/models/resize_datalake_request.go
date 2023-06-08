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

// ResizeDatalakeRequest Datalake resize request.
//
// swagger:model ResizeDatalakeRequest
type ResizeDatalakeRequest struct {

	// The name or CRN of the datalake.
	// Required: true
	DatalakeName *string `json:"datalakeName"`

	// The target size for the datalake.
	// Required: true
	// Enum: [MEDIUM_DUTY_HA]
	TargetSize *string `json:"targetSize"`
}

// Validate validates this resize datalake request
func (m *ResizeDatalakeRequest) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *ResizeDatalakeRequest) validateDatalakeName(formats strfmt.Registry) error {

	if err := validate.Required("datalakeName", "body", m.DatalakeName); err != nil {
		return err
	}

	return nil
}

var resizeDatalakeRequestTypeTargetSizePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MEDIUM_DUTY_HA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resizeDatalakeRequestTypeTargetSizePropEnum = append(resizeDatalakeRequestTypeTargetSizePropEnum, v)
	}
}

const (

	// ResizeDatalakeRequestTargetSizeMEDIUMDUTYHA captures enum value "MEDIUM_DUTY_HA"
	ResizeDatalakeRequestTargetSizeMEDIUMDUTYHA string = "MEDIUM_DUTY_HA"
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

// ContextValidate validates this resize datalake request based on context it is used
func (m *ResizeDatalakeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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
