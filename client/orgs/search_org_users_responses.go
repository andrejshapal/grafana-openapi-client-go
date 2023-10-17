// Code generated by go-swagger; DO NOT EDIT.

package orgs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// SearchOrgUsersReader is a Reader for the SearchOrgUsers structure.
type SearchOrgUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchOrgUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchOrgUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSearchOrgUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchOrgUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchOrgUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{org_id}/users/search] searchOrgUsers", response, response.Code())
	}
}

// NewSearchOrgUsersOK creates a SearchOrgUsersOK with default headers values
func NewSearchOrgUsersOK() *SearchOrgUsersOK {
	return &SearchOrgUsersOK{}
}

/*
SearchOrgUsersOK describes a response with status code 200, with default header values.

(empty)
*/
type SearchOrgUsersOK struct {
	Payload *models.SearchOrgUsersQueryResult
}

// IsSuccess returns true when this search org users o k response has a 2xx status code
func (o *SearchOrgUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search org users o k response has a 3xx status code
func (o *SearchOrgUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search org users o k response has a 4xx status code
func (o *SearchOrgUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search org users o k response has a 5xx status code
func (o *SearchOrgUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search org users o k response a status code equal to that given
func (o *SearchOrgUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search org users o k response
func (o *SearchOrgUsersOK) Code() int {
	return 200
}

func (o *SearchOrgUsersOK) Error() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/users/search][%d] searchOrgUsersOK  %+v", 200, o.Payload)
}

func (o *SearchOrgUsersOK) String() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/users/search][%d] searchOrgUsersOK  %+v", 200, o.Payload)
}

func (o *SearchOrgUsersOK) GetPayload() *models.SearchOrgUsersQueryResult {
	return o.Payload
}

func (o *SearchOrgUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SearchOrgUsersQueryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchOrgUsersUnauthorized creates a SearchOrgUsersUnauthorized with default headers values
func NewSearchOrgUsersUnauthorized() *SearchOrgUsersUnauthorized {
	return &SearchOrgUsersUnauthorized{}
}

/*
SearchOrgUsersUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type SearchOrgUsersUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this search org users unauthorized response has a 2xx status code
func (o *SearchOrgUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search org users unauthorized response has a 3xx status code
func (o *SearchOrgUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search org users unauthorized response has a 4xx status code
func (o *SearchOrgUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search org users unauthorized response has a 5xx status code
func (o *SearchOrgUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search org users unauthorized response a status code equal to that given
func (o *SearchOrgUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the search org users unauthorized response
func (o *SearchOrgUsersUnauthorized) Code() int {
	return 401
}

func (o *SearchOrgUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/users/search][%d] searchOrgUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchOrgUsersUnauthorized) String() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/users/search][%d] searchOrgUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchOrgUsersUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchOrgUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchOrgUsersForbidden creates a SearchOrgUsersForbidden with default headers values
func NewSearchOrgUsersForbidden() *SearchOrgUsersForbidden {
	return &SearchOrgUsersForbidden{}
}

/*
SearchOrgUsersForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type SearchOrgUsersForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this search org users forbidden response has a 2xx status code
func (o *SearchOrgUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search org users forbidden response has a 3xx status code
func (o *SearchOrgUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search org users forbidden response has a 4xx status code
func (o *SearchOrgUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search org users forbidden response has a 5xx status code
func (o *SearchOrgUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search org users forbidden response a status code equal to that given
func (o *SearchOrgUsersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the search org users forbidden response
func (o *SearchOrgUsersForbidden) Code() int {
	return 403
}

func (o *SearchOrgUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/users/search][%d] searchOrgUsersForbidden  %+v", 403, o.Payload)
}

func (o *SearchOrgUsersForbidden) String() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/users/search][%d] searchOrgUsersForbidden  %+v", 403, o.Payload)
}

func (o *SearchOrgUsersForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchOrgUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchOrgUsersInternalServerError creates a SearchOrgUsersInternalServerError with default headers values
func NewSearchOrgUsersInternalServerError() *SearchOrgUsersInternalServerError {
	return &SearchOrgUsersInternalServerError{}
}

/*
SearchOrgUsersInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SearchOrgUsersInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this search org users internal server error response has a 2xx status code
func (o *SearchOrgUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search org users internal server error response has a 3xx status code
func (o *SearchOrgUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search org users internal server error response has a 4xx status code
func (o *SearchOrgUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search org users internal server error response has a 5xx status code
func (o *SearchOrgUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search org users internal server error response a status code equal to that given
func (o *SearchOrgUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the search org users internal server error response
func (o *SearchOrgUsersInternalServerError) Code() int {
	return 500
}

func (o *SearchOrgUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/users/search][%d] searchOrgUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchOrgUsersInternalServerError) String() string {
	return fmt.Sprintf("[GET /orgs/{org_id}/users/search][%d] searchOrgUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchOrgUsersInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SearchOrgUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}