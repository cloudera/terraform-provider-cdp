// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// NewCreateMachineUserParams creates a new CreateMachineUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMachineUserParams() *CreateMachineUserParams {
	return &CreateMachineUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMachineUserParamsWithTimeout creates a new CreateMachineUserParams object
// with the ability to set a timeout on a request.
func NewCreateMachineUserParamsWithTimeout(timeout time.Duration) *CreateMachineUserParams {
	return &CreateMachineUserParams{
		timeout: timeout,
	}
}

// NewCreateMachineUserParamsWithContext creates a new CreateMachineUserParams object
// with the ability to set a context for a request.
func NewCreateMachineUserParamsWithContext(ctx context.Context) *CreateMachineUserParams {
	return &CreateMachineUserParams{
		Context: ctx,
	}
}

// NewCreateMachineUserParamsWithHTTPClient creates a new CreateMachineUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMachineUserParamsWithHTTPClient(client *http.Client) *CreateMachineUserParams {
	return &CreateMachineUserParams{
		HTTPClient: client,
	}
}

/*
CreateMachineUserParams contains all the parameters to send to the API endpoint

	for the create machine user operation.

	Typically these are written to a http.Request.
*/
type CreateMachineUserParams struct {

	// Input.
	Input *models.CreateMachineUserRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create machine user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMachineUserParams) WithDefaults() *CreateMachineUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create machine user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMachineUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create machine user params
func (o *CreateMachineUserParams) WithTimeout(timeout time.Duration) *CreateMachineUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create machine user params
func (o *CreateMachineUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create machine user params
func (o *CreateMachineUserParams) WithContext(ctx context.Context) *CreateMachineUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create machine user params
func (o *CreateMachineUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create machine user params
func (o *CreateMachineUserParams) WithHTTPClient(client *http.Client) *CreateMachineUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create machine user params
func (o *CreateMachineUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the create machine user params
func (o *CreateMachineUserParams) WithInput(input *models.CreateMachineUserRequest) *CreateMachineUserParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the create machine user params
func (o *CreateMachineUserParams) SetInput(input *models.CreateMachineUserRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMachineUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
