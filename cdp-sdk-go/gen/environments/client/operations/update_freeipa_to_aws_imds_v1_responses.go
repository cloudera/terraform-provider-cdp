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

// UpdateFreeipaToAwsImdsV1Reader is a Reader for the UpdateFreeipaToAwsImdsV1 structure.
type UpdateFreeipaToAwsImdsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFreeipaToAwsImdsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFreeipaToAwsImdsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateFreeipaToAwsImdsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateFreeipaToAwsImdsV1OK creates a UpdateFreeipaToAwsImdsV1OK with default headers values
func NewUpdateFreeipaToAwsImdsV1OK() *UpdateFreeipaToAwsImdsV1OK {
	return &UpdateFreeipaToAwsImdsV1OK{}
}

/*
UpdateFreeipaToAwsImdsV1OK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateFreeipaToAwsImdsV1OK struct {
	Payload *models.UpdateFreeipaToAwsImdsV1Response
}

// IsSuccess returns true when this update freeipa to aws imds v1 o k response has a 2xx status code
func (o *UpdateFreeipaToAwsImdsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update freeipa to aws imds v1 o k response has a 3xx status code
func (o *UpdateFreeipaToAwsImdsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update freeipa to aws imds v1 o k response has a 4xx status code
func (o *UpdateFreeipaToAwsImdsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update freeipa to aws imds v1 o k response has a 5xx status code
func (o *UpdateFreeipaToAwsImdsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this update freeipa to aws imds v1 o k response a status code equal to that given
func (o *UpdateFreeipaToAwsImdsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update freeipa to aws imds v1 o k response
func (o *UpdateFreeipaToAwsImdsV1OK) Code() int {
	return 200
}

func (o *UpdateFreeipaToAwsImdsV1OK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/updateFreeipaToAwsImdsV1][%d] updateFreeipaToAwsImdsV1OK %s", 200, payload)
}

func (o *UpdateFreeipaToAwsImdsV1OK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/updateFreeipaToAwsImdsV1][%d] updateFreeipaToAwsImdsV1OK %s", 200, payload)
}

func (o *UpdateFreeipaToAwsImdsV1OK) GetPayload() *models.UpdateFreeipaToAwsImdsV1Response {
	return o.Payload
}

func (o *UpdateFreeipaToAwsImdsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateFreeipaToAwsImdsV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFreeipaToAwsImdsV1Default creates a UpdateFreeipaToAwsImdsV1Default with default headers values
func NewUpdateFreeipaToAwsImdsV1Default(code int) *UpdateFreeipaToAwsImdsV1Default {
	return &UpdateFreeipaToAwsImdsV1Default{
		_statusCode: code,
	}
}

/*
UpdateFreeipaToAwsImdsV1Default describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateFreeipaToAwsImdsV1Default struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update freeipa to aws imds v1 default response has a 2xx status code
func (o *UpdateFreeipaToAwsImdsV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update freeipa to aws imds v1 default response has a 3xx status code
func (o *UpdateFreeipaToAwsImdsV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update freeipa to aws imds v1 default response has a 4xx status code
func (o *UpdateFreeipaToAwsImdsV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update freeipa to aws imds v1 default response has a 5xx status code
func (o *UpdateFreeipaToAwsImdsV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update freeipa to aws imds v1 default response a status code equal to that given
func (o *UpdateFreeipaToAwsImdsV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update freeipa to aws imds v1 default response
func (o *UpdateFreeipaToAwsImdsV1Default) Code() int {
	return o._statusCode
}

func (o *UpdateFreeipaToAwsImdsV1Default) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/updateFreeipaToAwsImdsV1][%d] updateFreeipaToAwsImdsV1 default %s", o._statusCode, payload)
}

func (o *UpdateFreeipaToAwsImdsV1Default) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/updateFreeipaToAwsImdsV1][%d] updateFreeipaToAwsImdsV1 default %s", o._statusCode, payload)
}

func (o *UpdateFreeipaToAwsImdsV1Default) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateFreeipaToAwsImdsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
