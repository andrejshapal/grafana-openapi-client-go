// Code generated by go-swagger; DO NOT EDIT.

package saml

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetMetadataReader is a Reader for the GetMetadata structure.
type GetMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /saml/metadata] getMetadata", response, response.Code())
	}
}

// NewGetMetadataOK creates a GetMetadataOK with default headers values
func NewGetMetadataOK() *GetMetadataOK {
	return &GetMetadataOK{}
}

/*
GetMetadataOK describes a response with status code 200, with default header values.

(empty)
*/
type GetMetadataOK struct {
	Payload []uint8
}

// IsSuccess returns true when this get metadata o k response has a 2xx status code
func (o *GetMetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get metadata o k response has a 3xx status code
func (o *GetMetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metadata o k response has a 4xx status code
func (o *GetMetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get metadata o k response has a 5xx status code
func (o *GetMetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get metadata o k response a status code equal to that given
func (o *GetMetadataOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get metadata o k response
func (o *GetMetadataOK) Code() int {
	return 200
}

func (o *GetMetadataOK) Error() string {
	return fmt.Sprintf("[GET /saml/metadata][%d] getMetadataOK  %+v", 200, o.Payload)
}

func (o *GetMetadataOK) String() string {
	return fmt.Sprintf("[GET /saml/metadata][%d] getMetadataOK  %+v", 200, o.Payload)
}

func (o *GetMetadataOK) GetPayload() []uint8 {
	return o.Payload
}

func (o *GetMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}