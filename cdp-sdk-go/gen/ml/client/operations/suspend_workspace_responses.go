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

// SuspendWorkspaceReader is a Reader for the SuspendWorkspace structure.
type SuspendWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuspendWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuspendWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSuspendWorkspaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSuspendWorkspaceOK creates a SuspendWorkspaceOK with default headers values
func NewSuspendWorkspaceOK() *SuspendWorkspaceOK {
	return &SuspendWorkspaceOK{}
}

/*
SuspendWorkspaceOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SuspendWorkspaceOK struct {
	Payload models.SuspendWorkspaceResponse
}

// IsSuccess returns true when this suspend workspace o k response has a 2xx status code
func (o *SuspendWorkspaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this suspend workspace o k response has a 3xx status code
func (o *SuspendWorkspaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this suspend workspace o k response has a 4xx status code
func (o *SuspendWorkspaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this suspend workspace o k response has a 5xx status code
func (o *SuspendWorkspaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this suspend workspace o k response a status code equal to that given
func (o *SuspendWorkspaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the suspend workspace o k response
func (o *SuspendWorkspaceOK) Code() int {
	return 200
}

func (o *SuspendWorkspaceOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/suspendWorkspace][%d] suspendWorkspaceOK  %+v", 200, o.Payload)
}

func (o *SuspendWorkspaceOK) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/suspendWorkspace][%d] suspendWorkspaceOK  %+v", 200, o.Payload)
}

func (o *SuspendWorkspaceOK) GetPayload() models.SuspendWorkspaceResponse {
	return o.Payload
}

func (o *SuspendWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuspendWorkspaceDefault creates a SuspendWorkspaceDefault with default headers values
func NewSuspendWorkspaceDefault(code int) *SuspendWorkspaceDefault {
	return &SuspendWorkspaceDefault{
		_statusCode: code,
	}
}

/*
SuspendWorkspaceDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SuspendWorkspaceDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this suspend workspace default response has a 2xx status code
func (o *SuspendWorkspaceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this suspend workspace default response has a 3xx status code
func (o *SuspendWorkspaceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this suspend workspace default response has a 4xx status code
func (o *SuspendWorkspaceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this suspend workspace default response has a 5xx status code
func (o *SuspendWorkspaceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this suspend workspace default response a status code equal to that given
func (o *SuspendWorkspaceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the suspend workspace default response
func (o *SuspendWorkspaceDefault) Code() int {
	return o._statusCode
}

func (o *SuspendWorkspaceDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/suspendWorkspace][%d] suspendWorkspace default  %+v", o._statusCode, o.Payload)
}

func (o *SuspendWorkspaceDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/suspendWorkspace][%d] suspendWorkspace default  %+v", o._statusCode, o.Payload)
}

func (o *SuspendWorkspaceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SuspendWorkspaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}