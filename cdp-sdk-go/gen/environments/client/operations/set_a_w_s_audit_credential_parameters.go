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

// NewSetAWSAuditCredentialParams creates a new SetAWSAuditCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetAWSAuditCredentialParams() *SetAWSAuditCredentialParams {
	return &SetAWSAuditCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetAWSAuditCredentialParamsWithTimeout creates a new SetAWSAuditCredentialParams object
// with the ability to set a timeout on a request.
func NewSetAWSAuditCredentialParamsWithTimeout(timeout time.Duration) *SetAWSAuditCredentialParams {
	return &SetAWSAuditCredentialParams{
		timeout: timeout,
	}
}

// NewSetAWSAuditCredentialParamsWithContext creates a new SetAWSAuditCredentialParams object
// with the ability to set a context for a request.
func NewSetAWSAuditCredentialParamsWithContext(ctx context.Context) *SetAWSAuditCredentialParams {
	return &SetAWSAuditCredentialParams{
		Context: ctx,
	}
}

// NewSetAWSAuditCredentialParamsWithHTTPClient creates a new SetAWSAuditCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetAWSAuditCredentialParamsWithHTTPClient(client *http.Client) *SetAWSAuditCredentialParams {
	return &SetAWSAuditCredentialParams{
		HTTPClient: client,
	}
}

/*
SetAWSAuditCredentialParams contains all the parameters to send to the API endpoint

	for the set a w s audit credential operation.

	Typically these are written to a http.Request.
*/
type SetAWSAuditCredentialParams struct {

	// Input.
	Input *models.SetAWSAuditCredentialRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set a w s audit credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAWSAuditCredentialParams) WithDefaults() *SetAWSAuditCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set a w s audit credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetAWSAuditCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set a w s audit credential params
func (o *SetAWSAuditCredentialParams) WithTimeout(timeout time.Duration) *SetAWSAuditCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set a w s audit credential params
func (o *SetAWSAuditCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set a w s audit credential params
func (o *SetAWSAuditCredentialParams) WithContext(ctx context.Context) *SetAWSAuditCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set a w s audit credential params
func (o *SetAWSAuditCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set a w s audit credential params
func (o *SetAWSAuditCredentialParams) WithHTTPClient(client *http.Client) *SetAWSAuditCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set a w s audit credential params
func (o *SetAWSAuditCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the set a w s audit credential params
func (o *SetAWSAuditCredentialParams) WithInput(input *models.SetAWSAuditCredentialRequest) *SetAWSAuditCredentialParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the set a w s audit credential params
func (o *SetAWSAuditCredentialParams) SetInput(input *models.SetAWSAuditCredentialRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *SetAWSAuditCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
