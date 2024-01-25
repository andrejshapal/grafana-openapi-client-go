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

// AdminUpdateUserPermissionsReader is a Reader for the AdminUpdateUserPermissions structure.
type AdminUpdateUserPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AdminUpdateUserPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAdminUpdateUserPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAdminUpdateUserPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAdminUpdateUserPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAdminUpdateUserPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAdminUpdateUserPermissionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /admin/users/{user_id}/permissions] adminUpdateUserPermissions", response, response.Code())
	}
}

// NewAdminUpdateUserPermissionsOK creates a AdminUpdateUserPermissionsOK with default headers values
func NewAdminUpdateUserPermissionsOK() *AdminUpdateUserPermissionsOK {
	return &AdminUpdateUserPermissionsOK{}
}

/*
AdminUpdateUserPermissionsOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type AdminUpdateUserPermissionsOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this admin update user permissions Ok response has a 2xx status code
func (o *AdminUpdateUserPermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this admin update user permissions Ok response has a 3xx status code
func (o *AdminUpdateUserPermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user permissions Ok response has a 4xx status code
func (o *AdminUpdateUserPermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin update user permissions Ok response has a 5xx status code
func (o *AdminUpdateUserPermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user permissions Ok response a status code equal to that given
func (o *AdminUpdateUserPermissionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the admin update user permissions Ok response
func (o *AdminUpdateUserPermissionsOK) Code() int {
	return 200
}

func (o *AdminUpdateUserPermissionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/permissions][%d] adminUpdateUserPermissionsOk %s", 200, payload)
}

func (o *AdminUpdateUserPermissionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/permissions][%d] adminUpdateUserPermissionsOk %s", 200, payload)
}

func (o *AdminUpdateUserPermissionsOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPermissionsBadRequest creates a AdminUpdateUserPermissionsBadRequest with default headers values
func NewAdminUpdateUserPermissionsBadRequest() *AdminUpdateUserPermissionsBadRequest {
	return &AdminUpdateUserPermissionsBadRequest{}
}

/*
AdminUpdateUserPermissionsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type AdminUpdateUserPermissionsBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin update user permissions bad request response has a 2xx status code
func (o *AdminUpdateUserPermissionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user permissions bad request response has a 3xx status code
func (o *AdminUpdateUserPermissionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user permissions bad request response has a 4xx status code
func (o *AdminUpdateUserPermissionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user permissions bad request response has a 5xx status code
func (o *AdminUpdateUserPermissionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user permissions bad request response a status code equal to that given
func (o *AdminUpdateUserPermissionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the admin update user permissions bad request response
func (o *AdminUpdateUserPermissionsBadRequest) Code() int {
	return 400
}

func (o *AdminUpdateUserPermissionsBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/permissions][%d] adminUpdateUserPermissionsBadRequest %s", 400, payload)
}

func (o *AdminUpdateUserPermissionsBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/permissions][%d] adminUpdateUserPermissionsBadRequest %s", 400, payload)
}

func (o *AdminUpdateUserPermissionsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPermissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPermissionsUnauthorized creates a AdminUpdateUserPermissionsUnauthorized with default headers values
func NewAdminUpdateUserPermissionsUnauthorized() *AdminUpdateUserPermissionsUnauthorized {
	return &AdminUpdateUserPermissionsUnauthorized{}
}

/*
AdminUpdateUserPermissionsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type AdminUpdateUserPermissionsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin update user permissions unauthorized response has a 2xx status code
func (o *AdminUpdateUserPermissionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user permissions unauthorized response has a 3xx status code
func (o *AdminUpdateUserPermissionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user permissions unauthorized response has a 4xx status code
func (o *AdminUpdateUserPermissionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user permissions unauthorized response has a 5xx status code
func (o *AdminUpdateUserPermissionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user permissions unauthorized response a status code equal to that given
func (o *AdminUpdateUserPermissionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the admin update user permissions unauthorized response
func (o *AdminUpdateUserPermissionsUnauthorized) Code() int {
	return 401
}

func (o *AdminUpdateUserPermissionsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/permissions][%d] adminUpdateUserPermissionsUnauthorized %s", 401, payload)
}

func (o *AdminUpdateUserPermissionsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/permissions][%d] adminUpdateUserPermissionsUnauthorized %s", 401, payload)
}

func (o *AdminUpdateUserPermissionsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPermissionsForbidden creates a AdminUpdateUserPermissionsForbidden with default headers values
func NewAdminUpdateUserPermissionsForbidden() *AdminUpdateUserPermissionsForbidden {
	return &AdminUpdateUserPermissionsForbidden{}
}

/*
AdminUpdateUserPermissionsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type AdminUpdateUserPermissionsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin update user permissions forbidden response has a 2xx status code
func (o *AdminUpdateUserPermissionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user permissions forbidden response has a 3xx status code
func (o *AdminUpdateUserPermissionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user permissions forbidden response has a 4xx status code
func (o *AdminUpdateUserPermissionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this admin update user permissions forbidden response has a 5xx status code
func (o *AdminUpdateUserPermissionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this admin update user permissions forbidden response a status code equal to that given
func (o *AdminUpdateUserPermissionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the admin update user permissions forbidden response
func (o *AdminUpdateUserPermissionsForbidden) Code() int {
	return 403
}

func (o *AdminUpdateUserPermissionsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/permissions][%d] adminUpdateUserPermissionsForbidden %s", 403, payload)
}

func (o *AdminUpdateUserPermissionsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/permissions][%d] adminUpdateUserPermissionsForbidden %s", 403, payload)
}

func (o *AdminUpdateUserPermissionsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPermissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAdminUpdateUserPermissionsInternalServerError creates a AdminUpdateUserPermissionsInternalServerError with default headers values
func NewAdminUpdateUserPermissionsInternalServerError() *AdminUpdateUserPermissionsInternalServerError {
	return &AdminUpdateUserPermissionsInternalServerError{}
}

/*
AdminUpdateUserPermissionsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type AdminUpdateUserPermissionsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this admin update user permissions internal server error response has a 2xx status code
func (o *AdminUpdateUserPermissionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this admin update user permissions internal server error response has a 3xx status code
func (o *AdminUpdateUserPermissionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this admin update user permissions internal server error response has a 4xx status code
func (o *AdminUpdateUserPermissionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this admin update user permissions internal server error response has a 5xx status code
func (o *AdminUpdateUserPermissionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this admin update user permissions internal server error response a status code equal to that given
func (o *AdminUpdateUserPermissionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the admin update user permissions internal server error response
func (o *AdminUpdateUserPermissionsInternalServerError) Code() int {
	return 500
}

func (o *AdminUpdateUserPermissionsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/permissions][%d] adminUpdateUserPermissionsInternalServerError %s", 500, payload)
}

func (o *AdminUpdateUserPermissionsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /admin/users/{user_id}/permissions][%d] adminUpdateUserPermissionsInternalServerError %s", 500, payload)
}

func (o *AdminUpdateUserPermissionsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *AdminUpdateUserPermissionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
