// Code generated by go-swagger; DO NOT EDIT.

package scim

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

// NewGetV1ScimV2UsersIDParams creates a new GetV1ScimV2UsersIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ScimV2UsersIDParams() *GetV1ScimV2UsersIDParams {
	return &GetV1ScimV2UsersIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ScimV2UsersIDParamsWithTimeout creates a new GetV1ScimV2UsersIDParams object
// with the ability to set a timeout on a request.
func NewGetV1ScimV2UsersIDParamsWithTimeout(timeout time.Duration) *GetV1ScimV2UsersIDParams {
	return &GetV1ScimV2UsersIDParams{
		timeout: timeout,
	}
}

// NewGetV1ScimV2UsersIDParamsWithContext creates a new GetV1ScimV2UsersIDParams object
// with the ability to set a context for a request.
func NewGetV1ScimV2UsersIDParamsWithContext(ctx context.Context) *GetV1ScimV2UsersIDParams {
	return &GetV1ScimV2UsersIDParams{
		Context: ctx,
	}
}

// NewGetV1ScimV2UsersIDParamsWithHTTPClient creates a new GetV1ScimV2UsersIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ScimV2UsersIDParamsWithHTTPClient(client *http.Client) *GetV1ScimV2UsersIDParams {
	return &GetV1ScimV2UsersIDParams{
		HTTPClient: client,
	}
}

/*
GetV1ScimV2UsersIDParams contains all the parameters to send to the API endpoint

	for the get v1 scim v2 users Id operation.

	Typically these are written to a http.Request.
*/
type GetV1ScimV2UsersIDParams struct {

	// ID.
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 scim v2 users Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScimV2UsersIDParams) WithDefaults() *GetV1ScimV2UsersIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 scim v2 users Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ScimV2UsersIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 scim v2 users Id params
func (o *GetV1ScimV2UsersIDParams) WithTimeout(timeout time.Duration) *GetV1ScimV2UsersIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 scim v2 users Id params
func (o *GetV1ScimV2UsersIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 scim v2 users Id params
func (o *GetV1ScimV2UsersIDParams) WithContext(ctx context.Context) *GetV1ScimV2UsersIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 scim v2 users Id params
func (o *GetV1ScimV2UsersIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 scim v2 users Id params
func (o *GetV1ScimV2UsersIDParams) WithHTTPClient(client *http.Client) *GetV1ScimV2UsersIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 scim v2 users Id params
func (o *GetV1ScimV2UsersIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 scim v2 users Id params
func (o *GetV1ScimV2UsersIDParams) WithID(id string) *GetV1ScimV2UsersIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 scim v2 users Id params
func (o *GetV1ScimV2UsersIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ScimV2UsersIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
