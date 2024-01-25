// Code generated by go-swagger; DO NOT EDIT.

package search

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

// ListSortOptionsReader is a Reader for the ListSortOptions structure.
type ListSortOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSortOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSortOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListSortOptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /search/sorting] listSortOptions", response, response.Code())
	}
}

// NewListSortOptionsOK creates a ListSortOptionsOK with default headers values
func NewListSortOptionsOK() *ListSortOptionsOK {
	return &ListSortOptionsOK{}
}

/*
ListSortOptionsOK describes a response with status code 200, with default header values.

(empty)
*/
type ListSortOptionsOK struct {
	Payload *models.ListSortOptionsOKBody
}

// IsSuccess returns true when this list sort options Ok response has a 2xx status code
func (o *ListSortOptionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list sort options Ok response has a 3xx status code
func (o *ListSortOptionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list sort options Ok response has a 4xx status code
func (o *ListSortOptionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list sort options Ok response has a 5xx status code
func (o *ListSortOptionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list sort options Ok response a status code equal to that given
func (o *ListSortOptionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list sort options Ok response
func (o *ListSortOptionsOK) Code() int {
	return 200
}

func (o *ListSortOptionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /search/sorting][%d] listSortOptionsOk %s", 200, payload)
}

func (o *ListSortOptionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /search/sorting][%d] listSortOptionsOk %s", 200, payload)
}

func (o *ListSortOptionsOK) GetPayload() *models.ListSortOptionsOKBody {
	return o.Payload
}

func (o *ListSortOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListSortOptionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSortOptionsUnauthorized creates a ListSortOptionsUnauthorized with default headers values
func NewListSortOptionsUnauthorized() *ListSortOptionsUnauthorized {
	return &ListSortOptionsUnauthorized{}
}

/*
ListSortOptionsUnauthorized describes a response with status code 401, with default header values.

UnauthorizedError is returned when the request is not authenticated.
*/
type ListSortOptionsUnauthorized struct {
	Payload *models.ErrorResponseBody
}

// IsSuccess returns true when this list sort options unauthorized response has a 2xx status code
func (o *ListSortOptionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list sort options unauthorized response has a 3xx status code
func (o *ListSortOptionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list sort options unauthorized response has a 4xx status code
func (o *ListSortOptionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list sort options unauthorized response has a 5xx status code
func (o *ListSortOptionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list sort options unauthorized response a status code equal to that given
func (o *ListSortOptionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list sort options unauthorized response
func (o *ListSortOptionsUnauthorized) Code() int {
	return 401
}

func (o *ListSortOptionsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /search/sorting][%d] listSortOptionsUnauthorized %s", 401, payload)
}

func (o *ListSortOptionsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /search/sorting][%d] listSortOptionsUnauthorized %s", 401, payload)
}

func (o *ListSortOptionsUnauthorized) GetPayload() *models.ErrorResponseBody {
	return o.Payload
}

func (o *ListSortOptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponseBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
