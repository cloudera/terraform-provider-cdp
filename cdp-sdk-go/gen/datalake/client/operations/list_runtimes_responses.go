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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// ListRuntimesReader is a Reader for the ListRuntimes structure.
type ListRuntimesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRuntimesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRuntimesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListRuntimesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRuntimesOK creates a ListRuntimesOK with default headers values
func NewListRuntimesOK() *ListRuntimesOK {
	return &ListRuntimesOK{}
}

/*
ListRuntimesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListRuntimesOK struct {
	Payload *models.ListRuntimesResponse
}

// IsSuccess returns true when this list runtimes o k response has a 2xx status code
func (o *ListRuntimesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list runtimes o k response has a 3xx status code
func (o *ListRuntimesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list runtimes o k response has a 4xx status code
func (o *ListRuntimesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list runtimes o k response has a 5xx status code
func (o *ListRuntimesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list runtimes o k response a status code equal to that given
func (o *ListRuntimesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list runtimes o k response
func (o *ListRuntimesOK) Code() int {
	return 200
}

func (o *ListRuntimesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/listRuntimes][%d] listRuntimesOK %s", 200, payload)
}

func (o *ListRuntimesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/listRuntimes][%d] listRuntimesOK %s", 200, payload)
}

func (o *ListRuntimesOK) GetPayload() *models.ListRuntimesResponse {
	return o.Payload
}

func (o *ListRuntimesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListRuntimesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRuntimesDefault creates a ListRuntimesDefault with default headers values
func NewListRuntimesDefault(code int) *ListRuntimesDefault {
	return &ListRuntimesDefault{
		_statusCode: code,
	}
}

/*
ListRuntimesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListRuntimesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list runtimes default response has a 2xx status code
func (o *ListRuntimesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list runtimes default response has a 3xx status code
func (o *ListRuntimesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list runtimes default response has a 4xx status code
func (o *ListRuntimesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list runtimes default response has a 5xx status code
func (o *ListRuntimesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list runtimes default response a status code equal to that given
func (o *ListRuntimesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list runtimes default response
func (o *ListRuntimesDefault) Code() int {
	return o._statusCode
}

func (o *ListRuntimesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/listRuntimes][%d] listRuntimes default %s", o._statusCode, payload)
}

func (o *ListRuntimesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/listRuntimes][%d] listRuntimes default %s", o._statusCode, payload)
}

func (o *ListRuntimesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListRuntimesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
