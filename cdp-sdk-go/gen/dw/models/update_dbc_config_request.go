// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateDbcConfigRequest Request object for the updateDbcConfig method.
//
// swagger:model UpdateDbcConfigRequest
type UpdateDbcConfigRequest struct {

	// ID of the cluster.
	// Required: true
	ClusterID *string `json:"clusterId"`

	// Database Catalog configuration component to update.
	// Required: true
	// Enum: [DasEventProcessor DatabusProducer HueQueryProcessor Metastore]
	Component *string `json:"component"`

	// ID of the Database Catalog.
	// Required: true
	DbcID *string `json:"dbcId"`

	// Configuration files of the selected component to update.
	Set []*ConfigBlock `json:"set"`
}

// Validate validates this update dbc config request
func (m *UpdateDbcConfigRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComponent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDbcID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateDbcConfigRequest) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

var updateDbcConfigRequestTypeComponentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DasEventProcessor","DatabusProducer","HueQueryProcessor","Metastore"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDbcConfigRequestTypeComponentPropEnum = append(updateDbcConfigRequestTypeComponentPropEnum, v)
	}
}

const (

	// UpdateDbcConfigRequestComponentDasEventProcessor captures enum value "DasEventProcessor"
	UpdateDbcConfigRequestComponentDasEventProcessor string = "DasEventProcessor"

	// UpdateDbcConfigRequestComponentDatabusProducer captures enum value "DatabusProducer"
	UpdateDbcConfigRequestComponentDatabusProducer string = "DatabusProducer"

	// UpdateDbcConfigRequestComponentHueQueryProcessor captures enum value "HueQueryProcessor"
	UpdateDbcConfigRequestComponentHueQueryProcessor string = "HueQueryProcessor"

	// UpdateDbcConfigRequestComponentMetastore captures enum value "Metastore"
	UpdateDbcConfigRequestComponentMetastore string = "Metastore"
)

// prop value enum
func (m *UpdateDbcConfigRequest) validateComponentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateDbcConfigRequestTypeComponentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateDbcConfigRequest) validateComponent(formats strfmt.Registry) error {

	if err := validate.Required("component", "body", m.Component); err != nil {
		return err
	}

	// value enum
	if err := m.validateComponentEnum("component", "body", *m.Component); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDbcConfigRequest) validateDbcID(formats strfmt.Registry) error {

	if err := validate.Required("dbcId", "body", m.DbcID); err != nil {
		return err
	}

	return nil
}

func (m *UpdateDbcConfigRequest) validateSet(formats strfmt.Registry) error {
	if swag.IsZero(m.Set) { // not required
		return nil
	}

	for i := 0; i < len(m.Set); i++ {
		if swag.IsZero(m.Set[i]) { // not required
			continue
		}

		if m.Set[i] != nil {
			if err := m.Set[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update dbc config request based on the context it is used
func (m *UpdateDbcConfigRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateDbcConfigRequest) contextValidateSet(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Set); i++ {

		if m.Set[i] != nil {

			if swag.IsZero(m.Set[i]) { // not required
				return nil
			}

			if err := m.Set[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("set" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("set" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateDbcConfigRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateDbcConfigRequest) UnmarshalBinary(b []byte) error {
	var res UpdateDbcConfigRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}