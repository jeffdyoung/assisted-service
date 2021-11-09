// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// UploadLogsNoContentCode is the HTTP code returned for type UploadLogsNoContent
const UploadLogsNoContentCode int = 204

/*UploadLogsNoContent Success.

swagger:response uploadLogsNoContent
*/
type UploadLogsNoContent struct {
}

// NewUploadLogsNoContent creates UploadLogsNoContent with default headers values
func NewUploadLogsNoContent() *UploadLogsNoContent {

	return &UploadLogsNoContent{}
}

// WriteResponse to the client
func (o *UploadLogsNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// UploadLogsBadRequestCode is the HTTP code returned for type UploadLogsBadRequest
const UploadLogsBadRequestCode int = 400

/*UploadLogsBadRequest Bad Request

swagger:response uploadLogsBadRequest
*/
type UploadLogsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewUploadLogsBadRequest creates UploadLogsBadRequest with default headers values
func NewUploadLogsBadRequest() *UploadLogsBadRequest {

	return &UploadLogsBadRequest{}
}

// WithPayload adds the payload to the upload logs bad request response
func (o *UploadLogsBadRequest) WithPayload(payload *models.InfraError) *UploadLogsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload logs bad request response
func (o *UploadLogsBadRequest) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadLogsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UploadLogsUnauthorizedCode is the HTTP code returned for type UploadLogsUnauthorized
const UploadLogsUnauthorizedCode int = 401

/*UploadLogsUnauthorized Unauthorized.

swagger:response uploadLogsUnauthorized
*/
type UploadLogsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewUploadLogsUnauthorized creates UploadLogsUnauthorized with default headers values
func NewUploadLogsUnauthorized() *UploadLogsUnauthorized {

	return &UploadLogsUnauthorized{}
}

// WithPayload adds the payload to the upload logs unauthorized response
func (o *UploadLogsUnauthorized) WithPayload(payload *models.InfraError) *UploadLogsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload logs unauthorized response
func (o *UploadLogsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadLogsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UploadLogsForbiddenCode is the HTTP code returned for type UploadLogsForbidden
const UploadLogsForbiddenCode int = 403

/*UploadLogsForbidden Forbidden.

swagger:response uploadLogsForbidden
*/
type UploadLogsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewUploadLogsForbidden creates UploadLogsForbidden with default headers values
func NewUploadLogsForbidden() *UploadLogsForbidden {

	return &UploadLogsForbidden{}
}

// WithPayload adds the payload to the upload logs forbidden response
func (o *UploadLogsForbidden) WithPayload(payload *models.InfraError) *UploadLogsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload logs forbidden response
func (o *UploadLogsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadLogsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UploadLogsNotFoundCode is the HTTP code returned for type UploadLogsNotFound
const UploadLogsNotFoundCode int = 404

/*UploadLogsNotFound Error.

swagger:response uploadLogsNotFound
*/
type UploadLogsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUploadLogsNotFound creates UploadLogsNotFound with default headers values
func NewUploadLogsNotFound() *UploadLogsNotFound {

	return &UploadLogsNotFound{}
}

// WithPayload adds the payload to the upload logs not found response
func (o *UploadLogsNotFound) WithPayload(payload *models.Error) *UploadLogsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload logs not found response
func (o *UploadLogsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadLogsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UploadLogsInternalServerErrorCode is the HTTP code returned for type UploadLogsInternalServerError
const UploadLogsInternalServerErrorCode int = 500

/*UploadLogsInternalServerError Error.

swagger:response uploadLogsInternalServerError
*/
type UploadLogsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUploadLogsInternalServerError creates UploadLogsInternalServerError with default headers values
func NewUploadLogsInternalServerError() *UploadLogsInternalServerError {

	return &UploadLogsInternalServerError{}
}

// WithPayload adds the payload to the upload logs internal server error response
func (o *UploadLogsInternalServerError) WithPayload(payload *models.Error) *UploadLogsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload logs internal server error response
func (o *UploadLogsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadLogsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UploadLogsServiceUnavailableCode is the HTTP code returned for type UploadLogsServiceUnavailable
const UploadLogsServiceUnavailableCode int = 503

/*UploadLogsServiceUnavailable Unavailable.

swagger:response uploadLogsServiceUnavailable
*/
type UploadLogsServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUploadLogsServiceUnavailable creates UploadLogsServiceUnavailable with default headers values
func NewUploadLogsServiceUnavailable() *UploadLogsServiceUnavailable {

	return &UploadLogsServiceUnavailable{}
}

// WithPayload adds the payload to the upload logs service unavailable response
func (o *UploadLogsServiceUnavailable) WithPayload(payload *models.Error) *UploadLogsServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload logs service unavailable response
func (o *UploadLogsServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadLogsServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
