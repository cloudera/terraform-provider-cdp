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

// ListRecipesReader is a Reader for the ListRecipes structure.
type ListRecipesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRecipesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRecipesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListRecipesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRecipesOK creates a ListRecipesOK with default headers values
func NewListRecipesOK() *ListRecipesOK {
	return &ListRecipesOK{}
}

/*
ListRecipesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListRecipesOK struct {
	Payload *models.ListRecipesResponse
}

// IsSuccess returns true when this list recipes o k response has a 2xx status code
func (o *ListRecipesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list recipes o k response has a 3xx status code
func (o *ListRecipesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list recipes o k response has a 4xx status code
func (o *ListRecipesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list recipes o k response has a 5xx status code
func (o *ListRecipesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list recipes o k response a status code equal to that given
func (o *ListRecipesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list recipes o k response
func (o *ListRecipesOK) Code() int {
	return 200
}

func (o *ListRecipesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/listRecipes][%d] listRecipesOK %s", 200, payload)
}

func (o *ListRecipesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/listRecipes][%d] listRecipesOK %s", 200, payload)
}

func (o *ListRecipesOK) GetPayload() *models.ListRecipesResponse {
	return o.Payload
}

func (o *ListRecipesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListRecipesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRecipesDefault creates a ListRecipesDefault with default headers values
func NewListRecipesDefault(code int) *ListRecipesDefault {
	return &ListRecipesDefault{
		_statusCode: code,
	}
}

/*
ListRecipesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListRecipesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list recipes default response has a 2xx status code
func (o *ListRecipesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list recipes default response has a 3xx status code
func (o *ListRecipesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list recipes default response has a 4xx status code
func (o *ListRecipesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list recipes default response has a 5xx status code
func (o *ListRecipesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list recipes default response a status code equal to that given
func (o *ListRecipesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list recipes default response
func (o *ListRecipesDefault) Code() int {
	return o._statusCode
}

func (o *ListRecipesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/listRecipes][%d] listRecipes default %s", o._statusCode, payload)
}

func (o *ListRecipesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/listRecipes][%d] listRecipes default %s", o._statusCode, payload)
}

func (o *ListRecipesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListRecipesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
