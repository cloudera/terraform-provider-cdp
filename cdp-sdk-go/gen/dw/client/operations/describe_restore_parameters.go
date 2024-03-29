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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// NewDescribeRestoreParams creates a new DescribeRestoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDescribeRestoreParams() *DescribeRestoreParams {
	return &DescribeRestoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeRestoreParamsWithTimeout creates a new DescribeRestoreParams object
// with the ability to set a timeout on a request.
func NewDescribeRestoreParamsWithTimeout(timeout time.Duration) *DescribeRestoreParams {
	return &DescribeRestoreParams{
		timeout: timeout,
	}
}

// NewDescribeRestoreParamsWithContext creates a new DescribeRestoreParams object
// with the ability to set a context for a request.
func NewDescribeRestoreParamsWithContext(ctx context.Context) *DescribeRestoreParams {
	return &DescribeRestoreParams{
		Context: ctx,
	}
}

// NewDescribeRestoreParamsWithHTTPClient creates a new DescribeRestoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewDescribeRestoreParamsWithHTTPClient(client *http.Client) *DescribeRestoreParams {
	return &DescribeRestoreParams{
		HTTPClient: client,
	}
}

/*
DescribeRestoreParams contains all the parameters to send to the API endpoint

	for the describe restore operation.

	Typically these are written to a http.Request.
*/
type DescribeRestoreParams struct {

	// Input.
	Input *models.DescribeRestoreRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the describe restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeRestoreParams) WithDefaults() *DescribeRestoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the describe restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DescribeRestoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the describe restore params
func (o *DescribeRestoreParams) WithTimeout(timeout time.Duration) *DescribeRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe restore params
func (o *DescribeRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe restore params
func (o *DescribeRestoreParams) WithContext(ctx context.Context) *DescribeRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe restore params
func (o *DescribeRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe restore params
func (o *DescribeRestoreParams) WithHTTPClient(client *http.Client) *DescribeRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe restore params
func (o *DescribeRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the describe restore params
func (o *DescribeRestoreParams) WithInput(input *models.DescribeRestoreRequest) *DescribeRestoreParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the describe restore params
func (o *DescribeRestoreParams) SetInput(input *models.DescribeRestoreRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
