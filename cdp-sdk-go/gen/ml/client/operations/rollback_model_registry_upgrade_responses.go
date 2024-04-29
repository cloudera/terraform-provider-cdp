// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// RollbackModelRegistryUpgradeReader is a Reader for the RollbackModelRegistryUpgrade structure.
type RollbackModelRegistryUpgradeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RollbackModelRegistryUpgradeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRollbackModelRegistryUpgradeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRollbackModelRegistryUpgradeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRollbackModelRegistryUpgradeOK creates a RollbackModelRegistryUpgradeOK with default headers values
func NewRollbackModelRegistryUpgradeOK() *RollbackModelRegistryUpgradeOK {
	return &RollbackModelRegistryUpgradeOK{}
}

/*
RollbackModelRegistryUpgradeOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type RollbackModelRegistryUpgradeOK struct {
	Payload models.RollbackModelRegistryUpgradeResponse
}

// IsSuccess returns true when this rollback model registry upgrade o k response has a 2xx status code
func (o *RollbackModelRegistryUpgradeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rollback model registry upgrade o k response has a 3xx status code
func (o *RollbackModelRegistryUpgradeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rollback model registry upgrade o k response has a 4xx status code
func (o *RollbackModelRegistryUpgradeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rollback model registry upgrade o k response has a 5xx status code
func (o *RollbackModelRegistryUpgradeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rollback model registry upgrade o k response a status code equal to that given
func (o *RollbackModelRegistryUpgradeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rollback model registry upgrade o k response
func (o *RollbackModelRegistryUpgradeOK) Code() int {
	return 200
}

func (o *RollbackModelRegistryUpgradeOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/rollbackModelRegistryUpgrade][%d] rollbackModelRegistryUpgradeOK  %+v", 200, o.Payload)
}

func (o *RollbackModelRegistryUpgradeOK) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/rollbackModelRegistryUpgrade][%d] rollbackModelRegistryUpgradeOK  %+v", 200, o.Payload)
}

func (o *RollbackModelRegistryUpgradeOK) GetPayload() models.RollbackModelRegistryUpgradeResponse {
	return o.Payload
}

func (o *RollbackModelRegistryUpgradeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRollbackModelRegistryUpgradeDefault creates a RollbackModelRegistryUpgradeDefault with default headers values
func NewRollbackModelRegistryUpgradeDefault(code int) *RollbackModelRegistryUpgradeDefault {
	return &RollbackModelRegistryUpgradeDefault{
		_statusCode: code,
	}
}

/*
RollbackModelRegistryUpgradeDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type RollbackModelRegistryUpgradeDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this rollback model registry upgrade default response has a 2xx status code
func (o *RollbackModelRegistryUpgradeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rollback model registry upgrade default response has a 3xx status code
func (o *RollbackModelRegistryUpgradeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rollback model registry upgrade default response has a 4xx status code
func (o *RollbackModelRegistryUpgradeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rollback model registry upgrade default response has a 5xx status code
func (o *RollbackModelRegistryUpgradeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rollback model registry upgrade default response a status code equal to that given
func (o *RollbackModelRegistryUpgradeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the rollback model registry upgrade default response
func (o *RollbackModelRegistryUpgradeDefault) Code() int {
	return o._statusCode
}

func (o *RollbackModelRegistryUpgradeDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/rollbackModelRegistryUpgrade][%d] rollbackModelRegistryUpgrade default  %+v", o._statusCode, o.Payload)
}

func (o *RollbackModelRegistryUpgradeDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/ml/rollbackModelRegistryUpgrade][%d] rollbackModelRegistryUpgrade default  %+v", o._statusCode, o.Payload)
}

func (o *RollbackModelRegistryUpgradeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RollbackModelRegistryUpgradeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
