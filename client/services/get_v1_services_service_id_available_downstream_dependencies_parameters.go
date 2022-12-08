// Code generated by go-swagger; DO NOT EDIT.

package services

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

// NewGetV1ServicesServiceIDAvailableDownstreamDependenciesParams creates a new GetV1ServicesServiceIDAvailableDownstreamDependenciesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ServicesServiceIDAvailableDownstreamDependenciesParams() *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams {
	return &GetV1ServicesServiceIDAvailableDownstreamDependenciesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ServicesServiceIDAvailableDownstreamDependenciesParamsWithTimeout creates a new GetV1ServicesServiceIDAvailableDownstreamDependenciesParams object
// with the ability to set a timeout on a request.
func NewGetV1ServicesServiceIDAvailableDownstreamDependenciesParamsWithTimeout(timeout time.Duration) *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams {
	return &GetV1ServicesServiceIDAvailableDownstreamDependenciesParams{
		timeout: timeout,
	}
}

// NewGetV1ServicesServiceIDAvailableDownstreamDependenciesParamsWithContext creates a new GetV1ServicesServiceIDAvailableDownstreamDependenciesParams object
// with the ability to set a context for a request.
func NewGetV1ServicesServiceIDAvailableDownstreamDependenciesParamsWithContext(ctx context.Context) *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams {
	return &GetV1ServicesServiceIDAvailableDownstreamDependenciesParams{
		Context: ctx,
	}
}

// NewGetV1ServicesServiceIDAvailableDownstreamDependenciesParamsWithHTTPClient creates a new GetV1ServicesServiceIDAvailableDownstreamDependenciesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ServicesServiceIDAvailableDownstreamDependenciesParamsWithHTTPClient(client *http.Client) *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams {
	return &GetV1ServicesServiceIDAvailableDownstreamDependenciesParams{
		HTTPClient: client,
	}
}

/*
GetV1ServicesServiceIDAvailableDownstreamDependenciesParams contains all the parameters to send to the API endpoint

	for the get v1 services service Id available downstream dependencies operation.

	Typically these are written to a http.Request.
*/
type GetV1ServicesServiceIDAvailableDownstreamDependenciesParams struct {

	// ServiceID.
	//
	// Format: int32
	ServiceID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 services service Id available downstream dependencies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams) WithDefaults() *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 services service Id available downstream dependencies params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 services service Id available downstream dependencies params
func (o *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams) WithTimeout(timeout time.Duration) *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 services service Id available downstream dependencies params
func (o *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 services service Id available downstream dependencies params
func (o *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams) WithContext(ctx context.Context) *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 services service Id available downstream dependencies params
func (o *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 services service Id available downstream dependencies params
func (o *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams) WithHTTPClient(client *http.Client) *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 services service Id available downstream dependencies params
func (o *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceID adds the serviceID to the get v1 services service Id available downstream dependencies params
func (o *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams) WithServiceID(serviceID int32) *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the get v1 services service Id available downstream dependencies params
func (o *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams) SetServiceID(serviceID int32) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ServicesServiceIDAvailableDownstreamDependenciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param service_id
	if err := r.SetPathParam("service_id", swag.FormatInt32(o.ServiceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
