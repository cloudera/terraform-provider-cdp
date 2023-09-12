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

// NewUpdateDataVisualizationParams creates a new UpdateDataVisualizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDataVisualizationParams() *UpdateDataVisualizationParams {
	return &UpdateDataVisualizationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDataVisualizationParamsWithTimeout creates a new UpdateDataVisualizationParams object
// with the ability to set a timeout on a request.
func NewUpdateDataVisualizationParamsWithTimeout(timeout time.Duration) *UpdateDataVisualizationParams {
	return &UpdateDataVisualizationParams{
		timeout: timeout,
	}
}

// NewUpdateDataVisualizationParamsWithContext creates a new UpdateDataVisualizationParams object
// with the ability to set a context for a request.
func NewUpdateDataVisualizationParamsWithContext(ctx context.Context) *UpdateDataVisualizationParams {
	return &UpdateDataVisualizationParams{
		Context: ctx,
	}
}

// NewUpdateDataVisualizationParamsWithHTTPClient creates a new UpdateDataVisualizationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDataVisualizationParamsWithHTTPClient(client *http.Client) *UpdateDataVisualizationParams {
	return &UpdateDataVisualizationParams{
		HTTPClient: client,
	}
}

/*
UpdateDataVisualizationParams contains all the parameters to send to the API endpoint

	for the update data visualization operation.

	Typically these are written to a http.Request.
*/
type UpdateDataVisualizationParams struct {

	// Input.
	Input *models.UpdateDataVisualizationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update data visualization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDataVisualizationParams) WithDefaults() *UpdateDataVisualizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update data visualization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDataVisualizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update data visualization params
func (o *UpdateDataVisualizationParams) WithTimeout(timeout time.Duration) *UpdateDataVisualizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update data visualization params
func (o *UpdateDataVisualizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update data visualization params
func (o *UpdateDataVisualizationParams) WithContext(ctx context.Context) *UpdateDataVisualizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update data visualization params
func (o *UpdateDataVisualizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update data visualization params
func (o *UpdateDataVisualizationParams) WithHTTPClient(client *http.Client) *UpdateDataVisualizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update data visualization params
func (o *UpdateDataVisualizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the update data visualization params
func (o *UpdateDataVisualizationParams) WithInput(input *models.UpdateDataVisualizationRequest) *UpdateDataVisualizationParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the update data visualization params
func (o *UpdateDataVisualizationParams) SetInput(input *models.UpdateDataVisualizationRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDataVisualizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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