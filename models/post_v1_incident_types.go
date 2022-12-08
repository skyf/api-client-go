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

// PostV1IncidentTypes Create a new incident type
//
// swagger:model postV1IncidentTypes
type PostV1IncidentTypes struct {

	// name
	// Required: true
	Name *string `json:"name"`

	// template
	// Required: true
	Template interface{} `json:"template"`
}

// Validate validates this post v1 incident types
func (m *PostV1IncidentTypes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1IncidentTypes) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PostV1IncidentTypes) validateTemplate(formats strfmt.Registry) error {

	if m.Template == nil {
		return errors.Required("template", "body", nil)
	}

	return nil
}

// ContextValidate validates this post v1 incident types based on context it is used
func (m *PostV1IncidentTypes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1IncidentTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1IncidentTypes) UnmarshalBinary(b []byte) error {
	var res PostV1IncidentTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
