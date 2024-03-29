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

// NewGetAccountTelemetryDefaultParams creates a new GetAccountTelemetryDefaultParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountTelemetryDefaultParams() *GetAccountTelemetryDefaultParams {
	return &GetAccountTelemetryDefaultParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountTelemetryDefaultParamsWithTimeout creates a new GetAccountTelemetryDefaultParams object
// with the ability to set a timeout on a request.
func NewGetAccountTelemetryDefaultParamsWithTimeout(timeout time.Duration) *GetAccountTelemetryDefaultParams {
	return &GetAccountTelemetryDefaultParams{
		timeout: timeout,
	}
}

// NewGetAccountTelemetryDefaultParamsWithContext creates a new GetAccountTelemetryDefaultParams object
// with the ability to set a context for a request.
func NewGetAccountTelemetryDefaultParamsWithContext(ctx context.Context) *GetAccountTelemetryDefaultParams {
	return &GetAccountTelemetryDefaultParams{
		Context: ctx,
	}
}

// NewGetAccountTelemetryDefaultParamsWithHTTPClient creates a new GetAccountTelemetryDefaultParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountTelemetryDefaultParamsWithHTTPClient(client *http.Client) *GetAccountTelemetryDefaultParams {
	return &GetAccountTelemetryDefaultParams{
		HTTPClient: client,
	}
}

/*
GetAccountTelemetryDefaultParams contains all the parameters to send to the API endpoint

	for the get account telemetry default operation.

	Typically these are written to a http.Request.
*/
type GetAccountTelemetryDefaultParams struct {

	// Input.
	Input models.GetAccountTelemetryDefaultRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get account telemetry default params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountTelemetryDefaultParams) WithDefaults() *GetAccountTelemetryDefaultParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get account telemetry default params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountTelemetryDefaultParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get account telemetry default params
func (o *GetAccountTelemetryDefaultParams) WithTimeout(timeout time.Duration) *GetAccountTelemetryDefaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account telemetry default params
func (o *GetAccountTelemetryDefaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account telemetry default params
func (o *GetAccountTelemetryDefaultParams) WithContext(ctx context.Context) *GetAccountTelemetryDefaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account telemetry default params
func (o *GetAccountTelemetryDefaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account telemetry default params
func (o *GetAccountTelemetryDefaultParams) WithHTTPClient(client *http.Client) *GetAccountTelemetryDefaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account telemetry default params
func (o *GetAccountTelemetryDefaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get account telemetry default params
func (o *GetAccountTelemetryDefaultParams) WithInput(input models.GetAccountTelemetryDefaultRequest) *GetAccountTelemetryDefaultParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get account telemetry default params
func (o *GetAccountTelemetryDefaultParams) SetInput(input models.GetAccountTelemetryDefaultRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountTelemetryDefaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
