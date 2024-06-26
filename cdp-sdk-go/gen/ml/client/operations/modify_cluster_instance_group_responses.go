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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// ModifyClusterInstanceGroupReader is a Reader for the ModifyClusterInstanceGroup structure.
type ModifyClusterInstanceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyClusterInstanceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModifyClusterInstanceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewModifyClusterInstanceGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModifyClusterInstanceGroupOK creates a ModifyClusterInstanceGroupOK with default headers values
func NewModifyClusterInstanceGroupOK() *ModifyClusterInstanceGroupOK {
	return &ModifyClusterInstanceGroupOK{}
}

/*
ModifyClusterInstanceGroupOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ModifyClusterInstanceGroupOK struct {
	Payload models.ModifyClusterInstanceGroupResponse
}

// IsSuccess returns true when this modify cluster instance group o k response has a 2xx status code
func (o *ModifyClusterInstanceGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modify cluster instance group o k response has a 3xx status code
func (o *ModifyClusterInstanceGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify cluster instance group o k response has a 4xx status code
func (o *ModifyClusterInstanceGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify cluster instance group o k response has a 5xx status code
func (o *ModifyClusterInstanceGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this modify cluster instance group o k response a status code equal to that given
func (o *ModifyClusterInstanceGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the modify cluster instance group o k response
func (o *ModifyClusterInstanceGroupOK) Code() int {
	return 200
}

func (o *ModifyClusterInstanceGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/modifyClusterInstanceGroup][%d] modifyClusterInstanceGroupOK %s", 200, payload)
}

func (o *ModifyClusterInstanceGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/modifyClusterInstanceGroup][%d] modifyClusterInstanceGroupOK %s", 200, payload)
}

func (o *ModifyClusterInstanceGroupOK) GetPayload() models.ModifyClusterInstanceGroupResponse {
	return o.Payload
}

func (o *ModifyClusterInstanceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyClusterInstanceGroupDefault creates a ModifyClusterInstanceGroupDefault with default headers values
func NewModifyClusterInstanceGroupDefault(code int) *ModifyClusterInstanceGroupDefault {
	return &ModifyClusterInstanceGroupDefault{
		_statusCode: code,
	}
}

/*
ModifyClusterInstanceGroupDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ModifyClusterInstanceGroupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this modify cluster instance group default response has a 2xx status code
func (o *ModifyClusterInstanceGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this modify cluster instance group default response has a 3xx status code
func (o *ModifyClusterInstanceGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this modify cluster instance group default response has a 4xx status code
func (o *ModifyClusterInstanceGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this modify cluster instance group default response has a 5xx status code
func (o *ModifyClusterInstanceGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this modify cluster instance group default response a status code equal to that given
func (o *ModifyClusterInstanceGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the modify cluster instance group default response
func (o *ModifyClusterInstanceGroupDefault) Code() int {
	return o._statusCode
}

func (o *ModifyClusterInstanceGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/modifyClusterInstanceGroup][%d] modifyClusterInstanceGroup default %s", o._statusCode, payload)
}

func (o *ModifyClusterInstanceGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/modifyClusterInstanceGroup][%d] modifyClusterInstanceGroup default %s", o._statusCode, payload)
}

func (o *ModifyClusterInstanceGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ModifyClusterInstanceGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
