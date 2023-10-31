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

// NewGrantWorkspaceAccessParams creates a new GrantWorkspaceAccessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGrantWorkspaceAccessParams() *GrantWorkspaceAccessParams {
	return &GrantWorkspaceAccessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGrantWorkspaceAccessParamsWithTimeout creates a new GrantWorkspaceAccessParams object
// with the ability to set a timeout on a request.
func NewGrantWorkspaceAccessParamsWithTimeout(timeout time.Duration) *GrantWorkspaceAccessParams {
	return &GrantWorkspaceAccessParams{
		timeout: timeout,
	}
}

// NewGrantWorkspaceAccessParamsWithContext creates a new GrantWorkspaceAccessParams object
// with the ability to set a context for a request.
func NewGrantWorkspaceAccessParamsWithContext(ctx context.Context) *GrantWorkspaceAccessParams {
	return &GrantWorkspaceAccessParams{
		Context: ctx,
	}
}

// NewGrantWorkspaceAccessParamsWithHTTPClient creates a new GrantWorkspaceAccessParams object
// with the ability to set a custom HTTPClient for a request.
func NewGrantWorkspaceAccessParamsWithHTTPClient(client *http.Client) *GrantWorkspaceAccessParams {
	return &GrantWorkspaceAccessParams{
		HTTPClient: client,
	}
}

/*
GrantWorkspaceAccessParams contains all the parameters to send to the API endpoint

	for the grant workspace access operation.

	Typically these are written to a http.Request.
*/
type GrantWorkspaceAccessParams struct {

	// Input.
	Input *models.GrantWorkspaceAccessRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the grant workspace access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GrantWorkspaceAccessParams) WithDefaults() *GrantWorkspaceAccessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the grant workspace access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GrantWorkspaceAccessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the grant workspace access params
func (o *GrantWorkspaceAccessParams) WithTimeout(timeout time.Duration) *GrantWorkspaceAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the grant workspace access params
func (o *GrantWorkspaceAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the grant workspace access params
func (o *GrantWorkspaceAccessParams) WithContext(ctx context.Context) *GrantWorkspaceAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the grant workspace access params
func (o *GrantWorkspaceAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the grant workspace access params
func (o *GrantWorkspaceAccessParams) WithHTTPClient(client *http.Client) *GrantWorkspaceAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the grant workspace access params
func (o *GrantWorkspaceAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the grant workspace access params
func (o *GrantWorkspaceAccessParams) WithInput(input *models.GrantWorkspaceAccessRequest) *GrantWorkspaceAccessParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the grant workspace access params
func (o *GrantWorkspaceAccessParams) SetInput(input *models.GrantWorkspaceAccessRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *GrantWorkspaceAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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