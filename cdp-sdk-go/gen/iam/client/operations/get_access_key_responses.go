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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// GetAccessKeyReader is a Reader for the GetAccessKey structure.
type GetAccessKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccessKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccessKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAccessKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAccessKeyOK creates a GetAccessKeyOK with default headers values
func NewGetAccessKeyOK() *GetAccessKeyOK {
	return &GetAccessKeyOK{}
}

/*
GetAccessKeyOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetAccessKeyOK struct {
	Payload *models.GetAccessKeyResponse
}

// IsSuccess returns true when this get access key o k response has a 2xx status code
func (o *GetAccessKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get access key o k response has a 3xx status code
func (o *GetAccessKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get access key o k response has a 4xx status code
func (o *GetAccessKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get access key o k response has a 5xx status code
func (o *GetAccessKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get access key o k response a status code equal to that given
func (o *GetAccessKeyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get access key o k response
func (o *GetAccessKeyOK) Code() int {
	return 200
}

func (o *GetAccessKeyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/getAccessKey][%d] getAccessKeyOK %s", 200, payload)
}

func (o *GetAccessKeyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/getAccessKey][%d] getAccessKeyOK %s", 200, payload)
}

func (o *GetAccessKeyOK) GetPayload() *models.GetAccessKeyResponse {
	return o.Payload
}

func (o *GetAccessKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAccessKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccessKeyDefault creates a GetAccessKeyDefault with default headers values
func NewGetAccessKeyDefault(code int) *GetAccessKeyDefault {
	return &GetAccessKeyDefault{
		_statusCode: code,
	}
}

/*
GetAccessKeyDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetAccessKeyDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get access key default response has a 2xx status code
func (o *GetAccessKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get access key default response has a 3xx status code
func (o *GetAccessKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get access key default response has a 4xx status code
func (o *GetAccessKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get access key default response has a 5xx status code
func (o *GetAccessKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get access key default response a status code equal to that given
func (o *GetAccessKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get access key default response
func (o *GetAccessKeyDefault) Code() int {
	return o._statusCode
}

func (o *GetAccessKeyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/getAccessKey][%d] getAccessKey default %s", o._statusCode, payload)
}

func (o *GetAccessKeyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/getAccessKey][%d] getAccessKey default %s", o._statusCode, payload)
}

func (o *GetAccessKeyDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccessKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
