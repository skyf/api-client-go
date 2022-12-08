// Code generated by go-swagger; DO NOT EDIT.

package ticketing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PatchV1TicketingTicketsTicketIDReader is a Reader for the PatchV1TicketingTicketsTicketID structure.
type PatchV1TicketingTicketsTicketIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchV1TicketingTicketsTicketIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchV1TicketingTicketsTicketIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchV1TicketingTicketsTicketIDOK creates a PatchV1TicketingTicketsTicketIDOK with default headers values
func NewPatchV1TicketingTicketsTicketIDOK() *PatchV1TicketingTicketsTicketIDOK {
	return &PatchV1TicketingTicketsTicketIDOK{}
}

/*
PatchV1TicketingTicketsTicketIDOK describes a response with status code 200, with default header values.

Update a ticket's attributes
*/
type PatchV1TicketingTicketsTicketIDOK struct {
	Payload *models.TicketEntity
}

// IsSuccess returns true when this patch v1 ticketing tickets ticket Id o k response has a 2xx status code
func (o *PatchV1TicketingTicketsTicketIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch v1 ticketing tickets ticket Id o k response has a 3xx status code
func (o *PatchV1TicketingTicketsTicketIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch v1 ticketing tickets ticket Id o k response has a 4xx status code
func (o *PatchV1TicketingTicketsTicketIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch v1 ticketing tickets ticket Id o k response has a 5xx status code
func (o *PatchV1TicketingTicketsTicketIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch v1 ticketing tickets ticket Id o k response a status code equal to that given
func (o *PatchV1TicketingTicketsTicketIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchV1TicketingTicketsTicketIDOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/ticketing/tickets/{ticket_id}][%d] patchV1TicketingTicketsTicketIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1TicketingTicketsTicketIDOK) String() string {
	return fmt.Sprintf("[PATCH /v1/ticketing/tickets/{ticket_id}][%d] patchV1TicketingTicketsTicketIdOK  %+v", 200, o.Payload)
}

func (o *PatchV1TicketingTicketsTicketIDOK) GetPayload() *models.TicketEntity {
	return o.Payload
}

func (o *PatchV1TicketingTicketsTicketIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TicketEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
