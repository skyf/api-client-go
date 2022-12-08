// Code generated by go-swagger; DO NOT EDIT.

package scheduled_maintenances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1ScheduledMaintenancesReader is a Reader for the PostV1ScheduledMaintenances structure.
type PostV1ScheduledMaintenancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1ScheduledMaintenancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1ScheduledMaintenancesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostV1ScheduledMaintenancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1ScheduledMaintenancesCreated creates a PostV1ScheduledMaintenancesCreated with default headers values
func NewPostV1ScheduledMaintenancesCreated() *PostV1ScheduledMaintenancesCreated {
	return &PostV1ScheduledMaintenancesCreated{}
}

/*
PostV1ScheduledMaintenancesCreated describes a response with status code 201, with default header values.

Create a new scheduled maintenance event
*/
type PostV1ScheduledMaintenancesCreated struct {
	Payload *models.ScheduledMaintenanceEntity
}

// IsSuccess returns true when this post v1 scheduled maintenances created response has a 2xx status code
func (o *PostV1ScheduledMaintenancesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 scheduled maintenances created response has a 3xx status code
func (o *PostV1ScheduledMaintenancesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 scheduled maintenances created response has a 4xx status code
func (o *PostV1ScheduledMaintenancesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 scheduled maintenances created response has a 5xx status code
func (o *PostV1ScheduledMaintenancesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 scheduled maintenances created response a status code equal to that given
func (o *PostV1ScheduledMaintenancesCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1ScheduledMaintenancesCreated) Error() string {
	return fmt.Sprintf("[POST /v1/scheduled_maintenances][%d] postV1ScheduledMaintenancesCreated  %+v", 201, o.Payload)
}

func (o *PostV1ScheduledMaintenancesCreated) String() string {
	return fmt.Sprintf("[POST /v1/scheduled_maintenances][%d] postV1ScheduledMaintenancesCreated  %+v", 201, o.Payload)
}

func (o *PostV1ScheduledMaintenancesCreated) GetPayload() *models.ScheduledMaintenanceEntity {
	return o.Payload
}

func (o *PostV1ScheduledMaintenancesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduledMaintenanceEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostV1ScheduledMaintenancesBadRequest creates a PostV1ScheduledMaintenancesBadRequest with default headers values
func NewPostV1ScheduledMaintenancesBadRequest() *PostV1ScheduledMaintenancesBadRequest {
	return &PostV1ScheduledMaintenancesBadRequest{}
}

/*
PostV1ScheduledMaintenancesBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostV1ScheduledMaintenancesBadRequest struct {
	Payload *models.ErrorEntity
}

// IsSuccess returns true when this post v1 scheduled maintenances bad request response has a 2xx status code
func (o *PostV1ScheduledMaintenancesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post v1 scheduled maintenances bad request response has a 3xx status code
func (o *PostV1ScheduledMaintenancesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 scheduled maintenances bad request response has a 4xx status code
func (o *PostV1ScheduledMaintenancesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post v1 scheduled maintenances bad request response has a 5xx status code
func (o *PostV1ScheduledMaintenancesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 scheduled maintenances bad request response a status code equal to that given
func (o *PostV1ScheduledMaintenancesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostV1ScheduledMaintenancesBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/scheduled_maintenances][%d] postV1ScheduledMaintenancesBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1ScheduledMaintenancesBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/scheduled_maintenances][%d] postV1ScheduledMaintenancesBadRequest  %+v", 400, o.Payload)
}

func (o *PostV1ScheduledMaintenancesBadRequest) GetPayload() *models.ErrorEntity {
	return o.Payload
}

func (o *PostV1ScheduledMaintenancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
