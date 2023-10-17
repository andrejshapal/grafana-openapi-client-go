// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetUserTeamsReader is a Reader for the GetUserTeams structure.
type GetUserTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserTeamsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserTeamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserTeamsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserTeamsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/{user_id}/teams] getUserTeams", response, response.Code())
	}
}

// NewGetUserTeamsOK creates a GetUserTeamsOK with default headers values
func NewGetUserTeamsOK() *GetUserTeamsOK {
	return &GetUserTeamsOK{}
}

/*
GetUserTeamsOK describes a response with status code 200, with default header values.

(empty)
*/
type GetUserTeamsOK struct {
	Payload []*models.TeamDTO
}

// IsSuccess returns true when this get user teams o k response has a 2xx status code
func (o *GetUserTeamsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user teams o k response has a 3xx status code
func (o *GetUserTeamsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user teams o k response has a 4xx status code
func (o *GetUserTeamsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user teams o k response has a 5xx status code
func (o *GetUserTeamsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user teams o k response a status code equal to that given
func (o *GetUserTeamsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user teams o k response
func (o *GetUserTeamsOK) Code() int {
	return 200
}

func (o *GetUserTeamsOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsOK  %+v", 200, o.Payload)
}

func (o *GetUserTeamsOK) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsOK  %+v", 200, o.Payload)
}

func (o *GetUserTeamsOK) GetPayload() []*models.TeamDTO {
	return o.Payload
}

func (o *GetUserTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTeamsUnauthorized creates a GetUserTeamsUnauthorized with default headers values
func NewGetUserTeamsUnauthorized() *GetUserTeamsUnauthorized {
	return &GetUserTeamsUnauthorized{}
}

/*
GetUserTeamsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetUserTeamsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user teams unauthorized response has a 2xx status code
func (o *GetUserTeamsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user teams unauthorized response has a 3xx status code
func (o *GetUserTeamsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user teams unauthorized response has a 4xx status code
func (o *GetUserTeamsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user teams unauthorized response has a 5xx status code
func (o *GetUserTeamsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user teams unauthorized response a status code equal to that given
func (o *GetUserTeamsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get user teams unauthorized response
func (o *GetUserTeamsUnauthorized) Code() int {
	return 401
}

func (o *GetUserTeamsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserTeamsUnauthorized) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserTeamsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserTeamsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTeamsForbidden creates a GetUserTeamsForbidden with default headers values
func NewGetUserTeamsForbidden() *GetUserTeamsForbidden {
	return &GetUserTeamsForbidden{}
}

/*
GetUserTeamsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetUserTeamsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user teams forbidden response has a 2xx status code
func (o *GetUserTeamsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user teams forbidden response has a 3xx status code
func (o *GetUserTeamsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user teams forbidden response has a 4xx status code
func (o *GetUserTeamsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user teams forbidden response has a 5xx status code
func (o *GetUserTeamsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user teams forbidden response a status code equal to that given
func (o *GetUserTeamsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get user teams forbidden response
func (o *GetUserTeamsForbidden) Code() int {
	return 403
}

func (o *GetUserTeamsForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserTeamsForbidden) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserTeamsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserTeamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTeamsNotFound creates a GetUserTeamsNotFound with default headers values
func NewGetUserTeamsNotFound() *GetUserTeamsNotFound {
	return &GetUserTeamsNotFound{}
}

/*
GetUserTeamsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetUserTeamsNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user teams not found response has a 2xx status code
func (o *GetUserTeamsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user teams not found response has a 3xx status code
func (o *GetUserTeamsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user teams not found response has a 4xx status code
func (o *GetUserTeamsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user teams not found response has a 5xx status code
func (o *GetUserTeamsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user teams not found response a status code equal to that given
func (o *GetUserTeamsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get user teams not found response
func (o *GetUserTeamsNotFound) Code() int {
	return 404
}

func (o *GetUserTeamsNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserTeamsNotFound) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserTeamsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserTeamsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTeamsInternalServerError creates a GetUserTeamsInternalServerError with default headers values
func NewGetUserTeamsInternalServerError() *GetUserTeamsInternalServerError {
	return &GetUserTeamsInternalServerError{}
}

/*
GetUserTeamsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetUserTeamsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user teams internal server error response has a 2xx status code
func (o *GetUserTeamsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user teams internal server error response has a 3xx status code
func (o *GetUserTeamsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user teams internal server error response has a 4xx status code
func (o *GetUserTeamsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user teams internal server error response has a 5xx status code
func (o *GetUserTeamsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user teams internal server error response a status code equal to that given
func (o *GetUserTeamsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get user teams internal server error response
func (o *GetUserTeamsInternalServerError) Code() int {
	return 500
}

func (o *GetUserTeamsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserTeamsInternalServerError) String() string {
	return fmt.Sprintf("[GET /users/{user_id}/teams][%d] getUserTeamsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserTeamsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserTeamsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}