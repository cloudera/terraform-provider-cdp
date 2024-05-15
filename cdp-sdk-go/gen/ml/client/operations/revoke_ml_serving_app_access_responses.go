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

// RevokeMlServingAppAccessReader is a Reader for the RevokeMlServingAppAccess structure.
type RevokeMlServingAppAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeMlServingAppAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeMlServingAppAccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRevokeMlServingAppAccessDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRevokeMlServingAppAccessOK creates a RevokeMlServingAppAccessOK with default headers values
func NewRevokeMlServingAppAccessOK() *RevokeMlServingAppAccessOK {
	return &RevokeMlServingAppAccessOK{}
}

/*
RevokeMlServingAppAccessOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type RevokeMlServingAppAccessOK struct {
	Payload models.RevokeMlServingAppAccessResponse
}

// IsSuccess returns true when this revoke ml serving app access o k response has a 2xx status code
func (o *RevokeMlServingAppAccessOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this revoke ml serving app access o k response has a 3xx status code
func (o *RevokeMlServingAppAccessOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this revoke ml serving app access o k response has a 4xx status code
func (o *RevokeMlServingAppAccessOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this revoke ml serving app access o k response has a 5xx status code
func (o *RevokeMlServingAppAccessOK) IsServerError() bool {
	return false
}

// IsCode returns true when this revoke ml serving app access o k response a status code equal to that given
func (o *RevokeMlServingAppAccessOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the revoke ml serving app access o k response
func (o *RevokeMlServingAppAccessOK) Code() int {
	return 200
}

func (o *RevokeMlServingAppAccessOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/revokeMlServingAppAccess][%d] revokeMlServingAppAccessOK %s", 200, payload)
}

func (o *RevokeMlServingAppAccessOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/revokeMlServingAppAccess][%d] revokeMlServingAppAccessOK %s", 200, payload)
}

func (o *RevokeMlServingAppAccessOK) GetPayload() models.RevokeMlServingAppAccessResponse {
	return o.Payload
}

func (o *RevokeMlServingAppAccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeMlServingAppAccessDefault creates a RevokeMlServingAppAccessDefault with default headers values
func NewRevokeMlServingAppAccessDefault(code int) *RevokeMlServingAppAccessDefault {
	return &RevokeMlServingAppAccessDefault{
		_statusCode: code,
	}
}

/*
RevokeMlServingAppAccessDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type RevokeMlServingAppAccessDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this revoke ml serving app access default response has a 2xx status code
func (o *RevokeMlServingAppAccessDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this revoke ml serving app access default response has a 3xx status code
func (o *RevokeMlServingAppAccessDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this revoke ml serving app access default response has a 4xx status code
func (o *RevokeMlServingAppAccessDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this revoke ml serving app access default response has a 5xx status code
func (o *RevokeMlServingAppAccessDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this revoke ml serving app access default response a status code equal to that given
func (o *RevokeMlServingAppAccessDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the revoke ml serving app access default response
func (o *RevokeMlServingAppAccessDefault) Code() int {
	return o._statusCode
}

func (o *RevokeMlServingAppAccessDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/revokeMlServingAppAccess][%d] revokeMlServingAppAccess default %s", o._statusCode, payload)
}

func (o *RevokeMlServingAppAccessDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/revokeMlServingAppAccess][%d] revokeMlServingAppAccess default %s", o._statusCode, payload)
}

func (o *RevokeMlServingAppAccessDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeMlServingAppAccessDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
