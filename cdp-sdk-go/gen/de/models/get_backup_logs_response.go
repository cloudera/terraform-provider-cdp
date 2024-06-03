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

// GetBackupLogsResponse Response object for the Get Backup Logs command.
//
// swagger:model GetBackupLogsResponse
type GetBackupLogsResponse struct {

	// Contains the logs of the backup operation.
	// Required: true
	Logs []string `json:"logs"`

	// The token to use when requesting the next set of results. If there are no additional results, the string is empty.
	NextToken string `json:"nextToken,omitempty"`
}

// Validate validates this get backup logs response
func (m *GetBackupLogsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBackupLogsResponse) validateLogs(formats strfmt.Registry) error {

	if err := validate.Required("logs", "body", m.Logs); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get backup logs response based on context it is used
func (m *GetBackupLogsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetBackupLogsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBackupLogsResponse) UnmarshalBinary(b []byte) error {
	var res GetBackupLogsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
