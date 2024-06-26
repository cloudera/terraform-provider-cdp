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

// DescribeServerSettingReader is a Reader for the DescribeServerSetting structure.
type DescribeServerSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeServerSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeServerSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeServerSettingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeServerSettingOK creates a DescribeServerSettingOK with default headers values
func NewDescribeServerSettingOK() *DescribeServerSettingOK {
	return &DescribeServerSettingOK{}
}

/*
DescribeServerSettingOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type DescribeServerSettingOK struct {
	Payload *models.DescribeServerSettingResponse
}

// IsSuccess returns true when this describe server setting o k response has a 2xx status code
func (o *DescribeServerSettingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this describe server setting o k response has a 3xx status code
func (o *DescribeServerSettingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this describe server setting o k response has a 4xx status code
func (o *DescribeServerSettingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this describe server setting o k response has a 5xx status code
func (o *DescribeServerSettingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this describe server setting o k response a status code equal to that given
func (o *DescribeServerSettingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the describe server setting o k response
func (o *DescribeServerSettingOK) Code() int {
	return 200
}

func (o *DescribeServerSettingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeServerSetting][%d] describeServerSettingOK %s", 200, payload)
}

func (o *DescribeServerSettingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeServerSetting][%d] describeServerSettingOK %s", 200, payload)
}

func (o *DescribeServerSettingOK) GetPayload() *models.DescribeServerSettingResponse {
	return o.Payload
}

func (o *DescribeServerSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeServerSettingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeServerSettingDefault creates a DescribeServerSettingDefault with default headers values
func NewDescribeServerSettingDefault(code int) *DescribeServerSettingDefault {
	return &DescribeServerSettingDefault{
		_statusCode: code,
	}
}

/*
DescribeServerSettingDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type DescribeServerSettingDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this describe server setting default response has a 2xx status code
func (o *DescribeServerSettingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this describe server setting default response has a 3xx status code
func (o *DescribeServerSettingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this describe server setting default response has a 4xx status code
func (o *DescribeServerSettingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this describe server setting default response has a 5xx status code
func (o *DescribeServerSettingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this describe server setting default response a status code equal to that given
func (o *DescribeServerSettingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the describe server setting default response
func (o *DescribeServerSettingDefault) Code() int {
	return o._statusCode
}

func (o *DescribeServerSettingDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeServerSetting][%d] describeServerSetting default %s", o._statusCode, payload)
}

func (o *DescribeServerSettingDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v1/dw/describeServerSetting][%d] describeServerSetting default %s", o._statusCode, payload)
}

func (o *DescribeServerSettingDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeServerSettingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
