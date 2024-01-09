// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// UpdateSSHKeyReader is a Reader for the UpdateSSHKey structure.
type UpdateSSHKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSSHKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSSHKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateSSHKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSSHKeyOK creates a UpdateSSHKeyOK with default headers values
func NewUpdateSSHKeyOK() *UpdateSSHKeyOK {
	return &UpdateSSHKeyOK{}
}

/*
UpdateSSHKeyOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateSSHKeyOK struct {
	Payload *models.UpdateSSHKeyResponse
}

// IsSuccess returns true when this update Ssh key o k response has a 2xx status code
func (o *UpdateSSHKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update Ssh key o k response has a 3xx status code
func (o *UpdateSSHKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update Ssh key o k response has a 4xx status code
func (o *UpdateSSHKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update Ssh key o k response has a 5xx status code
func (o *UpdateSSHKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update Ssh key o k response a status code equal to that given
func (o *UpdateSSHKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update Ssh key o k response
func (o *UpdateSSHKeyOK) Code() int {
	return 200
}

func (o *UpdateSSHKeyOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/updateSshKey][%d] updateSshKeyOK  %+v", 200, o.Payload)
}

func (o *UpdateSSHKeyOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/updateSshKey][%d] updateSshKeyOK  %+v", 200, o.Payload)
}

func (o *UpdateSSHKeyOK) GetPayload() *models.UpdateSSHKeyResponse {
	return o.Payload
}

func (o *UpdateSSHKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateSSHKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSSHKeyDefault creates a UpdateSSHKeyDefault with default headers values
func NewUpdateSSHKeyDefault(code int) *UpdateSSHKeyDefault {
	return &UpdateSSHKeyDefault{
		_statusCode: code,
	}
}

/*
UpdateSSHKeyDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateSSHKeyDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update Ssh key default response has a 2xx status code
func (o *UpdateSSHKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update Ssh key default response has a 3xx status code
func (o *UpdateSSHKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update Ssh key default response has a 4xx status code
func (o *UpdateSSHKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update Ssh key default response has a 5xx status code
func (o *UpdateSSHKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update Ssh key default response a status code equal to that given
func (o *UpdateSSHKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update Ssh key default response
func (o *UpdateSSHKeyDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSSHKeyDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/updateSshKey][%d] updateSshKey default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSSHKeyDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/updateSshKey][%d] updateSshKey default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSSHKeyDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSSHKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
