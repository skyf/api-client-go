// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PatchV1TaskLists Updates a task list's attributes and task list items
//
// swagger:model patchV1TaskLists
type PatchV1TaskLists struct {

	// description
	Description string `json:"description,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// task list items
	TaskListItems []*PatchV1TaskListsTaskListItemsItems0 `json:"task_list_items"`
}

// Validate validates this patch v1 task lists
func (m *PatchV1TaskLists) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTaskListItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1TaskLists) validateTaskListItems(formats strfmt.Registry) error {
	if swag.IsZero(m.TaskListItems) { // not required
		return nil
	}

	for i := 0; i < len(m.TaskListItems); i++ {
		if swag.IsZero(m.TaskListItems[i]) { // not required
			continue
		}

		if m.TaskListItems[i] != nil {
			if err := m.TaskListItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("task_list_items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("task_list_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this patch v1 task lists based on the context it is used
func (m *PatchV1TaskLists) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTaskListItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1TaskLists) contextValidateTaskListItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TaskListItems); i++ {

		if m.TaskListItems[i] != nil {
			if err := m.TaskListItems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("task_list_items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("task_list_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1TaskLists) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1TaskLists) UnmarshalBinary(b []byte) error {
	var res PatchV1TaskLists
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PatchV1TaskListsTaskListItemsItems0 patch v1 task lists task list items items0
//
// swagger:model PatchV1TaskListsTaskListItemsItems0
type PatchV1TaskListsTaskListItemsItems0 struct {

	// A long-form description for the task if additional context is helpful
	Description string `json:"description,omitempty"`

	// A summary of the task
	// Required: true
	Summary *string `json:"summary"`
}

// Validate validates this patch v1 task lists task list items items0
func (m *PatchV1TaskListsTaskListItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummary(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchV1TaskListsTaskListItemsItems0) validateSummary(formats strfmt.Registry) error {

	if err := validate.Required("summary", "body", m.Summary); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch v1 task lists task list items items0 based on context it is used
func (m *PatchV1TaskListsTaskListItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchV1TaskListsTaskListItemsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchV1TaskListsTaskListItemsItems0) UnmarshalBinary(b []byte) error {
	var res PatchV1TaskListsTaskListItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
