// Code generated by go-swagger; DO NOT EDIT.

package datasources

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

// NewGetDataSourceByNameParams creates a new GetDataSourceByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDataSourceByNameParams() *GetDataSourceByNameParams {
	return &GetDataSourceByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDataSourceByNameParamsWithTimeout creates a new GetDataSourceByNameParams object
// with the ability to set a timeout on a request.
func NewGetDataSourceByNameParamsWithTimeout(timeout time.Duration) *GetDataSourceByNameParams {
	return &GetDataSourceByNameParams{
		timeout: timeout,
	}
}

// NewGetDataSourceByNameParamsWithContext creates a new GetDataSourceByNameParams object
// with the ability to set a context for a request.
func NewGetDataSourceByNameParamsWithContext(ctx context.Context) *GetDataSourceByNameParams {
	return &GetDataSourceByNameParams{
		Context: ctx,
	}
}

// NewGetDataSourceByNameParamsWithHTTPClient creates a new GetDataSourceByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDataSourceByNameParamsWithHTTPClient(client *http.Client) *GetDataSourceByNameParams {
	return &GetDataSourceByNameParams{
		HTTPClient: client,
	}
}

/*
GetDataSourceByNameParams contains all the parameters to send to the API endpoint

	for the get data source by name operation.

	Typically these are written to a http.Request.
*/
type GetDataSourceByNameParams struct {

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get data source by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataSourceByNameParams) WithDefaults() *GetDataSourceByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get data source by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDataSourceByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get data source by name params
func (o *GetDataSourceByNameParams) WithTimeout(timeout time.Duration) *GetDataSourceByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get data source by name params
func (o *GetDataSourceByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get data source by name params
func (o *GetDataSourceByNameParams) WithContext(ctx context.Context) *GetDataSourceByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get data source by name params
func (o *GetDataSourceByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get data source by name params
func (o *GetDataSourceByNameParams) WithHTTPClient(client *http.Client) *GetDataSourceByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get data source by name params
func (o *GetDataSourceByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get data source by name params
func (o *GetDataSourceByNameParams) WithName(name string) *GetDataSourceByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get data source by name params
func (o *GetDataSourceByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetDataSourceByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}