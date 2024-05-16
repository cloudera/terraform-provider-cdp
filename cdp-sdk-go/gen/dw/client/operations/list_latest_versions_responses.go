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

// ListLatestVersionsReader is a Reader for the ListLatestVersions structure.
type ListLatestVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLatestVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLatestVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListLatestVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListLatestVersionsOK creates a ListLatestVersionsOK with default headers values
func NewListLatestVersionsOK() *ListLatestVersionsOK {
	return &ListLatestVersionsOK{}
}

/*
ListLatestVersionsOK describes a response with status code 200, with default header values.

successful operation
*/
type ListLatestVersionsOK struct {
	Payload *models.ListLatestVersionsResponse
}

// IsSuccess returns true when this list latest versions o k response has a 2xx status code
func (o *ListLatestVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list latest versions o k response has a 3xx status code
func (o *ListLatestVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list latest versions o k response has a 4xx status code
func (o *ListLatestVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list latest versions o k response has a 5xx status code
func (o *ListLatestVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list latest versions o k response a status code equal to that given
func (o *ListLatestVersionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list latest versions o k response
func (o *ListLatestVersionsOK) Code() int {
	return 200
}

func (o *ListLatestVersionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listLatestVersions][%d] listLatestVersionsOK %s", 200, payload)
}

func (o *ListLatestVersionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listLatestVersions][%d] listLatestVersionsOK %s", 200, payload)
}

func (o *ListLatestVersionsOK) GetPayload() *models.ListLatestVersionsResponse {
	return o.Payload
}

func (o *ListLatestVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListLatestVersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLatestVersionsDefault creates a ListLatestVersionsDefault with default headers values
func NewListLatestVersionsDefault(code int) *ListLatestVersionsDefault {
	return &ListLatestVersionsDefault{
		_statusCode: code,
	}
}

/*
ListLatestVersionsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListLatestVersionsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list latest versions default response has a 2xx status code
func (o *ListLatestVersionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list latest versions default response has a 3xx status code
func (o *ListLatestVersionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list latest versions default response has a 4xx status code
func (o *ListLatestVersionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list latest versions default response has a 5xx status code
func (o *ListLatestVersionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list latest versions default response a status code equal to that given
func (o *ListLatestVersionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list latest versions default response
func (o *ListLatestVersionsDefault) Code() int {
	return o._statusCode
}

func (o *ListLatestVersionsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listLatestVersions][%d] listLatestVersions default %s", o._statusCode, payload)
}

func (o *ListLatestVersionsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/listLatestVersions][%d] listLatestVersions default %s", o._statusCode, payload)
}

func (o *ListLatestVersionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListLatestVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
