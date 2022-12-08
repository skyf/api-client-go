// Code generated by go-swagger; DO NOT EDIT.

package severity_matrix

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

// NewGetV1SeverityMatrixConditionsParams creates a new GetV1SeverityMatrixConditionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1SeverityMatrixConditionsParams() *GetV1SeverityMatrixConditionsParams {
	return &GetV1SeverityMatrixConditionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1SeverityMatrixConditionsParamsWithTimeout creates a new GetV1SeverityMatrixConditionsParams object
// with the ability to set a timeout on a request.
func NewGetV1SeverityMatrixConditionsParamsWithTimeout(timeout time.Duration) *GetV1SeverityMatrixConditionsParams {
	return &GetV1SeverityMatrixConditionsParams{
		timeout: timeout,
	}
}

// NewGetV1SeverityMatrixConditionsParamsWithContext creates a new GetV1SeverityMatrixConditionsParams object
// with the ability to set a context for a request.
func NewGetV1SeverityMatrixConditionsParamsWithContext(ctx context.Context) *GetV1SeverityMatrixConditionsParams {
	return &GetV1SeverityMatrixConditionsParams{
		Context: ctx,
	}
}

// NewGetV1SeverityMatrixConditionsParamsWithHTTPClient creates a new GetV1SeverityMatrixConditionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1SeverityMatrixConditionsParamsWithHTTPClient(client *http.Client) *GetV1SeverityMatrixConditionsParams {
	return &GetV1SeverityMatrixConditionsParams{
		HTTPClient: client,
	}
}

/*
GetV1SeverityMatrixConditionsParams contains all the parameters to send to the API endpoint

	for the get v1 severity matrix conditions operation.

	Typically these are written to a http.Request.
*/
type GetV1SeverityMatrixConditionsParams struct {

	// Page.
	//
	// Format: int32
	Page *int32

	// PerPage.
	//
	// Format: int32
	PerPage *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 severity matrix conditions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SeverityMatrixConditionsParams) WithDefaults() *GetV1SeverityMatrixConditionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 severity matrix conditions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1SeverityMatrixConditionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 severity matrix conditions params
func (o *GetV1SeverityMatrixConditionsParams) WithTimeout(timeout time.Duration) *GetV1SeverityMatrixConditionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 severity matrix conditions params
func (o *GetV1SeverityMatrixConditionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 severity matrix conditions params
func (o *GetV1SeverityMatrixConditionsParams) WithContext(ctx context.Context) *GetV1SeverityMatrixConditionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 severity matrix conditions params
func (o *GetV1SeverityMatrixConditionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 severity matrix conditions params
func (o *GetV1SeverityMatrixConditionsParams) WithHTTPClient(client *http.Client) *GetV1SeverityMatrixConditionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 severity matrix conditions params
func (o *GetV1SeverityMatrixConditionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get v1 severity matrix conditions params
func (o *GetV1SeverityMatrixConditionsParams) WithPage(page *int32) *GetV1SeverityMatrixConditionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 severity matrix conditions params
func (o *GetV1SeverityMatrixConditionsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the get v1 severity matrix conditions params
func (o *GetV1SeverityMatrixConditionsParams) WithPerPage(perPage *int32) *GetV1SeverityMatrixConditionsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get v1 severity matrix conditions params
func (o *GetV1SeverityMatrixConditionsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1SeverityMatrixConditionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
