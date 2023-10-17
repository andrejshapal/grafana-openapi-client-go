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

// SearchUsersWithPagingReader is a Reader for the SearchUsersWithPaging structure.
type SearchUsersWithPagingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchUsersWithPagingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchUsersWithPagingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSearchUsersWithPagingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchUsersWithPagingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchUsersWithPagingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchUsersWithPagingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/search] searchUsersWithPaging", response, response.Code())
	}
}

// NewSearchUsersWithPagingOK creates a SearchUsersWithPagingOK with default headers values
func NewSearchUsersWithPagingOK() *SearchUsersWithPagingOK {
	return &SearchUsersWithPagingOK{}
}

/*
SearchUsersWithPagingOK describes a response with status code 200, with default header values.

(empty)
*/
type SearchUsersWithPagingOK struct {
	Payload *models.SearchUserQueryResult
}

// IsSuccess returns true when this search users with paging o k response has a 2xx status code
func (o *SearchUsersWithPagingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search users with paging o k response has a 3xx status code
func (o *SearchUsersWithPagingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search users with paging o k response has a 4xx status code
func (o *SearchUsersWithPagingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search users with paging o k response has a 5xx status code
func (o *SearchUsersWithPagingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search users with paging o k response a status code equal to that given
func (o *SearchUsersWithPagingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search users with paging o k response
func (o *SearchUsersWithPagingOK) Code() int {
	return 200
}

func (o *SearchUsersWithPagingOK) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingOK  %+v", 200, o.Payload)
}

func (o *SearchUsersWithPagingOK) String() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingOK  %+v", 200, o.Payload)
}

func (o *SearchUsersWithPagingOK) GetPayload() *models.SearchUserQueryResult {
	return o.Payload
}

func (o *SearchUsersWithPagingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchUserQueryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersWithPagingUnauthorized creates a SearchUsersWithPagingUnauthorized with default headers values
func NewSearchUsersWithPagingUnauthorized() *SearchUsersWithPagingUnauthorized {
	return &SearchUsersWithPagingUnauthorized{}
}

/*
SearchUsersWithPagingUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type SearchUsersWithPagingUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this search users with paging unauthorized response has a 2xx status code
func (o *SearchUsersWithPagingUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search users with paging unauthorized response has a 3xx status code
func (o *SearchUsersWithPagingUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search users with paging unauthorized response has a 4xx status code
func (o *SearchUsersWithPagingUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search users with paging unauthorized response has a 5xx status code
func (o *SearchUsersWithPagingUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search users with paging unauthorized response a status code equal to that given
func (o *SearchUsersWithPagingUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the search users with paging unauthorized response
func (o *SearchUsersWithPagingUnauthorized) Code() int {
	return 401
}

func (o *SearchUsersWithPagingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchUsersWithPagingUnauthorized) String() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchUsersWithPagingUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchUsersWithPagingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersWithPagingForbidden creates a SearchUsersWithPagingForbidden with default headers values
func NewSearchUsersWithPagingForbidden() *SearchUsersWithPagingForbidden {
	return &SearchUsersWithPagingForbidden{}
}

/*
SearchUsersWithPagingForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type SearchUsersWithPagingForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this search users with paging forbidden response has a 2xx status code
func (o *SearchUsersWithPagingForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search users with paging forbidden response has a 3xx status code
func (o *SearchUsersWithPagingForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search users with paging forbidden response has a 4xx status code
func (o *SearchUsersWithPagingForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search users with paging forbidden response has a 5xx status code
func (o *SearchUsersWithPagingForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search users with paging forbidden response a status code equal to that given
func (o *SearchUsersWithPagingForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the search users with paging forbidden response
func (o *SearchUsersWithPagingForbidden) Code() int {
	return 403
}

func (o *SearchUsersWithPagingForbidden) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingForbidden  %+v", 403, o.Payload)
}

func (o *SearchUsersWithPagingForbidden) String() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingForbidden  %+v", 403, o.Payload)
}

func (o *SearchUsersWithPagingForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchUsersWithPagingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersWithPagingNotFound creates a SearchUsersWithPagingNotFound with default headers values
func NewSearchUsersWithPagingNotFound() *SearchUsersWithPagingNotFound {
	return &SearchUsersWithPagingNotFound{}
}

/*
SearchUsersWithPagingNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type SearchUsersWithPagingNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this search users with paging not found response has a 2xx status code
func (o *SearchUsersWithPagingNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search users with paging not found response has a 3xx status code
func (o *SearchUsersWithPagingNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search users with paging not found response has a 4xx status code
func (o *SearchUsersWithPagingNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this search users with paging not found response has a 5xx status code
func (o *SearchUsersWithPagingNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this search users with paging not found response a status code equal to that given
func (o *SearchUsersWithPagingNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the search users with paging not found response
func (o *SearchUsersWithPagingNotFound) Code() int {
	return 404
}

func (o *SearchUsersWithPagingNotFound) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingNotFound  %+v", 404, o.Payload)
}

func (o *SearchUsersWithPagingNotFound) String() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingNotFound  %+v", 404, o.Payload)
}

func (o *SearchUsersWithPagingNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchUsersWithPagingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsersWithPagingInternalServerError creates a SearchUsersWithPagingInternalServerError with default headers values
func NewSearchUsersWithPagingInternalServerError() *SearchUsersWithPagingInternalServerError {
	return &SearchUsersWithPagingInternalServerError{}
}

/*
SearchUsersWithPagingInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SearchUsersWithPagingInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this search users with paging internal server error response has a 2xx status code
func (o *SearchUsersWithPagingInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search users with paging internal server error response has a 3xx status code
func (o *SearchUsersWithPagingInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search users with paging internal server error response has a 4xx status code
func (o *SearchUsersWithPagingInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search users with paging internal server error response has a 5xx status code
func (o *SearchUsersWithPagingInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search users with paging internal server error response a status code equal to that given
func (o *SearchUsersWithPagingInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the search users with paging internal server error response
func (o *SearchUsersWithPagingInternalServerError) Code() int {
	return 500
}

func (o *SearchUsersWithPagingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchUsersWithPagingInternalServerError) String() string {
	return fmt.Sprintf("[GET /users/search][%d] searchUsersWithPagingInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchUsersWithPagingInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchUsersWithPagingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}