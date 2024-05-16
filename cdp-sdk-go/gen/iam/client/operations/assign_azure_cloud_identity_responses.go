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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// AssignAzureCloudIdentityReader is a Reader for the AssignAzureCloudIdentity structure.
type AssignAzureCloudIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignAzureCloudIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignAzureCloudIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssignAzureCloudIdentityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssignAzureCloudIdentityOK creates a AssignAzureCloudIdentityOK with default headers values
func NewAssignAzureCloudIdentityOK() *AssignAzureCloudIdentityOK {
	return &AssignAzureCloudIdentityOK{}
}

/*
AssignAzureCloudIdentityOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type AssignAzureCloudIdentityOK struct {
	Payload models.AssignAzureCloudIdentityResponse
}

// IsSuccess returns true when this assign azure cloud identity o k response has a 2xx status code
func (o *AssignAzureCloudIdentityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this assign azure cloud identity o k response has a 3xx status code
func (o *AssignAzureCloudIdentityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this assign azure cloud identity o k response has a 4xx status code
func (o *AssignAzureCloudIdentityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this assign azure cloud identity o k response has a 5xx status code
func (o *AssignAzureCloudIdentityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this assign azure cloud identity o k response a status code equal to that given
func (o *AssignAzureCloudIdentityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the assign azure cloud identity o k response
func (o *AssignAzureCloudIdentityOK) Code() int {
	return 200
}

func (o *AssignAzureCloudIdentityOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/assignAzureCloudIdentity][%d] assignAzureCloudIdentityOK %s", 200, payload)
}

func (o *AssignAzureCloudIdentityOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/assignAzureCloudIdentity][%d] assignAzureCloudIdentityOK %s", 200, payload)
}

func (o *AssignAzureCloudIdentityOK) GetPayload() models.AssignAzureCloudIdentityResponse {
	return o.Payload
}

func (o *AssignAzureCloudIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignAzureCloudIdentityDefault creates a AssignAzureCloudIdentityDefault with default headers values
func NewAssignAzureCloudIdentityDefault(code int) *AssignAzureCloudIdentityDefault {
	return &AssignAzureCloudIdentityDefault{
		_statusCode: code,
	}
}

/*
AssignAzureCloudIdentityDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type AssignAzureCloudIdentityDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this assign azure cloud identity default response has a 2xx status code
func (o *AssignAzureCloudIdentityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this assign azure cloud identity default response has a 3xx status code
func (o *AssignAzureCloudIdentityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this assign azure cloud identity default response has a 4xx status code
func (o *AssignAzureCloudIdentityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this assign azure cloud identity default response has a 5xx status code
func (o *AssignAzureCloudIdentityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this assign azure cloud identity default response a status code equal to that given
func (o *AssignAzureCloudIdentityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the assign azure cloud identity default response
func (o *AssignAzureCloudIdentityDefault) Code() int {
	return o._statusCode
}

func (o *AssignAzureCloudIdentityDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/assignAzureCloudIdentity][%d] assignAzureCloudIdentity default %s", o._statusCode, payload)
}

func (o *AssignAzureCloudIdentityDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/assignAzureCloudIdentity][%d] assignAzureCloudIdentity default %s", o._statusCode, payload)
}

func (o *AssignAzureCloudIdentityDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AssignAzureCloudIdentityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
