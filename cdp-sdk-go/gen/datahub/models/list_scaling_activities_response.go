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

// ListScalingActivitiesResponse Response object for list scaling activities request.
//
// swagger:model ListScalingActivitiesResponse
type ListScalingActivitiesResponse struct {

	// The token to use when requesting the next set of results. If not present, there are no additional results.
	NextToken string `json:"nextToken,omitempty"`

	// The list of scaling activities.
	// Required: true
	ScalingActivity []*ScalingActivitySummary `json:"scalingActivity"`
}

// Validate validates this list scaling activities response
func (m *ListScalingActivitiesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScalingActivity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListScalingActivitiesResponse) validateScalingActivity(formats strfmt.Registry) error {

	if err := validate.Required("scalingActivity", "body", m.ScalingActivity); err != nil {
		return err
	}

	for i := 0; i < len(m.ScalingActivity); i++ {
		if swag.IsZero(m.ScalingActivity[i]) { // not required
			continue
		}

		if m.ScalingActivity[i] != nil {
			if err := m.ScalingActivity[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scalingActivity" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scalingActivity" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list scaling activities response based on the context it is used
func (m *ListScalingActivitiesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScalingActivity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListScalingActivitiesResponse) contextValidateScalingActivity(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ScalingActivity); i++ {

		if m.ScalingActivity[i] != nil {

			if swag.IsZero(m.ScalingActivity[i]) { // not required
				return nil
			}

			if err := m.ScalingActivity[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scalingActivity" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("scalingActivity" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListScalingActivitiesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListScalingActivitiesResponse) UnmarshalBinary(b []byte) error {
	var res ListScalingActivitiesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
