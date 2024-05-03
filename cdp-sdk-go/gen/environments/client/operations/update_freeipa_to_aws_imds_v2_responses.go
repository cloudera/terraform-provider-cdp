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

// UpdateFreeipaToAwsImdsV2Reader is a Reader for the UpdateFreeipaToAwsImdsV2 structure.
type UpdateFreeipaToAwsImdsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFreeipaToAwsImdsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFreeipaToAwsImdsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateFreeipaToAwsImdsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateFreeipaToAwsImdsV2OK creates a UpdateFreeipaToAwsImdsV2OK with default headers values
func NewUpdateFreeipaToAwsImdsV2OK() *UpdateFreeipaToAwsImdsV2OK {
	return &UpdateFreeipaToAwsImdsV2OK{}
}

/*
UpdateFreeipaToAwsImdsV2OK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateFreeipaToAwsImdsV2OK struct {
	Payload models.UpdateFreeipaToAwsImdsV2Response
}

// IsSuccess returns true when this update freeipa to aws imds v2 o k response has a 2xx status code
func (o *UpdateFreeipaToAwsImdsV2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update freeipa to aws imds v2 o k response has a 3xx status code
func (o *UpdateFreeipaToAwsImdsV2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update freeipa to aws imds v2 o k response has a 4xx status code
func (o *UpdateFreeipaToAwsImdsV2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update freeipa to aws imds v2 o k response has a 5xx status code
func (o *UpdateFreeipaToAwsImdsV2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this update freeipa to aws imds v2 o k response a status code equal to that given
func (o *UpdateFreeipaToAwsImdsV2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update freeipa to aws imds v2 o k response
func (o *UpdateFreeipaToAwsImdsV2OK) Code() int {
	return 200
}

func (o *UpdateFreeipaToAwsImdsV2OK) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/updateFreeipaToAwsImdsV2][%d] updateFreeipaToAwsImdsV2OK  %+v", 200, o.Payload)
}

func (o *UpdateFreeipaToAwsImdsV2OK) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/updateFreeipaToAwsImdsV2][%d] updateFreeipaToAwsImdsV2OK  %+v", 200, o.Payload)
}

func (o *UpdateFreeipaToAwsImdsV2OK) GetPayload() models.UpdateFreeipaToAwsImdsV2Response {
	return o.Payload
}

func (o *UpdateFreeipaToAwsImdsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFreeipaToAwsImdsV2Default creates a UpdateFreeipaToAwsImdsV2Default with default headers values
func NewUpdateFreeipaToAwsImdsV2Default(code int) *UpdateFreeipaToAwsImdsV2Default {
	return &UpdateFreeipaToAwsImdsV2Default{
		_statusCode: code,
	}
}

/*
UpdateFreeipaToAwsImdsV2Default describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateFreeipaToAwsImdsV2Default struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update freeipa to aws imds v2 default response has a 2xx status code
func (o *UpdateFreeipaToAwsImdsV2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update freeipa to aws imds v2 default response has a 3xx status code
func (o *UpdateFreeipaToAwsImdsV2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update freeipa to aws imds v2 default response has a 4xx status code
func (o *UpdateFreeipaToAwsImdsV2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update freeipa to aws imds v2 default response has a 5xx status code
func (o *UpdateFreeipaToAwsImdsV2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update freeipa to aws imds v2 default response a status code equal to that given
func (o *UpdateFreeipaToAwsImdsV2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update freeipa to aws imds v2 default response
func (o *UpdateFreeipaToAwsImdsV2Default) Code() int {
	return o._statusCode
}

func (o *UpdateFreeipaToAwsImdsV2Default) Error() string {
	return fmt.Sprintf("[POST /api/v1/environments2/updateFreeipaToAwsImdsV2][%d] updateFreeipaToAwsImdsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateFreeipaToAwsImdsV2Default) String() string {
	return fmt.Sprintf("[POST /api/v1/environments2/updateFreeipaToAwsImdsV2][%d] updateFreeipaToAwsImdsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateFreeipaToAwsImdsV2Default) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateFreeipaToAwsImdsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
