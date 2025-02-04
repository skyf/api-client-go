// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostV1Runbooks Create a new runbook for use with incidents.
//
// swagger:model postV1Runbooks
type PostV1Runbooks struct {

	// attachment rule
	AttachmentRule *PostV1RunbooksAttachmentRule `json:"attachment_rule,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// owner
	Owner *PostV1RunbooksOwner `json:"owner,omitempty"`

	// steps
	Steps []*PostV1RunbooksStepsItems0 `json:"steps"`

	// summary
	Summary string `json:"summary,omitempty"`

	// type
	// Enum: [incident general infrastructure incident_role]
	Type string `json:"type,omitempty"`
}

// Validate validates this post v1 runbooks
func (m *PostV1Runbooks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachmentRule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1Runbooks) validateAttachmentRule(formats strfmt.Registry) error {
	if swag.IsZero(m.AttachmentRule) { // not required
		return nil
	}

	if m.AttachmentRule != nil {
		if err := m.AttachmentRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachment_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attachment_rule")
			}
			return err
		}
	}

	return nil
}

func (m *PostV1Runbooks) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PostV1Runbooks) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *PostV1Runbooks) validateSteps(formats strfmt.Registry) error {
	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	for i := 0; i < len(m.Steps); i++ {
		if swag.IsZero(m.Steps[i]) { // not required
			continue
		}

		if m.Steps[i] != nil {
			if err := m.Steps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var postV1RunbooksTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["incident","general","infrastructure","incident_role"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postV1RunbooksTypeTypePropEnum = append(postV1RunbooksTypeTypePropEnum, v)
	}
}

const (

	// PostV1RunbooksTypeIncident captures enum value "incident"
	PostV1RunbooksTypeIncident string = "incident"

	// PostV1RunbooksTypeGeneral captures enum value "general"
	PostV1RunbooksTypeGeneral string = "general"

	// PostV1RunbooksTypeInfrastructure captures enum value "infrastructure"
	PostV1RunbooksTypeInfrastructure string = "infrastructure"

	// PostV1RunbooksTypeIncidentRole captures enum value "incident_role"
	PostV1RunbooksTypeIncidentRole string = "incident_role"
)

// prop value enum
func (m *PostV1Runbooks) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postV1RunbooksTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostV1Runbooks) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post v1 runbooks based on the context it is used
func (m *PostV1Runbooks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachmentRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSteps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1Runbooks) contextValidateAttachmentRule(ctx context.Context, formats strfmt.Registry) error {

	if m.AttachmentRule != nil {
		if err := m.AttachmentRule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachment_rule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attachment_rule")
			}
			return err
		}
	}

	return nil
}

func (m *PostV1Runbooks) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {
		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *PostV1Runbooks) contextValidateSteps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Steps); i++ {

		if m.Steps[i] != nil {
			if err := m.Steps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("steps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("steps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1Runbooks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1Runbooks) UnmarshalBinary(b []byte) error {
	var res PostV1Runbooks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1RunbooksAttachmentRule post v1 runbooks attachment rule
//
// swagger:model PostV1RunbooksAttachmentRule
type PostV1RunbooksAttachmentRule struct {

	// The JSON logic for the attaching the runbook
	// Required: true
	Logic *string `json:"logic"`

	// The user data for the rule
	UserData string `json:"user_data,omitempty"`
}

// Validate validates this post v1 runbooks attachment rule
func (m *PostV1RunbooksAttachmentRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1RunbooksAttachmentRule) validateLogic(formats strfmt.Registry) error {

	if err := validate.Required("attachment_rule"+"."+"logic", "body", m.Logic); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 runbooks attachment rule based on context it is used
func (m *PostV1RunbooksAttachmentRule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1RunbooksAttachmentRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1RunbooksAttachmentRule) UnmarshalBinary(b []byte) error {
	var res PostV1RunbooksAttachmentRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1RunbooksOwner An object representing a Team that owns the runbook
//
// swagger:model PostV1RunbooksOwner
type PostV1RunbooksOwner struct {

	// id
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this post v1 runbooks owner
func (m *PostV1RunbooksOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1RunbooksOwner) validateID(formats strfmt.Registry) error {

	if err := validate.Required("owner"+"."+"id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 runbooks owner based on context it is used
func (m *PostV1RunbooksOwner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1RunbooksOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1RunbooksOwner) UnmarshalBinary(b []byte) error {
	var res PostV1RunbooksOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1RunbooksStepsItems0 post v1 runbooks steps items0
//
// swagger:model PostV1RunbooksStepsItems0
type PostV1RunbooksStepsItems0 struct {

	// ID of action to use for this step.
	// Required: true
	ActionID *string `json:"action_id"`

	// Name for step
	// Required: true
	Name *string `json:"name"`

	// rule
	Rule []*PostV1RunbooksStepsItems0RuleItems0 `json:"rule"`
}

// Validate validates this post v1 runbooks steps items0
func (m *PostV1RunbooksStepsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1RunbooksStepsItems0) validateActionID(formats strfmt.Registry) error {

	if err := validate.Required("action_id", "body", m.ActionID); err != nil {
		return err
	}

	return nil
}

func (m *PostV1RunbooksStepsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PostV1RunbooksStepsItems0) validateRule(formats strfmt.Registry) error {
	if swag.IsZero(m.Rule) { // not required
		return nil
	}

	for i := 0; i < len(m.Rule); i++ {
		if swag.IsZero(m.Rule[i]) { // not required
			continue
		}

		if m.Rule[i] != nil {
			if err := m.Rule[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rule" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rule" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post v1 runbooks steps items0 based on the context it is used
func (m *PostV1RunbooksStepsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1RunbooksStepsItems0) contextValidateRule(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rule); i++ {

		if m.Rule[i] != nil {
			if err := m.Rule[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rule" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rule" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostV1RunbooksStepsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1RunbooksStepsItems0) UnmarshalBinary(b []byte) error {
	var res PostV1RunbooksStepsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PostV1RunbooksStepsItems0RuleItems0 post v1 runbooks steps items0 rule items0
//
// swagger:model PostV1RunbooksStepsItems0RuleItems0
type PostV1RunbooksStepsItems0RuleItems0 struct {

	// The JSON logic for the rule
	// Required: true
	Logic *string `json:"logic"`

	// The user data for the rule
	UserData string `json:"user_data,omitempty"`
}

// Validate validates this post v1 runbooks steps items0 rule items0
func (m *PostV1RunbooksStepsItems0RuleItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogic(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostV1RunbooksStepsItems0RuleItems0) validateLogic(formats strfmt.Registry) error {

	if err := validate.Required("logic", "body", m.Logic); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post v1 runbooks steps items0 rule items0 based on context it is used
func (m *PostV1RunbooksStepsItems0RuleItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostV1RunbooksStepsItems0RuleItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostV1RunbooksStepsItems0RuleItems0) UnmarshalBinary(b []byte) error {
	var res PostV1RunbooksStepsItems0RuleItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
