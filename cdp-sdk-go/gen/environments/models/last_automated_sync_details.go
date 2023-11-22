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

// LastAutomatedSyncDetails The details of the last sync performed by automated user sync.
//
// swagger:model LastAutomatedSyncDetails
type LastAutomatedSyncDetails struct {

	// The status of the sync.
	// Required: true
	// Enum: [UNKNOWN SUCCESS FAILED]
	Status *string `json:"status"`

	// Additional detail related to the status.
	StatusMessages []string `json:"statusMessages"`

	// The time when the sync was processed.
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`
}

// Validate validates this last automated sync details
func (m *LastAutomatedSyncDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var lastAutomatedSyncDetailsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","SUCCESS","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lastAutomatedSyncDetailsTypeStatusPropEnum = append(lastAutomatedSyncDetailsTypeStatusPropEnum, v)
	}
}

const (

	// LastAutomatedSyncDetailsStatusUNKNOWN captures enum value "UNKNOWN"
	LastAutomatedSyncDetailsStatusUNKNOWN string = "UNKNOWN"

	// LastAutomatedSyncDetailsStatusSUCCESS captures enum value "SUCCESS"
	LastAutomatedSyncDetailsStatusSUCCESS string = "SUCCESS"

	// LastAutomatedSyncDetailsStatusFAILED captures enum value "FAILED"
	LastAutomatedSyncDetailsStatusFAILED string = "FAILED"
)

// prop value enum
func (m *LastAutomatedSyncDetails) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, lastAutomatedSyncDetailsTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LastAutomatedSyncDetails) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *LastAutomatedSyncDetails) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this last automated sync details based on context it is used
func (m *LastAutomatedSyncDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LastAutomatedSyncDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LastAutomatedSyncDetails) UnmarshalBinary(b []byte) error {
	var res LastAutomatedSyncDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}