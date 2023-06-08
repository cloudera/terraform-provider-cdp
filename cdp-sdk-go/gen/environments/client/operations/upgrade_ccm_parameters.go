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

// NewUpgradeCcmParams creates a new UpgradeCcmParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpgradeCcmParams() *UpgradeCcmParams {
	return &UpgradeCcmParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeCcmParamsWithTimeout creates a new UpgradeCcmParams object
// with the ability to set a timeout on a request.
func NewUpgradeCcmParamsWithTimeout(timeout time.Duration) *UpgradeCcmParams {
	return &UpgradeCcmParams{
		timeout: timeout,
	}
}

// NewUpgradeCcmParamsWithContext creates a new UpgradeCcmParams object
// with the ability to set a context for a request.
func NewUpgradeCcmParamsWithContext(ctx context.Context) *UpgradeCcmParams {
	return &UpgradeCcmParams{
		Context: ctx,
	}
}

// NewUpgradeCcmParamsWithHTTPClient creates a new UpgradeCcmParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpgradeCcmParamsWithHTTPClient(client *http.Client) *UpgradeCcmParams {
	return &UpgradeCcmParams{
		HTTPClient: client,
	}
}

/*
UpgradeCcmParams contains all the parameters to send to the API endpoint

	for the upgrade ccm operation.

	Typically these are written to a http.Request.
*/
type UpgradeCcmParams struct {

	// Input.
	Input *models.UpgradeCcmRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upgrade ccm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeCcmParams) WithDefaults() *UpgradeCcmParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upgrade ccm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeCcmParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upgrade ccm params
func (o *UpgradeCcmParams) WithTimeout(timeout time.Duration) *UpgradeCcmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade ccm params
func (o *UpgradeCcmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade ccm params
func (o *UpgradeCcmParams) WithContext(ctx context.Context) *UpgradeCcmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade ccm params
func (o *UpgradeCcmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade ccm params
func (o *UpgradeCcmParams) WithHTTPClient(client *http.Client) *UpgradeCcmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade ccm params
func (o *UpgradeCcmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the upgrade ccm params
func (o *UpgradeCcmParams) WithInput(input *models.UpgradeCcmRequest) *UpgradeCcmParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the upgrade ccm params
func (o *UpgradeCcmParams) SetInput(input *models.UpgradeCcmRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeCcmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
