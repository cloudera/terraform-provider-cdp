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

// GetRootCertificateReader is a Reader for the GetRootCertificate structure.
type GetRootCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRootCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRootCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRootCertificateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRootCertificateOK creates a GetRootCertificateOK with default headers values
func NewGetRootCertificateOK() *GetRootCertificateOK {
	return &GetRootCertificateOK{}
}

/*
GetRootCertificateOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetRootCertificateOK struct {
	Payload *models.GetRootCertificateResponse
}

// IsSuccess returns true when this get root certificate o k response has a 2xx status code
func (o *GetRootCertificateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get root certificate o k response has a 3xx status code
func (o *GetRootCertificateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get root certificate o k response has a 4xx status code
func (o *GetRootCertificateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get root certificate o k response has a 5xx status code
func (o *GetRootCertificateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get root certificate o k response a status code equal to that given
func (o *GetRootCertificateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get root certificate o k response
func (o *GetRootCertificateOK) Code() int {
	return 200
}

func (o *GetRootCertificateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/getRootCertificate][%d] getRootCertificateOK %s", 200, payload)
}

func (o *GetRootCertificateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/getRootCertificate][%d] getRootCertificateOK %s", 200, payload)
}

func (o *GetRootCertificateOK) GetPayload() *models.GetRootCertificateResponse {
	return o.Payload
}

func (o *GetRootCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetRootCertificateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRootCertificateDefault creates a GetRootCertificateDefault with default headers values
func NewGetRootCertificateDefault(code int) *GetRootCertificateDefault {
	return &GetRootCertificateDefault{
		_statusCode: code,
	}
}

/*
GetRootCertificateDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetRootCertificateDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get root certificate default response has a 2xx status code
func (o *GetRootCertificateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get root certificate default response has a 3xx status code
func (o *GetRootCertificateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get root certificate default response has a 4xx status code
func (o *GetRootCertificateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get root certificate default response has a 5xx status code
func (o *GetRootCertificateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get root certificate default response a status code equal to that given
func (o *GetRootCertificateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get root certificate default response
func (o *GetRootCertificateDefault) Code() int {
	return o._statusCode
}

func (o *GetRootCertificateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/getRootCertificate][%d] getRootCertificate default %s", o._statusCode, payload)
}

func (o *GetRootCertificateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/environments2/getRootCertificate][%d] getRootCertificate default %s", o._statusCode, payload)
}

func (o *GetRootCertificateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRootCertificateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
