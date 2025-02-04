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

// PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateReader is a Reader for the PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptState structure.
type PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK creates a PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK with default headers values
func NewPutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK() *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK {
	return &PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK{}
}

/*
PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK describes a response with status code 200, with default header values.

Updates the execution's step.
*/
type PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK struct {
	Payload *models.ExecutionEntity
}

// IsSuccess returns true when this put v1 runbooks executions execution Id steps step Id script state o k response has a 2xx status code
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put v1 runbooks executions execution Id steps step Id script state o k response has a 3xx status code
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put v1 runbooks executions execution Id steps step Id script state o k response has a 4xx status code
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put v1 runbooks executions execution Id steps step Id script state o k response has a 5xx status code
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put v1 runbooks executions execution Id steps step Id script state o k response a status code equal to that given
func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK) Error() string {
	return fmt.Sprintf("[PUT /v1/runbooks/executions/{execution_id}/steps/{step_id}/script/{state}][%d] putV1RunbooksExecutionsExecutionIdStepsStepIdScriptStateOK  %+v", 200, o.Payload)
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK) String() string {
	return fmt.Sprintf("[PUT /v1/runbooks/executions/{execution_id}/steps/{step_id}/script/{state}][%d] putV1RunbooksExecutionsExecutionIdStepsStepIdScriptStateOK  %+v", 200, o.Payload)
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK) GetPayload() *models.ExecutionEntity {
	return o.Payload
}

func (o *PutV1RunbooksExecutionsExecutionIDStepsStepIDScriptStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExecutionEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
