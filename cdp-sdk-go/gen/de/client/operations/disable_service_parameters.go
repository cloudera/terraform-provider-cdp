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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/de/models"
)

// NewDisableServiceParams creates a new DisableServiceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDisableServiceParams() *DisableServiceParams {
	return &DisableServiceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDisableServiceParamsWithTimeout creates a new DisableServiceParams object
// with the ability to set a timeout on a request.
func NewDisableServiceParamsWithTimeout(timeout time.Duration) *DisableServiceParams {
	return &DisableServiceParams{
		timeout: timeout,
	}
}

// NewDisableServiceParamsWithContext creates a new DisableServiceParams object
// with the ability to set a context for a request.
func NewDisableServiceParamsWithContext(ctx context.Context) *DisableServiceParams {
	return &DisableServiceParams{
		Context: ctx,
	}
}

// NewDisableServiceParamsWithHTTPClient creates a new DisableServiceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDisableServiceParamsWithHTTPClient(client *http.Client) *DisableServiceParams {
	return &DisableServiceParams{
		HTTPClient: client,
	}
}

/*
DisableServiceParams contains all the parameters to send to the API endpoint

	for the disable service operation.

	Typically these are written to a http.Request.
*/
type DisableServiceParams struct {

	// Input.
	Input *models.DisableServiceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disable service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableServiceParams) WithDefaults() *DisableServiceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disable service params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DisableServiceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disable service params
func (o *DisableServiceParams) WithTimeout(timeout time.Duration) *DisableServiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable service params
func (o *DisableServiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable service params
func (o *DisableServiceParams) WithContext(ctx context.Context) *DisableServiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable service params
func (o *DisableServiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable service params
func (o *DisableServiceParams) WithHTTPClient(client *http.Client) *DisableServiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable service params
func (o *DisableServiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the disable service params
func (o *DisableServiceParams) WithInput(input *models.DisableServiceRequest) *DisableServiceParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the disable service params
func (o *DisableServiceParams) SetInput(input *models.DisableServiceRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DisableServiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
