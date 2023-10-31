// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// GetAuditEventsReader is a Reader for the GetAuditEvents structure.
type GetAuditEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAuditEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuditEventsOK creates a GetAuditEventsOK with default headers values
func NewGetAuditEventsOK() *GetAuditEventsOK {
	return &GetAuditEventsOK{}
}

/*
GetAuditEventsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetAuditEventsOK struct {
	Payload *models.GetAuditEventsResponse
}

// IsSuccess returns true when this get audit events o k response has a 2xx status code
func (o *GetAuditEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get audit events o k response has a 3xx status code
func (o *GetAuditEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audit events o k response has a 4xx status code
func (o *GetAuditEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get audit events o k response has a 5xx status code
func (o *GetAuditEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get audit events o k response a status code equal to that given
func (o *GetAuditEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get audit events o k response
func (o *GetAuditEventsOK) Code() int {
	return 200
}

func (o *GetAuditEventsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/getAuditEvents][%d] getAuditEventsOK  %+v", 200, o.Payload)
}

func (o *GetAuditEventsOK) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/getAuditEvents][%d] getAuditEventsOK  %+v", 200, o.Payload)
}

func (o *GetAuditEventsOK) GetPayload() *models.GetAuditEventsResponse {
	return o.Payload
}

func (o *GetAuditEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAuditEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditEventsDefault creates a GetAuditEventsDefault with default headers values
func NewGetAuditEventsDefault(code int) *GetAuditEventsDefault {
	return &GetAuditEventsDefault{
		_statusCode: code,
	}
}

/*
GetAuditEventsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetAuditEventsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get audit events default response has a 2xx status code
func (o *GetAuditEventsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get audit events default response has a 3xx status code
func (o *GetAuditEventsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get audit events default response has a 4xx status code
func (o *GetAuditEventsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get audit events default response has a 5xx status code
func (o *GetAuditEventsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get audit events default response a status code equal to that given
func (o *GetAuditEventsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get audit events default response
func (o *GetAuditEventsDefault) Code() int {
	return o._statusCode
}

func (o *GetAuditEventsDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/getAuditEvents][%d] getAuditEvents default  %+v", o._statusCode, o.Payload)
}

func (o *GetAuditEventsDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/getAuditEvents][%d] getAuditEvents default  %+v", o._statusCode, o.Payload)
}

func (o *GetAuditEventsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAuditEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}