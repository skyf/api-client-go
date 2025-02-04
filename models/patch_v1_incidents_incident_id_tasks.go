// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchV1IncidentsIncidentIDTasks Update a task's attributes
//
// swagger:model patchV1IncidentsIncidentIdTasks
type PatchV1IncidentsIncidentIDTasks struct {

	// The ID of the user assigned to the task.
	AssigneeID string `json:"assignee_id,omitempty"`

	// A description of what the task is for.
	Description string `json:"description,omitempty"`

	// The state of the task. One of: open, in_progress, cancelled, done
	State string `json:"state,omitempty"`

	// The title of the task.
	Title string `json:"title,omitempty"`
}

// Validate validates this patch v1 incidents incident Id tasks
func (m *PatchV1IncidentsIncidentIDTasks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch v1 incidents incident Id tasks based on context it is used
func (m *PatchV1IncidentsIncidentIDTasks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDTasks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1IncidentsIncidentIDTasks) UnmarshalBinary(b []byte) error {
	var res PatchV1IncidentsIncidentIDTasks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
