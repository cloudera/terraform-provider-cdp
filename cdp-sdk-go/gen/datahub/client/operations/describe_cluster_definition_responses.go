// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// DescribeClusterDefinitionReader is a Reader for the DescribeClusterDefinition structure.
type DescribeClusterDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeClusterDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeClusterDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeClusterDefinitionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeClusterDefinitionOK creates a DescribeClusterDefinitionOK with default headers values
func NewDescribeClusterDefinitionOK() *DescribeClusterDefinitionOK {
	return &DescribeClusterDefinitionOK{}
}

/*
DescribeClusterDefinitionOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeClusterDefinitionOK struct {
	Payload *models.DescribeClusterDefinitionResponse
}

// IsSuccess returns true when this describe cluster definition o k response has a 2xx status code
func (o *DescribeClusterDefinitionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe cluster definition o k response has a 3xx status code
func (o *DescribeClusterDefinitionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe cluster definition o k response has a 4xx status code
func (o *DescribeClusterDefinitionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe cluster definition o k response has a 5xx status code
func (o *DescribeClusterDefinitionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe cluster definition o k response a status code equal to that given
func (o *DescribeClusterDefinitionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe cluster definition o k response
func (o *DescribeClusterDefinitionOK) Code() int {
	return 200
}

func (o *DescribeClusterDefinitionOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/describeClusterDefinition][%d] describeClusterDefinitionOK  %+v", 200, o.Payload)
}

func (o *DescribeClusterDefinitionOK) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/describeClusterDefinition][%d] describeClusterDefinitionOK  %+v", 200, o.Payload)
}

func (o *DescribeClusterDefinitionOK) GetPayload() *models.DescribeClusterDefinitionResponse {
	return o.Payload
}

func (o *DescribeClusterDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeClusterDefinitionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeClusterDefinitionDefault creates a DescribeClusterDefinitionDefault with default headers values
func NewDescribeClusterDefinitionDefault(code int) *DescribeClusterDefinitionDefault {
	return &DescribeClusterDefinitionDefault{
		_statusCode: code,
	}
}

/*
DescribeClusterDefinitionDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeClusterDefinitionDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe cluster definition default response has a 2xx status code
func (o *DescribeClusterDefinitionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe cluster definition default response has a 3xx status code
func (o *DescribeClusterDefinitionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe cluster definition default response has a 4xx status code
func (o *DescribeClusterDefinitionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe cluster definition default response has a 5xx status code
func (o *DescribeClusterDefinitionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe cluster definition default response a status code equal to that given
func (o *DescribeClusterDefinitionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe cluster definition default response
func (o *DescribeClusterDefinitionDefault) Code() int {
	return o._statusCode
}

func (o *DescribeClusterDefinitionDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/describeClusterDefinition][%d] describeClusterDefinition default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeClusterDefinitionDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/describeClusterDefinition][%d] describeClusterDefinition default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeClusterDefinitionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeClusterDefinitionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}