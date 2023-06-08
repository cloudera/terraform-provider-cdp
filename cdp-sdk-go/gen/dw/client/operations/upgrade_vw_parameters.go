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

// NewUpgradeVwParams creates a new UpgradeVwParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpgradeVwParams() *UpgradeVwParams {
	return &UpgradeVwParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeVwParamsWithTimeout creates a new UpgradeVwParams object
// with the ability to set a timeout on a request.
func NewUpgradeVwParamsWithTimeout(timeout time.Duration) *UpgradeVwParams {
	return &UpgradeVwParams{
		timeout: timeout,
	}
}

// NewUpgradeVwParamsWithContext creates a new UpgradeVwParams object
// with the ability to set a context for a request.
func NewUpgradeVwParamsWithContext(ctx context.Context) *UpgradeVwParams {
	return &UpgradeVwParams{
		Context: ctx,
	}
}

// NewUpgradeVwParamsWithHTTPClient creates a new UpgradeVwParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpgradeVwParamsWithHTTPClient(client *http.Client) *UpgradeVwParams {
	return &UpgradeVwParams{
		HTTPClient: client,
	}
}

/*
UpgradeVwParams contains all the parameters to send to the API endpoint

	for the upgrade vw operation.

	Typically these are written to a http.Request.
*/
type UpgradeVwParams struct {

	// Input.
	Input *models.UpgradeVwRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upgrade vw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeVwParams) WithDefaults() *UpgradeVwParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upgrade vw params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeVwParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upgrade vw params
func (o *UpgradeVwParams) WithTimeout(timeout time.Duration) *UpgradeVwParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade vw params
func (o *UpgradeVwParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade vw params
func (o *UpgradeVwParams) WithContext(ctx context.Context) *UpgradeVwParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade vw params
func (o *UpgradeVwParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade vw params
func (o *UpgradeVwParams) WithHTTPClient(client *http.Client) *UpgradeVwParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade vw params
func (o *UpgradeVwParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the upgrade vw params
func (o *UpgradeVwParams) WithInput(input *models.UpgradeVwRequest) *UpgradeVwParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the upgrade vw params
func (o *UpgradeVwParams) SetInput(input *models.UpgradeVwRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeVwParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
