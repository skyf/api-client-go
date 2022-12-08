// Code generated by go-swagger; DO NOT EDIT.

package scheduled_maintenances

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

// NewPostV1ScheduledMaintenancesParams creates a new PostV1ScheduledMaintenancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1ScheduledMaintenancesParams() *PostV1ScheduledMaintenancesParams {
	return &PostV1ScheduledMaintenancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ScheduledMaintenancesParamsWithTimeout creates a new PostV1ScheduledMaintenancesParams object
// with the ability to set a timeout on a request.
func NewPostV1ScheduledMaintenancesParamsWithTimeout(timeout time.Duration) *PostV1ScheduledMaintenancesParams {
	return &PostV1ScheduledMaintenancesParams{
		timeout: timeout,
	}
}

// NewPostV1ScheduledMaintenancesParamsWithContext creates a new PostV1ScheduledMaintenancesParams object
// with the ability to set a context for a request.
func NewPostV1ScheduledMaintenancesParamsWithContext(ctx context.Context) *PostV1ScheduledMaintenancesParams {
	return &PostV1ScheduledMaintenancesParams{
		Context: ctx,
	}
}

// NewPostV1ScheduledMaintenancesParamsWithHTTPClient creates a new PostV1ScheduledMaintenancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1ScheduledMaintenancesParamsWithHTTPClient(client *http.Client) *PostV1ScheduledMaintenancesParams {
	return &PostV1ScheduledMaintenancesParams{
		HTTPClient: client,
	}
}

/*
PostV1ScheduledMaintenancesParams contains all the parameters to send to the API endpoint

	for the post v1 scheduled maintenances operation.

	Typically these are written to a http.Request.
*/
type PostV1ScheduledMaintenancesParams struct {

	// V1ScheduledMaintenances.
	V1ScheduledMaintenances *models.PostV1ScheduledMaintenances

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 scheduled maintenances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ScheduledMaintenancesParams) WithDefaults() *PostV1ScheduledMaintenancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 scheduled maintenances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ScheduledMaintenancesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 scheduled maintenances params
func (o *PostV1ScheduledMaintenancesParams) WithTimeout(timeout time.Duration) *PostV1ScheduledMaintenancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 scheduled maintenances params
func (o *PostV1ScheduledMaintenancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 scheduled maintenances params
func (o *PostV1ScheduledMaintenancesParams) WithContext(ctx context.Context) *PostV1ScheduledMaintenancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 scheduled maintenances params
func (o *PostV1ScheduledMaintenancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 scheduled maintenances params
func (o *PostV1ScheduledMaintenancesParams) WithHTTPClient(client *http.Client) *PostV1ScheduledMaintenancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 scheduled maintenances params
func (o *PostV1ScheduledMaintenancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1ScheduledMaintenances adds the v1ScheduledMaintenances to the post v1 scheduled maintenances params
func (o *PostV1ScheduledMaintenancesParams) WithV1ScheduledMaintenances(v1ScheduledMaintenances *models.PostV1ScheduledMaintenances) *PostV1ScheduledMaintenancesParams {
	o.SetV1ScheduledMaintenances(v1ScheduledMaintenances)
	return o
}

// SetV1ScheduledMaintenances adds the v1ScheduledMaintenances to the post v1 scheduled maintenances params
func (o *PostV1ScheduledMaintenancesParams) SetV1ScheduledMaintenances(v1ScheduledMaintenances *models.PostV1ScheduledMaintenances) {
	o.V1ScheduledMaintenances = v1ScheduledMaintenances
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ScheduledMaintenancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1ScheduledMaintenances != nil {
		if err := r.SetBodyParam(o.V1ScheduledMaintenances); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
