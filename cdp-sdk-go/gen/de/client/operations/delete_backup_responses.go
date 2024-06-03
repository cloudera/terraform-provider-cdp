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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/de/models"
)

// DeleteBackupReader is a Reader for the DeleteBackup structure.
type DeleteBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteBackupOK creates a DeleteBackupOK with default headers values
func NewDeleteBackupOK() *DeleteBackupOK {
	return &DeleteBackupOK{}
}

/*
DeleteBackupOK describes a response with status code 200, with default header values.

Response object for Delete Backup command.
*/
type DeleteBackupOK struct {
	Payload models.DeleteBackupResponse
}

// IsSuccess returns true when this delete backup o k response has a 2xx status code
func (o *DeleteBackupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete backup o k response has a 3xx status code
func (o *DeleteBackupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete backup o k response has a 4xx status code
func (o *DeleteBackupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete backup o k response has a 5xx status code
func (o *DeleteBackupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete backup o k response a status code equal to that given
func (o *DeleteBackupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete backup o k response
func (o *DeleteBackupOK) Code() int {
	return 200
}

func (o *DeleteBackupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/deleteBackup][%d] deleteBackupOK %s", 200, payload)
}

func (o *DeleteBackupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/deleteBackup][%d] deleteBackupOK %s", 200, payload)
}

func (o *DeleteBackupOK) GetPayload() models.DeleteBackupResponse {
	return o.Payload
}

func (o *DeleteBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBackupDefault creates a DeleteBackupDefault with default headers values
func NewDeleteBackupDefault(code int) *DeleteBackupDefault {
	return &DeleteBackupDefault{
		_statusCode: code,
	}
}

/*
DeleteBackupDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DeleteBackupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete backup default response has a 2xx status code
func (o *DeleteBackupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete backup default response has a 3xx status code
func (o *DeleteBackupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete backup default response has a 4xx status code
func (o *DeleteBackupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete backup default response has a 5xx status code
func (o *DeleteBackupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete backup default response a status code equal to that given
func (o *DeleteBackupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete backup default response
func (o *DeleteBackupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBackupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/deleteBackup][%d] deleteBackup default %s", o._statusCode, payload)
}

func (o *DeleteBackupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/deleteBackup][%d] deleteBackup default %s", o._statusCode, payload)
}

func (o *DeleteBackupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
