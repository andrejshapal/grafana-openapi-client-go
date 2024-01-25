// Code generated by go-swagger; DO NOT EDIT.

package legacy_alerts_notification_channels

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

// UpdateAlertNotificationChannelReader is a Reader for the UpdateAlertNotificationChannel structure.
type UpdateAlertNotificationChannelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAlertNotificationChannelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAlertNotificationChannelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateAlertNotificationChannelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAlertNotificationChannelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateAlertNotificationChannelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateAlertNotificationChannelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /alert-notifications/{notification_channel_id}] updateAlertNotificationChannel", response, response.Code())
	}
}

// NewUpdateAlertNotificationChannelOK creates a UpdateAlertNotificationChannelOK with default headers values
func NewUpdateAlertNotificationChannelOK() *UpdateAlertNotificationChannelOK {
	return &UpdateAlertNotificationChannelOK{}
}

/*
UpdateAlertNotificationChannelOK describes a response with status code 200, with default header values.

(empty)
*/
type UpdateAlertNotificationChannelOK struct {
	Payload *models.AlertNotification
}

// IsSuccess returns true when this update alert notification channel Ok response has a 2xx status code
func (o *UpdateAlertNotificationChannelOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update alert notification channel Ok response has a 3xx status code
func (o *UpdateAlertNotificationChannelOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update alert notification channel Ok response has a 4xx status code
func (o *UpdateAlertNotificationChannelOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update alert notification channel Ok response has a 5xx status code
func (o *UpdateAlertNotificationChannelOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update alert notification channel Ok response a status code equal to that given
func (o *UpdateAlertNotificationChannelOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update alert notification channel Ok response
func (o *UpdateAlertNotificationChannelOK) Code() int {
	return 200
}

func (o *UpdateAlertNotificationChannelOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelOk %s", 200, payload)
}

func (o *UpdateAlertNotificationChannelOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelOk %s", 200, payload)
}

func (o *UpdateAlertNotificationChannelOK) GetPayload() *models.AlertNotification {
	return o.Payload
}

func (o *UpdateAlertNotificationChannelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertNotification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAlertNotificationChannelUnauthorized creates a UpdateAlertNotificationChannelUnauthorized with default headers values
func NewUpdateAlertNotificationChannelUnauthorized() *UpdateAlertNotificationChannelUnauthorized {
	return &UpdateAlertNotificationChannelUnauthorized{}
}

/*
UpdateAlertNotificationChannelUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type UpdateAlertNotificationChannelUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update alert notification channel unauthorized response has a 2xx status code
func (o *UpdateAlertNotificationChannelUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update alert notification channel unauthorized response has a 3xx status code
func (o *UpdateAlertNotificationChannelUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update alert notification channel unauthorized response has a 4xx status code
func (o *UpdateAlertNotificationChannelUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update alert notification channel unauthorized response has a 5xx status code
func (o *UpdateAlertNotificationChannelUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update alert notification channel unauthorized response a status code equal to that given
func (o *UpdateAlertNotificationChannelUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update alert notification channel unauthorized response
func (o *UpdateAlertNotificationChannelUnauthorized) Code() int {
	return 401
}

func (o *UpdateAlertNotificationChannelUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelUnauthorized %s", 401, payload)
}

func (o *UpdateAlertNotificationChannelUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelUnauthorized %s", 401, payload)
}

func (o *UpdateAlertNotificationChannelUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateAlertNotificationChannelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAlertNotificationChannelForbidden creates a UpdateAlertNotificationChannelForbidden with default headers values
func NewUpdateAlertNotificationChannelForbidden() *UpdateAlertNotificationChannelForbidden {
	return &UpdateAlertNotificationChannelForbidden{}
}

/*
UpdateAlertNotificationChannelForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type UpdateAlertNotificationChannelForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update alert notification channel forbidden response has a 2xx status code
func (o *UpdateAlertNotificationChannelForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update alert notification channel forbidden response has a 3xx status code
func (o *UpdateAlertNotificationChannelForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update alert notification channel forbidden response has a 4xx status code
func (o *UpdateAlertNotificationChannelForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update alert notification channel forbidden response has a 5xx status code
func (o *UpdateAlertNotificationChannelForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update alert notification channel forbidden response a status code equal to that given
func (o *UpdateAlertNotificationChannelForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update alert notification channel forbidden response
func (o *UpdateAlertNotificationChannelForbidden) Code() int {
	return 403
}

func (o *UpdateAlertNotificationChannelForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelForbidden %s", 403, payload)
}

func (o *UpdateAlertNotificationChannelForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelForbidden %s", 403, payload)
}

func (o *UpdateAlertNotificationChannelForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateAlertNotificationChannelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAlertNotificationChannelNotFound creates a UpdateAlertNotificationChannelNotFound with default headers values
func NewUpdateAlertNotificationChannelNotFound() *UpdateAlertNotificationChannelNotFound {
	return &UpdateAlertNotificationChannelNotFound{}
}

/*
UpdateAlertNotificationChannelNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type UpdateAlertNotificationChannelNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update alert notification channel not found response has a 2xx status code
func (o *UpdateAlertNotificationChannelNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update alert notification channel not found response has a 3xx status code
func (o *UpdateAlertNotificationChannelNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update alert notification channel not found response has a 4xx status code
func (o *UpdateAlertNotificationChannelNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update alert notification channel not found response has a 5xx status code
func (o *UpdateAlertNotificationChannelNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update alert notification channel not found response a status code equal to that given
func (o *UpdateAlertNotificationChannelNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update alert notification channel not found response
func (o *UpdateAlertNotificationChannelNotFound) Code() int {
	return 404
}

func (o *UpdateAlertNotificationChannelNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelNotFound %s", 404, payload)
}

func (o *UpdateAlertNotificationChannelNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelNotFound %s", 404, payload)
}

func (o *UpdateAlertNotificationChannelNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateAlertNotificationChannelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAlertNotificationChannelInternalServerError creates a UpdateAlertNotificationChannelInternalServerError with default headers values
func NewUpdateAlertNotificationChannelInternalServerError() *UpdateAlertNotificationChannelInternalServerError {
	return &UpdateAlertNotificationChannelInternalServerError{}
}

/*
UpdateAlertNotificationChannelInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type UpdateAlertNotificationChannelInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this update alert notification channel internal server error response has a 2xx status code
func (o *UpdateAlertNotificationChannelInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update alert notification channel internal server error response has a 3xx status code
func (o *UpdateAlertNotificationChannelInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update alert notification channel internal server error response has a 4xx status code
func (o *UpdateAlertNotificationChannelInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update alert notification channel internal server error response has a 5xx status code
func (o *UpdateAlertNotificationChannelInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update alert notification channel internal server error response a status code equal to that given
func (o *UpdateAlertNotificationChannelInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update alert notification channel internal server error response
func (o *UpdateAlertNotificationChannelInternalServerError) Code() int {
	return 500
}

func (o *UpdateAlertNotificationChannelInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelInternalServerError %s", 500, payload)
}

func (o *UpdateAlertNotificationChannelInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /alert-notifications/{notification_channel_id}][%d] updateAlertNotificationChannelInternalServerError %s", 500, payload)
}

func (o *UpdateAlertNotificationChannelInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *UpdateAlertNotificationChannelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
