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

// ValidateAwsCloudStorageReader is a Reader for the ValidateAwsCloudStorage structure.
type ValidateAwsCloudStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateAwsCloudStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateAwsCloudStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewValidateAwsCloudStorageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidateAwsCloudStorageOK creates a ValidateAwsCloudStorageOK with default headers values
func NewValidateAwsCloudStorageOK() *ValidateAwsCloudStorageOK {
	return &ValidateAwsCloudStorageOK{}
}

/*
ValidateAwsCloudStorageOK describes a response with status code 200, with default header values.

AWS cloud storage validation result for Data Lake.
*/
type ValidateAwsCloudStorageOK struct {
	Payload *models.ValidateAwsCloudStorageResponse
}

// IsSuccess returns true when this validate aws cloud storage o k response has a 2xx status code
func (o *ValidateAwsCloudStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate aws cloud storage o k response has a 3xx status code
func (o *ValidateAwsCloudStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate aws cloud storage o k response has a 4xx status code
func (o *ValidateAwsCloudStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate aws cloud storage o k response has a 5xx status code
func (o *ValidateAwsCloudStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate aws cloud storage o k response a status code equal to that given
func (o *ValidateAwsCloudStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the validate aws cloud storage o k response
func (o *ValidateAwsCloudStorageOK) Code() int {
	return 200
}

func (o *ValidateAwsCloudStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/validateAwsCloudStorage][%d] validateAwsCloudStorageOK %s", 200, payload)
}

func (o *ValidateAwsCloudStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/validateAwsCloudStorage][%d] validateAwsCloudStorageOK %s", 200, payload)
}

func (o *ValidateAwsCloudStorageOK) GetPayload() *models.ValidateAwsCloudStorageResponse {
	return o.Payload
}

func (o *ValidateAwsCloudStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidateAwsCloudStorageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateAwsCloudStorageDefault creates a ValidateAwsCloudStorageDefault with default headers values
func NewValidateAwsCloudStorageDefault(code int) *ValidateAwsCloudStorageDefault {
	return &ValidateAwsCloudStorageDefault{
		_statusCode: code,
	}
}

/*
ValidateAwsCloudStorageDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ValidateAwsCloudStorageDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this validate aws cloud storage default response has a 2xx status code
func (o *ValidateAwsCloudStorageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this validate aws cloud storage default response has a 3xx status code
func (o *ValidateAwsCloudStorageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this validate aws cloud storage default response has a 4xx status code
func (o *ValidateAwsCloudStorageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this validate aws cloud storage default response has a 5xx status code
func (o *ValidateAwsCloudStorageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this validate aws cloud storage default response a status code equal to that given
func (o *ValidateAwsCloudStorageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the validate aws cloud storage default response
func (o *ValidateAwsCloudStorageDefault) Code() int {
	return o._statusCode
}

func (o *ValidateAwsCloudStorageDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/validateAwsCloudStorage][%d] validateAwsCloudStorage default %s", o._statusCode, payload)
}

func (o *ValidateAwsCloudStorageDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/validateAwsCloudStorage][%d] validateAwsCloudStorage default %s", o._statusCode, payload)
}

func (o *ValidateAwsCloudStorageDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateAwsCloudStorageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
