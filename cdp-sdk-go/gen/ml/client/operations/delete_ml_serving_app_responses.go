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

// DeleteMlServingAppReader is a Reader for the DeleteMlServingApp structure.
type DeleteMlServingAppReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMlServingAppReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMlServingAppOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteMlServingAppDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMlServingAppOK creates a DeleteMlServingAppOK with default headers values
func NewDeleteMlServingAppOK() *DeleteMlServingAppOK {
	return &DeleteMlServingAppOK{}
}

/*
DeleteMlServingAppOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DeleteMlServingAppOK struct {
	Payload models.DeleteMlServingAppResponse
}

// IsSuccess returns true when this delete ml serving app o k response has a 2xx status code
func (o *DeleteMlServingAppOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete ml serving app o k response has a 3xx status code
func (o *DeleteMlServingAppOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete ml serving app o k response has a 4xx status code
func (o *DeleteMlServingAppOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete ml serving app o k response has a 5xx status code
func (o *DeleteMlServingAppOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete ml serving app o k response a status code equal to that given
func (o *DeleteMlServingAppOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete ml serving app o k response
func (o *DeleteMlServingAppOK) Code() int {
	return 200
}

func (o *DeleteMlServingAppOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/deleteMlServingApp][%d] deleteMlServingAppOK %s", 200, payload)
}

func (o *DeleteMlServingAppOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/deleteMlServingApp][%d] deleteMlServingAppOK %s", 200, payload)
}

func (o *DeleteMlServingAppOK) GetPayload() models.DeleteMlServingAppResponse {
	return o.Payload
}

func (o *DeleteMlServingAppOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMlServingAppDefault creates a DeleteMlServingAppDefault with default headers values
func NewDeleteMlServingAppDefault(code int) *DeleteMlServingAppDefault {
	return &DeleteMlServingAppDefault{
		_statusCode: code,
	}
}

/*
DeleteMlServingAppDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DeleteMlServingAppDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete ml serving app default response has a 2xx status code
func (o *DeleteMlServingAppDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete ml serving app default response has a 3xx status code
func (o *DeleteMlServingAppDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete ml serving app default response has a 4xx status code
func (o *DeleteMlServingAppDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete ml serving app default response has a 5xx status code
func (o *DeleteMlServingAppDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete ml serving app default response a status code equal to that given
func (o *DeleteMlServingAppDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete ml serving app default response
func (o *DeleteMlServingAppDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMlServingAppDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/deleteMlServingApp][%d] deleteMlServingApp default %s", o._statusCode, payload)
}

func (o *DeleteMlServingAppDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/ml/deleteMlServingApp][%d] deleteMlServingApp default %s", o._statusCode, payload)
}

func (o *DeleteMlServingAppDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteMlServingAppDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
