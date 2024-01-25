// Code generated by go-swagger; DO NOT EDIT.

package admin_users

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

// AdminUpdateUserPasswordReader is a Reader for the AdminUpdateUserPassword structure.
type AdminUpdateUserPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUpdateUserPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminUpdateUserPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUpdateUserPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminUpdateUserPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminUpdateUserPasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminUpdateUserPasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /admin/users/{user_id}/password] adminUpdateUserPassword", response, response.Code())
	}
}

// NewAdminUpdateUserPasswordOK creates a AdminUpdateUserPasswordOK with default headers values
func NewAdminUpdateUserPasswordOK() *AdminUpdateUserPasswordOK {
	return &AdminUpdateUserPasswordOK{}
}

/*
AdminUpdateUserPasswordOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AdminUpdateUserPasswordOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this admin update user password Ok response has a 2xx status code
func (o *AdminUpdateUserPasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin update user password Ok response has a 3xx status code
func (o *AdminUpdateUserPasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user password Ok response has a 4xx status code
func (o *AdminUpdateUserPasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin update user password Ok response has a 5xx status code
func (o *AdminUpdateUserPasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user password Ok response a status code equal to that given
func (o *AdminUpdateUserPasswordOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin update user password Ok response
func (o *AdminUpdateUserPasswordOK) Code() int {
	return 200
}

func (o *AdminUpdateUserPasswordOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordOk %s", 200, payload)
}

func (o *AdminUpdateUserPasswordOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordOk %s", 200, payload)
}

func (o *AdminUpdateUserPasswordOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordBadRequest creates a AdminUpdateUserPasswordBadRequest with default headers values
func NewAdminUpdateUserPasswordBadRequest() *AdminUpdateUserPasswordBadRequest {
	return &AdminUpdateUserPasswordBadRequest{}
}

/*
AdminUpdateUserPasswordBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AdminUpdateUserPasswordBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin update user password bad request response has a 2xx status code
func (o *AdminUpdateUserPasswordBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user password bad request response has a 3xx status code
func (o *AdminUpdateUserPasswordBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user password bad request response has a 4xx status code
func (o *AdminUpdateUserPasswordBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user password bad request response has a 5xx status code
func (o *AdminUpdateUserPasswordBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user password bad request response a status code equal to that given
func (o *AdminUpdateUserPasswordBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the admin update user password bad request response
func (o *AdminUpdateUserPasswordBadRequest) Code() int {
	return 400
}

func (o *AdminUpdateUserPasswordBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordBadRequest %s", 400, payload)
}

func (o *AdminUpdateUserPasswordBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordBadRequest %s", 400, payload)
}

func (o *AdminUpdateUserPasswordBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordUnauthorized creates a AdminUpdateUserPasswordUnauthorized with default headers values
func NewAdminUpdateUserPasswordUnauthorized() *AdminUpdateUserPasswordUnauthorized {
	return &AdminUpdateUserPasswordUnauthorized{}
}

/*
AdminUpdateUserPasswordUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminUpdateUserPasswordUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin update user password unauthorized response has a 2xx status code
func (o *AdminUpdateUserPasswordUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user password unauthorized response has a 3xx status code
func (o *AdminUpdateUserPasswordUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user password unauthorized response has a 4xx status code
func (o *AdminUpdateUserPasswordUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user password unauthorized response has a 5xx status code
func (o *AdminUpdateUserPasswordUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user password unauthorized response a status code equal to that given
func (o *AdminUpdateUserPasswordUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin update user password unauthorized response
func (o *AdminUpdateUserPasswordUnauthorized) Code() int {
	return 401
}

func (o *AdminUpdateUserPasswordUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordUnauthorized %s", 401, payload)
}

func (o *AdminUpdateUserPasswordUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordUnauthorized %s", 401, payload)
}

func (o *AdminUpdateUserPasswordUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordForbidden creates a AdminUpdateUserPasswordForbidden with default headers values
func NewAdminUpdateUserPasswordForbidden() *AdminUpdateUserPasswordForbidden {
	return &AdminUpdateUserPasswordForbidden{}
}

/*
AdminUpdateUserPasswordForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminUpdateUserPasswordForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin update user password forbidden response has a 2xx status code
func (o *AdminUpdateUserPasswordForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user password forbidden response has a 3xx status code
func (o *AdminUpdateUserPasswordForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user password forbidden response has a 4xx status code
func (o *AdminUpdateUserPasswordForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user password forbidden response has a 5xx status code
func (o *AdminUpdateUserPasswordForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user password forbidden response a status code equal to that given
func (o *AdminUpdateUserPasswordForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin update user password forbidden response
func (o *AdminUpdateUserPasswordForbidden) Code() int {
	return 403
}

func (o *AdminUpdateUserPasswordForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordForbidden %s", 403, payload)
}

func (o *AdminUpdateUserPasswordForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordForbidden %s", 403, payload)
}

func (o *AdminUpdateUserPasswordForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPasswordInternalServerError creates a AdminUpdateUserPasswordInternalServerError with default headers values
func NewAdminUpdateUserPasswordInternalServerError() *AdminUpdateUserPasswordInternalServerError {
	return &AdminUpdateUserPasswordInternalServerError{}
}

/*
AdminUpdateUserPasswordInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminUpdateUserPasswordInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin update user password internal server error response has a 2xx status code
func (o *AdminUpdateUserPasswordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user password internal server error response has a 3xx status code
func (o *AdminUpdateUserPasswordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user password internal server error response has a 4xx status code
func (o *AdminUpdateUserPasswordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin update user password internal server error response has a 5xx status code
func (o *AdminUpdateUserPasswordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin update user password internal server error response a status code equal to that given
func (o *AdminUpdateUserPasswordInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin update user password internal server error response
func (o *AdminUpdateUserPasswordInternalServerError) Code() int {
	return 500
}

func (o *AdminUpdateUserPasswordInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordInternalServerError %s", 500, payload)
}

func (o *AdminUpdateUserPasswordInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/password][%d] adminUpdateUserPasswordInternalServerError %s", 500, payload)
}

func (o *AdminUpdateUserPasswordInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
