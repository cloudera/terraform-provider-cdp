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

// NewCreateAWSCredentialParams creates a new CreateAWSCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAWSCredentialParams() *CreateAWSCredentialParams {
	return &CreateAWSCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAWSCredentialParamsWithTimeout creates a new CreateAWSCredentialParams object
// with the ability to set a timeout on a request.
func NewCreateAWSCredentialParamsWithTimeout(timeout time.Duration) *CreateAWSCredentialParams {
	return &CreateAWSCredentialParams{
		timeout: timeout,
	}
}

// NewCreateAWSCredentialParamsWithContext creates a new CreateAWSCredentialParams object
// with the ability to set a context for a request.
func NewCreateAWSCredentialParamsWithContext(ctx context.Context) *CreateAWSCredentialParams {
	return &CreateAWSCredentialParams{
		Context: ctx,
	}
}

// NewCreateAWSCredentialParamsWithHTTPClient creates a new CreateAWSCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAWSCredentialParamsWithHTTPClient(client *http.Client) *CreateAWSCredentialParams {
	return &CreateAWSCredentialParams{
		HTTPClient: client,
	}
}

/*
CreateAWSCredentialParams contains all the parameters to send to the API endpoint

	for the create a w s credential operation.

	Typically these are written to a http.Request.
*/
type CreateAWSCredentialParams struct {

	// Input.
	Input *models.CreateAWSCredentialRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create a w s credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAWSCredentialParams) WithDefaults() *CreateAWSCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create a w s credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAWSCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create a w s credential params
func (o *CreateAWSCredentialParams) WithTimeout(timeout time.Duration) *CreateAWSCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create a w s credential params
func (o *CreateAWSCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create a w s credential params
func (o *CreateAWSCredentialParams) WithContext(ctx context.Context) *CreateAWSCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create a w s credential params
func (o *CreateAWSCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create a w s credential params
func (o *CreateAWSCredentialParams) WithHTTPClient(client *http.Client) *CreateAWSCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create a w s credential params
func (o *CreateAWSCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the create a w s credential params
func (o *CreateAWSCredentialParams) WithInput(input *models.CreateAWSCredentialRequest) *CreateAWSCredentialParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the create a w s credential params
func (o *CreateAWSCredentialParams) SetInput(input *models.CreateAWSCredentialRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAWSCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
