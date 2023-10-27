// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetTeamPreferencesReader is a Reader for the GetTeamPreferences structure.
type GetTeamPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTeamPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTeamPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTeamPreferencesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTeamPreferencesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /teams/{team_id}/preferences] getTeamPreferences", response, response.Code())
	}
}

// NewGetTeamPreferencesOK creates a GetTeamPreferencesOK with default headers values
func NewGetTeamPreferencesOK() *GetTeamPreferencesOK {
	return &GetTeamPreferencesOK{}
}

/*
GetTeamPreferencesOK describes a response with status code 200, with default header values.

(empty)
*/
type GetTeamPreferencesOK struct {
	Payload *models.Preferences
}

// IsSuccess returns true when this get team preferences Ok response has a 2xx status code
func (o *GetTeamPreferencesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get team preferences Ok response has a 3xx status code
func (o *GetTeamPreferencesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team preferences Ok response has a 4xx status code
func (o *GetTeamPreferencesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team preferences Ok response has a 5xx status code
func (o *GetTeamPreferencesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get team preferences Ok response a status code equal to that given
func (o *GetTeamPreferencesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get team preferences Ok response
func (o *GetTeamPreferencesOK) Code() int {
	return 200
}

func (o *GetTeamPreferencesOK) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/preferences][%d] getTeamPreferencesOk  %+v", 200, o.Payload)
}

func (o *GetTeamPreferencesOK) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}/preferences][%d] getTeamPreferencesOk  %+v", 200, o.Payload)
}

func (o *GetTeamPreferencesOK) GetPayload() *models.Preferences {
	return o.Payload
}

func (o *GetTeamPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Preferences)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamPreferencesUnauthorized creates a GetTeamPreferencesUnauthorized with default headers values
func NewGetTeamPreferencesUnauthorized() *GetTeamPreferencesUnauthorized {
	return &GetTeamPreferencesUnauthorized{}
}

/*
GetTeamPreferencesUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetTeamPreferencesUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team preferences unauthorized response has a 2xx status code
func (o *GetTeamPreferencesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team preferences unauthorized response has a 3xx status code
func (o *GetTeamPreferencesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team preferences unauthorized response has a 4xx status code
func (o *GetTeamPreferencesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get team preferences unauthorized response has a 5xx status code
func (o *GetTeamPreferencesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get team preferences unauthorized response a status code equal to that given
func (o *GetTeamPreferencesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get team preferences unauthorized response
func (o *GetTeamPreferencesUnauthorized) Code() int {
	return 401
}

func (o *GetTeamPreferencesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/preferences][%d] getTeamPreferencesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTeamPreferencesUnauthorized) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}/preferences][%d] getTeamPreferencesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTeamPreferencesUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamPreferencesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamPreferencesInternalServerError creates a GetTeamPreferencesInternalServerError with default headers values
func NewGetTeamPreferencesInternalServerError() *GetTeamPreferencesInternalServerError {
	return &GetTeamPreferencesInternalServerError{}
}

/*
GetTeamPreferencesInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetTeamPreferencesInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get team preferences internal server error response has a 2xx status code
func (o *GetTeamPreferencesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get team preferences internal server error response has a 3xx status code
func (o *GetTeamPreferencesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get team preferences internal server error response has a 4xx status code
func (o *GetTeamPreferencesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get team preferences internal server error response has a 5xx status code
func (o *GetTeamPreferencesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get team preferences internal server error response a status code equal to that given
func (o *GetTeamPreferencesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get team preferences internal server error response
func (o *GetTeamPreferencesInternalServerError) Code() int {
	return 500
}

func (o *GetTeamPreferencesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}/preferences][%d] getTeamPreferencesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTeamPreferencesInternalServerError) String() string {
	return fmt.Sprintf("[GET /teams/{team_id}/preferences][%d] getTeamPreferencesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTeamPreferencesInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetTeamPreferencesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
