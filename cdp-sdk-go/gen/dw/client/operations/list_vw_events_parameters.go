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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/dw/models"
)

// NewListVwEventsParams creates a new ListVwEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListVwEventsParams() *ListVwEventsParams {
	return &ListVwEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListVwEventsParamsWithTimeout creates a new ListVwEventsParams object
// with the ability to set a timeout on a request.
func NewListVwEventsParamsWithTimeout(timeout time.Duration) *ListVwEventsParams {
	return &ListVwEventsParams{
		timeout: timeout,
	}
}

// NewListVwEventsParamsWithContext creates a new ListVwEventsParams object
// with the ability to set a context for a request.
func NewListVwEventsParamsWithContext(ctx context.Context) *ListVwEventsParams {
	return &ListVwEventsParams{
		Context: ctx,
	}
}

// NewListVwEventsParamsWithHTTPClient creates a new ListVwEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListVwEventsParamsWithHTTPClient(client *http.Client) *ListVwEventsParams {
	return &ListVwEventsParams{
		HTTPClient: client,
	}
}

/*
ListVwEventsParams contains all the parameters to send to the API endpoint

	for the list vw events operation.

	Typically these are written to a http.Request.
*/
type ListVwEventsParams struct {

	// Input.
	Input *models.ListVwEventsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list vw events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVwEventsParams) WithDefaults() *ListVwEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list vw events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVwEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list vw events params
func (o *ListVwEventsParams) WithTimeout(timeout time.Duration) *ListVwEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list vw events params
func (o *ListVwEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list vw events params
func (o *ListVwEventsParams) WithContext(ctx context.Context) *ListVwEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list vw events params
func (o *ListVwEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list vw events params
func (o *ListVwEventsParams) WithHTTPClient(client *http.Client) *ListVwEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list vw events params
func (o *ListVwEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list vw events params
func (o *ListVwEventsParams) WithInput(input *models.ListVwEventsRequest) *ListVwEventsParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list vw events params
func (o *ListVwEventsParams) SetInput(input *models.ListVwEventsRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListVwEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
