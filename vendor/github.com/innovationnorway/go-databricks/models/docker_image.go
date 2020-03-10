// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DockerImage docker image
// swagger:model DockerImage
type DockerImage struct {

	// basic auth
	BasicAuth *DockerBasicAuth `json:"basic_auth,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this docker image
func (m *DockerImage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBasicAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DockerImage) validateBasicAuth(formats strfmt.Registry) error {

	if swag.IsZero(m.BasicAuth) { // not required
		return nil
	}

	if m.BasicAuth != nil {
		if err := m.BasicAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basic_auth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DockerImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DockerImage) UnmarshalBinary(b []byte) error {
	var res DockerImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
