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

// CheckDatabaseConnectivityReader is a Reader for the CheckDatabaseConnectivity structure.
type CheckDatabaseConnectivityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckDatabaseConnectivityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckDatabaseConnectivityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCheckDatabaseConnectivityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCheckDatabaseConnectivityOK creates a CheckDatabaseConnectivityOK with default headers values
func NewCheckDatabaseConnectivityOK() *CheckDatabaseConnectivityOK {
	return &CheckDatabaseConnectivityOK{}
}

/*
CheckDatabaseConnectivityOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CheckDatabaseConnectivityOK struct {
	Payload *models.CheckDatabaseConnectivityResponse
}

// IsSuccess returns true when this check database connectivity o k response has a 2xx status code
func (o *CheckDatabaseConnectivityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this check database connectivity o k response has a 3xx status code
func (o *CheckDatabaseConnectivityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this check database connectivity o k response has a 4xx status code
func (o *CheckDatabaseConnectivityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this check database connectivity o k response has a 5xx status code
func (o *CheckDatabaseConnectivityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this check database connectivity o k response a status code equal to that given
func (o *CheckDatabaseConnectivityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the check database connectivity o k response
func (o *CheckDatabaseConnectivityOK) Code() int {
	return 200
}

func (o *CheckDatabaseConnectivityOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/checkDatabaseConnectivity][%d] checkDatabaseConnectivityOK %s", 200, payload)
}

func (o *CheckDatabaseConnectivityOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/checkDatabaseConnectivity][%d] checkDatabaseConnectivityOK %s", 200, payload)
}

func (o *CheckDatabaseConnectivityOK) GetPayload() *models.CheckDatabaseConnectivityResponse {
	return o.Payload
}

func (o *CheckDatabaseConnectivityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CheckDatabaseConnectivityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckDatabaseConnectivityDefault creates a CheckDatabaseConnectivityDefault with default headers values
func NewCheckDatabaseConnectivityDefault(code int) *CheckDatabaseConnectivityDefault {
	return &CheckDatabaseConnectivityDefault{
		_statusCode: code,
	}
}

/*
CheckDatabaseConnectivityDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CheckDatabaseConnectivityDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this check database connectivity default response has a 2xx status code
func (o *CheckDatabaseConnectivityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this check database connectivity default response has a 3xx status code
func (o *CheckDatabaseConnectivityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this check database connectivity default response has a 4xx status code
func (o *CheckDatabaseConnectivityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this check database connectivity default response has a 5xx status code
func (o *CheckDatabaseConnectivityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this check database connectivity default response a status code equal to that given
func (o *CheckDatabaseConnectivityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the check database connectivity default response
func (o *CheckDatabaseConnectivityDefault) Code() int {
	return o._statusCode
}

func (o *CheckDatabaseConnectivityDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/checkDatabaseConnectivity][%d] checkDatabaseConnectivity default %s", o._statusCode, payload)
}

func (o *CheckDatabaseConnectivityDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/checkDatabaseConnectivity][%d] checkDatabaseConnectivity default %s", o._statusCode, payload)
}

func (o *CheckDatabaseConnectivityDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CheckDatabaseConnectivityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
