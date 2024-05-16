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

// AddUserReader is a Reader for the AddUser structure.
type AddUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddUserOK creates a AddUserOK with default headers values
func NewAddUserOK() *AddUserOK {
	return &AddUserOK{}
}

/*
AddUserOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type AddUserOK struct {
	Payload models.AddUserResponse
}

// IsSuccess returns true when this add user o k response has a 2xx status code
func (o *AddUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add user o k response has a 3xx status code
func (o *AddUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add user o k response has a 4xx status code
func (o *AddUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add user o k response has a 5xx status code
func (o *AddUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add user o k response a status code equal to that given
func (o *AddUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add user o k response
func (o *AddUserOK) Code() int {
	return 200
}

func (o *AddUserOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/addUser][%d] addUserOK %s", 200, payload)
}

func (o *AddUserOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/addUser][%d] addUserOK %s", 200, payload)
}

func (o *AddUserOK) GetPayload() models.AddUserResponse {
	return o.Payload
}

func (o *AddUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddUserDefault creates a AddUserDefault with default headers values
func NewAddUserDefault(code int) *AddUserDefault {
	return &AddUserDefault{
		_statusCode: code,
	}
}

/*
AddUserDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type AddUserDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this add user default response has a 2xx status code
func (o *AddUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add user default response has a 3xx status code
func (o *AddUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add user default response has a 4xx status code
func (o *AddUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add user default response has a 5xx status code
func (o *AddUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add user default response a status code equal to that given
func (o *AddUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the add user default response
func (o *AddUserDefault) Code() int {
	return o._statusCode
}

func (o *AddUserDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/addUser][%d] addUser default %s", o._statusCode, payload)
}

func (o *AddUserDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/addUser][%d] addUser default %s", o._statusCode, payload)
}

func (o *AddUserDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
