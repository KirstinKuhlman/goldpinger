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
)

// NewHealthzParams creates a new HealthzParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHealthzParams() *HealthzParams {
	return &HealthzParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHealthzParamsWithTimeout creates a new HealthzParams object
// with the ability to set a timeout on a request.
func NewHealthzParamsWithTimeout(timeout time.Duration) *HealthzParams {
	return &HealthzParams{
		timeout: timeout,
	}
}

// NewHealthzParamsWithContext creates a new HealthzParams object
// with the ability to set a context for a request.
func NewHealthzParamsWithContext(ctx context.Context) *HealthzParams {
	return &HealthzParams{
		Context: ctx,
	}
}

// NewHealthzParamsWithHTTPClient creates a new HealthzParams object
// with the ability to set a custom HTTPClient for a request.
func NewHealthzParamsWithHTTPClient(client *http.Client) *HealthzParams {
	return &HealthzParams{
		HTTPClient: client,
	}
}

/* HealthzParams contains all the parameters to send to the API endpoint
   for the healthz operation.

   Typically these are written to a http.Request.
*/
type HealthzParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the healthz params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HealthzParams) WithDefaults() *HealthzParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the healthz params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HealthzParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the healthz params
func (o *HealthzParams) WithTimeout(timeout time.Duration) *HealthzParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the healthz params
func (o *HealthzParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the healthz params
func (o *HealthzParams) WithContext(ctx context.Context) *HealthzParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the healthz params
func (o *HealthzParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the healthz params
func (o *HealthzParams) WithHTTPClient(client *http.Client) *HealthzParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the healthz params
func (o *HealthzParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HealthzParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
