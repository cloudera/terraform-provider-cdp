// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// BackupDatalakeStatusReader is a Reader for the BackupDatalakeStatus structure.
type BackupDatalakeStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupDatalakeStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBackupDatalakeStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBackupDatalakeStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBackupDatalakeStatusOK creates a BackupDatalakeStatusOK with default headers values
func NewBackupDatalakeStatusOK() *BackupDatalakeStatusOK {
	return &BackupDatalakeStatusOK{}
}

/*
BackupDatalakeStatusOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type BackupDatalakeStatusOK struct {
	Payload *models.BackupDatalakeStatusResponse
}

// IsSuccess returns true when this backup datalake status o k response has a 2xx status code
func (o *BackupDatalakeStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this backup datalake status o k response has a 3xx status code
func (o *BackupDatalakeStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this backup datalake status o k response has a 4xx status code
func (o *BackupDatalakeStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this backup datalake status o k response has a 5xx status code
func (o *BackupDatalakeStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this backup datalake status o k response a status code equal to that given
func (o *BackupDatalakeStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the backup datalake status o k response
func (o *BackupDatalakeStatusOK) Code() int {
	return 200
}

func (o *BackupDatalakeStatusOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/backupDatalakeStatus][%d] backupDatalakeStatusOK  %+v", 200, o.Payload)
}

func (o *BackupDatalakeStatusOK) String() string {
	return fmt.Sprintf("[POST /api/v1/datalake/backupDatalakeStatus][%d] backupDatalakeStatusOK  %+v", 200, o.Payload)
}

func (o *BackupDatalakeStatusOK) GetPayload() *models.BackupDatalakeStatusResponse {
	return o.Payload
}

func (o *BackupDatalakeStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupDatalakeStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBackupDatalakeStatusDefault creates a BackupDatalakeStatusDefault with default headers values
func NewBackupDatalakeStatusDefault(code int) *BackupDatalakeStatusDefault {
	return &BackupDatalakeStatusDefault{
		_statusCode: code,
	}
}

/*
BackupDatalakeStatusDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type BackupDatalakeStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this backup datalake status default response has a 2xx status code
func (o *BackupDatalakeStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this backup datalake status default response has a 3xx status code
func (o *BackupDatalakeStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this backup datalake status default response has a 4xx status code
func (o *BackupDatalakeStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this backup datalake status default response has a 5xx status code
func (o *BackupDatalakeStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this backup datalake status default response a status code equal to that given
func (o *BackupDatalakeStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the backup datalake status default response
func (o *BackupDatalakeStatusDefault) Code() int {
	return o._statusCode
}

func (o *BackupDatalakeStatusDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datalake/backupDatalakeStatus][%d] backupDatalakeStatus default  %+v", o._statusCode, o.Payload)
}

func (o *BackupDatalakeStatusDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/datalake/backupDatalakeStatus][%d] backupDatalakeStatus default  %+v", o._statusCode, o.Payload)
}

func (o *BackupDatalakeStatusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *BackupDatalakeStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}