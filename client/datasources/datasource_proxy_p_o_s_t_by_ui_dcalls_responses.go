// Code generated by go-swagger; DO NOT EDIT.

package datasources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// DatasourceProxyPOSTByUIDcallsReader is a Reader for the DatasourceProxyPOSTByUIDcalls structure.
type DatasourceProxyPOSTByUIDcallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DatasourceProxyPOSTByUIDcallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDatasourceProxyPOSTByUIDcallsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDatasourceProxyPOSTByUIDcallsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDatasourceProxyPOSTByUIDcallsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDatasourceProxyPOSTByUIDcallsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDatasourceProxyPOSTByUIDcallsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDatasourceProxyPOSTByUIDcallsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDatasourceProxyPOSTByUIDcallsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}] datasourceProxyPOSTByUIDcalls", response, response.Code())
	}
}

// NewDatasourceProxyPOSTByUIDcallsCreated creates a DatasourceProxyPOSTByUIDcallsCreated with default headers values
func NewDatasourceProxyPOSTByUIDcallsCreated() *DatasourceProxyPOSTByUIDcallsCreated {
	return &DatasourceProxyPOSTByUIDcallsCreated{}
}

/*
DatasourceProxyPOSTByUIDcallsCreated describes a response with status code 201, with default header values.

(empty)
*/
type DatasourceProxyPOSTByUIDcallsCreated struct {
}

// IsSuccess returns true when this datasource proxy p o s t by Ui dcalls created response has a 2xx status code
func (o *DatasourceProxyPOSTByUIDcallsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this datasource proxy p o s t by Ui dcalls created response has a 3xx status code
func (o *DatasourceProxyPOSTByUIDcallsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy p o s t by Ui dcalls created response has a 4xx status code
func (o *DatasourceProxyPOSTByUIDcallsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this datasource proxy p o s t by Ui dcalls created response has a 5xx status code
func (o *DatasourceProxyPOSTByUIDcallsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy p o s t by Ui dcalls created response a status code equal to that given
func (o *DatasourceProxyPOSTByUIDcallsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the datasource proxy p o s t by Ui dcalls created response
func (o *DatasourceProxyPOSTByUIDcallsCreated) Code() int {
	return 201
}

func (o *DatasourceProxyPOSTByUIDcallsCreated) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsCreated ", 201)
}

func (o *DatasourceProxyPOSTByUIDcallsCreated) String() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsCreated ", 201)
}

func (o *DatasourceProxyPOSTByUIDcallsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDatasourceProxyPOSTByUIDcallsAccepted creates a DatasourceProxyPOSTByUIDcallsAccepted with default headers values
func NewDatasourceProxyPOSTByUIDcallsAccepted() *DatasourceProxyPOSTByUIDcallsAccepted {
	return &DatasourceProxyPOSTByUIDcallsAccepted{}
}

/*
DatasourceProxyPOSTByUIDcallsAccepted describes a response with status code 202, with default header values.

(empty)
*/
type DatasourceProxyPOSTByUIDcallsAccepted struct {
}

// IsSuccess returns true when this datasource proxy p o s t by Ui dcalls accepted response has a 2xx status code
func (o *DatasourceProxyPOSTByUIDcallsAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this datasource proxy p o s t by Ui dcalls accepted response has a 3xx status code
func (o *DatasourceProxyPOSTByUIDcallsAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy p o s t by Ui dcalls accepted response has a 4xx status code
func (o *DatasourceProxyPOSTByUIDcallsAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this datasource proxy p o s t by Ui dcalls accepted response has a 5xx status code
func (o *DatasourceProxyPOSTByUIDcallsAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy p o s t by Ui dcalls accepted response a status code equal to that given
func (o *DatasourceProxyPOSTByUIDcallsAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the datasource proxy p o s t by Ui dcalls accepted response
func (o *DatasourceProxyPOSTByUIDcallsAccepted) Code() int {
	return 202
}

func (o *DatasourceProxyPOSTByUIDcallsAccepted) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsAccepted ", 202)
}

func (o *DatasourceProxyPOSTByUIDcallsAccepted) String() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsAccepted ", 202)
}

func (o *DatasourceProxyPOSTByUIDcallsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDatasourceProxyPOSTByUIDcallsBadRequest creates a DatasourceProxyPOSTByUIDcallsBadRequest with default headers values
func NewDatasourceProxyPOSTByUIDcallsBadRequest() *DatasourceProxyPOSTByUIDcallsBadRequest {
	return &DatasourceProxyPOSTByUIDcallsBadRequest{}
}

/*
DatasourceProxyPOSTByUIDcallsBadRequest describes a response with status code 400, with default header values.

BadRequestError is returned when the request is invalid and it cannot be processed.
*/
type DatasourceProxyPOSTByUIDcallsBadRequest struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy p o s t by Ui dcalls bad request response has a 2xx status code
func (o *DatasourceProxyPOSTByUIDcallsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy p o s t by Ui dcalls bad request response has a 3xx status code
func (o *DatasourceProxyPOSTByUIDcallsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy p o s t by Ui dcalls bad request response has a 4xx status code
func (o *DatasourceProxyPOSTByUIDcallsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this datasource proxy p o s t by Ui dcalls bad request response has a 5xx status code
func (o *DatasourceProxyPOSTByUIDcallsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy p o s t by Ui dcalls bad request response a status code equal to that given
func (o *DatasourceProxyPOSTByUIDcallsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the datasource proxy p o s t by Ui dcalls bad request response
func (o *DatasourceProxyPOSTByUIDcallsBadRequest) Code() int {
	return 400
}

func (o *DatasourceProxyPOSTByUIDcallsBadRequest) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsBadRequest  %+v", 400, o.Payload)
}

func (o *DatasourceProxyPOSTByUIDcallsBadRequest) String() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsBadRequest  %+v", 400, o.Payload)
}

func (o *DatasourceProxyPOSTByUIDcallsBadRequest) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyPOSTByUIDcallsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyPOSTByUIDcallsUnauthorized creates a DatasourceProxyPOSTByUIDcallsUnauthorized with default headers values
func NewDatasourceProxyPOSTByUIDcallsUnauthorized() *DatasourceProxyPOSTByUIDcallsUnauthorized {
	return &DatasourceProxyPOSTByUIDcallsUnauthorized{}
}

/*
DatasourceProxyPOSTByUIDcallsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type DatasourceProxyPOSTByUIDcallsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy p o s t by Ui dcalls unauthorized response has a 2xx status code
func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy p o s t by Ui dcalls unauthorized response has a 3xx status code
func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy p o s t by Ui dcalls unauthorized response has a 4xx status code
func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this datasource proxy p o s t by Ui dcalls unauthorized response has a 5xx status code
func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy p o s t by Ui dcalls unauthorized response a status code equal to that given
func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the datasource proxy p o s t by Ui dcalls unauthorized response
func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) Code() int {
	return 401
}

func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsUnauthorized  %+v", 401, o.Payload)
}

func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) String() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsUnauthorized  %+v", 401, o.Payload)
}

func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyPOSTByUIDcallsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyPOSTByUIDcallsForbidden creates a DatasourceProxyPOSTByUIDcallsForbidden with default headers values
func NewDatasourceProxyPOSTByUIDcallsForbidden() *DatasourceProxyPOSTByUIDcallsForbidden {
	return &DatasourceProxyPOSTByUIDcallsForbidden{}
}

/*
DatasourceProxyPOSTByUIDcallsForbidden describes a response with status code 403, with default header values.

ForbiddenError is returned if the user/token has insufficient permissions to access the requested resource.
*/
type DatasourceProxyPOSTByUIDcallsForbidden struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy p o s t by Ui dcalls forbidden response has a 2xx status code
func (o *DatasourceProxyPOSTByUIDcallsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy p o s t by Ui dcalls forbidden response has a 3xx status code
func (o *DatasourceProxyPOSTByUIDcallsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy p o s t by Ui dcalls forbidden response has a 4xx status code
func (o *DatasourceProxyPOSTByUIDcallsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this datasource proxy p o s t by Ui dcalls forbidden response has a 5xx status code
func (o *DatasourceProxyPOSTByUIDcallsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy p o s t by Ui dcalls forbidden response a status code equal to that given
func (o *DatasourceProxyPOSTByUIDcallsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the datasource proxy p o s t by Ui dcalls forbidden response
func (o *DatasourceProxyPOSTByUIDcallsForbidden) Code() int {
	return 403
}

func (o *DatasourceProxyPOSTByUIDcallsForbidden) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsForbidden  %+v", 403, o.Payload)
}

func (o *DatasourceProxyPOSTByUIDcallsForbidden) String() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsForbidden  %+v", 403, o.Payload)
}

func (o *DatasourceProxyPOSTByUIDcallsForbidden) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyPOSTByUIDcallsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyPOSTByUIDcallsNotFound creates a DatasourceProxyPOSTByUIDcallsNotFound with default headers values
func NewDatasourceProxyPOSTByUIDcallsNotFound() *DatasourceProxyPOSTByUIDcallsNotFound {
	return &DatasourceProxyPOSTByUIDcallsNotFound{}
}

/*
DatasourceProxyPOSTByUIDcallsNotFound describes a response with status code 404, with default header values.

NotFoundError is returned when the requested resource was not found.
*/
type DatasourceProxyPOSTByUIDcallsNotFound struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy p o s t by Ui dcalls not found response has a 2xx status code
func (o *DatasourceProxyPOSTByUIDcallsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy p o s t by Ui dcalls not found response has a 3xx status code
func (o *DatasourceProxyPOSTByUIDcallsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy p o s t by Ui dcalls not found response has a 4xx status code
func (o *DatasourceProxyPOSTByUIDcallsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this datasource proxy p o s t by Ui dcalls not found response has a 5xx status code
func (o *DatasourceProxyPOSTByUIDcallsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this datasource proxy p o s t by Ui dcalls not found response a status code equal to that given
func (o *DatasourceProxyPOSTByUIDcallsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the datasource proxy p o s t by Ui dcalls not found response
func (o *DatasourceProxyPOSTByUIDcallsNotFound) Code() int {
	return 404
}

func (o *DatasourceProxyPOSTByUIDcallsNotFound) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsNotFound  %+v", 404, o.Payload)
}

func (o *DatasourceProxyPOSTByUIDcallsNotFound) String() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsNotFound  %+v", 404, o.Payload)
}

func (o *DatasourceProxyPOSTByUIDcallsNotFound) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyPOSTByUIDcallsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDatasourceProxyPOSTByUIDcallsInternalServerError creates a DatasourceProxyPOSTByUIDcallsInternalServerError with default headers values
func NewDatasourceProxyPOSTByUIDcallsInternalServerError() *DatasourceProxyPOSTByUIDcallsInternalServerError {
	return &DatasourceProxyPOSTByUIDcallsInternalServerError{}
}

/*
DatasourceProxyPOSTByUIDcallsInternalServerError describes a response with status code 500, with default header values.

InternalServerError is a general error indicating something went wrong internally.
*/
type DatasourceProxyPOSTByUIDcallsInternalServerError struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this datasource proxy p o s t by Ui dcalls internal server error response has a 2xx status code
func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this datasource proxy p o s t by Ui dcalls internal server error response has a 3xx status code
func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this datasource proxy p o s t by Ui dcalls internal server error response has a 4xx status code
func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this datasource proxy p o s t by Ui dcalls internal server error response has a 5xx status code
func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this datasource proxy p o s t by Ui dcalls internal server error response a status code equal to that given
func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the datasource proxy p o s t by Ui dcalls internal server error response
func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) Code() int {
	return 500
}

func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsInternalServerError  %+v", 500, o.Payload)
}

func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) String() string {
	return fmt.Sprintf("[POST /datasources/proxy/uid/{uid}/{datasource_proxy_route}][%d] datasourceProxyPOSTByUiDcallsInternalServerError  %+v", 500, o.Payload)
}

func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *DatasourceProxyPOSTByUIDcallsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}