// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/project-flotta/flotta-operator/backend/models"
)

// GetRegistrationStatusReader is a Reader for the GetRegistrationStatus structure.
type GetRegistrationStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegistrationStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegistrationStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRegistrationStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRegistrationStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRegistrationStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegistrationStatusOK creates a GetRegistrationStatusOK with default headers values
func NewGetRegistrationStatusOK() *GetRegistrationStatusOK {
	return &GetRegistrationStatusOK{}
}

/* GetRegistrationStatusOK describes a response with status code 200, with default header values.

OK
*/
type GetRegistrationStatusOK struct {
	Payload *models.DeviceRegistrationStatusResponse
}

func (o *GetRegistrationStatusOK) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/devices/{device-id}/registration][%d] getRegistrationStatusOK  %+v", 200, o.Payload)
}
func (o *GetRegistrationStatusOK) GetPayload() *models.DeviceRegistrationStatusResponse {
	return o.Payload
}

func (o *GetRegistrationStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceRegistrationStatusResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegistrationStatusUnauthorized creates a GetRegistrationStatusUnauthorized with default headers values
func NewGetRegistrationStatusUnauthorized() *GetRegistrationStatusUnauthorized {
	return &GetRegistrationStatusUnauthorized{}
}

/* GetRegistrationStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRegistrationStatusUnauthorized struct {
}

func (o *GetRegistrationStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/devices/{device-id}/registration][%d] getRegistrationStatusUnauthorized ", 401)
}

func (o *GetRegistrationStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistrationStatusForbidden creates a GetRegistrationStatusForbidden with default headers values
func NewGetRegistrationStatusForbidden() *GetRegistrationStatusForbidden {
	return &GetRegistrationStatusForbidden{}
}

/* GetRegistrationStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRegistrationStatusForbidden struct {
}

func (o *GetRegistrationStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/devices/{device-id}/registration][%d] getRegistrationStatusForbidden ", 403)
}

func (o *GetRegistrationStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegistrationStatusInternalServerError creates a GetRegistrationStatusInternalServerError with default headers values
func NewGetRegistrationStatusInternalServerError() *GetRegistrationStatusInternalServerError {
	return &GetRegistrationStatusInternalServerError{}
}

/* GetRegistrationStatusInternalServerError describes a response with status code 500, with default header values.

Error
*/
type GetRegistrationStatusInternalServerError struct {
}

func (o *GetRegistrationStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /namespaces/{namespace}/devices/{device-id}/registration][%d] getRegistrationStatusInternalServerError ", 500)
}

func (o *GetRegistrationStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
