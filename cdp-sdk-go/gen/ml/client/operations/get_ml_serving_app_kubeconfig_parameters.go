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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// NewGetMlServingAppKubeconfigParams creates a new GetMlServingAppKubeconfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMlServingAppKubeconfigParams() *GetMlServingAppKubeconfigParams {
	return &GetMlServingAppKubeconfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMlServingAppKubeconfigParamsWithTimeout creates a new GetMlServingAppKubeconfigParams object
// with the ability to set a timeout on a request.
func NewGetMlServingAppKubeconfigParamsWithTimeout(timeout time.Duration) *GetMlServingAppKubeconfigParams {
	return &GetMlServingAppKubeconfigParams{
		timeout: timeout,
	}
}

// NewGetMlServingAppKubeconfigParamsWithContext creates a new GetMlServingAppKubeconfigParams object
// with the ability to set a context for a request.
func NewGetMlServingAppKubeconfigParamsWithContext(ctx context.Context) *GetMlServingAppKubeconfigParams {
	return &GetMlServingAppKubeconfigParams{
		Context: ctx,
	}
}

// NewGetMlServingAppKubeconfigParamsWithHTTPClient creates a new GetMlServingAppKubeconfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMlServingAppKubeconfigParamsWithHTTPClient(client *http.Client) *GetMlServingAppKubeconfigParams {
	return &GetMlServingAppKubeconfigParams{
		HTTPClient: client,
	}
}

/*
GetMlServingAppKubeconfigParams contains all the parameters to send to the API endpoint

	for the get ml serving app kubeconfig operation.

	Typically these are written to a http.Request.
*/
type GetMlServingAppKubeconfigParams struct {

	// Input.
	Input *models.GetMlServingAppKubeconfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ml serving app kubeconfig params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMlServingAppKubeconfigParams) WithDefaults() *GetMlServingAppKubeconfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ml serving app kubeconfig params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMlServingAppKubeconfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ml serving app kubeconfig params
func (o *GetMlServingAppKubeconfigParams) WithTimeout(timeout time.Duration) *GetMlServingAppKubeconfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ml serving app kubeconfig params
func (o *GetMlServingAppKubeconfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ml serving app kubeconfig params
func (o *GetMlServingAppKubeconfigParams) WithContext(ctx context.Context) *GetMlServingAppKubeconfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ml serving app kubeconfig params
func (o *GetMlServingAppKubeconfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ml serving app kubeconfig params
func (o *GetMlServingAppKubeconfigParams) WithHTTPClient(client *http.Client) *GetMlServingAppKubeconfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ml serving app kubeconfig params
func (o *GetMlServingAppKubeconfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the get ml serving app kubeconfig params
func (o *GetMlServingAppKubeconfigParams) WithInput(input *models.GetMlServingAppKubeconfigRequest) *GetMlServingAppKubeconfigParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the get ml serving app kubeconfig params
func (o *GetMlServingAppKubeconfigParams) SetInput(input *models.GetMlServingAppKubeconfigRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GetMlServingAppKubeconfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
