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

// AssignMachineUserResourceRoleReader is a Reader for the AssignMachineUserResourceRole structure.
type AssignMachineUserResourceRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignMachineUserResourceRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignMachineUserResourceRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssignMachineUserResourceRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssignMachineUserResourceRoleOK creates a AssignMachineUserResourceRoleOK with default headers values
func NewAssignMachineUserResourceRoleOK() *AssignMachineUserResourceRoleOK {
	return &AssignMachineUserResourceRoleOK{}
}

/*
AssignMachineUserResourceRoleOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type AssignMachineUserResourceRoleOK struct {
	Payload models.AssignMachineUserResourceRoleResponse
}

// IsSuccess returns true when this assign machine user resource role o k response has a 2xx status code
func (o *AssignMachineUserResourceRoleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign machine user resource role o k response has a 3xx status code
func (o *AssignMachineUserResourceRoleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign machine user resource role o k response has a 4xx status code
func (o *AssignMachineUserResourceRoleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign machine user resource role o k response has a 5xx status code
func (o *AssignMachineUserResourceRoleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign machine user resource role o k response a status code equal to that given
func (o *AssignMachineUserResourceRoleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the assign machine user resource role o k response
func (o *AssignMachineUserResourceRoleOK) Code() int {
	return 200
}

func (o *AssignMachineUserResourceRoleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/assignMachineUserResourceRole][%d] assignMachineUserResourceRoleOK %s", 200, payload)
}

func (o *AssignMachineUserResourceRoleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/assignMachineUserResourceRole][%d] assignMachineUserResourceRoleOK %s", 200, payload)
}

func (o *AssignMachineUserResourceRoleOK) GetPayload() models.AssignMachineUserResourceRoleResponse {
	return o.Payload
}

func (o *AssignMachineUserResourceRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignMachineUserResourceRoleDefault creates a AssignMachineUserResourceRoleDefault with default headers values
func NewAssignMachineUserResourceRoleDefault(code int) *AssignMachineUserResourceRoleDefault {
	return &AssignMachineUserResourceRoleDefault{
		_statusCode: code,
	}
}

/*
AssignMachineUserResourceRoleDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type AssignMachineUserResourceRoleDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this assign machine user resource role default response has a 2xx status code
func (o *AssignMachineUserResourceRoleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this assign machine user resource role default response has a 3xx status code
func (o *AssignMachineUserResourceRoleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this assign machine user resource role default response has a 4xx status code
func (o *AssignMachineUserResourceRoleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this assign machine user resource role default response has a 5xx status code
func (o *AssignMachineUserResourceRoleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this assign machine user resource role default response a status code equal to that given
func (o *AssignMachineUserResourceRoleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the assign machine user resource role default response
func (o *AssignMachineUserResourceRoleDefault) Code() int {
	return o._statusCode
}

func (o *AssignMachineUserResourceRoleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/assignMachineUserResourceRole][%d] assignMachineUserResourceRole default %s", o._statusCode, payload)
}

func (o *AssignMachineUserResourceRoleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/assignMachineUserResourceRole][%d] assignMachineUserResourceRole default %s", o._statusCode, payload)
}

func (o *AssignMachineUserResourceRoleDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignMachineUserResourceRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
