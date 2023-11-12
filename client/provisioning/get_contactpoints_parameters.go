// Code generated by go-swagger; DO NOT EDIT.

package provisioning

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

// NewGetContactpointsParams creates a new GetContactpointsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetContactpointsParams() *GetContactpointsParams {
	return &GetContactpointsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetContactpointsParamsWithTimeout creates a new GetContactpointsParams object
// with the ability to set a timeout on a request.
func NewGetContactpointsParamsWithTimeout(timeout time.Duration) *GetContactpointsParams {
	return &GetContactpointsParams{
		timeout: timeout,
	}
}

// NewGetContactpointsParamsWithContext creates a new GetContactpointsParams object
// with the ability to set a context for a request.
func NewGetContactpointsParamsWithContext(ctx context.Context) *GetContactpointsParams {
	return &GetContactpointsParams{
		Context: ctx,
	}
}

// NewGetContactpointsParamsWithHTTPClient creates a new GetContactpointsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetContactpointsParamsWithHTTPClient(client *http.Client) *GetContactpointsParams {
	return &GetContactpointsParams{
		HTTPClient: client,
	}
}

/*
GetContactpointsParams contains all the parameters to send to the API endpoint

	for the get contactpoints operation.

	Typically these are written to a http.Request.
*/
type GetContactpointsParams struct {

	/* Name.

	   Filter by name
	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get contactpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactpointsParams) WithDefaults() *GetContactpointsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get contactpoints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetContactpointsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get contactpoints params
func (o *GetContactpointsParams) WithTimeout(timeout time.Duration) *GetContactpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get contactpoints params
func (o *GetContactpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get contactpoints params
func (o *GetContactpointsParams) WithContext(ctx context.Context) *GetContactpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get contactpoints params
func (o *GetContactpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get contactpoints params
func (o *GetContactpointsParams) WithHTTPClient(client *http.Client) *GetContactpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get contactpoints params
func (o *GetContactpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get contactpoints params
func (o *GetContactpointsParams) WithName(name *string) *GetContactpointsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get contactpoints params
func (o *GetContactpointsParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetContactpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}