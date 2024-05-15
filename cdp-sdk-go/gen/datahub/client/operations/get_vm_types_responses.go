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

// GetVMTypesReader is a Reader for the GetVMTypes structure.
type GetVMTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVMTypesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVMTypesOK creates a GetVMTypesOK with default headers values
func NewGetVMTypesOK() *GetVMTypesOK {
	return &GetVMTypesOK{}
}

/*
GetVMTypesOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetVMTypesOK struct {
	Payload *models.GetVMTypesResponse
}

// IsSuccess returns true when this get Vm types o k response has a 2xx status code
func (o *GetVMTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Vm types o k response has a 3xx status code
func (o *GetVMTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Vm types o k response has a 4xx status code
func (o *GetVMTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Vm types o k response has a 5xx status code
func (o *GetVMTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Vm types o k response a status code equal to that given
func (o *GetVMTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Vm types o k response
func (o *GetVMTypesOK) Code() int {
	return 200
}

func (o *GetVMTypesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/getVmTypes][%d] getVmTypesOK %s", 200, payload)
}

func (o *GetVMTypesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/getVmTypes][%d] getVmTypesOK %s", 200, payload)
}

func (o *GetVMTypesOK) GetPayload() *models.GetVMTypesResponse {
	return o.Payload
}

func (o *GetVMTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetVMTypesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMTypesDefault creates a GetVMTypesDefault with default headers values
func NewGetVMTypesDefault(code int) *GetVMTypesDefault {
	return &GetVMTypesDefault{
		_statusCode: code,
	}
}

/*
GetVMTypesDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetVMTypesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get Vm types default response has a 2xx status code
func (o *GetVMTypesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get Vm types default response has a 3xx status code
func (o *GetVMTypesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get Vm types default response has a 4xx status code
func (o *GetVMTypesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get Vm types default response has a 5xx status code
func (o *GetVMTypesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get Vm types default response a status code equal to that given
func (o *GetVMTypesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get Vm types default response
func (o *GetVMTypesDefault) Code() int {
	return o._statusCode
}

func (o *GetVMTypesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/getVmTypes][%d] getVmTypes default %s", o._statusCode, payload)
}

func (o *GetVMTypesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datahub/getVmTypes][%d] getVmTypes default %s", o._statusCode, payload)
}

func (o *GetVMTypesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetVMTypesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
