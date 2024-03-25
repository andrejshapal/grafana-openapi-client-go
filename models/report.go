// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Report ConfigDTO is model representation in transfer
//
// swagger:model Report
type Report struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// dashboards
	Dashboards []*ReportDashboard `json:"dashboards"`

	// enable Csv
	EnableCSV bool `json:"enableCsv,omitempty"`

	// enable dashboard Url
	EnableDashboardURL bool `json:"enableDashboardUrl,omitempty"`

	// formats
	Formats []Type `json:"formats"`

	// id
	ID int64 `json:"id,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// options
	Options *ReportOptions `json:"options,omitempty"`

	// org Id
	OrgID int64 `json:"orgId,omitempty"`

	// recipients
	Recipients string `json:"recipients,omitempty"`

	// reply to
	ReplyTo string `json:"replyTo,omitempty"`

	// scale factor
	ScaleFactor int64 `json:"scaleFactor,omitempty"`

	// schedule
	Schedule *ReportSchedule `json:"schedule,omitempty"`

	// state
	State State `json:"state,omitempty"`

	// uid
	UID string `json:"uid,omitempty"`

	// updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// user Id
	UserID int64 `json:"userId,omitempty"`
}

// Validate validates this report
func (m *Report) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDashboards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Report) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Report) validateDashboards(formats strfmt.Registry) error {
	if swag.IsZero(m.Dashboards) { // not required
		return nil
	}

	for i := 0; i < len(m.Dashboards); i++ {
		if swag.IsZero(m.Dashboards[i]) { // not required
			continue
		}

		if m.Dashboards[i] != nil {
			if err := m.Dashboards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dashboards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dashboards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Report) validateFormats(formats strfmt.Registry) error {
	if swag.IsZero(m.Formats) { // not required
		return nil
	}

	for i := 0; i < len(m.Formats); i++ {

		if err := m.Formats[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("formats" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("formats" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Report) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Options) { // not required
		return nil
	}

	if m.Options != nil {
		if err := m.Options.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

func (m *Report) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *Report) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

func (m *Report) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this report based on the context it is used
func (m *Report) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDashboards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFormats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Report) contextValidateDashboards(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Dashboards); i++ {

		if m.Dashboards[i] != nil {

			if swag.IsZero(m.Dashboards[i]) { // not required
				return nil
			}

			if err := m.Dashboards[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dashboards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dashboards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Report) contextValidateFormats(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Formats); i++ {

		if swag.IsZero(m.Formats[i]) { // not required
			return nil
		}

		if err := m.Formats[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("formats" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("formats" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *Report) contextValidateOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.Options != nil {

		if swag.IsZero(m.Options) { // not required
			return nil
		}

		if err := m.Options.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("options")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("options")
			}
			return err
		}
	}

	return nil
}

func (m *Report) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {

		if swag.IsZero(m.Schedule) { // not required
			return nil
		}

		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *Report) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Report) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Report) UnmarshalBinary(b []byte) error {
	var res Report
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
