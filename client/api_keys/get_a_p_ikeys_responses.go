// Code generated by go-swagger; DO NOT EDIT.

package api_keys

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

// GetAPIkeysReader is a Reader for the GetAPIkeys structure.
type GetAPIkeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIkeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIkeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAPIkeysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAPIkeysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAPIkeysNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAPIkeysInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /auth/keys] getAPIkeys", response, response.Code())
	}
}

// NewGetAPIkeysOK creates a GetAPIkeysOK with default headers values
func NewGetAPIkeysOK() *GetAPIkeysOK {
	return &GetAPIkeysOK{}
}

/*
GetAPIkeysOK describes a response with status code 200, with default header values.

(empty)
*/
type GetAPIkeysOK struct {
	Payload []*models.APIKeyDTO
}

// IsSuccess returns true when this get a p ikeys Ok response has a 2xx status code
func (o *GetAPIkeysOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get a p ikeys Ok response has a 3xx status code
func (o *GetAPIkeysOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get a p ikeys Ok response has a 4xx status code
func (o *GetAPIkeysOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get a p ikeys Ok response has a 5xx status code
func (o *GetAPIkeysOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get a p ikeys Ok response a status code equal to that given
func (o *GetAPIkeysOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get a p ikeys Ok response
func (o *GetAPIkeysOK) Code() int {
	return 200
}

func (o *GetAPIkeysOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysOk %s", 200, payload)
}

func (o *GetAPIkeysOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysOk %s", 200, payload)
}

func (o *GetAPIkeysOK) GetPayload() []*models.APIKeyDTO {
	return o.Payload
}

func (o *GetAPIkeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIkeysUnauthorized creates a GetAPIkeysUnauthorized with default headers values
func NewGetAPIkeysUnauthorized() *GetAPIkeysUnauthorized {
	return &GetAPIkeysUnauthorized{}
}

/*
GetAPIkeysUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetAPIkeysUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get a p ikeys unauthorized response has a 2xx status code
func (o *GetAPIkeysUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get a p ikeys unauthorized response has a 3xx status code
func (o *GetAPIkeysUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get a p ikeys unauthorized response has a 4xx status code
func (o *GetAPIkeysUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get a p ikeys unauthorized response has a 5xx status code
func (o *GetAPIkeysUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get a p ikeys unauthorized response a status code equal to that given
func (o *GetAPIkeysUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get a p ikeys unauthorized response
func (o *GetAPIkeysUnauthorized) Code() int {
	return 401
}

func (o *GetAPIkeysUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysUnauthorized %s", 401, payload)
}

func (o *GetAPIkeysUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysUnauthorized %s", 401, payload)
}

func (o *GetAPIkeysUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAPIkeysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIkeysForbidden creates a GetAPIkeysForbidden with default headers values
func NewGetAPIkeysForbidden() *GetAPIkeysForbidden {
	return &GetAPIkeysForbidden{}
}

/*
GetAPIkeysForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetAPIkeysForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get a p ikeys forbidden response has a 2xx status code
func (o *GetAPIkeysForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get a p ikeys forbidden response has a 3xx status code
func (o *GetAPIkeysForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get a p ikeys forbidden response has a 4xx status code
func (o *GetAPIkeysForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get a p ikeys forbidden response has a 5xx status code
func (o *GetAPIkeysForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get a p ikeys forbidden response a status code equal to that given
func (o *GetAPIkeysForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get a p ikeys forbidden response
func (o *GetAPIkeysForbidden) Code() int {
	return 403
}

func (o *GetAPIkeysForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysForbidden %s", 403, payload)
}

func (o *GetAPIkeysForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysForbidden %s", 403, payload)
}

func (o *GetAPIkeysForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAPIkeysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIkeysNotFound creates a GetAPIkeysNotFound with default headers values
func NewGetAPIkeysNotFound() *GetAPIkeysNotFound {
	return &GetAPIkeysNotFound{}
}

/*
GetAPIkeysNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetAPIkeysNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get a p ikeys not found response has a 2xx status code
func (o *GetAPIkeysNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get a p ikeys not found response has a 3xx status code
func (o *GetAPIkeysNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get a p ikeys not found response has a 4xx status code
func (o *GetAPIkeysNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get a p ikeys not found response has a 5xx status code
func (o *GetAPIkeysNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get a p ikeys not found response a status code equal to that given
func (o *GetAPIkeysNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get a p ikeys not found response
func (o *GetAPIkeysNotFound) Code() int {
	return 404
}

func (o *GetAPIkeysNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysNotFound %s", 404, payload)
}

func (o *GetAPIkeysNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysNotFound %s", 404, payload)
}

func (o *GetAPIkeysNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAPIkeysNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIkeysInternalServerError creates a GetAPIkeysInternalServerError with default headers values
func NewGetAPIkeysInternalServerError() *GetAPIkeysInternalServerError {
	return &GetAPIkeysInternalServerError{}
}

/*
GetAPIkeysInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetAPIkeysInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get a p ikeys internal server error response has a 2xx status code
func (o *GetAPIkeysInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get a p ikeys internal server error response has a 3xx status code
func (o *GetAPIkeysInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get a p ikeys internal server error response has a 4xx status code
func (o *GetAPIkeysInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get a p ikeys internal server error response has a 5xx status code
func (o *GetAPIkeysInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get a p ikeys internal server error response a status code equal to that given
func (o *GetAPIkeysInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get a p ikeys internal server error response
func (o *GetAPIkeysInternalServerError) Code() int {
	return 500
}

func (o *GetAPIkeysInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysInternalServerError %s", 500, payload)
}

func (o *GetAPIkeysInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /auth/keys][%d] getAPIkeysInternalServerError %s", 500, payload)
}

func (o *GetAPIkeysInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAPIkeysInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
