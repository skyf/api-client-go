// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewGetV1IntegrationsStatuspageConnectionsIDPagesParams creates a new GetV1IntegrationsStatuspageConnectionsIDPagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1IntegrationsStatuspageConnectionsIDPagesParams() *GetV1IntegrationsStatuspageConnectionsIDPagesParams {
	return &GetV1IntegrationsStatuspageConnectionsIDPagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1IntegrationsStatuspageConnectionsIDPagesParamsWithTimeout creates a new GetV1IntegrationsStatuspageConnectionsIDPagesParams object
// with the ability to set a timeout on a request.
func NewGetV1IntegrationsStatuspageConnectionsIDPagesParamsWithTimeout(timeout time.Duration) *GetV1IntegrationsStatuspageConnectionsIDPagesParams {
	return &GetV1IntegrationsStatuspageConnectionsIDPagesParams{
		timeout: timeout,
	}
}

// NewGetV1IntegrationsStatuspageConnectionsIDPagesParamsWithContext creates a new GetV1IntegrationsStatuspageConnectionsIDPagesParams object
// with the ability to set a context for a request.
func NewGetV1IntegrationsStatuspageConnectionsIDPagesParamsWithContext(ctx context.Context) *GetV1IntegrationsStatuspageConnectionsIDPagesParams {
	return &GetV1IntegrationsStatuspageConnectionsIDPagesParams{
		Context: ctx,
	}
}

// NewGetV1IntegrationsStatuspageConnectionsIDPagesParamsWithHTTPClient creates a new GetV1IntegrationsStatuspageConnectionsIDPagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1IntegrationsStatuspageConnectionsIDPagesParamsWithHTTPClient(client *http.Client) *GetV1IntegrationsStatuspageConnectionsIDPagesParams {
	return &GetV1IntegrationsStatuspageConnectionsIDPagesParams{
		HTTPClient: client,
	}
}

/*
GetV1IntegrationsStatuspageConnectionsIDPagesParams contains all the parameters to send to the API endpoint

	for the get v1 integrations statuspage connections Id pages operation.

	Typically these are written to a http.Request.
*/
type GetV1IntegrationsStatuspageConnectionsIDPagesParams struct {

	/* ID.

	   Connection UUID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 integrations statuspage connections Id pages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsStatuspageConnectionsIDPagesParams) WithDefaults() *GetV1IntegrationsStatuspageConnectionsIDPagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 integrations statuspage connections Id pages params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1IntegrationsStatuspageConnectionsIDPagesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 integrations statuspage connections Id pages params
func (o *GetV1IntegrationsStatuspageConnectionsIDPagesParams) WithTimeout(timeout time.Duration) *GetV1IntegrationsStatuspageConnectionsIDPagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 integrations statuspage connections Id pages params
func (o *GetV1IntegrationsStatuspageConnectionsIDPagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 integrations statuspage connections Id pages params
func (o *GetV1IntegrationsStatuspageConnectionsIDPagesParams) WithContext(ctx context.Context) *GetV1IntegrationsStatuspageConnectionsIDPagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 integrations statuspage connections Id pages params
func (o *GetV1IntegrationsStatuspageConnectionsIDPagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 integrations statuspage connections Id pages params
func (o *GetV1IntegrationsStatuspageConnectionsIDPagesParams) WithHTTPClient(client *http.Client) *GetV1IntegrationsStatuspageConnectionsIDPagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 integrations statuspage connections Id pages params
func (o *GetV1IntegrationsStatuspageConnectionsIDPagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get v1 integrations statuspage connections Id pages params
func (o *GetV1IntegrationsStatuspageConnectionsIDPagesParams) WithID(id string) *GetV1IntegrationsStatuspageConnectionsIDPagesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get v1 integrations statuspage connections Id pages params
func (o *GetV1IntegrationsStatuspageConnectionsIDPagesParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1IntegrationsStatuspageConnectionsIDPagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
