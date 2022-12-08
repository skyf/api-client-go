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

	"github.com/firehydrant/api-client-go/models"
)

// NewPostV1SeverityMatrixImpactsParams creates a new PostV1SeverityMatrixImpactsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1SeverityMatrixImpactsParams() *PostV1SeverityMatrixImpactsParams {
	return &PostV1SeverityMatrixImpactsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1SeverityMatrixImpactsParamsWithTimeout creates a new PostV1SeverityMatrixImpactsParams object
// with the ability to set a timeout on a request.
func NewPostV1SeverityMatrixImpactsParamsWithTimeout(timeout time.Duration) *PostV1SeverityMatrixImpactsParams {
	return &PostV1SeverityMatrixImpactsParams{
		timeout: timeout,
	}
}

// NewPostV1SeverityMatrixImpactsParamsWithContext creates a new PostV1SeverityMatrixImpactsParams object
// with the ability to set a context for a request.
func NewPostV1SeverityMatrixImpactsParamsWithContext(ctx context.Context) *PostV1SeverityMatrixImpactsParams {
	return &PostV1SeverityMatrixImpactsParams{
		Context: ctx,
	}
}

// NewPostV1SeverityMatrixImpactsParamsWithHTTPClient creates a new PostV1SeverityMatrixImpactsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1SeverityMatrixImpactsParamsWithHTTPClient(client *http.Client) *PostV1SeverityMatrixImpactsParams {
	return &PostV1SeverityMatrixImpactsParams{
		HTTPClient: client,
	}
}

/*
PostV1SeverityMatrixImpactsParams contains all the parameters to send to the API endpoint

	for the post v1 severity matrix impacts operation.

	Typically these are written to a http.Request.
*/
type PostV1SeverityMatrixImpactsParams struct {

	// V1SeverityMatrixImpacts.
	V1SeverityMatrixImpacts *models.PostV1SeverityMatrixImpacts

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 severity matrix impacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1SeverityMatrixImpactsParams) WithDefaults() *PostV1SeverityMatrixImpactsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 severity matrix impacts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1SeverityMatrixImpactsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 severity matrix impacts params
func (o *PostV1SeverityMatrixImpactsParams) WithTimeout(timeout time.Duration) *PostV1SeverityMatrixImpactsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 severity matrix impacts params
func (o *PostV1SeverityMatrixImpactsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 severity matrix impacts params
func (o *PostV1SeverityMatrixImpactsParams) WithContext(ctx context.Context) *PostV1SeverityMatrixImpactsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 severity matrix impacts params
func (o *PostV1SeverityMatrixImpactsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 severity matrix impacts params
func (o *PostV1SeverityMatrixImpactsParams) WithHTTPClient(client *http.Client) *PostV1SeverityMatrixImpactsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 severity matrix impacts params
func (o *PostV1SeverityMatrixImpactsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1SeverityMatrixImpacts adds the v1SeverityMatrixImpacts to the post v1 severity matrix impacts params
func (o *PostV1SeverityMatrixImpactsParams) WithV1SeverityMatrixImpacts(v1SeverityMatrixImpacts *models.PostV1SeverityMatrixImpacts) *PostV1SeverityMatrixImpactsParams {
	o.SetV1SeverityMatrixImpacts(v1SeverityMatrixImpacts)
	return o
}

// SetV1SeverityMatrixImpacts adds the v1SeverityMatrixImpacts to the post v1 severity matrix impacts params
func (o *PostV1SeverityMatrixImpactsParams) SetV1SeverityMatrixImpacts(v1SeverityMatrixImpacts *models.PostV1SeverityMatrixImpacts) {
	o.V1SeverityMatrixImpacts = v1SeverityMatrixImpacts
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1SeverityMatrixImpactsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1SeverityMatrixImpacts != nil {
		if err := r.SetBodyParam(o.V1SeverityMatrixImpacts); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
