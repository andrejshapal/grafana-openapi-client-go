// Code generated by go-swagger; DO NOT EDIT.

package admin_ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// PostSyncUserWithLDAPReader is a Reader for the PostSyncUserWithLDAP structure.
type PostSyncUserWithLDAPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSyncUserWithLDAPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSyncUserWithLDAPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostSyncUserWithLDAPUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostSyncUserWithLDAPForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSyncUserWithLDAPInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /admin/ldap/sync/{user_id}] postSyncUserWithLDAP", response, response.Code())
	}
}

// NewPostSyncUserWithLDAPOK creates a PostSyncUserWithLDAPOK with default headers values
func NewPostSyncUserWithLDAPOK() *PostSyncUserWithLDAPOK {
	return &PostSyncUserWithLDAPOK{}
}

/*
PostSyncUserWithLDAPOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type PostSyncUserWithLDAPOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this post sync user with Ldap Ok response has a 2xx status code
func (o *PostSyncUserWithLDAPOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post sync user with Ldap Ok response has a 3xx status code
func (o *PostSyncUserWithLDAPOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post sync user with Ldap Ok response has a 4xx status code
func (o *PostSyncUserWithLDAPOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post sync user with Ldap Ok response has a 5xx status code
func (o *PostSyncUserWithLDAPOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post sync user with Ldap Ok response a status code equal to that given
func (o *PostSyncUserWithLDAPOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post sync user with Ldap Ok response
func (o *PostSyncUserWithLDAPOK) Code() int {
	return 200
}

func (o *PostSyncUserWithLDAPOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/ldap/sync/{user_id}][%d] postSyncUserWithLdapOk %s", 200, payload)
}

func (o *PostSyncUserWithLDAPOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/ldap/sync/{user_id}][%d] postSyncUserWithLdapOk %s", 200, payload)
}

func (o *PostSyncUserWithLDAPOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *PostSyncUserWithLDAPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSyncUserWithLDAPUnauthorized creates a PostSyncUserWithLDAPUnauthorized with default headers values
func NewPostSyncUserWithLDAPUnauthorized() *PostSyncUserWithLDAPUnauthorized {
	return &PostSyncUserWithLDAPUnauthorized{}
}

/*
PostSyncUserWithLDAPUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type PostSyncUserWithLDAPUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this post sync user with Ldap unauthorized response has a 2xx status code
func (o *PostSyncUserWithLDAPUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post sync user with Ldap unauthorized response has a 3xx status code
func (o *PostSyncUserWithLDAPUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post sync user with Ldap unauthorized response has a 4xx status code
func (o *PostSyncUserWithLDAPUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post sync user with Ldap unauthorized response has a 5xx status code
func (o *PostSyncUserWithLDAPUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post sync user with Ldap unauthorized response a status code equal to that given
func (o *PostSyncUserWithLDAPUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post sync user with Ldap unauthorized response
func (o *PostSyncUserWithLDAPUnauthorized) Code() int {
	return 401
}

func (o *PostSyncUserWithLDAPUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/ldap/sync/{user_id}][%d] postSyncUserWithLdapUnauthorized %s", 401, payload)
}

func (o *PostSyncUserWithLDAPUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/ldap/sync/{user_id}][%d] postSyncUserWithLdapUnauthorized %s", 401, payload)
}

func (o *PostSyncUserWithLDAPUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PostSyncUserWithLDAPUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSyncUserWithLDAPForbidden creates a PostSyncUserWithLDAPForbidden with default headers values
func NewPostSyncUserWithLDAPForbidden() *PostSyncUserWithLDAPForbidden {
	return &PostSyncUserWithLDAPForbidden{}
}

/*
PostSyncUserWithLDAPForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type PostSyncUserWithLDAPForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this post sync user with Ldap forbidden response has a 2xx status code
func (o *PostSyncUserWithLDAPForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post sync user with Ldap forbidden response has a 3xx status code
func (o *PostSyncUserWithLDAPForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post sync user with Ldap forbidden response has a 4xx status code
func (o *PostSyncUserWithLDAPForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post sync user with Ldap forbidden response has a 5xx status code
func (o *PostSyncUserWithLDAPForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post sync user with Ldap forbidden response a status code equal to that given
func (o *PostSyncUserWithLDAPForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post sync user with Ldap forbidden response
func (o *PostSyncUserWithLDAPForbidden) Code() int {
	return 403
}

func (o *PostSyncUserWithLDAPForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/ldap/sync/{user_id}][%d] postSyncUserWithLdapForbidden %s", 403, payload)
}

func (o *PostSyncUserWithLDAPForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/ldap/sync/{user_id}][%d] postSyncUserWithLdapForbidden %s", 403, payload)
}

func (o *PostSyncUserWithLDAPForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PostSyncUserWithLDAPForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSyncUserWithLDAPInternalServerError creates a PostSyncUserWithLDAPInternalServerError with default headers values
func NewPostSyncUserWithLDAPInternalServerError() *PostSyncUserWithLDAPInternalServerError {
	return &PostSyncUserWithLDAPInternalServerError{}
}

/*
PostSyncUserWithLDAPInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type PostSyncUserWithLDAPInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this post sync user with Ldap internal server error response has a 2xx status code
func (o *PostSyncUserWithLDAPInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post sync user with Ldap internal server error response has a 3xx status code
func (o *PostSyncUserWithLDAPInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post sync user with Ldap internal server error response has a 4xx status code
func (o *PostSyncUserWithLDAPInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post sync user with Ldap internal server error response has a 5xx status code
func (o *PostSyncUserWithLDAPInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post sync user with Ldap internal server error response a status code equal to that given
func (o *PostSyncUserWithLDAPInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post sync user with Ldap internal server error response
func (o *PostSyncUserWithLDAPInternalServerError) Code() int {
	return 500
}

func (o *PostSyncUserWithLDAPInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/ldap/sync/{user_id}][%d] postSyncUserWithLdapInternalServerError %s", 500, payload)
}

func (o *PostSyncUserWithLDAPInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /admin/ldap/sync/{user_id}][%d] postSyncUserWithLdapInternalServerError %s", 500, payload)
}

func (o *PostSyncUserWithLDAPInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *PostSyncUserWithLDAPInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
