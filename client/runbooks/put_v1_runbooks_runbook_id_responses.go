// Code generated by go-swagger; DO NOT EDIT.

package runbooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PutV1RunbooksRunbookIDReader is a Reader for the PutV1RunbooksRunbookID structure.
type PutV1RunbooksRunbookIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1RunbooksRunbookIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV1RunbooksRunbookIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutV1RunbooksRunbookIDOK creates a PutV1RunbooksRunbookIDOK with default headers values
func NewPutV1RunbooksRunbookIDOK() *PutV1RunbooksRunbookIDOK {
	return &PutV1RunbooksRunbookIDOK{}
}

/*
	PutV1RunbooksRunbookIDOK describes a response with status code 200, with default header values.

	Update a runbook and any attachment rules associated with it. This endpoint is used to configure nearly everything

about a runbook, including but not limited to the steps, environments, attachment rules, and severities.
*/
type PutV1RunbooksRunbookIDOK struct {
	Payload *models.RunbookEntity
}

// IsSuccess returns true when this put v1 runbooks runbook Id o k response has a 2xx status code
func (o *PutV1RunbooksRunbookIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put v1 runbooks runbook Id o k response has a 3xx status code
func (o *PutV1RunbooksRunbookIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 runbooks runbook Id o k response has a 4xx status code
func (o *PutV1RunbooksRunbookIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 runbooks runbook Id o k response has a 5xx status code
func (o *PutV1RunbooksRunbookIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 runbooks runbook Id o k response a status code equal to that given
func (o *PutV1RunbooksRunbookIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutV1RunbooksRunbookIDOK) Error() string {
	return fmt.Sprintf("[PUT /v1/runbooks/{runbook_id}][%d] putV1RunbooksRunbookIdOK  %+v", 200, o.Payload)
}

func (o *PutV1RunbooksRunbookIDOK) String() string {
	return fmt.Sprintf("[PUT /v1/runbooks/{runbook_id}][%d] putV1RunbooksRunbookIdOK  %+v", 200, o.Payload)
}

func (o *PutV1RunbooksRunbookIDOK) GetPayload() *models.RunbookEntity {
	return o.Payload
}

func (o *PutV1RunbooksRunbookIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunbookEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
