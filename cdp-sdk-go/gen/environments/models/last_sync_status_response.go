// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LastSyncStatusResponse Response object for Sync Operation.
//
// swagger:model LastSyncStatusResponse
type LastSyncStatusResponse struct {

	// Date when the sync operation ended. Omitted if operation has not ended.
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// If there is any error associated. The error will be populated on any error and it may be populated when the operation failure details are empty.
	Error string `json:"error,omitempty"`

	// List of sync operation details for all failed environments.
	Failure []*SyncOperationDetails `json:"failure"`

	// Unique operation ID assigned to this command execution. Use this identifier with 'get-operation' to track status and retrieve detailed results.
	// Required: true
	OperationID *string `json:"operationId"`

	// Operation type, set password or user sync
	OperationType OperationType `json:"operationType,omitempty"`

	// Date when the sync operation started.
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// Status of this operation. Status can be one of these values (REQUESTED, RUNNING, COMPLETED, FAILED, REJECTED, TIMEDOUT)
	Status SyncStatus `json:"status,omitempty"`

	// List of sync operation details for all succeeded environments.
	Success []*SyncOperationDetails `json:"success"`
}

// Validate validates this last sync status response
func (m *LastSyncStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFailure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LastSyncStatusResponse) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LastSyncStatusResponse) validateFailure(formats strfmt.Registry) error {
	if swag.IsZero(m.Failure) { // not required
		return nil
	}

	for i := 0; i < len(m.Failure); i++ {
		if swag.IsZero(m.Failure[i]) { // not required
			continue
		}

		if m.Failure[i] != nil {
			if err := m.Failure[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failure" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("failure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LastSyncStatusResponse) validateOperationID(formats strfmt.Registry) error {

	if err := validate.Required("operationId", "body", m.OperationID); err != nil {
		return err
	}

	return nil
}

func (m *LastSyncStatusResponse) validateOperationType(formats strfmt.Registry) error {
	if swag.IsZero(m.OperationType) { // not required
		return nil
	}

	if err := m.OperationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operationType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operationType")
		}
		return err
	}

	return nil
}

func (m *LastSyncStatusResponse) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LastSyncStatusResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *LastSyncStatusResponse) validateSuccess(formats strfmt.Registry) error {
	if swag.IsZero(m.Success) { // not required
		return nil
	}

	for i := 0; i < len(m.Success); i++ {
		if swag.IsZero(m.Success[i]) { // not required
			continue
		}

		if m.Success[i] != nil {
			if err := m.Success[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("success" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("success" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this last sync status response based on the context it is used
func (m *LastSyncStatusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFailure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOperationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSuccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LastSyncStatusResponse) contextValidateFailure(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Failure); i++ {

		if m.Failure[i] != nil {

			if swag.IsZero(m.Failure[i]) { // not required
				return nil
			}

			if err := m.Failure[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failure" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("failure" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LastSyncStatusResponse) contextValidateOperationType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.OperationType) { // not required
		return nil
	}

	if err := m.OperationType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("operationType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("operationType")
		}
		return err
	}

	return nil
}

func (m *LastSyncStatusResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *LastSyncStatusResponse) contextValidateSuccess(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Success); i++ {

		if m.Success[i] != nil {

			if swag.IsZero(m.Success[i]) { // not required
				return nil
			}

			if err := m.Success[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("success" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("success" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LastSyncStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LastSyncStatusResponse) UnmarshalBinary(b []byte) error {
	var res LastSyncStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
