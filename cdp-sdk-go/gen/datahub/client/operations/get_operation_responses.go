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

// GetOperationReader is a Reader for the GetOperation structure.
type GetOperationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOperationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOperationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetOperationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOperationOK creates a GetOperationOK with default headers values
func NewGetOperationOK() *GetOperationOK {
	return &GetOperationOK{}
}

/*
GetOperationOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetOperationOK struct {
	Payload *models.GetOperationResponse
}

// IsSuccess returns true when this get operation o k response has a 2xx status code
func (o *GetOperationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get operation o k response has a 3xx status code
func (o *GetOperationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get operation o k response has a 4xx status code
func (o *GetOperationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get operation o k response has a 5xx status code
func (o *GetOperationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get operation o k response a status code equal to that given
func (o *GetOperationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get operation o k response
func (o *GetOperationOK) Code() int {
	return 200
}

func (o *GetOperationOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/getOperation][%d] getOperationOK  %+v", 200, o.Payload)
}

func (o *GetOperationOK) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/getOperation][%d] getOperationOK  %+v", 200, o.Payload)
}

func (o *GetOperationOK) GetPayload() *models.GetOperationResponse {
	return o.Payload
}

func (o *GetOperationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetOperationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOperationDefault creates a GetOperationDefault with default headers values
func NewGetOperationDefault(code int) *GetOperationDefault {
	return &GetOperationDefault{
		_statusCode: code,
	}
}

/*
GetOperationDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetOperationDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get operation default response has a 2xx status code
func (o *GetOperationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get operation default response has a 3xx status code
func (o *GetOperationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get operation default response has a 4xx status code
func (o *GetOperationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get operation default response has a 5xx status code
func (o *GetOperationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get operation default response a status code equal to that given
func (o *GetOperationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get operation default response
func (o *GetOperationDefault) Code() int {
	return o._statusCode
}

func (o *GetOperationDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/getOperation][%d] getOperation default  %+v", o._statusCode, o.Payload)
}

func (o *GetOperationDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/getOperation][%d] getOperation default  %+v", o._statusCode, o.Payload)
}

func (o *GetOperationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetOperationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
