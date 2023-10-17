// Code generated by go-swagger; DO NOT EDIT.

package admin_provisioning

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

// NewAdminProvisioningReloadPluginsParams creates a new AdminProvisioningReloadPluginsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminProvisioningReloadPluginsParams() *AdminProvisioningReloadPluginsParams {
	return &AdminProvisioningReloadPluginsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminProvisioningReloadPluginsParamsWithTimeout creates a new AdminProvisioningReloadPluginsParams object
// with the ability to set a timeout on a request.
func NewAdminProvisioningReloadPluginsParamsWithTimeout(timeout time.Duration) *AdminProvisioningReloadPluginsParams {
	return &AdminProvisioningReloadPluginsParams{
		timeout: timeout,
	}
}

// NewAdminProvisioningReloadPluginsParamsWithContext creates a new AdminProvisioningReloadPluginsParams object
// with the ability to set a context for a request.
func NewAdminProvisioningReloadPluginsParamsWithContext(ctx context.Context) *AdminProvisioningReloadPluginsParams {
	return &AdminProvisioningReloadPluginsParams{
		Context: ctx,
	}
}

// NewAdminProvisioningReloadPluginsParamsWithHTTPClient creates a new AdminProvisioningReloadPluginsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminProvisioningReloadPluginsParamsWithHTTPClient(client *http.Client) *AdminProvisioningReloadPluginsParams {
	return &AdminProvisioningReloadPluginsParams{
		HTTPClient: client,
	}
}

/*
AdminProvisioningReloadPluginsParams contains all the parameters to send to the API endpoint

	for the admin provisioning reload plugins operation.

	Typically these are written to a http.Request.
*/
type AdminProvisioningReloadPluginsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin provisioning reload plugins params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminProvisioningReloadPluginsParams) WithDefaults() *AdminProvisioningReloadPluginsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin provisioning reload plugins params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminProvisioningReloadPluginsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin provisioning reload plugins params
func (o *AdminProvisioningReloadPluginsParams) WithTimeout(timeout time.Duration) *AdminProvisioningReloadPluginsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin provisioning reload plugins params
func (o *AdminProvisioningReloadPluginsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin provisioning reload plugins params
func (o *AdminProvisioningReloadPluginsParams) WithContext(ctx context.Context) *AdminProvisioningReloadPluginsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin provisioning reload plugins params
func (o *AdminProvisioningReloadPluginsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin provisioning reload plugins params
func (o *AdminProvisioningReloadPluginsParams) WithHTTPClient(client *http.Client) *AdminProvisioningReloadPluginsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin provisioning reload plugins params
func (o *AdminProvisioningReloadPluginsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AdminProvisioningReloadPluginsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}