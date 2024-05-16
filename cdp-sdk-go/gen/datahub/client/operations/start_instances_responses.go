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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// StartInstancesReader is a Reader for the StartInstances structure.
type StartInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartInstancesOK creates a StartInstancesOK with default headers values
func NewStartInstancesOK() *StartInstancesOK {
	return &StartInstancesOK{}
}

/*
StartInstancesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type StartInstancesOK struct {
	Payload models.StartInstancesResponse
}

// IsSuccess returns true when this start instances o k response has a 2xx status code
func (o *StartInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start instances o k response has a 3xx status code
func (o *StartInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start instances o k response has a 4xx status code
func (o *StartInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start instances o k response has a 5xx status code
func (o *StartInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start instances o k response a status code equal to that given
func (o *StartInstancesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start instances o k response
func (o *StartInstancesOK) Code() int {
	return 200
}

func (o *StartInstancesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/startInstances][%d] startInstancesOK %s", 200, payload)
}

func (o *StartInstancesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/startInstances][%d] startInstancesOK %s", 200, payload)
}

func (o *StartInstancesOK) GetPayload() models.StartInstancesResponse {
	return o.Payload
}

func (o *StartInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartInstancesDefault creates a StartInstancesDefault with default headers values
func NewStartInstancesDefault(code int) *StartInstancesDefault {
	return &StartInstancesDefault{
		_statusCode: code,
	}
}

/*
StartInstancesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type StartInstancesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this start instances default response has a 2xx status code
func (o *StartInstancesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this start instances default response has a 3xx status code
func (o *StartInstancesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this start instances default response has a 4xx status code
func (o *StartInstancesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this start instances default response has a 5xx status code
func (o *StartInstancesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this start instances default response a status code equal to that given
func (o *StartInstancesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the start instances default response
func (o *StartInstancesDefault) Code() int {
	return o._statusCode
}

func (o *StartInstancesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/startInstances][%d] startInstances default %s", o._statusCode, payload)
}

func (o *StartInstancesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/startInstances][%d] startInstances default %s", o._statusCode, payload)
}

func (o *StartInstancesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
