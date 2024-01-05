// Code generated by go-swagger; DO NOT EDIT.

package dashboard_public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// UpdatePublicDashboardReader is a Reader for the UpdatePublicDashboard structure.
type UpdatePublicDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePublicDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePublicDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePublicDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdatePublicDashboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdatePublicDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdatePublicDashboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /dashboards/uid/{dashboardUid}/public-dashboards/{uid}] updatePublicDashboard", response, response.Code())
	}
}

// NewUpdatePublicDashboardOK creates a UpdatePublicDashboardOK with default headers values
func NewUpdatePublicDashboardOK() *UpdatePublicDashboardOK {
	return &UpdatePublicDashboardOK{}
}

/*
UpdatePublicDashboardOK describes a response with status code 200, with default header values.

(empty)
*/
type UpdatePublicDashboardOK struct {
	Payload *models.PublicDashboard
}

// IsSuccess returns true when this update public dashboard Ok response has a 2xx status code
func (o *UpdatePublicDashboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update public dashboard Ok response has a 3xx status code
func (o *UpdatePublicDashboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update public dashboard Ok response has a 4xx status code
func (o *UpdatePublicDashboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update public dashboard Ok response has a 5xx status code
func (o *UpdatePublicDashboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update public dashboard Ok response a status code equal to that given
func (o *UpdatePublicDashboardOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update public dashboard Ok response
func (o *UpdatePublicDashboardOK) Code() int {
	return 200
}

func (o *UpdatePublicDashboardOK) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/uid/{dashboardUid}/public-dashboards/{uid}][%d] updatePublicDashboardOk  %+v", 200, o.Payload)
}

func (o *UpdatePublicDashboardOK) String() string {
	return fmt.Sprintf("[PATCH /dashboards/uid/{dashboardUid}/public-dashboards/{uid}][%d] updatePublicDashboardOk  %+v", 200, o.Payload)
}

func (o *UpdatePublicDashboardOK) GetPayload() *models.PublicDashboard {
	return o.Payload
}

func (o *UpdatePublicDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicDashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePublicDashboardBadRequest creates a UpdatePublicDashboardBadRequest with default headers values
func NewUpdatePublicDashboardBadRequest() *UpdatePublicDashboardBadRequest {
	return &UpdatePublicDashboardBadRequest{}
}

/*
UpdatePublicDashboardBadRequest describes a response with status code 400, with default header values.

BadRequestPublicError is returned when the request is invalid and it cannot be processed.
*/
type UpdatePublicDashboardBadRequest struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this update public dashboard bad request response has a 2xx status code
func (o *UpdatePublicDashboardBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update public dashboard bad request response has a 3xx status code
func (o *UpdatePublicDashboardBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update public dashboard bad request response has a 4xx status code
func (o *UpdatePublicDashboardBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update public dashboard bad request response has a 5xx status code
func (o *UpdatePublicDashboardBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update public dashboard bad request response a status code equal to that given
func (o *UpdatePublicDashboardBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update public dashboard bad request response
func (o *UpdatePublicDashboardBadRequest) Code() int {
	return 400
}

func (o *UpdatePublicDashboardBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/uid/{dashboardUid}/public-dashboards/{uid}][%d] updatePublicDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdatePublicDashboardBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dashboards/uid/{dashboardUid}/public-dashboards/{uid}][%d] updatePublicDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *UpdatePublicDashboardBadRequest) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *UpdatePublicDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePublicDashboardUnauthorized creates a UpdatePublicDashboardUnauthorized with default headers values
func NewUpdatePublicDashboardUnauthorized() *UpdatePublicDashboardUnauthorized {
	return &UpdatePublicDashboardUnauthorized{}
}

/*
UpdatePublicDashboardUnauthorized describes a response with status code 401, with default header values.

UnauthorisedPublicError is returned when the request is not authenticated.
*/
type UpdatePublicDashboardUnauthorized struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this update public dashboard unauthorized response has a 2xx status code
func (o *UpdatePublicDashboardUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update public dashboard unauthorized response has a 3xx status code
func (o *UpdatePublicDashboardUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update public dashboard unauthorized response has a 4xx status code
func (o *UpdatePublicDashboardUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update public dashboard unauthorized response has a 5xx status code
func (o *UpdatePublicDashboardUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update public dashboard unauthorized response a status code equal to that given
func (o *UpdatePublicDashboardUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update public dashboard unauthorized response
func (o *UpdatePublicDashboardUnauthorized) Code() int {
	return 401
}

func (o *UpdatePublicDashboardUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/uid/{dashboardUid}/public-dashboards/{uid}][%d] updatePublicDashboardUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdatePublicDashboardUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /dashboards/uid/{dashboardUid}/public-dashboards/{uid}][%d] updatePublicDashboardUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdatePublicDashboardUnauthorized) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *UpdatePublicDashboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePublicDashboardForbidden creates a UpdatePublicDashboardForbidden with default headers values
func NewUpdatePublicDashboardForbidden() *UpdatePublicDashboardForbidden {
	return &UpdatePublicDashboardForbidden{}
}

/*
UpdatePublicDashboardForbidden describes a response with status code 403, with default header values.

ForbiddenPublicError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdatePublicDashboardForbidden struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this update public dashboard forbidden response has a 2xx status code
func (o *UpdatePublicDashboardForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update public dashboard forbidden response has a 3xx status code
func (o *UpdatePublicDashboardForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update public dashboard forbidden response has a 4xx status code
func (o *UpdatePublicDashboardForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update public dashboard forbidden response has a 5xx status code
func (o *UpdatePublicDashboardForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update public dashboard forbidden response a status code equal to that given
func (o *UpdatePublicDashboardForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update public dashboard forbidden response
func (o *UpdatePublicDashboardForbidden) Code() int {
	return 403
}

func (o *UpdatePublicDashboardForbidden) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/uid/{dashboardUid}/public-dashboards/{uid}][%d] updatePublicDashboardForbidden  %+v", 403, o.Payload)
}

func (o *UpdatePublicDashboardForbidden) String() string {
	return fmt.Sprintf("[PATCH /dashboards/uid/{dashboardUid}/public-dashboards/{uid}][%d] updatePublicDashboardForbidden  %+v", 403, o.Payload)
}

func (o *UpdatePublicDashboardForbidden) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *UpdatePublicDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePublicDashboardInternalServerError creates a UpdatePublicDashboardInternalServerError with default headers values
func NewUpdatePublicDashboardInternalServerError() *UpdatePublicDashboardInternalServerError {
	return &UpdatePublicDashboardInternalServerError{}
}

/*
UpdatePublicDashboardInternalServerError describes a response with status code 500, with default header values.

InternalServerPublicError is a general error indicating something went wrong internally.
*/
type UpdatePublicDashboardInternalServerError struct {
	Payload *models.PublicError
}

// IsSuccess returns true when this update public dashboard internal server error response has a 2xx status code
func (o *UpdatePublicDashboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update public dashboard internal server error response has a 3xx status code
func (o *UpdatePublicDashboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update public dashboard internal server error response has a 4xx status code
func (o *UpdatePublicDashboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update public dashboard internal server error response has a 5xx status code
func (o *UpdatePublicDashboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update public dashboard internal server error response a status code equal to that given
func (o *UpdatePublicDashboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update public dashboard internal server error response
func (o *UpdatePublicDashboardInternalServerError) Code() int {
	return 500
}

func (o *UpdatePublicDashboardInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /dashboards/uid/{dashboardUid}/public-dashboards/{uid}][%d] updatePublicDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdatePublicDashboardInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /dashboards/uid/{dashboardUid}/public-dashboards/{uid}][%d] updatePublicDashboardInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdatePublicDashboardInternalServerError) GetPayload() *models.PublicError {
	return o.Payload
}

func (o *UpdatePublicDashboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}