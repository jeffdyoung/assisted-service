// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateInfraEnvHandlerFunc turns a function with the right signature into a update infra env handler
type UpdateInfraEnvHandlerFunc func(UpdateInfraEnvParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateInfraEnvHandlerFunc) Handle(params UpdateInfraEnvParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// UpdateInfraEnvHandler interface for that can handle valid update infra env params
type UpdateInfraEnvHandler interface {
	Handle(UpdateInfraEnvParams, interface{}) middleware.Responder
}

// NewUpdateInfraEnv creates a new http.Handler for the update infra env operation
func NewUpdateInfraEnv(ctx *middleware.Context, handler UpdateInfraEnvHandler) *UpdateInfraEnv {
	return &UpdateInfraEnv{Context: ctx, Handler: handler}
}

/* UpdateInfraEnv swagger:route PATCH /v2/infra-envs/{infra_env_id} installer updateInfraEnv

Updates an InfraEnv.

*/
type UpdateInfraEnv struct {
	Context *middleware.Context
	Handler UpdateInfraEnvHandler
}

func (o *UpdateInfraEnv) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateInfraEnvParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
