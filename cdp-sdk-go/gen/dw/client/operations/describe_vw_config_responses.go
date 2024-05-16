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

// DescribeVwConfigReader is a Reader for the DescribeVwConfig structure.
type DescribeVwConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeVwConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeVwConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeVwConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeVwConfigOK creates a DescribeVwConfigOK with default headers values
func NewDescribeVwConfigOK() *DescribeVwConfigOK {
	return &DescribeVwConfigOK{}
}

/*
DescribeVwConfigOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeVwConfigOK struct {
	Payload *models.DescribeVwConfigResponse
}

// IsSuccess returns true when this describe vw config o k response has a 2xx status code
func (o *DescribeVwConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe vw config o k response has a 3xx status code
func (o *DescribeVwConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe vw config o k response has a 4xx status code
func (o *DescribeVwConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe vw config o k response has a 5xx status code
func (o *DescribeVwConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe vw config o k response a status code equal to that given
func (o *DescribeVwConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe vw config o k response
func (o *DescribeVwConfigOK) Code() int {
	return 200
}

func (o *DescribeVwConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeVwConfig][%d] describeVwConfigOK %s", 200, payload)
}

func (o *DescribeVwConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeVwConfig][%d] describeVwConfigOK %s", 200, payload)
}

func (o *DescribeVwConfigOK) GetPayload() *models.DescribeVwConfigResponse {
	return o.Payload
}

func (o *DescribeVwConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeVwConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeVwConfigDefault creates a DescribeVwConfigDefault with default headers values
func NewDescribeVwConfigDefault(code int) *DescribeVwConfigDefault {
	return &DescribeVwConfigDefault{
		_statusCode: code,
	}
}

/*
DescribeVwConfigDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeVwConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe vw config default response has a 2xx status code
func (o *DescribeVwConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe vw config default response has a 3xx status code
func (o *DescribeVwConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe vw config default response has a 4xx status code
func (o *DescribeVwConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe vw config default response has a 5xx status code
func (o *DescribeVwConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe vw config default response a status code equal to that given
func (o *DescribeVwConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe vw config default response
func (o *DescribeVwConfigDefault) Code() int {
	return o._statusCode
}

func (o *DescribeVwConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeVwConfig][%d] describeVwConfig default %s", o._statusCode, payload)
}

func (o *DescribeVwConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeVwConfig][%d] describeVwConfig default %s", o._statusCode, payload)
}

func (o *DescribeVwConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeVwConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
