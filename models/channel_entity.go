// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChannelEntity Gives chat channel information for the specified incident
//
// swagger:model ChannelEntity
type ChannelEntity struct {

	// icon url
	IconURL string `json:"icon_url,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// source id
	SourceID string `json:"source_id,omitempty"`

	// source name
	SourceName string `json:"source_name,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this channel entity
func (m *ChannelEntity) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this channel entity based on context it is used
func (m *ChannelEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChannelEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelEntity) UnmarshalBinary(b []byte) error {
	var res ChannelEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
