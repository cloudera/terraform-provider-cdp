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

// ListGroupMembersReader is a Reader for the ListGroupMembers structure.
type ListGroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGroupMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListGroupMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGroupMembersOK creates a ListGroupMembersOK with default headers values
func NewListGroupMembersOK() *ListGroupMembersOK {
	return &ListGroupMembersOK{}
}

/*
ListGroupMembersOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListGroupMembersOK struct {
	Payload *models.ListGroupMembersResponse
}

// IsSuccess returns true when this list group members o k response has a 2xx status code
func (o *ListGroupMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list group members o k response has a 3xx status code
func (o *ListGroupMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list group members o k response has a 4xx status code
func (o *ListGroupMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list group members o k response has a 5xx status code
func (o *ListGroupMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list group members o k response a status code equal to that given
func (o *ListGroupMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list group members o k response
func (o *ListGroupMembersOK) Code() int {
	return 200
}

func (o *ListGroupMembersOK) Error() string {
	return fmt.Sprintf("[POST /iam/listGroupMembers][%d] listGroupMembersOK  %+v", 200, o.Payload)
}

func (o *ListGroupMembersOK) String() string {
	return fmt.Sprintf("[POST /iam/listGroupMembers][%d] listGroupMembersOK  %+v", 200, o.Payload)
}

func (o *ListGroupMembersOK) GetPayload() *models.ListGroupMembersResponse {
	return o.Payload
}

func (o *ListGroupMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListGroupMembersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupMembersDefault creates a ListGroupMembersDefault with default headers values
func NewListGroupMembersDefault(code int) *ListGroupMembersDefault {
	return &ListGroupMembersDefault{
		_statusCode: code,
	}
}

/*
ListGroupMembersDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListGroupMembersDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list group members default response has a 2xx status code
func (o *ListGroupMembersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list group members default response has a 3xx status code
func (o *ListGroupMembersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list group members default response has a 4xx status code
func (o *ListGroupMembersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list group members default response has a 5xx status code
func (o *ListGroupMembersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list group members default response a status code equal to that given
func (o *ListGroupMembersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list group members default response
func (o *ListGroupMembersDefault) Code() int {
	return o._statusCode
}

func (o *ListGroupMembersDefault) Error() string {
	return fmt.Sprintf("[POST /iam/listGroupMembers][%d] listGroupMembers default  %+v", o._statusCode, o.Payload)
}

func (o *ListGroupMembersDefault) String() string {
	return fmt.Sprintf("[POST /iam/listGroupMembers][%d] listGroupMembers default  %+v", o._statusCode, o.Payload)
}

func (o *ListGroupMembersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListGroupMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}