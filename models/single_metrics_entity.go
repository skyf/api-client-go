// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SingleMetricsEntity Return metrics for a specific infrastructure record
//
// swagger:model SingleMetricsEntity
type SingleMetricsEntity struct {

	// count
	Count string `json:"count,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// mtta
	Mtta string `json:"mtta,omitempty"`

	// mttd
	Mttd string `json:"mttd,omitempty"`

	// mttm
	Mttm string `json:"mttm,omitempty"`

	// mttr
	Mttr string `json:"mttr,omitempty"`

	// total time
	TotalTime string `json:"total_time,omitempty"`
}

// Validate validates this single metrics entity
func (m *SingleMetricsEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this single metrics entity based on context it is used
func (m *SingleMetricsEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SingleMetricsEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SingleMetricsEntity) UnmarshalBinary(b []byte) error {
	var res SingleMetricsEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
