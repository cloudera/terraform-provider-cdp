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

// CollectDatalakeDiagnosticsReader is a Reader for the CollectDatalakeDiagnostics structure.
type CollectDatalakeDiagnosticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectDatalakeDiagnosticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCollectDatalakeDiagnosticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCollectDatalakeDiagnosticsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCollectDatalakeDiagnosticsOK creates a CollectDatalakeDiagnosticsOK with default headers values
func NewCollectDatalakeDiagnosticsOK() *CollectDatalakeDiagnosticsOK {
	return &CollectDatalakeDiagnosticsOK{}
}

/*
CollectDatalakeDiagnosticsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CollectDatalakeDiagnosticsOK struct {
	Payload *models.CollectDatalakeDiagnosticsResponse
}

// IsSuccess returns true when this collect datalake diagnostics o k response has a 2xx status code
func (o *CollectDatalakeDiagnosticsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this collect datalake diagnostics o k response has a 3xx status code
func (o *CollectDatalakeDiagnosticsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect datalake diagnostics o k response has a 4xx status code
func (o *CollectDatalakeDiagnosticsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this collect datalake diagnostics o k response has a 5xx status code
func (o *CollectDatalakeDiagnosticsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this collect datalake diagnostics o k response a status code equal to that given
func (o *CollectDatalakeDiagnosticsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the collect datalake diagnostics o k response
func (o *CollectDatalakeDiagnosticsOK) Code() int {
	return 200
}

func (o *CollectDatalakeDiagnosticsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/collectDatalakeDiagnostics][%d] collectDatalakeDiagnosticsOK %s", 200, payload)
}

func (o *CollectDatalakeDiagnosticsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/collectDatalakeDiagnostics][%d] collectDatalakeDiagnosticsOK %s", 200, payload)
}

func (o *CollectDatalakeDiagnosticsOK) GetPayload() *models.CollectDatalakeDiagnosticsResponse {
	return o.Payload
}

func (o *CollectDatalakeDiagnosticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CollectDatalakeDiagnosticsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectDatalakeDiagnosticsDefault creates a CollectDatalakeDiagnosticsDefault with default headers values
func NewCollectDatalakeDiagnosticsDefault(code int) *CollectDatalakeDiagnosticsDefault {
	return &CollectDatalakeDiagnosticsDefault{
		_statusCode: code,
	}
}

/*
CollectDatalakeDiagnosticsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CollectDatalakeDiagnosticsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this collect datalake diagnostics default response has a 2xx status code
func (o *CollectDatalakeDiagnosticsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this collect datalake diagnostics default response has a 3xx status code
func (o *CollectDatalakeDiagnosticsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this collect datalake diagnostics default response has a 4xx status code
func (o *CollectDatalakeDiagnosticsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this collect datalake diagnostics default response has a 5xx status code
func (o *CollectDatalakeDiagnosticsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this collect datalake diagnostics default response a status code equal to that given
func (o *CollectDatalakeDiagnosticsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the collect datalake diagnostics default response
func (o *CollectDatalakeDiagnosticsDefault) Code() int {
	return o._statusCode
}

func (o *CollectDatalakeDiagnosticsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/collectDatalakeDiagnostics][%d] collectDatalakeDiagnostics default %s", o._statusCode, payload)
}

func (o *CollectDatalakeDiagnosticsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/collectDatalakeDiagnostics][%d] collectDatalakeDiagnostics default %s", o._statusCode, payload)
}

func (o *CollectDatalakeDiagnosticsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CollectDatalakeDiagnosticsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
