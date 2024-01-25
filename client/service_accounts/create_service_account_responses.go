// Code generated by go-swagger; DO NOT EDIT.

package service_accounts

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

// CreateServiceAccountReader is a Reader for the CreateServiceAccount structure.
type CreateServiceAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateServiceAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateServiceAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateServiceAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateServiceAccountUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateServiceAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateServiceAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /serviceaccounts] createServiceAccount", response, response.Code())
	}
}

// NewCreateServiceAccountCreated creates a CreateServiceAccountCreated with default headers values
func NewCreateServiceAccountCreated() *CreateServiceAccountCreated {
	return &CreateServiceAccountCreated{}
}

/*
CreateServiceAccountCreated describes a response with status code 201, with default header values.

(empty)
*/
type CreateServiceAccountCreated struct {
	Payload *models.ServiceAccountDTO
}

// IsSuccess returns true when this create service account created response has a 2xx status code
func (o *CreateServiceAccountCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create service account created response has a 3xx status code
func (o *CreateServiceAccountCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account created response has a 4xx status code
func (o *CreateServiceAccountCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create service account created response has a 5xx status code
func (o *CreateServiceAccountCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create service account created response a status code equal to that given
func (o *CreateServiceAccountCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create service account created response
func (o *CreateServiceAccountCreated) Code() int {
	return 201
}

func (o *CreateServiceAccountCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountCreated %s", 201, payload)
}

func (o *CreateServiceAccountCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountCreated %s", 201, payload)
}

func (o *CreateServiceAccountCreated) GetPayload() *models.ServiceAccountDTO {
	return o.Payload
}

func (o *CreateServiceAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceAccountDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountBadRequest creates a CreateServiceAccountBadRequest with default headers values
func NewCreateServiceAccountBadRequest() *CreateServiceAccountBadRequest {
	return &CreateServiceAccountBadRequest{}
}

/*
CreateServiceAccountBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type CreateServiceAccountBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create service account bad request response has a 2xx status code
func (o *CreateServiceAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service account bad request response has a 3xx status code
func (o *CreateServiceAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account bad request response has a 4xx status code
func (o *CreateServiceAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create service account bad request response has a 5xx status code
func (o *CreateServiceAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create service account bad request response a status code equal to that given
func (o *CreateServiceAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create service account bad request response
func (o *CreateServiceAccountBadRequest) Code() int {
	return 400
}

func (o *CreateServiceAccountBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountBadRequest %s", 400, payload)
}

func (o *CreateServiceAccountBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountBadRequest %s", 400, payload)
}

func (o *CreateServiceAccountBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateServiceAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountUnauthorized creates a CreateServiceAccountUnauthorized with default headers values
func NewCreateServiceAccountUnauthorized() *CreateServiceAccountUnauthorized {
	return &CreateServiceAccountUnauthorized{}
}

/*
CreateServiceAccountUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type CreateServiceAccountUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create service account unauthorized response has a 2xx status code
func (o *CreateServiceAccountUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service account unauthorized response has a 3xx status code
func (o *CreateServiceAccountUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account unauthorized response has a 4xx status code
func (o *CreateServiceAccountUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create service account unauthorized response has a 5xx status code
func (o *CreateServiceAccountUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create service account unauthorized response a status code equal to that given
func (o *CreateServiceAccountUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create service account unauthorized response
func (o *CreateServiceAccountUnauthorized) Code() int {
	return 401
}

func (o *CreateServiceAccountUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountUnauthorized %s", 401, payload)
}

func (o *CreateServiceAccountUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountUnauthorized %s", 401, payload)
}

func (o *CreateServiceAccountUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateServiceAccountUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountForbidden creates a CreateServiceAccountForbidden with default headers values
func NewCreateServiceAccountForbidden() *CreateServiceAccountForbidden {
	return &CreateServiceAccountForbidden{}
}

/*
CreateServiceAccountForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type CreateServiceAccountForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create service account forbidden response has a 2xx status code
func (o *CreateServiceAccountForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service account forbidden response has a 3xx status code
func (o *CreateServiceAccountForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account forbidden response has a 4xx status code
func (o *CreateServiceAccountForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create service account forbidden response has a 5xx status code
func (o *CreateServiceAccountForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create service account forbidden response a status code equal to that given
func (o *CreateServiceAccountForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create service account forbidden response
func (o *CreateServiceAccountForbidden) Code() int {
	return 403
}

func (o *CreateServiceAccountForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountForbidden %s", 403, payload)
}

func (o *CreateServiceAccountForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountForbidden %s", 403, payload)
}

func (o *CreateServiceAccountForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateServiceAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateServiceAccountInternalServerError creates a CreateServiceAccountInternalServerError with default headers values
func NewCreateServiceAccountInternalServerError() *CreateServiceAccountInternalServerError {
	return &CreateServiceAccountInternalServerError{}
}

/*
CreateServiceAccountInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type CreateServiceAccountInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this create service account internal server error response has a 2xx status code
func (o *CreateServiceAccountInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create service account internal server error response has a 3xx status code
func (o *CreateServiceAccountInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create service account internal server error response has a 4xx status code
func (o *CreateServiceAccountInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create service account internal server error response has a 5xx status code
func (o *CreateServiceAccountInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create service account internal server error response a status code equal to that given
func (o *CreateServiceAccountInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create service account internal server error response
func (o *CreateServiceAccountInternalServerError) Code() int {
	return 500
}

func (o *CreateServiceAccountInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountInternalServerError %s", 500, payload)
}

func (o *CreateServiceAccountInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /serviceaccounts][%d] createServiceAccountInternalServerError %s", 500, payload)
}

func (o *CreateServiceAccountInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *CreateServiceAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
