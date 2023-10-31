// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// AssignUserRoleReader is a Reader for the AssignUserRole structure.
type AssignUserRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignUserRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignUserRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssignUserRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssignUserRoleOK creates a AssignUserRoleOK with default headers values
func NewAssignUserRoleOK() *AssignUserRoleOK {
	return &AssignUserRoleOK{}
}

/*
AssignUserRoleOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type AssignUserRoleOK struct {
	Payload models.AssignUserRoleResponse
}

// IsSuccess returns true when this assign user role o k response has a 2xx status code
func (o *AssignUserRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign user role o k response has a 3xx status code
func (o *AssignUserRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign user role o k response has a 4xx status code
func (o *AssignUserRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign user role o k response has a 5xx status code
func (o *AssignUserRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign user role o k response a status code equal to that given
func (o *AssignUserRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the assign user role o k response
func (o *AssignUserRoleOK) Code() int {
	return 200
}

func (o *AssignUserRoleOK) Error() string {
	return fmt.Sprintf("[POST /iam/assignUserRole][%d] assignUserRoleOK  %+v", 200, o.Payload)
}

func (o *AssignUserRoleOK) String() string {
	return fmt.Sprintf("[POST /iam/assignUserRole][%d] assignUserRoleOK  %+v", 200, o.Payload)
}

func (o *AssignUserRoleOK) GetPayload() models.AssignUserRoleResponse {
	return o.Payload
}

func (o *AssignUserRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignUserRoleDefault creates a AssignUserRoleDefault with default headers values
func NewAssignUserRoleDefault(code int) *AssignUserRoleDefault {
	return &AssignUserRoleDefault{
		_statusCode: code,
	}
}

/*
AssignUserRoleDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type AssignUserRoleDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this assign user role default response has a 2xx status code
func (o *AssignUserRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this assign user role default response has a 3xx status code
func (o *AssignUserRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this assign user role default response has a 4xx status code
func (o *AssignUserRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this assign user role default response has a 5xx status code
func (o *AssignUserRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this assign user role default response a status code equal to that given
func (o *AssignUserRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the assign user role default response
func (o *AssignUserRoleDefault) Code() int {
	return o._statusCode
}

func (o *AssignUserRoleDefault) Error() string {
	return fmt.Sprintf("[POST /iam/assignUserRole][%d] assignUserRole default  %+v", o._statusCode, o.Payload)
}

func (o *AssignUserRoleDefault) String() string {
	return fmt.Sprintf("[POST /iam/assignUserRole][%d] assignUserRole default  %+v", o._statusCode, o.Payload)
}

func (o *AssignUserRoleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignUserRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}