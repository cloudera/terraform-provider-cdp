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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// ListModelRegistriesReader is a Reader for the ListModelRegistries structure.
type ListModelRegistriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListModelRegistriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListModelRegistriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListModelRegistriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListModelRegistriesOK creates a ListModelRegistriesOK with default headers values
func NewListModelRegistriesOK() *ListModelRegistriesOK {
	return &ListModelRegistriesOK{}
}

/*
ListModelRegistriesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListModelRegistriesOK struct {
	Payload *models.ListModelRegistriesResponse
}

// IsSuccess returns true when this list model registries o k response has a 2xx status code
func (o *ListModelRegistriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list model registries o k response has a 3xx status code
func (o *ListModelRegistriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list model registries o k response has a 4xx status code
func (o *ListModelRegistriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list model registries o k response has a 5xx status code
func (o *ListModelRegistriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list model registries o k response a status code equal to that given
func (o *ListModelRegistriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list model registries o k response
func (o *ListModelRegistriesOK) Code() int {
	return 200
}

func (o *ListModelRegistriesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/listModelRegistries][%d] listModelRegistriesOK %s", 200, payload)
}

func (o *ListModelRegistriesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/listModelRegistries][%d] listModelRegistriesOK %s", 200, payload)
}

func (o *ListModelRegistriesOK) GetPayload() *models.ListModelRegistriesResponse {
	return o.Payload
}

func (o *ListModelRegistriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListModelRegistriesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListModelRegistriesDefault creates a ListModelRegistriesDefault with default headers values
func NewListModelRegistriesDefault(code int) *ListModelRegistriesDefault {
	return &ListModelRegistriesDefault{
		_statusCode: code,
	}
}

/*
ListModelRegistriesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListModelRegistriesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list model registries default response has a 2xx status code
func (o *ListModelRegistriesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list model registries default response has a 3xx status code
func (o *ListModelRegistriesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list model registries default response has a 4xx status code
func (o *ListModelRegistriesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list model registries default response has a 5xx status code
func (o *ListModelRegistriesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list model registries default response a status code equal to that given
func (o *ListModelRegistriesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list model registries default response
func (o *ListModelRegistriesDefault) Code() int {
	return o._statusCode
}

func (o *ListModelRegistriesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/listModelRegistries][%d] listModelRegistries default %s", o._statusCode, payload)
}

func (o *ListModelRegistriesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/listModelRegistries][%d] listModelRegistries default %s", o._statusCode, payload)
}

func (o *ListModelRegistriesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListModelRegistriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
