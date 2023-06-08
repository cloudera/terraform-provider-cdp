// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// DisableClouderaSSOLoginReader is a Reader for the DisableClouderaSSOLogin structure.
type DisableClouderaSSOLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DisableClouderaSSOLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDisableClouderaSSOLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDisableClouderaSSOLoginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDisableClouderaSSOLoginOK creates a DisableClouderaSSOLoginOK with default headers values
func NewDisableClouderaSSOLoginOK() *DisableClouderaSSOLoginOK {
	return &DisableClouderaSSOLoginOK{}
}

/*
DisableClouderaSSOLoginOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DisableClouderaSSOLoginOK struct {
	Payload models.DisableClouderaSSOLoginResponse
}

// IsSuccess returns true when this disable cloudera s s o login o k response has a 2xx status code
func (o *DisableClouderaSSOLoginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disable cloudera s s o login o k response has a 3xx status code
func (o *DisableClouderaSSOLoginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disable cloudera s s o login o k response has a 4xx status code
func (o *DisableClouderaSSOLoginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this disable cloudera s s o login o k response has a 5xx status code
func (o *DisableClouderaSSOLoginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this disable cloudera s s o login o k response a status code equal to that given
func (o *DisableClouderaSSOLoginOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the disable cloudera s s o login o k response
func (o *DisableClouderaSSOLoginOK) Code() int {
	return 200
}

func (o *DisableClouderaSSOLoginOK) Error() string {
	return fmt.Sprintf("[POST /iam/disableClouderaSSOLogin][%d] disableClouderaSSOLoginOK  %+v", 200, o.Payload)
}

func (o *DisableClouderaSSOLoginOK) String() string {
	return fmt.Sprintf("[POST /iam/disableClouderaSSOLogin][%d] disableClouderaSSOLoginOK  %+v", 200, o.Payload)
}

func (o *DisableClouderaSSOLoginOK) GetPayload() models.DisableClouderaSSOLoginResponse {
	return o.Payload
}

func (o *DisableClouderaSSOLoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDisableClouderaSSOLoginDefault creates a DisableClouderaSSOLoginDefault with default headers values
func NewDisableClouderaSSOLoginDefault(code int) *DisableClouderaSSOLoginDefault {
	return &DisableClouderaSSOLoginDefault{
		_statusCode: code,
	}
}

/*
DisableClouderaSSOLoginDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DisableClouderaSSOLoginDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this disable cloudera s s o login default response has a 2xx status code
func (o *DisableClouderaSSOLoginDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this disable cloudera s s o login default response has a 3xx status code
func (o *DisableClouderaSSOLoginDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this disable cloudera s s o login default response has a 4xx status code
func (o *DisableClouderaSSOLoginDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this disable cloudera s s o login default response has a 5xx status code
func (o *DisableClouderaSSOLoginDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this disable cloudera s s o login default response a status code equal to that given
func (o *DisableClouderaSSOLoginDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the disable cloudera s s o login default response
func (o *DisableClouderaSSOLoginDefault) Code() int {
	return o._statusCode
}

func (o *DisableClouderaSSOLoginDefault) Error() string {
	return fmt.Sprintf("[POST /iam/disableClouderaSSOLogin][%d] disableClouderaSSOLogin default  %+v", o._statusCode, o.Payload)
}

func (o *DisableClouderaSSOLoginDefault) String() string {
	return fmt.Sprintf("[POST /iam/disableClouderaSSOLogin][%d] disableClouderaSSOLogin default  %+v", o._statusCode, o.Payload)
}

func (o *DisableClouderaSSOLoginDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DisableClouderaSSOLoginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
