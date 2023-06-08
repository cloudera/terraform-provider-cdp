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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// NewSetAccountTelemetryParams creates a new SetAccountTelemetryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetAccountTelemetryParams() *SetAccountTelemetryParams {
	return &SetAccountTelemetryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetAccountTelemetryParamsWithTimeout creates a new SetAccountTelemetryParams object
// with the ability to set a timeout on a request.
func NewSetAccountTelemetryParamsWithTimeout(timeout time.Duration) *SetAccountTelemetryParams {
	return &SetAccountTelemetryParams{
		timeout: timeout,
	}
}

// NewSetAccountTelemetryParamsWithContext creates a new SetAccountTelemetryParams object
// with the ability to set a context for a request.
func NewSetAccountTelemetryParamsWithContext(ctx context.Context) *SetAccountTelemetryParams {
	return &SetAccountTelemetryParams{
		Context: ctx,
	}
}

// NewSetAccountTelemetryParamsWithHTTPClient creates a new SetAccountTelemetryParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetAccountTelemetryParamsWithHTTPClient(client *http.Client) *SetAccountTelemetryParams {
	return &SetAccountTelemetryParams{
		HTTPClient: client,
	}
}

/*
SetAccountTelemetryParams contains all the parameters to send to the API endpoint

	for the set account telemetry operation.

	Typically these are written to a http.Request.
*/
type SetAccountTelemetryParams struct {

	// Input.
	Input *models.SetAccountTelemetryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set account telemetry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAccountTelemetryParams) WithDefaults() *SetAccountTelemetryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set account telemetry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAccountTelemetryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set account telemetry params
func (o *SetAccountTelemetryParams) WithTimeout(timeout time.Duration) *SetAccountTelemetryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set account telemetry params
func (o *SetAccountTelemetryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set account telemetry params
func (o *SetAccountTelemetryParams) WithContext(ctx context.Context) *SetAccountTelemetryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set account telemetry params
func (o *SetAccountTelemetryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set account telemetry params
func (o *SetAccountTelemetryParams) WithHTTPClient(client *http.Client) *SetAccountTelemetryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set account telemetry params
func (o *SetAccountTelemetryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the set account telemetry params
func (o *SetAccountTelemetryParams) WithInput(input *models.SetAccountTelemetryRequest) *SetAccountTelemetryParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the set account telemetry params
func (o *SetAccountTelemetryParams) SetInput(input *models.SetAccountTelemetryRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *SetAccountTelemetryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
