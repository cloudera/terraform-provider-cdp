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

// DescribeClusterDiagnosticDataJobReader is a Reader for the DescribeClusterDiagnosticDataJob structure.
type DescribeClusterDiagnosticDataJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeClusterDiagnosticDataJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeClusterDiagnosticDataJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeClusterDiagnosticDataJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeClusterDiagnosticDataJobOK creates a DescribeClusterDiagnosticDataJobOK with default headers values
func NewDescribeClusterDiagnosticDataJobOK() *DescribeClusterDiagnosticDataJobOK {
	return &DescribeClusterDiagnosticDataJobOK{}
}

/*
DescribeClusterDiagnosticDataJobOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeClusterDiagnosticDataJobOK struct {
	Payload *models.DescribeClusterDiagnosticDataJobResponse
}

// IsSuccess returns true when this describe cluster diagnostic data job o k response has a 2xx status code
func (o *DescribeClusterDiagnosticDataJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe cluster diagnostic data job o k response has a 3xx status code
func (o *DescribeClusterDiagnosticDataJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe cluster diagnostic data job o k response has a 4xx status code
func (o *DescribeClusterDiagnosticDataJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe cluster diagnostic data job o k response has a 5xx status code
func (o *DescribeClusterDiagnosticDataJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe cluster diagnostic data job o k response a status code equal to that given
func (o *DescribeClusterDiagnosticDataJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe cluster diagnostic data job o k response
func (o *DescribeClusterDiagnosticDataJobOK) Code() int {
	return 200
}

func (o *DescribeClusterDiagnosticDataJobOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeClusterDiagnosticDataJob][%d] describeClusterDiagnosticDataJobOK %s", 200, payload)
}

func (o *DescribeClusterDiagnosticDataJobOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeClusterDiagnosticDataJob][%d] describeClusterDiagnosticDataJobOK %s", 200, payload)
}

func (o *DescribeClusterDiagnosticDataJobOK) GetPayload() *models.DescribeClusterDiagnosticDataJobResponse {
	return o.Payload
}

func (o *DescribeClusterDiagnosticDataJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeClusterDiagnosticDataJobResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeClusterDiagnosticDataJobDefault creates a DescribeClusterDiagnosticDataJobDefault with default headers values
func NewDescribeClusterDiagnosticDataJobDefault(code int) *DescribeClusterDiagnosticDataJobDefault {
	return &DescribeClusterDiagnosticDataJobDefault{
		_statusCode: code,
	}
}

/*
DescribeClusterDiagnosticDataJobDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeClusterDiagnosticDataJobDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe cluster diagnostic data job default response has a 2xx status code
func (o *DescribeClusterDiagnosticDataJobDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe cluster diagnostic data job default response has a 3xx status code
func (o *DescribeClusterDiagnosticDataJobDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe cluster diagnostic data job default response has a 4xx status code
func (o *DescribeClusterDiagnosticDataJobDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe cluster diagnostic data job default response has a 5xx status code
func (o *DescribeClusterDiagnosticDataJobDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe cluster diagnostic data job default response a status code equal to that given
func (o *DescribeClusterDiagnosticDataJobDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe cluster diagnostic data job default response
func (o *DescribeClusterDiagnosticDataJobDefault) Code() int {
	return o._statusCode
}

func (o *DescribeClusterDiagnosticDataJobDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeClusterDiagnosticDataJob][%d] describeClusterDiagnosticDataJob default %s", o._statusCode, payload)
}

func (o *DescribeClusterDiagnosticDataJobDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeClusterDiagnosticDataJob][%d] describeClusterDiagnosticDataJob default %s", o._statusCode, payload)
}

func (o *DescribeClusterDiagnosticDataJobDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeClusterDiagnosticDataJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
