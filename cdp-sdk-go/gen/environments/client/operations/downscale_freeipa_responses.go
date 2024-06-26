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

// DownscaleFreeipaReader is a Reader for the DownscaleFreeipa structure.
type DownscaleFreeipaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownscaleFreeipaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownscaleFreeipaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDownscaleFreeipaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDownscaleFreeipaOK creates a DownscaleFreeipaOK with default headers values
func NewDownscaleFreeipaOK() *DownscaleFreeipaOK {
	return &DownscaleFreeipaOK{}
}

/*
DownscaleFreeipaOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DownscaleFreeipaOK struct {
	Payload *models.DownscaleFreeipaResponse
}

// IsSuccess returns true when this downscale freeipa o k response has a 2xx status code
func (o *DownscaleFreeipaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this downscale freeipa o k response has a 3xx status code
func (o *DownscaleFreeipaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this downscale freeipa o k response has a 4xx status code
func (o *DownscaleFreeipaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this downscale freeipa o k response has a 5xx status code
func (o *DownscaleFreeipaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this downscale freeipa o k response a status code equal to that given
func (o *DownscaleFreeipaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the downscale freeipa o k response
func (o *DownscaleFreeipaOK) Code() int {
	return 200
}

func (o *DownscaleFreeipaOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/downscaleFreeipa][%d] downscaleFreeipaOK %s", 200, payload)
}

func (o *DownscaleFreeipaOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/downscaleFreeipa][%d] downscaleFreeipaOK %s", 200, payload)
}

func (o *DownscaleFreeipaOK) GetPayload() *models.DownscaleFreeipaResponse {
	return o.Payload
}

func (o *DownscaleFreeipaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DownscaleFreeipaResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownscaleFreeipaDefault creates a DownscaleFreeipaDefault with default headers values
func NewDownscaleFreeipaDefault(code int) *DownscaleFreeipaDefault {
	return &DownscaleFreeipaDefault{
		_statusCode: code,
	}
}

/*
DownscaleFreeipaDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DownscaleFreeipaDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this downscale freeipa default response has a 2xx status code
func (o *DownscaleFreeipaDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this downscale freeipa default response has a 3xx status code
func (o *DownscaleFreeipaDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this downscale freeipa default response has a 4xx status code
func (o *DownscaleFreeipaDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this downscale freeipa default response has a 5xx status code
func (o *DownscaleFreeipaDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this downscale freeipa default response a status code equal to that given
func (o *DownscaleFreeipaDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the downscale freeipa default response
func (o *DownscaleFreeipaDefault) Code() int {
	return o._statusCode
}

func (o *DownscaleFreeipaDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/downscaleFreeipa][%d] downscaleFreeipa default %s", o._statusCode, payload)
}

func (o *DownscaleFreeipaDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/downscaleFreeipa][%d] downscaleFreeipa default %s", o._statusCode, payload)
}

func (o *DownscaleFreeipaDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DownscaleFreeipaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
