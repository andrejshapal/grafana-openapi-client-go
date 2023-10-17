// Code generated by go-swagger; DO NOT EDIT.

package orgs

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

	"github.com/grafana/grafana-openapi-client-go/models"
)

// NewUpdateOrgUserParams creates a new UpdateOrgUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateOrgUserParams() *UpdateOrgUserParams {
	return &UpdateOrgUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOrgUserParamsWithTimeout creates a new UpdateOrgUserParams object
// with the ability to set a timeout on a request.
func NewUpdateOrgUserParamsWithTimeout(timeout time.Duration) *UpdateOrgUserParams {
	return &UpdateOrgUserParams{
		timeout: timeout,
	}
}

// NewUpdateOrgUserParamsWithContext creates a new UpdateOrgUserParams object
// with the ability to set a context for a request.
func NewUpdateOrgUserParamsWithContext(ctx context.Context) *UpdateOrgUserParams {
	return &UpdateOrgUserParams{
		Context: ctx,
	}
}

// NewUpdateOrgUserParamsWithHTTPClient creates a new UpdateOrgUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateOrgUserParamsWithHTTPClient(client *http.Client) *UpdateOrgUserParams {
	return &UpdateOrgUserParams{
		HTTPClient: client,
	}
}

/*
UpdateOrgUserParams contains all the parameters to send to the API endpoint

	for the update org user operation.

	Typically these are written to a http.Request.
*/
type UpdateOrgUserParams struct {

	// Body.
	Body *models.UpdateOrgUserCommand

	// OrgID.
	//
	// Format: int64
	OrgID int64

	// UserID.
	//
	// Format: int64
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update org user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrgUserParams) WithDefaults() *UpdateOrgUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update org user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateOrgUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update org user params
func (o *UpdateOrgUserParams) WithTimeout(timeout time.Duration) *UpdateOrgUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update org user params
func (o *UpdateOrgUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update org user params
func (o *UpdateOrgUserParams) WithContext(ctx context.Context) *UpdateOrgUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update org user params
func (o *UpdateOrgUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update org user params
func (o *UpdateOrgUserParams) WithHTTPClient(client *http.Client) *UpdateOrgUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update org user params
func (o *UpdateOrgUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update org user params
func (o *UpdateOrgUserParams) WithBody(body *models.UpdateOrgUserCommand) *UpdateOrgUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update org user params
func (o *UpdateOrgUserParams) SetBody(body *models.UpdateOrgUserCommand) {
	o.Body = body
}

// WithOrgID adds the orgID to the update org user params
func (o *UpdateOrgUserParams) WithOrgID(orgID int64) *UpdateOrgUserParams {
	o.SetOrgID(orgID)
	return o
}

// SetOrgID adds the orgId to the update org user params
func (o *UpdateOrgUserParams) SetOrgID(orgID int64) {
	o.OrgID = orgID
}

// WithUserID adds the userID to the update org user params
func (o *UpdateOrgUserParams) WithUserID(userID int64) *UpdateOrgUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the update org user params
func (o *UpdateOrgUserParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOrgUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param org_id
	if err := r.SetPathParam("org_id", swag.FormatInt64(o.OrgID)); err != nil {
		return err
	}

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}