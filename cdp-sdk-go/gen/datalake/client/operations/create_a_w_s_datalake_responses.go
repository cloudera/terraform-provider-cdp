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

// CreateAWSDatalakeReader is a Reader for the CreateAWSDatalake structure.
type CreateAWSDatalakeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAWSDatalakeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAWSDatalakeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateAWSDatalakeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAWSDatalakeOK creates a CreateAWSDatalakeOK with default headers values
func NewCreateAWSDatalakeOK() *CreateAWSDatalakeOK {
	return &CreateAWSDatalakeOK{}
}

/*
CreateAWSDatalakeOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CreateAWSDatalakeOK struct {
	Payload *models.CreateAWSDatalakeResponse
}

// IsSuccess returns true when this create a w s datalake o k response has a 2xx status code
func (o *CreateAWSDatalakeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create a w s datalake o k response has a 3xx status code
func (o *CreateAWSDatalakeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create a w s datalake o k response has a 4xx status code
func (o *CreateAWSDatalakeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create a w s datalake o k response has a 5xx status code
func (o *CreateAWSDatalakeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create a w s datalake o k response a status code equal to that given
func (o *CreateAWSDatalakeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create a w s datalake o k response
func (o *CreateAWSDatalakeOK) Code() int {
	return 200
}

func (o *CreateAWSDatalakeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/createAWSDatalake][%d] createAWSDatalakeOK %s", 200, payload)
}

func (o *CreateAWSDatalakeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/createAWSDatalake][%d] createAWSDatalakeOK %s", 200, payload)
}

func (o *CreateAWSDatalakeOK) GetPayload() *models.CreateAWSDatalakeResponse {
	return o.Payload
}

func (o *CreateAWSDatalakeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateAWSDatalakeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAWSDatalakeDefault creates a CreateAWSDatalakeDefault with default headers values
func NewCreateAWSDatalakeDefault(code int) *CreateAWSDatalakeDefault {
	return &CreateAWSDatalakeDefault{
		_statusCode: code,
	}
}

/*
CreateAWSDatalakeDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CreateAWSDatalakeDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create a w s datalake default response has a 2xx status code
func (o *CreateAWSDatalakeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create a w s datalake default response has a 3xx status code
func (o *CreateAWSDatalakeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create a w s datalake default response has a 4xx status code
func (o *CreateAWSDatalakeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create a w s datalake default response has a 5xx status code
func (o *CreateAWSDatalakeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create a w s datalake default response a status code equal to that given
func (o *CreateAWSDatalakeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create a w s datalake default response
func (o *CreateAWSDatalakeDefault) Code() int {
	return o._statusCode
}

func (o *CreateAWSDatalakeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/createAWSDatalake][%d] createAWSDatalake default %s", o._statusCode, payload)
}

func (o *CreateAWSDatalakeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/createAWSDatalake][%d] createAWSDatalake default %s", o._statusCode, payload)
}

func (o *CreateAWSDatalakeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateAWSDatalakeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
