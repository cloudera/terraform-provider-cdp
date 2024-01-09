// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImpalaExecutorGroupSetUpdateRequest Re-configure independently scaling set of uniformly sized executor groups.
//
// swagger:model ImpalaExecutorGroupSetUpdateRequest
type ImpalaExecutorGroupSetUpdateRequest struct {

	// Set auto suspend threshold. If not provided defaults will apply.
	AutoSuspendTimeoutSeconds int32 `json:"autoSuspendTimeoutSeconds,omitempty"`

	// Delete the executor group set.
	DeleteGroupSet bool `json:"deleteGroupSet,omitempty"`

	// Turn off auto suspend. If not provided defaults will apply.
	DisableAutoSuspend bool `json:"disableAutoSuspend,omitempty"`

	// Set number of executors per executor group.
	ExecGroupSize int32 `json:"execGroupSize,omitempty"`

	// Set maximum number of executor groups allowed.
	MaxExecutorGroups int32 `json:"maxExecutorGroups,omitempty"`

	// Set minimum number of executor groups allowed.
	MinExecutorGroups int32 `json:"minExecutorGroups,omitempty"`

	// Set scale down threshold in seconds. If not provided defaults will apply.
	TriggerScaleDownDelay int32 `json:"triggerScaleDownDelay,omitempty"`

	// Set scale up threshold in seconds. If not provided defaults will apply.
	TriggerScaleUpDelay int32 `json:"triggerScaleUpDelay,omitempty"`
}

// Validate validates this impala executor group set update request
func (m *ImpalaExecutorGroupSetUpdateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this impala executor group set update request based on context it is used
func (m *ImpalaExecutorGroupSetUpdateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImpalaExecutorGroupSetUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImpalaExecutorGroupSetUpdateRequest) UnmarshalBinary(b []byte) error {
	var res ImpalaExecutorGroupSetUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
