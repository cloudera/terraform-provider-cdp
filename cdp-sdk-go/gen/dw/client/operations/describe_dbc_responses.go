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

// DescribeDbcReader is a Reader for the DescribeDbc structure.
type DescribeDbcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeDbcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeDbcOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeDbcDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeDbcOK creates a DescribeDbcOK with default headers values
func NewDescribeDbcOK() *DescribeDbcOK {
	return &DescribeDbcOK{}
}

/*
DescribeDbcOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeDbcOK struct {
	Payload *models.DescribeDbcResponse
}

// IsSuccess returns true when this describe dbc o k response has a 2xx status code
func (o *DescribeDbcOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe dbc o k response has a 3xx status code
func (o *DescribeDbcOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe dbc o k response has a 4xx status code
func (o *DescribeDbcOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe dbc o k response has a 5xx status code
func (o *DescribeDbcOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe dbc o k response a status code equal to that given
func (o *DescribeDbcOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe dbc o k response
func (o *DescribeDbcOK) Code() int {
	return 200
}

func (o *DescribeDbcOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeDbc][%d] describeDbcOK %s", 200, payload)
}

func (o *DescribeDbcOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeDbc][%d] describeDbcOK %s", 200, payload)
}

func (o *DescribeDbcOK) GetPayload() *models.DescribeDbcResponse {
	return o.Payload
}

func (o *DescribeDbcOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeDbcResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeDbcDefault creates a DescribeDbcDefault with default headers values
func NewDescribeDbcDefault(code int) *DescribeDbcDefault {
	return &DescribeDbcDefault{
		_statusCode: code,
	}
}

/*
DescribeDbcDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeDbcDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe dbc default response has a 2xx status code
func (o *DescribeDbcDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe dbc default response has a 3xx status code
func (o *DescribeDbcDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe dbc default response has a 4xx status code
func (o *DescribeDbcDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe dbc default response has a 5xx status code
func (o *DescribeDbcDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe dbc default response a status code equal to that given
func (o *DescribeDbcDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe dbc default response
func (o *DescribeDbcDefault) Code() int {
	return o._statusCode
}

func (o *DescribeDbcDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeDbc][%d] describeDbc default %s", o._statusCode, payload)
}

func (o *DescribeDbcDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeDbc][%d] describeDbc default %s", o._statusCode, payload)
}

func (o *DescribeDbcDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeDbcDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
