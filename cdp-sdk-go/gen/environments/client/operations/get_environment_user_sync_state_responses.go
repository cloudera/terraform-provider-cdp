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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// GetEnvironmentUserSyncStateReader is a Reader for the GetEnvironmentUserSyncState structure.
type GetEnvironmentUserSyncStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnvironmentUserSyncStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEnvironmentUserSyncStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEnvironmentUserSyncStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEnvironmentUserSyncStateOK creates a GetEnvironmentUserSyncStateOK with default headers values
func NewGetEnvironmentUserSyncStateOK() *GetEnvironmentUserSyncStateOK {
	return &GetEnvironmentUserSyncStateOK{}
}

/*
GetEnvironmentUserSyncStateOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetEnvironmentUserSyncStateOK struct {
	Payload *models.GetEnvironmentUserSyncStateResponse
}

// IsSuccess returns true when this get environment user sync state o k response has a 2xx status code
func (o *GetEnvironmentUserSyncStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get environment user sync state o k response has a 3xx status code
func (o *GetEnvironmentUserSyncStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get environment user sync state o k response has a 4xx status code
func (o *GetEnvironmentUserSyncStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get environment user sync state o k response has a 5xx status code
func (o *GetEnvironmentUserSyncStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get environment user sync state o k response a status code equal to that given
func (o *GetEnvironmentUserSyncStateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get environment user sync state o k response
func (o *GetEnvironmentUserSyncStateOK) Code() int {
	return 200
}

func (o *GetEnvironmentUserSyncStateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/getEnvironmentUserSyncState][%d] getEnvironmentUserSyncStateOK %s", 200, payload)
}

func (o *GetEnvironmentUserSyncStateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/getEnvironmentUserSyncState][%d] getEnvironmentUserSyncStateOK %s", 200, payload)
}

func (o *GetEnvironmentUserSyncStateOK) GetPayload() *models.GetEnvironmentUserSyncStateResponse {
	return o.Payload
}

func (o *GetEnvironmentUserSyncStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetEnvironmentUserSyncStateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnvironmentUserSyncStateDefault creates a GetEnvironmentUserSyncStateDefault with default headers values
func NewGetEnvironmentUserSyncStateDefault(code int) *GetEnvironmentUserSyncStateDefault {
	return &GetEnvironmentUserSyncStateDefault{
		_statusCode: code,
	}
}

/*
GetEnvironmentUserSyncStateDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetEnvironmentUserSyncStateDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get environment user sync state default response has a 2xx status code
func (o *GetEnvironmentUserSyncStateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get environment user sync state default response has a 3xx status code
func (o *GetEnvironmentUserSyncStateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get environment user sync state default response has a 4xx status code
func (o *GetEnvironmentUserSyncStateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get environment user sync state default response has a 5xx status code
func (o *GetEnvironmentUserSyncStateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get environment user sync state default response a status code equal to that given
func (o *GetEnvironmentUserSyncStateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get environment user sync state default response
func (o *GetEnvironmentUserSyncStateDefault) Code() int {
	return o._statusCode
}

func (o *GetEnvironmentUserSyncStateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/getEnvironmentUserSyncState][%d] getEnvironmentUserSyncState default %s", o._statusCode, payload)
}

func (o *GetEnvironmentUserSyncStateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/getEnvironmentUserSyncState][%d] getEnvironmentUserSyncState default %s", o._statusCode, payload)
}

func (o *GetEnvironmentUserSyncStateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetEnvironmentUserSyncStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
