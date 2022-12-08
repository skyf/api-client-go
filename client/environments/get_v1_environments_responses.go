// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/firehydrant/api-client-go/models"
)

// GetV1EnvironmentsReader is a Reader for the GetV1Environments structure.
type GetV1EnvironmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1EnvironmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1EnvironmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1EnvironmentsOK creates a GetV1EnvironmentsOK with default headers values
func NewGetV1EnvironmentsOK() *GetV1EnvironmentsOK {
	return &GetV1EnvironmentsOK{}
}

/*
GetV1EnvironmentsOK describes a response with status code 200, with default header values.

List all of the environments that have been added to the organiation
*/
type GetV1EnvironmentsOK struct {
	Payload *models.EnvironmentEntityPaginated
}

// IsSuccess returns true when this get v1 environments o k response has a 2xx status code
func (o *GetV1EnvironmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v1 environments o k response has a 3xx status code
func (o *GetV1EnvironmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v1 environments o k response has a 4xx status code
func (o *GetV1EnvironmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v1 environments o k response has a 5xx status code
func (o *GetV1EnvironmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v1 environments o k response a status code equal to that given
func (o *GetV1EnvironmentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetV1EnvironmentsOK) Error() string {
	return fmt.Sprintf("[GET /v1/environments][%d] getV1EnvironmentsOK  %+v", 200, o.Payload)
}

func (o *GetV1EnvironmentsOK) String() string {
	return fmt.Sprintf("[GET /v1/environments][%d] getV1EnvironmentsOK  %+v", 200, o.Payload)
}

func (o *GetV1EnvironmentsOK) GetPayload() *models.EnvironmentEntityPaginated {
	return o.Payload
}

func (o *GetV1EnvironmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentEntityPaginated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
