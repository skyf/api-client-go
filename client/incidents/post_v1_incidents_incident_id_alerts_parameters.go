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
	"github.com/go-openapi/swag"
)

// NewPostV1IncidentsIncidentIDAlertsParams creates a new PostV1IncidentsIncidentIDAlertsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1IncidentsIncidentIDAlertsParams() *PostV1IncidentsIncidentIDAlertsParams {
	return &PostV1IncidentsIncidentIDAlertsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1IncidentsIncidentIDAlertsParamsWithTimeout creates a new PostV1IncidentsIncidentIDAlertsParams object
// with the ability to set a timeout on a request.
func NewPostV1IncidentsIncidentIDAlertsParamsWithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDAlertsParams {
	return &PostV1IncidentsIncidentIDAlertsParams{
		timeout: timeout,
	}
}

// NewPostV1IncidentsIncidentIDAlertsParamsWithContext creates a new PostV1IncidentsIncidentIDAlertsParams object
// with the ability to set a context for a request.
func NewPostV1IncidentsIncidentIDAlertsParamsWithContext(ctx context.Context) *PostV1IncidentsIncidentIDAlertsParams {
	return &PostV1IncidentsIncidentIDAlertsParams{
		Context: ctx,
	}
}

// NewPostV1IncidentsIncidentIDAlertsParamsWithHTTPClient creates a new PostV1IncidentsIncidentIDAlertsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1IncidentsIncidentIDAlertsParamsWithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDAlertsParams {
	return &PostV1IncidentsIncidentIDAlertsParams{
		HTTPClient: client,
	}
}

/*
PostV1IncidentsIncidentIDAlertsParams contains all the parameters to send to the API endpoint

	for the post v1 incidents incident Id alerts operation.

	Typically these are written to a http.Request.
*/
type PostV1IncidentsIncidentIDAlertsParams struct {

	/* AlertIds.

	   Array of alert IDs to be assigned to the incident
	*/
	AlertIds []string

	// IncidentID.
	IncidentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 incidents incident Id alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDAlertsParams) WithDefaults() *PostV1IncidentsIncidentIDAlertsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 incidents incident Id alerts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1IncidentsIncidentIDAlertsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 incidents incident Id alerts params
func (o *PostV1IncidentsIncidentIDAlertsParams) WithTimeout(timeout time.Duration) *PostV1IncidentsIncidentIDAlertsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 incidents incident Id alerts params
func (o *PostV1IncidentsIncidentIDAlertsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 incidents incident Id alerts params
func (o *PostV1IncidentsIncidentIDAlertsParams) WithContext(ctx context.Context) *PostV1IncidentsIncidentIDAlertsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 incidents incident Id alerts params
func (o *PostV1IncidentsIncidentIDAlertsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 incidents incident Id alerts params
func (o *PostV1IncidentsIncidentIDAlertsParams) WithHTTPClient(client *http.Client) *PostV1IncidentsIncidentIDAlertsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 incidents incident Id alerts params
func (o *PostV1IncidentsIncidentIDAlertsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertIds adds the alertIds to the post v1 incidents incident Id alerts params
func (o *PostV1IncidentsIncidentIDAlertsParams) WithAlertIds(alertIds []string) *PostV1IncidentsIncidentIDAlertsParams {
	o.SetAlertIds(alertIds)
	return o
}

// SetAlertIds adds the alertIds to the post v1 incidents incident Id alerts params
func (o *PostV1IncidentsIncidentIDAlertsParams) SetAlertIds(alertIds []string) {
	o.AlertIds = alertIds
}

// WithIncidentID adds the incidentID to the post v1 incidents incident Id alerts params
func (o *PostV1IncidentsIncidentIDAlertsParams) WithIncidentID(incidentID string) *PostV1IncidentsIncidentIDAlertsParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the post v1 incidents incident Id alerts params
func (o *PostV1IncidentsIncidentIDAlertsParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1IncidentsIncidentIDAlertsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AlertIds != nil {

		// binding items for alert_ids
		joinedAlertIds := o.bindParamAlertIds(reg)

		// form array param alert_ids
		if err := r.SetFormParam("alert_ids", joinedAlertIds...); err != nil {
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

// bindParamPostV1IncidentsIncidentIDAlerts binds the parameter alert_ids
func (o *PostV1IncidentsIncidentIDAlertsParams) bindParamAlertIds(formats strfmt.Registry) []string {
	alertIdsIR := o.AlertIds

	var alertIdsIC []string
	for _, alertIdsIIR := range alertIdsIR { // explode []string

		alertIdsIIV := alertIdsIIR // string as string
		alertIdsIC = append(alertIdsIC, alertIdsIIV)
	}

	// items.CollectionFormat: ""
	alertIdsIS := swag.JoinByFormat(alertIdsIC, "")

	return alertIdsIS
}
