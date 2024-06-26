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

// DetachFreeIpaRecipesReader is a Reader for the DetachFreeIpaRecipes structure.
type DetachFreeIpaRecipesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetachFreeIpaRecipesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetachFreeIpaRecipesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDetachFreeIpaRecipesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDetachFreeIpaRecipesOK creates a DetachFreeIpaRecipesOK with default headers values
func NewDetachFreeIpaRecipesOK() *DetachFreeIpaRecipesOK {
	return &DetachFreeIpaRecipesOK{}
}

/*
DetachFreeIpaRecipesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DetachFreeIpaRecipesOK struct {
	Payload models.DetachFreeIpaRecipesResponse
}

// IsSuccess returns true when this detach free ipa recipes o k response has a 2xx status code
func (o *DetachFreeIpaRecipesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this detach free ipa recipes o k response has a 3xx status code
func (o *DetachFreeIpaRecipesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this detach free ipa recipes o k response has a 4xx status code
func (o *DetachFreeIpaRecipesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this detach free ipa recipes o k response has a 5xx status code
func (o *DetachFreeIpaRecipesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this detach free ipa recipes o k response a status code equal to that given
func (o *DetachFreeIpaRecipesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the detach free ipa recipes o k response
func (o *DetachFreeIpaRecipesOK) Code() int {
	return 200
}

func (o *DetachFreeIpaRecipesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/detachFreeIpaRecipes][%d] detachFreeIpaRecipesOK %s", 200, payload)
}

func (o *DetachFreeIpaRecipesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/detachFreeIpaRecipes][%d] detachFreeIpaRecipesOK %s", 200, payload)
}

func (o *DetachFreeIpaRecipesOK) GetPayload() models.DetachFreeIpaRecipesResponse {
	return o.Payload
}

func (o *DetachFreeIpaRecipesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetachFreeIpaRecipesDefault creates a DetachFreeIpaRecipesDefault with default headers values
func NewDetachFreeIpaRecipesDefault(code int) *DetachFreeIpaRecipesDefault {
	return &DetachFreeIpaRecipesDefault{
		_statusCode: code,
	}
}

/*
DetachFreeIpaRecipesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DetachFreeIpaRecipesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this detach free ipa recipes default response has a 2xx status code
func (o *DetachFreeIpaRecipesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this detach free ipa recipes default response has a 3xx status code
func (o *DetachFreeIpaRecipesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this detach free ipa recipes default response has a 4xx status code
func (o *DetachFreeIpaRecipesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this detach free ipa recipes default response has a 5xx status code
func (o *DetachFreeIpaRecipesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this detach free ipa recipes default response a status code equal to that given
func (o *DetachFreeIpaRecipesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the detach free ipa recipes default response
func (o *DetachFreeIpaRecipesDefault) Code() int {
	return o._statusCode
}

func (o *DetachFreeIpaRecipesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/detachFreeIpaRecipes][%d] detachFreeIpaRecipes default %s", o._statusCode, payload)
}

func (o *DetachFreeIpaRecipesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/detachFreeIpaRecipes][%d] detachFreeIpaRecipes default %s", o._statusCode, payload)
}

func (o *DetachFreeIpaRecipesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DetachFreeIpaRecipesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
