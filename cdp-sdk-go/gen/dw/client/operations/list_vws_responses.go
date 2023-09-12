// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// ListVwsReader is a Reader for the ListVws structure.
type ListVwsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVwsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVwsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVwsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVwsOK creates a ListVwsOK with default headers values
func NewListVwsOK() *ListVwsOK {
	return &ListVwsOK{}
}

/*
ListVwsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListVwsOK struct {
	Payload *models.ListVwsResponse
}

// IsSuccess returns true when this list vws o k response has a 2xx status code
func (o *ListVwsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list vws o k response has a 3xx status code
func (o *ListVwsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list vws o k response has a 4xx status code
func (o *ListVwsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list vws o k response has a 5xx status code
func (o *ListVwsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list vws o k response a status code equal to that given
func (o *ListVwsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list vws o k response
func (o *ListVwsOK) Code() int {
	return 200
}

func (o *ListVwsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/listVws][%d] listVwsOK  %+v", 200, o.Payload)
}

func (o *ListVwsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/listVws][%d] listVwsOK  %+v", 200, o.Payload)
}

func (o *ListVwsOK) GetPayload() *models.ListVwsResponse {
	return o.Payload
}

func (o *ListVwsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListVwsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVwsDefault creates a ListVwsDefault with default headers values
func NewListVwsDefault(code int) *ListVwsDefault {
	return &ListVwsDefault{
		_statusCode: code,
	}
}

/*
ListVwsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListVwsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list vws default response has a 2xx status code
func (o *ListVwsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list vws default response has a 3xx status code
func (o *ListVwsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list vws default response has a 4xx status code
func (o *ListVwsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list vws default response has a 5xx status code
func (o *ListVwsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list vws default response a status code equal to that given
func (o *ListVwsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list vws default response
func (o *ListVwsDefault) Code() int {
	return o._statusCode
}

func (o *ListVwsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/listVws][%d] listVws default  %+v", o._statusCode, o.Payload)
}

func (o *ListVwsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/listVws][%d] listVws default  %+v", o._statusCode, o.Payload)
}

func (o *ListVwsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListVwsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}