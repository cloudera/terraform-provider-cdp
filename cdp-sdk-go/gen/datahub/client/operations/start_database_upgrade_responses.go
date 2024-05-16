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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// StartDatabaseUpgradeReader is a Reader for the StartDatabaseUpgrade structure.
type StartDatabaseUpgradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartDatabaseUpgradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartDatabaseUpgradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartDatabaseUpgradeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartDatabaseUpgradeOK creates a StartDatabaseUpgradeOK with default headers values
func NewStartDatabaseUpgradeOK() *StartDatabaseUpgradeOK {
	return &StartDatabaseUpgradeOK{}
}

/*
StartDatabaseUpgradeOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type StartDatabaseUpgradeOK struct {
	Payload *models.StartDatabaseUpgradeResponse
}

// IsSuccess returns true when this start database upgrade o k response has a 2xx status code
func (o *StartDatabaseUpgradeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start database upgrade o k response has a 3xx status code
func (o *StartDatabaseUpgradeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start database upgrade o k response has a 4xx status code
func (o *StartDatabaseUpgradeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this start database upgrade o k response has a 5xx status code
func (o *StartDatabaseUpgradeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this start database upgrade o k response a status code equal to that given
func (o *StartDatabaseUpgradeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the start database upgrade o k response
func (o *StartDatabaseUpgradeOK) Code() int {
	return 200
}

func (o *StartDatabaseUpgradeOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/startDatabaseUpgrade][%d] startDatabaseUpgradeOK %s", 200, payload)
}

func (o *StartDatabaseUpgradeOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/startDatabaseUpgrade][%d] startDatabaseUpgradeOK %s", 200, payload)
}

func (o *StartDatabaseUpgradeOK) GetPayload() *models.StartDatabaseUpgradeResponse {
	return o.Payload
}

func (o *StartDatabaseUpgradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StartDatabaseUpgradeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDatabaseUpgradeDefault creates a StartDatabaseUpgradeDefault with default headers values
func NewStartDatabaseUpgradeDefault(code int) *StartDatabaseUpgradeDefault {
	return &StartDatabaseUpgradeDefault{
		_statusCode: code,
	}
}

/*
StartDatabaseUpgradeDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type StartDatabaseUpgradeDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this start database upgrade default response has a 2xx status code
func (o *StartDatabaseUpgradeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this start database upgrade default response has a 3xx status code
func (o *StartDatabaseUpgradeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this start database upgrade default response has a 4xx status code
func (o *StartDatabaseUpgradeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this start database upgrade default response has a 5xx status code
func (o *StartDatabaseUpgradeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this start database upgrade default response a status code equal to that given
func (o *StartDatabaseUpgradeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the start database upgrade default response
func (o *StartDatabaseUpgradeDefault) Code() int {
	return o._statusCode
}

func (o *StartDatabaseUpgradeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/startDatabaseUpgrade][%d] startDatabaseUpgrade default %s", o._statusCode, payload)
}

func (o *StartDatabaseUpgradeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/startDatabaseUpgrade][%d] startDatabaseUpgrade default %s", o._statusCode, payload)
}

func (o *StartDatabaseUpgradeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *StartDatabaseUpgradeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
