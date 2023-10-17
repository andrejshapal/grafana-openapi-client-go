// Code generated by go-swagger; DO NOT EDIT.

package dashboards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// ImportDashboardReader is a Reader for the ImportDashboard structure.
type ImportDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImportDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImportDashboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewImportDashboardPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewImportDashboardUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImportDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /dashboards/import] importDashboard", response, response.Code())
	}
}

// NewImportDashboardOK creates a ImportDashboardOK with default headers values
func NewImportDashboardOK() *ImportDashboardOK {
	return &ImportDashboardOK{}
}

/*
ImportDashboardOK describes a response with status code 200, with default header values.

(empty)
*/
type ImportDashboardOK struct {
	Payload *models.ImportDashboardResponse
}

// IsSuccess returns true when this import dashboard o k response has a 2xx status code
func (o *ImportDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this import dashboard o k response has a 3xx status code
func (o *ImportDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import dashboard o k response has a 4xx status code
func (o *ImportDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this import dashboard o k response has a 5xx status code
func (o *ImportDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this import dashboard o k response a status code equal to that given
func (o *ImportDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the import dashboard o k response
func (o *ImportDashboardOK) Code() int {
	return 200
}

func (o *ImportDashboardOK) Error() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardOK  %+v", 200, o.Payload)
}

func (o *ImportDashboardOK) String() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardOK  %+v", 200, o.Payload)
}

func (o *ImportDashboardOK) GetPayload() *models.ImportDashboardResponse {
	return o.Payload
}

func (o *ImportDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImportDashboardResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportDashboardBadRequest creates a ImportDashboardBadRequest with default headers values
func NewImportDashboardBadRequest() *ImportDashboardBadRequest {
	return &ImportDashboardBadRequest{}
}

/*
ImportDashboardBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type ImportDashboardBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this import dashboard bad request response has a 2xx status code
func (o *ImportDashboardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import dashboard bad request response has a 3xx status code
func (o *ImportDashboardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import dashboard bad request response has a 4xx status code
func (o *ImportDashboardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this import dashboard bad request response has a 5xx status code
func (o *ImportDashboardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this import dashboard bad request response a status code equal to that given
func (o *ImportDashboardBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the import dashboard bad request response
func (o *ImportDashboardBadRequest) Code() int {
	return 400
}

func (o *ImportDashboardBadRequest) Error() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *ImportDashboardBadRequest) String() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *ImportDashboardBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ImportDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportDashboardUnauthorized creates a ImportDashboardUnauthorized with default headers values
func NewImportDashboardUnauthorized() *ImportDashboardUnauthorized {
	return &ImportDashboardUnauthorized{}
}

/*
ImportDashboardUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type ImportDashboardUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this import dashboard unauthorized response has a 2xx status code
func (o *ImportDashboardUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import dashboard unauthorized response has a 3xx status code
func (o *ImportDashboardUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import dashboard unauthorized response has a 4xx status code
func (o *ImportDashboardUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this import dashboard unauthorized response has a 5xx status code
func (o *ImportDashboardUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this import dashboard unauthorized response a status code equal to that given
func (o *ImportDashboardUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the import dashboard unauthorized response
func (o *ImportDashboardUnauthorized) Code() int {
	return 401
}

func (o *ImportDashboardUnauthorized) Error() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardUnauthorized  %+v", 401, o.Payload)
}

func (o *ImportDashboardUnauthorized) String() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardUnauthorized  %+v", 401, o.Payload)
}

func (o *ImportDashboardUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ImportDashboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportDashboardPreconditionFailed creates a ImportDashboardPreconditionFailed with default headers values
func NewImportDashboardPreconditionFailed() *ImportDashboardPreconditionFailed {
	return &ImportDashboardPreconditionFailed{}
}

/*
ImportDashboardPreconditionFailed describes a response with status code 412, with default header values.

PreconditionFailedError
*/
type ImportDashboardPreconditionFailed struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this import dashboard precondition failed response has a 2xx status code
func (o *ImportDashboardPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import dashboard precondition failed response has a 3xx status code
func (o *ImportDashboardPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import dashboard precondition failed response has a 4xx status code
func (o *ImportDashboardPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this import dashboard precondition failed response has a 5xx status code
func (o *ImportDashboardPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this import dashboard precondition failed response a status code equal to that given
func (o *ImportDashboardPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the import dashboard precondition failed response
func (o *ImportDashboardPreconditionFailed) Code() int {
	return 412
}

func (o *ImportDashboardPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardPreconditionFailed  %+v", 412, o.Payload)
}

func (o *ImportDashboardPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardPreconditionFailed  %+v", 412, o.Payload)
}

func (o *ImportDashboardPreconditionFailed) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ImportDashboardPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportDashboardUnprocessableEntity creates a ImportDashboardUnprocessableEntity with default headers values
func NewImportDashboardUnprocessableEntity() *ImportDashboardUnprocessableEntity {
	return &ImportDashboardUnprocessableEntity{}
}

/*
ImportDashboardUnprocessableEntity describes a response with status code 422, with default header values.

UnprocessableEntityError
*/
type ImportDashboardUnprocessableEntity struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this import dashboard unprocessable entity response has a 2xx status code
func (o *ImportDashboardUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import dashboard unprocessable entity response has a 3xx status code
func (o *ImportDashboardUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import dashboard unprocessable entity response has a 4xx status code
func (o *ImportDashboardUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this import dashboard unprocessable entity response has a 5xx status code
func (o *ImportDashboardUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this import dashboard unprocessable entity response a status code equal to that given
func (o *ImportDashboardUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the import dashboard unprocessable entity response
func (o *ImportDashboardUnprocessableEntity) Code() int {
	return 422
}

func (o *ImportDashboardUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ImportDashboardUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ImportDashboardUnprocessableEntity) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ImportDashboardUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImportDashboardInternalServerError creates a ImportDashboardInternalServerError with default headers values
func NewImportDashboardInternalServerError() *ImportDashboardInternalServerError {
	return &ImportDashboardInternalServerError{}
}

/*
ImportDashboardInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type ImportDashboardInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this import dashboard internal server error response has a 2xx status code
func (o *ImportDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this import dashboard internal server error response has a 3xx status code
func (o *ImportDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this import dashboard internal server error response has a 4xx status code
func (o *ImportDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this import dashboard internal server error response has a 5xx status code
func (o *ImportDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this import dashboard internal server error response a status code equal to that given
func (o *ImportDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the import dashboard internal server error response
func (o *ImportDashboardInternalServerError) Code() int {
	return 500
}

func (o *ImportDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *ImportDashboardInternalServerError) String() string {
	return fmt.Sprintf("[POST /dashboards/import][%d] importDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *ImportDashboardInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ImportDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}