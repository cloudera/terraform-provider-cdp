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

// NewListWorkspacesParams creates a new ListWorkspacesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListWorkspacesParams() *ListWorkspacesParams {
	return &ListWorkspacesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListWorkspacesParamsWithTimeout creates a new ListWorkspacesParams object
// with the ability to set a timeout on a request.
func NewListWorkspacesParamsWithTimeout(timeout time.Duration) *ListWorkspacesParams {
	return &ListWorkspacesParams{
		timeout: timeout,
	}
}

// NewListWorkspacesParamsWithContext creates a new ListWorkspacesParams object
// with the ability to set a context for a request.
func NewListWorkspacesParamsWithContext(ctx context.Context) *ListWorkspacesParams {
	return &ListWorkspacesParams{
		Context: ctx,
	}
}

// NewListWorkspacesParamsWithHTTPClient creates a new ListWorkspacesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListWorkspacesParamsWithHTTPClient(client *http.Client) *ListWorkspacesParams {
	return &ListWorkspacesParams{
		HTTPClient: client,
	}
}

/*
ListWorkspacesParams contains all the parameters to send to the API endpoint

	for the list workspaces operation.

	Typically these are written to a http.Request.
*/
type ListWorkspacesParams struct {

	// Input.
	Input *models.ListWorkspacesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list workspaces params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListWorkspacesParams) WithDefaults() *ListWorkspacesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list workspaces params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListWorkspacesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list workspaces params
func (o *ListWorkspacesParams) WithTimeout(timeout time.Duration) *ListWorkspacesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list workspaces params
func (o *ListWorkspacesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list workspaces params
func (o *ListWorkspacesParams) WithContext(ctx context.Context) *ListWorkspacesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list workspaces params
func (o *ListWorkspacesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list workspaces params
func (o *ListWorkspacesParams) WithHTTPClient(client *http.Client) *ListWorkspacesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list workspaces params
func (o *ListWorkspacesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list workspaces params
func (o *ListWorkspacesParams) WithInput(input *models.ListWorkspacesRequest) *ListWorkspacesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list workspaces params
func (o *ListWorkspacesParams) SetInput(input *models.ListWorkspacesRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListWorkspacesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
