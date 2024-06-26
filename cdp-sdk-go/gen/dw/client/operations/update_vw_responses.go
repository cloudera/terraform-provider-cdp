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

// UpdateVwReader is a Reader for the UpdateVw structure.
type UpdateVwReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVwReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVwOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateVwDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateVwOK creates a UpdateVwOK with default headers values
func NewUpdateVwOK() *UpdateVwOK {
	return &UpdateVwOK{}
}

/*
UpdateVwOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateVwOK struct {
	Payload models.UpdateVwResponse
}

// IsSuccess returns true when this update vw o k response has a 2xx status code
func (o *UpdateVwOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update vw o k response has a 3xx status code
func (o *UpdateVwOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update vw o k response has a 4xx status code
func (o *UpdateVwOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update vw o k response has a 5xx status code
func (o *UpdateVwOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update vw o k response a status code equal to that given
func (o *UpdateVwOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update vw o k response
func (o *UpdateVwOK) Code() int {
	return 200
}

func (o *UpdateVwOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/updateVw][%d] updateVwOK %s", 200, payload)
}

func (o *UpdateVwOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/updateVw][%d] updateVwOK %s", 200, payload)
}

func (o *UpdateVwOK) GetPayload() models.UpdateVwResponse {
	return o.Payload
}

func (o *UpdateVwOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVwDefault creates a UpdateVwDefault with default headers values
func NewUpdateVwDefault(code int) *UpdateVwDefault {
	return &UpdateVwDefault{
		_statusCode: code,
	}
}

/*
UpdateVwDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateVwDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update vw default response has a 2xx status code
func (o *UpdateVwDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update vw default response has a 3xx status code
func (o *UpdateVwDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update vw default response has a 4xx status code
func (o *UpdateVwDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update vw default response has a 5xx status code
func (o *UpdateVwDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update vw default response a status code equal to that given
func (o *UpdateVwDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update vw default response
func (o *UpdateVwDefault) Code() int {
	return o._statusCode
}

func (o *UpdateVwDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/updateVw][%d] updateVw default %s", o._statusCode, payload)
}

func (o *UpdateVwDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/updateVw][%d] updateVw default %s", o._statusCode, payload)
}

func (o *UpdateVwDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateVwDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
