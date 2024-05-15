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

// RecoverDatalakeRequest Datalake recover request.
//
// swagger:model RecoverDatalakeRequest
type RecoverDatalakeRequest struct {

	// The name or CRN of the datalake.
	// Required: true
	DatalakeName *string `json:"datalakeName"`

	// The type of the recovery. The default value is RECOVER_WITHOUT_DATA. The recovery always runs with RECOVER_WITH_DATA if the on resize failure.
	// Enum: ["RECOVER_WITH_DATA","RECOVER_WITHOUT_DATA"]
	RecoveryType string `json:"recoveryType,omitempty"`
}

// Validate validates this recover datalake request
func (m *RecoverDatalakeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatalakeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoveryType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverDatalakeRequest) validateDatalakeName(formats strfmt.Registry) error {

	if err := validate.Required("datalakeName", "body", m.DatalakeName); err != nil {
		return err
	}

	return nil
}

var recoverDatalakeRequestTypeRecoveryTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RECOVER_WITH_DATA","RECOVER_WITHOUT_DATA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recoverDatalakeRequestTypeRecoveryTypePropEnum = append(recoverDatalakeRequestTypeRecoveryTypePropEnum, v)
	}
}

const (

	// RecoverDatalakeRequestRecoveryTypeRECOVERWITHDATA captures enum value "RECOVER_WITH_DATA"
	RecoverDatalakeRequestRecoveryTypeRECOVERWITHDATA string = "RECOVER_WITH_DATA"

	// RecoverDatalakeRequestRecoveryTypeRECOVERWITHOUTDATA captures enum value "RECOVER_WITHOUT_DATA"
	RecoverDatalakeRequestRecoveryTypeRECOVERWITHOUTDATA string = "RECOVER_WITHOUT_DATA"
)

// prop value enum
func (m *RecoverDatalakeRequest) validateRecoveryTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recoverDatalakeRequestTypeRecoveryTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecoverDatalakeRequest) validateRecoveryType(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoveryType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecoveryTypeEnum("recoveryType", "body", m.RecoveryType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this recover datalake request based on context it is used
func (m *RecoverDatalakeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecoverDatalakeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverDatalakeRequest) UnmarshalBinary(b []byte) error {
	var res RecoverDatalakeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
