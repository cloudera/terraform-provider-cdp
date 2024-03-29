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

// NewGetKeytabParams creates a new GetKeytabParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetKeytabParams() *GetKeytabParams {
	return &GetKeytabParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeytabParamsWithTimeout creates a new GetKeytabParams object
// with the ability to set a timeout on a request.
func NewGetKeytabParamsWithTimeout(timeout time.Duration) *GetKeytabParams {
	return &GetKeytabParams{
		timeout: timeout,
	}
}

// NewGetKeytabParamsWithContext creates a new GetKeytabParams object
// with the ability to set a context for a request.
func NewGetKeytabParamsWithContext(ctx context.Context) *GetKeytabParams {
	return &GetKeytabParams{
		Context: ctx,
	}
}

// NewGetKeytabParamsWithHTTPClient creates a new GetKeytabParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetKeytabParamsWithHTTPClient(client *http.Client) *GetKeytabParams {
	return &GetKeytabParams{
		HTTPClient: client,
	}
}

/*
GetKeytabParams contains all the parameters to send to the API endpoint

	for the get keytab operation.

	Typically these are written to a http.Request.
*/
type GetKeytabParams struct {

	// Input.
	Input *models.GetKeytabRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get keytab params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKeytabParams) WithDefaults() *GetKeytabParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get keytab params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetKeytabParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get keytab params
func (o *GetKeytabParams) WithTimeout(timeout time.Duration) *GetKeytabParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get keytab params
func (o *GetKeytabParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get keytab params
func (o *GetKeytabParams) WithContext(ctx context.Context) *GetKeytabParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get keytab params
func (o *GetKeytabParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get keytab params
func (o *GetKeytabParams) WithHTTPClient(client *http.Client) *GetKeytabParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get keytab params
func (o *GetKeytabParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get keytab params
func (o *GetKeytabParams) WithInput(input *models.GetKeytabRequest) *GetKeytabParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get keytab params
func (o *GetKeytabParams) SetInput(input *models.GetKeytabRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeytabParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
