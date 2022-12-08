// Code generated by go-swagger; DO NOT EDIT.

package functionalities

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

// NewPostV1FunctionalitiesParams creates a new PostV1FunctionalitiesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1FunctionalitiesParams() *PostV1FunctionalitiesParams {
	return &PostV1FunctionalitiesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1FunctionalitiesParamsWithTimeout creates a new PostV1FunctionalitiesParams object
// with the ability to set a timeout on a request.
func NewPostV1FunctionalitiesParamsWithTimeout(timeout time.Duration) *PostV1FunctionalitiesParams {
	return &PostV1FunctionalitiesParams{
		timeout: timeout,
	}
}

// NewPostV1FunctionalitiesParamsWithContext creates a new PostV1FunctionalitiesParams object
// with the ability to set a context for a request.
func NewPostV1FunctionalitiesParamsWithContext(ctx context.Context) *PostV1FunctionalitiesParams {
	return &PostV1FunctionalitiesParams{
		Context: ctx,
	}
}

// NewPostV1FunctionalitiesParamsWithHTTPClient creates a new PostV1FunctionalitiesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1FunctionalitiesParamsWithHTTPClient(client *http.Client) *PostV1FunctionalitiesParams {
	return &PostV1FunctionalitiesParams{
		HTTPClient: client,
	}
}

/*
PostV1FunctionalitiesParams contains all the parameters to send to the API endpoint

	for the post v1 functionalities operation.

	Typically these are written to a http.Request.
*/
type PostV1FunctionalitiesParams struct {

	// V1Functionalities.
	V1Functionalities *models.PostV1Functionalities

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 functionalities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1FunctionalitiesParams) WithDefaults() *PostV1FunctionalitiesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 functionalities params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1FunctionalitiesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 functionalities params
func (o *PostV1FunctionalitiesParams) WithTimeout(timeout time.Duration) *PostV1FunctionalitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 functionalities params
func (o *PostV1FunctionalitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 functionalities params
func (o *PostV1FunctionalitiesParams) WithContext(ctx context.Context) *PostV1FunctionalitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 functionalities params
func (o *PostV1FunctionalitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 functionalities params
func (o *PostV1FunctionalitiesParams) WithHTTPClient(client *http.Client) *PostV1FunctionalitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 functionalities params
func (o *PostV1FunctionalitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1Functionalities adds the v1Functionalities to the post v1 functionalities params
func (o *PostV1FunctionalitiesParams) WithV1Functionalities(v1Functionalities *models.PostV1Functionalities) *PostV1FunctionalitiesParams {
	o.SetV1Functionalities(v1Functionalities)
	return o
}

// SetV1Functionalities adds the v1Functionalities to the post v1 functionalities params
func (o *PostV1FunctionalitiesParams) SetV1Functionalities(v1Functionalities *models.PostV1Functionalities) {
	o.V1Functionalities = v1Functionalities
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1FunctionalitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1Functionalities != nil {
		if err := r.SetBodyParam(o.V1Functionalities); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
