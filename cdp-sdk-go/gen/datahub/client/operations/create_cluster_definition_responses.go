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

// CreateClusterDefinitionReader is a Reader for the CreateClusterDefinition structure.
type CreateClusterDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateClusterDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateClusterDefinitionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClusterDefinitionOK creates a CreateClusterDefinitionOK with default headers values
func NewCreateClusterDefinitionOK() *CreateClusterDefinitionOK {
	return &CreateClusterDefinitionOK{}
}

/*
CreateClusterDefinitionOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CreateClusterDefinitionOK struct {
	Payload *models.CreateClusterDefinitionResponse
}

// IsSuccess returns true when this create cluster definition o k response has a 2xx status code
func (o *CreateClusterDefinitionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cluster definition o k response has a 3xx status code
func (o *CreateClusterDefinitionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster definition o k response has a 4xx status code
func (o *CreateClusterDefinitionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cluster definition o k response has a 5xx status code
func (o *CreateClusterDefinitionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster definition o k response a status code equal to that given
func (o *CreateClusterDefinitionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create cluster definition o k response
func (o *CreateClusterDefinitionOK) Code() int {
	return 200
}

func (o *CreateClusterDefinitionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/createClusterDefinition][%d] createClusterDefinitionOK %s", 200, payload)
}

func (o *CreateClusterDefinitionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/createClusterDefinition][%d] createClusterDefinitionOK %s", 200, payload)
}

func (o *CreateClusterDefinitionOK) GetPayload() *models.CreateClusterDefinitionResponse {
	return o.Payload
}

func (o *CreateClusterDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateClusterDefinitionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterDefinitionDefault creates a CreateClusterDefinitionDefault with default headers values
func NewCreateClusterDefinitionDefault(code int) *CreateClusterDefinitionDefault {
	return &CreateClusterDefinitionDefault{
		_statusCode: code,
	}
}

/*
CreateClusterDefinitionDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CreateClusterDefinitionDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create cluster definition default response has a 2xx status code
func (o *CreateClusterDefinitionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create cluster definition default response has a 3xx status code
func (o *CreateClusterDefinitionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create cluster definition default response has a 4xx status code
func (o *CreateClusterDefinitionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create cluster definition default response has a 5xx status code
func (o *CreateClusterDefinitionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create cluster definition default response a status code equal to that given
func (o *CreateClusterDefinitionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create cluster definition default response
func (o *CreateClusterDefinitionDefault) Code() int {
	return o._statusCode
}

func (o *CreateClusterDefinitionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/createClusterDefinition][%d] createClusterDefinition default %s", o._statusCode, payload)
}

func (o *CreateClusterDefinitionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/createClusterDefinition][%d] createClusterDefinition default %s", o._statusCode, payload)
}

func (o *CreateClusterDefinitionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateClusterDefinitionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
