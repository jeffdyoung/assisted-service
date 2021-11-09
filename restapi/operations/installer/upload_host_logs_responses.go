// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// UploadHostLogsNoContentCode is the HTTP code returned for type UploadHostLogsNoContent
const UploadHostLogsNoContentCode int = 204

/*UploadHostLogsNoContent Success.

swagger:response uploadHostLogsNoContent
*/
type UploadHostLogsNoContent struct {
}

// NewUploadHostLogsNoContent creates UploadHostLogsNoContent with default headers values
func NewUploadHostLogsNoContent() *UploadHostLogsNoContent {

	return &UploadHostLogsNoContent{}
}

// WriteResponse to the client
func (o *UploadHostLogsNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// UploadHostLogsBadRequestCode is the HTTP code returned for type UploadHostLogsBadRequest
const UploadHostLogsBadRequestCode int = 400

/*UploadHostLogsBadRequest Bad Request

swagger:response uploadHostLogsBadRequest
*/
type UploadHostLogsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewUploadHostLogsBadRequest creates UploadHostLogsBadRequest with default headers values
func NewUploadHostLogsBadRequest() *UploadHostLogsBadRequest {

	return &UploadHostLogsBadRequest{}
}

// WithPayload adds the payload to the upload host logs bad request response
func (o *UploadHostLogsBadRequest) WithPayload(payload *models.InfraError) *UploadHostLogsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload host logs bad request response
func (o *UploadHostLogsBadRequest) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadHostLogsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UploadHostLogsUnauthorizedCode is the HTTP code returned for type UploadHostLogsUnauthorized
const UploadHostLogsUnauthorizedCode int = 401

/*UploadHostLogsUnauthorized Unauthorized.

swagger:response uploadHostLogsUnauthorized
*/
type UploadHostLogsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewUploadHostLogsUnauthorized creates UploadHostLogsUnauthorized with default headers values
func NewUploadHostLogsUnauthorized() *UploadHostLogsUnauthorized {

	return &UploadHostLogsUnauthorized{}
}

// WithPayload adds the payload to the upload host logs unauthorized response
func (o *UploadHostLogsUnauthorized) WithPayload(payload *models.InfraError) *UploadHostLogsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload host logs unauthorized response
func (o *UploadHostLogsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadHostLogsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UploadHostLogsForbiddenCode is the HTTP code returned for type UploadHostLogsForbidden
const UploadHostLogsForbiddenCode int = 403

/*UploadHostLogsForbidden Forbidden.

swagger:response uploadHostLogsForbidden
*/
type UploadHostLogsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewUploadHostLogsForbidden creates UploadHostLogsForbidden with default headers values
func NewUploadHostLogsForbidden() *UploadHostLogsForbidden {

	return &UploadHostLogsForbidden{}
}

// WithPayload adds the payload to the upload host logs forbidden response
func (o *UploadHostLogsForbidden) WithPayload(payload *models.InfraError) *UploadHostLogsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload host logs forbidden response
func (o *UploadHostLogsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadHostLogsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UploadHostLogsNotFoundCode is the HTTP code returned for type UploadHostLogsNotFound
const UploadHostLogsNotFoundCode int = 404

/*UploadHostLogsNotFound Error.

swagger:response uploadHostLogsNotFound
*/
type UploadHostLogsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUploadHostLogsNotFound creates UploadHostLogsNotFound with default headers values
func NewUploadHostLogsNotFound() *UploadHostLogsNotFound {

	return &UploadHostLogsNotFound{}
}

// WithPayload adds the payload to the upload host logs not found response
func (o *UploadHostLogsNotFound) WithPayload(payload *models.Error) *UploadHostLogsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload host logs not found response
func (o *UploadHostLogsNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadHostLogsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UploadHostLogsInternalServerErrorCode is the HTTP code returned for type UploadHostLogsInternalServerError
const UploadHostLogsInternalServerErrorCode int = 500

/*UploadHostLogsInternalServerError Error.

swagger:response uploadHostLogsInternalServerError
*/
type UploadHostLogsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUploadHostLogsInternalServerError creates UploadHostLogsInternalServerError with default headers values
func NewUploadHostLogsInternalServerError() *UploadHostLogsInternalServerError {

	return &UploadHostLogsInternalServerError{}
}

// WithPayload adds the payload to the upload host logs internal server error response
func (o *UploadHostLogsInternalServerError) WithPayload(payload *models.Error) *UploadHostLogsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload host logs internal server error response
func (o *UploadHostLogsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadHostLogsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UploadHostLogsServiceUnavailableCode is the HTTP code returned for type UploadHostLogsServiceUnavailable
const UploadHostLogsServiceUnavailableCode int = 503

/*UploadHostLogsServiceUnavailable Unavailable.

swagger:response uploadHostLogsServiceUnavailable
*/
type UploadHostLogsServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUploadHostLogsServiceUnavailable creates UploadHostLogsServiceUnavailable with default headers values
func NewUploadHostLogsServiceUnavailable() *UploadHostLogsServiceUnavailable {

	return &UploadHostLogsServiceUnavailable{}
}

// WithPayload adds the payload to the upload host logs service unavailable response
func (o *UploadHostLogsServiceUnavailable) WithPayload(payload *models.Error) *UploadHostLogsServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload host logs service unavailable response
func (o *UploadHostLogsServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadHostLogsServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
