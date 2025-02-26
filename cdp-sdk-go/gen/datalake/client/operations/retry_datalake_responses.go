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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// RetryDatalakeReader is a Reader for the RetryDatalake structure.
type RetryDatalakeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetryDatalakeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetryDatalakeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRetryDatalakeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRetryDatalakeOK creates a RetryDatalakeOK with default headers values
func NewRetryDatalakeOK() *RetryDatalakeOK {
	return &RetryDatalakeOK{}
}

/*
RetryDatalakeOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type RetryDatalakeOK struct {
	Payload *models.RetryDatalakeResponse
}

// IsSuccess returns true when this retry datalake o k response has a 2xx status code
func (o *RetryDatalakeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this retry datalake o k response has a 3xx status code
func (o *RetryDatalakeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this retry datalake o k response has a 4xx status code
func (o *RetryDatalakeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this retry datalake o k response has a 5xx status code
func (o *RetryDatalakeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this retry datalake o k response a status code equal to that given
func (o *RetryDatalakeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the retry datalake o k response
func (o *RetryDatalakeOK) Code() int {
	return 200
}

func (o *RetryDatalakeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/retryDatalake][%d] retryDatalakeOK %s", 200, payload)
}

func (o *RetryDatalakeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/retryDatalake][%d] retryDatalakeOK %s", 200, payload)
}

func (o *RetryDatalakeOK) GetPayload() *models.RetryDatalakeResponse {
	return o.Payload
}

func (o *RetryDatalakeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RetryDatalakeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetryDatalakeDefault creates a RetryDatalakeDefault with default headers values
func NewRetryDatalakeDefault(code int) *RetryDatalakeDefault {
	return &RetryDatalakeDefault{
		_statusCode: code,
	}
}

/*
RetryDatalakeDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type RetryDatalakeDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this retry datalake default response has a 2xx status code
func (o *RetryDatalakeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this retry datalake default response has a 3xx status code
func (o *RetryDatalakeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this retry datalake default response has a 4xx status code
func (o *RetryDatalakeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this retry datalake default response has a 5xx status code
func (o *RetryDatalakeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this retry datalake default response a status code equal to that given
func (o *RetryDatalakeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the retry datalake default response
func (o *RetryDatalakeDefault) Code() int {
	return o._statusCode
}

func (o *RetryDatalakeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/retryDatalake][%d] retryDatalake default %s", o._statusCode, payload)
}

func (o *RetryDatalakeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/retryDatalake][%d] retryDatalake default %s", o._statusCode, payload)
}

func (o *RetryDatalakeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RetryDatalakeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
