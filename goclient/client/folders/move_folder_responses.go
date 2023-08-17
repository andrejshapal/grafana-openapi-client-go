// Code generated by go-swagger; DO NOT EDIT.

package folders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/goclient/models"
)

// MoveFolderReader is a Reader for the MoveFolder structure.
type MoveFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MoveFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMoveFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewMoveFolderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMoveFolderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMoveFolderNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMoveFolderInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /folders/{folder_uid}/move] moveFolder", response, response.Code())
	}
}

// NewMoveFolderOK creates a MoveFolderOK with default headers values
func NewMoveFolderOK() *MoveFolderOK {
	return &MoveFolderOK{}
}

/*
MoveFolderOK describes a response with status code 200, with default header values.

(empty)
*/
type MoveFolderOK struct {
	Payload *models.Folder
}

// IsSuccess returns true when this move folder o k response has a 2xx status code
func (o *MoveFolderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this move folder o k response has a 3xx status code
func (o *MoveFolderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move folder o k response has a 4xx status code
func (o *MoveFolderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this move folder o k response has a 5xx status code
func (o *MoveFolderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this move folder o k response a status code equal to that given
func (o *MoveFolderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the move folder o k response
func (o *MoveFolderOK) Code() int {
	return 200
}

func (o *MoveFolderOK) Error() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/move][%d] moveFolderOK  %+v", 200, o.Payload)
}

func (o *MoveFolderOK) String() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/move][%d] moveFolderOK  %+v", 200, o.Payload)
}

func (o *MoveFolderOK) GetPayload() *models.Folder {
	return o.Payload
}

func (o *MoveFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Folder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveFolderUnauthorized creates a MoveFolderUnauthorized with default headers values
func NewMoveFolderUnauthorized() *MoveFolderUnauthorized {
	return &MoveFolderUnauthorized{}
}

/*
MoveFolderUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type MoveFolderUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this move folder unauthorized response has a 2xx status code
func (o *MoveFolderUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this move folder unauthorized response has a 3xx status code
func (o *MoveFolderUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move folder unauthorized response has a 4xx status code
func (o *MoveFolderUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this move folder unauthorized response has a 5xx status code
func (o *MoveFolderUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this move folder unauthorized response a status code equal to that given
func (o *MoveFolderUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the move folder unauthorized response
func (o *MoveFolderUnauthorized) Code() int {
	return 401
}

func (o *MoveFolderUnauthorized) Error() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/move][%d] moveFolderUnauthorized  %+v", 401, o.Payload)
}

func (o *MoveFolderUnauthorized) String() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/move][%d] moveFolderUnauthorized  %+v", 401, o.Payload)
}

func (o *MoveFolderUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *MoveFolderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveFolderForbidden creates a MoveFolderForbidden with default headers values
func NewMoveFolderForbidden() *MoveFolderForbidden {
	return &MoveFolderForbidden{}
}

/*
MoveFolderForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type MoveFolderForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this move folder forbidden response has a 2xx status code
func (o *MoveFolderForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this move folder forbidden response has a 3xx status code
func (o *MoveFolderForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move folder forbidden response has a 4xx status code
func (o *MoveFolderForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this move folder forbidden response has a 5xx status code
func (o *MoveFolderForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this move folder forbidden response a status code equal to that given
func (o *MoveFolderForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the move folder forbidden response
func (o *MoveFolderForbidden) Code() int {
	return 403
}

func (o *MoveFolderForbidden) Error() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/move][%d] moveFolderForbidden  %+v", 403, o.Payload)
}

func (o *MoveFolderForbidden) String() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/move][%d] moveFolderForbidden  %+v", 403, o.Payload)
}

func (o *MoveFolderForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *MoveFolderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveFolderNotFound creates a MoveFolderNotFound with default headers values
func NewMoveFolderNotFound() *MoveFolderNotFound {
	return &MoveFolderNotFound{}
}

/*
MoveFolderNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type MoveFolderNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this move folder not found response has a 2xx status code
func (o *MoveFolderNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this move folder not found response has a 3xx status code
func (o *MoveFolderNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move folder not found response has a 4xx status code
func (o *MoveFolderNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this move folder not found response has a 5xx status code
func (o *MoveFolderNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this move folder not found response a status code equal to that given
func (o *MoveFolderNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the move folder not found response
func (o *MoveFolderNotFound) Code() int {
	return 404
}

func (o *MoveFolderNotFound) Error() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/move][%d] moveFolderNotFound  %+v", 404, o.Payload)
}

func (o *MoveFolderNotFound) String() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/move][%d] moveFolderNotFound  %+v", 404, o.Payload)
}

func (o *MoveFolderNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *MoveFolderNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveFolderInternalServerError creates a MoveFolderInternalServerError with default headers values
func NewMoveFolderInternalServerError() *MoveFolderInternalServerError {
	return &MoveFolderInternalServerError{}
}

/*
MoveFolderInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type MoveFolderInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this move folder internal server error response has a 2xx status code
func (o *MoveFolderInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this move folder internal server error response has a 3xx status code
func (o *MoveFolderInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move folder internal server error response has a 4xx status code
func (o *MoveFolderInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this move folder internal server error response has a 5xx status code
func (o *MoveFolderInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this move folder internal server error response a status code equal to that given
func (o *MoveFolderInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the move folder internal server error response
func (o *MoveFolderInternalServerError) Code() int {
	return 500
}

func (o *MoveFolderInternalServerError) Error() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/move][%d] moveFolderInternalServerError  %+v", 500, o.Payload)
}

func (o *MoveFolderInternalServerError) String() string {
	return fmt.Sprintf("[POST /folders/{folder_uid}/move][%d] moveFolderInternalServerError  %+v", 500, o.Payload)
}

func (o *MoveFolderInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *MoveFolderInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
