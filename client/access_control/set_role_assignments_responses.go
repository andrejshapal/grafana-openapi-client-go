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

// SetRoleAssignmentsReader is a Reader for the SetRoleAssignments structure.
type SetRoleAssignmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRoleAssignmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetRoleAssignmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSetRoleAssignmentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetRoleAssignmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetRoleAssignmentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /access-control/roles/{roleUID}/assignments] setRoleAssignments", response, response.Code())
	}
}

// NewSetRoleAssignmentsOK creates a SetRoleAssignmentsOK with default headers values
func NewSetRoleAssignmentsOK() *SetRoleAssignmentsOK {
	return &SetRoleAssignmentsOK{}
}

/*
SetRoleAssignmentsOK describes a response with status code 200, with default header values.

(empty)
*/
type SetRoleAssignmentsOK struct {
	Payload *models.RoleAssignmentsDTO
}

// IsSuccess returns true when this set role assignments Ok response has a 2xx status code
func (o *SetRoleAssignmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set role assignments Ok response has a 3xx status code
func (o *SetRoleAssignmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set role assignments Ok response has a 4xx status code
func (o *SetRoleAssignmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set role assignments Ok response has a 5xx status code
func (o *SetRoleAssignmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set role assignments Ok response a status code equal to that given
func (o *SetRoleAssignmentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set role assignments Ok response
func (o *SetRoleAssignmentsOK) Code() int {
	return 200
}

func (o *SetRoleAssignmentsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}/assignments][%d] setRoleAssignmentsOk %s", 200, payload)
}

func (o *SetRoleAssignmentsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}/assignments][%d] setRoleAssignmentsOk %s", 200, payload)
}

func (o *SetRoleAssignmentsOK) GetPayload() *models.RoleAssignmentsDTO {
	return o.Payload
}

func (o *SetRoleAssignmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleAssignmentsDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRoleAssignmentsForbidden creates a SetRoleAssignmentsForbidden with default headers values
func NewSetRoleAssignmentsForbidden() *SetRoleAssignmentsForbidden {
	return &SetRoleAssignmentsForbidden{}
}

/*
SetRoleAssignmentsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type SetRoleAssignmentsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set role assignments forbidden response has a 2xx status code
func (o *SetRoleAssignmentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set role assignments forbidden response has a 3xx status code
func (o *SetRoleAssignmentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set role assignments forbidden response has a 4xx status code
func (o *SetRoleAssignmentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this set role assignments forbidden response has a 5xx status code
func (o *SetRoleAssignmentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this set role assignments forbidden response a status code equal to that given
func (o *SetRoleAssignmentsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the set role assignments forbidden response
func (o *SetRoleAssignmentsForbidden) Code() int {
	return 403
}

func (o *SetRoleAssignmentsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}/assignments][%d] setRoleAssignmentsForbidden %s", 403, payload)
}

func (o *SetRoleAssignmentsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}/assignments][%d] setRoleAssignmentsForbidden %s", 403, payload)
}

func (o *SetRoleAssignmentsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetRoleAssignmentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRoleAssignmentsNotFound creates a SetRoleAssignmentsNotFound with default headers values
func NewSetRoleAssignmentsNotFound() *SetRoleAssignmentsNotFound {
	return &SetRoleAssignmentsNotFound{}
}

/*
SetRoleAssignmentsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type SetRoleAssignmentsNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set role assignments not found response has a 2xx status code
func (o *SetRoleAssignmentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set role assignments not found response has a 3xx status code
func (o *SetRoleAssignmentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set role assignments not found response has a 4xx status code
func (o *SetRoleAssignmentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this set role assignments not found response has a 5xx status code
func (o *SetRoleAssignmentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this set role assignments not found response a status code equal to that given
func (o *SetRoleAssignmentsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the set role assignments not found response
func (o *SetRoleAssignmentsNotFound) Code() int {
	return 404
}

func (o *SetRoleAssignmentsNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}/assignments][%d] setRoleAssignmentsNotFound %s", 404, payload)
}

func (o *SetRoleAssignmentsNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}/assignments][%d] setRoleAssignmentsNotFound %s", 404, payload)
}

func (o *SetRoleAssignmentsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetRoleAssignmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRoleAssignmentsInternalServerError creates a SetRoleAssignmentsInternalServerError with default headers values
func NewSetRoleAssignmentsInternalServerError() *SetRoleAssignmentsInternalServerError {
	return &SetRoleAssignmentsInternalServerError{}
}

/*
SetRoleAssignmentsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type SetRoleAssignmentsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this set role assignments internal server error response has a 2xx status code
func (o *SetRoleAssignmentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set role assignments internal server error response has a 3xx status code
func (o *SetRoleAssignmentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set role assignments internal server error response has a 4xx status code
func (o *SetRoleAssignmentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this set role assignments internal server error response has a 5xx status code
func (o *SetRoleAssignmentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this set role assignments internal server error response a status code equal to that given
func (o *SetRoleAssignmentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the set role assignments internal server error response
func (o *SetRoleAssignmentsInternalServerError) Code() int {
	return 500
}

func (o *SetRoleAssignmentsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}/assignments][%d] setRoleAssignmentsInternalServerError %s", 500, payload)
}

func (o *SetRoleAssignmentsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /access-control/roles/{roleUID}/assignments][%d] setRoleAssignmentsInternalServerError %s", 500, payload)
}

func (o *SetRoleAssignmentsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *SetRoleAssignmentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
