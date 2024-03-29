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

// NewListDatalakeDiagnosticsParams creates a new ListDatalakeDiagnosticsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDatalakeDiagnosticsParams() *ListDatalakeDiagnosticsParams {
	return &ListDatalakeDiagnosticsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDatalakeDiagnosticsParamsWithTimeout creates a new ListDatalakeDiagnosticsParams object
// with the ability to set a timeout on a request.
func NewListDatalakeDiagnosticsParamsWithTimeout(timeout time.Duration) *ListDatalakeDiagnosticsParams {
	return &ListDatalakeDiagnosticsParams{
		timeout: timeout,
	}
}

// NewListDatalakeDiagnosticsParamsWithContext creates a new ListDatalakeDiagnosticsParams object
// with the ability to set a context for a request.
func NewListDatalakeDiagnosticsParamsWithContext(ctx context.Context) *ListDatalakeDiagnosticsParams {
	return &ListDatalakeDiagnosticsParams{
		Context: ctx,
	}
}

// NewListDatalakeDiagnosticsParamsWithHTTPClient creates a new ListDatalakeDiagnosticsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDatalakeDiagnosticsParamsWithHTTPClient(client *http.Client) *ListDatalakeDiagnosticsParams {
	return &ListDatalakeDiagnosticsParams{
		HTTPClient: client,
	}
}

/*
ListDatalakeDiagnosticsParams contains all the parameters to send to the API endpoint

	for the list datalake diagnostics operation.

	Typically these are written to a http.Request.
*/
type ListDatalakeDiagnosticsParams struct {

	// Input.
	Input *models.ListDatalakeDiagnosticsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list datalake diagnostics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDatalakeDiagnosticsParams) WithDefaults() *ListDatalakeDiagnosticsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list datalake diagnostics params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDatalakeDiagnosticsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list datalake diagnostics params
func (o *ListDatalakeDiagnosticsParams) WithTimeout(timeout time.Duration) *ListDatalakeDiagnosticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list datalake diagnostics params
func (o *ListDatalakeDiagnosticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list datalake diagnostics params
func (o *ListDatalakeDiagnosticsParams) WithContext(ctx context.Context) *ListDatalakeDiagnosticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list datalake diagnostics params
func (o *ListDatalakeDiagnosticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list datalake diagnostics params
func (o *ListDatalakeDiagnosticsParams) WithHTTPClient(client *http.Client) *ListDatalakeDiagnosticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list datalake diagnostics params
func (o *ListDatalakeDiagnosticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list datalake diagnostics params
func (o *ListDatalakeDiagnosticsParams) WithInput(input *models.ListDatalakeDiagnosticsRequest) *ListDatalakeDiagnosticsParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list datalake diagnostics params
func (o *ListDatalakeDiagnosticsParams) SetInput(input *models.ListDatalakeDiagnosticsRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListDatalakeDiagnosticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
