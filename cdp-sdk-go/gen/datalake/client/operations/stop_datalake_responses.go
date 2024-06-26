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

// StopDatalakeReader is a Reader for the StopDatalake structure.
type StopDatalakeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopDatalakeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopDatalakeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStopDatalakeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStopDatalakeOK creates a StopDatalakeOK with default headers values
func NewStopDatalakeOK() *StopDatalakeOK {
	return &StopDatalakeOK{}
}

/*
StopDatalakeOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type StopDatalakeOK struct {
	Payload models.StopDatalakeResponse
}

// IsSuccess returns true when this stop datalake o k response has a 2xx status code
func (o *StopDatalakeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop datalake o k response has a 3xx status code
func (o *StopDatalakeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop datalake o k response has a 4xx status code
func (o *StopDatalakeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop datalake o k response has a 5xx status code
func (o *StopDatalakeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop datalake o k response a status code equal to that given
func (o *StopDatalakeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop datalake o k response
func (o *StopDatalakeOK) Code() int {
	return 200
}

func (o *StopDatalakeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/stopDatalake][%d] stopDatalakeOK %s", 200, payload)
}

func (o *StopDatalakeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/stopDatalake][%d] stopDatalakeOK %s", 200, payload)
}

func (o *StopDatalakeOK) GetPayload() models.StopDatalakeResponse {
	return o.Payload
}

func (o *StopDatalakeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopDatalakeDefault creates a StopDatalakeDefault with default headers values
func NewStopDatalakeDefault(code int) *StopDatalakeDefault {
	return &StopDatalakeDefault{
		_statusCode: code,
	}
}

/*
StopDatalakeDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type StopDatalakeDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this stop datalake default response has a 2xx status code
func (o *StopDatalakeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this stop datalake default response has a 3xx status code
func (o *StopDatalakeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this stop datalake default response has a 4xx status code
func (o *StopDatalakeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this stop datalake default response has a 5xx status code
func (o *StopDatalakeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this stop datalake default response a status code equal to that given
func (o *StopDatalakeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the stop datalake default response
func (o *StopDatalakeDefault) Code() int {
	return o._statusCode
}

func (o *StopDatalakeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/stopDatalake][%d] stopDatalake default %s", o._statusCode, payload)
}

func (o *StopDatalakeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/stopDatalake][%d] stopDatalake default %s", o._statusCode, payload)
}

func (o *StopDatalakeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopDatalakeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
