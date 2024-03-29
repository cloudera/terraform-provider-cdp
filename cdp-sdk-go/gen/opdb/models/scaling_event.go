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

// ScalingEvent Details of a scaling event.
//
// swagger:model ScalingEvent
type ScalingEvent struct {

	// The time of the event creation.
	// Format: date-time
	EventTimestamp strfmt.DateTime `json:"eventTimestamp,omitempty"`

	// JSON of the metric name and value that triggered the scaling event.
	Metric string `json:"metric,omitempty"`

	// Node count after scaling event.
	NodeCountAfter int32 `json:"nodeCountAfter,omitempty"`

	// Node count before scaling event.
	NodeCountBefore int32 `json:"nodeCountBefore,omitempty"`

	// Reason for the scaling event in a readable format.
	Reason string `json:"reason,omitempty"`

	// Scaling event status.
	Status string `json:"status,omitempty"`
}

// Validate validates this scaling event
func (m *ScalingEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScalingEvent) validateEventTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.EventTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("eventTimestamp", "body", "date-time", m.EventTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this scaling event based on context it is used
func (m *ScalingEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScalingEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScalingEvent) UnmarshalBinary(b []byte) error {
	var res ScalingEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
