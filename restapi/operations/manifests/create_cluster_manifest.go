// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateClusterManifestHandlerFunc turns a function with the right signature into a create cluster manifest handler
type CreateClusterManifestHandlerFunc func(CreateClusterManifestParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateClusterManifestHandlerFunc) Handle(params CreateClusterManifestParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateClusterManifestHandler interface for that can handle valid create cluster manifest params
type CreateClusterManifestHandler interface {
	Handle(CreateClusterManifestParams, interface{}) middleware.Responder
}

// NewCreateClusterManifest creates a new http.Handler for the create cluster manifest operation
func NewCreateClusterManifest(ctx *middleware.Context, handler CreateClusterManifestHandler) *CreateClusterManifest {
	return &CreateClusterManifest{Context: ctx, Handler: handler}
}

/* CreateClusterManifest swagger:route POST /v1/clusters/{cluster_id}/manifests manifests createClusterManifest

Creates a manifest for customizing cluster installation.

*/
type CreateClusterManifest struct {
	Context *middleware.Context
	Handler CreateClusterManifestHandler
}

func (o *CreateClusterManifest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateClusterManifestParams()
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
