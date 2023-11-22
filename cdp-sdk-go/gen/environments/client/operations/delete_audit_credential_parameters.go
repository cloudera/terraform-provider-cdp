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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/environments/models"
)

// NewDeleteAuditCredentialParams creates a new DeleteAuditCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAuditCredentialParams() *DeleteAuditCredentialParams {
	return &DeleteAuditCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAuditCredentialParamsWithTimeout creates a new DeleteAuditCredentialParams object
// with the ability to set a timeout on a request.
func NewDeleteAuditCredentialParamsWithTimeout(timeout time.Duration) *DeleteAuditCredentialParams {
	return &DeleteAuditCredentialParams{
		timeout: timeout,
	}
}

// NewDeleteAuditCredentialParamsWithContext creates a new DeleteAuditCredentialParams object
// with the ability to set a context for a request.
func NewDeleteAuditCredentialParamsWithContext(ctx context.Context) *DeleteAuditCredentialParams {
	return &DeleteAuditCredentialParams{
		Context: ctx,
	}
}

// NewDeleteAuditCredentialParamsWithHTTPClient creates a new DeleteAuditCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAuditCredentialParamsWithHTTPClient(client *http.Client) *DeleteAuditCredentialParams {
	return &DeleteAuditCredentialParams{
		HTTPClient: client,
	}
}

/*
DeleteAuditCredentialParams contains all the parameters to send to the API endpoint

	for the delete audit credential operation.

	Typically these are written to a http.Request.
*/
type DeleteAuditCredentialParams struct {

	// Input.
	Input *models.DeleteAuditCredentialRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete audit credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAuditCredentialParams) WithDefaults() *DeleteAuditCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete audit credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAuditCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete audit credential params
func (o *DeleteAuditCredentialParams) WithTimeout(timeout time.Duration) *DeleteAuditCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete audit credential params
func (o *DeleteAuditCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete audit credential params
func (o *DeleteAuditCredentialParams) WithContext(ctx context.Context) *DeleteAuditCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete audit credential params
func (o *DeleteAuditCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete audit credential params
func (o *DeleteAuditCredentialParams) WithHTTPClient(client *http.Client) *DeleteAuditCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete audit credential params
func (o *DeleteAuditCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the delete audit credential params
func (o *DeleteAuditCredentialParams) WithInput(input *models.DeleteAuditCredentialRequest) *DeleteAuditCredentialParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the delete audit credential params
func (o *DeleteAuditCredentialParams) SetInput(input *models.DeleteAuditCredentialRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAuditCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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