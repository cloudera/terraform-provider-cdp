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

// DescribeUpgradeDatabaseResponse Response with upgrade availability of CDP Runtime and Operating System for a database.
//
// swagger:model DescribeUpgradeDatabaseResponse
type DescribeUpgradeDatabaseResponse struct {

	// List of available versions for upgrade.
	AvailableComponentVersions []*ComponentsVersion `json:"availableComponentVersions"`

	// Versions of currently deployed CDP runtime and operating system.
	CurrentComponentVersion *ComponentsVersion `json:"currentComponentVersion,omitempty"`

	// Is an OS upgrade available.
	IsOSUpgradeAvailable bool `json:"isOSUpgradeAvailable,omitempty"`

	// Is a CDP Runtime upgrade available.
	IsRuntimeUpgradeAvailable bool `json:"isRuntimeUpgradeAvailable,omitempty"`

	// The reason whether upgrade request is accepted or why it is not possible.
	StatusReason string `json:"statusReason,omitempty"`
}

// Validate validates this describe upgrade database response
func (m *DescribeUpgradeDatabaseResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableComponentVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentComponentVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeUpgradeDatabaseResponse) validateAvailableComponentVersions(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailableComponentVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailableComponentVersions); i++ {
		if swag.IsZero(m.AvailableComponentVersions[i]) { // not required
			continue
		}

		if m.AvailableComponentVersions[i] != nil {
			if err := m.AvailableComponentVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availableComponentVersions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availableComponentVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DescribeUpgradeDatabaseResponse) validateCurrentComponentVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentComponentVersion) { // not required
		return nil
	}

	if m.CurrentComponentVersion != nil {
		if err := m.CurrentComponentVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentComponentVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currentComponentVersion")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe upgrade database response based on the context it is used
func (m *DescribeUpgradeDatabaseResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailableComponentVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCurrentComponentVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeUpgradeDatabaseResponse) contextValidateAvailableComponentVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AvailableComponentVersions); i++ {

		if m.AvailableComponentVersions[i] != nil {

			if swag.IsZero(m.AvailableComponentVersions[i]) { // not required
				return nil
			}

			if err := m.AvailableComponentVersions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("availableComponentVersions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("availableComponentVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DescribeUpgradeDatabaseResponse) contextValidateCurrentComponentVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.CurrentComponentVersion != nil {

		if swag.IsZero(m.CurrentComponentVersion) { // not required
			return nil
		}

		if err := m.CurrentComponentVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentComponentVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("currentComponentVersion")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeUpgradeDatabaseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeUpgradeDatabaseResponse) UnmarshalBinary(b []byte) error {
	var res DescribeUpgradeDatabaseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
