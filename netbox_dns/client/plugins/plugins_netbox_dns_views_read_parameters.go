// Code generated by go-swagger; DO NOT EDIT.

package plugins

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
	"github.com/go-openapi/swag"
)

// NewPluginsNetboxDNSViewsReadParams creates a new PluginsNetboxDNSViewsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsNetboxDNSViewsReadParams() *PluginsNetboxDNSViewsReadParams {
	return &PluginsNetboxDNSViewsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsNetboxDNSViewsReadParamsWithTimeout creates a new PluginsNetboxDNSViewsReadParams object
// with the ability to set a timeout on a request.
func NewPluginsNetboxDNSViewsReadParamsWithTimeout(timeout time.Duration) *PluginsNetboxDNSViewsReadParams {
	return &PluginsNetboxDNSViewsReadParams{
		timeout: timeout,
	}
}

// NewPluginsNetboxDNSViewsReadParamsWithContext creates a new PluginsNetboxDNSViewsReadParams object
// with the ability to set a context for a request.
func NewPluginsNetboxDNSViewsReadParamsWithContext(ctx context.Context) *PluginsNetboxDNSViewsReadParams {
	return &PluginsNetboxDNSViewsReadParams{
		Context: ctx,
	}
}

// NewPluginsNetboxDNSViewsReadParamsWithHTTPClient creates a new PluginsNetboxDNSViewsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsNetboxDNSViewsReadParamsWithHTTPClient(client *http.Client) *PluginsNetboxDNSViewsReadParams {
	return &PluginsNetboxDNSViewsReadParams{
		HTTPClient: client,
	}
}

/*
PluginsNetboxDNSViewsReadParams contains all the parameters to send to the API endpoint

	for the plugins netbox dns views read operation.

	Typically these are written to a http.Request.
*/
type PluginsNetboxDNSViewsReadParams struct {

	/* ID.

	   A unique integer value identifying this view.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins netbox dns views read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSViewsReadParams) WithDefaults() *PluginsNetboxDNSViewsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins netbox dns views read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsNetboxDNSViewsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins netbox dns views read params
func (o *PluginsNetboxDNSViewsReadParams) WithTimeout(timeout time.Duration) *PluginsNetboxDNSViewsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins netbox dns views read params
func (o *PluginsNetboxDNSViewsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins netbox dns views read params
func (o *PluginsNetboxDNSViewsReadParams) WithContext(ctx context.Context) *PluginsNetboxDNSViewsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins netbox dns views read params
func (o *PluginsNetboxDNSViewsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins netbox dns views read params
func (o *PluginsNetboxDNSViewsReadParams) WithHTTPClient(client *http.Client) *PluginsNetboxDNSViewsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins netbox dns views read params
func (o *PluginsNetboxDNSViewsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins netbox dns views read params
func (o *PluginsNetboxDNSViewsReadParams) WithID(id int64) *PluginsNetboxDNSViewsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins netbox dns views read params
func (o *PluginsNetboxDNSViewsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsNetboxDNSViewsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}