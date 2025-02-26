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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// RestartDatalakeInstancesReader is a Reader for the RestartDatalakeInstances structure.
type RestartDatalakeInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartDatalakeInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestartDatalakeInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestartDatalakeInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestartDatalakeInstancesOK creates a RestartDatalakeInstancesOK with default headers values
func NewRestartDatalakeInstancesOK() *RestartDatalakeInstancesOK {
	return &RestartDatalakeInstancesOK{}
}

/*
RestartDatalakeInstancesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type RestartDatalakeInstancesOK struct {
	Payload *models.RestartDatalakeInstancesResponse
}

// IsSuccess returns true when this restart datalake instances o k response has a 2xx status code
func (o *RestartDatalakeInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restart datalake instances o k response has a 3xx status code
func (o *RestartDatalakeInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restart datalake instances o k response has a 4xx status code
func (o *RestartDatalakeInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restart datalake instances o k response has a 5xx status code
func (o *RestartDatalakeInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this restart datalake instances o k response a status code equal to that given
func (o *RestartDatalakeInstancesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the restart datalake instances o k response
func (o *RestartDatalakeInstancesOK) Code() int {
	return 200
}

func (o *RestartDatalakeInstancesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/restartDatalakeInstances][%d] restartDatalakeInstancesOK %s", 200, payload)
}

func (o *RestartDatalakeInstancesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/restartDatalakeInstances][%d] restartDatalakeInstancesOK %s", 200, payload)
}

func (o *RestartDatalakeInstancesOK) GetPayload() *models.RestartDatalakeInstancesResponse {
	return o.Payload
}

func (o *RestartDatalakeInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestartDatalakeInstancesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartDatalakeInstancesDefault creates a RestartDatalakeInstancesDefault with default headers values
func NewRestartDatalakeInstancesDefault(code int) *RestartDatalakeInstancesDefault {
	return &RestartDatalakeInstancesDefault{
		_statusCode: code,
	}
}

/*
RestartDatalakeInstancesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type RestartDatalakeInstancesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this restart datalake instances default response has a 2xx status code
func (o *RestartDatalakeInstancesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this restart datalake instances default response has a 3xx status code
func (o *RestartDatalakeInstancesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this restart datalake instances default response has a 4xx status code
func (o *RestartDatalakeInstancesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this restart datalake instances default response has a 5xx status code
func (o *RestartDatalakeInstancesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this restart datalake instances default response a status code equal to that given
func (o *RestartDatalakeInstancesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the restart datalake instances default response
func (o *RestartDatalakeInstancesDefault) Code() int {
	return o._statusCode
}

func (o *RestartDatalakeInstancesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/restartDatalakeInstances][%d] restartDatalakeInstances default %s", o._statusCode, payload)
}

func (o *RestartDatalakeInstancesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/restartDatalakeInstances][%d] restartDatalakeInstances default %s", o._statusCode, payload)
}

func (o *RestartDatalakeInstancesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RestartDatalakeInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
