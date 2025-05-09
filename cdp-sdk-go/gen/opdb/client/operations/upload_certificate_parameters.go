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

// NewUploadCertificateParams creates a new UploadCertificateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadCertificateParams() *UploadCertificateParams {
	return &UploadCertificateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadCertificateParamsWithTimeout creates a new UploadCertificateParams object
// with the ability to set a timeout on a request.
func NewUploadCertificateParamsWithTimeout(timeout time.Duration) *UploadCertificateParams {
	return &UploadCertificateParams{
		timeout: timeout,
	}
}

// NewUploadCertificateParamsWithContext creates a new UploadCertificateParams object
// with the ability to set a context for a request.
func NewUploadCertificateParamsWithContext(ctx context.Context) *UploadCertificateParams {
	return &UploadCertificateParams{
		Context: ctx,
	}
}

// NewUploadCertificateParamsWithHTTPClient creates a new UploadCertificateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadCertificateParamsWithHTTPClient(client *http.Client) *UploadCertificateParams {
	return &UploadCertificateParams{
		HTTPClient: client,
	}
}

/*
UploadCertificateParams contains all the parameters to send to the API endpoint

	for the upload certificate operation.

	Typically these are written to a http.Request.
*/
type UploadCertificateParams struct {

	// Input.
	Input *models.UploadCertificateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadCertificateParams) WithDefaults() *UploadCertificateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload certificate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadCertificateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload certificate params
func (o *UploadCertificateParams) WithTimeout(timeout time.Duration) *UploadCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload certificate params
func (o *UploadCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload certificate params
func (o *UploadCertificateParams) WithContext(ctx context.Context) *UploadCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload certificate params
func (o *UploadCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload certificate params
func (o *UploadCertificateParams) WithHTTPClient(client *http.Client) *UploadCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload certificate params
func (o *UploadCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the upload certificate params
func (o *UploadCertificateParams) WithInput(input *models.UploadCertificateRequest) *UploadCertificateParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the upload certificate params
func (o *UploadCertificateParams) SetInput(input *models.UploadCertificateRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UploadCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
