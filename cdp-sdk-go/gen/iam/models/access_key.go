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

// AccessKey Information about a Cloudera CDP access key.
//
// swagger:model AccessKey
type AccessKey struct {

	// The ID of the access key.
	// Required: true
	AccessKeyID *string `json:"accessKeyId"`

	// The CRN of the actor with which this access key is associated.
	// Required: true
	ActorCrn *string `json:"actorCrn"`

	// The date when the access key was created.
	// Required: true
	// Format: date-time
	CreationDate *strfmt.DateTime `json:"creationDate"`

	// The CRN of the access key.
	// Required: true
	Crn *string `json:"crn"`

	// Information on the last time this access key was used.
	LastUsage *AccessKeyLastUsage `json:"lastUsage,omitempty"`

	// The status of an access key.
	// Enum: [ACTIVE INACTIVE]
	Status string `json:"status,omitempty"`

	// The type of an access key.
	Type AccessKeyType `json:"type,omitempty"`
}

// Validate validates this access key
func (m *AccessKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActorCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUsage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessKey) validateAccessKeyID(formats strfmt.Registry) error {

	if err := validate.Required("accessKeyId", "body", m.AccessKeyID); err != nil {
		return err
	}

	return nil
}

func (m *AccessKey) validateActorCrn(formats strfmt.Registry) error {

	if err := validate.Required("actorCrn", "body", m.ActorCrn); err != nil {
		return err
	}

	return nil
}

func (m *AccessKey) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccessKey) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *AccessKey) validateLastUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUsage) { // not required
		return nil
	}

	if m.LastUsage != nil {
		if err := m.LastUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastUsage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastUsage")
			}
			return err
		}
	}

	return nil
}

var accessKeyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accessKeyTypeStatusPropEnum = append(accessKeyTypeStatusPropEnum, v)
	}
}

const (

	// AccessKeyStatusACTIVE captures enum value "ACTIVE"
	AccessKeyStatusACTIVE string = "ACTIVE"

	// AccessKeyStatusINACTIVE captures enum value "INACTIVE"
	AccessKeyStatusINACTIVE string = "INACTIVE"
)

// prop value enum
func (m *AccessKey) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accessKeyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccessKey) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *AccessKey) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// ContextValidate validate this access key based on the context it is used
func (m *AccessKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessKey) contextValidateLastUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.LastUsage != nil {

		if swag.IsZero(m.LastUsage) { // not required
			return nil
		}

		if err := m.LastUsage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastUsage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lastUsage")
			}
			return err
		}
	}

	return nil
}

func (m *AccessKey) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccessKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccessKey) UnmarshalBinary(b []byte) error {
	var res AccessKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
