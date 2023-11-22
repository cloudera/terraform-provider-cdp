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

// SuspendVwReader is a Reader for the SuspendVw structure.
type SuspendVwReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuspendVwReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuspendVwOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSuspendVwDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSuspendVwOK creates a SuspendVwOK with default headers values
func NewSuspendVwOK() *SuspendVwOK {
	return &SuspendVwOK{}
}

/*
SuspendVwOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SuspendVwOK struct {
	Payload models.SuspendVwResponse
}

// IsSuccess returns true when this suspend vw o k response has a 2xx status code
func (o *SuspendVwOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this suspend vw o k response has a 3xx status code
func (o *SuspendVwOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this suspend vw o k response has a 4xx status code
func (o *SuspendVwOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this suspend vw o k response has a 5xx status code
func (o *SuspendVwOK) IsServerError() bool {
	return false
}

// IsCode returns true when this suspend vw o k response a status code equal to that given
func (o *SuspendVwOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the suspend vw o k response
func (o *SuspendVwOK) Code() int {
	return 200
}

func (o *SuspendVwOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/suspendVw][%d] suspendVwOK  %+v", 200, o.Payload)
}

func (o *SuspendVwOK) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/suspendVw][%d] suspendVwOK  %+v", 200, o.Payload)
}

func (o *SuspendVwOK) GetPayload() models.SuspendVwResponse {
	return o.Payload
}

func (o *SuspendVwOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuspendVwDefault creates a SuspendVwDefault with default headers values
func NewSuspendVwDefault(code int) *SuspendVwDefault {
	return &SuspendVwDefault{
		_statusCode: code,
	}
}

/*
SuspendVwDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SuspendVwDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this suspend vw default response has a 2xx status code
func (o *SuspendVwDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this suspend vw default response has a 3xx status code
func (o *SuspendVwDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this suspend vw default response has a 4xx status code
func (o *SuspendVwDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this suspend vw default response has a 5xx status code
func (o *SuspendVwDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this suspend vw default response a status code equal to that given
func (o *SuspendVwDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the suspend vw default response
func (o *SuspendVwDefault) Code() int {
	return o._statusCode
}

func (o *SuspendVwDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/suspendVw][%d] suspendVw default  %+v", o._statusCode, o.Payload)
}

func (o *SuspendVwDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/suspendVw][%d] suspendVw default  %+v", o._statusCode, o.Payload)
}

func (o *SuspendVwDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SuspendVwDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}