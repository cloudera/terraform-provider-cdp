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

// SyncUserResponse Response Object for single user sync operation.
//
// swagger:model SyncUserResponse
type SyncUserResponse struct {

	// Sync operation end timestamp.
	EndTime string `json:"endTime,omitempty"`

	// If there is any error associated.
	Error string `json:"error,omitempty"`

	// List of sync operation details for all failed envs.
	Failure []*SyncOperationDetails `json:"failure"`

	// UUID of the request for this operation. This Id can be used for geting status on the operation.
	// Required: true
	OperationID *string `json:"operationId"`

	// Operation type, set password or user sync
	OperationType OperationType `json:"operationType,omitempty"`

	// Sync operation start timestamp.
	StartTime string `json:"startTime,omitempty"`

	// Status of this operation. Status can be one of these values (REQUESTED, RUNNING, COMPLETED, FAILED, REJECTED, TIMEDOUT)
	Status SyncStatus `json:"status,omitempty"`

	// List of sync operation details for all succeeded environments.
	Success []*SyncOperationDetails `json:"success"`
}

// Validate validates this sync user response
func (m *SyncUserResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperationType(formats); err != nil {
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

func (m *SyncUserResponse) validateFailure(formats strfmt.Registry) error {
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

func (m *SyncUserResponse) validateOperationID(formats strfmt.Registry) error {

	if err := validate.Required("operationId", "body", m.OperationID); err != nil {
		return err
	}

	return nil
}

func (m *SyncUserResponse) validateOperationType(formats strfmt.Registry) error {
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

func (m *SyncUserResponse) validateStatus(formats strfmt.Registry) error {
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

func (m *SyncUserResponse) validateSuccess(formats strfmt.Registry) error {
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

// ContextValidate validate this sync user response based on the context it is used
func (m *SyncUserResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *SyncUserResponse) contextValidateFailure(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SyncUserResponse) contextValidateOperationType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SyncUserResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SyncUserResponse) contextValidateSuccess(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SyncUserResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SyncUserResponse) UnmarshalBinary(b []byte) error {
	var res SyncUserResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
