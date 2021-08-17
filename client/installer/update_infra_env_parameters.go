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

// NewUpdateInfraEnvParams creates a new UpdateInfraEnvParams object
// with the default values initialized.
func NewUpdateInfraEnvParams() *UpdateInfraEnvParams {
	var ()
	return &UpdateInfraEnvParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInfraEnvParamsWithTimeout creates a new UpdateInfraEnvParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateInfraEnvParamsWithTimeout(timeout time.Duration) *UpdateInfraEnvParams {
	var ()
	return &UpdateInfraEnvParams{

		timeout: timeout,
	}
}

// NewUpdateInfraEnvParamsWithContext creates a new UpdateInfraEnvParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateInfraEnvParamsWithContext(ctx context.Context) *UpdateInfraEnvParams {
	var ()
	return &UpdateInfraEnvParams{

		Context: ctx,
	}
}

// NewUpdateInfraEnvParamsWithHTTPClient creates a new UpdateInfraEnvParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateInfraEnvParamsWithHTTPClient(client *http.Client) *UpdateInfraEnvParams {
	var ()
	return &UpdateInfraEnvParams{
		HTTPClient: client,
	}
}

/*UpdateInfraEnvParams contains all the parameters to send to the API endpoint
for the update infra env operation typically these are written to a http.Request
*/
type UpdateInfraEnvParams struct {

	/*InfraEnvUpdateParams
	  The properties to update.

	*/
	InfraEnvUpdateParams *models.InfraEnvUpdateParams
	/*InfraEnvID
	  The InfraEnv to be updated.

	*/
	InfraEnvID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update infra env params
func (o *UpdateInfraEnvParams) WithTimeout(timeout time.Duration) *UpdateInfraEnvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update infra env params
func (o *UpdateInfraEnvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update infra env params
func (o *UpdateInfraEnvParams) WithContext(ctx context.Context) *UpdateInfraEnvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update infra env params
func (o *UpdateInfraEnvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update infra env params
func (o *UpdateInfraEnvParams) WithHTTPClient(client *http.Client) *UpdateInfraEnvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update infra env params
func (o *UpdateInfraEnvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfraEnvUpdateParams adds the infraEnvUpdateParams to the update infra env params
func (o *UpdateInfraEnvParams) WithInfraEnvUpdateParams(infraEnvUpdateParams *models.InfraEnvUpdateParams) *UpdateInfraEnvParams {
	o.SetInfraEnvUpdateParams(infraEnvUpdateParams)
	return o
}

// SetInfraEnvUpdateParams adds the infraEnvUpdateParams to the update infra env params
func (o *UpdateInfraEnvParams) SetInfraEnvUpdateParams(infraEnvUpdateParams *models.InfraEnvUpdateParams) {
	o.InfraEnvUpdateParams = infraEnvUpdateParams
}

// WithInfraEnvID adds the infraEnvID to the update infra env params
func (o *UpdateInfraEnvParams) WithInfraEnvID(infraEnvID strfmt.UUID) *UpdateInfraEnvParams {
	o.SetInfraEnvID(infraEnvID)
	return o
}

// SetInfraEnvID adds the infraEnvId to the update infra env params
func (o *UpdateInfraEnvParams) SetInfraEnvID(infraEnvID strfmt.UUID) {
	o.InfraEnvID = infraEnvID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInfraEnvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.InfraEnvUpdateParams != nil {
		if err := r.SetBodyParam(o.InfraEnvUpdateParams); err != nil {
			return err
		}
	}

	// path param infra_env_id
	if err := r.SetPathParam("infra_env_id", o.InfraEnvID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}