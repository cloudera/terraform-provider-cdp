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

// ListBackupEntitiesReader is a Reader for the ListBackupEntities structure.
type ListBackupEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBackupEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBackupEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListBackupEntitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListBackupEntitiesOK creates a ListBackupEntitiesOK with default headers values
func NewListBackupEntitiesOK() *ListBackupEntitiesOK {
	return &ListBackupEntitiesOK{}
}

/*
ListBackupEntitiesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListBackupEntitiesOK struct {
	Payload *models.ListBackupEntitiesResponse
}

// IsSuccess returns true when this list backup entities o k response has a 2xx status code
func (o *ListBackupEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list backup entities o k response has a 3xx status code
func (o *ListBackupEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list backup entities o k response has a 4xx status code
func (o *ListBackupEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list backup entities o k response has a 5xx status code
func (o *ListBackupEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list backup entities o k response a status code equal to that given
func (o *ListBackupEntitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list backup entities o k response
func (o *ListBackupEntitiesOK) Code() int {
	return 200
}

func (o *ListBackupEntitiesOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/listBackupEntities][%d] listBackupEntitiesOK  %+v", 200, o.Payload)
}

func (o *ListBackupEntitiesOK) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/listBackupEntities][%d] listBackupEntitiesOK  %+v", 200, o.Payload)
}

func (o *ListBackupEntitiesOK) GetPayload() *models.ListBackupEntitiesResponse {
	return o.Payload
}

func (o *ListBackupEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListBackupEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBackupEntitiesDefault creates a ListBackupEntitiesDefault with default headers values
func NewListBackupEntitiesDefault(code int) *ListBackupEntitiesDefault {
	return &ListBackupEntitiesDefault{
		_statusCode: code,
	}
}

/*
ListBackupEntitiesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListBackupEntitiesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list backup entities default response has a 2xx status code
func (o *ListBackupEntitiesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list backup entities default response has a 3xx status code
func (o *ListBackupEntitiesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list backup entities default response has a 4xx status code
func (o *ListBackupEntitiesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list backup entities default response has a 5xx status code
func (o *ListBackupEntitiesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list backup entities default response a status code equal to that given
func (o *ListBackupEntitiesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list backup entities default response
func (o *ListBackupEntitiesDefault) Code() int {
	return o._statusCode
}

func (o *ListBackupEntitiesDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/listBackupEntities][%d] listBackupEntities default  %+v", o._statusCode, o.Payload)
}

func (o *ListBackupEntitiesDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/listBackupEntities][%d] listBackupEntities default  %+v", o._statusCode, o.Payload)
}

func (o *ListBackupEntitiesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListBackupEntitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}