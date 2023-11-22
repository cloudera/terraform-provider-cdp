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

// NewUnassignMachineUserResourceRoleParams creates a new UnassignMachineUserResourceRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnassignMachineUserResourceRoleParams() *UnassignMachineUserResourceRoleParams {
	return &UnassignMachineUserResourceRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnassignMachineUserResourceRoleParamsWithTimeout creates a new UnassignMachineUserResourceRoleParams object
// with the ability to set a timeout on a request.
func NewUnassignMachineUserResourceRoleParamsWithTimeout(timeout time.Duration) *UnassignMachineUserResourceRoleParams {
	return &UnassignMachineUserResourceRoleParams{
		timeout: timeout,
	}
}

// NewUnassignMachineUserResourceRoleParamsWithContext creates a new UnassignMachineUserResourceRoleParams object
// with the ability to set a context for a request.
func NewUnassignMachineUserResourceRoleParamsWithContext(ctx context.Context) *UnassignMachineUserResourceRoleParams {
	return &UnassignMachineUserResourceRoleParams{
		Context: ctx,
	}
}

// NewUnassignMachineUserResourceRoleParamsWithHTTPClient creates a new UnassignMachineUserResourceRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnassignMachineUserResourceRoleParamsWithHTTPClient(client *http.Client) *UnassignMachineUserResourceRoleParams {
	return &UnassignMachineUserResourceRoleParams{
		HTTPClient: client,
	}
}

/*
UnassignMachineUserResourceRoleParams contains all the parameters to send to the API endpoint

	for the unassign machine user resource role operation.

	Typically these are written to a http.Request.
*/
type UnassignMachineUserResourceRoleParams struct {

	// Input.
	Input *models.UnassignMachineUserResourceRoleRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unassign machine user resource role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnassignMachineUserResourceRoleParams) WithDefaults() *UnassignMachineUserResourceRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unassign machine user resource role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnassignMachineUserResourceRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unassign machine user resource role params
func (o *UnassignMachineUserResourceRoleParams) WithTimeout(timeout time.Duration) *UnassignMachineUserResourceRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unassign machine user resource role params
func (o *UnassignMachineUserResourceRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unassign machine user resource role params
func (o *UnassignMachineUserResourceRoleParams) WithContext(ctx context.Context) *UnassignMachineUserResourceRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unassign machine user resource role params
func (o *UnassignMachineUserResourceRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unassign machine user resource role params
func (o *UnassignMachineUserResourceRoleParams) WithHTTPClient(client *http.Client) *UnassignMachineUserResourceRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unassign machine user resource role params
func (o *UnassignMachineUserResourceRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the unassign machine user resource role params
func (o *UnassignMachineUserResourceRoleParams) WithInput(input *models.UnassignMachineUserResourceRoleRequest) *UnassignMachineUserResourceRoleParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the unassign machine user resource role params
func (o *UnassignMachineUserResourceRoleParams) SetInput(input *models.UnassignMachineUserResourceRoleRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UnassignMachineUserResourceRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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