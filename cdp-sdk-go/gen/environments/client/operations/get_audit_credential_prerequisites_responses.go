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

// GetAuditCredentialPrerequisitesReader is a Reader for the GetAuditCredentialPrerequisites structure.
type GetAuditCredentialPrerequisitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditCredentialPrerequisitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditCredentialPrerequisitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAuditCredentialPrerequisitesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuditCredentialPrerequisitesOK creates a GetAuditCredentialPrerequisitesOK with default headers values
func NewGetAuditCredentialPrerequisitesOK() *GetAuditCredentialPrerequisitesOK {
	return &GetAuditCredentialPrerequisitesOK{}
}

/*
GetAuditCredentialPrerequisitesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetAuditCredentialPrerequisitesOK struct {
	Payload *models.GetAuditCredentialPrerequisitesResponse
}

// IsSuccess returns true when this get audit credential prerequisites o k response has a 2xx status code
func (o *GetAuditCredentialPrerequisitesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get audit credential prerequisites o k response has a 3xx status code
func (o *GetAuditCredentialPrerequisitesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audit credential prerequisites o k response has a 4xx status code
func (o *GetAuditCredentialPrerequisitesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get audit credential prerequisites o k response has a 5xx status code
func (o *GetAuditCredentialPrerequisitesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get audit credential prerequisites o k response a status code equal to that given
func (o *GetAuditCredentialPrerequisitesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get audit credential prerequisites o k response
func (o *GetAuditCredentialPrerequisitesOK) Code() int {
	return 200
}

func (o *GetAuditCredentialPrerequisitesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/getAuditCredentialPrerequisites][%d] getAuditCredentialPrerequisitesOK %s", 200, payload)
}

func (o *GetAuditCredentialPrerequisitesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/getAuditCredentialPrerequisites][%d] getAuditCredentialPrerequisitesOK %s", 200, payload)
}

func (o *GetAuditCredentialPrerequisitesOK) GetPayload() *models.GetAuditCredentialPrerequisitesResponse {
	return o.Payload
}

func (o *GetAuditCredentialPrerequisitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAuditCredentialPrerequisitesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditCredentialPrerequisitesDefault creates a GetAuditCredentialPrerequisitesDefault with default headers values
func NewGetAuditCredentialPrerequisitesDefault(code int) *GetAuditCredentialPrerequisitesDefault {
	return &GetAuditCredentialPrerequisitesDefault{
		_statusCode: code,
	}
}

/*
GetAuditCredentialPrerequisitesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetAuditCredentialPrerequisitesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get audit credential prerequisites default response has a 2xx status code
func (o *GetAuditCredentialPrerequisitesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get audit credential prerequisites default response has a 3xx status code
func (o *GetAuditCredentialPrerequisitesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get audit credential prerequisites default response has a 4xx status code
func (o *GetAuditCredentialPrerequisitesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get audit credential prerequisites default response has a 5xx status code
func (o *GetAuditCredentialPrerequisitesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get audit credential prerequisites default response a status code equal to that given
func (o *GetAuditCredentialPrerequisitesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get audit credential prerequisites default response
func (o *GetAuditCredentialPrerequisitesDefault) Code() int {
	return o._statusCode
}

func (o *GetAuditCredentialPrerequisitesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/getAuditCredentialPrerequisites][%d] getAuditCredentialPrerequisites default %s", o._statusCode, payload)
}

func (o *GetAuditCredentialPrerequisitesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/getAuditCredentialPrerequisites][%d] getAuditCredentialPrerequisites default %s", o._statusCode, payload)
}

func (o *GetAuditCredentialPrerequisitesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAuditCredentialPrerequisitesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
