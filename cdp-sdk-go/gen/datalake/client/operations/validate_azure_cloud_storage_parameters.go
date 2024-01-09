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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// NewValidateAzureCloudStorageParams creates a new ValidateAzureCloudStorageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateAzureCloudStorageParams() *ValidateAzureCloudStorageParams {
	return &ValidateAzureCloudStorageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateAzureCloudStorageParamsWithTimeout creates a new ValidateAzureCloudStorageParams object
// with the ability to set a timeout on a request.
func NewValidateAzureCloudStorageParamsWithTimeout(timeout time.Duration) *ValidateAzureCloudStorageParams {
	return &ValidateAzureCloudStorageParams{
		timeout: timeout,
	}
}

// NewValidateAzureCloudStorageParamsWithContext creates a new ValidateAzureCloudStorageParams object
// with the ability to set a context for a request.
func NewValidateAzureCloudStorageParamsWithContext(ctx context.Context) *ValidateAzureCloudStorageParams {
	return &ValidateAzureCloudStorageParams{
		Context: ctx,
	}
}

// NewValidateAzureCloudStorageParamsWithHTTPClient creates a new ValidateAzureCloudStorageParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateAzureCloudStorageParamsWithHTTPClient(client *http.Client) *ValidateAzureCloudStorageParams {
	return &ValidateAzureCloudStorageParams{
		HTTPClient: client,
	}
}

/*
ValidateAzureCloudStorageParams contains all the parameters to send to the API endpoint

	for the validate azure cloud storage operation.

	Typically these are written to a http.Request.
*/
type ValidateAzureCloudStorageParams struct {

	// Input.
	Input *models.ValidateAzureCloudStorageRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate azure cloud storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateAzureCloudStorageParams) WithDefaults() *ValidateAzureCloudStorageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate azure cloud storage params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateAzureCloudStorageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate azure cloud storage params
func (o *ValidateAzureCloudStorageParams) WithTimeout(timeout time.Duration) *ValidateAzureCloudStorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate azure cloud storage params
func (o *ValidateAzureCloudStorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate azure cloud storage params
func (o *ValidateAzureCloudStorageParams) WithContext(ctx context.Context) *ValidateAzureCloudStorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate azure cloud storage params
func (o *ValidateAzureCloudStorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate azure cloud storage params
func (o *ValidateAzureCloudStorageParams) WithHTTPClient(client *http.Client) *ValidateAzureCloudStorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate azure cloud storage params
func (o *ValidateAzureCloudStorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the validate azure cloud storage params
func (o *ValidateAzureCloudStorageParams) WithInput(input *models.ValidateAzureCloudStorageRequest) *ValidateAzureCloudStorageParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the validate azure cloud storage params
func (o *ValidateAzureCloudStorageParams) SetInput(input *models.ValidateAzureCloudStorageRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateAzureCloudStorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
