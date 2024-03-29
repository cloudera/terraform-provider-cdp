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

// NewRevokeModelRegistryAccessParams creates a new RevokeModelRegistryAccessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevokeModelRegistryAccessParams() *RevokeModelRegistryAccessParams {
	return &RevokeModelRegistryAccessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeModelRegistryAccessParamsWithTimeout creates a new RevokeModelRegistryAccessParams object
// with the ability to set a timeout on a request.
func NewRevokeModelRegistryAccessParamsWithTimeout(timeout time.Duration) *RevokeModelRegistryAccessParams {
	return &RevokeModelRegistryAccessParams{
		timeout: timeout,
	}
}

// NewRevokeModelRegistryAccessParamsWithContext creates a new RevokeModelRegistryAccessParams object
// with the ability to set a context for a request.
func NewRevokeModelRegistryAccessParamsWithContext(ctx context.Context) *RevokeModelRegistryAccessParams {
	return &RevokeModelRegistryAccessParams{
		Context: ctx,
	}
}

// NewRevokeModelRegistryAccessParamsWithHTTPClient creates a new RevokeModelRegistryAccessParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevokeModelRegistryAccessParamsWithHTTPClient(client *http.Client) *RevokeModelRegistryAccessParams {
	return &RevokeModelRegistryAccessParams{
		HTTPClient: client,
	}
}

/*
RevokeModelRegistryAccessParams contains all the parameters to send to the API endpoint

	for the revoke model registry access operation.

	Typically these are written to a http.Request.
*/
type RevokeModelRegistryAccessParams struct {

	// Input.
	Input *models.RevokeModelRegistryAccessRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revoke model registry access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeModelRegistryAccessParams) WithDefaults() *RevokeModelRegistryAccessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revoke model registry access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeModelRegistryAccessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revoke model registry access params
func (o *RevokeModelRegistryAccessParams) WithTimeout(timeout time.Duration) *RevokeModelRegistryAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke model registry access params
func (o *RevokeModelRegistryAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke model registry access params
func (o *RevokeModelRegistryAccessParams) WithContext(ctx context.Context) *RevokeModelRegistryAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke model registry access params
func (o *RevokeModelRegistryAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke model registry access params
func (o *RevokeModelRegistryAccessParams) WithHTTPClient(client *http.Client) *RevokeModelRegistryAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke model registry access params
func (o *RevokeModelRegistryAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the revoke model registry access params
func (o *RevokeModelRegistryAccessParams) WithInput(input *models.RevokeModelRegistryAccessRequest) *RevokeModelRegistryAccessParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the revoke model registry access params
func (o *RevokeModelRegistryAccessParams) SetInput(input *models.RevokeModelRegistryAccessRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeModelRegistryAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
