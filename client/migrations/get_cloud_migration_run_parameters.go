// Code generated by go-swagger; DO NOT EDIT.

package migrations

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

// NewGetCloudMigrationRunParams creates a new GetCloudMigrationRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCloudMigrationRunParams() *GetCloudMigrationRunParams {
	return &GetCloudMigrationRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudMigrationRunParamsWithTimeout creates a new GetCloudMigrationRunParams object
// with the ability to set a timeout on a request.
func NewGetCloudMigrationRunParamsWithTimeout(timeout time.Duration) *GetCloudMigrationRunParams {
	return &GetCloudMigrationRunParams{
		timeout: timeout,
	}
}

// NewGetCloudMigrationRunParamsWithContext creates a new GetCloudMigrationRunParams object
// with the ability to set a context for a request.
func NewGetCloudMigrationRunParamsWithContext(ctx context.Context) *GetCloudMigrationRunParams {
	return &GetCloudMigrationRunParams{
		Context: ctx,
	}
}

// NewGetCloudMigrationRunParamsWithHTTPClient creates a new GetCloudMigrationRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCloudMigrationRunParamsWithHTTPClient(client *http.Client) *GetCloudMigrationRunParams {
	return &GetCloudMigrationRunParams{
		HTTPClient: client,
	}
}

/*
GetCloudMigrationRunParams contains all the parameters to send to the API endpoint

	for the get cloud migration run operation.

	Typically these are written to a http.Request.
*/
type GetCloudMigrationRunParams struct {

	/* ID.

	   ID of an migration

	   Format: int64
	*/
	ID int64

	/* RunID.

	   Run ID of a migration run

	   Format: int64
	*/
	RunID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cloud migration run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCloudMigrationRunParams) WithDefaults() *GetCloudMigrationRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cloud migration run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCloudMigrationRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get cloud migration run params
func (o *GetCloudMigrationRunParams) WithTimeout(timeout time.Duration) *GetCloudMigrationRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud migration run params
func (o *GetCloudMigrationRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud migration run params
func (o *GetCloudMigrationRunParams) WithContext(ctx context.Context) *GetCloudMigrationRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud migration run params
func (o *GetCloudMigrationRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud migration run params
func (o *GetCloudMigrationRunParams) WithHTTPClient(client *http.Client) *GetCloudMigrationRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud migration run params
func (o *GetCloudMigrationRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get cloud migration run params
func (o *GetCloudMigrationRunParams) WithID(id int64) *GetCloudMigrationRunParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get cloud migration run params
func (o *GetCloudMigrationRunParams) SetID(id int64) {
	o.ID = id
}

// WithRunID adds the runID to the get cloud migration run params
func (o *GetCloudMigrationRunParams) WithRunID(runID int64) *GetCloudMigrationRunParams {
	o.SetRunID(runID)
	return o
}

// SetRunID adds the runId to the get cloud migration run params
func (o *GetCloudMigrationRunParams) SetRunID(runID int64) {
	o.RunID = runID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudMigrationRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param runID
	if err := r.SetPathParam("runID", swag.FormatInt64(o.RunID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}