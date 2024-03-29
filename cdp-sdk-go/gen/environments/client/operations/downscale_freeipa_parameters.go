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

// NewDownscaleFreeipaParams creates a new DownscaleFreeipaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownscaleFreeipaParams() *DownscaleFreeipaParams {
	return &DownscaleFreeipaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownscaleFreeipaParamsWithTimeout creates a new DownscaleFreeipaParams object
// with the ability to set a timeout on a request.
func NewDownscaleFreeipaParamsWithTimeout(timeout time.Duration) *DownscaleFreeipaParams {
	return &DownscaleFreeipaParams{
		timeout: timeout,
	}
}

// NewDownscaleFreeipaParamsWithContext creates a new DownscaleFreeipaParams object
// with the ability to set a context for a request.
func NewDownscaleFreeipaParamsWithContext(ctx context.Context) *DownscaleFreeipaParams {
	return &DownscaleFreeipaParams{
		Context: ctx,
	}
}

// NewDownscaleFreeipaParamsWithHTTPClient creates a new DownscaleFreeipaParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownscaleFreeipaParamsWithHTTPClient(client *http.Client) *DownscaleFreeipaParams {
	return &DownscaleFreeipaParams{
		HTTPClient: client,
	}
}

/*
DownscaleFreeipaParams contains all the parameters to send to the API endpoint

	for the downscale freeipa operation.

	Typically these are written to a http.Request.
*/
type DownscaleFreeipaParams struct {

	// Input.
	Input *models.DownscaleFreeipaRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the downscale freeipa params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownscaleFreeipaParams) WithDefaults() *DownscaleFreeipaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the downscale freeipa params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownscaleFreeipaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the downscale freeipa params
func (o *DownscaleFreeipaParams) WithTimeout(timeout time.Duration) *DownscaleFreeipaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the downscale freeipa params
func (o *DownscaleFreeipaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the downscale freeipa params
func (o *DownscaleFreeipaParams) WithContext(ctx context.Context) *DownscaleFreeipaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the downscale freeipa params
func (o *DownscaleFreeipaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the downscale freeipa params
func (o *DownscaleFreeipaParams) WithHTTPClient(client *http.Client) *DownscaleFreeipaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the downscale freeipa params
func (o *DownscaleFreeipaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the downscale freeipa params
func (o *DownscaleFreeipaParams) WithInput(input *models.DownscaleFreeipaRequest) *DownscaleFreeipaParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the downscale freeipa params
func (o *DownscaleFreeipaParams) SetInput(input *models.DownscaleFreeipaRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DownscaleFreeipaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
