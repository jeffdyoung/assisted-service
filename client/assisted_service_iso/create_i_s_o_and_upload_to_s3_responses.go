// Code generated by go-swagger; DO NOT EDIT.

package assisted_service_iso

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// CreateISOAndUploadToS3Reader is a Reader for the CreateISOAndUploadToS3 structure.
type CreateISOAndUploadToS3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateISOAndUploadToS3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateISOAndUploadToS3Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateISOAndUploadToS3BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateISOAndUploadToS3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateISOAndUploadToS3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateISOAndUploadToS3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateISOAndUploadToS3Created creates a CreateISOAndUploadToS3Created with default headers values
func NewCreateISOAndUploadToS3Created() *CreateISOAndUploadToS3Created {
	return &CreateISOAndUploadToS3Created{}
}

/* CreateISOAndUploadToS3Created describes a response with status code 201, with default header values.

Success.
*/
type CreateISOAndUploadToS3Created struct {
}

func (o *CreateISOAndUploadToS3Created) Error() string {
	return fmt.Sprintf("[POST /v1/assisted-service-iso][%d] createISOAndUploadToS3Created ", 201)
}

func (o *CreateISOAndUploadToS3Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateISOAndUploadToS3BadRequest creates a CreateISOAndUploadToS3BadRequest with default headers values
func NewCreateISOAndUploadToS3BadRequest() *CreateISOAndUploadToS3BadRequest {
	return &CreateISOAndUploadToS3BadRequest{}
}

/* CreateISOAndUploadToS3BadRequest describes a response with status code 400, with default header values.

Error.
*/
type CreateISOAndUploadToS3BadRequest struct {
	Payload *models.Error
}

func (o *CreateISOAndUploadToS3BadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/assisted-service-iso][%d] createISOAndUploadToS3BadRequest  %+v", 400, o.Payload)
}
func (o *CreateISOAndUploadToS3BadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateISOAndUploadToS3BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateISOAndUploadToS3Unauthorized creates a CreateISOAndUploadToS3Unauthorized with default headers values
func NewCreateISOAndUploadToS3Unauthorized() *CreateISOAndUploadToS3Unauthorized {
	return &CreateISOAndUploadToS3Unauthorized{}
}

/* CreateISOAndUploadToS3Unauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type CreateISOAndUploadToS3Unauthorized struct {
	Payload *models.InfraError
}

func (o *CreateISOAndUploadToS3Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/assisted-service-iso][%d] createISOAndUploadToS3Unauthorized  %+v", 401, o.Payload)
}
func (o *CreateISOAndUploadToS3Unauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *CreateISOAndUploadToS3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateISOAndUploadToS3Forbidden creates a CreateISOAndUploadToS3Forbidden with default headers values
func NewCreateISOAndUploadToS3Forbidden() *CreateISOAndUploadToS3Forbidden {
	return &CreateISOAndUploadToS3Forbidden{}
}

/* CreateISOAndUploadToS3Forbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type CreateISOAndUploadToS3Forbidden struct {
	Payload *models.InfraError
}

func (o *CreateISOAndUploadToS3Forbidden) Error() string {
	return fmt.Sprintf("[POST /v1/assisted-service-iso][%d] createISOAndUploadToS3Forbidden  %+v", 403, o.Payload)
}
func (o *CreateISOAndUploadToS3Forbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *CreateISOAndUploadToS3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateISOAndUploadToS3InternalServerError creates a CreateISOAndUploadToS3InternalServerError with default headers values
func NewCreateISOAndUploadToS3InternalServerError() *CreateISOAndUploadToS3InternalServerError {
	return &CreateISOAndUploadToS3InternalServerError{}
}

/* CreateISOAndUploadToS3InternalServerError describes a response with status code 500, with default header values.

Error.
*/
type CreateISOAndUploadToS3InternalServerError struct {
	Payload *models.Error
}

func (o *CreateISOAndUploadToS3InternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/assisted-service-iso][%d] createISOAndUploadToS3InternalServerError  %+v", 500, o.Payload)
}
func (o *CreateISOAndUploadToS3InternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateISOAndUploadToS3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
