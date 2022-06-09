// Code generated by go-swagger; DO NOT EDIT.

package backend

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
)

// NewGetRegistrationStatusParams creates a new GetRegistrationStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRegistrationStatusParams() *GetRegistrationStatusParams {
	return &GetRegistrationStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegistrationStatusParamsWithTimeout creates a new GetRegistrationStatusParams object
// with the ability to set a timeout on a request.
func NewGetRegistrationStatusParamsWithTimeout(timeout time.Duration) *GetRegistrationStatusParams {
	return &GetRegistrationStatusParams{
		timeout: timeout,
	}
}

// NewGetRegistrationStatusParamsWithContext creates a new GetRegistrationStatusParams object
// with the ability to set a context for a request.
func NewGetRegistrationStatusParamsWithContext(ctx context.Context) *GetRegistrationStatusParams {
	return &GetRegistrationStatusParams{
		Context: ctx,
	}
}

// NewGetRegistrationStatusParamsWithHTTPClient creates a new GetRegistrationStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRegistrationStatusParamsWithHTTPClient(client *http.Client) *GetRegistrationStatusParams {
	return &GetRegistrationStatusParams{
		HTTPClient: client,
	}
}

/* GetRegistrationStatusParams contains all the parameters to send to the API endpoint
   for the get registration status operation.

   Typically these are written to a http.Request.
*/
type GetRegistrationStatusParams struct {

	/* DeviceID.

	   Device ID
	*/
	DeviceID string

	/* Namespace.

	   Namespace where the device resides
	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get registration status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistrationStatusParams) WithDefaults() *GetRegistrationStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get registration status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistrationStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get registration status params
func (o *GetRegistrationStatusParams) WithTimeout(timeout time.Duration) *GetRegistrationStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get registration status params
func (o *GetRegistrationStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get registration status params
func (o *GetRegistrationStatusParams) WithContext(ctx context.Context) *GetRegistrationStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get registration status params
func (o *GetRegistrationStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get registration status params
func (o *GetRegistrationStatusParams) WithHTTPClient(client *http.Client) *GetRegistrationStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get registration status params
func (o *GetRegistrationStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get registration status params
func (o *GetRegistrationStatusParams) WithDeviceID(deviceID string) *GetRegistrationStatusParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get registration status params
func (o *GetRegistrationStatusParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WithNamespace adds the namespace to the get registration status params
func (o *GetRegistrationStatusParams) WithNamespace(namespace string) *GetRegistrationStatusParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get registration status params
func (o *GetRegistrationStatusParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegistrationStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device-id
	if err := r.SetPathParam("device-id", o.DeviceID); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
