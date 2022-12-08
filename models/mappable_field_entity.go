// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MappableFieldEntity Returns metadata for the fields that are available for field mapping.
//
// swagger:model MappableFieldEntity
type MappableFieldEntity struct {

	// The allowed values of the field
	AllowedValues string `json:"allowed_values,omitempty"`

	// The human-readable name of the field
	Label string `json:"label,omitempty"`

	// If the field is required to be mapped
	Required string `json:"required,omitempty"`

	// The allowed type of the field
	Type string `json:"type,omitempty"`

	// The ID of the field
	Value string `json:"value,omitempty"`
}

// Validate validates this mappable field entity
func (m *MappableFieldEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mappable field entity based on context it is used
func (m *MappableFieldEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MappableFieldEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MappableFieldEntity) UnmarshalBinary(b []byte) error {
	var res MappableFieldEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
