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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// NewGetClusterHostStatusParams creates a new GetClusterHostStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterHostStatusParams() *GetClusterHostStatusParams {
	return &GetClusterHostStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterHostStatusParamsWithTimeout creates a new GetClusterHostStatusParams object
// with the ability to set a timeout on a request.
func NewGetClusterHostStatusParamsWithTimeout(timeout time.Duration) *GetClusterHostStatusParams {
	return &GetClusterHostStatusParams{
		timeout: timeout,
	}
}

// NewGetClusterHostStatusParamsWithContext creates a new GetClusterHostStatusParams object
// with the ability to set a context for a request.
func NewGetClusterHostStatusParamsWithContext(ctx context.Context) *GetClusterHostStatusParams {
	return &GetClusterHostStatusParams{
		Context: ctx,
	}
}

// NewGetClusterHostStatusParamsWithHTTPClient creates a new GetClusterHostStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterHostStatusParamsWithHTTPClient(client *http.Client) *GetClusterHostStatusParams {
	return &GetClusterHostStatusParams{
		HTTPClient: client,
	}
}

/*
GetClusterHostStatusParams contains all the parameters to send to the API endpoint

	for the get cluster host status operation.

	Typically these are written to a http.Request.
*/
type GetClusterHostStatusParams struct {

	// Input.
	Input *models.GetClusterHostStatusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster host status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterHostStatusParams) WithDefaults() *GetClusterHostStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster host status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterHostStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cluster host status params
func (o *GetClusterHostStatusParams) WithTimeout(timeout time.Duration) *GetClusterHostStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster host status params
func (o *GetClusterHostStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster host status params
func (o *GetClusterHostStatusParams) WithContext(ctx context.Context) *GetClusterHostStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster host status params
func (o *GetClusterHostStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster host status params
func (o *GetClusterHostStatusParams) WithHTTPClient(client *http.Client) *GetClusterHostStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster host status params
func (o *GetClusterHostStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get cluster host status params
func (o *GetClusterHostStatusParams) WithInput(input *models.GetClusterHostStatusRequest) *GetClusterHostStatusParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get cluster host status params
func (o *GetClusterHostStatusParams) SetInput(input *models.GetClusterHostStatusRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterHostStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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