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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// SetAzureAuditCredentialReader is a Reader for the SetAzureAuditCredential structure.
type SetAzureAuditCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAzureAuditCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetAzureAuditCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetAzureAuditCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetAzureAuditCredentialOK creates a SetAzureAuditCredentialOK with default headers values
func NewSetAzureAuditCredentialOK() *SetAzureAuditCredentialOK {
	return &SetAzureAuditCredentialOK{}
}

/*
SetAzureAuditCredentialOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SetAzureAuditCredentialOK struct {
	Payload *models.SetAzureAuditCredentialResponse
}

// IsSuccess returns true when this set azure audit credential o k response has a 2xx status code
func (o *SetAzureAuditCredentialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set azure audit credential o k response has a 3xx status code
func (o *SetAzureAuditCredentialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set azure audit credential o k response has a 4xx status code
func (o *SetAzureAuditCredentialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set azure audit credential o k response has a 5xx status code
func (o *SetAzureAuditCredentialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set azure audit credential o k response a status code equal to that given
func (o *SetAzureAuditCredentialOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set azure audit credential o k response
func (o *SetAzureAuditCredentialOK) Code() int {
	return 200
}

func (o *SetAzureAuditCredentialOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/setAzureAuditCredential][%d] setAzureAuditCredentialOK %s", 200, payload)
}

func (o *SetAzureAuditCredentialOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/setAzureAuditCredential][%d] setAzureAuditCredentialOK %s", 200, payload)
}

func (o *SetAzureAuditCredentialOK) GetPayload() *models.SetAzureAuditCredentialResponse {
	return o.Payload
}

func (o *SetAzureAuditCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SetAzureAuditCredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAzureAuditCredentialDefault creates a SetAzureAuditCredentialDefault with default headers values
func NewSetAzureAuditCredentialDefault(code int) *SetAzureAuditCredentialDefault {
	return &SetAzureAuditCredentialDefault{
		_statusCode: code,
	}
}

/*
SetAzureAuditCredentialDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SetAzureAuditCredentialDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this set azure audit credential default response has a 2xx status code
func (o *SetAzureAuditCredentialDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this set azure audit credential default response has a 3xx status code
func (o *SetAzureAuditCredentialDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this set azure audit credential default response has a 4xx status code
func (o *SetAzureAuditCredentialDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this set azure audit credential default response has a 5xx status code
func (o *SetAzureAuditCredentialDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this set azure audit credential default response a status code equal to that given
func (o *SetAzureAuditCredentialDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the set azure audit credential default response
func (o *SetAzureAuditCredentialDefault) Code() int {
	return o._statusCode
}

func (o *SetAzureAuditCredentialDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/setAzureAuditCredential][%d] setAzureAuditCredential default %s", o._statusCode, payload)
}

func (o *SetAzureAuditCredentialDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/setAzureAuditCredential][%d] setAzureAuditCredential default %s", o._statusCode, payload)
}

func (o *SetAzureAuditCredentialDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAzureAuditCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
