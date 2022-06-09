// Code generated by go-swagger; DO NOT EDIT.

package yggdrasil

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

// NewGetControlMessageForDeviceParams creates a new GetControlMessageForDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetControlMessageForDeviceParams() *GetControlMessageForDeviceParams {
	return &GetControlMessageForDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetControlMessageForDeviceParamsWithTimeout creates a new GetControlMessageForDeviceParams object
// with the ability to set a timeout on a request.
func NewGetControlMessageForDeviceParamsWithTimeout(timeout time.Duration) *GetControlMessageForDeviceParams {
	return &GetControlMessageForDeviceParams{
		timeout: timeout,
	}
}

// NewGetControlMessageForDeviceParamsWithContext creates a new GetControlMessageForDeviceParams object
// with the ability to set a context for a request.
func NewGetControlMessageForDeviceParamsWithContext(ctx context.Context) *GetControlMessageForDeviceParams {
	return &GetControlMessageForDeviceParams{
		Context: ctx,
	}
}

// NewGetControlMessageForDeviceParamsWithHTTPClient creates a new GetControlMessageForDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetControlMessageForDeviceParamsWithHTTPClient(client *http.Client) *GetControlMessageForDeviceParams {
	return &GetControlMessageForDeviceParams{
		HTTPClient: client,
	}
}

/* GetControlMessageForDeviceParams contains all the parameters to send to the API endpoint
   for the get control message for device operation.

   Typically these are written to a http.Request.
*/
type GetControlMessageForDeviceParams struct {

	/* DeviceID.

	   Device ID
	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get control message for device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetControlMessageForDeviceParams) WithDefaults() *GetControlMessageForDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get control message for device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetControlMessageForDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get control message for device params
func (o *GetControlMessageForDeviceParams) WithTimeout(timeout time.Duration) *GetControlMessageForDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get control message for device params
func (o *GetControlMessageForDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get control message for device params
func (o *GetControlMessageForDeviceParams) WithContext(ctx context.Context) *GetControlMessageForDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get control message for device params
func (o *GetControlMessageForDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get control message for device params
func (o *GetControlMessageForDeviceParams) WithHTTPClient(client *http.Client) *GetControlMessageForDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get control message for device params
func (o *GetControlMessageForDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get control message for device params
func (o *GetControlMessageForDeviceParams) WithDeviceID(deviceID string) *GetControlMessageForDeviceParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get control message for device params
func (o *GetControlMessageForDeviceParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetControlMessageForDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device_id
	if err := r.SetPathParam("device_id", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
