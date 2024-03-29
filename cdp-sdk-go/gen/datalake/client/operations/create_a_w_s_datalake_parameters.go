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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// NewCreateAWSDatalakeParams creates a new CreateAWSDatalakeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAWSDatalakeParams() *CreateAWSDatalakeParams {
	return &CreateAWSDatalakeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAWSDatalakeParamsWithTimeout creates a new CreateAWSDatalakeParams object
// with the ability to set a timeout on a request.
func NewCreateAWSDatalakeParamsWithTimeout(timeout time.Duration) *CreateAWSDatalakeParams {
	return &CreateAWSDatalakeParams{
		timeout: timeout,
	}
}

// NewCreateAWSDatalakeParamsWithContext creates a new CreateAWSDatalakeParams object
// with the ability to set a context for a request.
func NewCreateAWSDatalakeParamsWithContext(ctx context.Context) *CreateAWSDatalakeParams {
	return &CreateAWSDatalakeParams{
		Context: ctx,
	}
}

// NewCreateAWSDatalakeParamsWithHTTPClient creates a new CreateAWSDatalakeParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAWSDatalakeParamsWithHTTPClient(client *http.Client) *CreateAWSDatalakeParams {
	return &CreateAWSDatalakeParams{
		HTTPClient: client,
	}
}

/*
CreateAWSDatalakeParams contains all the parameters to send to the API endpoint

	for the create a w s datalake operation.

	Typically these are written to a http.Request.
*/
type CreateAWSDatalakeParams struct {

	// Input.
	Input *models.CreateAWSDatalakeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create a w s datalake params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAWSDatalakeParams) WithDefaults() *CreateAWSDatalakeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create a w s datalake params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAWSDatalakeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create a w s datalake params
func (o *CreateAWSDatalakeParams) WithTimeout(timeout time.Duration) *CreateAWSDatalakeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create a w s datalake params
func (o *CreateAWSDatalakeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create a w s datalake params
func (o *CreateAWSDatalakeParams) WithContext(ctx context.Context) *CreateAWSDatalakeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create a w s datalake params
func (o *CreateAWSDatalakeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create a w s datalake params
func (o *CreateAWSDatalakeParams) WithHTTPClient(client *http.Client) *CreateAWSDatalakeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create a w s datalake params
func (o *CreateAWSDatalakeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the create a w s datalake params
func (o *CreateAWSDatalakeParams) WithInput(input *models.CreateAWSDatalakeRequest) *CreateAWSDatalakeParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the create a w s datalake params
func (o *CreateAWSDatalakeParams) SetInput(input *models.CreateAWSDatalakeRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAWSDatalakeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
