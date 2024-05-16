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

// PrepareUpgradeDatabaseReader is a Reader for the PrepareUpgradeDatabase structure.
type PrepareUpgradeDatabaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrepareUpgradeDatabaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrepareUpgradeDatabaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPrepareUpgradeDatabaseDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPrepareUpgradeDatabaseOK creates a PrepareUpgradeDatabaseOK with default headers values
func NewPrepareUpgradeDatabaseOK() *PrepareUpgradeDatabaseOK {
	return &PrepareUpgradeDatabaseOK{}
}

/*
PrepareUpgradeDatabaseOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type PrepareUpgradeDatabaseOK struct {
	Payload *models.PrepareUpgradeDatabaseResponse
}

// IsSuccess returns true when this prepare upgrade database o k response has a 2xx status code
func (o *PrepareUpgradeDatabaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this prepare upgrade database o k response has a 3xx status code
func (o *PrepareUpgradeDatabaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this prepare upgrade database o k response has a 4xx status code
func (o *PrepareUpgradeDatabaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this prepare upgrade database o k response has a 5xx status code
func (o *PrepareUpgradeDatabaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this prepare upgrade database o k response a status code equal to that given
func (o *PrepareUpgradeDatabaseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the prepare upgrade database o k response
func (o *PrepareUpgradeDatabaseOK) Code() int {
	return 200
}

func (o *PrepareUpgradeDatabaseOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/prepareUpgradeDatabase][%d] prepareUpgradeDatabaseOK %s", 200, payload)
}

func (o *PrepareUpgradeDatabaseOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/prepareUpgradeDatabase][%d] prepareUpgradeDatabaseOK %s", 200, payload)
}

func (o *PrepareUpgradeDatabaseOK) GetPayload() *models.PrepareUpgradeDatabaseResponse {
	return o.Payload
}

func (o *PrepareUpgradeDatabaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrepareUpgradeDatabaseResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrepareUpgradeDatabaseDefault creates a PrepareUpgradeDatabaseDefault with default headers values
func NewPrepareUpgradeDatabaseDefault(code int) *PrepareUpgradeDatabaseDefault {
	return &PrepareUpgradeDatabaseDefault{
		_statusCode: code,
	}
}

/*
PrepareUpgradeDatabaseDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type PrepareUpgradeDatabaseDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this prepare upgrade database default response has a 2xx status code
func (o *PrepareUpgradeDatabaseDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this prepare upgrade database default response has a 3xx status code
func (o *PrepareUpgradeDatabaseDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this prepare upgrade database default response has a 4xx status code
func (o *PrepareUpgradeDatabaseDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this prepare upgrade database default response has a 5xx status code
func (o *PrepareUpgradeDatabaseDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this prepare upgrade database default response a status code equal to that given
func (o *PrepareUpgradeDatabaseDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the prepare upgrade database default response
func (o *PrepareUpgradeDatabaseDefault) Code() int {
	return o._statusCode
}

func (o *PrepareUpgradeDatabaseDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/prepareUpgradeDatabase][%d] prepareUpgradeDatabase default %s", o._statusCode, payload)
}

func (o *PrepareUpgradeDatabaseDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/opdb/prepareUpgradeDatabase][%d] prepareUpgradeDatabase default %s", o._statusCode, payload)
}

func (o *PrepareUpgradeDatabaseDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PrepareUpgradeDatabaseDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
