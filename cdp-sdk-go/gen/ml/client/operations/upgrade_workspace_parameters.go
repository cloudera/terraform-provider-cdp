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

// NewUpgradeWorkspaceParams creates a new UpgradeWorkspaceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpgradeWorkspaceParams() *UpgradeWorkspaceParams {
	return &UpgradeWorkspaceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeWorkspaceParamsWithTimeout creates a new UpgradeWorkspaceParams object
// with the ability to set a timeout on a request.
func NewUpgradeWorkspaceParamsWithTimeout(timeout time.Duration) *UpgradeWorkspaceParams {
	return &UpgradeWorkspaceParams{
		timeout: timeout,
	}
}

// NewUpgradeWorkspaceParamsWithContext creates a new UpgradeWorkspaceParams object
// with the ability to set a context for a request.
func NewUpgradeWorkspaceParamsWithContext(ctx context.Context) *UpgradeWorkspaceParams {
	return &UpgradeWorkspaceParams{
		Context: ctx,
	}
}

// NewUpgradeWorkspaceParamsWithHTTPClient creates a new UpgradeWorkspaceParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpgradeWorkspaceParamsWithHTTPClient(client *http.Client) *UpgradeWorkspaceParams {
	return &UpgradeWorkspaceParams{
		HTTPClient: client,
	}
}

/*
UpgradeWorkspaceParams contains all the parameters to send to the API endpoint

	for the upgrade workspace operation.

	Typically these are written to a http.Request.
*/
type UpgradeWorkspaceParams struct {

	// Input.
	Input *models.UpgradeWorkspaceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upgrade workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeWorkspaceParams) WithDefaults() *UpgradeWorkspaceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upgrade workspace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeWorkspaceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upgrade workspace params
func (o *UpgradeWorkspaceParams) WithTimeout(timeout time.Duration) *UpgradeWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade workspace params
func (o *UpgradeWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade workspace params
func (o *UpgradeWorkspaceParams) WithContext(ctx context.Context) *UpgradeWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade workspace params
func (o *UpgradeWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade workspace params
func (o *UpgradeWorkspaceParams) WithHTTPClient(client *http.Client) *UpgradeWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade workspace params
func (o *UpgradeWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the upgrade workspace params
func (o *UpgradeWorkspaceParams) WithInput(input *models.UpgradeWorkspaceRequest) *UpgradeWorkspaceParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the upgrade workspace params
func (o *UpgradeWorkspaceParams) SetInput(input *models.UpgradeWorkspaceRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
