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

// DeleteVwReader is a Reader for the DeleteVw structure.
type DeleteVwReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVwReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVwOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteVwDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteVwOK creates a DeleteVwOK with default headers values
func NewDeleteVwOK() *DeleteVwOK {
	return &DeleteVwOK{}
}

/*
DeleteVwOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DeleteVwOK struct {
	Payload models.DeleteVwResponse
}

// IsSuccess returns true when this delete vw o k response has a 2xx status code
func (o *DeleteVwOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete vw o k response has a 3xx status code
func (o *DeleteVwOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete vw o k response has a 4xx status code
func (o *DeleteVwOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete vw o k response has a 5xx status code
func (o *DeleteVwOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete vw o k response a status code equal to that given
func (o *DeleteVwOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete vw o k response
func (o *DeleteVwOK) Code() int {
	return 200
}

func (o *DeleteVwOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/deleteVw][%d] deleteVwOK %s", 200, payload)
}

func (o *DeleteVwOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/deleteVw][%d] deleteVwOK %s", 200, payload)
}

func (o *DeleteVwOK) GetPayload() models.DeleteVwResponse {
	return o.Payload
}

func (o *DeleteVwOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVwDefault creates a DeleteVwDefault with default headers values
func NewDeleteVwDefault(code int) *DeleteVwDefault {
	return &DeleteVwDefault{
		_statusCode: code,
	}
}

/*
DeleteVwDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DeleteVwDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete vw default response has a 2xx status code
func (o *DeleteVwDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete vw default response has a 3xx status code
func (o *DeleteVwDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete vw default response has a 4xx status code
func (o *DeleteVwDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete vw default response has a 5xx status code
func (o *DeleteVwDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete vw default response a status code equal to that given
func (o *DeleteVwDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete vw default response
func (o *DeleteVwDefault) Code() int {
	return o._statusCode
}

func (o *DeleteVwDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/deleteVw][%d] deleteVw default %s", o._statusCode, payload)
}

func (o *DeleteVwDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/deleteVw][%d] deleteVw default %s", o._statusCode, payload)
}

func (o *DeleteVwDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteVwDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
