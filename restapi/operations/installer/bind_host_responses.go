// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// BindHostOKCode is the HTTP code returned for type BindHostOK
const BindHostOKCode int = 200

/*BindHostOK Success.

swagger:response bindHostOK
*/
type BindHostOK struct {

	/*
	  In: Body
	*/
	Payload *models.Host `json:"body,omitempty"`
}

// NewBindHostOK creates BindHostOK with default headers values
func NewBindHostOK() *BindHostOK {

	return &BindHostOK{}
}

// WithPayload adds the payload to the bind host o k response
func (o *BindHostOK) WithPayload(payload *models.Host) *BindHostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bind host o k response
func (o *BindHostOK) SetPayload(payload *models.Host) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BindHostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BindHostBadRequestCode is the HTTP code returned for type BindHostBadRequest
const BindHostBadRequestCode int = 400

/*BindHostBadRequest Error.

swagger:response bindHostBadRequest
*/
type BindHostBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBindHostBadRequest creates BindHostBadRequest with default headers values
func NewBindHostBadRequest() *BindHostBadRequest {

	return &BindHostBadRequest{}
}

// WithPayload adds the payload to the bind host bad request response
func (o *BindHostBadRequest) WithPayload(payload *models.Error) *BindHostBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bind host bad request response
func (o *BindHostBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BindHostBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BindHostUnauthorizedCode is the HTTP code returned for type BindHostUnauthorized
const BindHostUnauthorizedCode int = 401

/*BindHostUnauthorized Unauthorized.

swagger:response bindHostUnauthorized
*/
type BindHostUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewBindHostUnauthorized creates BindHostUnauthorized with default headers values
func NewBindHostUnauthorized() *BindHostUnauthorized {

	return &BindHostUnauthorized{}
}

// WithPayload adds the payload to the bind host unauthorized response
func (o *BindHostUnauthorized) WithPayload(payload *models.InfraError) *BindHostUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bind host unauthorized response
func (o *BindHostUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BindHostUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BindHostForbiddenCode is the HTTP code returned for type BindHostForbidden
const BindHostForbiddenCode int = 403

/*BindHostForbidden Forbidden.

swagger:response bindHostForbidden
*/
type BindHostForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewBindHostForbidden creates BindHostForbidden with default headers values
func NewBindHostForbidden() *BindHostForbidden {

	return &BindHostForbidden{}
}

// WithPayload adds the payload to the bind host forbidden response
func (o *BindHostForbidden) WithPayload(payload *models.InfraError) *BindHostForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bind host forbidden response
func (o *BindHostForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BindHostForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BindHostNotFoundCode is the HTTP code returned for type BindHostNotFound
const BindHostNotFoundCode int = 404

/*BindHostNotFound Error.

swagger:response bindHostNotFound
*/
type BindHostNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBindHostNotFound creates BindHostNotFound with default headers values
func NewBindHostNotFound() *BindHostNotFound {

	return &BindHostNotFound{}
}

// WithPayload adds the payload to the bind host not found response
func (o *BindHostNotFound) WithPayload(payload *models.Error) *BindHostNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bind host not found response
func (o *BindHostNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BindHostNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BindHostMethodNotAllowedCode is the HTTP code returned for type BindHostMethodNotAllowed
const BindHostMethodNotAllowedCode int = 405

/*BindHostMethodNotAllowed Method Not Allowed.

swagger:response bindHostMethodNotAllowed
*/
type BindHostMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBindHostMethodNotAllowed creates BindHostMethodNotAllowed with default headers values
func NewBindHostMethodNotAllowed() *BindHostMethodNotAllowed {

	return &BindHostMethodNotAllowed{}
}

// WithPayload adds the payload to the bind host method not allowed response
func (o *BindHostMethodNotAllowed) WithPayload(payload *models.Error) *BindHostMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bind host method not allowed response
func (o *BindHostMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BindHostMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BindHostInternalServerErrorCode is the HTTP code returned for type BindHostInternalServerError
const BindHostInternalServerErrorCode int = 500

/*BindHostInternalServerError Error.

swagger:response bindHostInternalServerError
*/
type BindHostInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBindHostInternalServerError creates BindHostInternalServerError with default headers values
func NewBindHostInternalServerError() *BindHostInternalServerError {

	return &BindHostInternalServerError{}
}

// WithPayload adds the payload to the bind host internal server error response
func (o *BindHostInternalServerError) WithPayload(payload *models.Error) *BindHostInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bind host internal server error response
func (o *BindHostInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BindHostInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BindHostNotImplementedCode is the HTTP code returned for type BindHostNotImplemented
const BindHostNotImplementedCode int = 501

/*BindHostNotImplemented Not implemented.

swagger:response bindHostNotImplemented
*/
type BindHostNotImplemented struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBindHostNotImplemented creates BindHostNotImplemented with default headers values
func NewBindHostNotImplemented() *BindHostNotImplemented {

	return &BindHostNotImplemented{}
}

// WithPayload adds the payload to the bind host not implemented response
func (o *BindHostNotImplemented) WithPayload(payload *models.Error) *BindHostNotImplemented {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bind host not implemented response
func (o *BindHostNotImplemented) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BindHostNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BindHostServiceUnavailableCode is the HTTP code returned for type BindHostServiceUnavailable
const BindHostServiceUnavailableCode int = 503

/*BindHostServiceUnavailable Unavailable.

swagger:response bindHostServiceUnavailable
*/
type BindHostServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBindHostServiceUnavailable creates BindHostServiceUnavailable with default headers values
func NewBindHostServiceUnavailable() *BindHostServiceUnavailable {

	return &BindHostServiceUnavailable{}
}

// WithPayload adds the payload to the bind host service unavailable response
func (o *BindHostServiceUnavailable) WithPayload(payload *models.Error) *BindHostServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the bind host service unavailable response
func (o *BindHostServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BindHostServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
