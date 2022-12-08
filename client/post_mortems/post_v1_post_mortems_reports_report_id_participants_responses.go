// Code generated by go-swagger; DO NOT EDIT.

package post_mortems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// PostV1PostMortemsReportsReportIDParticipantsReader is a Reader for the PostV1PostMortemsReportsReportIDParticipants structure.
type PostV1PostMortemsReportsReportIDParticipantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostV1PostMortemsReportsReportIDParticipantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostV1PostMortemsReportsReportIDParticipantsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostV1PostMortemsReportsReportIDParticipantsCreated creates a PostV1PostMortemsReportsReportIDParticipantsCreated with default headers values
func NewPostV1PostMortemsReportsReportIDParticipantsCreated() *PostV1PostMortemsReportsReportIDParticipantsCreated {
	return &PostV1PostMortemsReportsReportIDParticipantsCreated{}
}

/*
PostV1PostMortemsReportsReportIDParticipantsCreated describes a response with status code 201, with default header values.

Add a participant to the retrospective report
*/
type PostV1PostMortemsReportsReportIDParticipantsCreated struct {
	Payload *models.ParticipantEntityPaginated
}

// IsSuccess returns true when this post v1 post mortems reports report Id participants created response has a 2xx status code
func (o *PostV1PostMortemsReportsReportIDParticipantsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post v1 post mortems reports report Id participants created response has a 3xx status code
func (o *PostV1PostMortemsReportsReportIDParticipantsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post v1 post mortems reports report Id participants created response has a 4xx status code
func (o *PostV1PostMortemsReportsReportIDParticipantsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post v1 post mortems reports report Id participants created response has a 5xx status code
func (o *PostV1PostMortemsReportsReportIDParticipantsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post v1 post mortems reports report Id participants created response a status code equal to that given
func (o *PostV1PostMortemsReportsReportIDParticipantsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostV1PostMortemsReportsReportIDParticipantsCreated) Error() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/participants][%d] postV1PostMortemsReportsReportIdParticipantsCreated  %+v", 201, o.Payload)
}

func (o *PostV1PostMortemsReportsReportIDParticipantsCreated) String() string {
	return fmt.Sprintf("[POST /v1/post_mortems/reports/{report_id}/participants][%d] postV1PostMortemsReportsReportIdParticipantsCreated  %+v", 201, o.Payload)
}

func (o *PostV1PostMortemsReportsReportIDParticipantsCreated) GetPayload() *models.ParticipantEntityPaginated {
	return o.Payload
}

func (o *PostV1PostMortemsReportsReportIDParticipantsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ParticipantEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
