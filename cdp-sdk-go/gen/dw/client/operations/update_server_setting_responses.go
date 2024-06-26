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

// UpdateServerSettingReader is a Reader for the UpdateServerSetting structure.
type UpdateServerSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServerSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateServerSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateServerSettingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateServerSettingOK creates a UpdateServerSettingOK with default headers values
func NewUpdateServerSettingOK() *UpdateServerSettingOK {
	return &UpdateServerSettingOK{}
}

/*
UpdateServerSettingOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateServerSettingOK struct {
	Payload *models.UpdateServerSettingResponse
}

// IsSuccess returns true when this update server setting o k response has a 2xx status code
func (o *UpdateServerSettingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update server setting o k response has a 3xx status code
func (o *UpdateServerSettingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update server setting o k response has a 4xx status code
func (o *UpdateServerSettingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update server setting o k response has a 5xx status code
func (o *UpdateServerSettingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update server setting o k response a status code equal to that given
func (o *UpdateServerSettingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update server setting o k response
func (o *UpdateServerSettingOK) Code() int {
	return 200
}

func (o *UpdateServerSettingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/updateServerSetting][%d] updateServerSettingOK %s", 200, payload)
}

func (o *UpdateServerSettingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/updateServerSetting][%d] updateServerSettingOK %s", 200, payload)
}

func (o *UpdateServerSettingOK) GetPayload() *models.UpdateServerSettingResponse {
	return o.Payload
}

func (o *UpdateServerSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateServerSettingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServerSettingDefault creates a UpdateServerSettingDefault with default headers values
func NewUpdateServerSettingDefault(code int) *UpdateServerSettingDefault {
	return &UpdateServerSettingDefault{
		_statusCode: code,
	}
}

/*
UpdateServerSettingDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateServerSettingDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update server setting default response has a 2xx status code
func (o *UpdateServerSettingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update server setting default response has a 3xx status code
func (o *UpdateServerSettingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update server setting default response has a 4xx status code
func (o *UpdateServerSettingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update server setting default response has a 5xx status code
func (o *UpdateServerSettingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update server setting default response a status code equal to that given
func (o *UpdateServerSettingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update server setting default response
func (o *UpdateServerSettingDefault) Code() int {
	return o._statusCode
}

func (o *UpdateServerSettingDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/updateServerSetting][%d] updateServerSetting default %s", o._statusCode, payload)
}

func (o *UpdateServerSettingDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/updateServerSetting][%d] updateServerSetting default %s", o._statusCode, payload)
}

func (o *UpdateServerSettingDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateServerSettingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
