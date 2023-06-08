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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// NewStopDatalakeParams creates a new StopDatalakeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStopDatalakeParams() *StopDatalakeParams {
	return &StopDatalakeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStopDatalakeParamsWithTimeout creates a new StopDatalakeParams object
// with the ability to set a timeout on a request.
func NewStopDatalakeParamsWithTimeout(timeout time.Duration) *StopDatalakeParams {
	return &StopDatalakeParams{
		timeout: timeout,
	}
}

// NewStopDatalakeParamsWithContext creates a new StopDatalakeParams object
// with the ability to set a context for a request.
func NewStopDatalakeParamsWithContext(ctx context.Context) *StopDatalakeParams {
	return &StopDatalakeParams{
		Context: ctx,
	}
}

// NewStopDatalakeParamsWithHTTPClient creates a new StopDatalakeParams object
// with the ability to set a custom HTTPClient for a request.
func NewStopDatalakeParamsWithHTTPClient(client *http.Client) *StopDatalakeParams {
	return &StopDatalakeParams{
		HTTPClient: client,
	}
}

/*
StopDatalakeParams contains all the parameters to send to the API endpoint

	for the stop datalake operation.

	Typically these are written to a http.Request.
*/
type StopDatalakeParams struct {

	// Input.
	Input *models.StopDatalakeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the stop datalake params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopDatalakeParams) WithDefaults() *StopDatalakeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the stop datalake params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StopDatalakeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the stop datalake params
func (o *StopDatalakeParams) WithTimeout(timeout time.Duration) *StopDatalakeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop datalake params
func (o *StopDatalakeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop datalake params
func (o *StopDatalakeParams) WithContext(ctx context.Context) *StopDatalakeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop datalake params
func (o *StopDatalakeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop datalake params
func (o *StopDatalakeParams) WithHTTPClient(client *http.Client) *StopDatalakeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop datalake params
func (o *StopDatalakeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the stop datalake params
func (o *StopDatalakeParams) WithInput(input *models.StopDatalakeRequest) *StopDatalakeParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the stop datalake params
func (o *StopDatalakeParams) SetInput(input *models.StopDatalakeRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *StopDatalakeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
