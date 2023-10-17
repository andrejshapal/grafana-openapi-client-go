// Code generated by go-swagger; DO NOT EDIT.

package user_preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetUserPreferencesReader is a Reader for the GetUserPreferences structure.
type GetUserPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserPreferencesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserPreferencesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /user/preferences] getUserPreferences", response, response.Code())
	}
}

// NewGetUserPreferencesOK creates a GetUserPreferencesOK with default headers values
func NewGetUserPreferencesOK() *GetUserPreferencesOK {
	return &GetUserPreferencesOK{}
}

/*
GetUserPreferencesOK describes a response with status code 200, with default header values.

(empty)
*/
type GetUserPreferencesOK struct {
	Payload *models.Spec
}

// IsSuccess returns true when this get user preferences o k response has a 2xx status code
func (o *GetUserPreferencesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user preferences o k response has a 3xx status code
func (o *GetUserPreferencesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user preferences o k response has a 4xx status code
func (o *GetUserPreferencesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user preferences o k response has a 5xx status code
func (o *GetUserPreferencesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user preferences o k response a status code equal to that given
func (o *GetUserPreferencesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user preferences o k response
func (o *GetUserPreferencesOK) Code() int {
	return 200
}

func (o *GetUserPreferencesOK) Error() string {
	return fmt.Sprintf("[GET /user/preferences][%d] getUserPreferencesOK  %+v", 200, o.Payload)
}

func (o *GetUserPreferencesOK) String() string {
	return fmt.Sprintf("[GET /user/preferences][%d] getUserPreferencesOK  %+v", 200, o.Payload)
}

func (o *GetUserPreferencesOK) GetPayload() *models.Spec {
	return o.Payload
}

func (o *GetUserPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Spec)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPreferencesUnauthorized creates a GetUserPreferencesUnauthorized with default headers values
func NewGetUserPreferencesUnauthorized() *GetUserPreferencesUnauthorized {
	return &GetUserPreferencesUnauthorized{}
}

/*
GetUserPreferencesUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetUserPreferencesUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user preferences unauthorized response has a 2xx status code
func (o *GetUserPreferencesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user preferences unauthorized response has a 3xx status code
func (o *GetUserPreferencesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user preferences unauthorized response has a 4xx status code
func (o *GetUserPreferencesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user preferences unauthorized response has a 5xx status code
func (o *GetUserPreferencesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user preferences unauthorized response a status code equal to that given
func (o *GetUserPreferencesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get user preferences unauthorized response
func (o *GetUserPreferencesUnauthorized) Code() int {
	return 401
}

func (o *GetUserPreferencesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /user/preferences][%d] getUserPreferencesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserPreferencesUnauthorized) String() string {
	return fmt.Sprintf("[GET /user/preferences][%d] getUserPreferencesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserPreferencesUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserPreferencesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserPreferencesInternalServerError creates a GetUserPreferencesInternalServerError with default headers values
func NewGetUserPreferencesInternalServerError() *GetUserPreferencesInternalServerError {
	return &GetUserPreferencesInternalServerError{}
}

/*
GetUserPreferencesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetUserPreferencesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user preferences internal server error response has a 2xx status code
func (o *GetUserPreferencesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user preferences internal server error response has a 3xx status code
func (o *GetUserPreferencesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user preferences internal server error response has a 4xx status code
func (o *GetUserPreferencesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user preferences internal server error response has a 5xx status code
func (o *GetUserPreferencesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user preferences internal server error response a status code equal to that given
func (o *GetUserPreferencesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get user preferences internal server error response
func (o *GetUserPreferencesInternalServerError) Code() int {
	return 500
}

func (o *GetUserPreferencesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user/preferences][%d] getUserPreferencesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserPreferencesInternalServerError) String() string {
	return fmt.Sprintf("[GET /user/preferences][%d] getUserPreferencesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserPreferencesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserPreferencesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}