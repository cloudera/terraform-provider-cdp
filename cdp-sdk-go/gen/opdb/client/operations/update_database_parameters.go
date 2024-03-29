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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

// NewUpdateDatabaseParams creates a new UpdateDatabaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDatabaseParams() *UpdateDatabaseParams {
	return &UpdateDatabaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDatabaseParamsWithTimeout creates a new UpdateDatabaseParams object
// with the ability to set a timeout on a request.
func NewUpdateDatabaseParamsWithTimeout(timeout time.Duration) *UpdateDatabaseParams {
	return &UpdateDatabaseParams{
		timeout: timeout,
	}
}

// NewUpdateDatabaseParamsWithContext creates a new UpdateDatabaseParams object
// with the ability to set a context for a request.
func NewUpdateDatabaseParamsWithContext(ctx context.Context) *UpdateDatabaseParams {
	return &UpdateDatabaseParams{
		Context: ctx,
	}
}

// NewUpdateDatabaseParamsWithHTTPClient creates a new UpdateDatabaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDatabaseParamsWithHTTPClient(client *http.Client) *UpdateDatabaseParams {
	return &UpdateDatabaseParams{
		HTTPClient: client,
	}
}

/*
UpdateDatabaseParams contains all the parameters to send to the API endpoint

	for the update database operation.

	Typically these are written to a http.Request.
*/
type UpdateDatabaseParams struct {

	// Input.
	Input *models.UpdateDatabaseRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update database params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDatabaseParams) WithDefaults() *UpdateDatabaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update database params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDatabaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update database params
func (o *UpdateDatabaseParams) WithTimeout(timeout time.Duration) *UpdateDatabaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update database params
func (o *UpdateDatabaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update database params
func (o *UpdateDatabaseParams) WithContext(ctx context.Context) *UpdateDatabaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update database params
func (o *UpdateDatabaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update database params
func (o *UpdateDatabaseParams) WithHTTPClient(client *http.Client) *UpdateDatabaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update database params
func (o *UpdateDatabaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the update database params
func (o *UpdateDatabaseParams) WithInput(input *models.UpdateDatabaseRequest) *UpdateDatabaseParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update database params
func (o *UpdateDatabaseParams) SetInput(input *models.UpdateDatabaseRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDatabaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
