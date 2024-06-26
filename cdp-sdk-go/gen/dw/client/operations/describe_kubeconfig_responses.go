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

// DescribeKubeconfigReader is a Reader for the DescribeKubeconfig structure.
type DescribeKubeconfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeKubeconfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeKubeconfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeKubeconfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeKubeconfigOK creates a DescribeKubeconfigOK with default headers values
func NewDescribeKubeconfigOK() *DescribeKubeconfigOK {
	return &DescribeKubeconfigOK{}
}

/*
DescribeKubeconfigOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeKubeconfigOK struct {
	Payload *models.DescribeKubeconfigResponse
}

// IsSuccess returns true when this describe kubeconfig o k response has a 2xx status code
func (o *DescribeKubeconfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe kubeconfig o k response has a 3xx status code
func (o *DescribeKubeconfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe kubeconfig o k response has a 4xx status code
func (o *DescribeKubeconfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe kubeconfig o k response has a 5xx status code
func (o *DescribeKubeconfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe kubeconfig o k response a status code equal to that given
func (o *DescribeKubeconfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe kubeconfig o k response
func (o *DescribeKubeconfigOK) Code() int {
	return 200
}

func (o *DescribeKubeconfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeKubeconfig][%d] describeKubeconfigOK %s", 200, payload)
}

func (o *DescribeKubeconfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeKubeconfig][%d] describeKubeconfigOK %s", 200, payload)
}

func (o *DescribeKubeconfigOK) GetPayload() *models.DescribeKubeconfigResponse {
	return o.Payload
}

func (o *DescribeKubeconfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeKubeconfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeKubeconfigDefault creates a DescribeKubeconfigDefault with default headers values
func NewDescribeKubeconfigDefault(code int) *DescribeKubeconfigDefault {
	return &DescribeKubeconfigDefault{
		_statusCode: code,
	}
}

/*
DescribeKubeconfigDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeKubeconfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe kubeconfig default response has a 2xx status code
func (o *DescribeKubeconfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe kubeconfig default response has a 3xx status code
func (o *DescribeKubeconfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe kubeconfig default response has a 4xx status code
func (o *DescribeKubeconfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe kubeconfig default response has a 5xx status code
func (o *DescribeKubeconfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe kubeconfig default response a status code equal to that given
func (o *DescribeKubeconfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe kubeconfig default response
func (o *DescribeKubeconfigDefault) Code() int {
	return o._statusCode
}

func (o *DescribeKubeconfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeKubeconfig][%d] describeKubeconfig default %s", o._statusCode, payload)
}

func (o *DescribeKubeconfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeKubeconfig][%d] describeKubeconfig default %s", o._statusCode, payload)
}

func (o *DescribeKubeconfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeKubeconfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
