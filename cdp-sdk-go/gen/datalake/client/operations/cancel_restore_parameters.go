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

// NewCancelRestoreParams creates a new CancelRestoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelRestoreParams() *CancelRestoreParams {
	return &CancelRestoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelRestoreParamsWithTimeout creates a new CancelRestoreParams object
// with the ability to set a timeout on a request.
func NewCancelRestoreParamsWithTimeout(timeout time.Duration) *CancelRestoreParams {
	return &CancelRestoreParams{
		timeout: timeout,
	}
}

// NewCancelRestoreParamsWithContext creates a new CancelRestoreParams object
// with the ability to set a context for a request.
func NewCancelRestoreParamsWithContext(ctx context.Context) *CancelRestoreParams {
	return &CancelRestoreParams{
		Context: ctx,
	}
}

// NewCancelRestoreParamsWithHTTPClient creates a new CancelRestoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelRestoreParamsWithHTTPClient(client *http.Client) *CancelRestoreParams {
	return &CancelRestoreParams{
		HTTPClient: client,
	}
}

/*
CancelRestoreParams contains all the parameters to send to the API endpoint

	for the cancel restore operation.

	Typically these are written to a http.Request.
*/
type CancelRestoreParams struct {

	// Input.
	Input *models.CancelRestoreRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelRestoreParams) WithDefaults() *CancelRestoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel restore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelRestoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel restore params
func (o *CancelRestoreParams) WithTimeout(timeout time.Duration) *CancelRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel restore params
func (o *CancelRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel restore params
func (o *CancelRestoreParams) WithContext(ctx context.Context) *CancelRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel restore params
func (o *CancelRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel restore params
func (o *CancelRestoreParams) WithHTTPClient(client *http.Client) *CancelRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel restore params
func (o *CancelRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the cancel restore params
func (o *CancelRestoreParams) WithInput(input *models.CancelRestoreRequest) *CancelRestoreParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the cancel restore params
func (o *CancelRestoreParams) SetInput(input *models.CancelRestoreRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *CancelRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
