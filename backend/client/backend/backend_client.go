// Code generated by go-swagger; DO NOT EDIT.

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the backend client
type API interface {
	/*
	   EnrolDevice Initiates the process of enrolling the device*/
	EnrolDevice(ctx context.Context, params *EnrolDeviceParams) (*EnrolDeviceOK, *EnrolDeviceCreated, error)
	/*
	   GetDeviceConfiguration Returns the device configuration*/
	GetDeviceConfiguration(ctx context.Context, params *GetDeviceConfigurationParams) (*GetDeviceConfigurationOK, error)
	/*
	   GetRegistrationStatus Returns a device registration status, which can be registered, unregistered or unknown.*/
	GetRegistrationStatus(ctx context.Context, params *GetRegistrationStatusParams) (*GetRegistrationStatusOK, error)
	/*
	   RegisterDevice Registers the device by providing its hardware configuration*/
	RegisterDevice(ctx context.Context, params *RegisterDeviceParams) (*RegisterDeviceOK, error)
	/*
	   UpdateHeartBeat Updates the heartbeat information of the device.*/
	UpdateHeartBeat(ctx context.Context, params *UpdateHeartBeatParams) (*UpdateHeartBeatOK, error)
}

// New creates a new backend API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for backend API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
EnrolDevice Initiates the process of enrolling the device
*/
func (a *Client) EnrolDevice(ctx context.Context, params *EnrolDeviceParams) (*EnrolDeviceOK, *EnrolDeviceCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EnrolDevice",
		Method:             "POST",
		PathPattern:        "/namespaces/{namespace}/devices/{device-id}/enrolment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnrolDeviceReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *EnrolDeviceOK:
		return value, nil, nil
	case *EnrolDeviceCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GetDeviceConfiguration Returns the device configuration
*/
func (a *Client) GetDeviceConfiguration(ctx context.Context, params *GetDeviceConfigurationParams) (*GetDeviceConfigurationOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDeviceConfiguration",
		Method:             "GET",
		PathPattern:        "/namespaces/{namespace}/devices/{device-id}/configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceConfigurationReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeviceConfigurationOK), nil

}

/*
GetRegistrationStatus Returns a device registration status, which can be registered, unregistered or unknown.
*/
func (a *Client) GetRegistrationStatus(ctx context.Context, params *GetRegistrationStatusParams) (*GetRegistrationStatusOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRegistrationStatus",
		Method:             "GET",
		PathPattern:        "/namespaces/{namespace}/devices/{device-id}/registration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRegistrationStatusReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRegistrationStatusOK), nil

}

/*
RegisterDevice Registers the device by providing its hardware configuration
*/
func (a *Client) RegisterDevice(ctx context.Context, params *RegisterDeviceParams) (*RegisterDeviceOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegisterDevice",
		Method:             "PUT",
		PathPattern:        "/namespaces/{namespace}/devices/{device-id}/registration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RegisterDeviceReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegisterDeviceOK), nil

}

/*
UpdateHeartBeat Updates the heartbeat information of the device.
*/
func (a *Client) UpdateHeartBeat(ctx context.Context, params *UpdateHeartBeatParams) (*UpdateHeartBeatOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateHeartBeat",
		Method:             "PUT",
		PathPattern:        "/namespaces/{namespace}/devices/{device-id}/heartbeat",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateHeartBeatReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateHeartBeatOK), nil

}
