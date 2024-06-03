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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/de/models"
)

// GetServiceInitLogsReader is a Reader for the GetServiceInitLogs structure.
type GetServiceInitLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceInitLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceInitLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServiceInitLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceInitLogsOK creates a GetServiceInitLogsOK with default headers values
func NewGetServiceInitLogsOK() *GetServiceInitLogsOK {
	return &GetServiceInitLogsOK{}
}

/*
GetServiceInitLogsOK describes a response with status code 200, with default header values.

Response object for Get Service Init Logs command.
*/
type GetServiceInitLogsOK struct {
	Payload *models.GetServiceInitLogsResponse
}

// IsSuccess returns true when this get service init logs o k response has a 2xx status code
func (o *GetServiceInitLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get service init logs o k response has a 3xx status code
func (o *GetServiceInitLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service init logs o k response has a 4xx status code
func (o *GetServiceInitLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get service init logs o k response has a 5xx status code
func (o *GetServiceInitLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get service init logs o k response a status code equal to that given
func (o *GetServiceInitLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get service init logs o k response
func (o *GetServiceInitLogsOK) Code() int {
	return 200
}

func (o *GetServiceInitLogsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/getServiceInitLogs][%d] getServiceInitLogsOK %s", 200, payload)
}

func (o *GetServiceInitLogsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/getServiceInitLogs][%d] getServiceInitLogsOK %s", 200, payload)
}

func (o *GetServiceInitLogsOK) GetPayload() *models.GetServiceInitLogsResponse {
	return o.Payload
}

func (o *GetServiceInitLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetServiceInitLogsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceInitLogsDefault creates a GetServiceInitLogsDefault with default headers values
func NewGetServiceInitLogsDefault(code int) *GetServiceInitLogsDefault {
	return &GetServiceInitLogsDefault{
		_statusCode: code,
	}
}

/*
GetServiceInitLogsDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetServiceInitLogsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get service init logs default response has a 2xx status code
func (o *GetServiceInitLogsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get service init logs default response has a 3xx status code
func (o *GetServiceInitLogsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get service init logs default response has a 4xx status code
func (o *GetServiceInitLogsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get service init logs default response has a 5xx status code
func (o *GetServiceInitLogsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get service init logs default response a status code equal to that given
func (o *GetServiceInitLogsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get service init logs default response
func (o *GetServiceInitLogsDefault) Code() int {
	return o._statusCode
}

func (o *GetServiceInitLogsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/getServiceInitLogs][%d] getServiceInitLogs default %s", o._statusCode, payload)
}

func (o *GetServiceInitLogsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/de/getServiceInitLogs][%d] getServiceInitLogs default %s", o._statusCode, payload)
}

func (o *GetServiceInitLogsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetServiceInitLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
