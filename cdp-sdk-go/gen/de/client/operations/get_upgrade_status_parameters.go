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

// NewGetUpgradeStatusParams creates a new GetUpgradeStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUpgradeStatusParams() *GetUpgradeStatusParams {
	return &GetUpgradeStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUpgradeStatusParamsWithTimeout creates a new GetUpgradeStatusParams object
// with the ability to set a timeout on a request.
func NewGetUpgradeStatusParamsWithTimeout(timeout time.Duration) *GetUpgradeStatusParams {
	return &GetUpgradeStatusParams{
		timeout: timeout,
	}
}

// NewGetUpgradeStatusParamsWithContext creates a new GetUpgradeStatusParams object
// with the ability to set a context for a request.
func NewGetUpgradeStatusParamsWithContext(ctx context.Context) *GetUpgradeStatusParams {
	return &GetUpgradeStatusParams{
		Context: ctx,
	}
}

// NewGetUpgradeStatusParamsWithHTTPClient creates a new GetUpgradeStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUpgradeStatusParamsWithHTTPClient(client *http.Client) *GetUpgradeStatusParams {
	return &GetUpgradeStatusParams{
		HTTPClient: client,
	}
}

/*
GetUpgradeStatusParams contains all the parameters to send to the API endpoint

	for the get upgrade status operation.

	Typically these are written to a http.Request.
*/
type GetUpgradeStatusParams struct {

	// Input.
	Input *models.GetUpgradeStatusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get upgrade status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpgradeStatusParams) WithDefaults() *GetUpgradeStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get upgrade status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUpgradeStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get upgrade status params
func (o *GetUpgradeStatusParams) WithTimeout(timeout time.Duration) *GetUpgradeStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get upgrade status params
func (o *GetUpgradeStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get upgrade status params
func (o *GetUpgradeStatusParams) WithContext(ctx context.Context) *GetUpgradeStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get upgrade status params
func (o *GetUpgradeStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get upgrade status params
func (o *GetUpgradeStatusParams) WithHTTPClient(client *http.Client) *GetUpgradeStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get upgrade status params
func (o *GetUpgradeStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get upgrade status params
func (o *GetUpgradeStatusParams) WithInput(input *models.GetUpgradeStatusRequest) *GetUpgradeStatusParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get upgrade status params
func (o *GetUpgradeStatusParams) SetInput(input *models.GetUpgradeStatusRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetUpgradeStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
