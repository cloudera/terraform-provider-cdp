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

// NewDeleteDataVisualizationParams creates a new DeleteDataVisualizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDataVisualizationParams() *DeleteDataVisualizationParams {
	return &DeleteDataVisualizationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDataVisualizationParamsWithTimeout creates a new DeleteDataVisualizationParams object
// with the ability to set a timeout on a request.
func NewDeleteDataVisualizationParamsWithTimeout(timeout time.Duration) *DeleteDataVisualizationParams {
	return &DeleteDataVisualizationParams{
		timeout: timeout,
	}
}

// NewDeleteDataVisualizationParamsWithContext creates a new DeleteDataVisualizationParams object
// with the ability to set a context for a request.
func NewDeleteDataVisualizationParamsWithContext(ctx context.Context) *DeleteDataVisualizationParams {
	return &DeleteDataVisualizationParams{
		Context: ctx,
	}
}

// NewDeleteDataVisualizationParamsWithHTTPClient creates a new DeleteDataVisualizationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDataVisualizationParamsWithHTTPClient(client *http.Client) *DeleteDataVisualizationParams {
	return &DeleteDataVisualizationParams{
		HTTPClient: client,
	}
}

/*
DeleteDataVisualizationParams contains all the parameters to send to the API endpoint

	for the delete data visualization operation.

	Typically these are written to a http.Request.
*/
type DeleteDataVisualizationParams struct {

	// Input.
	Input *models.DeleteDataVisualizationRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete data visualization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDataVisualizationParams) WithDefaults() *DeleteDataVisualizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete data visualization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDataVisualizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete data visualization params
func (o *DeleteDataVisualizationParams) WithTimeout(timeout time.Duration) *DeleteDataVisualizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete data visualization params
func (o *DeleteDataVisualizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete data visualization params
func (o *DeleteDataVisualizationParams) WithContext(ctx context.Context) *DeleteDataVisualizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete data visualization params
func (o *DeleteDataVisualizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete data visualization params
func (o *DeleteDataVisualizationParams) WithHTTPClient(client *http.Client) *DeleteDataVisualizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete data visualization params
func (o *DeleteDataVisualizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the delete data visualization params
func (o *DeleteDataVisualizationParams) WithInput(input *models.DeleteDataVisualizationRequest) *DeleteDataVisualizationParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the delete data visualization params
func (o *DeleteDataVisualizationParams) SetInput(input *models.DeleteDataVisualizationRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDataVisualizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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