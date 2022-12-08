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

	"github.com/firehydrant/api-client-go/models"
)

// NewPatchV1IntegrationsAwsConnectionsIDParams creates a new PatchV1IntegrationsAwsConnectionsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1IntegrationsAwsConnectionsIDParams() *PatchV1IntegrationsAwsConnectionsIDParams {
	return &PatchV1IntegrationsAwsConnectionsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1IntegrationsAwsConnectionsIDParamsWithTimeout creates a new PatchV1IntegrationsAwsConnectionsIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1IntegrationsAwsConnectionsIDParamsWithTimeout(timeout time.Duration) *PatchV1IntegrationsAwsConnectionsIDParams {
	return &PatchV1IntegrationsAwsConnectionsIDParams{
		timeout: timeout,
	}
}

// NewPatchV1IntegrationsAwsConnectionsIDParamsWithContext creates a new PatchV1IntegrationsAwsConnectionsIDParams object
// with the ability to set a context for a request.
func NewPatchV1IntegrationsAwsConnectionsIDParamsWithContext(ctx context.Context) *PatchV1IntegrationsAwsConnectionsIDParams {
	return &PatchV1IntegrationsAwsConnectionsIDParams{
		Context: ctx,
	}
}

// NewPatchV1IntegrationsAwsConnectionsIDParamsWithHTTPClient creates a new PatchV1IntegrationsAwsConnectionsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1IntegrationsAwsConnectionsIDParamsWithHTTPClient(client *http.Client) *PatchV1IntegrationsAwsConnectionsIDParams {
	return &PatchV1IntegrationsAwsConnectionsIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1IntegrationsAwsConnectionsIDParams contains all the parameters to send to the API endpoint

	for the patch v1 integrations aws connections Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1IntegrationsAwsConnectionsIDParams struct {

	// V1IntegrationsAwsConnections.
	V1IntegrationsAwsConnections *models.PatchV1IntegrationsAwsConnections

	/* ID.

	   Connection UUID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 integrations aws connections Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IntegrationsAwsConnectionsIDParams) WithDefaults() *PatchV1IntegrationsAwsConnectionsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 integrations aws connections Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IntegrationsAwsConnectionsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 integrations aws connections Id params
func (o *PatchV1IntegrationsAwsConnectionsIDParams) WithTimeout(timeout time.Duration) *PatchV1IntegrationsAwsConnectionsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 integrations aws connections Id params
func (o *PatchV1IntegrationsAwsConnectionsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 integrations aws connections Id params
func (o *PatchV1IntegrationsAwsConnectionsIDParams) WithContext(ctx context.Context) *PatchV1IntegrationsAwsConnectionsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 integrations aws connections Id params
func (o *PatchV1IntegrationsAwsConnectionsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 integrations aws connections Id params
func (o *PatchV1IntegrationsAwsConnectionsIDParams) WithHTTPClient(client *http.Client) *PatchV1IntegrationsAwsConnectionsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 integrations aws connections Id params
func (o *PatchV1IntegrationsAwsConnectionsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1IntegrationsAwsConnections adds the v1IntegrationsAwsConnections to the patch v1 integrations aws connections Id params
func (o *PatchV1IntegrationsAwsConnectionsIDParams) WithV1IntegrationsAwsConnections(v1IntegrationsAwsConnections *models.PatchV1IntegrationsAwsConnections) *PatchV1IntegrationsAwsConnectionsIDParams {
	o.SetV1IntegrationsAwsConnections(v1IntegrationsAwsConnections)
	return o
}

// SetV1IntegrationsAwsConnections adds the v1IntegrationsAwsConnections to the patch v1 integrations aws connections Id params
func (o *PatchV1IntegrationsAwsConnectionsIDParams) SetV1IntegrationsAwsConnections(v1IntegrationsAwsConnections *models.PatchV1IntegrationsAwsConnections) {
	o.V1IntegrationsAwsConnections = v1IntegrationsAwsConnections
}

// WithID adds the id to the patch v1 integrations aws connections Id params
func (o *PatchV1IntegrationsAwsConnectionsIDParams) WithID(id string) *PatchV1IntegrationsAwsConnectionsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch v1 integrations aws connections Id params
func (o *PatchV1IntegrationsAwsConnectionsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1IntegrationsAwsConnectionsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1IntegrationsAwsConnections != nil {
		if err := r.SetBodyParam(o.V1IntegrationsAwsConnections); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
