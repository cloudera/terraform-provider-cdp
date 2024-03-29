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

// NewDisengageAutoAdminParams creates a new DisengageAutoAdminParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDisengageAutoAdminParams() *DisengageAutoAdminParams {
	return &DisengageAutoAdminParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDisengageAutoAdminParamsWithTimeout creates a new DisengageAutoAdminParams object
// with the ability to set a timeout on a request.
func NewDisengageAutoAdminParamsWithTimeout(timeout time.Duration) *DisengageAutoAdminParams {
	return &DisengageAutoAdminParams{
		timeout: timeout,
	}
}

// NewDisengageAutoAdminParamsWithContext creates a new DisengageAutoAdminParams object
// with the ability to set a context for a request.
func NewDisengageAutoAdminParamsWithContext(ctx context.Context) *DisengageAutoAdminParams {
	return &DisengageAutoAdminParams{
		Context: ctx,
	}
}

// NewDisengageAutoAdminParamsWithHTTPClient creates a new DisengageAutoAdminParams object
// with the ability to set a custom HTTPClient for a request.
func NewDisengageAutoAdminParamsWithHTTPClient(client *http.Client) *DisengageAutoAdminParams {
	return &DisengageAutoAdminParams{
		HTTPClient: client,
	}
}

/*
DisengageAutoAdminParams contains all the parameters to send to the API endpoint

	for the disengage auto admin operation.

	Typically these are written to a http.Request.
*/
type DisengageAutoAdminParams struct {

	// Input.
	Input *models.DisengageAutoAdminRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disengage auto admin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisengageAutoAdminParams) WithDefaults() *DisengageAutoAdminParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disengage auto admin params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisengageAutoAdminParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disengage auto admin params
func (o *DisengageAutoAdminParams) WithTimeout(timeout time.Duration) *DisengageAutoAdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disengage auto admin params
func (o *DisengageAutoAdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disengage auto admin params
func (o *DisengageAutoAdminParams) WithContext(ctx context.Context) *DisengageAutoAdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disengage auto admin params
func (o *DisengageAutoAdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disengage auto admin params
func (o *DisengageAutoAdminParams) WithHTTPClient(client *http.Client) *DisengageAutoAdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disengage auto admin params
func (o *DisengageAutoAdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the disengage auto admin params
func (o *DisengageAutoAdminParams) WithInput(input *models.DisengageAutoAdminRequest) *DisengageAutoAdminParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the disengage auto admin params
func (o *DisengageAutoAdminParams) SetInput(input *models.DisengageAutoAdminRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DisengageAutoAdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
