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

// NewBackupDatalakeStatusParams creates a new BackupDatalakeStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupDatalakeStatusParams() *BackupDatalakeStatusParams {
	return &BackupDatalakeStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupDatalakeStatusParamsWithTimeout creates a new BackupDatalakeStatusParams object
// with the ability to set a timeout on a request.
func NewBackupDatalakeStatusParamsWithTimeout(timeout time.Duration) *BackupDatalakeStatusParams {
	return &BackupDatalakeStatusParams{
		timeout: timeout,
	}
}

// NewBackupDatalakeStatusParamsWithContext creates a new BackupDatalakeStatusParams object
// with the ability to set a context for a request.
func NewBackupDatalakeStatusParamsWithContext(ctx context.Context) *BackupDatalakeStatusParams {
	return &BackupDatalakeStatusParams{
		Context: ctx,
	}
}

// NewBackupDatalakeStatusParamsWithHTTPClient creates a new BackupDatalakeStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupDatalakeStatusParamsWithHTTPClient(client *http.Client) *BackupDatalakeStatusParams {
	return &BackupDatalakeStatusParams{
		HTTPClient: client,
	}
}

/*
BackupDatalakeStatusParams contains all the parameters to send to the API endpoint

	for the backup datalake status operation.

	Typically these are written to a http.Request.
*/
type BackupDatalakeStatusParams struct {

	// Input.
	Input *models.BackupDatalakeStatusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backup datalake status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupDatalakeStatusParams) WithDefaults() *BackupDatalakeStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backup datalake status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupDatalakeStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backup datalake status params
func (o *BackupDatalakeStatusParams) WithTimeout(timeout time.Duration) *BackupDatalakeStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup datalake status params
func (o *BackupDatalakeStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup datalake status params
func (o *BackupDatalakeStatusParams) WithContext(ctx context.Context) *BackupDatalakeStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup datalake status params
func (o *BackupDatalakeStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup datalake status params
func (o *BackupDatalakeStatusParams) WithHTTPClient(client *http.Client) *BackupDatalakeStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup datalake status params
func (o *BackupDatalakeStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the backup datalake status params
func (o *BackupDatalakeStatusParams) WithInput(input *models.BackupDatalakeStatusRequest) *BackupDatalakeStatusParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the backup datalake status params
func (o *BackupDatalakeStatusParams) SetInput(input *models.BackupDatalakeStatusRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *BackupDatalakeStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
