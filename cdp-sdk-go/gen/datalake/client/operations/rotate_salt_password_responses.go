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

// RotateSaltPasswordReader is a Reader for the RotateSaltPassword structure.
type RotateSaltPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RotateSaltPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRotateSaltPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRotateSaltPasswordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRotateSaltPasswordOK creates a RotateSaltPasswordOK with default headers values
func NewRotateSaltPasswordOK() *RotateSaltPasswordOK {
	return &RotateSaltPasswordOK{}
}

/*
RotateSaltPasswordOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type RotateSaltPasswordOK struct {
	Payload *models.RotateSaltPasswordResponse
}

// IsSuccess returns true when this rotate salt password o k response has a 2xx status code
func (o *RotateSaltPasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rotate salt password o k response has a 3xx status code
func (o *RotateSaltPasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rotate salt password o k response has a 4xx status code
func (o *RotateSaltPasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rotate salt password o k response has a 5xx status code
func (o *RotateSaltPasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rotate salt password o k response a status code equal to that given
func (o *RotateSaltPasswordOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rotate salt password o k response
func (o *RotateSaltPasswordOK) Code() int {
	return 200
}

func (o *RotateSaltPasswordOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/rotateSaltPassword][%d] rotateSaltPasswordOK %s", 200, payload)
}

func (o *RotateSaltPasswordOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/rotateSaltPassword][%d] rotateSaltPasswordOK %s", 200, payload)
}

func (o *RotateSaltPasswordOK) GetPayload() *models.RotateSaltPasswordResponse {
	return o.Payload
}

func (o *RotateSaltPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RotateSaltPasswordResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRotateSaltPasswordDefault creates a RotateSaltPasswordDefault with default headers values
func NewRotateSaltPasswordDefault(code int) *RotateSaltPasswordDefault {
	return &RotateSaltPasswordDefault{
		_statusCode: code,
	}
}

/*
RotateSaltPasswordDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type RotateSaltPasswordDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this rotate salt password default response has a 2xx status code
func (o *RotateSaltPasswordDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rotate salt password default response has a 3xx status code
func (o *RotateSaltPasswordDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rotate salt password default response has a 4xx status code
func (o *RotateSaltPasswordDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rotate salt password default response has a 5xx status code
func (o *RotateSaltPasswordDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rotate salt password default response a status code equal to that given
func (o *RotateSaltPasswordDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the rotate salt password default response
func (o *RotateSaltPasswordDefault) Code() int {
	return o._statusCode
}

func (o *RotateSaltPasswordDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/rotateSaltPassword][%d] rotateSaltPassword default %s", o._statusCode, payload)
}

func (o *RotateSaltPasswordDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/datalake/rotateSaltPassword][%d] rotateSaltPassword default %s", o._statusCode, payload)
}

func (o *RotateSaltPasswordDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RotateSaltPasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
