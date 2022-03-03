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

// NewPatchV1IncidentsIncidentIDNotesNoteIDParams creates a new PatchV1IncidentsIncidentIDNotesNoteIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchV1IncidentsIncidentIDNotesNoteIDParams() *PatchV1IncidentsIncidentIDNotesNoteIDParams {
	return &PatchV1IncidentsIncidentIDNotesNoteIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchV1IncidentsIncidentIDNotesNoteIDParamsWithTimeout creates a new PatchV1IncidentsIncidentIDNotesNoteIDParams object
// with the ability to set a timeout on a request.
func NewPatchV1IncidentsIncidentIDNotesNoteIDParamsWithTimeout(timeout time.Duration) *PatchV1IncidentsIncidentIDNotesNoteIDParams {
	return &PatchV1IncidentsIncidentIDNotesNoteIDParams{
		timeout: timeout,
	}
}

// NewPatchV1IncidentsIncidentIDNotesNoteIDParamsWithContext creates a new PatchV1IncidentsIncidentIDNotesNoteIDParams object
// with the ability to set a context for a request.
func NewPatchV1IncidentsIncidentIDNotesNoteIDParamsWithContext(ctx context.Context) *PatchV1IncidentsIncidentIDNotesNoteIDParams {
	return &PatchV1IncidentsIncidentIDNotesNoteIDParams{
		Context: ctx,
	}
}

// NewPatchV1IncidentsIncidentIDNotesNoteIDParamsWithHTTPClient creates a new PatchV1IncidentsIncidentIDNotesNoteIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchV1IncidentsIncidentIDNotesNoteIDParamsWithHTTPClient(client *http.Client) *PatchV1IncidentsIncidentIDNotesNoteIDParams {
	return &PatchV1IncidentsIncidentIDNotesNoteIDParams{
		HTTPClient: client,
	}
}

/* PatchV1IncidentsIncidentIDNotesNoteIDParams contains all the parameters to send to the API endpoint
   for the patch v1 incidents incident Id notes note Id operation.

   Typically these are written to a http.Request.
*/
type PatchV1IncidentsIncidentIDNotesNoteIDParams struct {

	// V1IncidentsIncidentIDNotes.
	V1IncidentsIncidentIDNotes *models.PatchV1IncidentsIncidentIDNotes

	// IncidentID.
	IncidentID string

	// NoteID.
	NoteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch v1 incidents incident Id notes note Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) WithDefaults() *PatchV1IncidentsIncidentIDNotesNoteIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch v1 incidents incident Id notes note Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch v1 incidents incident Id notes note Id params
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) WithTimeout(timeout time.Duration) *PatchV1IncidentsIncidentIDNotesNoteIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch v1 incidents incident Id notes note Id params
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch v1 incidents incident Id notes note Id params
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) WithContext(ctx context.Context) *PatchV1IncidentsIncidentIDNotesNoteIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch v1 incidents incident Id notes note Id params
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch v1 incidents incident Id notes note Id params
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) WithHTTPClient(client *http.Client) *PatchV1IncidentsIncidentIDNotesNoteIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch v1 incidents incident Id notes note Id params
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1IncidentsIncidentIDNotes adds the v1IncidentsIncidentIDNotes to the patch v1 incidents incident Id notes note Id params
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) WithV1IncidentsIncidentIDNotes(v1IncidentsIncidentIDNotes *models.PatchV1IncidentsIncidentIDNotes) *PatchV1IncidentsIncidentIDNotesNoteIDParams {
	o.SetV1IncidentsIncidentIDNotes(v1IncidentsIncidentIDNotes)
	return o
}

// SetV1IncidentsIncidentIDNotes adds the v1IncidentsIncidentIdNotes to the patch v1 incidents incident Id notes note Id params
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) SetV1IncidentsIncidentIDNotes(v1IncidentsIncidentIDNotes *models.PatchV1IncidentsIncidentIDNotes) {
	o.V1IncidentsIncidentIDNotes = v1IncidentsIncidentIDNotes
}

// WithIncidentID adds the incidentID to the patch v1 incidents incident Id notes note Id params
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) WithIncidentID(incidentID string) *PatchV1IncidentsIncidentIDNotesNoteIDParams {
	o.SetIncidentID(incidentID)
	return o
}

// SetIncidentID adds the incidentId to the patch v1 incidents incident Id notes note Id params
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) SetIncidentID(incidentID string) {
	o.IncidentID = incidentID
}

// WithNoteID adds the noteID to the patch v1 incidents incident Id notes note Id params
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) WithNoteID(noteID string) *PatchV1IncidentsIncidentIDNotesNoteIDParams {
	o.SetNoteID(noteID)
	return o
}

// SetNoteID adds the noteId to the patch v1 incidents incident Id notes note Id params
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) SetNoteID(noteID string) {
	o.NoteID = noteID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchV1IncidentsIncidentIDNotesNoteIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1IncidentsIncidentIDNotes != nil {
		if err := r.SetBodyParam(o.V1IncidentsIncidentIDNotes); err != nil {
			return err
		}
	}

	// path param incident_id
	if err := r.SetPathParam("incident_id", o.IncidentID); err != nil {
		return err
	}

	// path param note_id
	if err := r.SetPathParam("note_id", o.NoteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
