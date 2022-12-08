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

// PostV1ServiceDependencies Creates a service dependency relationship between two services
//
// swagger:model postV1ServiceDependencies
type PostV1ServiceDependencies struct {

	// connected service id
	// Required: true
	ConnectedServiceID *string `json:"connected_service_id"`

	// A note to describe the service dependency relationship
	Notes string `json:"notes,omitempty"`

	// service id
	// Required: true
	ServiceID *string `json:"service_id"`
}

// Validate validates this post v1 service dependencies
func (m *PostV1ServiceDependencies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectedServiceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1ServiceDependencies) validateConnectedServiceID(formats strfmt.Registry) error {

	if err := validate.Required("connected_service_id", "body", m.ConnectedServiceID); err != nil {
		return err
	}

	return nil
}

func (m *PostV1ServiceDependencies) validateServiceID(formats strfmt.Registry) error {

	if err := validate.Required("service_id", "body", m.ServiceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 service dependencies based on context it is used
func (m *PostV1ServiceDependencies) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1ServiceDependencies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1ServiceDependencies) UnmarshalBinary(b []byte) error {
	var res PostV1ServiceDependencies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
