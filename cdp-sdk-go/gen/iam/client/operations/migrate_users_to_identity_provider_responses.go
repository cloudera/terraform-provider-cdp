// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// MigrateUsersToIdentityProviderReader is a Reader for the MigrateUsersToIdentityProvider structure.
type MigrateUsersToIdentityProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MigrateUsersToIdentityProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMigrateUsersToIdentityProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMigrateUsersToIdentityProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMigrateUsersToIdentityProviderOK creates a MigrateUsersToIdentityProviderOK with default headers values
func NewMigrateUsersToIdentityProviderOK() *MigrateUsersToIdentityProviderOK {
	return &MigrateUsersToIdentityProviderOK{}
}

/*
MigrateUsersToIdentityProviderOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type MigrateUsersToIdentityProviderOK struct {
	Payload *models.MigrateUsersToIdentityProviderResponse
}

// IsSuccess returns true when this migrate users to identity provider o k response has a 2xx status code
func (o *MigrateUsersToIdentityProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this migrate users to identity provider o k response has a 3xx status code
func (o *MigrateUsersToIdentityProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this migrate users to identity provider o k response has a 4xx status code
func (o *MigrateUsersToIdentityProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this migrate users to identity provider o k response has a 5xx status code
func (o *MigrateUsersToIdentityProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this migrate users to identity provider o k response a status code equal to that given
func (o *MigrateUsersToIdentityProviderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the migrate users to identity provider o k response
func (o *MigrateUsersToIdentityProviderOK) Code() int {
	return 200
}

func (o *MigrateUsersToIdentityProviderOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/migrateUsersToIdentityProvider][%d] migrateUsersToIdentityProviderOK %s", 200, payload)
}

func (o *MigrateUsersToIdentityProviderOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/migrateUsersToIdentityProvider][%d] migrateUsersToIdentityProviderOK %s", 200, payload)
}

func (o *MigrateUsersToIdentityProviderOK) GetPayload() *models.MigrateUsersToIdentityProviderResponse {
	return o.Payload
}

func (o *MigrateUsersToIdentityProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MigrateUsersToIdentityProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateUsersToIdentityProviderDefault creates a MigrateUsersToIdentityProviderDefault with default headers values
func NewMigrateUsersToIdentityProviderDefault(code int) *MigrateUsersToIdentityProviderDefault {
	return &MigrateUsersToIdentityProviderDefault{
		_statusCode: code,
	}
}

/*
MigrateUsersToIdentityProviderDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type MigrateUsersToIdentityProviderDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this migrate users to identity provider default response has a 2xx status code
func (o *MigrateUsersToIdentityProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this migrate users to identity provider default response has a 3xx status code
func (o *MigrateUsersToIdentityProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this migrate users to identity provider default response has a 4xx status code
func (o *MigrateUsersToIdentityProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this migrate users to identity provider default response has a 5xx status code
func (o *MigrateUsersToIdentityProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this migrate users to identity provider default response a status code equal to that given
func (o *MigrateUsersToIdentityProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the migrate users to identity provider default response
func (o *MigrateUsersToIdentityProviderDefault) Code() int {
	return o._statusCode
}

func (o *MigrateUsersToIdentityProviderDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/migrateUsersToIdentityProvider][%d] migrateUsersToIdentityProvider default %s", o._statusCode, payload)
}

func (o *MigrateUsersToIdentityProviderDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/migrateUsersToIdentityProvider][%d] migrateUsersToIdentityProvider default %s", o._statusCode, payload)
}

func (o *MigrateUsersToIdentityProviderDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *MigrateUsersToIdentityProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
