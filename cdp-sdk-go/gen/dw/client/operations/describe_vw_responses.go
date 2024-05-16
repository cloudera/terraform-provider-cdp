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

// DescribeVwReader is a Reader for the DescribeVw structure.
type DescribeVwReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeVwReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeVwOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeVwDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeVwOK creates a DescribeVwOK with default headers values
func NewDescribeVwOK() *DescribeVwOK {
	return &DescribeVwOK{}
}

/*
DescribeVwOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeVwOK struct {
	Payload *models.DescribeVwResponse
}

// IsSuccess returns true when this describe vw o k response has a 2xx status code
func (o *DescribeVwOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe vw o k response has a 3xx status code
func (o *DescribeVwOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe vw o k response has a 4xx status code
func (o *DescribeVwOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe vw o k response has a 5xx status code
func (o *DescribeVwOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe vw o k response a status code equal to that given
func (o *DescribeVwOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe vw o k response
func (o *DescribeVwOK) Code() int {
	return 200
}

func (o *DescribeVwOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeVw][%d] describeVwOK %s", 200, payload)
}

func (o *DescribeVwOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeVw][%d] describeVwOK %s", 200, payload)
}

func (o *DescribeVwOK) GetPayload() *models.DescribeVwResponse {
	return o.Payload
}

func (o *DescribeVwOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeVwResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeVwDefault creates a DescribeVwDefault with default headers values
func NewDescribeVwDefault(code int) *DescribeVwDefault {
	return &DescribeVwDefault{
		_statusCode: code,
	}
}

/*
DescribeVwDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeVwDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe vw default response has a 2xx status code
func (o *DescribeVwDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe vw default response has a 3xx status code
func (o *DescribeVwDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe vw default response has a 4xx status code
func (o *DescribeVwDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe vw default response has a 5xx status code
func (o *DescribeVwDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe vw default response a status code equal to that given
func (o *DescribeVwDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe vw default response
func (o *DescribeVwDefault) Code() int {
	return o._statusCode
}

func (o *DescribeVwDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeVw][%d] describeVw default %s", o._statusCode, payload)
}

func (o *DescribeVwDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeVw][%d] describeVw default %s", o._statusCode, payload)
}

func (o *DescribeVwDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeVwDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
