// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// SetEnvironmentSettingReader is a Reader for the SetEnvironmentSetting structure.
type SetEnvironmentSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetEnvironmentSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetEnvironmentSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetEnvironmentSettingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetEnvironmentSettingOK creates a SetEnvironmentSettingOK with default headers values
func NewSetEnvironmentSettingOK() *SetEnvironmentSettingOK {
	return &SetEnvironmentSettingOK{}
}

/*
SetEnvironmentSettingOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type SetEnvironmentSettingOK struct {
	Payload models.SetEnvironmentSettingResponse
}

// IsSuccess returns true when this set environment setting o k response has a 2xx status code
func (o *SetEnvironmentSettingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set environment setting o k response has a 3xx status code
func (o *SetEnvironmentSettingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set environment setting o k response has a 4xx status code
func (o *SetEnvironmentSettingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set environment setting o k response has a 5xx status code
func (o *SetEnvironmentSettingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set environment setting o k response a status code equal to that given
func (o *SetEnvironmentSettingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set environment setting o k response
func (o *SetEnvironmentSettingOK) Code() int {
	return 200
}

func (o *SetEnvironmentSettingOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setEnvironmentSetting][%d] setEnvironmentSettingOK  %+v", 200, o.Payload)
}

func (o *SetEnvironmentSettingOK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setEnvironmentSetting][%d] setEnvironmentSettingOK  %+v", 200, o.Payload)
}

func (o *SetEnvironmentSettingOK) GetPayload() models.SetEnvironmentSettingResponse {
	return o.Payload
}

func (o *SetEnvironmentSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetEnvironmentSettingDefault creates a SetEnvironmentSettingDefault with default headers values
func NewSetEnvironmentSettingDefault(code int) *SetEnvironmentSettingDefault {
	return &SetEnvironmentSettingDefault{
		_statusCode: code,
	}
}

/*
SetEnvironmentSettingDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type SetEnvironmentSettingDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this set environment setting default response has a 2xx status code
func (o *SetEnvironmentSettingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this set environment setting default response has a 3xx status code
func (o *SetEnvironmentSettingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this set environment setting default response has a 4xx status code
func (o *SetEnvironmentSettingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this set environment setting default response has a 5xx status code
func (o *SetEnvironmentSettingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this set environment setting default response a status code equal to that given
func (o *SetEnvironmentSettingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the set environment setting default response
func (o *SetEnvironmentSettingDefault) Code() int {
	return o._statusCode
}

func (o *SetEnvironmentSettingDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setEnvironmentSetting][%d] setEnvironmentSetting default  %+v", o._statusCode, o.Payload)
}

func (o *SetEnvironmentSettingDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/setEnvironmentSetting][%d] setEnvironmentSetting default  %+v", o._statusCode, o.Payload)
}

func (o *SetEnvironmentSettingDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetEnvironmentSettingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
