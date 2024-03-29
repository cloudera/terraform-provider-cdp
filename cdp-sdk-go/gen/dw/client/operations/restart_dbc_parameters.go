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

// NewRestartDbcParams creates a new RestartDbcParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestartDbcParams() *RestartDbcParams {
	return &RestartDbcParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestartDbcParamsWithTimeout creates a new RestartDbcParams object
// with the ability to set a timeout on a request.
func NewRestartDbcParamsWithTimeout(timeout time.Duration) *RestartDbcParams {
	return &RestartDbcParams{
		timeout: timeout,
	}
}

// NewRestartDbcParamsWithContext creates a new RestartDbcParams object
// with the ability to set a context for a request.
func NewRestartDbcParamsWithContext(ctx context.Context) *RestartDbcParams {
	return &RestartDbcParams{
		Context: ctx,
	}
}

// NewRestartDbcParamsWithHTTPClient creates a new RestartDbcParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestartDbcParamsWithHTTPClient(client *http.Client) *RestartDbcParams {
	return &RestartDbcParams{
		HTTPClient: client,
	}
}

/*
RestartDbcParams contains all the parameters to send to the API endpoint

	for the restart dbc operation.

	Typically these are written to a http.Request.
*/
type RestartDbcParams struct {

	// Input.
	Input *models.RestartDbcRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restart dbc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartDbcParams) WithDefaults() *RestartDbcParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restart dbc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestartDbcParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restart dbc params
func (o *RestartDbcParams) WithTimeout(timeout time.Duration) *RestartDbcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restart dbc params
func (o *RestartDbcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restart dbc params
func (o *RestartDbcParams) WithContext(ctx context.Context) *RestartDbcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restart dbc params
func (o *RestartDbcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restart dbc params
func (o *RestartDbcParams) WithHTTPClient(client *http.Client) *RestartDbcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restart dbc params
func (o *RestartDbcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the restart dbc params
func (o *RestartDbcParams) WithInput(input *models.RestartDbcRequest) *RestartDbcParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the restart dbc params
func (o *RestartDbcParams) SetInput(input *models.RestartDbcRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *RestartDbcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
