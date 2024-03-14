// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// GetK8sCertPEMReader is a Reader for the GetK8sCertPEM structure.
type GetK8sCertPEMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetK8sCertPEMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetK8sCertPEMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetK8sCertPEMDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetK8sCertPEMOK creates a GetK8sCertPEMOK with default headers values
func NewGetK8sCertPEMOK() *GetK8sCertPEMOK {
	return &GetK8sCertPEMOK{}
}

/*
GetK8sCertPEMOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetK8sCertPEMOK struct {
	Payload *models.GetK8sCertPEMResponse
}

// IsSuccess returns true when this get k8s cert p e m o k response has a 2xx status code
func (o *GetK8sCertPEMOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get k8s cert p e m o k response has a 3xx status code
func (o *GetK8sCertPEMOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get k8s cert p e m o k response has a 4xx status code
func (o *GetK8sCertPEMOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get k8s cert p e m o k response has a 5xx status code
func (o *GetK8sCertPEMOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get k8s cert p e m o k response a status code equal to that given
func (o *GetK8sCertPEMOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get k8s cert p e m o k response
func (o *GetK8sCertPEMOK) Code() int {
	return 200
}

func (o *GetK8sCertPEMOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/getK8sCertPEM][%d] getK8sCertPEMOK  %+v", 200, o.Payload)
}

func (o *GetK8sCertPEMOK) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/getK8sCertPEM][%d] getK8sCertPEMOK  %+v", 200, o.Payload)
}

func (o *GetK8sCertPEMOK) GetPayload() *models.GetK8sCertPEMResponse {
	return o.Payload
}

func (o *GetK8sCertPEMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetK8sCertPEMResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetK8sCertPEMDefault creates a GetK8sCertPEMDefault with default headers values
func NewGetK8sCertPEMDefault(code int) *GetK8sCertPEMDefault {
	return &GetK8sCertPEMDefault{
		_statusCode: code,
	}
}

/*
GetK8sCertPEMDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetK8sCertPEMDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get k8s cert p e m default response has a 2xx status code
func (o *GetK8sCertPEMDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get k8s cert p e m default response has a 3xx status code
func (o *GetK8sCertPEMDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get k8s cert p e m default response has a 4xx status code
func (o *GetK8sCertPEMDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get k8s cert p e m default response has a 5xx status code
func (o *GetK8sCertPEMDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get k8s cert p e m default response a status code equal to that given
func (o *GetK8sCertPEMDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get k8s cert p e m default response
func (o *GetK8sCertPEMDefault) Code() int {
	return o._statusCode
}

func (o *GetK8sCertPEMDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/dw/getK8sCertPEM][%d] getK8sCertPEM default  %+v", o._statusCode, o.Payload)
}

func (o *GetK8sCertPEMDefault) String() string {
	return fmt.Sprintf("[POST /api/v1/dw/getK8sCertPEM][%d] getK8sCertPEM default  %+v", o._statusCode, o.Payload)
}

func (o *GetK8sCertPEMDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetK8sCertPEMDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
