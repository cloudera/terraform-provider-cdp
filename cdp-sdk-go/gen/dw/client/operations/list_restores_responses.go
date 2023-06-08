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

// ListRestoresReader is a Reader for the ListRestores structure.
type ListRestoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRestoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRestoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListRestoresDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRestoresOK creates a ListRestoresOK with default headers values
func NewListRestoresOK() *ListRestoresOK {
	return &ListRestoresOK{}
}

/*
ListRestoresOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListRestoresOK struct {
	Payload *models.ListRestoresResponse
}

// IsSuccess returns true when this list restores o k response has a 2xx status code
func (o *ListRestoresOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list restores o k response has a 3xx status code
func (o *ListRestoresOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list restores o k response has a 4xx status code
func (o *ListRestoresOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list restores o k response has a 5xx status code
func (o *ListRestoresOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list restores o k response a status code equal to that given
func (o *ListRestoresOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list restores o k response
func (o *ListRestoresOK) Code() int {
	return 200
}

func (o *ListRestoresOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/listRestores][%d] listRestoresOK  %+v", 200, o.Payload)
}

func (o *ListRestoresOK) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/listRestores][%d] listRestoresOK  %+v", 200, o.Payload)
}

func (o *ListRestoresOK) GetPayload() *models.ListRestoresResponse {
	return o.Payload
}

func (o *ListRestoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListRestoresResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRestoresDefault creates a ListRestoresDefault with default headers values
func NewListRestoresDefault(code int) *ListRestoresDefault {
	return &ListRestoresDefault{
		_statusCode: code,
	}
}

/*
ListRestoresDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListRestoresDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list restores default response has a 2xx status code
func (o *ListRestoresDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list restores default response has a 3xx status code
func (o *ListRestoresDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list restores default response has a 4xx status code
func (o *ListRestoresDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list restores default response has a 5xx status code
func (o *ListRestoresDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list restores default response a status code equal to that given
func (o *ListRestoresDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list restores default response
func (o *ListRestoresDefault) Code() int {
	return o._statusCode
}

func (o *ListRestoresDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/listRestores][%d] listRestores default  %+v", o._statusCode, o.Payload)
}

func (o *ListRestoresDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/listRestores][%d] listRestores default  %+v", o._statusCode, o.Payload)
}

func (o *ListRestoresDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListRestoresDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
