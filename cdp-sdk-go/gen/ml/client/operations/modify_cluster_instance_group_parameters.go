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

// NewModifyClusterInstanceGroupParams creates a new ModifyClusterInstanceGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyClusterInstanceGroupParams() *ModifyClusterInstanceGroupParams {
	return &ModifyClusterInstanceGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyClusterInstanceGroupParamsWithTimeout creates a new ModifyClusterInstanceGroupParams object
// with the ability to set a timeout on a request.
func NewModifyClusterInstanceGroupParamsWithTimeout(timeout time.Duration) *ModifyClusterInstanceGroupParams {
	return &ModifyClusterInstanceGroupParams{
		timeout: timeout,
	}
}

// NewModifyClusterInstanceGroupParamsWithContext creates a new ModifyClusterInstanceGroupParams object
// with the ability to set a context for a request.
func NewModifyClusterInstanceGroupParamsWithContext(ctx context.Context) *ModifyClusterInstanceGroupParams {
	return &ModifyClusterInstanceGroupParams{
		Context: ctx,
	}
}

// NewModifyClusterInstanceGroupParamsWithHTTPClient creates a new ModifyClusterInstanceGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyClusterInstanceGroupParamsWithHTTPClient(client *http.Client) *ModifyClusterInstanceGroupParams {
	return &ModifyClusterInstanceGroupParams{
		HTTPClient: client,
	}
}

/*
ModifyClusterInstanceGroupParams contains all the parameters to send to the API endpoint

	for the modify cluster instance group operation.

	Typically these are written to a http.Request.
*/
type ModifyClusterInstanceGroupParams struct {

	// Input.
	Input *models.ModifyClusterInstanceGroupRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the modify cluster instance group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyClusterInstanceGroupParams) WithDefaults() *ModifyClusterInstanceGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify cluster instance group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyClusterInstanceGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify cluster instance group params
func (o *ModifyClusterInstanceGroupParams) WithTimeout(timeout time.Duration) *ModifyClusterInstanceGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify cluster instance group params
func (o *ModifyClusterInstanceGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify cluster instance group params
func (o *ModifyClusterInstanceGroupParams) WithContext(ctx context.Context) *ModifyClusterInstanceGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify cluster instance group params
func (o *ModifyClusterInstanceGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify cluster instance group params
func (o *ModifyClusterInstanceGroupParams) WithHTTPClient(client *http.Client) *ModifyClusterInstanceGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify cluster instance group params
func (o *ModifyClusterInstanceGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the modify cluster instance group params
func (o *ModifyClusterInstanceGroupParams) WithInput(input *models.ModifyClusterInstanceGroupRequest) *ModifyClusterInstanceGroupParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the modify cluster instance group params
func (o *ModifyClusterInstanceGroupParams) SetInput(input *models.ModifyClusterInstanceGroupRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyClusterInstanceGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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