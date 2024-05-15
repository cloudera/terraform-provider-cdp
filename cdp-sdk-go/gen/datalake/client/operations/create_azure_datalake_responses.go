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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// CreateAzureDatalakeReader is a Reader for the CreateAzureDatalake structure.
type CreateAzureDatalakeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAzureDatalakeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAzureDatalakeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateAzureDatalakeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAzureDatalakeOK creates a CreateAzureDatalakeOK with default headers values
func NewCreateAzureDatalakeOK() *CreateAzureDatalakeOK {
	return &CreateAzureDatalakeOK{}
}

/*
CreateAzureDatalakeOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CreateAzureDatalakeOK struct {
	Payload *models.CreateAzureDatalakeResponse
}

// IsSuccess returns true when this create azure datalake o k response has a 2xx status code
func (o *CreateAzureDatalakeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create azure datalake o k response has a 3xx status code
func (o *CreateAzureDatalakeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create azure datalake o k response has a 4xx status code
func (o *CreateAzureDatalakeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create azure datalake o k response has a 5xx status code
func (o *CreateAzureDatalakeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create azure datalake o k response a status code equal to that given
func (o *CreateAzureDatalakeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create azure datalake o k response
func (o *CreateAzureDatalakeOK) Code() int {
	return 200
}

func (o *CreateAzureDatalakeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/createAzureDatalake][%d] createAzureDatalakeOK %s", 200, payload)
}

func (o *CreateAzureDatalakeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/createAzureDatalake][%d] createAzureDatalakeOK %s", 200, payload)
}

func (o *CreateAzureDatalakeOK) GetPayload() *models.CreateAzureDatalakeResponse {
	return o.Payload
}

func (o *CreateAzureDatalakeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateAzureDatalakeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAzureDatalakeDefault creates a CreateAzureDatalakeDefault with default headers values
func NewCreateAzureDatalakeDefault(code int) *CreateAzureDatalakeDefault {
	return &CreateAzureDatalakeDefault{
		_statusCode: code,
	}
}

/*
CreateAzureDatalakeDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CreateAzureDatalakeDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create azure datalake default response has a 2xx status code
func (o *CreateAzureDatalakeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create azure datalake default response has a 3xx status code
func (o *CreateAzureDatalakeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create azure datalake default response has a 4xx status code
func (o *CreateAzureDatalakeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create azure datalake default response has a 5xx status code
func (o *CreateAzureDatalakeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create azure datalake default response a status code equal to that given
func (o *CreateAzureDatalakeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create azure datalake default response
func (o *CreateAzureDatalakeDefault) Code() int {
	return o._statusCode
}

func (o *CreateAzureDatalakeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/createAzureDatalake][%d] createAzureDatalake default %s", o._statusCode, payload)
}

func (o *CreateAzureDatalakeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/createAzureDatalake][%d] createAzureDatalake default %s", o._statusCode, payload)
}

func (o *CreateAzureDatalakeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAzureDatalakeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
