// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// GetAutomatedSyncEnvironmentStatusReader is a Reader for the GetAutomatedSyncEnvironmentStatus structure.
type GetAutomatedSyncEnvironmentStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAutomatedSyncEnvironmentStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAutomatedSyncEnvironmentStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAutomatedSyncEnvironmentStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAutomatedSyncEnvironmentStatusOK creates a GetAutomatedSyncEnvironmentStatusOK with default headers values
func NewGetAutomatedSyncEnvironmentStatusOK() *GetAutomatedSyncEnvironmentStatusOK {
	return &GetAutomatedSyncEnvironmentStatusOK{}
}

/*
GetAutomatedSyncEnvironmentStatusOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetAutomatedSyncEnvironmentStatusOK struct {
	Payload *models.GetAutomatedSyncEnvironmentStatusResponse
}

// IsSuccess returns true when this get automated sync environment status o k response has a 2xx status code
func (o *GetAutomatedSyncEnvironmentStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get automated sync environment status o k response has a 3xx status code
func (o *GetAutomatedSyncEnvironmentStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get automated sync environment status o k response has a 4xx status code
func (o *GetAutomatedSyncEnvironmentStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get automated sync environment status o k response has a 5xx status code
func (o *GetAutomatedSyncEnvironmentStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get automated sync environment status o k response a status code equal to that given
func (o *GetAutomatedSyncEnvironmentStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get automated sync environment status o k response
func (o *GetAutomatedSyncEnvironmentStatusOK) Code() int {
	return 200
}

func (o *GetAutomatedSyncEnvironmentStatusOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getAutomatedSyncEnvironmentStatus][%d] getAutomatedSyncEnvironmentStatusOK  %+v", 200, o.Payload)
}

func (o *GetAutomatedSyncEnvironmentStatusOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getAutomatedSyncEnvironmentStatus][%d] getAutomatedSyncEnvironmentStatusOK  %+v", 200, o.Payload)
}

func (o *GetAutomatedSyncEnvironmentStatusOK) GetPayload() *models.GetAutomatedSyncEnvironmentStatusResponse {
	return o.Payload
}

func (o *GetAutomatedSyncEnvironmentStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAutomatedSyncEnvironmentStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAutomatedSyncEnvironmentStatusDefault creates a GetAutomatedSyncEnvironmentStatusDefault with default headers values
func NewGetAutomatedSyncEnvironmentStatusDefault(code int) *GetAutomatedSyncEnvironmentStatusDefault {
	return &GetAutomatedSyncEnvironmentStatusDefault{
		_statusCode: code,
	}
}

/*
GetAutomatedSyncEnvironmentStatusDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetAutomatedSyncEnvironmentStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get automated sync environment status default response has a 2xx status code
func (o *GetAutomatedSyncEnvironmentStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get automated sync environment status default response has a 3xx status code
func (o *GetAutomatedSyncEnvironmentStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get automated sync environment status default response has a 4xx status code
func (o *GetAutomatedSyncEnvironmentStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get automated sync environment status default response has a 5xx status code
func (o *GetAutomatedSyncEnvironmentStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get automated sync environment status default response a status code equal to that given
func (o *GetAutomatedSyncEnvironmentStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get automated sync environment status default response
func (o *GetAutomatedSyncEnvironmentStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetAutomatedSyncEnvironmentStatusDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getAutomatedSyncEnvironmentStatus][%d] getAutomatedSyncEnvironmentStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetAutomatedSyncEnvironmentStatusDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/getAutomatedSyncEnvironmentStatus][%d] getAutomatedSyncEnvironmentStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetAutomatedSyncEnvironmentStatusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAutomatedSyncEnvironmentStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
