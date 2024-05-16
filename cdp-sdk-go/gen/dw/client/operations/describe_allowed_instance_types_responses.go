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

// DescribeAllowedInstanceTypesReader is a Reader for the DescribeAllowedInstanceTypes structure.
type DescribeAllowedInstanceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeAllowedInstanceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeAllowedInstanceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeAllowedInstanceTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeAllowedInstanceTypesOK creates a DescribeAllowedInstanceTypesOK with default headers values
func NewDescribeAllowedInstanceTypesOK() *DescribeAllowedInstanceTypesOK {
	return &DescribeAllowedInstanceTypesOK{}
}

/*
DescribeAllowedInstanceTypesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeAllowedInstanceTypesOK struct {
	Payload *models.DescribeAllowedInstanceTypesResponse
}

// IsSuccess returns true when this describe allowed instance types o k response has a 2xx status code
func (o *DescribeAllowedInstanceTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe allowed instance types o k response has a 3xx status code
func (o *DescribeAllowedInstanceTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe allowed instance types o k response has a 4xx status code
func (o *DescribeAllowedInstanceTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe allowed instance types o k response has a 5xx status code
func (o *DescribeAllowedInstanceTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe allowed instance types o k response a status code equal to that given
func (o *DescribeAllowedInstanceTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe allowed instance types o k response
func (o *DescribeAllowedInstanceTypesOK) Code() int {
	return 200
}

func (o *DescribeAllowedInstanceTypesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeAllowedInstanceTypes][%d] describeAllowedInstanceTypesOK %s", 200, payload)
}

func (o *DescribeAllowedInstanceTypesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeAllowedInstanceTypes][%d] describeAllowedInstanceTypesOK %s", 200, payload)
}

func (o *DescribeAllowedInstanceTypesOK) GetPayload() *models.DescribeAllowedInstanceTypesResponse {
	return o.Payload
}

func (o *DescribeAllowedInstanceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeAllowedInstanceTypesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeAllowedInstanceTypesDefault creates a DescribeAllowedInstanceTypesDefault with default headers values
func NewDescribeAllowedInstanceTypesDefault(code int) *DescribeAllowedInstanceTypesDefault {
	return &DescribeAllowedInstanceTypesDefault{
		_statusCode: code,
	}
}

/*
DescribeAllowedInstanceTypesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeAllowedInstanceTypesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe allowed instance types default response has a 2xx status code
func (o *DescribeAllowedInstanceTypesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe allowed instance types default response has a 3xx status code
func (o *DescribeAllowedInstanceTypesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe allowed instance types default response has a 4xx status code
func (o *DescribeAllowedInstanceTypesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe allowed instance types default response has a 5xx status code
func (o *DescribeAllowedInstanceTypesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe allowed instance types default response a status code equal to that given
func (o *DescribeAllowedInstanceTypesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe allowed instance types default response
func (o *DescribeAllowedInstanceTypesDefault) Code() int {
	return o._statusCode
}

func (o *DescribeAllowedInstanceTypesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeAllowedInstanceTypes][%d] describeAllowedInstanceTypes default %s", o._statusCode, payload)
}

func (o *DescribeAllowedInstanceTypesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeAllowedInstanceTypes][%d] describeAllowedInstanceTypes default %s", o._statusCode, payload)
}

func (o *DescribeAllowedInstanceTypesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeAllowedInstanceTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
