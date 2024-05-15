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

// ListMachineUserAssignedResourceRolesReader is a Reader for the ListMachineUserAssignedResourceRoles structure.
type ListMachineUserAssignedResourceRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListMachineUserAssignedResourceRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListMachineUserAssignedResourceRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListMachineUserAssignedResourceRolesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListMachineUserAssignedResourceRolesOK creates a ListMachineUserAssignedResourceRolesOK with default headers values
func NewListMachineUserAssignedResourceRolesOK() *ListMachineUserAssignedResourceRolesOK {
	return &ListMachineUserAssignedResourceRolesOK{}
}

/*
ListMachineUserAssignedResourceRolesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListMachineUserAssignedResourceRolesOK struct {
	Payload *models.ListMachineUserAssignedResourceRolesResponse
}

// IsSuccess returns true when this list machine user assigned resource roles o k response has a 2xx status code
func (o *ListMachineUserAssignedResourceRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list machine user assigned resource roles o k response has a 3xx status code
func (o *ListMachineUserAssignedResourceRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list machine user assigned resource roles o k response has a 4xx status code
func (o *ListMachineUserAssignedResourceRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list machine user assigned resource roles o k response has a 5xx status code
func (o *ListMachineUserAssignedResourceRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list machine user assigned resource roles o k response a status code equal to that given
func (o *ListMachineUserAssignedResourceRolesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list machine user assigned resource roles o k response
func (o *ListMachineUserAssignedResourceRolesOK) Code() int {
	return 200
}

func (o *ListMachineUserAssignedResourceRolesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/listMachineUserAssignedResourceRoles][%d] listMachineUserAssignedResourceRolesOK %s", 200, payload)
}

func (o *ListMachineUserAssignedResourceRolesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/listMachineUserAssignedResourceRoles][%d] listMachineUserAssignedResourceRolesOK %s", 200, payload)
}

func (o *ListMachineUserAssignedResourceRolesOK) GetPayload() *models.ListMachineUserAssignedResourceRolesResponse {
	return o.Payload
}

func (o *ListMachineUserAssignedResourceRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListMachineUserAssignedResourceRolesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListMachineUserAssignedResourceRolesDefault creates a ListMachineUserAssignedResourceRolesDefault with default headers values
func NewListMachineUserAssignedResourceRolesDefault(code int) *ListMachineUserAssignedResourceRolesDefault {
	return &ListMachineUserAssignedResourceRolesDefault{
		_statusCode: code,
	}
}

/*
ListMachineUserAssignedResourceRolesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListMachineUserAssignedResourceRolesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list machine user assigned resource roles default response has a 2xx status code
func (o *ListMachineUserAssignedResourceRolesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list machine user assigned resource roles default response has a 3xx status code
func (o *ListMachineUserAssignedResourceRolesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list machine user assigned resource roles default response has a 4xx status code
func (o *ListMachineUserAssignedResourceRolesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list machine user assigned resource roles default response has a 5xx status code
func (o *ListMachineUserAssignedResourceRolesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list machine user assigned resource roles default response a status code equal to that given
func (o *ListMachineUserAssignedResourceRolesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list machine user assigned resource roles default response
func (o *ListMachineUserAssignedResourceRolesDefault) Code() int {
	return o._statusCode
}

func (o *ListMachineUserAssignedResourceRolesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/listMachineUserAssignedResourceRoles][%d] listMachineUserAssignedResourceRoles default %s", o._statusCode, payload)
}

func (o *ListMachineUserAssignedResourceRolesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/listMachineUserAssignedResourceRoles][%d] listMachineUserAssignedResourceRoles default %s", o._statusCode, payload)
}

func (o *ListMachineUserAssignedResourceRolesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListMachineUserAssignedResourceRolesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
