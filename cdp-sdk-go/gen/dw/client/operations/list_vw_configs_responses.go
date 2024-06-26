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

// ListVwConfigsReader is a Reader for the ListVwConfigs structure.
type ListVwConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVwConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVwConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVwConfigsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVwConfigsOK creates a ListVwConfigsOK with default headers values
func NewListVwConfigsOK() *ListVwConfigsOK {
	return &ListVwConfigsOK{}
}

/*
ListVwConfigsOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListVwConfigsOK struct {
	Payload *models.ListVwConfigsResponse
}

// IsSuccess returns true when this list vw configs o k response has a 2xx status code
func (o *ListVwConfigsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list vw configs o k response has a 3xx status code
func (o *ListVwConfigsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list vw configs o k response has a 4xx status code
func (o *ListVwConfigsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list vw configs o k response has a 5xx status code
func (o *ListVwConfigsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list vw configs o k response a status code equal to that given
func (o *ListVwConfigsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list vw configs o k response
func (o *ListVwConfigsOK) Code() int {
	return 200
}

func (o *ListVwConfigsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listVwConfigs][%d] listVwConfigsOK %s", 200, payload)
}

func (o *ListVwConfigsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listVwConfigs][%d] listVwConfigsOK %s", 200, payload)
}

func (o *ListVwConfigsOK) GetPayload() *models.ListVwConfigsResponse {
	return o.Payload
}

func (o *ListVwConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListVwConfigsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVwConfigsDefault creates a ListVwConfigsDefault with default headers values
func NewListVwConfigsDefault(code int) *ListVwConfigsDefault {
	return &ListVwConfigsDefault{
		_statusCode: code,
	}
}

/*
ListVwConfigsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListVwConfigsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list vw configs default response has a 2xx status code
func (o *ListVwConfigsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list vw configs default response has a 3xx status code
func (o *ListVwConfigsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list vw configs default response has a 4xx status code
func (o *ListVwConfigsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list vw configs default response has a 5xx status code
func (o *ListVwConfigsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list vw configs default response a status code equal to that given
func (o *ListVwConfigsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list vw configs default response
func (o *ListVwConfigsDefault) Code() int {
	return o._statusCode
}

func (o *ListVwConfigsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listVwConfigs][%d] listVwConfigs default %s", o._statusCode, payload)
}

func (o *ListVwConfigsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listVwConfigs][%d] listVwConfigs default %s", o._statusCode, payload)
}

func (o *ListVwConfigsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListVwConfigsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
