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

// GetK8sCertJKSReader is a Reader for the GetK8sCertJKS structure.
type GetK8sCertJKSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetK8sCertJKSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetK8sCertJKSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetK8sCertJKSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetK8sCertJKSOK creates a GetK8sCertJKSOK with default headers values
func NewGetK8sCertJKSOK() *GetK8sCertJKSOK {
	return &GetK8sCertJKSOK{}
}

/*
GetK8sCertJKSOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetK8sCertJKSOK struct {
	Payload *models.GetK8sCertJKSResponse
}

// IsSuccess returns true when this get k8s cert j k s o k response has a 2xx status code
func (o *GetK8sCertJKSOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get k8s cert j k s o k response has a 3xx status code
func (o *GetK8sCertJKSOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get k8s cert j k s o k response has a 4xx status code
func (o *GetK8sCertJKSOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get k8s cert j k s o k response has a 5xx status code
func (o *GetK8sCertJKSOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get k8s cert j k s o k response a status code equal to that given
func (o *GetK8sCertJKSOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get k8s cert j k s o k response
func (o *GetK8sCertJKSOK) Code() int {
	return 200
}

func (o *GetK8sCertJKSOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/getK8sCertJKS][%d] getK8sCertJKSOK %s", 200, payload)
}

func (o *GetK8sCertJKSOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/getK8sCertJKS][%d] getK8sCertJKSOK %s", 200, payload)
}

func (o *GetK8sCertJKSOK) GetPayload() *models.GetK8sCertJKSResponse {
	return o.Payload
}

func (o *GetK8sCertJKSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetK8sCertJKSResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetK8sCertJKSDefault creates a GetK8sCertJKSDefault with default headers values
func NewGetK8sCertJKSDefault(code int) *GetK8sCertJKSDefault {
	return &GetK8sCertJKSDefault{
		_statusCode: code,
	}
}

/*
GetK8sCertJKSDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetK8sCertJKSDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get k8s cert j k s default response has a 2xx status code
func (o *GetK8sCertJKSDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get k8s cert j k s default response has a 3xx status code
func (o *GetK8sCertJKSDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get k8s cert j k s default response has a 4xx status code
func (o *GetK8sCertJKSDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get k8s cert j k s default response has a 5xx status code
func (o *GetK8sCertJKSDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get k8s cert j k s default response a status code equal to that given
func (o *GetK8sCertJKSDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get k8s cert j k s default response
func (o *GetK8sCertJKSDefault) Code() int {
	return o._statusCode
}

func (o *GetK8sCertJKSDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/getK8sCertJKS][%d] getK8sCertJKS default %s", o._statusCode, payload)
}

func (o *GetK8sCertJKSDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/getK8sCertJKS][%d] getK8sCertJKS default %s", o._statusCode, payload)
}

func (o *GetK8sCertJKSDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetK8sCertJKSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
