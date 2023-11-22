// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// GenerateWorkloadAuthTokenReader is a Reader for the GenerateWorkloadAuthToken structure.
type GenerateWorkloadAuthTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateWorkloadAuthTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateWorkloadAuthTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGenerateWorkloadAuthTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGenerateWorkloadAuthTokenOK creates a GenerateWorkloadAuthTokenOK with default headers values
func NewGenerateWorkloadAuthTokenOK() *GenerateWorkloadAuthTokenOK {
	return &GenerateWorkloadAuthTokenOK{}
}

/*
GenerateWorkloadAuthTokenOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GenerateWorkloadAuthTokenOK struct {
	Payload *models.GenerateWorkloadAuthTokenResponse
}

// IsSuccess returns true when this generate workload auth token o k response has a 2xx status code
func (o *GenerateWorkloadAuthTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate workload auth token o k response has a 3xx status code
func (o *GenerateWorkloadAuthTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate workload auth token o k response has a 4xx status code
func (o *GenerateWorkloadAuthTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate workload auth token o k response has a 5xx status code
func (o *GenerateWorkloadAuthTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate workload auth token o k response a status code equal to that given
func (o *GenerateWorkloadAuthTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the generate workload auth token o k response
func (o *GenerateWorkloadAuthTokenOK) Code() int {
	return 200
}

func (o *GenerateWorkloadAuthTokenOK) Error() string {
	return fmt.Sprintf("[POST /iam/generateWorkloadAuthToken][%d] generateWorkloadAuthTokenOK  %+v", 200, o.Payload)
}

func (o *GenerateWorkloadAuthTokenOK) String() string {
	return fmt.Sprintf("[POST /iam/generateWorkloadAuthToken][%d] generateWorkloadAuthTokenOK  %+v", 200, o.Payload)
}

func (o *GenerateWorkloadAuthTokenOK) GetPayload() *models.GenerateWorkloadAuthTokenResponse {
	return o.Payload
}

func (o *GenerateWorkloadAuthTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenerateWorkloadAuthTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateWorkloadAuthTokenDefault creates a GenerateWorkloadAuthTokenDefault with default headers values
func NewGenerateWorkloadAuthTokenDefault(code int) *GenerateWorkloadAuthTokenDefault {
	return &GenerateWorkloadAuthTokenDefault{
		_statusCode: code,
	}
}

/*
GenerateWorkloadAuthTokenDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GenerateWorkloadAuthTokenDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this generate workload auth token default response has a 2xx status code
func (o *GenerateWorkloadAuthTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this generate workload auth token default response has a 3xx status code
func (o *GenerateWorkloadAuthTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this generate workload auth token default response has a 4xx status code
func (o *GenerateWorkloadAuthTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this generate workload auth token default response has a 5xx status code
func (o *GenerateWorkloadAuthTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this generate workload auth token default response a status code equal to that given
func (o *GenerateWorkloadAuthTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the generate workload auth token default response
func (o *GenerateWorkloadAuthTokenDefault) Code() int {
	return o._statusCode
}

func (o *GenerateWorkloadAuthTokenDefault) Error() string {
	return fmt.Sprintf("[POST /iam/generateWorkloadAuthToken][%d] generateWorkloadAuthToken default  %+v", o._statusCode, o.Payload)
}

func (o *GenerateWorkloadAuthTokenDefault) String() string {
	return fmt.Sprintf("[POST /iam/generateWorkloadAuthToken][%d] generateWorkloadAuthToken default  %+v", o._statusCode, o.Payload)
}

func (o *GenerateWorkloadAuthTokenDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GenerateWorkloadAuthTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}