// Code generated by go-swagger; DO NOT EDIT.

package users

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

// GetUserByIDReader is a Reader for the GetUserByID structure.
type GetUserByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUserByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /users/{user_id}] getUserByID", response, response.Code())
	}
}

// NewGetUserByIDOK creates a GetUserByIDOK with default headers values
func NewGetUserByIDOK() *GetUserByIDOK {
	return &GetUserByIDOK{}
}

/*
GetUserByIDOK describes a response with status code 200, with default header values.

(empty)
*/
type GetUserByIDOK struct {
	Payload *models.UserProfileDTO
}

// IsSuccess returns true when this get user by Id Ok response has a 2xx status code
func (o *GetUserByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user by Id Ok response has a 3xx status code
func (o *GetUserByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user by Id Ok response has a 4xx status code
func (o *GetUserByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user by Id Ok response has a 5xx status code
func (o *GetUserByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user by Id Ok response a status code equal to that given
func (o *GetUserByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get user by Id Ok response
func (o *GetUserByIDOK) Code() int {
	return 200
}

func (o *GetUserByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdOk %s", 200, payload)
}

func (o *GetUserByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdOk %s", 200, payload)
}

func (o *GetUserByIDOK) GetPayload() *models.UserProfileDTO {
	return o.Payload
}

func (o *GetUserByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserProfileDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByIDUnauthorized creates a GetUserByIDUnauthorized with default headers values
func NewGetUserByIDUnauthorized() *GetUserByIDUnauthorized {
	return &GetUserByIDUnauthorized{}
}

/*
GetUserByIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetUserByIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user by Id unauthorized response has a 2xx status code
func (o *GetUserByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user by Id unauthorized response has a 3xx status code
func (o *GetUserByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user by Id unauthorized response has a 4xx status code
func (o *GetUserByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user by Id unauthorized response has a 5xx status code
func (o *GetUserByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user by Id unauthorized response a status code equal to that given
func (o *GetUserByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get user by Id unauthorized response
func (o *GetUserByIDUnauthorized) Code() int {
	return 401
}

func (o *GetUserByIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdUnauthorized %s", 401, payload)
}

func (o *GetUserByIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdUnauthorized %s", 401, payload)
}

func (o *GetUserByIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByIDForbidden creates a GetUserByIDForbidden with default headers values
func NewGetUserByIDForbidden() *GetUserByIDForbidden {
	return &GetUserByIDForbidden{}
}

/*
GetUserByIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetUserByIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user by Id forbidden response has a 2xx status code
func (o *GetUserByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user by Id forbidden response has a 3xx status code
func (o *GetUserByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user by Id forbidden response has a 4xx status code
func (o *GetUserByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user by Id forbidden response has a 5xx status code
func (o *GetUserByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user by Id forbidden response a status code equal to that given
func (o *GetUserByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get user by Id forbidden response
func (o *GetUserByIDForbidden) Code() int {
	return 403
}

func (o *GetUserByIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdForbidden %s", 403, payload)
}

func (o *GetUserByIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdForbidden %s", 403, payload)
}

func (o *GetUserByIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByIDNotFound creates a GetUserByIDNotFound with default headers values
func NewGetUserByIDNotFound() *GetUserByIDNotFound {
	return &GetUserByIDNotFound{}
}

/*
GetUserByIDNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetUserByIDNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user by Id not found response has a 2xx status code
func (o *GetUserByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user by Id not found response has a 3xx status code
func (o *GetUserByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user by Id not found response has a 4xx status code
func (o *GetUserByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user by Id not found response has a 5xx status code
func (o *GetUserByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user by Id not found response a status code equal to that given
func (o *GetUserByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get user by Id not found response
func (o *GetUserByIDNotFound) Code() int {
	return 404
}

func (o *GetUserByIDNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdNotFound %s", 404, payload)
}

func (o *GetUserByIDNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdNotFound %s", 404, payload)
}

func (o *GetUserByIDNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserByIDInternalServerError creates a GetUserByIDInternalServerError with default headers values
func NewGetUserByIDInternalServerError() *GetUserByIDInternalServerError {
	return &GetUserByIDInternalServerError{}
}

/*
GetUserByIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetUserByIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get user by Id internal server error response has a 2xx status code
func (o *GetUserByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user by Id internal server error response has a 3xx status code
func (o *GetUserByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user by Id internal server error response has a 4xx status code
func (o *GetUserByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user by Id internal server error response has a 5xx status code
func (o *GetUserByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user by Id internal server error response a status code equal to that given
func (o *GetUserByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get user by Id internal server error response
func (o *GetUserByIDInternalServerError) Code() int {
	return 500
}

func (o *GetUserByIDInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdInternalServerError %s", 500, payload)
}

func (o *GetUserByIDInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUserByIdInternalServerError %s", 500, payload)
}

func (o *GetUserByIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetUserByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
