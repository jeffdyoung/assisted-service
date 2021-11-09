// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"github.com/openshift/assisted-service/models"
)

// NewV2RegisterHostParams creates a new V2RegisterHostParams object
//
// There are no default values defined in the spec.
func NewV2RegisterHostParams() V2RegisterHostParams {

	return V2RegisterHostParams{}
}

// V2RegisterHostParams contains all the bound params for the v2 register host operation
// typically these are obtained from a http.Request
//
// swagger:parameters v2RegisterHost
type V2RegisterHostParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The software version of the discovery agent that is registering the agent.
	  In: header
	*/
	DiscoveryAgentVersion *string
	/*The InfraEnv that the agent is associated with.
	  Required: true
	  In: path
	*/
	InfraEnvID strfmt.UUID
	/*The description of the agent being registered.
	  Required: true
	  In: body
	*/
	NewHostParams *models.HostCreateParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewV2RegisterHostParams() beforehand.
func (o *V2RegisterHostParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindDiscoveryAgentVersion(r.Header[http.CanonicalHeaderKey("discovery_agent_version")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rInfraEnvID, rhkInfraEnvID, _ := route.Params.GetOK("infra_env_id")
	if err := o.bindInfraEnvID(rInfraEnvID, rhkInfraEnvID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.HostCreateParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("newHostParams", "body", ""))
			} else {
				res = append(res, errors.NewParseError("newHostParams", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.NewHostParams = &body
			}
		}
	} else {
		res = append(res, errors.Required("newHostParams", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDiscoveryAgentVersion binds and validates parameter DiscoveryAgentVersion from header.
func (o *V2RegisterHostParams) bindDiscoveryAgentVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.DiscoveryAgentVersion = &raw

	return nil
}

// bindInfraEnvID binds and validates parameter InfraEnvID from path.
func (o *V2RegisterHostParams) bindInfraEnvID(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *V2RegisterHostParams) validateInfraEnvID(formats strfmt.Registry) error {

	if err := validate.FormatOf("infra_env_id", "path", "uuid", o.InfraEnvID.String(), formats); err != nil {
		return err
	}
	return nil
}
