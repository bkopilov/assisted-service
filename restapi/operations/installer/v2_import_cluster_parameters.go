// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/openshift/assisted-service/models"
)

// NewV2ImportClusterParams creates a new V2ImportClusterParams object
// no default values defined in spec.
func NewV2ImportClusterParams() V2ImportClusterParams {

	return V2ImportClusterParams{}
}

// V2ImportClusterParams contains all the bound params for the v2 import cluster operation
// typically these are obtained from a http.Request
//
// swagger:parameters v2ImportCluster
type V2ImportClusterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Parameters for importing a OCP cluster for adding nodes.
	  Required: true
	  In: body
	*/
	NewImportClusterParams *models.ImportClusterParams
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewV2ImportClusterParams() beforehand.
func (o *V2ImportClusterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ImportClusterParams
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("newImportClusterParams", "body", ""))
			} else {
				res = append(res, errors.NewParseError("newImportClusterParams", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.NewImportClusterParams = &body
			}
		}
	} else {
		res = append(res, errors.Required("newImportClusterParams", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
