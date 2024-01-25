// Code generated by go-swagger; DO NOT EDIT.

package access_control

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

// ListUsersRolesReader is a Reader for the ListUsersRoles structure.
type ListUsersRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUsersRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUsersRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListUsersRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListUsersRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListUsersRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /access-control/users/roles/search] listUsersRoles", response, response.Code())
	}
}

// NewListUsersRolesOK creates a ListUsersRolesOK with default headers values
func NewListUsersRolesOK() *ListUsersRolesOK {
	return &ListUsersRolesOK{}
}

/*
ListUsersRolesOK describes a response with status code 200, with default header values.

(empty)
*/
type ListUsersRolesOK struct {
	Payload map[string][]models.RoleDTO
}

// IsSuccess returns true when this list users roles Ok response has a 2xx status code
func (o *ListUsersRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list users roles Ok response has a 3xx status code
func (o *ListUsersRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list users roles Ok response has a 4xx status code
func (o *ListUsersRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list users roles Ok response has a 5xx status code
func (o *ListUsersRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list users roles Ok response a status code equal to that given
func (o *ListUsersRolesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list users roles Ok response
func (o *ListUsersRolesOK) Code() int {
	return 200
}

func (o *ListUsersRolesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/users/roles/search][%d] listUsersRolesOk %s", 200, payload)
}

func (o *ListUsersRolesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/users/roles/search][%d] listUsersRolesOk %s", 200, payload)
}

func (o *ListUsersRolesOK) GetPayload() map[string][]models.RoleDTO {
	return o.Payload
}

func (o *ListUsersRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUsersRolesBadRequest creates a ListUsersRolesBadRequest with default headers values
func NewListUsersRolesBadRequest() *ListUsersRolesBadRequest {
	return &ListUsersRolesBadRequest{}
}

/*
ListUsersRolesBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type ListUsersRolesBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this list users roles bad request response has a 2xx status code
func (o *ListUsersRolesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list users roles bad request response has a 3xx status code
func (o *ListUsersRolesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list users roles bad request response has a 4xx status code
func (o *ListUsersRolesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list users roles bad request response has a 5xx status code
func (o *ListUsersRolesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list users roles bad request response a status code equal to that given
func (o *ListUsersRolesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list users roles bad request response
func (o *ListUsersRolesBadRequest) Code() int {
	return 400
}

func (o *ListUsersRolesBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/users/roles/search][%d] listUsersRolesBadRequest %s", 400, payload)
}

func (o *ListUsersRolesBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/users/roles/search][%d] listUsersRolesBadRequest %s", 400, payload)
}

func (o *ListUsersRolesBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListUsersRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUsersRolesForbidden creates a ListUsersRolesForbidden with default headers values
func NewListUsersRolesForbidden() *ListUsersRolesForbidden {
	return &ListUsersRolesForbidden{}
}

/*
ListUsersRolesForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type ListUsersRolesForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this list users roles forbidden response has a 2xx status code
func (o *ListUsersRolesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list users roles forbidden response has a 3xx status code
func (o *ListUsersRolesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list users roles forbidden response has a 4xx status code
func (o *ListUsersRolesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list users roles forbidden response has a 5xx status code
func (o *ListUsersRolesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list users roles forbidden response a status code equal to that given
func (o *ListUsersRolesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list users roles forbidden response
func (o *ListUsersRolesForbidden) Code() int {
	return 403
}

func (o *ListUsersRolesForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/users/roles/search][%d] listUsersRolesForbidden %s", 403, payload)
}

func (o *ListUsersRolesForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/users/roles/search][%d] listUsersRolesForbidden %s", 403, payload)
}

func (o *ListUsersRolesForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListUsersRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUsersRolesInternalServerError creates a ListUsersRolesInternalServerError with default headers values
func NewListUsersRolesInternalServerError() *ListUsersRolesInternalServerError {
	return &ListUsersRolesInternalServerError{}
}

/*
ListUsersRolesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type ListUsersRolesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this list users roles internal server error response has a 2xx status code
func (o *ListUsersRolesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list users roles internal server error response has a 3xx status code
func (o *ListUsersRolesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list users roles internal server error response has a 4xx status code
func (o *ListUsersRolesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list users roles internal server error response has a 5xx status code
func (o *ListUsersRolesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list users roles internal server error response a status code equal to that given
func (o *ListUsersRolesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list users roles internal server error response
func (o *ListUsersRolesInternalServerError) Code() int {
	return 500
}

func (o *ListUsersRolesInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/users/roles/search][%d] listUsersRolesInternalServerError %s", 500, payload)
}

func (o *ListUsersRolesInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /access-control/users/roles/search][%d] listUsersRolesInternalServerError %s", 500, payload)
}

func (o *ListUsersRolesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListUsersRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
