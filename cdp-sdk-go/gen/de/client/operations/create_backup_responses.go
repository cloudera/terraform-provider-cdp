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

// CreateBackupReader is a Reader for the CreateBackup structure.
type CreateBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateBackupOK creates a CreateBackupOK with default headers values
func NewCreateBackupOK() *CreateBackupOK {
	return &CreateBackupOK{}
}

/*
CreateBackupOK describes a response with status code 200, with default header values.

Response object for Create Backup command.
*/
type CreateBackupOK struct {
	Payload *models.CreateBackupResponse
}

// IsSuccess returns true when this create backup o k response has a 2xx status code
func (o *CreateBackupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create backup o k response has a 3xx status code
func (o *CreateBackupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create backup o k response has a 4xx status code
func (o *CreateBackupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create backup o k response has a 5xx status code
func (o *CreateBackupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create backup o k response a status code equal to that given
func (o *CreateBackupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create backup o k response
func (o *CreateBackupOK) Code() int {
	return 200
}

func (o *CreateBackupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/createBackup][%d] createBackupOK %s", 200, payload)
}

func (o *CreateBackupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/createBackup][%d] createBackupOK %s", 200, payload)
}

func (o *CreateBackupOK) GetPayload() *models.CreateBackupResponse {
	return o.Payload
}

func (o *CreateBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateBackupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBackupDefault creates a CreateBackupDefault with default headers values
func NewCreateBackupDefault(code int) *CreateBackupDefault {
	return &CreateBackupDefault{
		_statusCode: code,
	}
}

/*
CreateBackupDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CreateBackupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create backup default response has a 2xx status code
func (o *CreateBackupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create backup default response has a 3xx status code
func (o *CreateBackupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create backup default response has a 4xx status code
func (o *CreateBackupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create backup default response has a 5xx status code
func (o *CreateBackupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create backup default response a status code equal to that given
func (o *CreateBackupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create backup default response
func (o *CreateBackupDefault) Code() int {
	return o._statusCode
}

func (o *CreateBackupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/createBackup][%d] createBackup default %s", o._statusCode, payload)
}

func (o *CreateBackupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/createBackup][%d] createBackup default %s", o._statusCode, payload)
}

func (o *CreateBackupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
