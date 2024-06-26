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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// RestoreBackupReader is a Reader for the RestoreBackup structure.
type RestoreBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestoreBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestoreBackupOK creates a RestoreBackupOK with default headers values
func NewRestoreBackupOK() *RestoreBackupOK {
	return &RestoreBackupOK{}
}

/*
RestoreBackupOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type RestoreBackupOK struct {
	Payload *models.RestoreBackupResponse
}

// IsSuccess returns true when this restore backup o k response has a 2xx status code
func (o *RestoreBackupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restore backup o k response has a 3xx status code
func (o *RestoreBackupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore backup o k response has a 4xx status code
func (o *RestoreBackupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restore backup o k response has a 5xx status code
func (o *RestoreBackupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this restore backup o k response a status code equal to that given
func (o *RestoreBackupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the restore backup o k response
func (o *RestoreBackupOK) Code() int {
	return 200
}

func (o *RestoreBackupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/restoreBackup][%d] restoreBackupOK %s", 200, payload)
}

func (o *RestoreBackupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/restoreBackup][%d] restoreBackupOK %s", 200, payload)
}

func (o *RestoreBackupOK) GetPayload() *models.RestoreBackupResponse {
	return o.Payload
}

func (o *RestoreBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestoreBackupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreBackupDefault creates a RestoreBackupDefault with default headers values
func NewRestoreBackupDefault(code int) *RestoreBackupDefault {
	return &RestoreBackupDefault{
		_statusCode: code,
	}
}

/*
RestoreBackupDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type RestoreBackupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this restore backup default response has a 2xx status code
func (o *RestoreBackupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this restore backup default response has a 3xx status code
func (o *RestoreBackupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this restore backup default response has a 4xx status code
func (o *RestoreBackupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this restore backup default response has a 5xx status code
func (o *RestoreBackupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this restore backup default response a status code equal to that given
func (o *RestoreBackupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the restore backup default response
func (o *RestoreBackupDefault) Code() int {
	return o._statusCode
}

func (o *RestoreBackupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/restoreBackup][%d] restoreBackup default %s", o._statusCode, payload)
}

func (o *RestoreBackupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/restoreBackup][%d] restoreBackup default %s", o._statusCode, payload)
}

func (o *RestoreBackupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RestoreBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
