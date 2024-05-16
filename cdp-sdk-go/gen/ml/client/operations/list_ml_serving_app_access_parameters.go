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

// NewListMlServingAppAccessParams creates a new ListMlServingAppAccessParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListMlServingAppAccessParams() *ListMlServingAppAccessParams {
	return &ListMlServingAppAccessParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListMlServingAppAccessParamsWithTimeout creates a new ListMlServingAppAccessParams object
// with the ability to set a timeout on a request.
func NewListMlServingAppAccessParamsWithTimeout(timeout time.Duration) *ListMlServingAppAccessParams {
	return &ListMlServingAppAccessParams{
		timeout: timeout,
	}
}

// NewListMlServingAppAccessParamsWithContext creates a new ListMlServingAppAccessParams object
// with the ability to set a context for a request.
func NewListMlServingAppAccessParamsWithContext(ctx context.Context) *ListMlServingAppAccessParams {
	return &ListMlServingAppAccessParams{
		Context: ctx,
	}
}

// NewListMlServingAppAccessParamsWithHTTPClient creates a new ListMlServingAppAccessParams object
// with the ability to set a custom HTTPClient for a request.
func NewListMlServingAppAccessParamsWithHTTPClient(client *http.Client) *ListMlServingAppAccessParams {
	return &ListMlServingAppAccessParams{
		HTTPClient: client,
	}
}

/*
ListMlServingAppAccessParams contains all the parameters to send to the API endpoint

	for the list ml serving app access operation.

	Typically these are written to a http.Request.
*/
type ListMlServingAppAccessParams struct {

	// Input.
	Input *models.ListMlServingAppAccessRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list ml serving app access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMlServingAppAccessParams) WithDefaults() *ListMlServingAppAccessParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list ml serving app access params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMlServingAppAccessParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list ml serving app access params
func (o *ListMlServingAppAccessParams) WithTimeout(timeout time.Duration) *ListMlServingAppAccessParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list ml serving app access params
func (o *ListMlServingAppAccessParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list ml serving app access params
func (o *ListMlServingAppAccessParams) WithContext(ctx context.Context) *ListMlServingAppAccessParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list ml serving app access params
func (o *ListMlServingAppAccessParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list ml serving app access params
func (o *ListMlServingAppAccessParams) WithHTTPClient(client *http.Client) *ListMlServingAppAccessParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list ml serving app access params
func (o *ListMlServingAppAccessParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list ml serving app access params
func (o *ListMlServingAppAccessParams) WithInput(input *models.ListMlServingAppAccessRequest) *ListMlServingAppAccessParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list ml serving app access params
func (o *ListMlServingAppAccessParams) SetInput(input *models.ListMlServingAppAccessRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListMlServingAppAccessParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
