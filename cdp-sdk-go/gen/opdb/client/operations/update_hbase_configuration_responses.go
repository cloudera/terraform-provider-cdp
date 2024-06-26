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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

// UpdateHbaseConfigurationReader is a Reader for the UpdateHbaseConfiguration structure.
type UpdateHbaseConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateHbaseConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateHbaseConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateHbaseConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateHbaseConfigurationOK creates a UpdateHbaseConfigurationOK with default headers values
func NewUpdateHbaseConfigurationOK() *UpdateHbaseConfigurationOK {
	return &UpdateHbaseConfigurationOK{}
}

/*
UpdateHbaseConfigurationOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateHbaseConfigurationOK struct {
	Payload *models.UpdateHbaseConfigurationResponse
}

// IsSuccess returns true when this update hbase configuration o k response has a 2xx status code
func (o *UpdateHbaseConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update hbase configuration o k response has a 3xx status code
func (o *UpdateHbaseConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update hbase configuration o k response has a 4xx status code
func (o *UpdateHbaseConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update hbase configuration o k response has a 5xx status code
func (o *UpdateHbaseConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update hbase configuration o k response a status code equal to that given
func (o *UpdateHbaseConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update hbase configuration o k response
func (o *UpdateHbaseConfigurationOK) Code() int {
	return 200
}

func (o *UpdateHbaseConfigurationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/updateHbaseConfiguration][%d] updateHbaseConfigurationOK %s", 200, payload)
}

func (o *UpdateHbaseConfigurationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/updateHbaseConfiguration][%d] updateHbaseConfigurationOK %s", 200, payload)
}

func (o *UpdateHbaseConfigurationOK) GetPayload() *models.UpdateHbaseConfigurationResponse {
	return o.Payload
}

func (o *UpdateHbaseConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateHbaseConfigurationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateHbaseConfigurationDefault creates a UpdateHbaseConfigurationDefault with default headers values
func NewUpdateHbaseConfigurationDefault(code int) *UpdateHbaseConfigurationDefault {
	return &UpdateHbaseConfigurationDefault{
		_statusCode: code,
	}
}

/*
UpdateHbaseConfigurationDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateHbaseConfigurationDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update hbase configuration default response has a 2xx status code
func (o *UpdateHbaseConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update hbase configuration default response has a 3xx status code
func (o *UpdateHbaseConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update hbase configuration default response has a 4xx status code
func (o *UpdateHbaseConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update hbase configuration default response has a 5xx status code
func (o *UpdateHbaseConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update hbase configuration default response a status code equal to that given
func (o *UpdateHbaseConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update hbase configuration default response
func (o *UpdateHbaseConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateHbaseConfigurationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/updateHbaseConfiguration][%d] updateHbaseConfiguration default %s", o._statusCode, payload)
}

func (o *UpdateHbaseConfigurationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/updateHbaseConfiguration][%d] updateHbaseConfiguration default %s", o._statusCode, payload)
}

func (o *UpdateHbaseConfigurationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateHbaseConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
