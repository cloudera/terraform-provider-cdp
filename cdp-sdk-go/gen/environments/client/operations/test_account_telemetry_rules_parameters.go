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

// NewTestAccountTelemetryRulesParams creates a new TestAccountTelemetryRulesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTestAccountTelemetryRulesParams() *TestAccountTelemetryRulesParams {
	return &TestAccountTelemetryRulesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTestAccountTelemetryRulesParamsWithTimeout creates a new TestAccountTelemetryRulesParams object
// with the ability to set a timeout on a request.
func NewTestAccountTelemetryRulesParamsWithTimeout(timeout time.Duration) *TestAccountTelemetryRulesParams {
	return &TestAccountTelemetryRulesParams{
		timeout: timeout,
	}
}

// NewTestAccountTelemetryRulesParamsWithContext creates a new TestAccountTelemetryRulesParams object
// with the ability to set a context for a request.
func NewTestAccountTelemetryRulesParamsWithContext(ctx context.Context) *TestAccountTelemetryRulesParams {
	return &TestAccountTelemetryRulesParams{
		Context: ctx,
	}
}

// NewTestAccountTelemetryRulesParamsWithHTTPClient creates a new TestAccountTelemetryRulesParams object
// with the ability to set a custom HTTPClient for a request.
func NewTestAccountTelemetryRulesParamsWithHTTPClient(client *http.Client) *TestAccountTelemetryRulesParams {
	return &TestAccountTelemetryRulesParams{
		HTTPClient: client,
	}
}

/*
TestAccountTelemetryRulesParams contains all the parameters to send to the API endpoint

	for the test account telemetry rules operation.

	Typically these are written to a http.Request.
*/
type TestAccountTelemetryRulesParams struct {

	// Input.
	Input *models.TestAccountTelemetryRulesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the test account telemetry rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestAccountTelemetryRulesParams) WithDefaults() *TestAccountTelemetryRulesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the test account telemetry rules params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TestAccountTelemetryRulesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the test account telemetry rules params
func (o *TestAccountTelemetryRulesParams) WithTimeout(timeout time.Duration) *TestAccountTelemetryRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test account telemetry rules params
func (o *TestAccountTelemetryRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test account telemetry rules params
func (o *TestAccountTelemetryRulesParams) WithContext(ctx context.Context) *TestAccountTelemetryRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test account telemetry rules params
func (o *TestAccountTelemetryRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test account telemetry rules params
func (o *TestAccountTelemetryRulesParams) WithHTTPClient(client *http.Client) *TestAccountTelemetryRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test account telemetry rules params
func (o *TestAccountTelemetryRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the test account telemetry rules params
func (o *TestAccountTelemetryRulesParams) WithInput(input *models.TestAccountTelemetryRulesRequest) *TestAccountTelemetryRulesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the test account telemetry rules params
func (o *TestAccountTelemetryRulesParams) SetInput(input *models.TestAccountTelemetryRulesRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *TestAccountTelemetryRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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