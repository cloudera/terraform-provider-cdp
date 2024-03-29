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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/iam/models"
)

// NewListMachineUsersParams creates a new ListMachineUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListMachineUsersParams() *ListMachineUsersParams {
	return &ListMachineUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListMachineUsersParamsWithTimeout creates a new ListMachineUsersParams object
// with the ability to set a timeout on a request.
func NewListMachineUsersParamsWithTimeout(timeout time.Duration) *ListMachineUsersParams {
	return &ListMachineUsersParams{
		timeout: timeout,
	}
}

// NewListMachineUsersParamsWithContext creates a new ListMachineUsersParams object
// with the ability to set a context for a request.
func NewListMachineUsersParamsWithContext(ctx context.Context) *ListMachineUsersParams {
	return &ListMachineUsersParams{
		Context: ctx,
	}
}

// NewListMachineUsersParamsWithHTTPClient creates a new ListMachineUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListMachineUsersParamsWithHTTPClient(client *http.Client) *ListMachineUsersParams {
	return &ListMachineUsersParams{
		HTTPClient: client,
	}
}

/*
ListMachineUsersParams contains all the parameters to send to the API endpoint

	for the list machine users operation.

	Typically these are written to a http.Request.
*/
type ListMachineUsersParams struct {

	// Input.
	Input *models.ListMachineUsersRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list machine users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMachineUsersParams) WithDefaults() *ListMachineUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list machine users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListMachineUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list machine users params
func (o *ListMachineUsersParams) WithTimeout(timeout time.Duration) *ListMachineUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list machine users params
func (o *ListMachineUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list machine users params
func (o *ListMachineUsersParams) WithContext(ctx context.Context) *ListMachineUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list machine users params
func (o *ListMachineUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list machine users params
func (o *ListMachineUsersParams) WithHTTPClient(client *http.Client) *ListMachineUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list machine users params
func (o *ListMachineUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list machine users params
func (o *ListMachineUsersParams) WithInput(input *models.ListMachineUsersRequest) *ListMachineUsersParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list machine users params
func (o *ListMachineUsersParams) SetInput(input *models.ListMachineUsersRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListMachineUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
