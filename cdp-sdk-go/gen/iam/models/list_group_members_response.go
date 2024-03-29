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

// ListGroupMembersResponse Response object for a list group members request.
//
// swagger:model ListGroupMembersResponse
type ListGroupMembersResponse struct {

	// The list of group members.
	// Required: true
	MemberCrns []string `json:"memberCrns"`

	// The token to use when requesting the next set of results. If not present, there are no additional results.
	NextToken string `json:"nextToken,omitempty"`
}

// Validate validates this list group members response
func (m *ListGroupMembersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMemberCrns(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListGroupMembersResponse) validateMemberCrns(formats strfmt.Registry) error {

	if err := validate.Required("memberCrns", "body", m.MemberCrns); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list group members response based on context it is used
func (m *ListGroupMembersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListGroupMembersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListGroupMembersResponse) UnmarshalBinary(b []byte) error {
	var res ListGroupMembersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
