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

// CreatePrivateEnvironmentReader is a Reader for the CreatePrivateEnvironment structure.
type CreatePrivateEnvironmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePrivateEnvironmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePrivateEnvironmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreatePrivateEnvironmentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePrivateEnvironmentOK creates a CreatePrivateEnvironmentOK with default headers values
func NewCreatePrivateEnvironmentOK() *CreatePrivateEnvironmentOK {
	return &CreatePrivateEnvironmentOK{}
}

/*
CreatePrivateEnvironmentOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CreatePrivateEnvironmentOK struct {
	Payload *models.CreatePrivateEnvironmentResponse
}

// IsSuccess returns true when this create private environment o k response has a 2xx status code
func (o *CreatePrivateEnvironmentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create private environment o k response has a 3xx status code
func (o *CreatePrivateEnvironmentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create private environment o k response has a 4xx status code
func (o *CreatePrivateEnvironmentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create private environment o k response has a 5xx status code
func (o *CreatePrivateEnvironmentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create private environment o k response a status code equal to that given
func (o *CreatePrivateEnvironmentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create private environment o k response
func (o *CreatePrivateEnvironmentOK) Code() int {
	return 200
}

func (o *CreatePrivateEnvironmentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/createPrivateEnvironment][%d] createPrivateEnvironmentOK %s", 200, payload)
}

func (o *CreatePrivateEnvironmentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/createPrivateEnvironment][%d] createPrivateEnvironmentOK %s", 200, payload)
}

func (o *CreatePrivateEnvironmentOK) GetPayload() *models.CreatePrivateEnvironmentResponse {
	return o.Payload
}

func (o *CreatePrivateEnvironmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreatePrivateEnvironmentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePrivateEnvironmentDefault creates a CreatePrivateEnvironmentDefault with default headers values
func NewCreatePrivateEnvironmentDefault(code int) *CreatePrivateEnvironmentDefault {
	return &CreatePrivateEnvironmentDefault{
		_statusCode: code,
	}
}

/*
CreatePrivateEnvironmentDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CreatePrivateEnvironmentDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create private environment default response has a 2xx status code
func (o *CreatePrivateEnvironmentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create private environment default response has a 3xx status code
func (o *CreatePrivateEnvironmentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create private environment default response has a 4xx status code
func (o *CreatePrivateEnvironmentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create private environment default response has a 5xx status code
func (o *CreatePrivateEnvironmentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create private environment default response a status code equal to that given
func (o *CreatePrivateEnvironmentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create private environment default response
func (o *CreatePrivateEnvironmentDefault) Code() int {
	return o._statusCode
}

func (o *CreatePrivateEnvironmentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/createPrivateEnvironment][%d] createPrivateEnvironment default %s", o._statusCode, payload)
}

func (o *CreatePrivateEnvironmentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/createPrivateEnvironment][%d] createPrivateEnvironment default %s", o._statusCode, payload)
}

func (o *CreatePrivateEnvironmentDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePrivateEnvironmentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
