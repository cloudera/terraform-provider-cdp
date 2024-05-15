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

// CreateProxyConfigReader is a Reader for the CreateProxyConfig structure.
type CreateProxyConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProxyConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProxyConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateProxyConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateProxyConfigOK creates a CreateProxyConfigOK with default headers values
func NewCreateProxyConfigOK() *CreateProxyConfigOK {
	return &CreateProxyConfigOK{}
}

/*
CreateProxyConfigOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CreateProxyConfigOK struct {
	Payload *models.CreateProxyConfigResponse
}

// IsSuccess returns true when this create proxy config o k response has a 2xx status code
func (o *CreateProxyConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create proxy config o k response has a 3xx status code
func (o *CreateProxyConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create proxy config o k response has a 4xx status code
func (o *CreateProxyConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create proxy config o k response has a 5xx status code
func (o *CreateProxyConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create proxy config o k response a status code equal to that given
func (o *CreateProxyConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create proxy config o k response
func (o *CreateProxyConfigOK) Code() int {
	return 200
}

func (o *CreateProxyConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/createProxyConfig][%d] createProxyConfigOK %s", 200, payload)
}

func (o *CreateProxyConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/createProxyConfig][%d] createProxyConfigOK %s", 200, payload)
}

func (o *CreateProxyConfigOK) GetPayload() *models.CreateProxyConfigResponse {
	return o.Payload
}

func (o *CreateProxyConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateProxyConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProxyConfigDefault creates a CreateProxyConfigDefault with default headers values
func NewCreateProxyConfigDefault(code int) *CreateProxyConfigDefault {
	return &CreateProxyConfigDefault{
		_statusCode: code,
	}
}

/*
CreateProxyConfigDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CreateProxyConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create proxy config default response has a 2xx status code
func (o *CreateProxyConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create proxy config default response has a 3xx status code
func (o *CreateProxyConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create proxy config default response has a 4xx status code
func (o *CreateProxyConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create proxy config default response has a 5xx status code
func (o *CreateProxyConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create proxy config default response a status code equal to that given
func (o *CreateProxyConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create proxy config default response
func (o *CreateProxyConfigDefault) Code() int {
	return o._statusCode
}

func (o *CreateProxyConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/createProxyConfig][%d] createProxyConfig default %s", o._statusCode, payload)
}

func (o *CreateProxyConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/createProxyConfig][%d] createProxyConfig default %s", o._statusCode, payload)
}

func (o *CreateProxyConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateProxyConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
