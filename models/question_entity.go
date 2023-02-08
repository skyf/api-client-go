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
)

// QuestionEntity question entity
//
// swagger:model QuestionEntity
type QuestionEntity struct {

	// available options
	AvailableOptions []string `json:"available_options"`

	// body
	Body string `json:"body,omitempty"`

	// conversations
	Conversations []*Reference `json:"conversations"`

	// id
	ID string `json:"id,omitempty"`

	// is required
	IsRequired string `json:"is_required,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// tooltip
	Tooltip string `json:"tooltip,omitempty"`
}

// Validate validates this question entity
func (m *QuestionEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConversations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuestionEntity) validateConversations(formats strfmt.Registry) error {
	if swag.IsZero(m.Conversations) { // not required
		return nil
	}

	for i := 0; i < len(m.Conversations); i++ {
		if swag.IsZero(m.Conversations[i]) { // not required
			continue
		}

		if m.Conversations[i] != nil {
			if err := m.Conversations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conversations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conversations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this question entity based on the context it is used
func (m *QuestionEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConversations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuestionEntity) contextValidateConversations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Conversations); i++ {

		if m.Conversations[i] != nil {
			if err := m.Conversations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conversations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("conversations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuestionEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuestionEntity) UnmarshalBinary(b []byte) error {
	var res QuestionEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
