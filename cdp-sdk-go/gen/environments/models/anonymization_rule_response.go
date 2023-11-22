// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AnonymizationRuleResponse Anonymization rule response object rule that is applied on logs that are sent to Cloudera.
//
// swagger:model AnonymizationRuleResponse
type AnonymizationRuleResponse struct {

	// If rule pattern (value) matches, that will be replaced for this string (default [REDACTED])
	Replacement string `json:"replacement,omitempty"`

	// Pattern of the rule that should be redacted.
	Value string `json:"value,omitempty"`
}

// Validate validates this anonymization rule response
func (m *AnonymizationRuleResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this anonymization rule response based on context it is used
func (m *AnonymizationRuleResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnonymizationRuleResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnonymizationRuleResponse) UnmarshalBinary(b []byte) error {
	var res AnonymizationRuleResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}