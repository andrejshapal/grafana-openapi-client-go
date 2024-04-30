// Code generated by go-swagger; DO NOT EDIT.

package enterprise

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

// NewCleanDataSourceCacheParams creates a new CleanDataSourceCacheParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCleanDataSourceCacheParams() *CleanDataSourceCacheParams {
	return &CleanDataSourceCacheParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCleanDataSourceCacheParamsWithTimeout creates a new CleanDataSourceCacheParams object
// with the ability to set a timeout on a request.
func NewCleanDataSourceCacheParamsWithTimeout(timeout time.Duration) *CleanDataSourceCacheParams {
	return &CleanDataSourceCacheParams{
		timeout: timeout,
	}
}

// NewCleanDataSourceCacheParamsWithContext creates a new CleanDataSourceCacheParams object
// with the ability to set a context for a request.
func NewCleanDataSourceCacheParamsWithContext(ctx context.Context) *CleanDataSourceCacheParams {
	return &CleanDataSourceCacheParams{
		Context: ctx,
	}
}

// NewCleanDataSourceCacheParamsWithHTTPClient creates a new CleanDataSourceCacheParams object
// with the ability to set a custom HTTPClient for a request.
func NewCleanDataSourceCacheParamsWithHTTPClient(client *http.Client) *CleanDataSourceCacheParams {
	return &CleanDataSourceCacheParams{
		HTTPClient: client,
	}
}

/*
CleanDataSourceCacheParams contains all the parameters to send to the API endpoint

	for the clean data source cache operation.

	Typically these are written to a http.Request.
*/
type CleanDataSourceCacheParams struct {

	// DataSourceUID.
	DataSourceUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the clean data source cache params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CleanDataSourceCacheParams) WithDefaults() *CleanDataSourceCacheParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the clean data source cache params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CleanDataSourceCacheParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the clean data source cache params
func (o *CleanDataSourceCacheParams) WithTimeout(timeout time.Duration) *CleanDataSourceCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clean data source cache params
func (o *CleanDataSourceCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clean data source cache params
func (o *CleanDataSourceCacheParams) WithContext(ctx context.Context) *CleanDataSourceCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clean data source cache params
func (o *CleanDataSourceCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clean data source cache params
func (o *CleanDataSourceCacheParams) WithHTTPClient(client *http.Client) *CleanDataSourceCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clean data source cache params
func (o *CleanDataSourceCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataSourceUID adds the dataSourceUID to the clean data source cache params
func (o *CleanDataSourceCacheParams) WithDataSourceUID(dataSourceUID string) *CleanDataSourceCacheParams {
	o.SetDataSourceUID(dataSourceUID)
	return o
}

// SetDataSourceUID adds the dataSourceUid to the clean data source cache params
func (o *CleanDataSourceCacheParams) SetDataSourceUID(dataSourceUID string) {
	o.DataSourceUID = dataSourceUID
}

// WriteToRequest writes these params to a swagger request
func (o *CleanDataSourceCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dataSourceUID
	if err := r.SetPathParam("dataSourceUID", o.DataSourceUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
