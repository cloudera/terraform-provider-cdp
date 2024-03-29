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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// NewRepairClusterParams creates a new RepairClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRepairClusterParams() *RepairClusterParams {
	return &RepairClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRepairClusterParamsWithTimeout creates a new RepairClusterParams object
// with the ability to set a timeout on a request.
func NewRepairClusterParamsWithTimeout(timeout time.Duration) *RepairClusterParams {
	return &RepairClusterParams{
		timeout: timeout,
	}
}

// NewRepairClusterParamsWithContext creates a new RepairClusterParams object
// with the ability to set a context for a request.
func NewRepairClusterParamsWithContext(ctx context.Context) *RepairClusterParams {
	return &RepairClusterParams{
		Context: ctx,
	}
}

// NewRepairClusterParamsWithHTTPClient creates a new RepairClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewRepairClusterParamsWithHTTPClient(client *http.Client) *RepairClusterParams {
	return &RepairClusterParams{
		HTTPClient: client,
	}
}

/*
RepairClusterParams contains all the parameters to send to the API endpoint

	for the repair cluster operation.

	Typically these are written to a http.Request.
*/
type RepairClusterParams struct {

	// Input.
	Input *models.RepairClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the repair cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepairClusterParams) WithDefaults() *RepairClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the repair cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RepairClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the repair cluster params
func (o *RepairClusterParams) WithTimeout(timeout time.Duration) *RepairClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the repair cluster params
func (o *RepairClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the repair cluster params
func (o *RepairClusterParams) WithContext(ctx context.Context) *RepairClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the repair cluster params
func (o *RepairClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the repair cluster params
func (o *RepairClusterParams) WithHTTPClient(client *http.Client) *RepairClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the repair cluster params
func (o *RepairClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the repair cluster params
func (o *RepairClusterParams) WithInput(input *models.RepairClusterRequest) *RepairClusterParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the repair cluster params
func (o *RepairClusterParams) SetInput(input *models.RepairClusterRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *RepairClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
