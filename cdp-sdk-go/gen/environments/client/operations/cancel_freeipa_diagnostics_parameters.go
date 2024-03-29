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

// NewCancelFreeipaDiagnosticsParams creates a new CancelFreeipaDiagnosticsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelFreeipaDiagnosticsParams() *CancelFreeipaDiagnosticsParams {
	return &CancelFreeipaDiagnosticsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelFreeipaDiagnosticsParamsWithTimeout creates a new CancelFreeipaDiagnosticsParams object
// with the ability to set a timeout on a request.
func NewCancelFreeipaDiagnosticsParamsWithTimeout(timeout time.Duration) *CancelFreeipaDiagnosticsParams {
	return &CancelFreeipaDiagnosticsParams{
		timeout: timeout,
	}
}

// NewCancelFreeipaDiagnosticsParamsWithContext creates a new CancelFreeipaDiagnosticsParams object
// with the ability to set a context for a request.
func NewCancelFreeipaDiagnosticsParamsWithContext(ctx context.Context) *CancelFreeipaDiagnosticsParams {
	return &CancelFreeipaDiagnosticsParams{
		Context: ctx,
	}
}

// NewCancelFreeipaDiagnosticsParamsWithHTTPClient creates a new CancelFreeipaDiagnosticsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelFreeipaDiagnosticsParamsWithHTTPClient(client *http.Client) *CancelFreeipaDiagnosticsParams {
	return &CancelFreeipaDiagnosticsParams{
		HTTPClient: client,
	}
}

/*
CancelFreeipaDiagnosticsParams contains all the parameters to send to the API endpoint

	for the cancel freeipa diagnostics operation.

	Typically these are written to a http.Request.
*/
type CancelFreeipaDiagnosticsParams struct {

	// Input.
	Input *models.CancelFreeipaDiagnosticsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel freeipa diagnostics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelFreeipaDiagnosticsParams) WithDefaults() *CancelFreeipaDiagnosticsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel freeipa diagnostics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelFreeipaDiagnosticsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel freeipa diagnostics params
func (o *CancelFreeipaDiagnosticsParams) WithTimeout(timeout time.Duration) *CancelFreeipaDiagnosticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel freeipa diagnostics params
func (o *CancelFreeipaDiagnosticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel freeipa diagnostics params
func (o *CancelFreeipaDiagnosticsParams) WithContext(ctx context.Context) *CancelFreeipaDiagnosticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel freeipa diagnostics params
func (o *CancelFreeipaDiagnosticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel freeipa diagnostics params
func (o *CancelFreeipaDiagnosticsParams) WithHTTPClient(client *http.Client) *CancelFreeipaDiagnosticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel freeipa diagnostics params
func (o *CancelFreeipaDiagnosticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the cancel freeipa diagnostics params
func (o *CancelFreeipaDiagnosticsParams) WithInput(input *models.CancelFreeipaDiagnosticsRequest) *CancelFreeipaDiagnosticsParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the cancel freeipa diagnostics params
func (o *CancelFreeipaDiagnosticsParams) SetInput(input *models.CancelFreeipaDiagnosticsRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *CancelFreeipaDiagnosticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
