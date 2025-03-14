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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// StopEnvironmentReader is a Reader for the StopEnvironment structure.
type StopEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStopEnvironmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStopEnvironmentOK creates a StopEnvironmentOK with default headers values
func NewStopEnvironmentOK() *StopEnvironmentOK {
	return &StopEnvironmentOK{}
}

/*
StopEnvironmentOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type StopEnvironmentOK struct {
	Payload *models.StopEnvironmentResponse
}

// IsSuccess returns true when this stop environment o k response has a 2xx status code
func (o *StopEnvironmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop environment o k response has a 3xx status code
func (o *StopEnvironmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop environment o k response has a 4xx status code
func (o *StopEnvironmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop environment o k response has a 5xx status code
func (o *StopEnvironmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop environment o k response a status code equal to that given
func (o *StopEnvironmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop environment o k response
func (o *StopEnvironmentOK) Code() int {
	return 200
}

func (o *StopEnvironmentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/stopEnvironment][%d] stopEnvironmentOK %s", 200, payload)
}

func (o *StopEnvironmentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/stopEnvironment][%d] stopEnvironmentOK %s", 200, payload)
}

func (o *StopEnvironmentOK) GetPayload() *models.StopEnvironmentResponse {
	return o.Payload
}

func (o *StopEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StopEnvironmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopEnvironmentDefault creates a StopEnvironmentDefault with default headers values
func NewStopEnvironmentDefault(code int) *StopEnvironmentDefault {
	return &StopEnvironmentDefault{
		_statusCode: code,
	}
}

/*
StopEnvironmentDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type StopEnvironmentDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this stop environment default response has a 2xx status code
func (o *StopEnvironmentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this stop environment default response has a 3xx status code
func (o *StopEnvironmentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this stop environment default response has a 4xx status code
func (o *StopEnvironmentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this stop environment default response has a 5xx status code
func (o *StopEnvironmentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this stop environment default response a status code equal to that given
func (o *StopEnvironmentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the stop environment default response
func (o *StopEnvironmentDefault) Code() int {
	return o._statusCode
}

func (o *StopEnvironmentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/stopEnvironment][%d] stopEnvironment default %s", o._statusCode, payload)
}

func (o *StopEnvironmentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/stopEnvironment][%d] stopEnvironment default %s", o._statusCode, payload)
}

func (o *StopEnvironmentDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopEnvironmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
