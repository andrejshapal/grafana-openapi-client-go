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

// GetAnnotationByIDReader is a Reader for the GetAnnotationByID structure.
type GetAnnotationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnnotationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnnotationByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAnnotationByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnnotationByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /annotations/{annotation_id}] getAnnotationByID", response, response.Code())
	}
}

// NewGetAnnotationByIDOK creates a GetAnnotationByIDOK with default headers values
func NewGetAnnotationByIDOK() *GetAnnotationByIDOK {
	return &GetAnnotationByIDOK{}
}

/*
GetAnnotationByIDOK describes a response with status code 200, with default header values.

(empty)
*/
type GetAnnotationByIDOK struct {
	Payload *models.Annotation
}

// IsSuccess returns true when this get annotation by Id Ok response has a 2xx status code
func (o *GetAnnotationByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get annotation by Id Ok response has a 3xx status code
func (o *GetAnnotationByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get annotation by Id Ok response has a 4xx status code
func (o *GetAnnotationByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get annotation by Id Ok response has a 5xx status code
func (o *GetAnnotationByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get annotation by Id Ok response a status code equal to that given
func (o *GetAnnotationByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get annotation by Id Ok response
func (o *GetAnnotationByIDOK) Code() int {
	return 200
}

func (o *GetAnnotationByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /annotations/{annotation_id}][%d] getAnnotationByIdOk %s", 200, payload)
}

func (o *GetAnnotationByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /annotations/{annotation_id}][%d] getAnnotationByIdOk %s", 200, payload)
}

func (o *GetAnnotationByIDOK) GetPayload() *models.Annotation {
	return o.Payload
}

func (o *GetAnnotationByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Annotation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnnotationByIDUnauthorized creates a GetAnnotationByIDUnauthorized with default headers values
func NewGetAnnotationByIDUnauthorized() *GetAnnotationByIDUnauthorized {
	return &GetAnnotationByIDUnauthorized{}
}

/*
GetAnnotationByIDUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetAnnotationByIDUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get annotation by Id unauthorized response has a 2xx status code
func (o *GetAnnotationByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get annotation by Id unauthorized response has a 3xx status code
func (o *GetAnnotationByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get annotation by Id unauthorized response has a 4xx status code
func (o *GetAnnotationByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get annotation by Id unauthorized response has a 5xx status code
func (o *GetAnnotationByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get annotation by Id unauthorized response a status code equal to that given
func (o *GetAnnotationByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get annotation by Id unauthorized response
func (o *GetAnnotationByIDUnauthorized) Code() int {
	return 401
}

func (o *GetAnnotationByIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /annotations/{annotation_id}][%d] getAnnotationByIdUnauthorized %s", 401, payload)
}

func (o *GetAnnotationByIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /annotations/{annotation_id}][%d] getAnnotationByIdUnauthorized %s", 401, payload)
}

func (o *GetAnnotationByIDUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAnnotationByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnnotationByIDInternalServerError creates a GetAnnotationByIDInternalServerError with default headers values
func NewGetAnnotationByIDInternalServerError() *GetAnnotationByIDInternalServerError {
	return &GetAnnotationByIDInternalServerError{}
}

/*
GetAnnotationByIDInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetAnnotationByIDInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get annotation by Id internal server error response has a 2xx status code
func (o *GetAnnotationByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get annotation by Id internal server error response has a 3xx status code
func (o *GetAnnotationByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get annotation by Id internal server error response has a 4xx status code
func (o *GetAnnotationByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get annotation by Id internal server error response has a 5xx status code
func (o *GetAnnotationByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get annotation by Id internal server error response a status code equal to that given
func (o *GetAnnotationByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get annotation by Id internal server error response
func (o *GetAnnotationByIDInternalServerError) Code() int {
	return 500
}

func (o *GetAnnotationByIDInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /annotations/{annotation_id}][%d] getAnnotationByIdInternalServerError %s", 500, payload)
}

func (o *GetAnnotationByIDInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /annotations/{annotation_id}][%d] getAnnotationByIdInternalServerError %s", 500, payload)
}

func (o *GetAnnotationByIDInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetAnnotationByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
