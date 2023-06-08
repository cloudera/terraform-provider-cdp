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

// NewRestoreDatalakeStatusParams creates a new RestoreDatalakeStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestoreDatalakeStatusParams() *RestoreDatalakeStatusParams {
	return &RestoreDatalakeStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreDatalakeStatusParamsWithTimeout creates a new RestoreDatalakeStatusParams object
// with the ability to set a timeout on a request.
func NewRestoreDatalakeStatusParamsWithTimeout(timeout time.Duration) *RestoreDatalakeStatusParams {
	return &RestoreDatalakeStatusParams{
		timeout: timeout,
	}
}

// NewRestoreDatalakeStatusParamsWithContext creates a new RestoreDatalakeStatusParams object
// with the ability to set a context for a request.
func NewRestoreDatalakeStatusParamsWithContext(ctx context.Context) *RestoreDatalakeStatusParams {
	return &RestoreDatalakeStatusParams{
		Context: ctx,
	}
}

// NewRestoreDatalakeStatusParamsWithHTTPClient creates a new RestoreDatalakeStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestoreDatalakeStatusParamsWithHTTPClient(client *http.Client) *RestoreDatalakeStatusParams {
	return &RestoreDatalakeStatusParams{
		HTTPClient: client,
	}
}

/*
RestoreDatalakeStatusParams contains all the parameters to send to the API endpoint

	for the restore datalake status operation.

	Typically these are written to a http.Request.
*/
type RestoreDatalakeStatusParams struct {

	// Input.
	Input *models.RestoreDatalakeStatusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restore datalake status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreDatalakeStatusParams) WithDefaults() *RestoreDatalakeStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restore datalake status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreDatalakeStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restore datalake status params
func (o *RestoreDatalakeStatusParams) WithTimeout(timeout time.Duration) *RestoreDatalakeStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore datalake status params
func (o *RestoreDatalakeStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore datalake status params
func (o *RestoreDatalakeStatusParams) WithContext(ctx context.Context) *RestoreDatalakeStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore datalake status params
func (o *RestoreDatalakeStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore datalake status params
func (o *RestoreDatalakeStatusParams) WithHTTPClient(client *http.Client) *RestoreDatalakeStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore datalake status params
func (o *RestoreDatalakeStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the restore datalake status params
func (o *RestoreDatalakeStatusParams) WithInput(input *models.RestoreDatalakeStatusRequest) *RestoreDatalakeStatusParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the restore datalake status params
func (o *RestoreDatalakeStatusParams) SetInput(input *models.RestoreDatalakeStatusRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreDatalakeStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
