// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutV1IncidentsIncidentIDLinks Updates the external incident link attributes
//
// swagger:model putV1IncidentsIncidentIdLinks
type PutV1IncidentsIncidentIDLinks struct {

	// display text
	DisplayText string `json:"display_text,omitempty"`

	// href url
	HrefURL string `json:"href_url,omitempty"`

	// icon url
	IconURL string `json:"icon_url,omitempty"`
}

// Validate validates this put v1 incidents incident Id links
func (m *PutV1IncidentsIncidentIDLinks) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put v1 incidents incident Id links based on context it is used
func (m *PutV1IncidentsIncidentIDLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutV1IncidentsIncidentIDLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutV1IncidentsIncidentIDLinks) UnmarshalBinary(b []byte) error {
	var res PutV1IncidentsIncidentIDLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
