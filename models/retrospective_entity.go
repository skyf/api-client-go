// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetrospectiveEntity Returns a report with retrospective analytics data
//
// swagger:model RetrospectiveEntity
type RetrospectiveEntity struct {

	// data
	Data string `json:"data,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`
}

// Validate validates this retrospective entity
func (m *RetrospectiveEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this retrospective entity based on context it is used
func (m *RetrospectiveEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RetrospectiveEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetrospectiveEntity) UnmarshalBinary(b []byte) error {
	var res RetrospectiveEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
