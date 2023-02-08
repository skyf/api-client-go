// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1CommentingComments Creates a comment for a project
//
// swagger:model postV1CommentingComments
type PostV1CommentingComments struct {

	// Text body of comment
	// Required: true
	Body *string `json:"body"`

	// Object that the comment relates to
	// Required: true
	RelatedTo *string `json:"related_to"`
}

// Validate validates this post v1 commenting comments
func (m *PostV1CommentingComments) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelatedTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1CommentingComments) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

func (m *PostV1CommentingComments) validateRelatedTo(formats strfmt.Registry) error {

	if err := validate.Required("related_to", "body", m.RelatedTo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 commenting comments based on context it is used
func (m *PostV1CommentingComments) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1CommentingComments) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1CommentingComments) UnmarshalBinary(b []byte) error {
	var res PostV1CommentingComments
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
