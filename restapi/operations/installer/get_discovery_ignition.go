// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetDiscoveryIgnitionHandlerFunc turns a function with the right signature into a get discovery ignition handler
type GetDiscoveryIgnitionHandlerFunc func(GetDiscoveryIgnitionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDiscoveryIgnitionHandlerFunc) Handle(params GetDiscoveryIgnitionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetDiscoveryIgnitionHandler interface for that can handle valid get discovery ignition params
type GetDiscoveryIgnitionHandler interface {
	Handle(GetDiscoveryIgnitionParams, interface{}) middleware.Responder
}

// NewGetDiscoveryIgnition creates a new http.Handler for the get discovery ignition operation
func NewGetDiscoveryIgnition(ctx *middleware.Context, handler GetDiscoveryIgnitionHandler) *GetDiscoveryIgnition {
	return &GetDiscoveryIgnition{Context: ctx, Handler: handler}
}

/* GetDiscoveryIgnition swagger:route GET /v1/clusters/{cluster_id}/discovery-ignition installer getDiscoveryIgnition

Get the discovery ignition for the cluster based on its attributes and overridden ignition value before generating the discovery ISO.
Used to test the validity of the discovery ignition when it is being overridden.
For downloading the generated discovery ignition use /clusters/$CLUSTER_ID/downloads/files?file_name=discovery.ign


*/
type GetDiscoveryIgnition struct {
	Context *middleware.Context
	Handler GetDiscoveryIgnitionHandler
}

func (o *GetDiscoveryIgnition) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetDiscoveryIgnitionParams()
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
