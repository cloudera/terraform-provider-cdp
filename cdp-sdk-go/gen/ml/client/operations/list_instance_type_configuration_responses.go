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

// ListInstanceTypeConfigurationReader is a Reader for the ListInstanceTypeConfiguration structure.
type ListInstanceTypeConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListInstanceTypeConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListInstanceTypeConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListInstanceTypeConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListInstanceTypeConfigurationOK creates a ListInstanceTypeConfigurationOK with default headers values
func NewListInstanceTypeConfigurationOK() *ListInstanceTypeConfigurationOK {
	return &ListInstanceTypeConfigurationOK{}
}

/*
ListInstanceTypeConfigurationOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListInstanceTypeConfigurationOK struct {
	Payload *models.ListInstanceTypeConfigurationResponse
}

// IsSuccess returns true when this list instance type configuration o k response has a 2xx status code
func (o *ListInstanceTypeConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list instance type configuration o k response has a 3xx status code
func (o *ListInstanceTypeConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list instance type configuration o k response has a 4xx status code
func (o *ListInstanceTypeConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list instance type configuration o k response has a 5xx status code
func (o *ListInstanceTypeConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list instance type configuration o k response a status code equal to that given
func (o *ListInstanceTypeConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list instance type configuration o k response
func (o *ListInstanceTypeConfigurationOK) Code() int {
	return 200
}

func (o *ListInstanceTypeConfigurationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/listInstanceTypeConfiguration][%d] listInstanceTypeConfigurationOK %s", 200, payload)
}

func (o *ListInstanceTypeConfigurationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/listInstanceTypeConfiguration][%d] listInstanceTypeConfigurationOK %s", 200, payload)
}

func (o *ListInstanceTypeConfigurationOK) GetPayload() *models.ListInstanceTypeConfigurationResponse {
	return o.Payload
}

func (o *ListInstanceTypeConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListInstanceTypeConfigurationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInstanceTypeConfigurationDefault creates a ListInstanceTypeConfigurationDefault with default headers values
func NewListInstanceTypeConfigurationDefault(code int) *ListInstanceTypeConfigurationDefault {
	return &ListInstanceTypeConfigurationDefault{
		_statusCode: code,
	}
}

/*
ListInstanceTypeConfigurationDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListInstanceTypeConfigurationDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list instance type configuration default response has a 2xx status code
func (o *ListInstanceTypeConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list instance type configuration default response has a 3xx status code
func (o *ListInstanceTypeConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list instance type configuration default response has a 4xx status code
func (o *ListInstanceTypeConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list instance type configuration default response has a 5xx status code
func (o *ListInstanceTypeConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list instance type configuration default response a status code equal to that given
func (o *ListInstanceTypeConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list instance type configuration default response
func (o *ListInstanceTypeConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *ListInstanceTypeConfigurationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/listInstanceTypeConfiguration][%d] listInstanceTypeConfiguration default %s", o._statusCode, payload)
}

func (o *ListInstanceTypeConfigurationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/listInstanceTypeConfiguration][%d] listInstanceTypeConfiguration default %s", o._statusCode, payload)
}

func (o *ListInstanceTypeConfigurationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListInstanceTypeConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
