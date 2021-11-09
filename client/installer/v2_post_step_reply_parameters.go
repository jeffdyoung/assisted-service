// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// NewV2PostStepReplyParams creates a new V2PostStepReplyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2PostStepReplyParams() *V2PostStepReplyParams {
	return &V2PostStepReplyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2PostStepReplyParamsWithTimeout creates a new V2PostStepReplyParams object
// with the ability to set a timeout on a request.
func NewV2PostStepReplyParamsWithTimeout(timeout time.Duration) *V2PostStepReplyParams {
	return &V2PostStepReplyParams{
		timeout: timeout,
	}
}

// NewV2PostStepReplyParamsWithContext creates a new V2PostStepReplyParams object
// with the ability to set a context for a request.
func NewV2PostStepReplyParamsWithContext(ctx context.Context) *V2PostStepReplyParams {
	return &V2PostStepReplyParams{
		Context: ctx,
	}
}

// NewV2PostStepReplyParamsWithHTTPClient creates a new V2PostStepReplyParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2PostStepReplyParamsWithHTTPClient(client *http.Client) *V2PostStepReplyParams {
	return &V2PostStepReplyParams{
		HTTPClient: client,
	}
}

/* V2PostStepReplyParams contains all the parameters to send to the API endpoint
   for the v2 post step reply operation.

   Typically these are written to a http.Request.
*/
type V2PostStepReplyParams struct {

	/* DiscoveryAgentVersion.

	   The software version of the discovery agent that is posting results.
	*/
	DiscoveryAgentVersion *string

	/* HostID.

	   The host that is posting results.

	   Format: uuid
	*/
	HostID strfmt.UUID

	/* InfraEnvID.

	   The infra env of the host that is posting results.

	   Format: uuid
	*/
	InfraEnvID strfmt.UUID

	/* Reply.

	   The results to be posted.
	*/
	Reply *models.StepReply

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 post step reply params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2PostStepReplyParams) WithDefaults() *V2PostStepReplyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 post step reply params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2PostStepReplyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 post step reply params
func (o *V2PostStepReplyParams) WithTimeout(timeout time.Duration) *V2PostStepReplyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 post step reply params
func (o *V2PostStepReplyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 post step reply params
func (o *V2PostStepReplyParams) WithContext(ctx context.Context) *V2PostStepReplyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 post step reply params
func (o *V2PostStepReplyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 post step reply params
func (o *V2PostStepReplyParams) WithHTTPClient(client *http.Client) *V2PostStepReplyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 post step reply params
func (o *V2PostStepReplyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDiscoveryAgentVersion adds the discoveryAgentVersion to the v2 post step reply params
func (o *V2PostStepReplyParams) WithDiscoveryAgentVersion(discoveryAgentVersion *string) *V2PostStepReplyParams {
	o.SetDiscoveryAgentVersion(discoveryAgentVersion)
	return o
}

// SetDiscoveryAgentVersion adds the discoveryAgentVersion to the v2 post step reply params
func (o *V2PostStepReplyParams) SetDiscoveryAgentVersion(discoveryAgentVersion *string) {
	o.DiscoveryAgentVersion = discoveryAgentVersion
}

// WithHostID adds the hostID to the v2 post step reply params
func (o *V2PostStepReplyParams) WithHostID(hostID strfmt.UUID) *V2PostStepReplyParams {
	o.SetHostID(hostID)
	return o
}

// SetHostID adds the hostId to the v2 post step reply params
func (o *V2PostStepReplyParams) SetHostID(hostID strfmt.UUID) {
	o.HostID = hostID
}

// WithInfraEnvID adds the infraEnvID to the v2 post step reply params
func (o *V2PostStepReplyParams) WithInfraEnvID(infraEnvID strfmt.UUID) *V2PostStepReplyParams {
	o.SetInfraEnvID(infraEnvID)
	return o
}

// SetInfraEnvID adds the infraEnvId to the v2 post step reply params
func (o *V2PostStepReplyParams) SetInfraEnvID(infraEnvID strfmt.UUID) {
	o.InfraEnvID = infraEnvID
}

// WithReply adds the reply to the v2 post step reply params
func (o *V2PostStepReplyParams) WithReply(reply *models.StepReply) *V2PostStepReplyParams {
	o.SetReply(reply)
	return o
}

// SetReply adds the reply to the v2 post step reply params
func (o *V2PostStepReplyParams) SetReply(reply *models.StepReply) {
	o.Reply = reply
}

// WriteToRequest writes these params to a swagger request
func (o *V2PostStepReplyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DiscoveryAgentVersion != nil {

		// header param discovery_agent_version
		if err := r.SetHeaderParam("discovery_agent_version", *o.DiscoveryAgentVersion); err != nil {
			return err
		}
	}

	// path param host_id
	if err := r.SetPathParam("host_id", o.HostID.String()); err != nil {
		return err
	}

	// path param infra_env_id
	if err := r.SetPathParam("infra_env_id", o.InfraEnvID.String()); err != nil {
		return err
	}
	if o.Reply != nil {
		if err := r.SetBodyParam(o.Reply); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
