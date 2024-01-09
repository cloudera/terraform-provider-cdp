// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// DescribeModelRegistryReader is a Reader for the DescribeModelRegistry structure.
type DescribeModelRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeModelRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeModelRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeModelRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeModelRegistryOK creates a DescribeModelRegistryOK with default headers values
func NewDescribeModelRegistryOK() *DescribeModelRegistryOK {
	return &DescribeModelRegistryOK{}
}

/*
DescribeModelRegistryOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeModelRegistryOK struct {
	Payload *models.DescribeModelRegistryResponse
}

// IsSuccess returns true when this describe model registry o k response has a 2xx status code
func (o *DescribeModelRegistryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe model registry o k response has a 3xx status code
func (o *DescribeModelRegistryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe model registry o k response has a 4xx status code
func (o *DescribeModelRegistryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe model registry o k response has a 5xx status code
func (o *DescribeModelRegistryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe model registry o k response a status code equal to that given
func (o *DescribeModelRegistryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe model registry o k response
func (o *DescribeModelRegistryOK) Code() int {
	return 200
}

func (o *DescribeModelRegistryOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/describeModelRegistry][%d] describeModelRegistryOK  %+v", 200, o.Payload)
}

func (o *DescribeModelRegistryOK) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/describeModelRegistry][%d] describeModelRegistryOK  %+v", 200, o.Payload)
}

func (o *DescribeModelRegistryOK) GetPayload() *models.DescribeModelRegistryResponse {
	return o.Payload
}

func (o *DescribeModelRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeModelRegistryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeModelRegistryDefault creates a DescribeModelRegistryDefault with default headers values
func NewDescribeModelRegistryDefault(code int) *DescribeModelRegistryDefault {
	return &DescribeModelRegistryDefault{
		_statusCode: code,
	}
}

/*
DescribeModelRegistryDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeModelRegistryDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe model registry default response has a 2xx status code
func (o *DescribeModelRegistryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe model registry default response has a 3xx status code
func (o *DescribeModelRegistryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe model registry default response has a 4xx status code
func (o *DescribeModelRegistryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe model registry default response has a 5xx status code
func (o *DescribeModelRegistryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe model registry default response a status code equal to that given
func (o *DescribeModelRegistryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe model registry default response
func (o *DescribeModelRegistryDefault) Code() int {
	return o._statusCode
}

func (o *DescribeModelRegistryDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/describeModelRegistry][%d] describeModelRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeModelRegistryDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/describeModelRegistry][%d] describeModelRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeModelRegistryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeModelRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
