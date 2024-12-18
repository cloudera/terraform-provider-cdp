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

// StorageType Storage type for clusters.
//
//	`CLOUD_WITH_EPHEMERAL` - Cloud with ephemeral storage. `CLOUD` - Cloud storage without ephemeral storage. `HDFS` - HDFS storage. `CLOUD_WITH_EPHEMERAL_DATATIERING` - Cloud with Ephemeral Storage and Datatiering.
//
// swagger:model StorageType
type StorageType string

func NewStorageType(value StorageType) *StorageType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated StorageType.
func (m StorageType) Pointer() *StorageType {
	return &m
}

const (

	// StorageTypeCLOUDWITHEPHEMERAL captures enum value "CLOUD_WITH_EPHEMERAL"
	StorageTypeCLOUDWITHEPHEMERAL StorageType = "CLOUD_WITH_EPHEMERAL"

	// StorageTypeCLOUD captures enum value "CLOUD"
	StorageTypeCLOUD StorageType = "CLOUD"

	// StorageTypeHDFS captures enum value "HDFS"
	StorageTypeHDFS StorageType = "HDFS"
)

// for schema
var storageTypeEnum []interface{}

func init() {
	var res []StorageType
	if err := json.Unmarshal([]byte(`["CLOUD_WITH_EPHEMERAL","CLOUD","HDFS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		storageTypeEnum = append(storageTypeEnum, v)
	}
}

func (m StorageType) validateStorageTypeEnum(path, location string, value StorageType) error {
	if err := validate.EnumCase(path, location, value, storageTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this storage type
func (m StorageType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStorageTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this storage type based on context it is used
func (m StorageType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
