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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

// NewCreateDatabaseParams creates a new CreateDatabaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateDatabaseParams() *CreateDatabaseParams {
	return &CreateDatabaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateDatabaseParamsWithTimeout creates a new CreateDatabaseParams object
// with the ability to set a timeout on a request.
func NewCreateDatabaseParamsWithTimeout(timeout time.Duration) *CreateDatabaseParams {
	return &CreateDatabaseParams{
		timeout: timeout,
	}
}

// NewCreateDatabaseParamsWithContext creates a new CreateDatabaseParams object
// with the ability to set a context for a request.
func NewCreateDatabaseParamsWithContext(ctx context.Context) *CreateDatabaseParams {
	return &CreateDatabaseParams{
		Context: ctx,
	}
}

// NewCreateDatabaseParamsWithHTTPClient creates a new CreateDatabaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateDatabaseParamsWithHTTPClient(client *http.Client) *CreateDatabaseParams {
	return &CreateDatabaseParams{
		HTTPClient: client,
	}
}

/*
CreateDatabaseParams contains all the parameters to send to the API endpoint

	for the create database operation.

	Typically these are written to a http.Request.
*/
type CreateDatabaseParams struct {

	// Input.
	Input *models.CreateDatabaseRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create database params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDatabaseParams) WithDefaults() *CreateDatabaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create database params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateDatabaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create database params
func (o *CreateDatabaseParams) WithTimeout(timeout time.Duration) *CreateDatabaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create database params
func (o *CreateDatabaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create database params
func (o *CreateDatabaseParams) WithContext(ctx context.Context) *CreateDatabaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create database params
func (o *CreateDatabaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create database params
func (o *CreateDatabaseParams) WithHTTPClient(client *http.Client) *CreateDatabaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create database params
func (o *CreateDatabaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the create database params
func (o *CreateDatabaseParams) WithInput(input *models.CreateDatabaseRequest) *CreateDatabaseParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the create database params
func (o *CreateDatabaseParams) SetInput(input *models.CreateDatabaseRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDatabaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
