// Code generated by go-swagger; DO NOT EDIT.

package incidents

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

// NewPostV1IncidentsIncidentIDLinksParams creates a new PostV1IncidentsIncidentIDLinksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1IncidentsIncidentIDLinksParams() *PostV1IncidentsIncidentIDLinksParams {
	return &PostV1IncidentsIncidentIDLinksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1IncidentsIncidentIDLinksParamsWithTimeout creates a new PostV1IncidentsIncidentIDLinksParams object
// with the ability to set a timeout on a request.
func NewPostV1IncidentsIncidentIDLinksParamsWithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDLinksParams {
	return &PostV1IncidentsIncidentIDLinksParams{
		timeout: timeout,
	}
}

// NewPostV1IncidentsIncidentIDLinksParamsWithContext creates a new PostV1IncidentsIncidentIDLinksParams object
// with the ability to set a context for a request.
func NewPostV1IncidentsIncidentIDLinksParamsWithContext(ctx context.Context) *PostV1IncidentsIncidentIDLinksParams {
	return &PostV1IncidentsIncidentIDLinksParams{
		Context: ctx,
	}
}

// NewPostV1IncidentsIncidentIDLinksParamsWithHTTPClient creates a new PostV1IncidentsIncidentIDLinksParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1IncidentsIncidentIDLinksParamsWithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDLinksParams {
	return &PostV1IncidentsIncidentIDLinksParams{
		HTTPClient: client,
	}
}

/*
PostV1IncidentsIncidentIDLinksParams contains all the parameters to send to the API endpoint

	for the post v1 incidents incident Id links operation.

	Typically these are written to a http.Request.
*/
type PostV1IncidentsIncidentIDLinksParams struct {

	// V1IncidentsIncidentIDLinks.
	V1IncidentsIncidentIDLinks *models.PostV1IncidentsIncidentIDLinks

	// IncidentID.
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 incidents incident Id links params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDLinksParams) WithDefaults() *PostV1IncidentsIncidentIDLinksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 incidents incident Id links params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDLinksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 incidents incident Id links params
func (o *PostV1IncidentsIncidentIDLinksParams) WithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDLinksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 incidents incident Id links params
func (o *PostV1IncidentsIncidentIDLinksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 incidents incident Id links params
func (o *PostV1IncidentsIncidentIDLinksParams) WithContext(ctx context.Context) *PostV1IncidentsIncidentIDLinksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 incidents incident Id links params
func (o *PostV1IncidentsIncidentIDLinksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 incidents incident Id links params
func (o *PostV1IncidentsIncidentIDLinksParams) WithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDLinksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 incidents incident Id links params
func (o *PostV1IncidentsIncidentIDLinksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1IncidentsIncidentIDLinks adds the v1IncidentsIncidentIDLinks to the post v1 incidents incident Id links params
func (o *PostV1IncidentsIncidentIDLinksParams) WithV1IncidentsIncidentIDLinks(v1IncidentsIncidentIDLinks *models.PostV1IncidentsIncidentIDLinks) *PostV1IncidentsIncidentIDLinksParams {
	o.SetV1IncidentsIncidentIDLinks(v1IncidentsIncidentIDLinks)
	return o
}

// SetV1IncidentsIncidentIDLinks adds the v1IncidentsIncidentIdLinks to the post v1 incidents incident Id links params
func (o *PostV1IncidentsIncidentIDLinksParams) SetV1IncidentsIncidentIDLinks(v1IncidentsIncidentIDLinks *models.PostV1IncidentsIncidentIDLinks) {
	o.V1IncidentsIncidentIDLinks = v1IncidentsIncidentIDLinks
}

// WithIncidentID adds the incidentID to the post v1 incidents incident Id links params
func (o *PostV1IncidentsIncidentIDLinksParams) WithIncidentID(incidentID string) *PostV1IncidentsIncidentIDLinksParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the post v1 incidents incident Id links params
func (o *PostV1IncidentsIncidentIDLinksParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1IncidentsIncidentIDLinksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1IncidentsIncidentIDLinks != nil {
		if err := r.SetBodyParam(o.V1IncidentsIncidentIDLinks); err != nil {
			return err
		}
	}

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
