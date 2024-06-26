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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

// RestoreSnapshotReader is a Reader for the RestoreSnapshot structure.
type RestoreSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestoreSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestoreSnapshotOK creates a RestoreSnapshotOK with default headers values
func NewRestoreSnapshotOK() *RestoreSnapshotOK {
	return &RestoreSnapshotOK{}
}

/*
RestoreSnapshotOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type RestoreSnapshotOK struct {
	Payload *models.RestoreSnapshotResponse
}

// IsSuccess returns true when this restore snapshot o k response has a 2xx status code
func (o *RestoreSnapshotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restore snapshot o k response has a 3xx status code
func (o *RestoreSnapshotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore snapshot o k response has a 4xx status code
func (o *RestoreSnapshotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restore snapshot o k response has a 5xx status code
func (o *RestoreSnapshotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this restore snapshot o k response a status code equal to that given
func (o *RestoreSnapshotOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the restore snapshot o k response
func (o *RestoreSnapshotOK) Code() int {
	return 200
}

func (o *RestoreSnapshotOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/restoreSnapshot][%d] restoreSnapshotOK %s", 200, payload)
}

func (o *RestoreSnapshotOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/restoreSnapshot][%d] restoreSnapshotOK %s", 200, payload)
}

func (o *RestoreSnapshotOK) GetPayload() *models.RestoreSnapshotResponse {
	return o.Payload
}

func (o *RestoreSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestoreSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreSnapshotDefault creates a RestoreSnapshotDefault with default headers values
func NewRestoreSnapshotDefault(code int) *RestoreSnapshotDefault {
	return &RestoreSnapshotDefault{
		_statusCode: code,
	}
}

/*
RestoreSnapshotDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type RestoreSnapshotDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this restore snapshot default response has a 2xx status code
func (o *RestoreSnapshotDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this restore snapshot default response has a 3xx status code
func (o *RestoreSnapshotDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this restore snapshot default response has a 4xx status code
func (o *RestoreSnapshotDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this restore snapshot default response has a 5xx status code
func (o *RestoreSnapshotDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this restore snapshot default response a status code equal to that given
func (o *RestoreSnapshotDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the restore snapshot default response
func (o *RestoreSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *RestoreSnapshotDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/restoreSnapshot][%d] restoreSnapshot default %s", o._statusCode, payload)
}

func (o *RestoreSnapshotDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/restoreSnapshot][%d] restoreSnapshot default %s", o._statusCode, payload)
}

func (o *RestoreSnapshotDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RestoreSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
