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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// ListEventsReader is a Reader for the ListEvents structure.
type ListEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEventsOK creates a ListEventsOK with default headers values
func NewListEventsOK() *ListEventsOK {
	return &ListEventsOK{}
}

/*
ListEventsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListEventsOK struct {
	Payload *models.ListEventsResponse
}

// IsSuccess returns true when this list events o k response has a 2xx status code
func (o *ListEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list events o k response has a 3xx status code
func (o *ListEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list events o k response has a 4xx status code
func (o *ListEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list events o k response has a 5xx status code
func (o *ListEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list events o k response a status code equal to that given
func (o *ListEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list events o k response
func (o *ListEventsOK) Code() int {
	return 200
}

func (o *ListEventsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listEvents][%d] listEventsOK %s", 200, payload)
}

func (o *ListEventsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listEvents][%d] listEventsOK %s", 200, payload)
}

func (o *ListEventsOK) GetPayload() *models.ListEventsResponse {
	return o.Payload
}

func (o *ListEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEventsDefault creates a ListEventsDefault with default headers values
func NewListEventsDefault(code int) *ListEventsDefault {
	return &ListEventsDefault{
		_statusCode: code,
	}
}

/*
ListEventsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListEventsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list events default response has a 2xx status code
func (o *ListEventsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list events default response has a 3xx status code
func (o *ListEventsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list events default response has a 4xx status code
func (o *ListEventsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list events default response has a 5xx status code
func (o *ListEventsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list events default response a status code equal to that given
func (o *ListEventsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list events default response
func (o *ListEventsDefault) Code() int {
	return o._statusCode
}

func (o *ListEventsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listEvents][%d] listEvents default %s", o._statusCode, payload)
}

func (o *ListEventsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listEvents][%d] listEvents default %s", o._statusCode, payload)
}

func (o *ListEventsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
