// Code generated by go-swagger; DO NOT EDIT.

package changes

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

// NewGetV1ChangesEventsChangeEventIDParams creates a new GetV1ChangesEventsChangeEventIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1ChangesEventsChangeEventIDParams() *GetV1ChangesEventsChangeEventIDParams {
	return &GetV1ChangesEventsChangeEventIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1ChangesEventsChangeEventIDParamsWithTimeout creates a new GetV1ChangesEventsChangeEventIDParams object
// with the ability to set a timeout on a request.
func NewGetV1ChangesEventsChangeEventIDParamsWithTimeout(timeout time.Duration) *GetV1ChangesEventsChangeEventIDParams {
	return &GetV1ChangesEventsChangeEventIDParams{
		timeout: timeout,
	}
}

// NewGetV1ChangesEventsChangeEventIDParamsWithContext creates a new GetV1ChangesEventsChangeEventIDParams object
// with the ability to set a context for a request.
func NewGetV1ChangesEventsChangeEventIDParamsWithContext(ctx context.Context) *GetV1ChangesEventsChangeEventIDParams {
	return &GetV1ChangesEventsChangeEventIDParams{
		Context: ctx,
	}
}

// NewGetV1ChangesEventsChangeEventIDParamsWithHTTPClient creates a new GetV1ChangesEventsChangeEventIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1ChangesEventsChangeEventIDParamsWithHTTPClient(client *http.Client) *GetV1ChangesEventsChangeEventIDParams {
	return &GetV1ChangesEventsChangeEventIDParams{
		HTTPClient: client,
	}
}

/*
GetV1ChangesEventsChangeEventIDParams contains all the parameters to send to the API endpoint

	for the get v1 changes events change event Id operation.

	Typically these are written to a http.Request.
*/
type GetV1ChangesEventsChangeEventIDParams struct {

	// ChangeEventID.
	ChangeEventID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 changes events change event Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ChangesEventsChangeEventIDParams) WithDefaults() *GetV1ChangesEventsChangeEventIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 changes events change event Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1ChangesEventsChangeEventIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 changes events change event Id params
func (o *GetV1ChangesEventsChangeEventIDParams) WithTimeout(timeout time.Duration) *GetV1ChangesEventsChangeEventIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 changes events change event Id params
func (o *GetV1ChangesEventsChangeEventIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 changes events change event Id params
func (o *GetV1ChangesEventsChangeEventIDParams) WithContext(ctx context.Context) *GetV1ChangesEventsChangeEventIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 changes events change event Id params
func (o *GetV1ChangesEventsChangeEventIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 changes events change event Id params
func (o *GetV1ChangesEventsChangeEventIDParams) WithHTTPClient(client *http.Client) *GetV1ChangesEventsChangeEventIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 changes events change event Id params
func (o *GetV1ChangesEventsChangeEventIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeEventID adds the changeEventID to the get v1 changes events change event Id params
func (o *GetV1ChangesEventsChangeEventIDParams) WithChangeEventID(changeEventID string) *GetV1ChangesEventsChangeEventIDParams {
	o.SetChangeEventID(changeEventID)
	return o
}

// SetChangeEventID adds the changeEventId to the get v1 changes events change event Id params
func (o *GetV1ChangesEventsChangeEventIDParams) SetChangeEventID(changeEventID string) {
	o.ChangeEventID = changeEventID
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1ChangesEventsChangeEventIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param change_event_id
	if err := r.SetPathParam("change_event_id", o.ChangeEventID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
