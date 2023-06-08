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

// NewSuspendVwParams creates a new SuspendVwParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSuspendVwParams() *SuspendVwParams {
	return &SuspendVwParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSuspendVwParamsWithTimeout creates a new SuspendVwParams object
// with the ability to set a timeout on a request.
func NewSuspendVwParamsWithTimeout(timeout time.Duration) *SuspendVwParams {
	return &SuspendVwParams{
		timeout: timeout,
	}
}

// NewSuspendVwParamsWithContext creates a new SuspendVwParams object
// with the ability to set a context for a request.
func NewSuspendVwParamsWithContext(ctx context.Context) *SuspendVwParams {
	return &SuspendVwParams{
		Context: ctx,
	}
}

// NewSuspendVwParamsWithHTTPClient creates a new SuspendVwParams object
// with the ability to set a custom HTTPClient for a request.
func NewSuspendVwParamsWithHTTPClient(client *http.Client) *SuspendVwParams {
	return &SuspendVwParams{
		HTTPClient: client,
	}
}

/*
SuspendVwParams contains all the parameters to send to the API endpoint

	for the suspend vw operation.

	Typically these are written to a http.Request.
*/
type SuspendVwParams struct {

	// Input.
	Input *models.SuspendVwRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the suspend vw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuspendVwParams) WithDefaults() *SuspendVwParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the suspend vw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuspendVwParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the suspend vw params
func (o *SuspendVwParams) WithTimeout(timeout time.Duration) *SuspendVwParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the suspend vw params
func (o *SuspendVwParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the suspend vw params
func (o *SuspendVwParams) WithContext(ctx context.Context) *SuspendVwParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the suspend vw params
func (o *SuspendVwParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the suspend vw params
func (o *SuspendVwParams) WithHTTPClient(client *http.Client) *SuspendVwParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the suspend vw params
func (o *SuspendVwParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the suspend vw params
func (o *SuspendVwParams) WithInput(input *models.SuspendVwRequest) *SuspendVwParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the suspend vw params
func (o *SuspendVwParams) SetInput(input *models.SuspendVwRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *SuspendVwParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
