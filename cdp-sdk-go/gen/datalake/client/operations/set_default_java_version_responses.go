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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// SetDefaultJavaVersionReader is a Reader for the SetDefaultJavaVersion structure.
type SetDefaultJavaVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDefaultJavaVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetDefaultJavaVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetDefaultJavaVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetDefaultJavaVersionOK creates a SetDefaultJavaVersionOK with default headers values
func NewSetDefaultJavaVersionOK() *SetDefaultJavaVersionOK {
	return &SetDefaultJavaVersionOK{}
}

/*
SetDefaultJavaVersionOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SetDefaultJavaVersionOK struct {
	Payload *models.SetDefaultJavaVersionResponse
}

// IsSuccess returns true when this set default java version o k response has a 2xx status code
func (o *SetDefaultJavaVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set default java version o k response has a 3xx status code
func (o *SetDefaultJavaVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set default java version o k response has a 4xx status code
func (o *SetDefaultJavaVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set default java version o k response has a 5xx status code
func (o *SetDefaultJavaVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set default java version o k response a status code equal to that given
func (o *SetDefaultJavaVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set default java version o k response
func (o *SetDefaultJavaVersionOK) Code() int {
	return 200
}

func (o *SetDefaultJavaVersionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/setDefaultJavaVersion][%d] setDefaultJavaVersionOK %s", 200, payload)
}

func (o *SetDefaultJavaVersionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/setDefaultJavaVersion][%d] setDefaultJavaVersionOK %s", 200, payload)
}

func (o *SetDefaultJavaVersionOK) GetPayload() *models.SetDefaultJavaVersionResponse {
	return o.Payload
}

func (o *SetDefaultJavaVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SetDefaultJavaVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDefaultJavaVersionDefault creates a SetDefaultJavaVersionDefault with default headers values
func NewSetDefaultJavaVersionDefault(code int) *SetDefaultJavaVersionDefault {
	return &SetDefaultJavaVersionDefault{
		_statusCode: code,
	}
}

/*
SetDefaultJavaVersionDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SetDefaultJavaVersionDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this set default java version default response has a 2xx status code
func (o *SetDefaultJavaVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this set default java version default response has a 3xx status code
func (o *SetDefaultJavaVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this set default java version default response has a 4xx status code
func (o *SetDefaultJavaVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this set default java version default response has a 5xx status code
func (o *SetDefaultJavaVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this set default java version default response a status code equal to that given
func (o *SetDefaultJavaVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the set default java version default response
func (o *SetDefaultJavaVersionDefault) Code() int {
	return o._statusCode
}

func (o *SetDefaultJavaVersionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/setDefaultJavaVersion][%d] setDefaultJavaVersion default %s", o._statusCode, payload)
}

func (o *SetDefaultJavaVersionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/setDefaultJavaVersion][%d] setDefaultJavaVersion default %s", o._statusCode, payload)
}

func (o *SetDefaultJavaVersionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetDefaultJavaVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
