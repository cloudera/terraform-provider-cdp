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

// NewListBackupsParams creates a new ListBackupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListBackupsParams() *ListBackupsParams {
	return &ListBackupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListBackupsParamsWithTimeout creates a new ListBackupsParams object
// with the ability to set a timeout on a request.
func NewListBackupsParamsWithTimeout(timeout time.Duration) *ListBackupsParams {
	return &ListBackupsParams{
		timeout: timeout,
	}
}

// NewListBackupsParamsWithContext creates a new ListBackupsParams object
// with the ability to set a context for a request.
func NewListBackupsParamsWithContext(ctx context.Context) *ListBackupsParams {
	return &ListBackupsParams{
		Context: ctx,
	}
}

// NewListBackupsParamsWithHTTPClient creates a new ListBackupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListBackupsParamsWithHTTPClient(client *http.Client) *ListBackupsParams {
	return &ListBackupsParams{
		HTTPClient: client,
	}
}

/*
ListBackupsParams contains all the parameters to send to the API endpoint

	for the list backups operation.

	Typically these are written to a http.Request.
*/
type ListBackupsParams struct {

	// Input.
	Input *models.ListBackupsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list backups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBackupsParams) WithDefaults() *ListBackupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list backups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBackupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list backups params
func (o *ListBackupsParams) WithTimeout(timeout time.Duration) *ListBackupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list backups params
func (o *ListBackupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list backups params
func (o *ListBackupsParams) WithContext(ctx context.Context) *ListBackupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list backups params
func (o *ListBackupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list backups params
func (o *ListBackupsParams) WithHTTPClient(client *http.Client) *ListBackupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list backups params
func (o *ListBackupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list backups params
func (o *ListBackupsParams) WithInput(input *models.ListBackupsRequest) *ListBackupsParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list backups params
func (o *ListBackupsParams) SetInput(input *models.ListBackupsRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListBackupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
