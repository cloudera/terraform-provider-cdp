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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// EnableClouderaSSOLoginReader is a Reader for the EnableClouderaSSOLogin structure.
type EnableClouderaSSOLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableClouderaSSOLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableClouderaSSOLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEnableClouderaSSOLoginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEnableClouderaSSOLoginOK creates a EnableClouderaSSOLoginOK with default headers values
func NewEnableClouderaSSOLoginOK() *EnableClouderaSSOLoginOK {
	return &EnableClouderaSSOLoginOK{}
}

/*
EnableClouderaSSOLoginOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type EnableClouderaSSOLoginOK struct {
	Payload models.EnableClouderaSSOLoginResponse
}

// IsSuccess returns true when this enable cloudera s s o login o k response has a 2xx status code
func (o *EnableClouderaSSOLoginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this enable cloudera s s o login o k response has a 3xx status code
func (o *EnableClouderaSSOLoginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this enable cloudera s s o login o k response has a 4xx status code
func (o *EnableClouderaSSOLoginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this enable cloudera s s o login o k response has a 5xx status code
func (o *EnableClouderaSSOLoginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this enable cloudera s s o login o k response a status code equal to that given
func (o *EnableClouderaSSOLoginOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the enable cloudera s s o login o k response
func (o *EnableClouderaSSOLoginOK) Code() int {
	return 200
}

func (o *EnableClouderaSSOLoginOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/enableClouderaSSOLogin][%d] enableClouderaSSOLoginOK %s", 200, payload)
}

func (o *EnableClouderaSSOLoginOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/enableClouderaSSOLogin][%d] enableClouderaSSOLoginOK %s", 200, payload)
}

func (o *EnableClouderaSSOLoginOK) GetPayload() models.EnableClouderaSSOLoginResponse {
	return o.Payload
}

func (o *EnableClouderaSSOLoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableClouderaSSOLoginDefault creates a EnableClouderaSSOLoginDefault with default headers values
func NewEnableClouderaSSOLoginDefault(code int) *EnableClouderaSSOLoginDefault {
	return &EnableClouderaSSOLoginDefault{
		_statusCode: code,
	}
}

/*
EnableClouderaSSOLoginDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type EnableClouderaSSOLoginDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this enable cloudera s s o login default response has a 2xx status code
func (o *EnableClouderaSSOLoginDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this enable cloudera s s o login default response has a 3xx status code
func (o *EnableClouderaSSOLoginDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this enable cloudera s s o login default response has a 4xx status code
func (o *EnableClouderaSSOLoginDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this enable cloudera s s o login default response has a 5xx status code
func (o *EnableClouderaSSOLoginDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this enable cloudera s s o login default response a status code equal to that given
func (o *EnableClouderaSSOLoginDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the enable cloudera s s o login default response
func (o *EnableClouderaSSOLoginDefault) Code() int {
	return o._statusCode
}

func (o *EnableClouderaSSOLoginDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/enableClouderaSSOLogin][%d] enableClouderaSSOLogin default %s", o._statusCode, payload)
}

func (o *EnableClouderaSSOLoginDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /iam/enableClouderaSSOLogin][%d] enableClouderaSSOLogin default %s", o._statusCode, payload)
}

func (o *EnableClouderaSSOLoginDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnableClouderaSSOLoginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
