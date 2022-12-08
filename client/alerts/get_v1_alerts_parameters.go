// Code generated by go-swagger; DO NOT EDIT.

package alerts

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

// NewGetV1AlertsParams creates a new GetV1AlertsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1AlertsParams() *GetV1AlertsParams {
	return &GetV1AlertsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1AlertsParamsWithTimeout creates a new GetV1AlertsParams object
// with the ability to set a timeout on a request.
func NewGetV1AlertsParamsWithTimeout(timeout time.Duration) *GetV1AlertsParams {
	return &GetV1AlertsParams{
		timeout: timeout,
	}
}

// NewGetV1AlertsParamsWithContext creates a new GetV1AlertsParams object
// with the ability to set a context for a request.
func NewGetV1AlertsParamsWithContext(ctx context.Context) *GetV1AlertsParams {
	return &GetV1AlertsParams{
		Context: ctx,
	}
}

// NewGetV1AlertsParamsWithHTTPClient creates a new GetV1AlertsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1AlertsParamsWithHTTPClient(client *http.Client) *GetV1AlertsParams {
	return &GetV1AlertsParams{
		HTTPClient: client,
	}
}

/*
GetV1AlertsParams contains all the parameters to send to the API endpoint

	for the get v1 alerts operation.

	Typically these are written to a http.Request.
*/
type GetV1AlertsParams struct {

	/* Environments.

	   A comma separated list of environment IDs
	*/
	Environments *string

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	/* Query.

	   A text query for alerts
	*/
	Query *string

	/* Services.

	   A comma separated list of service IDs
	*/
	Services *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AlertsParams) WithDefaults() *GetV1AlertsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1AlertsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 alerts params
func (o *GetV1AlertsParams) WithTimeout(timeout time.Duration) *GetV1AlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 alerts params
func (o *GetV1AlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 alerts params
func (o *GetV1AlertsParams) WithContext(ctx context.Context) *GetV1AlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 alerts params
func (o *GetV1AlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 alerts params
func (o *GetV1AlertsParams) WithHTTPClient(client *http.Client) *GetV1AlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 alerts params
func (o *GetV1AlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnvironments adds the environments to the get v1 alerts params
func (o *GetV1AlertsParams) WithEnvironments(environments *string) *GetV1AlertsParams {
	o.SetEnvironments(environments)
	return o
}

// SetEnvironments adds the environments to the get v1 alerts params
func (o *GetV1AlertsParams) SetEnvironments(environments *string) {
	o.Environments = environments
}

// WithPage adds the page to the get v1 alerts params
func (o *GetV1AlertsParams) WithPage(page *int32) *GetV1AlertsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 alerts params
func (o *GetV1AlertsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 alerts params
func (o *GetV1AlertsParams) WithPerPage(perPage *int32) *GetV1AlertsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 alerts params
func (o *GetV1AlertsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithQuery adds the query to the get v1 alerts params
func (o *GetV1AlertsParams) WithQuery(query *string) *GetV1AlertsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get v1 alerts params
func (o *GetV1AlertsParams) SetQuery(query *string) {
	o.Query = query
}

// WithServices adds the services to the get v1 alerts params
func (o *GetV1AlertsParams) WithServices(services *string) *GetV1AlertsParams {
	o.SetServices(services)
	return o
}

// SetServices adds the services to the get v1 alerts params
func (o *GetV1AlertsParams) SetServices(services *string) {
	o.Services = services
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1AlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Environments != nil {

		// query param environments
		var qrEnvironments string

		if o.Environments != nil {
			qrEnvironments = *o.Environments
		}
		qEnvironments := qrEnvironments
		if qEnvironments != "" {

			if err := r.SetQueryParam("environments", qEnvironments); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int32

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if o.Services != nil {

		// query param services
		var qrServices string

		if o.Services != nil {
			qrServices = *o.Services
		}
		qServices := qrServices
		if qServices != "" {

			if err := r.SetQueryParam("services", qServices); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
