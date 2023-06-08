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

// NewUpdateUserParams creates a new UpdateUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUserParams() *UpdateUserParams {
	return &UpdateUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserParamsWithTimeout creates a new UpdateUserParams object
// with the ability to set a timeout on a request.
func NewUpdateUserParamsWithTimeout(timeout time.Duration) *UpdateUserParams {
	return &UpdateUserParams{
		timeout: timeout,
	}
}

// NewUpdateUserParamsWithContext creates a new UpdateUserParams object
// with the ability to set a context for a request.
func NewUpdateUserParamsWithContext(ctx context.Context) *UpdateUserParams {
	return &UpdateUserParams{
		Context: ctx,
	}
}

// NewUpdateUserParamsWithHTTPClient creates a new UpdateUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUserParamsWithHTTPClient(client *http.Client) *UpdateUserParams {
	return &UpdateUserParams{
		HTTPClient: client,
	}
}

/*
UpdateUserParams contains all the parameters to send to the API endpoint

	for the update user operation.

	Typically these are written to a http.Request.
*/
type UpdateUserParams struct {

	// Input.
	Input *models.UpdateUserRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserParams) WithDefaults() *UpdateUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update user params
func (o *UpdateUserParams) WithTimeout(timeout time.Duration) *UpdateUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user params
func (o *UpdateUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user params
func (o *UpdateUserParams) WithContext(ctx context.Context) *UpdateUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user params
func (o *UpdateUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user params
func (o *UpdateUserParams) WithHTTPClient(client *http.Client) *UpdateUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user params
func (o *UpdateUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the update user params
func (o *UpdateUserParams) WithInput(input *models.UpdateUserRequest) *UpdateUserParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update user params
func (o *UpdateUserParams) SetInput(input *models.UpdateUserRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
