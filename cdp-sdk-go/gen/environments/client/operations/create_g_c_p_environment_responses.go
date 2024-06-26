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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// CreateGCPEnvironmentReader is a Reader for the CreateGCPEnvironment structure.
type CreateGCPEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGCPEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateGCPEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateGCPEnvironmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateGCPEnvironmentOK creates a CreateGCPEnvironmentOK with default headers values
func NewCreateGCPEnvironmentOK() *CreateGCPEnvironmentOK {
	return &CreateGCPEnvironmentOK{}
}

/*
CreateGCPEnvironmentOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CreateGCPEnvironmentOK struct {
	Payload *models.CreateGCPEnvironmentResponse
}

// IsSuccess returns true when this create g c p environment o k response has a 2xx status code
func (o *CreateGCPEnvironmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create g c p environment o k response has a 3xx status code
func (o *CreateGCPEnvironmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create g c p environment o k response has a 4xx status code
func (o *CreateGCPEnvironmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create g c p environment o k response has a 5xx status code
func (o *CreateGCPEnvironmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create g c p environment o k response a status code equal to that given
func (o *CreateGCPEnvironmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create g c p environment o k response
func (o *CreateGCPEnvironmentOK) Code() int {
	return 200
}

func (o *CreateGCPEnvironmentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/createGCPEnvironment][%d] createGCPEnvironmentOK %s", 200, payload)
}

func (o *CreateGCPEnvironmentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/createGCPEnvironment][%d] createGCPEnvironmentOK %s", 200, payload)
}

func (o *CreateGCPEnvironmentOK) GetPayload() *models.CreateGCPEnvironmentResponse {
	return o.Payload
}

func (o *CreateGCPEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateGCPEnvironmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGCPEnvironmentDefault creates a CreateGCPEnvironmentDefault with default headers values
func NewCreateGCPEnvironmentDefault(code int) *CreateGCPEnvironmentDefault {
	return &CreateGCPEnvironmentDefault{
		_statusCode: code,
	}
}

/*
CreateGCPEnvironmentDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CreateGCPEnvironmentDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create g c p environment default response has a 2xx status code
func (o *CreateGCPEnvironmentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create g c p environment default response has a 3xx status code
func (o *CreateGCPEnvironmentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create g c p environment default response has a 4xx status code
func (o *CreateGCPEnvironmentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create g c p environment default response has a 5xx status code
func (o *CreateGCPEnvironmentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create g c p environment default response a status code equal to that given
func (o *CreateGCPEnvironmentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create g c p environment default response
func (o *CreateGCPEnvironmentDefault) Code() int {
	return o._statusCode
}

func (o *CreateGCPEnvironmentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/createGCPEnvironment][%d] createGCPEnvironment default %s", o._statusCode, payload)
}

func (o *CreateGCPEnvironmentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/createGCPEnvironment][%d] createGCPEnvironment default %s", o._statusCode, payload)
}

func (o *CreateGCPEnvironmentDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGCPEnvironmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
