// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// ListSSHPublicKeysReader is a Reader for the ListSSHPublicKeys structure.
type ListSSHPublicKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSSHPublicKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSSHPublicKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSSHPublicKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSSHPublicKeysOK creates a ListSSHPublicKeysOK with default headers values
func NewListSSHPublicKeysOK() *ListSSHPublicKeysOK {
	return &ListSSHPublicKeysOK{}
}

/*
ListSSHPublicKeysOK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type ListSSHPublicKeysOK struct {
	Payload *models.ListSSHPublicKeysResponse
}

// IsSuccess returns true when this list Ssh public keys o k response has a 2xx status code
func (o *ListSSHPublicKeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list Ssh public keys o k response has a 3xx status code
func (o *ListSSHPublicKeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list Ssh public keys o k response has a 4xx status code
func (o *ListSSHPublicKeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list Ssh public keys o k response has a 5xx status code
func (o *ListSSHPublicKeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list Ssh public keys o k response a status code equal to that given
func (o *ListSSHPublicKeysOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list Ssh public keys o k response
func (o *ListSSHPublicKeysOK) Code() int {
	return 200
}

func (o *ListSSHPublicKeysOK) Error() string {
	return fmt.Sprintf("[POST /iam/listSshPublicKeys][%d] listSshPublicKeysOK  %+v", 200, o.Payload)
}

func (o *ListSSHPublicKeysOK) String() string {
	return fmt.Sprintf("[POST /iam/listSshPublicKeys][%d] listSshPublicKeysOK  %+v", 200, o.Payload)
}

func (o *ListSSHPublicKeysOK) GetPayload() *models.ListSSHPublicKeysResponse {
	return o.Payload
}

func (o *ListSSHPublicKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListSSHPublicKeysResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSSHPublicKeysDefault creates a ListSSHPublicKeysDefault with default headers values
func NewListSSHPublicKeysDefault(code int) *ListSSHPublicKeysDefault {
	return &ListSSHPublicKeysDefault{
		_statusCode: code,
	}
}

/*
ListSSHPublicKeysDefault describes a response with status code -1, with default header values.

The default response on an error.
*/
type ListSSHPublicKeysDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list Ssh public keys default response has a 2xx status code
func (o *ListSSHPublicKeysDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list Ssh public keys default response has a 3xx status code
func (o *ListSSHPublicKeysDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list Ssh public keys default response has a 4xx status code
func (o *ListSSHPublicKeysDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list Ssh public keys default response has a 5xx status code
func (o *ListSSHPublicKeysDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list Ssh public keys default response a status code equal to that given
func (o *ListSSHPublicKeysDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list Ssh public keys default response
func (o *ListSSHPublicKeysDefault) Code() int {
	return o._statusCode
}

func (o *ListSSHPublicKeysDefault) Error() string {
	return fmt.Sprintf("[POST /iam/listSshPublicKeys][%d] listSshPublicKeys default  %+v", o._statusCode, o.Payload)
}

func (o *ListSSHPublicKeysDefault) String() string {
	return fmt.Sprintf("[POST /iam/listSshPublicKeys][%d] listSshPublicKeys default  %+v", o._statusCode, o.Payload)
}

func (o *ListSSHPublicKeysDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListSSHPublicKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}