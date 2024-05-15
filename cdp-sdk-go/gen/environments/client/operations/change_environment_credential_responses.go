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

// ChangeEnvironmentCredentialReader is a Reader for the ChangeEnvironmentCredential structure.
type ChangeEnvironmentCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeEnvironmentCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeEnvironmentCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeEnvironmentCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeEnvironmentCredentialOK creates a ChangeEnvironmentCredentialOK with default headers values
func NewChangeEnvironmentCredentialOK() *ChangeEnvironmentCredentialOK {
	return &ChangeEnvironmentCredentialOK{}
}

/*
ChangeEnvironmentCredentialOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ChangeEnvironmentCredentialOK struct {
	Payload *models.ChangeEnvironmentCredentialResponse
}

// IsSuccess returns true when this change environment credential o k response has a 2xx status code
func (o *ChangeEnvironmentCredentialOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this change environment credential o k response has a 3xx status code
func (o *ChangeEnvironmentCredentialOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this change environment credential o k response has a 4xx status code
func (o *ChangeEnvironmentCredentialOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this change environment credential o k response has a 5xx status code
func (o *ChangeEnvironmentCredentialOK) IsServerError() bool {
	return false
}

// IsCode returns true when this change environment credential o k response a status code equal to that given
func (o *ChangeEnvironmentCredentialOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the change environment credential o k response
func (o *ChangeEnvironmentCredentialOK) Code() int {
	return 200
}

func (o *ChangeEnvironmentCredentialOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/changeEnvironmentCredential][%d] changeEnvironmentCredentialOK %s", 200, payload)
}

func (o *ChangeEnvironmentCredentialOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/changeEnvironmentCredential][%d] changeEnvironmentCredentialOK %s", 200, payload)
}

func (o *ChangeEnvironmentCredentialOK) GetPayload() *models.ChangeEnvironmentCredentialResponse {
	return o.Payload
}

func (o *ChangeEnvironmentCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChangeEnvironmentCredentialResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeEnvironmentCredentialDefault creates a ChangeEnvironmentCredentialDefault with default headers values
func NewChangeEnvironmentCredentialDefault(code int) *ChangeEnvironmentCredentialDefault {
	return &ChangeEnvironmentCredentialDefault{
		_statusCode: code,
	}
}

/*
ChangeEnvironmentCredentialDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ChangeEnvironmentCredentialDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this change environment credential default response has a 2xx status code
func (o *ChangeEnvironmentCredentialDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this change environment credential default response has a 3xx status code
func (o *ChangeEnvironmentCredentialDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this change environment credential default response has a 4xx status code
func (o *ChangeEnvironmentCredentialDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this change environment credential default response has a 5xx status code
func (o *ChangeEnvironmentCredentialDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this change environment credential default response a status code equal to that given
func (o *ChangeEnvironmentCredentialDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the change environment credential default response
func (o *ChangeEnvironmentCredentialDefault) Code() int {
	return o._statusCode
}

func (o *ChangeEnvironmentCredentialDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/changeEnvironmentCredential][%d] changeEnvironmentCredential default %s", o._statusCode, payload)
}

func (o *ChangeEnvironmentCredentialDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/changeEnvironmentCredential][%d] changeEnvironmentCredential default %s", o._statusCode, payload)
}

func (o *ChangeEnvironmentCredentialDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ChangeEnvironmentCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
