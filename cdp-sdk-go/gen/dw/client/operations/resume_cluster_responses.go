// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// ResumeClusterReader is a Reader for the ResumeCluster structure.
type ResumeClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResumeClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResumeClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResumeClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResumeClusterOK creates a ResumeClusterOK with default headers values
func NewResumeClusterOK() *ResumeClusterOK {
	return &ResumeClusterOK{}
}

/*
ResumeClusterOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ResumeClusterOK struct {
	Payload *models.ResumeClusterResponse
}

// IsSuccess returns true when this resume cluster o k response has a 2xx status code
func (o *ResumeClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resume cluster o k response has a 3xx status code
func (o *ResumeClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resume cluster o k response has a 4xx status code
func (o *ResumeClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resume cluster o k response has a 5xx status code
func (o *ResumeClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resume cluster o k response a status code equal to that given
func (o *ResumeClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resume cluster o k response
func (o *ResumeClusterOK) Code() int {
	return 200
}

func (o *ResumeClusterOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/resumeCluster][%d] resumeClusterOK  %+v", 200, o.Payload)
}

func (o *ResumeClusterOK) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/resumeCluster][%d] resumeClusterOK  %+v", 200, o.Payload)
}

func (o *ResumeClusterOK) GetPayload() *models.ResumeClusterResponse {
	return o.Payload
}

func (o *ResumeClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResumeClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResumeClusterDefault creates a ResumeClusterDefault with default headers values
func NewResumeClusterDefault(code int) *ResumeClusterDefault {
	return &ResumeClusterDefault{
		_statusCode: code,
	}
}

/*
ResumeClusterDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ResumeClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this resume cluster default response has a 2xx status code
func (o *ResumeClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resume cluster default response has a 3xx status code
func (o *ResumeClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resume cluster default response has a 4xx status code
func (o *ResumeClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resume cluster default response has a 5xx status code
func (o *ResumeClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resume cluster default response a status code equal to that given
func (o *ResumeClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resume cluster default response
func (o *ResumeClusterDefault) Code() int {
	return o._statusCode
}

func (o *ResumeClusterDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/resumeCluster][%d] resumeCluster default  %+v", o._statusCode, o.Payload)
}

func (o *ResumeClusterDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/resumeCluster][%d] resumeCluster default  %+v", o._statusCode, o.Payload)
}

func (o *ResumeClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ResumeClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
