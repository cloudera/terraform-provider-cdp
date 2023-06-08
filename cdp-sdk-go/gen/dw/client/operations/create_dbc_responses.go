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

// CreateDbcReader is a Reader for the CreateDbc structure.
type CreateDbcReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDbcReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDbcOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateDbcDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDbcOK creates a CreateDbcOK with default headers values
func NewCreateDbcOK() *CreateDbcOK {
	return &CreateDbcOK{}
}

/*
CreateDbcOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CreateDbcOK struct {
	Payload *models.CreateDbcResponse
}

// IsSuccess returns true when this create dbc o k response has a 2xx status code
func (o *CreateDbcOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create dbc o k response has a 3xx status code
func (o *CreateDbcOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create dbc o k response has a 4xx status code
func (o *CreateDbcOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create dbc o k response has a 5xx status code
func (o *CreateDbcOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create dbc o k response a status code equal to that given
func (o *CreateDbcOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create dbc o k response
func (o *CreateDbcOK) Code() int {
	return 200
}

func (o *CreateDbcOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/createDbc][%d] createDbcOK  %+v", 200, o.Payload)
}

func (o *CreateDbcOK) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/createDbc][%d] createDbcOK  %+v", 200, o.Payload)
}

func (o *CreateDbcOK) GetPayload() *models.CreateDbcResponse {
	return o.Payload
}

func (o *CreateDbcOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDbcResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDbcDefault creates a CreateDbcDefault with default headers values
func NewCreateDbcDefault(code int) *CreateDbcDefault {
	return &CreateDbcDefault{
		_statusCode: code,
	}
}

/*
CreateDbcDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CreateDbcDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create dbc default response has a 2xx status code
func (o *CreateDbcDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create dbc default response has a 3xx status code
func (o *CreateDbcDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create dbc default response has a 4xx status code
func (o *CreateDbcDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create dbc default response has a 5xx status code
func (o *CreateDbcDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create dbc default response a status code equal to that given
func (o *CreateDbcDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create dbc default response
func (o *CreateDbcDefault) Code() int {
	return o._statusCode
}

func (o *CreateDbcDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/createDbc][%d] createDbc default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDbcDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/createDbc][%d] createDbc default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDbcDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateDbcDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
