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
	"github.com/go-openapi/swag"
)

// NewV2GetClusterParams creates a new V2GetClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2GetClusterParams() *V2GetClusterParams {
	return &V2GetClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2GetClusterParamsWithTimeout creates a new V2GetClusterParams object
// with the ability to set a timeout on a request.
func NewV2GetClusterParamsWithTimeout(timeout time.Duration) *V2GetClusterParams {
	return &V2GetClusterParams{
		timeout: timeout,
	}
}

// NewV2GetClusterParamsWithContext creates a new V2GetClusterParams object
// with the ability to set a context for a request.
func NewV2GetClusterParamsWithContext(ctx context.Context) *V2GetClusterParams {
	return &V2GetClusterParams{
		Context: ctx,
	}
}

// NewV2GetClusterParamsWithHTTPClient creates a new V2GetClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2GetClusterParamsWithHTTPClient(client *http.Client) *V2GetClusterParams {
	return &V2GetClusterParams{
		HTTPClient: client,
	}
}

/* V2GetClusterParams contains all the parameters to send to the API endpoint
   for the v2 get cluster operation.

   Typically these are written to a http.Request.
*/
type V2GetClusterParams struct {

	/* ClusterID.

	   The cluster to be retrieved.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	/* DiscoveryAgentVersion.

	   The software version of the discovery agent that is retrieving the cluster details.
	*/
	DiscoveryAgentVersion *string

	/* GetUnregisteredClusters.

	   Whether to return clusters that have been unregistered.
	*/
	GetUnregisteredClusters *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 get cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2GetClusterParams) WithDefaults() *V2GetClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 get cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2GetClusterParams) SetDefaults() {
	var (
		getUnregisteredClustersDefault = bool(false)
	)

	val := V2GetClusterParams{
		GetUnregisteredClusters: &getUnregisteredClustersDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the v2 get cluster params
func (o *V2GetClusterParams) WithTimeout(timeout time.Duration) *V2GetClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 get cluster params
func (o *V2GetClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 get cluster params
func (o *V2GetClusterParams) WithContext(ctx context.Context) *V2GetClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 get cluster params
func (o *V2GetClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 get cluster params
func (o *V2GetClusterParams) WithHTTPClient(client *http.Client) *V2GetClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 get cluster params
func (o *V2GetClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the v2 get cluster params
func (o *V2GetClusterParams) WithClusterID(clusterID strfmt.UUID) *V2GetClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the v2 get cluster params
func (o *V2GetClusterParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithDiscoveryAgentVersion adds the discoveryAgentVersion to the v2 get cluster params
func (o *V2GetClusterParams) WithDiscoveryAgentVersion(discoveryAgentVersion *string) *V2GetClusterParams {
	o.SetDiscoveryAgentVersion(discoveryAgentVersion)
	return o
}

// SetDiscoveryAgentVersion adds the discoveryAgentVersion to the v2 get cluster params
func (o *V2GetClusterParams) SetDiscoveryAgentVersion(discoveryAgentVersion *string) {
	o.DiscoveryAgentVersion = discoveryAgentVersion
}

// WithGetUnregisteredClusters adds the getUnregisteredClusters to the v2 get cluster params
func (o *V2GetClusterParams) WithGetUnregisteredClusters(getUnregisteredClusters *bool) *V2GetClusterParams {
	o.SetGetUnregisteredClusters(getUnregisteredClusters)
	return o
}

// SetGetUnregisteredClusters adds the getUnregisteredClusters to the v2 get cluster params
func (o *V2GetClusterParams) SetGetUnregisteredClusters(getUnregisteredClusters *bool) {
	o.GetUnregisteredClusters = getUnregisteredClusters
}

// WriteToRequest writes these params to a swagger request
func (o *V2GetClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if o.DiscoveryAgentVersion != nil {

		// header param discovery_agent_version
		if err := r.SetHeaderParam("discovery_agent_version", *o.DiscoveryAgentVersion); err != nil {
			return err
		}
	}

	if o.GetUnregisteredClusters != nil {

		// header param get_unregistered_clusters
		if err := r.SetHeaderParam("get_unregistered_clusters", swag.FormatBool(*o.GetUnregisteredClusters)); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
