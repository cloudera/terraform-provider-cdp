// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// CreateCustomConfigurationsReader is a Reader for the CreateCustomConfigurations structure.
type CreateCustomConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCustomConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateCustomConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateCustomConfigurationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCustomConfigurationsOK creates a CreateCustomConfigurationsOK with default headers values
func NewCreateCustomConfigurationsOK() *CreateCustomConfigurationsOK {
	return &CreateCustomConfigurationsOK{}
}

/*
CreateCustomConfigurationsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CreateCustomConfigurationsOK struct {
	Payload *models.CreateCustomConfigurationsResponse
}

// IsSuccess returns true when this create custom configurations o k response has a 2xx status code
func (o *CreateCustomConfigurationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create custom configurations o k response has a 3xx status code
func (o *CreateCustomConfigurationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create custom configurations o k response has a 4xx status code
func (o *CreateCustomConfigurationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create custom configurations o k response has a 5xx status code
func (o *CreateCustomConfigurationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create custom configurations o k response a status code equal to that given
func (o *CreateCustomConfigurationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create custom configurations o k response
func (o *CreateCustomConfigurationsOK) Code() int {
	return 200
}

func (o *CreateCustomConfigurationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/createCustomConfigurations][%d] createCustomConfigurationsOK  %+v", 200, o.Payload)
}

func (o *CreateCustomConfigurationsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/createCustomConfigurations][%d] createCustomConfigurationsOK  %+v", 200, o.Payload)
}

func (o *CreateCustomConfigurationsOK) GetPayload() *models.CreateCustomConfigurationsResponse {
	return o.Payload
}

func (o *CreateCustomConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateCustomConfigurationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomConfigurationsDefault creates a CreateCustomConfigurationsDefault with default headers values
func NewCreateCustomConfigurationsDefault(code int) *CreateCustomConfigurationsDefault {
	return &CreateCustomConfigurationsDefault{
		_statusCode: code,
	}
}

/*
CreateCustomConfigurationsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CreateCustomConfigurationsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create custom configurations default response has a 2xx status code
func (o *CreateCustomConfigurationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create custom configurations default response has a 3xx status code
func (o *CreateCustomConfigurationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create custom configurations default response has a 4xx status code
func (o *CreateCustomConfigurationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create custom configurations default response has a 5xx status code
func (o *CreateCustomConfigurationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create custom configurations default response a status code equal to that given
func (o *CreateCustomConfigurationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create custom configurations default response
func (o *CreateCustomConfigurationsDefault) Code() int {
	return o._statusCode
}

func (o *CreateCustomConfigurationsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/createCustomConfigurations][%d] createCustomConfigurations default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCustomConfigurationsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/createCustomConfigurations][%d] createCustomConfigurations default  %+v", o._statusCode, o.Payload)
}

func (o *CreateCustomConfigurationsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateCustomConfigurationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
