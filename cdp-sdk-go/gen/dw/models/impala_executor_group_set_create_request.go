// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ImpalaExecutorGroupSetCreateRequest Configure independently scaling set of uniformly sized executor groups.
//
// swagger:model ImpalaExecutorGroupSetCreateRequest
type ImpalaExecutorGroupSetCreateRequest struct {

	// Set auto suspend threshold. If not provided defaults will apply.
	AutoSuspendTimeoutSeconds int32 `json:"autoSuspendTimeoutSeconds,omitempty"`

	// Turn off auto suspend. If not provided defaults will apply.
	DisableAutoSuspend bool `json:"disableAutoSuspend,omitempty"`

	// Set number of executors per executor group.
	ExecGroupSize int32 `json:"execGroupSize"`

	// Set maximum number of executor groups allowed.
	MaxExecutorGroups int32 `json:"maxExecutorGroups"`

	// Set minimum number of executor groups allowed.
	MinExecutorGroups int32 `json:"minExecutorGroups"`

	// Set scale down threshold in seconds. If not provided defaults will apply.
	TriggerScaleDownDelay int32 `json:"triggerScaleDownDelay,omitempty"`

	// Set scale up threshold in seconds. If not provided defaults will apply.
	TriggerScaleUpDelay int32 `json:"triggerScaleUpDelay,omitempty"`
}

// Validate validates this impala executor group set create request
func (m *ImpalaExecutorGroupSetCreateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this impala executor group set create request based on context it is used
func (m *ImpalaExecutorGroupSetCreateRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImpalaExecutorGroupSetCreateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImpalaExecutorGroupSetCreateRequest) UnmarshalBinary(b []byte) error {
	var res ImpalaExecutorGroupSetCreateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
