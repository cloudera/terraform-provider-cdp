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

// ListDatahubDiagnosticsReader is a Reader for the ListDatahubDiagnostics structure.
type ListDatahubDiagnosticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDatahubDiagnosticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDatahubDiagnosticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDatahubDiagnosticsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDatahubDiagnosticsOK creates a ListDatahubDiagnosticsOK with default headers values
func NewListDatahubDiagnosticsOK() *ListDatahubDiagnosticsOK {
	return &ListDatahubDiagnosticsOK{}
}

/*
ListDatahubDiagnosticsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListDatahubDiagnosticsOK struct {
	Payload *models.ListDatahubDiagnosticsResponse
}

// IsSuccess returns true when this list datahub diagnostics o k response has a 2xx status code
func (o *ListDatahubDiagnosticsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list datahub diagnostics o k response has a 3xx status code
func (o *ListDatahubDiagnosticsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list datahub diagnostics o k response has a 4xx status code
func (o *ListDatahubDiagnosticsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list datahub diagnostics o k response has a 5xx status code
func (o *ListDatahubDiagnosticsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list datahub diagnostics o k response a status code equal to that given
func (o *ListDatahubDiagnosticsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list datahub diagnostics o k response
func (o *ListDatahubDiagnosticsOK) Code() int {
	return 200
}

func (o *ListDatahubDiagnosticsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/listDatahubDiagnostics][%d] listDatahubDiagnosticsOK %s", 200, payload)
}

func (o *ListDatahubDiagnosticsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/listDatahubDiagnostics][%d] listDatahubDiagnosticsOK %s", 200, payload)
}

func (o *ListDatahubDiagnosticsOK) GetPayload() *models.ListDatahubDiagnosticsResponse {
	return o.Payload
}

func (o *ListDatahubDiagnosticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDatahubDiagnosticsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDatahubDiagnosticsDefault creates a ListDatahubDiagnosticsDefault with default headers values
func NewListDatahubDiagnosticsDefault(code int) *ListDatahubDiagnosticsDefault {
	return &ListDatahubDiagnosticsDefault{
		_statusCode: code,
	}
}

/*
ListDatahubDiagnosticsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListDatahubDiagnosticsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list datahub diagnostics default response has a 2xx status code
func (o *ListDatahubDiagnosticsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list datahub diagnostics default response has a 3xx status code
func (o *ListDatahubDiagnosticsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list datahub diagnostics default response has a 4xx status code
func (o *ListDatahubDiagnosticsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list datahub diagnostics default response has a 5xx status code
func (o *ListDatahubDiagnosticsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list datahub diagnostics default response a status code equal to that given
func (o *ListDatahubDiagnosticsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list datahub diagnostics default response
func (o *ListDatahubDiagnosticsDefault) Code() int {
	return o._statusCode
}

func (o *ListDatahubDiagnosticsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/listDatahubDiagnostics][%d] listDatahubDiagnostics default %s", o._statusCode, payload)
}

func (o *ListDatahubDiagnosticsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/listDatahubDiagnostics][%d] listDatahubDiagnostics default %s", o._statusCode, payload)
}

func (o *ListDatahubDiagnosticsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDatahubDiagnosticsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
