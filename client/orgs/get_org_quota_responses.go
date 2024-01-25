// Code generated by go-swagger; DO NOT EDIT.

package orgs

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

// GetOrgQuotaReader is a Reader for the GetOrgQuota structure.
type GetOrgQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOrgQuotaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgQuotaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgQuotaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgQuotaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /orgs/{org_id}/quotas] getOrgQuota", response, response.Code())
	}
}

// NewGetOrgQuotaOK creates a GetOrgQuotaOK with default headers values
func NewGetOrgQuotaOK() *GetOrgQuotaOK {
	return &GetOrgQuotaOK{}
}

/*
GetOrgQuotaOK describes a response with status code 200, with default header values.

(empty)
*/
type GetOrgQuotaOK struct {
	Payload []*models.QuotaDTO
}

// IsSuccess returns true when this get org quota Ok response has a 2xx status code
func (o *GetOrgQuotaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get org quota Ok response has a 3xx status code
func (o *GetOrgQuotaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org quota Ok response has a 4xx status code
func (o *GetOrgQuotaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org quota Ok response has a 5xx status code
func (o *GetOrgQuotaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get org quota Ok response a status code equal to that given
func (o *GetOrgQuotaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get org quota Ok response
func (o *GetOrgQuotaOK) Code() int {
	return 200
}

func (o *GetOrgQuotaOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaOk %s", 200, payload)
}

func (o *GetOrgQuotaOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaOk %s", 200, payload)
}

func (o *GetOrgQuotaOK) GetPayload() []*models.QuotaDTO {
	return o.Payload
}

func (o *GetOrgQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgQuotaUnauthorized creates a GetOrgQuotaUnauthorized with default headers values
func NewGetOrgQuotaUnauthorized() *GetOrgQuotaUnauthorized {
	return &GetOrgQuotaUnauthorized{}
}

/*
GetOrgQuotaUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type GetOrgQuotaUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org quota unauthorized response has a 2xx status code
func (o *GetOrgQuotaUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org quota unauthorized response has a 3xx status code
func (o *GetOrgQuotaUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org quota unauthorized response has a 4xx status code
func (o *GetOrgQuotaUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org quota unauthorized response has a 5xx status code
func (o *GetOrgQuotaUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get org quota unauthorized response a status code equal to that given
func (o *GetOrgQuotaUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get org quota unauthorized response
func (o *GetOrgQuotaUnauthorized) Code() int {
	return 401
}

func (o *GetOrgQuotaUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaUnauthorized %s", 401, payload)
}

func (o *GetOrgQuotaUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaUnauthorized %s", 401, payload)
}

func (o *GetOrgQuotaUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgQuotaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgQuotaForbidden creates a GetOrgQuotaForbidden with default headers values
func NewGetOrgQuotaForbidden() *GetOrgQuotaForbidden {
	return &GetOrgQuotaForbidden{}
}

/*
GetOrgQuotaForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type GetOrgQuotaForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org quota forbidden response has a 2xx status code
func (o *GetOrgQuotaForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org quota forbidden response has a 3xx status code
func (o *GetOrgQuotaForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org quota forbidden response has a 4xx status code
func (o *GetOrgQuotaForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org quota forbidden response has a 5xx status code
func (o *GetOrgQuotaForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get org quota forbidden response a status code equal to that given
func (o *GetOrgQuotaForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get org quota forbidden response
func (o *GetOrgQuotaForbidden) Code() int {
	return 403
}

func (o *GetOrgQuotaForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaForbidden %s", 403, payload)
}

func (o *GetOrgQuotaForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaForbidden %s", 403, payload)
}

func (o *GetOrgQuotaForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgQuotaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgQuotaNotFound creates a GetOrgQuotaNotFound with default headers values
func NewGetOrgQuotaNotFound() *GetOrgQuotaNotFound {
	return &GetOrgQuotaNotFound{}
}

/*
GetOrgQuotaNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type GetOrgQuotaNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org quota not found response has a 2xx status code
func (o *GetOrgQuotaNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org quota not found response has a 3xx status code
func (o *GetOrgQuotaNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org quota not found response has a 4xx status code
func (o *GetOrgQuotaNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get org quota not found response has a 5xx status code
func (o *GetOrgQuotaNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get org quota not found response a status code equal to that given
func (o *GetOrgQuotaNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get org quota not found response
func (o *GetOrgQuotaNotFound) Code() int {
	return 404
}

func (o *GetOrgQuotaNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaNotFound %s", 404, payload)
}

func (o *GetOrgQuotaNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaNotFound %s", 404, payload)
}

func (o *GetOrgQuotaNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgQuotaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgQuotaInternalServerError creates a GetOrgQuotaInternalServerError with default headers values
func NewGetOrgQuotaInternalServerError() *GetOrgQuotaInternalServerError {
	return &GetOrgQuotaInternalServerError{}
}

/*
GetOrgQuotaInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type GetOrgQuotaInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this get org quota internal server error response has a 2xx status code
func (o *GetOrgQuotaInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get org quota internal server error response has a 3xx status code
func (o *GetOrgQuotaInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get org quota internal server error response has a 4xx status code
func (o *GetOrgQuotaInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get org quota internal server error response has a 5xx status code
func (o *GetOrgQuotaInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get org quota internal server error response a status code equal to that given
func (o *GetOrgQuotaInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get org quota internal server error response
func (o *GetOrgQuotaInternalServerError) Code() int {
	return 500
}

func (o *GetOrgQuotaInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaInternalServerError %s", 500, payload)
}

func (o *GetOrgQuotaInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /orgs/{org_id}/quotas][%d] getOrgQuotaInternalServerError %s", 500, payload)
}

func (o *GetOrgQuotaInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *GetOrgQuotaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
