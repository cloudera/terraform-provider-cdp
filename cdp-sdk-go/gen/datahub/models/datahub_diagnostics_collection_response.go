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

// DatahubDiagnosticsCollectionResponse Response object for diagnostic collection flow details.
//
// swagger:model DatahubDiagnosticsCollectionResponse
type DatahubDiagnosticsCollectionResponse struct {

	// Additional details about the diagnostics collection.
	CollectionDetails *DatahubDiagnosticsCollectionDetailsResponse `json:"collectionDetails,omitempty"`

	// Creation date of the diagnostics collection flow.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Flow ID of the diagnostics collection flow.
	FlowID string `json:"flowId,omitempty"`

	// Current state of the diagnostics collection flow.
	FlowState string `json:"flowState,omitempty"`

	// Progress percentage of the diagnostics collection flow (maximum value if finished).
	ProgressPercentage int32 `json:"progressPercentage,omitempty"`

	// Status of the diagnostics collection flow.
	// Enum: ["RUNNING","FAILED","FINISHED","CANCELLED"]
	Status string `json:"status,omitempty"`
}

// Validate validates this datahub diagnostics collection response
func (m *DatahubDiagnosticsCollectionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollectionDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatahubDiagnosticsCollectionResponse) validateCollectionDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.CollectionDetails) { // not required
		return nil
	}

	if m.CollectionDetails != nil {
		if err := m.CollectionDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collectionDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collectionDetails")
			}
			return err
		}
	}

	return nil
}

func (m *DatahubDiagnosticsCollectionResponse) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

var datahubDiagnosticsCollectionResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RUNNING","FAILED","FINISHED","CANCELLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		datahubDiagnosticsCollectionResponseTypeStatusPropEnum = append(datahubDiagnosticsCollectionResponseTypeStatusPropEnum, v)
	}
}

const (

	// DatahubDiagnosticsCollectionResponseStatusRUNNING captures enum value "RUNNING"
	DatahubDiagnosticsCollectionResponseStatusRUNNING string = "RUNNING"

	// DatahubDiagnosticsCollectionResponseStatusFAILED captures enum value "FAILED"
	DatahubDiagnosticsCollectionResponseStatusFAILED string = "FAILED"

	// DatahubDiagnosticsCollectionResponseStatusFINISHED captures enum value "FINISHED"
	DatahubDiagnosticsCollectionResponseStatusFINISHED string = "FINISHED"

	// DatahubDiagnosticsCollectionResponseStatusCANCELLED captures enum value "CANCELLED"
	DatahubDiagnosticsCollectionResponseStatusCANCELLED string = "CANCELLED"
)

// prop value enum
func (m *DatahubDiagnosticsCollectionResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, datahubDiagnosticsCollectionResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DatahubDiagnosticsCollectionResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this datahub diagnostics collection response based on the context it is used
func (m *DatahubDiagnosticsCollectionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCollectionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatahubDiagnosticsCollectionResponse) contextValidateCollectionDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.CollectionDetails != nil {

		if swag.IsZero(m.CollectionDetails) { // not required
			return nil
		}

		if err := m.CollectionDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collectionDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collectionDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatahubDiagnosticsCollectionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatahubDiagnosticsCollectionResponse) UnmarshalBinary(b []byte) error {
	var res DatahubDiagnosticsCollectionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
