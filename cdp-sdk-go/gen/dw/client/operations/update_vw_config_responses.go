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

// UpdateVwConfigReader is a Reader for the UpdateVwConfig structure.
type UpdateVwConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVwConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVwConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateVwConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateVwConfigOK creates a UpdateVwConfigOK with default headers values
func NewUpdateVwConfigOK() *UpdateVwConfigOK {
	return &UpdateVwConfigOK{}
}

/*
UpdateVwConfigOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateVwConfigOK struct {
	Payload models.UpdateVwConfigResponse
}

// IsSuccess returns true when this update vw config o k response has a 2xx status code
func (o *UpdateVwConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update vw config o k response has a 3xx status code
func (o *UpdateVwConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update vw config o k response has a 4xx status code
func (o *UpdateVwConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update vw config o k response has a 5xx status code
func (o *UpdateVwConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update vw config o k response a status code equal to that given
func (o *UpdateVwConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update vw config o k response
func (o *UpdateVwConfigOK) Code() int {
	return 200
}

func (o *UpdateVwConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/updateVwConfig][%d] updateVwConfigOK %s", 200, payload)
}

func (o *UpdateVwConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/updateVwConfig][%d] updateVwConfigOK %s", 200, payload)
}

func (o *UpdateVwConfigOK) GetPayload() models.UpdateVwConfigResponse {
	return o.Payload
}

func (o *UpdateVwConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVwConfigDefault creates a UpdateVwConfigDefault with default headers values
func NewUpdateVwConfigDefault(code int) *UpdateVwConfigDefault {
	return &UpdateVwConfigDefault{
		_statusCode: code,
	}
}

/*
UpdateVwConfigDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateVwConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update vw config default response has a 2xx status code
func (o *UpdateVwConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update vw config default response has a 3xx status code
func (o *UpdateVwConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update vw config default response has a 4xx status code
func (o *UpdateVwConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update vw config default response has a 5xx status code
func (o *UpdateVwConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update vw config default response a status code equal to that given
func (o *UpdateVwConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update vw config default response
func (o *UpdateVwConfigDefault) Code() int {
	return o._statusCode
}

func (o *UpdateVwConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/updateVwConfig][%d] updateVwConfig default %s", o._statusCode, payload)
}

func (o *UpdateVwConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/updateVwConfig][%d] updateVwConfig default %s", o._statusCode, payload)
}

func (o *UpdateVwConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateVwConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
