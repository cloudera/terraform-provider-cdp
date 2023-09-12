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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// NewUnassignAzureCloudIdentityParams creates a new UnassignAzureCloudIdentityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnassignAzureCloudIdentityParams() *UnassignAzureCloudIdentityParams {
	return &UnassignAzureCloudIdentityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnassignAzureCloudIdentityParamsWithTimeout creates a new UnassignAzureCloudIdentityParams object
// with the ability to set a timeout on a request.
func NewUnassignAzureCloudIdentityParamsWithTimeout(timeout time.Duration) *UnassignAzureCloudIdentityParams {
	return &UnassignAzureCloudIdentityParams{
		timeout: timeout,
	}
}

// NewUnassignAzureCloudIdentityParamsWithContext creates a new UnassignAzureCloudIdentityParams object
// with the ability to set a context for a request.
func NewUnassignAzureCloudIdentityParamsWithContext(ctx context.Context) *UnassignAzureCloudIdentityParams {
	return &UnassignAzureCloudIdentityParams{
		Context: ctx,
	}
}

// NewUnassignAzureCloudIdentityParamsWithHTTPClient creates a new UnassignAzureCloudIdentityParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnassignAzureCloudIdentityParamsWithHTTPClient(client *http.Client) *UnassignAzureCloudIdentityParams {
	return &UnassignAzureCloudIdentityParams{
		HTTPClient: client,
	}
}

/*
UnassignAzureCloudIdentityParams contains all the parameters to send to the API endpoint

	for the unassign azure cloud identity operation.

	Typically these are written to a http.Request.
*/
type UnassignAzureCloudIdentityParams struct {

	// Input.
	Input *models.UnassignAzureCloudIdentityRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unassign azure cloud identity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnassignAzureCloudIdentityParams) WithDefaults() *UnassignAzureCloudIdentityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unassign azure cloud identity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnassignAzureCloudIdentityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unassign azure cloud identity params
func (o *UnassignAzureCloudIdentityParams) WithTimeout(timeout time.Duration) *UnassignAzureCloudIdentityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unassign azure cloud identity params
func (o *UnassignAzureCloudIdentityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unassign azure cloud identity params
func (o *UnassignAzureCloudIdentityParams) WithContext(ctx context.Context) *UnassignAzureCloudIdentityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unassign azure cloud identity params
func (o *UnassignAzureCloudIdentityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unassign azure cloud identity params
func (o *UnassignAzureCloudIdentityParams) WithHTTPClient(client *http.Client) *UnassignAzureCloudIdentityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unassign azure cloud identity params
func (o *UnassignAzureCloudIdentityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the unassign azure cloud identity params
func (o *UnassignAzureCloudIdentityParams) WithInput(input *models.UnassignAzureCloudIdentityRequest) *UnassignAzureCloudIdentityParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the unassign azure cloud identity params
func (o *UnassignAzureCloudIdentityParams) SetInput(input *models.UnassignAzureCloudIdentityRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UnassignAzureCloudIdentityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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