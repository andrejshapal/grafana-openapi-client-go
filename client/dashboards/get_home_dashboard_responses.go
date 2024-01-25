// Code generated by go-swagger; DO NOT EDIT.

package dashboards

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

// GetHomeDashboardReader is a Reader for the GetHomeDashboard structure.
type GetHomeDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHomeDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHomeDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetHomeDashboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetHomeDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /dashboards/home] getHomeDashboard", response, response.Code())
	}
}

// NewGetHomeDashboardOK creates a GetHomeDashboardOK with default headers values
func NewGetHomeDashboardOK() *GetHomeDashboardOK {
	return &GetHomeDashboardOK{}
}

/*
GetHomeDashboardOK describes a response with status code 200, with default header values.

(empty)
*/
type GetHomeDashboardOK struct {
	Payload *models.GetHomeDashboardResponse
}

// IsSuccess returns true when this get home dashboard Ok response has a 2xx status code
func (o *GetHomeDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get home dashboard Ok response has a 3xx status code
func (o *GetHomeDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get home dashboard Ok response has a 4xx status code
func (o *GetHomeDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get home dashboard Ok response has a 5xx status code
func (o *GetHomeDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get home dashboard Ok response a status code equal to that given
func (o *GetHomeDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get home dashboard Ok response
func (o *GetHomeDashboardOK) Code() int {
	return 200
}

func (o *GetHomeDashboardOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/home][%d] getHomeDashboardOk %s", 200, payload)
}

func (o *GetHomeDashboardOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/home][%d] getHomeDashboardOk %s", 200, payload)
}

func (o *GetHomeDashboardOK) GetPayload() *models.GetHomeDashboardResponse {
	return o.Payload
}

func (o *GetHomeDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetHomeDashboardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHomeDashboardUnauthorized creates a GetHomeDashboardUnauthorized with default headers values
func NewGetHomeDashboardUnauthorized() *GetHomeDashboardUnauthorized {
	return &GetHomeDashboardUnauthorized{}
}

/*
GetHomeDashboardUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetHomeDashboardUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get home dashboard unauthorized response has a 2xx status code
func (o *GetHomeDashboardUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get home dashboard unauthorized response has a 3xx status code
func (o *GetHomeDashboardUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get home dashboard unauthorized response has a 4xx status code
func (o *GetHomeDashboardUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get home dashboard unauthorized response has a 5xx status code
func (o *GetHomeDashboardUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get home dashboard unauthorized response a status code equal to that given
func (o *GetHomeDashboardUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get home dashboard unauthorized response
func (o *GetHomeDashboardUnauthorized) Code() int {
	return 401
}

func (o *GetHomeDashboardUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/home][%d] getHomeDashboardUnauthorized %s", 401, payload)
}

func (o *GetHomeDashboardUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/home][%d] getHomeDashboardUnauthorized %s", 401, payload)
}

func (o *GetHomeDashboardUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetHomeDashboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHomeDashboardInternalServerError creates a GetHomeDashboardInternalServerError with default headers values
func NewGetHomeDashboardInternalServerError() *GetHomeDashboardInternalServerError {
	return &GetHomeDashboardInternalServerError{}
}

/*
GetHomeDashboardInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetHomeDashboardInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get home dashboard internal server error response has a 2xx status code
func (o *GetHomeDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get home dashboard internal server error response has a 3xx status code
func (o *GetHomeDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get home dashboard internal server error response has a 4xx status code
func (o *GetHomeDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get home dashboard internal server error response has a 5xx status code
func (o *GetHomeDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get home dashboard internal server error response a status code equal to that given
func (o *GetHomeDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get home dashboard internal server error response
func (o *GetHomeDashboardInternalServerError) Code() int {
	return 500
}

func (o *GetHomeDashboardInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/home][%d] getHomeDashboardInternalServerError %s", 500, payload)
}

func (o *GetHomeDashboardInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /dashboards/home][%d] getHomeDashboardInternalServerError %s", 500, payload)
}

func (o *GetHomeDashboardInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetHomeDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
