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

// InitializeAWSComputeClusterReader is a Reader for the InitializeAWSComputeCluster structure.
type InitializeAWSComputeClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitializeAWSComputeClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInitializeAWSComputeClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInitializeAWSComputeClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInitializeAWSComputeClusterOK creates a InitializeAWSComputeClusterOK with default headers values
func NewInitializeAWSComputeClusterOK() *InitializeAWSComputeClusterOK {
	return &InitializeAWSComputeClusterOK{}
}

/*
InitializeAWSComputeClusterOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type InitializeAWSComputeClusterOK struct {
	Payload models.InitializeAWSComputeClusterResponse
}

// IsSuccess returns true when this initialize a w s compute cluster o k response has a 2xx status code
func (o *InitializeAWSComputeClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this initialize a w s compute cluster o k response has a 3xx status code
func (o *InitializeAWSComputeClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this initialize a w s compute cluster o k response has a 4xx status code
func (o *InitializeAWSComputeClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this initialize a w s compute cluster o k response has a 5xx status code
func (o *InitializeAWSComputeClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this initialize a w s compute cluster o k response a status code equal to that given
func (o *InitializeAWSComputeClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the initialize a w s compute cluster o k response
func (o *InitializeAWSComputeClusterOK) Code() int {
	return 200
}

func (o *InitializeAWSComputeClusterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/initializeAWSComputeCluster][%d] initializeAWSComputeClusterOK %s", 200, payload)
}

func (o *InitializeAWSComputeClusterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/initializeAWSComputeCluster][%d] initializeAWSComputeClusterOK %s", 200, payload)
}

func (o *InitializeAWSComputeClusterOK) GetPayload() models.InitializeAWSComputeClusterResponse {
	return o.Payload
}

func (o *InitializeAWSComputeClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInitializeAWSComputeClusterDefault creates a InitializeAWSComputeClusterDefault with default headers values
func NewInitializeAWSComputeClusterDefault(code int) *InitializeAWSComputeClusterDefault {
	return &InitializeAWSComputeClusterDefault{
		_statusCode: code,
	}
}

/*
InitializeAWSComputeClusterDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type InitializeAWSComputeClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this initialize a w s compute cluster default response has a 2xx status code
func (o *InitializeAWSComputeClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this initialize a w s compute cluster default response has a 3xx status code
func (o *InitializeAWSComputeClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this initialize a w s compute cluster default response has a 4xx status code
func (o *InitializeAWSComputeClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this initialize a w s compute cluster default response has a 5xx status code
func (o *InitializeAWSComputeClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this initialize a w s compute cluster default response a status code equal to that given
func (o *InitializeAWSComputeClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the initialize a w s compute cluster default response
func (o *InitializeAWSComputeClusterDefault) Code() int {
	return o._statusCode
}

func (o *InitializeAWSComputeClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/initializeAWSComputeCluster][%d] initializeAWSComputeCluster default %s", o._statusCode, payload)
}

func (o *InitializeAWSComputeClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/initializeAWSComputeCluster][%d] initializeAWSComputeCluster default %s", o._statusCode, payload)
}

func (o *InitializeAWSComputeClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *InitializeAWSComputeClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
