// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// GetReportsReader is a Reader for the GetReports structure.
type GetReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReportsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReportsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReportsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /reports] getReports", response, response.Code())
	}
}

// NewGetReportsOK creates a GetReportsOK with default headers values
func NewGetReportsOK() *GetReportsOK {
	return &GetReportsOK{}
}

/*
GetReportsOK describes a response with status code 200, with default header values.

(empty)
*/
type GetReportsOK struct {
	Payload []*models.Report
}

// IsSuccess returns true when this get reports Ok response has a 2xx status code
func (o *GetReportsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get reports Ok response has a 3xx status code
func (o *GetReportsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get reports Ok response has a 4xx status code
func (o *GetReportsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get reports Ok response has a 5xx status code
func (o *GetReportsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get reports Ok response a status code equal to that given
func (o *GetReportsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get reports Ok response
func (o *GetReportsOK) Code() int {
	return 200
}

func (o *GetReportsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports][%d] getReportsOk %s", 200, payload)
}

func (o *GetReportsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports][%d] getReportsOk %s", 200, payload)
}

func (o *GetReportsOK) GetPayload() []*models.Report {
	return o.Payload
}

func (o *GetReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsUnauthorized creates a GetReportsUnauthorized with default headers values
func NewGetReportsUnauthorized() *GetReportsUnauthorized {
	return &GetReportsUnauthorized{}
}

/*
GetReportsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetReportsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get reports unauthorized response has a 2xx status code
func (o *GetReportsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get reports unauthorized response has a 3xx status code
func (o *GetReportsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get reports unauthorized response has a 4xx status code
func (o *GetReportsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get reports unauthorized response has a 5xx status code
func (o *GetReportsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get reports unauthorized response a status code equal to that given
func (o *GetReportsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get reports unauthorized response
func (o *GetReportsUnauthorized) Code() int {
	return 401
}

func (o *GetReportsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports][%d] getReportsUnauthorized %s", 401, payload)
}

func (o *GetReportsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports][%d] getReportsUnauthorized %s", 401, payload)
}

func (o *GetReportsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetReportsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsForbidden creates a GetReportsForbidden with default headers values
func NewGetReportsForbidden() *GetReportsForbidden {
	return &GetReportsForbidden{}
}

/*
GetReportsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetReportsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get reports forbidden response has a 2xx status code
func (o *GetReportsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get reports forbidden response has a 3xx status code
func (o *GetReportsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get reports forbidden response has a 4xx status code
func (o *GetReportsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get reports forbidden response has a 5xx status code
func (o *GetReportsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get reports forbidden response a status code equal to that given
func (o *GetReportsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get reports forbidden response
func (o *GetReportsForbidden) Code() int {
	return 403
}

func (o *GetReportsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports][%d] getReportsForbidden %s", 403, payload)
}

func (o *GetReportsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports][%d] getReportsForbidden %s", 403, payload)
}

func (o *GetReportsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetReportsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportsInternalServerError creates a GetReportsInternalServerError with default headers values
func NewGetReportsInternalServerError() *GetReportsInternalServerError {
	return &GetReportsInternalServerError{}
}

/*
GetReportsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetReportsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get reports internal server error response has a 2xx status code
func (o *GetReportsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get reports internal server error response has a 3xx status code
func (o *GetReportsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get reports internal server error response has a 4xx status code
func (o *GetReportsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get reports internal server error response has a 5xx status code
func (o *GetReportsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get reports internal server error response a status code equal to that given
func (o *GetReportsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get reports internal server error response
func (o *GetReportsInternalServerError) Code() int {
	return 500
}

func (o *GetReportsInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports][%d] getReportsInternalServerError %s", 500, payload)
}

func (o *GetReportsInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /reports][%d] getReportsInternalServerError %s", 500, payload)
}

func (o *GetReportsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetReportsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
