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

// PutAlertRuleGroupReader is a Reader for the PutAlertRuleGroup structure.
type PutAlertRuleGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAlertRuleGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAlertRuleGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAlertRuleGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /v1/provisioning/folder/{FolderUID}/rule-groups/{Group}] PutAlertRuleGroup", response, response.Code())
	}
}

// NewPutAlertRuleGroupOK creates a PutAlertRuleGroupOK with default headers values
func NewPutAlertRuleGroupOK() *PutAlertRuleGroupOK {
	return &PutAlertRuleGroupOK{}
}

/*
PutAlertRuleGroupOK describes a response with status code 200, with default header values.

AlertRuleGroup
*/
type PutAlertRuleGroupOK struct {
	Payload *models.AlertRuleGroup
}

// IsSuccess returns true when this put alert rule group Ok response has a 2xx status code
func (o *PutAlertRuleGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put alert rule group Ok response has a 3xx status code
func (o *PutAlertRuleGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put alert rule group Ok response has a 4xx status code
func (o *PutAlertRuleGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put alert rule group Ok response has a 5xx status code
func (o *PutAlertRuleGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put alert rule group Ok response a status code equal to that given
func (o *PutAlertRuleGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put alert rule group Ok response
func (o *PutAlertRuleGroupOK) Code() int {
	return 200
}

func (o *PutAlertRuleGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] putAlertRuleGroupOk %s", 200, payload)
}

func (o *PutAlertRuleGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] putAlertRuleGroupOk %s", 200, payload)
}

func (o *PutAlertRuleGroupOK) GetPayload() *models.AlertRuleGroup {
	return o.Payload
}

func (o *PutAlertRuleGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertRuleGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAlertRuleGroupBadRequest creates a PutAlertRuleGroupBadRequest with default headers values
func NewPutAlertRuleGroupBadRequest() *PutAlertRuleGroupBadRequest {
	return &PutAlertRuleGroupBadRequest{}
}

/*
PutAlertRuleGroupBadRequest describes a response with status code 400, with default header values.

ValidationError
*/
type PutAlertRuleGroupBadRequest struct {
	Payload *models.ValidationError
}

// IsSuccess returns true when this put alert rule group bad request response has a 2xx status code
func (o *PutAlertRuleGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put alert rule group bad request response has a 3xx status code
func (o *PutAlertRuleGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put alert rule group bad request response has a 4xx status code
func (o *PutAlertRuleGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put alert rule group bad request response has a 5xx status code
func (o *PutAlertRuleGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put alert rule group bad request response a status code equal to that given
func (o *PutAlertRuleGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put alert rule group bad request response
func (o *PutAlertRuleGroupBadRequest) Code() int {
	return 400
}

func (o *PutAlertRuleGroupBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] putAlertRuleGroupBadRequest %s", 400, payload)
}

func (o *PutAlertRuleGroupBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /v1/provisioning/folder/{FolderUID}/rule-groups/{Group}][%d] putAlertRuleGroupBadRequest %s", 400, payload)
}

func (o *PutAlertRuleGroupBadRequest) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *PutAlertRuleGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
