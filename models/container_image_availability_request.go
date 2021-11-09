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

// ContainerImageAvailabilityRequest container image availability request
//
// swagger:model container_image_availability_request
type ContainerImageAvailabilityRequest struct {

	// List of image names to be checked.
	// Required: true
	Images []string `json:"images"`

	// Positive number represents a timeout in seconds for a pull operation.
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this container image availability request
func (m *ContainerImageAvailabilityRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerImageAvailabilityRequest) validateImages(formats strfmt.Registry) error {

	if err := validate.Required("images", "body", m.Images); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this container image availability request based on context it is used
func (m *ContainerImageAvailabilityRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContainerImageAvailabilityRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerImageAvailabilityRequest) UnmarshalBinary(b []byte) error {
	var res ContainerImageAvailabilityRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
