// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TeamDTO team d t o
//
// swagger:model TeamDTO
type TeamDTO struct {

	// access control
	AccessControl map[string]bool `json:"accessControl,omitempty"`

	// avatar Url
	AvatarURL string `json:"avatarUrl,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// member count
	MemberCount int64 `json:"memberCount,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// org Id
	OrgID int64 `json:"orgId,omitempty"`

	// permission
	Permission PermissionType `json:"permission,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`
}

// Validate validates this team d t o
func (m *TeamDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamDTO) validatePermission(formats strfmt.Registry) error {
	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	if err := m.Permission.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("permission")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("permission")
		}
		return err
	}

	return nil
}

// ContextValidate validate this team d t o based on the context it is used
func (m *TeamDTO) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamDTO) contextValidatePermission(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Permission) { // not required
		return nil
	}

	if err := m.Permission.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("permission")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("permission")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamDTO) UnmarshalBinary(b []byte) error {
	var res TeamDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}