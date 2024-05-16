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

// DescribeConfigReader is a Reader for the DescribeConfig structure.
type DescribeConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeConfigOK creates a DescribeConfigOK with default headers values
func NewDescribeConfigOK() *DescribeConfigOK {
	return &DescribeConfigOK{}
}

/*
DescribeConfigOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeConfigOK struct {
	Payload *models.DescribeConfigResponse
}

// IsSuccess returns true when this describe config o k response has a 2xx status code
func (o *DescribeConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe config o k response has a 3xx status code
func (o *DescribeConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe config o k response has a 4xx status code
func (o *DescribeConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe config o k response has a 5xx status code
func (o *DescribeConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe config o k response a status code equal to that given
func (o *DescribeConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe config o k response
func (o *DescribeConfigOK) Code() int {
	return 200
}

func (o *DescribeConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeConfig][%d] describeConfigOK %s", 200, payload)
}

func (o *DescribeConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeConfig][%d] describeConfigOK %s", 200, payload)
}

func (o *DescribeConfigOK) GetPayload() *models.DescribeConfigResponse {
	return o.Payload
}

func (o *DescribeConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeConfigDefault creates a DescribeConfigDefault with default headers values
func NewDescribeConfigDefault(code int) *DescribeConfigDefault {
	return &DescribeConfigDefault{
		_statusCode: code,
	}
}

/*
DescribeConfigDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe config default response has a 2xx status code
func (o *DescribeConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe config default response has a 3xx status code
func (o *DescribeConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe config default response has a 4xx status code
func (o *DescribeConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe config default response has a 5xx status code
func (o *DescribeConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe config default response a status code equal to that given
func (o *DescribeConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe config default response
func (o *DescribeConfigDefault) Code() int {
	return o._statusCode
}

func (o *DescribeConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeConfig][%d] describeConfig default %s", o._statusCode, payload)
}

func (o *DescribeConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeConfig][%d] describeConfig default %s", o._statusCode, payload)
}

func (o *DescribeConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
