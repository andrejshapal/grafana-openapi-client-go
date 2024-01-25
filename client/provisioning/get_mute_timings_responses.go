// Code generated by go-swagger; DO NOT EDIT.

package provisioning

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

// GetMuteTimingsReader is a Reader for the GetMuteTimings structure.
type GetMuteTimingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMuteTimingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMuteTimingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /v1/provisioning/mute-timings] GetMuteTimings", response, response.Code())
	}
}

// NewGetMuteTimingsOK creates a GetMuteTimingsOK with default headers values
func NewGetMuteTimingsOK() *GetMuteTimingsOK {
	return &GetMuteTimingsOK{}
}

/*
GetMuteTimingsOK describes a response with status code 200, with default header values.

MuteTimings
*/
type GetMuteTimingsOK struct {
	Payload models.MuteTimings
}

// IsSuccess returns true when this get mute timings Ok response has a 2xx status code
func (o *GetMuteTimingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get mute timings Ok response has a 3xx status code
func (o *GetMuteTimingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get mute timings Ok response has a 4xx status code
func (o *GetMuteTimingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get mute timings Ok response has a 5xx status code
func (o *GetMuteTimingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get mute timings Ok response a status code equal to that given
func (o *GetMuteTimingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get mute timings Ok response
func (o *GetMuteTimingsOK) Code() int {
	return 200
}

func (o *GetMuteTimingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/mute-timings][%d] getMuteTimingsOk %s", 200, payload)
}

func (o *GetMuteTimingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/provisioning/mute-timings][%d] getMuteTimingsOk %s", 200, payload)
}

func (o *GetMuteTimingsOK) GetPayload() models.MuteTimings {
	return o.Payload
}

func (o *GetMuteTimingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
