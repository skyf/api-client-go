// Code generated by go-swagger; DO NOT EDIT.

package changes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1ChangesChangeIDReader is a Reader for the DeleteV1ChangesChangeID structure.
type DeleteV1ChangesChangeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1ChangesChangeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1ChangesChangeIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1ChangesChangeIDNoContent creates a DeleteV1ChangesChangeIDNoContent with default headers values
func NewDeleteV1ChangesChangeIDNoContent() *DeleteV1ChangesChangeIDNoContent {
	return &DeleteV1ChangesChangeIDNoContent{}
}

/*
DeleteV1ChangesChangeIDNoContent describes a response with status code 204, with default header values.

Archive a change entry
*/
type DeleteV1ChangesChangeIDNoContent struct {
}

// IsSuccess returns true when this delete v1 changes change Id no content response has a 2xx status code
func (o *DeleteV1ChangesChangeIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 changes change Id no content response has a 3xx status code
func (o *DeleteV1ChangesChangeIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 changes change Id no content response has a 4xx status code
func (o *DeleteV1ChangesChangeIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 changes change Id no content response has a 5xx status code
func (o *DeleteV1ChangesChangeIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 changes change Id no content response a status code equal to that given
func (o *DeleteV1ChangesChangeIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteV1ChangesChangeIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/changes/{change_id}][%d] deleteV1ChangesChangeIdNoContent ", 204)
}

func (o *DeleteV1ChangesChangeIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/changes/{change_id}][%d] deleteV1ChangesChangeIdNoContent ", 204)
}

func (o *DeleteV1ChangesChangeIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
