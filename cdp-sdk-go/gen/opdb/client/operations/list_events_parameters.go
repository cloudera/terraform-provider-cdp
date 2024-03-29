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

// NewListEventsParams creates a new ListEventsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListEventsParams() *ListEventsParams {
	return &ListEventsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListEventsParamsWithTimeout creates a new ListEventsParams object
// with the ability to set a timeout on a request.
func NewListEventsParamsWithTimeout(timeout time.Duration) *ListEventsParams {
	return &ListEventsParams{
		timeout: timeout,
	}
}

// NewListEventsParamsWithContext creates a new ListEventsParams object
// with the ability to set a context for a request.
func NewListEventsParamsWithContext(ctx context.Context) *ListEventsParams {
	return &ListEventsParams{
		Context: ctx,
	}
}

// NewListEventsParamsWithHTTPClient creates a new ListEventsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListEventsParamsWithHTTPClient(client *http.Client) *ListEventsParams {
	return &ListEventsParams{
		HTTPClient: client,
	}
}

/*
ListEventsParams contains all the parameters to send to the API endpoint

	for the list events operation.

	Typically these are written to a http.Request.
*/
type ListEventsParams struct {

	// Input.
	Input *models.ListEventsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEventsParams) WithDefaults() *ListEventsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list events params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListEventsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list events params
func (o *ListEventsParams) WithTimeout(timeout time.Duration) *ListEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list events params
func (o *ListEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list events params
func (o *ListEventsParams) WithContext(ctx context.Context) *ListEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list events params
func (o *ListEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list events params
func (o *ListEventsParams) WithHTTPClient(client *http.Client) *ListEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list events params
func (o *ListEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list events params
func (o *ListEventsParams) WithInput(input *models.ListEventsRequest) *ListEventsParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list events params
func (o *ListEventsParams) SetInput(input *models.ListEventsRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
