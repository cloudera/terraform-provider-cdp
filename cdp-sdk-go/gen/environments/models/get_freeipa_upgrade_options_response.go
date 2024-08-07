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
)

// GetFreeipaUpgradeOptionsResponse The response object with available FreeIPA upgrade candidates.
//
// swagger:model GetFreeipaUpgradeOptionsResponse
type GetFreeipaUpgradeOptionsResponse struct {

	// The current image.
	CurrentImage *ImageInfoResponse `json:"currentImage,omitempty"`

	// The list of the upgrade candidates.
	Images []*ImageInfoResponse `json:"images"`
}

// Validate validates this get freeipa upgrade options response
func (m *GetFreeipaUpgradeOptionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetFreeipaUpgradeOptionsResponse) validateCurrentImage(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentImage) { // not required
		return nil
	}

	if m.CurrentImage != nil {
		if err := m.CurrentImage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currentImage")
			}
			return err
		}
	}

	return nil
}

func (m *GetFreeipaUpgradeOptionsResponse) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get freeipa upgrade options response based on the context it is used
func (m *GetFreeipaUpgradeOptionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCurrentImage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetFreeipaUpgradeOptionsResponse) contextValidateCurrentImage(ctx context.Context, formats strfmt.Registry) error {

	if m.CurrentImage != nil {

		if swag.IsZero(m.CurrentImage) { // not required
			return nil
		}

		if err := m.CurrentImage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentImage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currentImage")
			}
			return err
		}
	}

	return nil
}

func (m *GetFreeipaUpgradeOptionsResponse) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Images); i++ {

		if m.Images[i] != nil {

			if swag.IsZero(m.Images[i]) { // not required
				return nil
			}

			if err := m.Images[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetFreeipaUpgradeOptionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFreeipaUpgradeOptionsResponse) UnmarshalBinary(b []byte) error {
	var res GetFreeipaUpgradeOptionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
