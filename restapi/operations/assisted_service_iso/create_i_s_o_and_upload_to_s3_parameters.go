// Code generated by go-swagger; DO NOT EDIT.

package assisted_service_iso

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/openshift/assisted-service/models"
)

// NewCreateISOAndUploadToS3Params creates a new CreateISOAndUploadToS3Params object
//
// There are no default values defined in the spec.
func NewCreateISOAndUploadToS3Params() CreateISOAndUploadToS3Params {

	return CreateISOAndUploadToS3Params{}
}

// CreateISOAndUploadToS3Params contains all the bound params for the create i s o and upload to s3 operation
// typically these are obtained from a http.Request
//
// swagger:parameters CreateISOAndUploadToS3
type CreateISOAndUploadToS3Params struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Parameters for creating an Assisted Service ISO.
	  Required: true
	  In: body
	*/
	AssistedServiceIsoCreateParams *models.AssistedServiceIsoCreateParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateISOAndUploadToS3Params() beforehand.
func (o *CreateISOAndUploadToS3Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.AssistedServiceIsoCreateParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("assistedServiceIsoCreateParams", "body", ""))
			} else {
				res = append(res, errors.NewParseError("assistedServiceIsoCreateParams", "body", "", err))
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
				o.AssistedServiceIsoCreateParams = &body
			}
		}
	} else {
		res = append(res, errors.Required("assistedServiceIsoCreateParams", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
