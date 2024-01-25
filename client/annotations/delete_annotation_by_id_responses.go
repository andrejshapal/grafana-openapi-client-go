// Code generated by go-swagger; DO NOT EDIT.

package annotations

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

// DeleteAnnotationByIDReader is a Reader for the DeleteAnnotationByID structure.
type DeleteAnnotationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAnnotationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAnnotationByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAnnotationByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAnnotationByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAnnotationByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /annotations/{annotation_id}] deleteAnnotationByID", response, response.Code())
	}
}

// NewDeleteAnnotationByIDOK creates a DeleteAnnotationByIDOK with default headers values
func NewDeleteAnnotationByIDOK() *DeleteAnnotationByIDOK {
	return &DeleteAnnotationByIDOK{}
}

/*
DeleteAnnotationByIDOK describes a response with status code 200, with default header values.

An OKResponse is returned if the request was successful.
*/
type DeleteAnnotationByIDOK struct {
	Payload *models.SuccessResponseBody
}

// IsSuccess returns true when this delete annotation by Id Ok response has a 2xx status code
func (o *DeleteAnnotationByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete annotation by Id Ok response has a 3xx status code
func (o *DeleteAnnotationByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete annotation by Id Ok response has a 4xx status code
func (o *DeleteAnnotationByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete annotation by Id Ok response has a 5xx status code
func (o *DeleteAnnotationByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete annotation by Id Ok response a status code equal to that given
func (o *DeleteAnnotationByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete annotation by Id Ok response
func (o *DeleteAnnotationByIDOK) Code() int {
	return 200
}

func (o *DeleteAnnotationByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /annotations/{annotation_id}][%d] deleteAnnotationByIdOk %s", 200, payload)
}

func (o *DeleteAnnotationByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /annotations/{annotation_id}][%d] deleteAnnotationByIdOk %s", 200, payload)
}

func (o *DeleteAnnotationByIDOK) GetPayload() *models.SuccessResponseBody {
	return o.Payload
}

func (o *DeleteAnnotationByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SuccessResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnnotationByIDUnauthorized creates a DeleteAnnotationByIDUnauthorized with default headers values
func NewDeleteAnnotationByIDUnauthorized() *DeleteAnnotationByIDUnauthorized {
	return &DeleteAnnotationByIDUnauthorized{}
}

/*
DeleteAnnotationByIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DeleteAnnotationByIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete annotation by Id unauthorized response has a 2xx status code
func (o *DeleteAnnotationByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete annotation by Id unauthorized response has a 3xx status code
func (o *DeleteAnnotationByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete annotation by Id unauthorized response has a 4xx status code
func (o *DeleteAnnotationByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete annotation by Id unauthorized response has a 5xx status code
func (o *DeleteAnnotationByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete annotation by Id unauthorized response a status code equal to that given
func (o *DeleteAnnotationByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete annotation by Id unauthorized response
func (o *DeleteAnnotationByIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteAnnotationByIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /annotations/{annotation_id}][%d] deleteAnnotationByIdUnauthorized %s", 401, payload)
}

func (o *DeleteAnnotationByIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /annotations/{annotation_id}][%d] deleteAnnotationByIdUnauthorized %s", 401, payload)
}

func (o *DeleteAnnotationByIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAnnotationByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnnotationByIDForbidden creates a DeleteAnnotationByIDForbidden with default headers values
func NewDeleteAnnotationByIDForbidden() *DeleteAnnotationByIDForbidden {
	return &DeleteAnnotationByIDForbidden{}
}

/*
DeleteAnnotationByIDForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DeleteAnnotationByIDForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete annotation by Id forbidden response has a 2xx status code
func (o *DeleteAnnotationByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete annotation by Id forbidden response has a 3xx status code
func (o *DeleteAnnotationByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete annotation by Id forbidden response has a 4xx status code
func (o *DeleteAnnotationByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete annotation by Id forbidden response has a 5xx status code
func (o *DeleteAnnotationByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete annotation by Id forbidden response a status code equal to that given
func (o *DeleteAnnotationByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete annotation by Id forbidden response
func (o *DeleteAnnotationByIDForbidden) Code() int {
	return 403
}

func (o *DeleteAnnotationByIDForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /annotations/{annotation_id}][%d] deleteAnnotationByIdForbidden %s", 403, payload)
}

func (o *DeleteAnnotationByIDForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /annotations/{annotation_id}][%d] deleteAnnotationByIdForbidden %s", 403, payload)
}

func (o *DeleteAnnotationByIDForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAnnotationByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnnotationByIDInternalServerError creates a DeleteAnnotationByIDInternalServerError with default headers values
func NewDeleteAnnotationByIDInternalServerError() *DeleteAnnotationByIDInternalServerError {
	return &DeleteAnnotationByIDInternalServerError{}
}

/*
DeleteAnnotationByIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DeleteAnnotationByIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this delete annotation by Id internal server error response has a 2xx status code
func (o *DeleteAnnotationByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete annotation by Id internal server error response has a 3xx status code
func (o *DeleteAnnotationByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete annotation by Id internal server error response has a 4xx status code
func (o *DeleteAnnotationByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete annotation by Id internal server error response has a 5xx status code
func (o *DeleteAnnotationByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete annotation by Id internal server error response a status code equal to that given
func (o *DeleteAnnotationByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete annotation by Id internal server error response
func (o *DeleteAnnotationByIDInternalServerError) Code() int {
	return 500
}

func (o *DeleteAnnotationByIDInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /annotations/{annotation_id}][%d] deleteAnnotationByIdInternalServerError %s", 500, payload)
}

func (o *DeleteAnnotationByIDInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /annotations/{annotation_id}][%d] deleteAnnotationByIdInternalServerError %s", 500, payload)
}

func (o *DeleteAnnotationByIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DeleteAnnotationByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
