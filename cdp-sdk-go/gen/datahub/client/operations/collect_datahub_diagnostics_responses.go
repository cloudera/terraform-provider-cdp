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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// CollectDatahubDiagnosticsReader is a Reader for the CollectDatahubDiagnostics structure.
type CollectDatahubDiagnosticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectDatahubDiagnosticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCollectDatahubDiagnosticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCollectDatahubDiagnosticsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCollectDatahubDiagnosticsOK creates a CollectDatahubDiagnosticsOK with default headers values
func NewCollectDatahubDiagnosticsOK() *CollectDatahubDiagnosticsOK {
	return &CollectDatahubDiagnosticsOK{}
}

/*
CollectDatahubDiagnosticsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CollectDatahubDiagnosticsOK struct {
	Payload models.CollectDatahubDiagnosticsResponse
}

// IsSuccess returns true when this collect datahub diagnostics o k response has a 2xx status code
func (o *CollectDatahubDiagnosticsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this collect datahub diagnostics o k response has a 3xx status code
func (o *CollectDatahubDiagnosticsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect datahub diagnostics o k response has a 4xx status code
func (o *CollectDatahubDiagnosticsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this collect datahub diagnostics o k response has a 5xx status code
func (o *CollectDatahubDiagnosticsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this collect datahub diagnostics o k response a status code equal to that given
func (o *CollectDatahubDiagnosticsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the collect datahub diagnostics o k response
func (o *CollectDatahubDiagnosticsOK) Code() int {
	return 200
}

func (o *CollectDatahubDiagnosticsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/collectDatahubDiagnostics][%d] collectDatahubDiagnosticsOK %s", 200, payload)
}

func (o *CollectDatahubDiagnosticsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/collectDatahubDiagnostics][%d] collectDatahubDiagnosticsOK %s", 200, payload)
}

func (o *CollectDatahubDiagnosticsOK) GetPayload() models.CollectDatahubDiagnosticsResponse {
	return o.Payload
}

func (o *CollectDatahubDiagnosticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectDatahubDiagnosticsDefault creates a CollectDatahubDiagnosticsDefault with default headers values
func NewCollectDatahubDiagnosticsDefault(code int) *CollectDatahubDiagnosticsDefault {
	return &CollectDatahubDiagnosticsDefault{
		_statusCode: code,
	}
}

/*
CollectDatahubDiagnosticsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CollectDatahubDiagnosticsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this collect datahub diagnostics default response has a 2xx status code
func (o *CollectDatahubDiagnosticsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this collect datahub diagnostics default response has a 3xx status code
func (o *CollectDatahubDiagnosticsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this collect datahub diagnostics default response has a 4xx status code
func (o *CollectDatahubDiagnosticsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this collect datahub diagnostics default response has a 5xx status code
func (o *CollectDatahubDiagnosticsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this collect datahub diagnostics default response a status code equal to that given
func (o *CollectDatahubDiagnosticsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the collect datahub diagnostics default response
func (o *CollectDatahubDiagnosticsDefault) Code() int {
	return o._statusCode
}

func (o *CollectDatahubDiagnosticsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/collectDatahubDiagnostics][%d] collectDatahubDiagnostics default %s", o._statusCode, payload)
}

func (o *CollectDatahubDiagnosticsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/collectDatahubDiagnostics][%d] collectDatahubDiagnostics default %s", o._statusCode, payload)
}

func (o *CollectDatahubDiagnosticsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CollectDatahubDiagnosticsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
