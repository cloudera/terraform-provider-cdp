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

// ListDatalakeDiagnosticsReader is a Reader for the ListDatalakeDiagnostics structure.
type ListDatalakeDiagnosticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDatalakeDiagnosticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDatalakeDiagnosticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDatalakeDiagnosticsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDatalakeDiagnosticsOK creates a ListDatalakeDiagnosticsOK with default headers values
func NewListDatalakeDiagnosticsOK() *ListDatalakeDiagnosticsOK {
	return &ListDatalakeDiagnosticsOK{}
}

/*
ListDatalakeDiagnosticsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListDatalakeDiagnosticsOK struct {
	Payload *models.ListDatalakeDiagnosticsResponse
}

// IsSuccess returns true when this list datalake diagnostics o k response has a 2xx status code
func (o *ListDatalakeDiagnosticsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list datalake diagnostics o k response has a 3xx status code
func (o *ListDatalakeDiagnosticsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list datalake diagnostics o k response has a 4xx status code
func (o *ListDatalakeDiagnosticsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list datalake diagnostics o k response has a 5xx status code
func (o *ListDatalakeDiagnosticsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list datalake diagnostics o k response a status code equal to that given
func (o *ListDatalakeDiagnosticsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list datalake diagnostics o k response
func (o *ListDatalakeDiagnosticsOK) Code() int {
	return 200
}

func (o *ListDatalakeDiagnosticsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/listDatalakeDiagnostics][%d] listDatalakeDiagnosticsOK %s", 200, payload)
}

func (o *ListDatalakeDiagnosticsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/listDatalakeDiagnostics][%d] listDatalakeDiagnosticsOK %s", 200, payload)
}

func (o *ListDatalakeDiagnosticsOK) GetPayload() *models.ListDatalakeDiagnosticsResponse {
	return o.Payload
}

func (o *ListDatalakeDiagnosticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDatalakeDiagnosticsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDatalakeDiagnosticsDefault creates a ListDatalakeDiagnosticsDefault with default headers values
func NewListDatalakeDiagnosticsDefault(code int) *ListDatalakeDiagnosticsDefault {
	return &ListDatalakeDiagnosticsDefault{
		_statusCode: code,
	}
}

/*
ListDatalakeDiagnosticsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListDatalakeDiagnosticsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list datalake diagnostics default response has a 2xx status code
func (o *ListDatalakeDiagnosticsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list datalake diagnostics default response has a 3xx status code
func (o *ListDatalakeDiagnosticsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list datalake diagnostics default response has a 4xx status code
func (o *ListDatalakeDiagnosticsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list datalake diagnostics default response has a 5xx status code
func (o *ListDatalakeDiagnosticsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list datalake diagnostics default response a status code equal to that given
func (o *ListDatalakeDiagnosticsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list datalake diagnostics default response
func (o *ListDatalakeDiagnosticsDefault) Code() int {
	return o._statusCode
}

func (o *ListDatalakeDiagnosticsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/listDatalakeDiagnostics][%d] listDatalakeDiagnostics default %s", o._statusCode, payload)
}

func (o *ListDatalakeDiagnosticsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/listDatalakeDiagnostics][%d] listDatalakeDiagnostics default %s", o._statusCode, payload)
}

func (o *ListDatalakeDiagnosticsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDatalakeDiagnosticsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
