// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDeregisterInfraEnvParams creates a new DeregisterInfraEnvParams object
//
// There are no default values defined in the spec.
func NewDeregisterInfraEnvParams() DeregisterInfraEnvParams {

	return DeregisterInfraEnvParams{}
}

// DeregisterInfraEnvParams contains all the bound params for the deregister infra env operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeregisterInfraEnv
type DeregisterInfraEnvParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The InfraEnv to be deleted.
	  Required: true
	  In: path
	*/
	InfraEnvID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeregisterInfraEnvParams() beforehand.
func (o *DeregisterInfraEnvParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rInfraEnvID, rhkInfraEnvID, _ := route.Params.GetOK("infra_env_id")
	if err := o.bindInfraEnvID(rInfraEnvID, rhkInfraEnvID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindInfraEnvID binds and validates parameter InfraEnvID from path.
func (o *DeregisterInfraEnvParams) bindInfraEnvID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("infra_env_id", "path", "strfmt.UUID", raw)
	}
	o.InfraEnvID = *(value.(*strfmt.UUID))

	if err := o.validateInfraEnvID(formats); err != nil {
		return err
	}

	return nil
}

// validateInfraEnvID carries on validations for parameter InfraEnvID
func (o *DeregisterInfraEnvParams) validateInfraEnvID(formats strfmt.Registry) error {

	if err := validate.FormatOf("infra_env_id", "path", "uuid", o.InfraEnvID.String(), formats); err != nil {
		return err
	}
	return nil
}
