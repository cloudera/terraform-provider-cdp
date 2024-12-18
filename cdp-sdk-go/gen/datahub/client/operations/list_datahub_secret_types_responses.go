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

// ListDatahubSecretTypesReader is a Reader for the ListDatahubSecretTypes structure.
type ListDatahubSecretTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDatahubSecretTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDatahubSecretTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDatahubSecretTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDatahubSecretTypesOK creates a ListDatahubSecretTypesOK with default headers values
func NewListDatahubSecretTypesOK() *ListDatahubSecretTypesOK {
	return &ListDatahubSecretTypesOK{}
}

/*
ListDatahubSecretTypesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListDatahubSecretTypesOK struct {
	Payload *models.ListDatahubSecretTypesResponse
}

// IsSuccess returns true when this list datahub secret types o k response has a 2xx status code
func (o *ListDatahubSecretTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list datahub secret types o k response has a 3xx status code
func (o *ListDatahubSecretTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list datahub secret types o k response has a 4xx status code
func (o *ListDatahubSecretTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list datahub secret types o k response has a 5xx status code
func (o *ListDatahubSecretTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list datahub secret types o k response a status code equal to that given
func (o *ListDatahubSecretTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list datahub secret types o k response
func (o *ListDatahubSecretTypesOK) Code() int {
	return 200
}

func (o *ListDatahubSecretTypesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/listDatahubSecretTypes][%d] listDatahubSecretTypesOK %s", 200, payload)
}

func (o *ListDatahubSecretTypesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/listDatahubSecretTypes][%d] listDatahubSecretTypesOK %s", 200, payload)
}

func (o *ListDatahubSecretTypesOK) GetPayload() *models.ListDatahubSecretTypesResponse {
	return o.Payload
}

func (o *ListDatahubSecretTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListDatahubSecretTypesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDatahubSecretTypesDefault creates a ListDatahubSecretTypesDefault with default headers values
func NewListDatahubSecretTypesDefault(code int) *ListDatahubSecretTypesDefault {
	return &ListDatahubSecretTypesDefault{
		_statusCode: code,
	}
}

/*
ListDatahubSecretTypesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListDatahubSecretTypesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list datahub secret types default response has a 2xx status code
func (o *ListDatahubSecretTypesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list datahub secret types default response has a 3xx status code
func (o *ListDatahubSecretTypesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list datahub secret types default response has a 4xx status code
func (o *ListDatahubSecretTypesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list datahub secret types default response has a 5xx status code
func (o *ListDatahubSecretTypesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list datahub secret types default response a status code equal to that given
func (o *ListDatahubSecretTypesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list datahub secret types default response
func (o *ListDatahubSecretTypesDefault) Code() int {
	return o._statusCode
}

func (o *ListDatahubSecretTypesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/listDatahubSecretTypes][%d] listDatahubSecretTypes default %s", o._statusCode, payload)
}

func (o *ListDatahubSecretTypesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/listDatahubSecretTypes][%d] listDatahubSecretTypes default %s", o._statusCode, payload)
}

func (o *ListDatahubSecretTypesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListDatahubSecretTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
