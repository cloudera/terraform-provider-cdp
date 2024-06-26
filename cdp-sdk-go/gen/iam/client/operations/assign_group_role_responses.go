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

// AssignGroupRoleReader is a Reader for the AssignGroupRole structure.
type AssignGroupRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignGroupRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignGroupRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssignGroupRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssignGroupRoleOK creates a AssignGroupRoleOK with default headers values
func NewAssignGroupRoleOK() *AssignGroupRoleOK {
	return &AssignGroupRoleOK{}
}

/*
AssignGroupRoleOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type AssignGroupRoleOK struct {
	Payload models.AssignGroupRoleResponse
}

// IsSuccess returns true when this assign group role o k response has a 2xx status code
func (o *AssignGroupRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign group role o k response has a 3xx status code
func (o *AssignGroupRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign group role o k response has a 4xx status code
func (o *AssignGroupRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign group role o k response has a 5xx status code
func (o *AssignGroupRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign group role o k response a status code equal to that given
func (o *AssignGroupRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the assign group role o k response
func (o *AssignGroupRoleOK) Code() int {
	return 200
}

func (o *AssignGroupRoleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/assignGroupRole][%d] assignGroupRoleOK %s", 200, payload)
}

func (o *AssignGroupRoleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/assignGroupRole][%d] assignGroupRoleOK %s", 200, payload)
}

func (o *AssignGroupRoleOK) GetPayload() models.AssignGroupRoleResponse {
	return o.Payload
}

func (o *AssignGroupRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignGroupRoleDefault creates a AssignGroupRoleDefault with default headers values
func NewAssignGroupRoleDefault(code int) *AssignGroupRoleDefault {
	return &AssignGroupRoleDefault{
		_statusCode: code,
	}
}

/*
AssignGroupRoleDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type AssignGroupRoleDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this assign group role default response has a 2xx status code
func (o *AssignGroupRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this assign group role default response has a 3xx status code
func (o *AssignGroupRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this assign group role default response has a 4xx status code
func (o *AssignGroupRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this assign group role default response has a 5xx status code
func (o *AssignGroupRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this assign group role default response a status code equal to that given
func (o *AssignGroupRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the assign group role default response
func (o *AssignGroupRoleDefault) Code() int {
	return o._statusCode
}

func (o *AssignGroupRoleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/assignGroupRole][%d] assignGroupRole default %s", o._statusCode, payload)
}

func (o *AssignGroupRoleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/assignGroupRole][%d] assignGroupRole default %s", o._statusCode, payload)
}

func (o *AssignGroupRoleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignGroupRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
