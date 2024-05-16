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

// CreatePrivateClusterReader is a Reader for the CreatePrivateCluster structure.
type CreatePrivateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePrivateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePrivateClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreatePrivateClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreatePrivateClusterOK creates a CreatePrivateClusterOK with default headers values
func NewCreatePrivateClusterOK() *CreatePrivateClusterOK {
	return &CreatePrivateClusterOK{}
}

/*
CreatePrivateClusterOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type CreatePrivateClusterOK struct {
	Payload *models.CreatePrivateClusterResponse
}

// IsSuccess returns true when this create private cluster o k response has a 2xx status code
func (o *CreatePrivateClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create private cluster o k response has a 3xx status code
func (o *CreatePrivateClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create private cluster o k response has a 4xx status code
func (o *CreatePrivateClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create private cluster o k response has a 5xx status code
func (o *CreatePrivateClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create private cluster o k response a status code equal to that given
func (o *CreatePrivateClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create private cluster o k response
func (o *CreatePrivateClusterOK) Code() int {
	return 200
}

func (o *CreatePrivateClusterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/createPrivateCluster][%d] createPrivateClusterOK %s", 200, payload)
}

func (o *CreatePrivateClusterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/createPrivateCluster][%d] createPrivateClusterOK %s", 200, payload)
}

func (o *CreatePrivateClusterOK) GetPayload() *models.CreatePrivateClusterResponse {
	return o.Payload
}

func (o *CreatePrivateClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreatePrivateClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePrivateClusterDefault creates a CreatePrivateClusterDefault with default headers values
func NewCreatePrivateClusterDefault(code int) *CreatePrivateClusterDefault {
	return &CreatePrivateClusterDefault{
		_statusCode: code,
	}
}

/*
CreatePrivateClusterDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type CreatePrivateClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create private cluster default response has a 2xx status code
func (o *CreatePrivateClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create private cluster default response has a 3xx status code
func (o *CreatePrivateClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create private cluster default response has a 4xx status code
func (o *CreatePrivateClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create private cluster default response has a 5xx status code
func (o *CreatePrivateClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create private cluster default response a status code equal to that given
func (o *CreatePrivateClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create private cluster default response
func (o *CreatePrivateClusterDefault) Code() int {
	return o._statusCode
}

func (o *CreatePrivateClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/createPrivateCluster][%d] createPrivateCluster default %s", o._statusCode, payload)
}

func (o *CreatePrivateClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/createPrivateCluster][%d] createPrivateCluster default %s", o._statusCode, payload)
}

func (o *CreatePrivateClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreatePrivateClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
