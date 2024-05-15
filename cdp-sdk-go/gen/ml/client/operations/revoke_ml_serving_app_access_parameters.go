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

// NewRevokeMlServingAppAccessParams creates a new RevokeMlServingAppAccessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevokeMlServingAppAccessParams() *RevokeMlServingAppAccessParams {
	return &RevokeMlServingAppAccessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeMlServingAppAccessParamsWithTimeout creates a new RevokeMlServingAppAccessParams object
// with the ability to set a timeout on a request.
func NewRevokeMlServingAppAccessParamsWithTimeout(timeout time.Duration) *RevokeMlServingAppAccessParams {
	return &RevokeMlServingAppAccessParams{
		timeout: timeout,
	}
}

// NewRevokeMlServingAppAccessParamsWithContext creates a new RevokeMlServingAppAccessParams object
// with the ability to set a context for a request.
func NewRevokeMlServingAppAccessParamsWithContext(ctx context.Context) *RevokeMlServingAppAccessParams {
	return &RevokeMlServingAppAccessParams{
		Context: ctx,
	}
}

// NewRevokeMlServingAppAccessParamsWithHTTPClient creates a new RevokeMlServingAppAccessParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevokeMlServingAppAccessParamsWithHTTPClient(client *http.Client) *RevokeMlServingAppAccessParams {
	return &RevokeMlServingAppAccessParams{
		HTTPClient: client,
	}
}

/*
RevokeMlServingAppAccessParams contains all the parameters to send to the API endpoint

	for the revoke ml serving app access operation.

	Typically these are written to a http.Request.
*/
type RevokeMlServingAppAccessParams struct {

	// Input.
	Input *models.RevokeMlServingAppAccessRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the revoke ml serving app access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeMlServingAppAccessParams) WithDefaults() *RevokeMlServingAppAccessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revoke ml serving app access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevokeMlServingAppAccessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revoke ml serving app access params
func (o *RevokeMlServingAppAccessParams) WithTimeout(timeout time.Duration) *RevokeMlServingAppAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke ml serving app access params
func (o *RevokeMlServingAppAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke ml serving app access params
func (o *RevokeMlServingAppAccessParams) WithContext(ctx context.Context) *RevokeMlServingAppAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke ml serving app access params
func (o *RevokeMlServingAppAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke ml serving app access params
func (o *RevokeMlServingAppAccessParams) WithHTTPClient(client *http.Client) *RevokeMlServingAppAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke ml serving app access params
func (o *RevokeMlServingAppAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the revoke ml serving app access params
func (o *RevokeMlServingAppAccessParams) WithInput(input *models.RevokeMlServingAppAccessRequest) *RevokeMlServingAppAccessParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the revoke ml serving app access params
func (o *RevokeMlServingAppAccessParams) SetInput(input *models.RevokeMlServingAppAccessRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeMlServingAppAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
