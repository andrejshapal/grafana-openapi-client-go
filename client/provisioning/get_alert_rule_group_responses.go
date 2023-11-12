// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/grafana/grafana-openapi-client-go/models"
)

// GetAlertRuleGroupReader is a Reader for the GetAlertRuleGroup structure.
type GetAlertRuleGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertRuleGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertRuleGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAlertRuleGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}] GetAlertRuleGroup", response, response.Code())
	}
}

// NewGetAlertRuleGroupOK creates a GetAlertRuleGroupOK with default headers values
func NewGetAlertRuleGroupOK() *GetAlertRuleGroupOK {
	return &GetAlertRuleGroupOK{}
}

/*
GetAlertRuleGroupOK describes a response with status code 200, with default header values.

AlertRuleGroup
*/
type GetAlertRuleGroupOK struct {
	Payload *models.AlertRuleGroup
}

// IsSuccess returns true when this get alert rule group Ok response has a 2xx status code
func (o *GetAlertRuleGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert rule group Ok response has a 3xx status code
func (o *GetAlertRuleGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert rule group Ok response has a 4xx status code
func (o *GetAlertRuleGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert rule group Ok response has a 5xx status code
func (o *GetAlertRuleGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert rule group Ok response a status code equal to that given
func (o *GetAlertRuleGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alert rule group Ok response
func (o *GetAlertRuleGroupOK) Code() int {
	return 200
}

func (o *GetAlertRuleGroupOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] getAlertRuleGroupOk  %+v", 200, o.Payload)
}

func (o *GetAlertRuleGroupOK) String() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] getAlertRuleGroupOk  %+v", 200, o.Payload)
}

func (o *GetAlertRuleGroupOK) GetPayload() *models.AlertRuleGroup {
	return o.Payload
}

func (o *GetAlertRuleGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertRuleGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertRuleGroupNotFound creates a GetAlertRuleGroupNotFound with default headers values
func NewGetAlertRuleGroupNotFound() *GetAlertRuleGroupNotFound {
	return &GetAlertRuleGroupNotFound{}
}

/*
GetAlertRuleGroupNotFound describes a response with status code 404, with default header values.

	Not found.
*/
type GetAlertRuleGroupNotFound struct {
}

// IsSuccess returns true when this get alert rule group not found response has a 2xx status code
func (o *GetAlertRuleGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get alert rule group not found response has a 3xx status code
func (o *GetAlertRuleGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert rule group not found response has a 4xx status code
func (o *GetAlertRuleGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get alert rule group not found response has a 5xx status code
func (o *GetAlertRuleGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert rule group not found response a status code equal to that given
func (o *GetAlertRuleGroupNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get alert rule group not found response
func (o *GetAlertRuleGroupNotFound) Code() int {
	return 404
}

func (o *GetAlertRuleGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] getAlertRuleGroupNotFound ", 404)
}

func (o *GetAlertRuleGroupNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] getAlertRuleGroupNotFound ", 404)
}

func (o *GetAlertRuleGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}