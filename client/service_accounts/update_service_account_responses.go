// Code generated by go-swagger; DO NOT EDIT.

package service_accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// UpdateServiceAccountReader is a Reader for the UpdateServiceAccount structure.
type UpdateServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateServiceAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateServiceAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateServiceAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateServiceAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateServiceAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /serviceaccounts/{serviceAccountId}] updateServiceAccount", response, response.Code())
	}
}

// NewUpdateServiceAccountOK creates a UpdateServiceAccountOK with default headers values
func NewUpdateServiceAccountOK() *UpdateServiceAccountOK {
	return &UpdateServiceAccountOK{}
}

/*
UpdateServiceAccountOK describes a response with status code 200, with default header values.

(empty)
*/
type UpdateServiceAccountOK struct {
	Payload *models.UpdateServiceAccountOKBody
}

// IsSuccess returns true when this update service account o k response has a 2xx status code
func (o *UpdateServiceAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update service account o k response has a 3xx status code
func (o *UpdateServiceAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service account o k response has a 4xx status code
func (o *UpdateServiceAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update service account o k response has a 5xx status code
func (o *UpdateServiceAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update service account o k response a status code equal to that given
func (o *UpdateServiceAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update service account o k response
func (o *UpdateServiceAccountOK) Code() int {
	return 200
}

func (o *UpdateServiceAccountOK) Error() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountOK  %+v", 200, o.Payload)
}

func (o *UpdateServiceAccountOK) String() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountOK  %+v", 200, o.Payload)
}

func (o *UpdateServiceAccountOK) GetPayload() *models.UpdateServiceAccountOKBody {
	return o.Payload
}

func (o *UpdateServiceAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateServiceAccountOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountBadRequest creates a UpdateServiceAccountBadRequest with default headers values
func NewUpdateServiceAccountBadRequest() *UpdateServiceAccountBadRequest {
	return &UpdateServiceAccountBadRequest{}
}

/*
UpdateServiceAccountBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type UpdateServiceAccountBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update service account bad request response has a 2xx status code
func (o *UpdateServiceAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service account bad request response has a 3xx status code
func (o *UpdateServiceAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service account bad request response has a 4xx status code
func (o *UpdateServiceAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service account bad request response has a 5xx status code
func (o *UpdateServiceAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update service account bad request response a status code equal to that given
func (o *UpdateServiceAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update service account bad request response
func (o *UpdateServiceAccountBadRequest) Code() int {
	return 400
}

func (o *UpdateServiceAccountBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateServiceAccountBadRequest) String() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateServiceAccountBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateServiceAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountUnauthorized creates a UpdateServiceAccountUnauthorized with default headers values
func NewUpdateServiceAccountUnauthorized() *UpdateServiceAccountUnauthorized {
	return &UpdateServiceAccountUnauthorized{}
}

/*
UpdateServiceAccountUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateServiceAccountUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update service account unauthorized response has a 2xx status code
func (o *UpdateServiceAccountUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service account unauthorized response has a 3xx status code
func (o *UpdateServiceAccountUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service account unauthorized response has a 4xx status code
func (o *UpdateServiceAccountUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service account unauthorized response has a 5xx status code
func (o *UpdateServiceAccountUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update service account unauthorized response a status code equal to that given
func (o *UpdateServiceAccountUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update service account unauthorized response
func (o *UpdateServiceAccountUnauthorized) Code() int {
	return 401
}

func (o *UpdateServiceAccountUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateServiceAccountUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateServiceAccountUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateServiceAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountForbidden creates a UpdateServiceAccountForbidden with default headers values
func NewUpdateServiceAccountForbidden() *UpdateServiceAccountForbidden {
	return &UpdateServiceAccountForbidden{}
}

/*
UpdateServiceAccountForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateServiceAccountForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update service account forbidden response has a 2xx status code
func (o *UpdateServiceAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service account forbidden response has a 3xx status code
func (o *UpdateServiceAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service account forbidden response has a 4xx status code
func (o *UpdateServiceAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service account forbidden response has a 5xx status code
func (o *UpdateServiceAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update service account forbidden response a status code equal to that given
func (o *UpdateServiceAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update service account forbidden response
func (o *UpdateServiceAccountForbidden) Code() int {
	return 403
}

func (o *UpdateServiceAccountForbidden) Error() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountForbidden  %+v", 403, o.Payload)
}

func (o *UpdateServiceAccountForbidden) String() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountForbidden  %+v", 403, o.Payload)
}

func (o *UpdateServiceAccountForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountNotFound creates a UpdateServiceAccountNotFound with default headers values
func NewUpdateServiceAccountNotFound() *UpdateServiceAccountNotFound {
	return &UpdateServiceAccountNotFound{}
}

/*
UpdateServiceAccountNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateServiceAccountNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update service account not found response has a 2xx status code
func (o *UpdateServiceAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service account not found response has a 3xx status code
func (o *UpdateServiceAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service account not found response has a 4xx status code
func (o *UpdateServiceAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update service account not found response has a 5xx status code
func (o *UpdateServiceAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update service account not found response a status code equal to that given
func (o *UpdateServiceAccountNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update service account not found response
func (o *UpdateServiceAccountNotFound) Code() int {
	return 404
}

func (o *UpdateServiceAccountNotFound) Error() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *UpdateServiceAccountNotFound) String() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountNotFound  %+v", 404, o.Payload)
}

func (o *UpdateServiceAccountNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateServiceAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateServiceAccountInternalServerError creates a UpdateServiceAccountInternalServerError with default headers values
func NewUpdateServiceAccountInternalServerError() *UpdateServiceAccountInternalServerError {
	return &UpdateServiceAccountInternalServerError{}
}

/*
UpdateServiceAccountInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateServiceAccountInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update service account internal server error response has a 2xx status code
func (o *UpdateServiceAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update service account internal server error response has a 3xx status code
func (o *UpdateServiceAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update service account internal server error response has a 4xx status code
func (o *UpdateServiceAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update service account internal server error response has a 5xx status code
func (o *UpdateServiceAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update service account internal server error response a status code equal to that given
func (o *UpdateServiceAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update service account internal server error response
func (o *UpdateServiceAccountInternalServerError) Code() int {
	return 500
}

func (o *UpdateServiceAccountInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateServiceAccountInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /serviceaccounts/{serviceAccountId}][%d] updateServiceAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateServiceAccountInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateServiceAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}