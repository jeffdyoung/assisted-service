// Code generated by go-swagger; DO NOT EDIT.

package manifests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// CreateClusterManifestReader is a Reader for the CreateClusterManifest structure.
type CreateClusterManifestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterManifestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateClusterManifestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateClusterManifestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateClusterManifestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateClusterManifestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateClusterManifestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCreateClusterManifestMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateClusterManifestConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateClusterManifestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateClusterManifestCreated creates a CreateClusterManifestCreated with default headers values
func NewCreateClusterManifestCreated() *CreateClusterManifestCreated {
	return &CreateClusterManifestCreated{}
}

/* CreateClusterManifestCreated describes a response with status code 201, with default header values.

Success.
*/
type CreateClusterManifestCreated struct {
	Payload *models.Manifest
}

func (o *CreateClusterManifestCreated) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/manifests][%d] createClusterManifestCreated  %+v", 201, o.Payload)
}
func (o *CreateClusterManifestCreated) GetPayload() *models.Manifest {
	return o.Payload
}

func (o *CreateClusterManifestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manifest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterManifestBadRequest creates a CreateClusterManifestBadRequest with default headers values
func NewCreateClusterManifestBadRequest() *CreateClusterManifestBadRequest {
	return &CreateClusterManifestBadRequest{}
}

/* CreateClusterManifestBadRequest describes a response with status code 400, with default header values.

Error.
*/
type CreateClusterManifestBadRequest struct {
	Payload *models.Error
}

func (o *CreateClusterManifestBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/manifests][%d] createClusterManifestBadRequest  %+v", 400, o.Payload)
}
func (o *CreateClusterManifestBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateClusterManifestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterManifestUnauthorized creates a CreateClusterManifestUnauthorized with default headers values
func NewCreateClusterManifestUnauthorized() *CreateClusterManifestUnauthorized {
	return &CreateClusterManifestUnauthorized{}
}

/* CreateClusterManifestUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type CreateClusterManifestUnauthorized struct {
	Payload *models.InfraError
}

func (o *CreateClusterManifestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/manifests][%d] createClusterManifestUnauthorized  %+v", 401, o.Payload)
}
func (o *CreateClusterManifestUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *CreateClusterManifestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterManifestForbidden creates a CreateClusterManifestForbidden with default headers values
func NewCreateClusterManifestForbidden() *CreateClusterManifestForbidden {
	return &CreateClusterManifestForbidden{}
}

/* CreateClusterManifestForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type CreateClusterManifestForbidden struct {
	Payload *models.InfraError
}

func (o *CreateClusterManifestForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/manifests][%d] createClusterManifestForbidden  %+v", 403, o.Payload)
}
func (o *CreateClusterManifestForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *CreateClusterManifestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterManifestNotFound creates a CreateClusterManifestNotFound with default headers values
func NewCreateClusterManifestNotFound() *CreateClusterManifestNotFound {
	return &CreateClusterManifestNotFound{}
}

/* CreateClusterManifestNotFound describes a response with status code 404, with default header values.

Error.
*/
type CreateClusterManifestNotFound struct {
	Payload *models.Error
}

func (o *CreateClusterManifestNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/manifests][%d] createClusterManifestNotFound  %+v", 404, o.Payload)
}
func (o *CreateClusterManifestNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateClusterManifestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterManifestMethodNotAllowed creates a CreateClusterManifestMethodNotAllowed with default headers values
func NewCreateClusterManifestMethodNotAllowed() *CreateClusterManifestMethodNotAllowed {
	return &CreateClusterManifestMethodNotAllowed{}
}

/* CreateClusterManifestMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type CreateClusterManifestMethodNotAllowed struct {
	Payload *models.Error
}

func (o *CreateClusterManifestMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/manifests][%d] createClusterManifestMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *CreateClusterManifestMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateClusterManifestMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterManifestConflict creates a CreateClusterManifestConflict with default headers values
func NewCreateClusterManifestConflict() *CreateClusterManifestConflict {
	return &CreateClusterManifestConflict{}
}

/* CreateClusterManifestConflict describes a response with status code 409, with default header values.

Error.
*/
type CreateClusterManifestConflict struct {
	Payload *models.Error
}

func (o *CreateClusterManifestConflict) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/manifests][%d] createClusterManifestConflict  %+v", 409, o.Payload)
}
func (o *CreateClusterManifestConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateClusterManifestConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterManifestInternalServerError creates a CreateClusterManifestInternalServerError with default headers values
func NewCreateClusterManifestInternalServerError() *CreateClusterManifestInternalServerError {
	return &CreateClusterManifestInternalServerError{}
}

/* CreateClusterManifestInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type CreateClusterManifestInternalServerError struct {
	Payload *models.Error
}

func (o *CreateClusterManifestInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/manifests][%d] createClusterManifestInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateClusterManifestInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateClusterManifestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
