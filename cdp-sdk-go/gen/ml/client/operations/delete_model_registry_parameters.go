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

// NewDeleteModelRegistryParams creates a new DeleteModelRegistryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteModelRegistryParams() *DeleteModelRegistryParams {
	return &DeleteModelRegistryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteModelRegistryParamsWithTimeout creates a new DeleteModelRegistryParams object
// with the ability to set a timeout on a request.
func NewDeleteModelRegistryParamsWithTimeout(timeout time.Duration) *DeleteModelRegistryParams {
	return &DeleteModelRegistryParams{
		timeout: timeout,
	}
}

// NewDeleteModelRegistryParamsWithContext creates a new DeleteModelRegistryParams object
// with the ability to set a context for a request.
func NewDeleteModelRegistryParamsWithContext(ctx context.Context) *DeleteModelRegistryParams {
	return &DeleteModelRegistryParams{
		Context: ctx,
	}
}

// NewDeleteModelRegistryParamsWithHTTPClient creates a new DeleteModelRegistryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteModelRegistryParamsWithHTTPClient(client *http.Client) *DeleteModelRegistryParams {
	return &DeleteModelRegistryParams{
		HTTPClient: client,
	}
}

/*
DeleteModelRegistryParams contains all the parameters to send to the API endpoint

	for the delete model registry operation.

	Typically these are written to a http.Request.
*/
type DeleteModelRegistryParams struct {

	// Input.
	Input *models.DeleteModelRegistryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete model registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteModelRegistryParams) WithDefaults() *DeleteModelRegistryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete model registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteModelRegistryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete model registry params
func (o *DeleteModelRegistryParams) WithTimeout(timeout time.Duration) *DeleteModelRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete model registry params
func (o *DeleteModelRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete model registry params
func (o *DeleteModelRegistryParams) WithContext(ctx context.Context) *DeleteModelRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete model registry params
func (o *DeleteModelRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete model registry params
func (o *DeleteModelRegistryParams) WithHTTPClient(client *http.Client) *DeleteModelRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete model registry params
func (o *DeleteModelRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the delete model registry params
func (o *DeleteModelRegistryParams) WithInput(input *models.DeleteModelRegistryRequest) *DeleteModelRegistryParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the delete model registry params
func (o *DeleteModelRegistryParams) SetInput(input *models.DeleteModelRegistryRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteModelRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
