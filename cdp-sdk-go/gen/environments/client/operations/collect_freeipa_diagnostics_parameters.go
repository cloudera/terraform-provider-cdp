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

// NewCollectFreeipaDiagnosticsParams creates a new CollectFreeipaDiagnosticsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCollectFreeipaDiagnosticsParams() *CollectFreeipaDiagnosticsParams {
	return &CollectFreeipaDiagnosticsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCollectFreeipaDiagnosticsParamsWithTimeout creates a new CollectFreeipaDiagnosticsParams object
// with the ability to set a timeout on a request.
func NewCollectFreeipaDiagnosticsParamsWithTimeout(timeout time.Duration) *CollectFreeipaDiagnosticsParams {
	return &CollectFreeipaDiagnosticsParams{
		timeout: timeout,
	}
}

// NewCollectFreeipaDiagnosticsParamsWithContext creates a new CollectFreeipaDiagnosticsParams object
// with the ability to set a context for a request.
func NewCollectFreeipaDiagnosticsParamsWithContext(ctx context.Context) *CollectFreeipaDiagnosticsParams {
	return &CollectFreeipaDiagnosticsParams{
		Context: ctx,
	}
}

// NewCollectFreeipaDiagnosticsParamsWithHTTPClient creates a new CollectFreeipaDiagnosticsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCollectFreeipaDiagnosticsParamsWithHTTPClient(client *http.Client) *CollectFreeipaDiagnosticsParams {
	return &CollectFreeipaDiagnosticsParams{
		HTTPClient: client,
	}
}

/*
CollectFreeipaDiagnosticsParams contains all the parameters to send to the API endpoint

	for the collect freeipa diagnostics operation.

	Typically these are written to a http.Request.
*/
type CollectFreeipaDiagnosticsParams struct {

	// Input.
	Input *models.CollectFreeipaDiagnosticsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the collect freeipa diagnostics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CollectFreeipaDiagnosticsParams) WithDefaults() *CollectFreeipaDiagnosticsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the collect freeipa diagnostics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CollectFreeipaDiagnosticsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the collect freeipa diagnostics params
func (o *CollectFreeipaDiagnosticsParams) WithTimeout(timeout time.Duration) *CollectFreeipaDiagnosticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the collect freeipa diagnostics params
func (o *CollectFreeipaDiagnosticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the collect freeipa diagnostics params
func (o *CollectFreeipaDiagnosticsParams) WithContext(ctx context.Context) *CollectFreeipaDiagnosticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the collect freeipa diagnostics params
func (o *CollectFreeipaDiagnosticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the collect freeipa diagnostics params
func (o *CollectFreeipaDiagnosticsParams) WithHTTPClient(client *http.Client) *CollectFreeipaDiagnosticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the collect freeipa diagnostics params
func (o *CollectFreeipaDiagnosticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the collect freeipa diagnostics params
func (o *CollectFreeipaDiagnosticsParams) WithInput(input *models.CollectFreeipaDiagnosticsRequest) *CollectFreeipaDiagnosticsParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the collect freeipa diagnostics params
func (o *CollectFreeipaDiagnosticsParams) SetInput(input *models.CollectFreeipaDiagnosticsRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *CollectFreeipaDiagnosticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
