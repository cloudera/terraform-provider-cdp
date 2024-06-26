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

// UpdateAwsDiskEncryptionParametersReader is a Reader for the UpdateAwsDiskEncryptionParameters structure.
type UpdateAwsDiskEncryptionParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAwsDiskEncryptionParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAwsDiskEncryptionParametersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateAwsDiskEncryptionParametersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAwsDiskEncryptionParametersOK creates a UpdateAwsDiskEncryptionParametersOK with default headers values
func NewUpdateAwsDiskEncryptionParametersOK() *UpdateAwsDiskEncryptionParametersOK {
	return &UpdateAwsDiskEncryptionParametersOK{}
}

/*
UpdateAwsDiskEncryptionParametersOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateAwsDiskEncryptionParametersOK struct {
	Payload *models.UpdateAwsDiskEncryptionParametersResponse
}

// IsSuccess returns true when this update aws disk encryption parameters o k response has a 2xx status code
func (o *UpdateAwsDiskEncryptionParametersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update aws disk encryption parameters o k response has a 3xx status code
func (o *UpdateAwsDiskEncryptionParametersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update aws disk encryption parameters o k response has a 4xx status code
func (o *UpdateAwsDiskEncryptionParametersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update aws disk encryption parameters o k response has a 5xx status code
func (o *UpdateAwsDiskEncryptionParametersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update aws disk encryption parameters o k response a status code equal to that given
func (o *UpdateAwsDiskEncryptionParametersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update aws disk encryption parameters o k response
func (o *UpdateAwsDiskEncryptionParametersOK) Code() int {
	return 200
}

func (o *UpdateAwsDiskEncryptionParametersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/updateAwsDiskEncryptionParameters][%d] updateAwsDiskEncryptionParametersOK %s", 200, payload)
}

func (o *UpdateAwsDiskEncryptionParametersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/updateAwsDiskEncryptionParameters][%d] updateAwsDiskEncryptionParametersOK %s", 200, payload)
}

func (o *UpdateAwsDiskEncryptionParametersOK) GetPayload() *models.UpdateAwsDiskEncryptionParametersResponse {
	return o.Payload
}

func (o *UpdateAwsDiskEncryptionParametersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateAwsDiskEncryptionParametersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAwsDiskEncryptionParametersDefault creates a UpdateAwsDiskEncryptionParametersDefault with default headers values
func NewUpdateAwsDiskEncryptionParametersDefault(code int) *UpdateAwsDiskEncryptionParametersDefault {
	return &UpdateAwsDiskEncryptionParametersDefault{
		_statusCode: code,
	}
}

/*
UpdateAwsDiskEncryptionParametersDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateAwsDiskEncryptionParametersDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update aws disk encryption parameters default response has a 2xx status code
func (o *UpdateAwsDiskEncryptionParametersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update aws disk encryption parameters default response has a 3xx status code
func (o *UpdateAwsDiskEncryptionParametersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update aws disk encryption parameters default response has a 4xx status code
func (o *UpdateAwsDiskEncryptionParametersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update aws disk encryption parameters default response has a 5xx status code
func (o *UpdateAwsDiskEncryptionParametersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update aws disk encryption parameters default response a status code equal to that given
func (o *UpdateAwsDiskEncryptionParametersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update aws disk encryption parameters default response
func (o *UpdateAwsDiskEncryptionParametersDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAwsDiskEncryptionParametersDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/updateAwsDiskEncryptionParameters][%d] updateAwsDiskEncryptionParameters default %s", o._statusCode, payload)
}

func (o *UpdateAwsDiskEncryptionParametersDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/updateAwsDiskEncryptionParameters][%d] updateAwsDiskEncryptionParameters default %s", o._statusCode, payload)
}

func (o *UpdateAwsDiskEncryptionParametersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateAwsDiskEncryptionParametersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
