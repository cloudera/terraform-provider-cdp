// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SyncStatus Status of a sync operation.
//
// swagger:model SyncStatus
type SyncStatus string

func NewSyncStatus(value SyncStatus) *SyncStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated SyncStatus.
func (m SyncStatus) Pointer() *SyncStatus {
	return &m
}

const (

	// SyncStatusNEVERRUN captures enum value "NEVER_RUN"
	SyncStatusNEVERRUN SyncStatus = "NEVER_RUN"

	// SyncStatusREQUESTED captures enum value "REQUESTED"
	SyncStatusREQUESTED SyncStatus = "REQUESTED"

	// SyncStatusREJECTED captures enum value "REJECTED"
	SyncStatusREJECTED SyncStatus = "REJECTED"

	// SyncStatusRUNNING captures enum value "RUNNING"
	SyncStatusRUNNING SyncStatus = "RUNNING"

	// SyncStatusCOMPLETED captures enum value "COMPLETED"
	SyncStatusCOMPLETED SyncStatus = "COMPLETED"

	// SyncStatusFAILED captures enum value "FAILED"
	SyncStatusFAILED SyncStatus = "FAILED"

	// SyncStatusTIMEDOUT captures enum value "TIMEDOUT"
	SyncStatusTIMEDOUT SyncStatus = "TIMEDOUT"
)

// for schema
var syncStatusEnum []interface{}

func init() {
	var res []SyncStatus
	if err := json.Unmarshal([]byte(`["NEVER_RUN","REQUESTED","REJECTED","RUNNING","COMPLETED","FAILED","TIMEDOUT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		syncStatusEnum = append(syncStatusEnum, v)
	}
}

func (m SyncStatus) validateSyncStatusEnum(path, location string, value SyncStatus) error {
	if err := validate.EnumCase(path, location, value, syncStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this sync status
func (m SyncStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSyncStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this sync status based on context it is used
func (m SyncStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
