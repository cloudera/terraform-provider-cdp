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

// ValidateAzureCloudStorageReader is a Reader for the ValidateAzureCloudStorage structure.
type ValidateAzureCloudStorageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateAzureCloudStorageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateAzureCloudStorageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewValidateAzureCloudStorageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidateAzureCloudStorageOK creates a ValidateAzureCloudStorageOK with default headers values
func NewValidateAzureCloudStorageOK() *ValidateAzureCloudStorageOK {
	return &ValidateAzureCloudStorageOK{}
}

/*
ValidateAzureCloudStorageOK describes a response with status code 200, with default header values.

Azure cloud storage validation result for Data Lake.
*/
type ValidateAzureCloudStorageOK struct {
	Payload *models.ValidateAzureCloudStorageResponse
}

// IsSuccess returns true when this validate azure cloud storage o k response has a 2xx status code
func (o *ValidateAzureCloudStorageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this validate azure cloud storage o k response has a 3xx status code
func (o *ValidateAzureCloudStorageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this validate azure cloud storage o k response has a 4xx status code
func (o *ValidateAzureCloudStorageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this validate azure cloud storage o k response has a 5xx status code
func (o *ValidateAzureCloudStorageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this validate azure cloud storage o k response a status code equal to that given
func (o *ValidateAzureCloudStorageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the validate azure cloud storage o k response
func (o *ValidateAzureCloudStorageOK) Code() int {
	return 200
}

func (o *ValidateAzureCloudStorageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/validateAzureCloudStorage][%d] validateAzureCloudStorageOK %s", 200, payload)
}

func (o *ValidateAzureCloudStorageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/validateAzureCloudStorage][%d] validateAzureCloudStorageOK %s", 200, payload)
}

func (o *ValidateAzureCloudStorageOK) GetPayload() *models.ValidateAzureCloudStorageResponse {
	return o.Payload
}

func (o *ValidateAzureCloudStorageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidateAzureCloudStorageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateAzureCloudStorageDefault creates a ValidateAzureCloudStorageDefault with default headers values
func NewValidateAzureCloudStorageDefault(code int) *ValidateAzureCloudStorageDefault {
	return &ValidateAzureCloudStorageDefault{
		_statusCode: code,
	}
}

/*
ValidateAzureCloudStorageDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ValidateAzureCloudStorageDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this validate azure cloud storage default response has a 2xx status code
func (o *ValidateAzureCloudStorageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this validate azure cloud storage default response has a 3xx status code
func (o *ValidateAzureCloudStorageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this validate azure cloud storage default response has a 4xx status code
func (o *ValidateAzureCloudStorageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this validate azure cloud storage default response has a 5xx status code
func (o *ValidateAzureCloudStorageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this validate azure cloud storage default response a status code equal to that given
func (o *ValidateAzureCloudStorageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the validate azure cloud storage default response
func (o *ValidateAzureCloudStorageDefault) Code() int {
	return o._statusCode
}

func (o *ValidateAzureCloudStorageDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/validateAzureCloudStorage][%d] validateAzureCloudStorage default %s", o._statusCode, payload)
}

func (o *ValidateAzureCloudStorageDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/validateAzureCloudStorage][%d] validateAzureCloudStorage default %s", o._statusCode, payload)
}

func (o *ValidateAzureCloudStorageDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ValidateAzureCloudStorageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
