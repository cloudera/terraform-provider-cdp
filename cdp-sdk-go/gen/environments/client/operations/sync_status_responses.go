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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// SyncStatusReader is a Reader for the SyncStatus structure.
type SyncStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSyncStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSyncStatusOK creates a SyncStatusOK with default headers values
func NewSyncStatusOK() *SyncStatusOK {
	return &SyncStatusOK{}
}

/*
SyncStatusOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SyncStatusOK struct {
	Payload *models.SyncStatusResponse
}

// IsSuccess returns true when this sync status o k response has a 2xx status code
func (o *SyncStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sync status o k response has a 3xx status code
func (o *SyncStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sync status o k response has a 4xx status code
func (o *SyncStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sync status o k response has a 5xx status code
func (o *SyncStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sync status o k response a status code equal to that given
func (o *SyncStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sync status o k response
func (o *SyncStatusOK) Code() int {
	return 200
}

func (o *SyncStatusOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/syncStatus][%d] syncStatusOK %s", 200, payload)
}

func (o *SyncStatusOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/syncStatus][%d] syncStatusOK %s", 200, payload)
}

func (o *SyncStatusOK) GetPayload() *models.SyncStatusResponse {
	return o.Payload
}

func (o *SyncStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SyncStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncStatusDefault creates a SyncStatusDefault with default headers values
func NewSyncStatusDefault(code int) *SyncStatusDefault {
	return &SyncStatusDefault{
		_statusCode: code,
	}
}

/*
SyncStatusDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SyncStatusDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this sync status default response has a 2xx status code
func (o *SyncStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sync status default response has a 3xx status code
func (o *SyncStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sync status default response has a 4xx status code
func (o *SyncStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sync status default response has a 5xx status code
func (o *SyncStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sync status default response a status code equal to that given
func (o *SyncStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sync status default response
func (o *SyncStatusDefault) Code() int {
	return o._statusCode
}

func (o *SyncStatusDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/syncStatus][%d] syncStatus default %s", o._statusCode, payload)
}

func (o *SyncStatusDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/syncStatus][%d] syncStatus default %s", o._statusCode, payload)
}

func (o *SyncStatusDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SyncStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
