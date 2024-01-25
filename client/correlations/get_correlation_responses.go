// Code generated by go-swagger; DO NOT EDIT.

package correlations

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

// GetCorrelationReader is a Reader for the GetCorrelation structure.
type GetCorrelationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorrelationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCorrelationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCorrelationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCorrelationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCorrelationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /datasources/uid/{sourceUID}/correlations/{correlationUID}] getCorrelation", response, response.Code())
	}
}

// NewGetCorrelationOK creates a GetCorrelationOK with default headers values
func NewGetCorrelationOK() *GetCorrelationOK {
	return &GetCorrelationOK{}
}

/*
GetCorrelationOK describes a response with status code 200, with default header values.

(empty)
*/
type GetCorrelationOK struct {
	Payload *models.Correlation
}

// IsSuccess returns true when this get correlation Ok response has a 2xx status code
func (o *GetCorrelationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get correlation Ok response has a 3xx status code
func (o *GetCorrelationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get correlation Ok response has a 4xx status code
func (o *GetCorrelationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get correlation Ok response has a 5xx status code
func (o *GetCorrelationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get correlation Ok response a status code equal to that given
func (o *GetCorrelationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get correlation Ok response
func (o *GetCorrelationOK) Code() int {
	return 200
}

func (o *GetCorrelationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] getCorrelationOk %s", 200, payload)
}

func (o *GetCorrelationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] getCorrelationOk %s", 200, payload)
}

func (o *GetCorrelationOK) GetPayload() *models.Correlation {
	return o.Payload
}

func (o *GetCorrelationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Correlation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorrelationUnauthorized creates a GetCorrelationUnauthorized with default headers values
func NewGetCorrelationUnauthorized() *GetCorrelationUnauthorized {
	return &GetCorrelationUnauthorized{}
}

/*
GetCorrelationUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetCorrelationUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get correlation unauthorized response has a 2xx status code
func (o *GetCorrelationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get correlation unauthorized response has a 3xx status code
func (o *GetCorrelationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get correlation unauthorized response has a 4xx status code
func (o *GetCorrelationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get correlation unauthorized response has a 5xx status code
func (o *GetCorrelationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get correlation unauthorized response a status code equal to that given
func (o *GetCorrelationUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get correlation unauthorized response
func (o *GetCorrelationUnauthorized) Code() int {
	return 401
}

func (o *GetCorrelationUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] getCorrelationUnauthorized %s", 401, payload)
}

func (o *GetCorrelationUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] getCorrelationUnauthorized %s", 401, payload)
}

func (o *GetCorrelationUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCorrelationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorrelationNotFound creates a GetCorrelationNotFound with default headers values
func NewGetCorrelationNotFound() *GetCorrelationNotFound {
	return &GetCorrelationNotFound{}
}

/*
GetCorrelationNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetCorrelationNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get correlation not found response has a 2xx status code
func (o *GetCorrelationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get correlation not found response has a 3xx status code
func (o *GetCorrelationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get correlation not found response has a 4xx status code
func (o *GetCorrelationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get correlation not found response has a 5xx status code
func (o *GetCorrelationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get correlation not found response a status code equal to that given
func (o *GetCorrelationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get correlation not found response
func (o *GetCorrelationNotFound) Code() int {
	return 404
}

func (o *GetCorrelationNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] getCorrelationNotFound %s", 404, payload)
}

func (o *GetCorrelationNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] getCorrelationNotFound %s", 404, payload)
}

func (o *GetCorrelationNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCorrelationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorrelationInternalServerError creates a GetCorrelationInternalServerError with default headers values
func NewGetCorrelationInternalServerError() *GetCorrelationInternalServerError {
	return &GetCorrelationInternalServerError{}
}

/*
GetCorrelationInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetCorrelationInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get correlation internal server error response has a 2xx status code
func (o *GetCorrelationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get correlation internal server error response has a 3xx status code
func (o *GetCorrelationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get correlation internal server error response has a 4xx status code
func (o *GetCorrelationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get correlation internal server error response has a 5xx status code
func (o *GetCorrelationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get correlation internal server error response a status code equal to that given
func (o *GetCorrelationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get correlation internal server error response
func (o *GetCorrelationInternalServerError) Code() int {
	return 500
}

func (o *GetCorrelationInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] getCorrelationInternalServerError %s", 500, payload)
}

func (o *GetCorrelationInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /datasources/uid/{sourceUID}/correlations/{correlationUID}][%d] getCorrelationInternalServerError %s", 500, payload)
}

func (o *GetCorrelationInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetCorrelationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
