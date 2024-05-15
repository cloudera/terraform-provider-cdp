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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// ListDbcDiagnosticDataJobsReader is a Reader for the ListDbcDiagnosticDataJobs structure.
type ListDbcDiagnosticDataJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDbcDiagnosticDataJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDbcDiagnosticDataJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDbcDiagnosticDataJobsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDbcDiagnosticDataJobsOK creates a ListDbcDiagnosticDataJobsOK with default headers values
func NewListDbcDiagnosticDataJobsOK() *ListDbcDiagnosticDataJobsOK {
	return &ListDbcDiagnosticDataJobsOK{}
}

/*
ListDbcDiagnosticDataJobsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListDbcDiagnosticDataJobsOK struct {
	Payload *models.ListDbcDiagnosticDataJobsResponse
}

// IsSuccess returns true when this list dbc diagnostic data jobs o k response has a 2xx status code
func (o *ListDbcDiagnosticDataJobsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list dbc diagnostic data jobs o k response has a 3xx status code
func (o *ListDbcDiagnosticDataJobsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list dbc diagnostic data jobs o k response has a 4xx status code
func (o *ListDbcDiagnosticDataJobsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list dbc diagnostic data jobs o k response has a 5xx status code
func (o *ListDbcDiagnosticDataJobsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list dbc diagnostic data jobs o k response a status code equal to that given
func (o *ListDbcDiagnosticDataJobsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list dbc diagnostic data jobs o k response
func (o *ListDbcDiagnosticDataJobsOK) Code() int {
	return 200
}

func (o *ListDbcDiagnosticDataJobsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listDbcDiagnosticDataJobs][%d] listDbcDiagnosticDataJobsOK %s", 200, payload)
}

func (o *ListDbcDiagnosticDataJobsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listDbcDiagnosticDataJobs][%d] listDbcDiagnosticDataJobsOK %s", 200, payload)
}

func (o *ListDbcDiagnosticDataJobsOK) GetPayload() *models.ListDbcDiagnosticDataJobsResponse {
	return o.Payload
}

func (o *ListDbcDiagnosticDataJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDbcDiagnosticDataJobsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDbcDiagnosticDataJobsDefault creates a ListDbcDiagnosticDataJobsDefault with default headers values
func NewListDbcDiagnosticDataJobsDefault(code int) *ListDbcDiagnosticDataJobsDefault {
	return &ListDbcDiagnosticDataJobsDefault{
		_statusCode: code,
	}
}

/*
ListDbcDiagnosticDataJobsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListDbcDiagnosticDataJobsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list dbc diagnostic data jobs default response has a 2xx status code
func (o *ListDbcDiagnosticDataJobsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list dbc diagnostic data jobs default response has a 3xx status code
func (o *ListDbcDiagnosticDataJobsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list dbc diagnostic data jobs default response has a 4xx status code
func (o *ListDbcDiagnosticDataJobsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list dbc diagnostic data jobs default response has a 5xx status code
func (o *ListDbcDiagnosticDataJobsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list dbc diagnostic data jobs default response a status code equal to that given
func (o *ListDbcDiagnosticDataJobsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list dbc diagnostic data jobs default response
func (o *ListDbcDiagnosticDataJobsDefault) Code() int {
	return o._statusCode
}

func (o *ListDbcDiagnosticDataJobsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listDbcDiagnosticDataJobs][%d] listDbcDiagnosticDataJobs default %s", o._statusCode, payload)
}

func (o *ListDbcDiagnosticDataJobsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listDbcDiagnosticDataJobs][%d] listDbcDiagnosticDataJobs default %s", o._statusCode, payload)
}

func (o *ListDbcDiagnosticDataJobsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDbcDiagnosticDataJobsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
