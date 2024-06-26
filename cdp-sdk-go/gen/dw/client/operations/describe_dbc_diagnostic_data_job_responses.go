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

// DescribeDbcDiagnosticDataJobReader is a Reader for the DescribeDbcDiagnosticDataJob structure.
type DescribeDbcDiagnosticDataJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeDbcDiagnosticDataJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeDbcDiagnosticDataJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeDbcDiagnosticDataJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeDbcDiagnosticDataJobOK creates a DescribeDbcDiagnosticDataJobOK with default headers values
func NewDescribeDbcDiagnosticDataJobOK() *DescribeDbcDiagnosticDataJobOK {
	return &DescribeDbcDiagnosticDataJobOK{}
}

/*
DescribeDbcDiagnosticDataJobOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeDbcDiagnosticDataJobOK struct {
	Payload *models.DescribeDbcDiagnosticDataJobResponse
}

// IsSuccess returns true when this describe dbc diagnostic data job o k response has a 2xx status code
func (o *DescribeDbcDiagnosticDataJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe dbc diagnostic data job o k response has a 3xx status code
func (o *DescribeDbcDiagnosticDataJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe dbc diagnostic data job o k response has a 4xx status code
func (o *DescribeDbcDiagnosticDataJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe dbc diagnostic data job o k response has a 5xx status code
func (o *DescribeDbcDiagnosticDataJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe dbc diagnostic data job o k response a status code equal to that given
func (o *DescribeDbcDiagnosticDataJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe dbc diagnostic data job o k response
func (o *DescribeDbcDiagnosticDataJobOK) Code() int {
	return 200
}

func (o *DescribeDbcDiagnosticDataJobOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeDbcDiagnosticDataJob][%d] describeDbcDiagnosticDataJobOK %s", 200, payload)
}

func (o *DescribeDbcDiagnosticDataJobOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeDbcDiagnosticDataJob][%d] describeDbcDiagnosticDataJobOK %s", 200, payload)
}

func (o *DescribeDbcDiagnosticDataJobOK) GetPayload() *models.DescribeDbcDiagnosticDataJobResponse {
	return o.Payload
}

func (o *DescribeDbcDiagnosticDataJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeDbcDiagnosticDataJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeDbcDiagnosticDataJobDefault creates a DescribeDbcDiagnosticDataJobDefault with default headers values
func NewDescribeDbcDiagnosticDataJobDefault(code int) *DescribeDbcDiagnosticDataJobDefault {
	return &DescribeDbcDiagnosticDataJobDefault{
		_statusCode: code,
	}
}

/*
DescribeDbcDiagnosticDataJobDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeDbcDiagnosticDataJobDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe dbc diagnostic data job default response has a 2xx status code
func (o *DescribeDbcDiagnosticDataJobDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe dbc diagnostic data job default response has a 3xx status code
func (o *DescribeDbcDiagnosticDataJobDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe dbc diagnostic data job default response has a 4xx status code
func (o *DescribeDbcDiagnosticDataJobDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe dbc diagnostic data job default response has a 5xx status code
func (o *DescribeDbcDiagnosticDataJobDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe dbc diagnostic data job default response a status code equal to that given
func (o *DescribeDbcDiagnosticDataJobDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe dbc diagnostic data job default response
func (o *DescribeDbcDiagnosticDataJobDefault) Code() int {
	return o._statusCode
}

func (o *DescribeDbcDiagnosticDataJobDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeDbcDiagnosticDataJob][%d] describeDbcDiagnosticDataJob default %s", o._statusCode, payload)
}

func (o *DescribeDbcDiagnosticDataJobDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeDbcDiagnosticDataJob][%d] describeDbcDiagnosticDataJob default %s", o._statusCode, payload)
}

func (o *DescribeDbcDiagnosticDataJobDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeDbcDiagnosticDataJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
