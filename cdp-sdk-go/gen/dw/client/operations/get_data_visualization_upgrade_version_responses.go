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

// GetDataVisualizationUpgradeVersionReader is a Reader for the GetDataVisualizationUpgradeVersion structure.
type GetDataVisualizationUpgradeVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataVisualizationUpgradeVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataVisualizationUpgradeVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDataVisualizationUpgradeVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDataVisualizationUpgradeVersionOK creates a GetDataVisualizationUpgradeVersionOK with default headers values
func NewGetDataVisualizationUpgradeVersionOK() *GetDataVisualizationUpgradeVersionOK {
	return &GetDataVisualizationUpgradeVersionOK{}
}

/*
GetDataVisualizationUpgradeVersionOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type GetDataVisualizationUpgradeVersionOK struct {
	Payload *models.GetDataVisualizationUpgradeVersionResponse
}

// IsSuccess returns true when this get data visualization upgrade version o k response has a 2xx status code
func (o *GetDataVisualizationUpgradeVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get data visualization upgrade version o k response has a 3xx status code
func (o *GetDataVisualizationUpgradeVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get data visualization upgrade version o k response has a 4xx status code
func (o *GetDataVisualizationUpgradeVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get data visualization upgrade version o k response has a 5xx status code
func (o *GetDataVisualizationUpgradeVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get data visualization upgrade version o k response a status code equal to that given
func (o *GetDataVisualizationUpgradeVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get data visualization upgrade version o k response
func (o *GetDataVisualizationUpgradeVersionOK) Code() int {
	return 200
}

func (o *GetDataVisualizationUpgradeVersionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/getDataVisualizationUpgradeVersion][%d] getDataVisualizationUpgradeVersionOK %s", 200, payload)
}

func (o *GetDataVisualizationUpgradeVersionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/getDataVisualizationUpgradeVersion][%d] getDataVisualizationUpgradeVersionOK %s", 200, payload)
}

func (o *GetDataVisualizationUpgradeVersionOK) GetPayload() *models.GetDataVisualizationUpgradeVersionResponse {
	return o.Payload
}

func (o *GetDataVisualizationUpgradeVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetDataVisualizationUpgradeVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataVisualizationUpgradeVersionDefault creates a GetDataVisualizationUpgradeVersionDefault with default headers values
func NewGetDataVisualizationUpgradeVersionDefault(code int) *GetDataVisualizationUpgradeVersionDefault {
	return &GetDataVisualizationUpgradeVersionDefault{
		_statusCode: code,
	}
}

/*
GetDataVisualizationUpgradeVersionDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type GetDataVisualizationUpgradeVersionDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get data visualization upgrade version default response has a 2xx status code
func (o *GetDataVisualizationUpgradeVersionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get data visualization upgrade version default response has a 3xx status code
func (o *GetDataVisualizationUpgradeVersionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get data visualization upgrade version default response has a 4xx status code
func (o *GetDataVisualizationUpgradeVersionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get data visualization upgrade version default response has a 5xx status code
func (o *GetDataVisualizationUpgradeVersionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get data visualization upgrade version default response a status code equal to that given
func (o *GetDataVisualizationUpgradeVersionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get data visualization upgrade version default response
func (o *GetDataVisualizationUpgradeVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetDataVisualizationUpgradeVersionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/getDataVisualizationUpgradeVersion][%d] getDataVisualizationUpgradeVersion default %s", o._statusCode, payload)
}

func (o *GetDataVisualizationUpgradeVersionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/getDataVisualizationUpgradeVersion][%d] getDataVisualizationUpgradeVersion default %s", o._statusCode, payload)
}

func (o *GetDataVisualizationUpgradeVersionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDataVisualizationUpgradeVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
