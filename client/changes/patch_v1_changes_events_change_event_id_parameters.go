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
	"github.com/go-openapi/swag"

	"github.com/firehydrant/api-client-go/models"
)

// NewPatchV1ChangesEventsChangeEventIDParams creates a new PatchV1ChangesEventsChangeEventIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1ChangesEventsChangeEventIDParams() *PatchV1ChangesEventsChangeEventIDParams {
	return &PatchV1ChangesEventsChangeEventIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1ChangesEventsChangeEventIDParamsWithTimeout creates a new PatchV1ChangesEventsChangeEventIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1ChangesEventsChangeEventIDParamsWithTimeout(timeout time.Duration) *PatchV1ChangesEventsChangeEventIDParams {
	return &PatchV1ChangesEventsChangeEventIDParams{
		timeout: timeout,
	}
}

// NewPatchV1ChangesEventsChangeEventIDParamsWithContext creates a new PatchV1ChangesEventsChangeEventIDParams object
// with the ability to set a context for a request.
func NewPatchV1ChangesEventsChangeEventIDParamsWithContext(ctx context.Context) *PatchV1ChangesEventsChangeEventIDParams {
	return &PatchV1ChangesEventsChangeEventIDParams{
		Context: ctx,
	}
}

// NewPatchV1ChangesEventsChangeEventIDParamsWithHTTPClient creates a new PatchV1ChangesEventsChangeEventIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1ChangesEventsChangeEventIDParamsWithHTTPClient(client *http.Client) *PatchV1ChangesEventsChangeEventIDParams {
	return &PatchV1ChangesEventsChangeEventIDParams{
		HTTPClient: client,
	}
}

/*
PatchV1ChangesEventsChangeEventIDParams contains all the parameters to send to the API endpoint

	for the patch v1 changes events change event Id operation.

	Typically these are written to a http.Request.
*/
type PatchV1ChangesEventsChangeEventIDParams struct {

	// V1ChangesEvents.
	V1ChangesEvents *models.PatchV1ChangesEvents

	// ChangeEventID.
	//
	// Format: int32
	ChangeEventID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 changes events change event Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ChangesEventsChangeEventIDParams) WithDefaults() *PatchV1ChangesEventsChangeEventIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 changes events change event Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1ChangesEventsChangeEventIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 changes events change event Id params
func (o *PatchV1ChangesEventsChangeEventIDParams) WithTimeout(timeout time.Duration) *PatchV1ChangesEventsChangeEventIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 changes events change event Id params
func (o *PatchV1ChangesEventsChangeEventIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 changes events change event Id params
func (o *PatchV1ChangesEventsChangeEventIDParams) WithContext(ctx context.Context) *PatchV1ChangesEventsChangeEventIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 changes events change event Id params
func (o *PatchV1ChangesEventsChangeEventIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 changes events change event Id params
func (o *PatchV1ChangesEventsChangeEventIDParams) WithHTTPClient(client *http.Client) *PatchV1ChangesEventsChangeEventIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 changes events change event Id params
func (o *PatchV1ChangesEventsChangeEventIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1ChangesEvents adds the v1ChangesEvents to the patch v1 changes events change event Id params
func (o *PatchV1ChangesEventsChangeEventIDParams) WithV1ChangesEvents(v1ChangesEvents *models.PatchV1ChangesEvents) *PatchV1ChangesEventsChangeEventIDParams {
	o.SetV1ChangesEvents(v1ChangesEvents)
	return o
}

// SetV1ChangesEvents adds the v1ChangesEvents to the patch v1 changes events change event Id params
func (o *PatchV1ChangesEventsChangeEventIDParams) SetV1ChangesEvents(v1ChangesEvents *models.PatchV1ChangesEvents) {
	o.V1ChangesEvents = v1ChangesEvents
}

// WithChangeEventID adds the changeEventID to the patch v1 changes events change event Id params
func (o *PatchV1ChangesEventsChangeEventIDParams) WithChangeEventID(changeEventID int32) *PatchV1ChangesEventsChangeEventIDParams {
	o.SetChangeEventID(changeEventID)
	return o
}

// SetChangeEventID adds the changeEventId to the patch v1 changes events change event Id params
func (o *PatchV1ChangesEventsChangeEventIDParams) SetChangeEventID(changeEventID int32) {
	o.ChangeEventID = changeEventID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1ChangesEventsChangeEventIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1ChangesEvents != nil {
		if err := r.SetBodyParam(o.V1ChangesEvents); err != nil {
			return err
		}
	}

	// path param change_event_id
	if err := r.SetPathParam("change_event_id", swag.FormatInt32(o.ChangeEventID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
