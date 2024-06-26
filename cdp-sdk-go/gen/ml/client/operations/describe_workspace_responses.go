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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// DescribeWorkspaceReader is a Reader for the DescribeWorkspace structure.
type DescribeWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeWorkspaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeWorkspaceOK creates a DescribeWorkspaceOK with default headers values
func NewDescribeWorkspaceOK() *DescribeWorkspaceOK {
	return &DescribeWorkspaceOK{}
}

/*
DescribeWorkspaceOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeWorkspaceOK struct {
	Payload *models.DescribeWorkspaceResponse
}

// IsSuccess returns true when this describe workspace o k response has a 2xx status code
func (o *DescribeWorkspaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe workspace o k response has a 3xx status code
func (o *DescribeWorkspaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe workspace o k response has a 4xx status code
func (o *DescribeWorkspaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe workspace o k response has a 5xx status code
func (o *DescribeWorkspaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe workspace o k response a status code equal to that given
func (o *DescribeWorkspaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe workspace o k response
func (o *DescribeWorkspaceOK) Code() int {
	return 200
}

func (o *DescribeWorkspaceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/describeWorkspace][%d] describeWorkspaceOK %s", 200, payload)
}

func (o *DescribeWorkspaceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/describeWorkspace][%d] describeWorkspaceOK %s", 200, payload)
}

func (o *DescribeWorkspaceOK) GetPayload() *models.DescribeWorkspaceResponse {
	return o.Payload
}

func (o *DescribeWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeWorkspaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeWorkspaceDefault creates a DescribeWorkspaceDefault with default headers values
func NewDescribeWorkspaceDefault(code int) *DescribeWorkspaceDefault {
	return &DescribeWorkspaceDefault{
		_statusCode: code,
	}
}

/*
DescribeWorkspaceDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeWorkspaceDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe workspace default response has a 2xx status code
func (o *DescribeWorkspaceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe workspace default response has a 3xx status code
func (o *DescribeWorkspaceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe workspace default response has a 4xx status code
func (o *DescribeWorkspaceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe workspace default response has a 5xx status code
func (o *DescribeWorkspaceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe workspace default response a status code equal to that given
func (o *DescribeWorkspaceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe workspace default response
func (o *DescribeWorkspaceDefault) Code() int {
	return o._statusCode
}

func (o *DescribeWorkspaceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/describeWorkspace][%d] describeWorkspace default %s", o._statusCode, payload)
}

func (o *DescribeWorkspaceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/describeWorkspace][%d] describeWorkspace default %s", o._statusCode, payload)
}

func (o *DescribeWorkspaceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeWorkspaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
