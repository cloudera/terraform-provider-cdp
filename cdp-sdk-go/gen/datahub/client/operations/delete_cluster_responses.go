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

// DeleteClusterReader is a Reader for the DeleteCluster structure.
type DeleteClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteClusterOK creates a DeleteClusterOK with default headers values
func NewDeleteClusterOK() *DeleteClusterOK {
	return &DeleteClusterOK{}
}

/*
DeleteClusterOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DeleteClusterOK struct {
	Payload *models.DeleteClusterResponse
}

// IsSuccess returns true when this delete cluster o k response has a 2xx status code
func (o *DeleteClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete cluster o k response has a 3xx status code
func (o *DeleteClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete cluster o k response has a 4xx status code
func (o *DeleteClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete cluster o k response has a 5xx status code
func (o *DeleteClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete cluster o k response a status code equal to that given
func (o *DeleteClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete cluster o k response
func (o *DeleteClusterOK) Code() int {
	return 200
}

func (o *DeleteClusterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/deleteCluster][%d] deleteClusterOK %s", 200, payload)
}

func (o *DeleteClusterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/deleteCluster][%d] deleteClusterOK %s", 200, payload)
}

func (o *DeleteClusterOK) GetPayload() *models.DeleteClusterResponse {
	return o.Payload
}

func (o *DeleteClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterDefault creates a DeleteClusterDefault with default headers values
func NewDeleteClusterDefault(code int) *DeleteClusterDefault {
	return &DeleteClusterDefault{
		_statusCode: code,
	}
}

/*
DeleteClusterDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DeleteClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete cluster default response has a 2xx status code
func (o *DeleteClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete cluster default response has a 3xx status code
func (o *DeleteClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete cluster default response has a 4xx status code
func (o *DeleteClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete cluster default response has a 5xx status code
func (o *DeleteClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete cluster default response a status code equal to that given
func (o *DeleteClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete cluster default response
func (o *DeleteClusterDefault) Code() int {
	return o._statusCode
}

func (o *DeleteClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/deleteCluster][%d] deleteCluster default %s", o._statusCode, payload)
}

func (o *DeleteClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/deleteCluster][%d] deleteCluster default %s", o._statusCode, payload)
}

func (o *DeleteClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
