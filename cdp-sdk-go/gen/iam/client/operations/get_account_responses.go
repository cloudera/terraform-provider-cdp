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

// GetAccountReader is a Reader for the GetAccount structure.
type GetAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAccountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAccountOK creates a GetAccountOK with default headers values
func NewGetAccountOK() *GetAccountOK {
	return &GetAccountOK{}
}

/*
GetAccountOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetAccountOK struct {
	Payload *models.GetAccountResponse
}

// IsSuccess returns true when this get account o k response has a 2xx status code
func (o *GetAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get account o k response has a 3xx status code
func (o *GetAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get account o k response has a 4xx status code
func (o *GetAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get account o k response has a 5xx status code
func (o *GetAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get account o k response a status code equal to that given
func (o *GetAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get account o k response
func (o *GetAccountOK) Code() int {
	return 200
}

func (o *GetAccountOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/getAccount][%d] getAccountOK %s", 200, payload)
}

func (o *GetAccountOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/getAccount][%d] getAccountOK %s", 200, payload)
}

func (o *GetAccountOK) GetPayload() *models.GetAccountResponse {
	return o.Payload
}

func (o *GetAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAccountResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountDefault creates a GetAccountDefault with default headers values
func NewGetAccountDefault(code int) *GetAccountDefault {
	return &GetAccountDefault{
		_statusCode: code,
	}
}

/*
GetAccountDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetAccountDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get account default response has a 2xx status code
func (o *GetAccountDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get account default response has a 3xx status code
func (o *GetAccountDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get account default response has a 4xx status code
func (o *GetAccountDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get account default response has a 5xx status code
func (o *GetAccountDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get account default response a status code equal to that given
func (o *GetAccountDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get account default response
func (o *GetAccountDefault) Code() int {
	return o._statusCode
}

func (o *GetAccountDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/getAccount][%d] getAccount default %s", o._statusCode, payload)
}

func (o *GetAccountDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/getAccount][%d] getAccount default %s", o._statusCode, payload)
}

func (o *GetAccountDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
