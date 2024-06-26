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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/de/models"
)

// EnableServiceReader is a Reader for the EnableService structure.
type EnableServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEnableServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnableServiceOK creates a EnableServiceOK with default headers values
func NewEnableServiceOK() *EnableServiceOK {
	return &EnableServiceOK{}
}

/*
EnableServiceOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type EnableServiceOK struct {
	Payload *models.EnableServiceResponse
}

// IsSuccess returns true when this enable service o k response has a 2xx status code
func (o *EnableServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enable service o k response has a 3xx status code
func (o *EnableServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable service o k response has a 4xx status code
func (o *EnableServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable service o k response has a 5xx status code
func (o *EnableServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enable service o k response a status code equal to that given
func (o *EnableServiceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the enable service o k response
func (o *EnableServiceOK) Code() int {
	return 200
}

func (o *EnableServiceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/enableService][%d] enableServiceOK %s", 200, payload)
}

func (o *EnableServiceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/enableService][%d] enableServiceOK %s", 200, payload)
}

func (o *EnableServiceOK) GetPayload() *models.EnableServiceResponse {
	return o.Payload
}

func (o *EnableServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnableServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableServiceDefault creates a EnableServiceDefault with default headers values
func NewEnableServiceDefault(code int) *EnableServiceDefault {
	return &EnableServiceDefault{
		_statusCode: code,
	}
}

/*
EnableServiceDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type EnableServiceDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this enable service default response has a 2xx status code
func (o *EnableServiceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this enable service default response has a 3xx status code
func (o *EnableServiceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this enable service default response has a 4xx status code
func (o *EnableServiceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this enable service default response has a 5xx status code
func (o *EnableServiceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this enable service default response a status code equal to that given
func (o *EnableServiceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the enable service default response
func (o *EnableServiceDefault) Code() int {
	return o._statusCode
}

func (o *EnableServiceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/enableService][%d] enableService default %s", o._statusCode, payload)
}

func (o *EnableServiceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/enableService][%d] enableService default %s", o._statusCode, payload)
}

func (o *EnableServiceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnableServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
