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
	"github.com/go-openapi/swag"
)

// NewPatchV1IntegrationsConnectionsSlugConnectionIDParams creates a new PatchV1IntegrationsConnectionsSlugConnectionIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1IntegrationsConnectionsSlugConnectionIDParams() *PatchV1IntegrationsConnectionsSlugConnectionIDParams {
	return &PatchV1IntegrationsConnectionsSlugConnectionIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1IntegrationsConnectionsSlugConnectionIDParamsWithTimeout creates a new PatchV1IntegrationsConnectionsSlugConnectionIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1IntegrationsConnectionsSlugConnectionIDParamsWithTimeout(timeout time.Duration) *PatchV1IntegrationsConnectionsSlugConnectionIDParams {
	return &PatchV1IntegrationsConnectionsSlugConnectionIDParams{
		timeout: timeout,
	}
}

// NewPatchV1IntegrationsConnectionsSlugConnectionIDParamsWithContext creates a new PatchV1IntegrationsConnectionsSlugConnectionIDParams object
// with the ability to set a context for a request.
func NewPatchV1IntegrationsConnectionsSlugConnectionIDParamsWithContext(ctx context.Context) *PatchV1IntegrationsConnectionsSlugConnectionIDParams {
	return &PatchV1IntegrationsConnectionsSlugConnectionIDParams{
		Context: ctx,
	}
}

// NewPatchV1IntegrationsConnectionsSlugConnectionIDParamsWithHTTPClient creates a new PatchV1IntegrationsConnectionsSlugConnectionIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1IntegrationsConnectionsSlugConnectionIDParamsWithHTTPClient(client *http.Client) *PatchV1IntegrationsConnectionsSlugConnectionIDParams {
	return &PatchV1IntegrationsConnectionsSlugConnectionIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1IntegrationsConnectionsSlugConnectionIDParams contains all the parameters to send to the API endpoint

	for the patch v1 integrations connections slug connection Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1IntegrationsConnectionsSlugConnectionIDParams struct {

	// ConnectionID.
	//
	// Format: int32
	ConnectionID int32

	// Slug.
	//
	// Format: int32
	Slug int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 integrations connections slug connection Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) WithDefaults() *PatchV1IntegrationsConnectionsSlugConnectionIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 integrations connections slug connection Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 integrations connections slug connection Id params
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) WithTimeout(timeout time.Duration) *PatchV1IntegrationsConnectionsSlugConnectionIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 integrations connections slug connection Id params
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 integrations connections slug connection Id params
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) WithContext(ctx context.Context) *PatchV1IntegrationsConnectionsSlugConnectionIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 integrations connections slug connection Id params
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 integrations connections slug connection Id params
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) WithHTTPClient(client *http.Client) *PatchV1IntegrationsConnectionsSlugConnectionIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 integrations connections slug connection Id params
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionID adds the connectionID to the patch v1 integrations connections slug connection Id params
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) WithConnectionID(connectionID int32) *PatchV1IntegrationsConnectionsSlugConnectionIDParams {
	o.SetConnectionID(connectionID)
	return o
}

// SetConnectionID adds the connectionId to the patch v1 integrations connections slug connection Id params
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) SetConnectionID(connectionID int32) {
	o.ConnectionID = connectionID
}

// WithSlug adds the slug to the patch v1 integrations connections slug connection Id params
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) WithSlug(slug int32) *PatchV1IntegrationsConnectionsSlugConnectionIDParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the patch v1 integrations connections slug connection Id params
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) SetSlug(slug int32) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1IntegrationsConnectionsSlugConnectionIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_id
	if err := r.SetPathParam("connection_id", swag.FormatInt32(o.ConnectionID)); err != nil {
		return err
	}

	// path param slug
	if err := r.SetPathParam("slug", swag.FormatInt32(o.Slug)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
