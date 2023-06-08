// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// RestoreClusterReader is a Reader for the RestoreCluster structure.
type RestoreClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestoreClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestoreClusterOK creates a RestoreClusterOK with default headers values
func NewRestoreClusterOK() *RestoreClusterOK {
	return &RestoreClusterOK{}
}

/*
RestoreClusterOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type RestoreClusterOK struct {
	Payload *models.RestoreClusterResponse
}

// IsSuccess returns true when this restore cluster o k response has a 2xx status code
func (o *RestoreClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restore cluster o k response has a 3xx status code
func (o *RestoreClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore cluster o k response has a 4xx status code
func (o *RestoreClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restore cluster o k response has a 5xx status code
func (o *RestoreClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this restore cluster o k response a status code equal to that given
func (o *RestoreClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the restore cluster o k response
func (o *RestoreClusterOK) Code() int {
	return 200
}

func (o *RestoreClusterOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/restoreCluster][%d] restoreClusterOK  %+v", 200, o.Payload)
}

func (o *RestoreClusterOK) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/restoreCluster][%d] restoreClusterOK  %+v", 200, o.Payload)
}

func (o *RestoreClusterOK) GetPayload() *models.RestoreClusterResponse {
	return o.Payload
}

func (o *RestoreClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestoreClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreClusterDefault creates a RestoreClusterDefault with default headers values
func NewRestoreClusterDefault(code int) *RestoreClusterDefault {
	return &RestoreClusterDefault{
		_statusCode: code,
	}
}

/*
RestoreClusterDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type RestoreClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this restore cluster default response has a 2xx status code
func (o *RestoreClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this restore cluster default response has a 3xx status code
func (o *RestoreClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this restore cluster default response has a 4xx status code
func (o *RestoreClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this restore cluster default response has a 5xx status code
func (o *RestoreClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this restore cluster default response a status code equal to that given
func (o *RestoreClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the restore cluster default response
func (o *RestoreClusterDefault) Code() int {
	return o._statusCode
}

func (o *RestoreClusterDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/restoreCluster][%d] restoreCluster default  %+v", o._statusCode, o.Payload)
}

func (o *RestoreClusterDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/restoreCluster][%d] restoreCluster default  %+v", o._statusCode, o.Payload)
}

func (o *RestoreClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RestoreClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
