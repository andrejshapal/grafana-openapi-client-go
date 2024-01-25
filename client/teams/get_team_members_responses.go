// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// GetTeamMembersReader is a Reader for the GetTeamMembers structure.
type GetTeamMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTeamMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTeamMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTeamMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTeamMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /teams/{team_id}/members] getTeamMembers", response, response.Code())
	}
}

// NewGetTeamMembersOK creates a GetTeamMembersOK with default headers values
func NewGetTeamMembersOK() *GetTeamMembersOK {
	return &GetTeamMembersOK{}
}

/*
GetTeamMembersOK describes a response with status code 200, with default header values.

(empty)
*/
type GetTeamMembersOK struct {
	Payload []*models.TeamMemberDTO
}

// IsSuccess returns true when this get team members Ok response has a 2xx status code
func (o *GetTeamMembersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team members Ok response has a 3xx status code
func (o *GetTeamMembersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members Ok response has a 4xx status code
func (o *GetTeamMembersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team members Ok response has a 5xx status code
func (o *GetTeamMembersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members Ok response a status code equal to that given
func (o *GetTeamMembersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team members Ok response
func (o *GetTeamMembersOK) Code() int {
	return 200
}

func (o *GetTeamMembersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{team_id}/members][%d] getTeamMembersOk %s", 200, payload)
}

func (o *GetTeamMembersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{team_id}/members][%d] getTeamMembersOk %s", 200, payload)
}

func (o *GetTeamMembersOK) GetPayload() []*models.TeamMemberDTO {
	return o.Payload
}

func (o *GetTeamMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersUnauthorized creates a GetTeamMembersUnauthorized with default headers values
func NewGetTeamMembersUnauthorized() *GetTeamMembersUnauthorized {
	return &GetTeamMembersUnauthorized{}
}

/*
GetTeamMembersUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetTeamMembersUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team members unauthorized response has a 2xx status code
func (o *GetTeamMembersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members unauthorized response has a 3xx status code
func (o *GetTeamMembersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members unauthorized response has a 4xx status code
func (o *GetTeamMembersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team members unauthorized response has a 5xx status code
func (o *GetTeamMembersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members unauthorized response a status code equal to that given
func (o *GetTeamMembersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get team members unauthorized response
func (o *GetTeamMembersUnauthorized) Code() int {
	return 401
}

func (o *GetTeamMembersUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{team_id}/members][%d] getTeamMembersUnauthorized %s", 401, payload)
}

func (o *GetTeamMembersUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{team_id}/members][%d] getTeamMembersUnauthorized %s", 401, payload)
}

func (o *GetTeamMembersUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersForbidden creates a GetTeamMembersForbidden with default headers values
func NewGetTeamMembersForbidden() *GetTeamMembersForbidden {
	return &GetTeamMembersForbidden{}
}

/*
GetTeamMembersForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetTeamMembersForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team members forbidden response has a 2xx status code
func (o *GetTeamMembersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members forbidden response has a 3xx status code
func (o *GetTeamMembersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members forbidden response has a 4xx status code
func (o *GetTeamMembersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team members forbidden response has a 5xx status code
func (o *GetTeamMembersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members forbidden response a status code equal to that given
func (o *GetTeamMembersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get team members forbidden response
func (o *GetTeamMembersForbidden) Code() int {
	return 403
}

func (o *GetTeamMembersForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{team_id}/members][%d] getTeamMembersForbidden %s", 403, payload)
}

func (o *GetTeamMembersForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{team_id}/members][%d] getTeamMembersForbidden %s", 403, payload)
}

func (o *GetTeamMembersForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersNotFound creates a GetTeamMembersNotFound with default headers values
func NewGetTeamMembersNotFound() *GetTeamMembersNotFound {
	return &GetTeamMembersNotFound{}
}

/*
GetTeamMembersNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetTeamMembersNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team members not found response has a 2xx status code
func (o *GetTeamMembersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members not found response has a 3xx status code
func (o *GetTeamMembersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members not found response has a 4xx status code
func (o *GetTeamMembersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team members not found response has a 5xx status code
func (o *GetTeamMembersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get team members not found response a status code equal to that given
func (o *GetTeamMembersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get team members not found response
func (o *GetTeamMembersNotFound) Code() int {
	return 404
}

func (o *GetTeamMembersNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{team_id}/members][%d] getTeamMembersNotFound %s", 404, payload)
}

func (o *GetTeamMembersNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{team_id}/members][%d] getTeamMembersNotFound %s", 404, payload)
}

func (o *GetTeamMembersNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamMembersInternalServerError creates a GetTeamMembersInternalServerError with default headers values
func NewGetTeamMembersInternalServerError() *GetTeamMembersInternalServerError {
	return &GetTeamMembersInternalServerError{}
}

/*
GetTeamMembersInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetTeamMembersInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team members internal server error response has a 2xx status code
func (o *GetTeamMembersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team members internal server error response has a 3xx status code
func (o *GetTeamMembersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team members internal server error response has a 4xx status code
func (o *GetTeamMembersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team members internal server error response has a 5xx status code
func (o *GetTeamMembersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get team members internal server error response a status code equal to that given
func (o *GetTeamMembersInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get team members internal server error response
func (o *GetTeamMembersInternalServerError) Code() int {
	return 500
}

func (o *GetTeamMembersInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{team_id}/members][%d] getTeamMembersInternalServerError %s", 500, payload)
}

func (o *GetTeamMembersInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /teams/{team_id}/members][%d] getTeamMembersInternalServerError %s", 500, payload)
}

func (o *GetTeamMembersInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
