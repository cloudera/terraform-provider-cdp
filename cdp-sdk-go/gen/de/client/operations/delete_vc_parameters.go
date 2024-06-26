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

// NewDeleteVcParams creates a new DeleteVcParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVcParams() *DeleteVcParams {
	return &DeleteVcParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVcParamsWithTimeout creates a new DeleteVcParams object
// with the ability to set a timeout on a request.
func NewDeleteVcParamsWithTimeout(timeout time.Duration) *DeleteVcParams {
	return &DeleteVcParams{
		timeout: timeout,
	}
}

// NewDeleteVcParamsWithContext creates a new DeleteVcParams object
// with the ability to set a context for a request.
func NewDeleteVcParamsWithContext(ctx context.Context) *DeleteVcParams {
	return &DeleteVcParams{
		Context: ctx,
	}
}

// NewDeleteVcParamsWithHTTPClient creates a new DeleteVcParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVcParamsWithHTTPClient(client *http.Client) *DeleteVcParams {
	return &DeleteVcParams{
		HTTPClient: client,
	}
}

/*
DeleteVcParams contains all the parameters to send to the API endpoint

	for the delete vc operation.

	Typically these are written to a http.Request.
*/
type DeleteVcParams struct {

	// Input.
	Input *models.DeleteVcRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete vc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVcParams) WithDefaults() *DeleteVcParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete vc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVcParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete vc params
func (o *DeleteVcParams) WithTimeout(timeout time.Duration) *DeleteVcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete vc params
func (o *DeleteVcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete vc params
func (o *DeleteVcParams) WithContext(ctx context.Context) *DeleteVcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete vc params
func (o *DeleteVcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete vc params
func (o *DeleteVcParams) WithHTTPClient(client *http.Client) *DeleteVcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete vc params
func (o *DeleteVcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the delete vc params
func (o *DeleteVcParams) WithInput(input *models.DeleteVcRequest) *DeleteVcParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the delete vc params
func (o *DeleteVcParams) SetInput(input *models.DeleteVcRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
