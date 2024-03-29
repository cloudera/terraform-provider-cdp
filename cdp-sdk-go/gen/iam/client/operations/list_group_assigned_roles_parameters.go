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

// NewListGroupAssignedRolesParams creates a new ListGroupAssignedRolesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListGroupAssignedRolesParams() *ListGroupAssignedRolesParams {
	return &ListGroupAssignedRolesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListGroupAssignedRolesParamsWithTimeout creates a new ListGroupAssignedRolesParams object
// with the ability to set a timeout on a request.
func NewListGroupAssignedRolesParamsWithTimeout(timeout time.Duration) *ListGroupAssignedRolesParams {
	return &ListGroupAssignedRolesParams{
		timeout: timeout,
	}
}

// NewListGroupAssignedRolesParamsWithContext creates a new ListGroupAssignedRolesParams object
// with the ability to set a context for a request.
func NewListGroupAssignedRolesParamsWithContext(ctx context.Context) *ListGroupAssignedRolesParams {
	return &ListGroupAssignedRolesParams{
		Context: ctx,
	}
}

// NewListGroupAssignedRolesParamsWithHTTPClient creates a new ListGroupAssignedRolesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListGroupAssignedRolesParamsWithHTTPClient(client *http.Client) *ListGroupAssignedRolesParams {
	return &ListGroupAssignedRolesParams{
		HTTPClient: client,
	}
}

/*
ListGroupAssignedRolesParams contains all the parameters to send to the API endpoint

	for the list group assigned roles operation.

	Typically these are written to a http.Request.
*/
type ListGroupAssignedRolesParams struct {

	// Input.
	Input *models.ListGroupAssignedRolesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list group assigned roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGroupAssignedRolesParams) WithDefaults() *ListGroupAssignedRolesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list group assigned roles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGroupAssignedRolesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list group assigned roles params
func (o *ListGroupAssignedRolesParams) WithTimeout(timeout time.Duration) *ListGroupAssignedRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list group assigned roles params
func (o *ListGroupAssignedRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list group assigned roles params
func (o *ListGroupAssignedRolesParams) WithContext(ctx context.Context) *ListGroupAssignedRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list group assigned roles params
func (o *ListGroupAssignedRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list group assigned roles params
func (o *ListGroupAssignedRolesParams) WithHTTPClient(client *http.Client) *ListGroupAssignedRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list group assigned roles params
func (o *ListGroupAssignedRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the list group assigned roles params
func (o *ListGroupAssignedRolesParams) WithInput(input *models.ListGroupAssignedRolesRequest) *ListGroupAssignedRolesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the list group assigned roles params
func (o *ListGroupAssignedRolesParams) SetInput(input *models.ListGroupAssignedRolesRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ListGroupAssignedRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
