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

// SuspendClusterReader is a Reader for the SuspendCluster structure.
type SuspendClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuspendClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuspendClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSuspendClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSuspendClusterOK creates a SuspendClusterOK with default headers values
func NewSuspendClusterOK() *SuspendClusterOK {
	return &SuspendClusterOK{}
}

/*
SuspendClusterOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SuspendClusterOK struct {
	Payload *models.SuspendClusterResponse
}

// IsSuccess returns true when this suspend cluster o k response has a 2xx status code
func (o *SuspendClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this suspend cluster o k response has a 3xx status code
func (o *SuspendClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this suspend cluster o k response has a 4xx status code
func (o *SuspendClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this suspend cluster o k response has a 5xx status code
func (o *SuspendClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this suspend cluster o k response a status code equal to that given
func (o *SuspendClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the suspend cluster o k response
func (o *SuspendClusterOK) Code() int {
	return 200
}

func (o *SuspendClusterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/suspendCluster][%d] suspendClusterOK %s", 200, payload)
}

func (o *SuspendClusterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/suspendCluster][%d] suspendClusterOK %s", 200, payload)
}

func (o *SuspendClusterOK) GetPayload() *models.SuspendClusterResponse {
	return o.Payload
}

func (o *SuspendClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuspendClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuspendClusterDefault creates a SuspendClusterDefault with default headers values
func NewSuspendClusterDefault(code int) *SuspendClusterDefault {
	return &SuspendClusterDefault{
		_statusCode: code,
	}
}

/*
SuspendClusterDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SuspendClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this suspend cluster default response has a 2xx status code
func (o *SuspendClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this suspend cluster default response has a 3xx status code
func (o *SuspendClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this suspend cluster default response has a 4xx status code
func (o *SuspendClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this suspend cluster default response has a 5xx status code
func (o *SuspendClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this suspend cluster default response a status code equal to that given
func (o *SuspendClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the suspend cluster default response
func (o *SuspendClusterDefault) Code() int {
	return o._statusCode
}

func (o *SuspendClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/suspendCluster][%d] suspendCluster default %s", o._statusCode, payload)
}

func (o *SuspendClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/suspendCluster][%d] suspendCluster default %s", o._statusCode, payload)
}

func (o *SuspendClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SuspendClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
