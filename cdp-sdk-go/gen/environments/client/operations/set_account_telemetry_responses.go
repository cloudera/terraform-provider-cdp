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

// SetAccountTelemetryReader is a Reader for the SetAccountTelemetry structure.
type SetAccountTelemetryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAccountTelemetryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetAccountTelemetryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetAccountTelemetryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetAccountTelemetryOK creates a SetAccountTelemetryOK with default headers values
func NewSetAccountTelemetryOK() *SetAccountTelemetryOK {
	return &SetAccountTelemetryOK{}
}

/*
SetAccountTelemetryOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SetAccountTelemetryOK struct {
	Payload *models.SetAccountTelemetryResponse
}

// IsSuccess returns true when this set account telemetry o k response has a 2xx status code
func (o *SetAccountTelemetryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set account telemetry o k response has a 3xx status code
func (o *SetAccountTelemetryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set account telemetry o k response has a 4xx status code
func (o *SetAccountTelemetryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set account telemetry o k response has a 5xx status code
func (o *SetAccountTelemetryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set account telemetry o k response a status code equal to that given
func (o *SetAccountTelemetryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set account telemetry o k response
func (o *SetAccountTelemetryOK) Code() int {
	return 200
}

func (o *SetAccountTelemetryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/setAccountTelemetry][%d] setAccountTelemetryOK %s", 200, payload)
}

func (o *SetAccountTelemetryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/setAccountTelemetry][%d] setAccountTelemetryOK %s", 200, payload)
}

func (o *SetAccountTelemetryOK) GetPayload() *models.SetAccountTelemetryResponse {
	return o.Payload
}

func (o *SetAccountTelemetryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SetAccountTelemetryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAccountTelemetryDefault creates a SetAccountTelemetryDefault with default headers values
func NewSetAccountTelemetryDefault(code int) *SetAccountTelemetryDefault {
	return &SetAccountTelemetryDefault{
		_statusCode: code,
	}
}

/*
SetAccountTelemetryDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SetAccountTelemetryDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this set account telemetry default response has a 2xx status code
func (o *SetAccountTelemetryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this set account telemetry default response has a 3xx status code
func (o *SetAccountTelemetryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this set account telemetry default response has a 4xx status code
func (o *SetAccountTelemetryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this set account telemetry default response has a 5xx status code
func (o *SetAccountTelemetryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this set account telemetry default response a status code equal to that given
func (o *SetAccountTelemetryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the set account telemetry default response
func (o *SetAccountTelemetryDefault) Code() int {
	return o._statusCode
}

func (o *SetAccountTelemetryDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/setAccountTelemetry][%d] setAccountTelemetry default %s", o._statusCode, payload)
}

func (o *SetAccountTelemetryDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/setAccountTelemetry][%d] setAccountTelemetry default %s", o._statusCode, payload)
}

func (o *SetAccountTelemetryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetAccountTelemetryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
