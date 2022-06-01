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

// LogsCollectionInformation Log collection information
//
// swagger:model logsCollectionInformation
type LogsCollectionInformation struct {

	// buffer size
	BufferSize int32 `json:"buffer_size,omitempty"`

	// kind
	Kind string `json:"kind,omitempty"`

	// syslog config
	SyslogConfig *LogsCollectionInformationSyslogConfig `json:"syslog_config,omitempty"`
}

// Validate validates this logs collection information
func (m *LogsCollectionInformation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSyslogConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogsCollectionInformation) validateSyslogConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.SyslogConfig) { // not required
		return nil
	}

	if m.SyslogConfig != nil {
		if err := m.SyslogConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syslog_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("syslog_config")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this logs collection information based on the context it is used
func (m *LogsCollectionInformation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSyslogConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LogsCollectionInformation) contextValidateSyslogConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.SyslogConfig != nil {
		if err := m.SyslogConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("syslog_config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("syslog_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogsCollectionInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogsCollectionInformation) UnmarshalBinary(b []byte) error {
	var res LogsCollectionInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LogsCollectionInformationSyslogConfig logs collection information syslog config
//
// swagger:model LogsCollectionInformationSyslogConfig
type LogsCollectionInformationSyslogConfig struct {

	// address
	Address string `json:"address,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`
}

// Validate validates this logs collection information syslog config
func (m *LogsCollectionInformationSyslogConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this logs collection information syslog config based on context it is used
func (m *LogsCollectionInformationSyslogConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LogsCollectionInformationSyslogConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogsCollectionInformationSyslogConfig) UnmarshalBinary(b []byte) error {
	var res LogsCollectionInformationSyslogConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
