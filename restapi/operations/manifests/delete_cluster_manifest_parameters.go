// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDeleteClusterManifestParams creates a new DeleteClusterManifestParams object
// with the default values initialized.
func NewDeleteClusterManifestParams() DeleteClusterManifestParams {

	var (
		// initialize parameters with default values

		folderDefault = string("manifests")
	)

	return DeleteClusterManifestParams{
		Folder: &folderDefault,
	}
}

// DeleteClusterManifestParams contains all the bound params for the delete cluster manifest operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteClusterManifest
type DeleteClusterManifestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The cluster whose manifest should be deleted.
	  Required: true
	  In: path
	*/
	ClusterID strfmt.UUID
	/*The manifest file name to delete from the cluster.
	  Required: true
	  In: query
	*/
	FileName string
	/*The folder that contains the files. Manifests can be placed in 'manifests' or 'openshift' directories.
	  In: query
	  Default: "manifests"
	*/
	Folder *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteClusterManifestParams() beforehand.
func (o *DeleteClusterManifestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rClusterID, rhkClusterID, _ := route.Params.GetOK("cluster_id")
	if err := o.bindClusterID(rClusterID, rhkClusterID, route.Formats); err != nil {
		res = append(res, err)
	}

	qFileName, qhkFileName, _ := qs.GetOK("file_name")
	if err := o.bindFileName(qFileName, qhkFileName, route.Formats); err != nil {
		res = append(res, err)
	}

	qFolder, qhkFolder, _ := qs.GetOK("folder")
	if err := o.bindFolder(qFolder, qhkFolder, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClusterID binds and validates parameter ClusterID from path.
func (o *DeleteClusterManifestParams) bindClusterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("cluster_id", "path", "strfmt.UUID", raw)
	}
	o.ClusterID = *(value.(*strfmt.UUID))

	if err := o.validateClusterID(formats); err != nil {
		return err
	}

	return nil
}

// validateClusterID carries on validations for parameter ClusterID
func (o *DeleteClusterManifestParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.FormatOf("cluster_id", "path", "uuid", o.ClusterID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindFileName binds and validates parameter FileName from query.
func (o *DeleteClusterManifestParams) bindFileName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("file_name", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("file_name", "query", raw); err != nil {
		return err
	}
	o.FileName = raw

	return nil
}

// bindFolder binds and validates parameter Folder from query.
func (o *DeleteClusterManifestParams) bindFolder(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewDeleteClusterManifestParams()
		return nil
	}
	o.Folder = &raw

	if err := o.validateFolder(formats); err != nil {
		return err
	}

	return nil
}

// validateFolder carries on validations for parameter Folder
func (o *DeleteClusterManifestParams) validateFolder(formats strfmt.Registry) error {

	if err := validate.EnumCase("folder", "query", *o.Folder, []interface{}{"manifests", "openshift"}, true); err != nil {
		return err
	}

	return nil
}
