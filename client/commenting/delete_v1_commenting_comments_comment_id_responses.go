// Code generated by go-swagger; DO NOT EDIT.

package commenting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteV1CommentingCommentsCommentIDReader is a Reader for the DeleteV1CommentingCommentsCommentID structure.
type DeleteV1CommentingCommentsCommentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteV1CommentingCommentsCommentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteV1CommentingCommentsCommentIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteV1CommentingCommentsCommentIDNoContent creates a DeleteV1CommentingCommentsCommentIDNoContent with default headers values
func NewDeleteV1CommentingCommentsCommentIDNoContent() *DeleteV1CommentingCommentsCommentIDNoContent {
	return &DeleteV1CommentingCommentsCommentIDNoContent{}
}

/*
DeleteV1CommentingCommentsCommentIDNoContent describes a response with status code 204, with default header values.

Archive a comment
*/
type DeleteV1CommentingCommentsCommentIDNoContent struct {
}

// IsSuccess returns true when this delete v1 commenting comments comment Id no content response has a 2xx status code
func (o *DeleteV1CommentingCommentsCommentIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete v1 commenting comments comment Id no content response has a 3xx status code
func (o *DeleteV1CommentingCommentsCommentIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete v1 commenting comments comment Id no content response has a 4xx status code
func (o *DeleteV1CommentingCommentsCommentIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete v1 commenting comments comment Id no content response has a 5xx status code
func (o *DeleteV1CommentingCommentsCommentIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete v1 commenting comments comment Id no content response a status code equal to that given
func (o *DeleteV1CommentingCommentsCommentIDNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteV1CommentingCommentsCommentIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/commenting/comments/{comment_id}][%d] deleteV1CommentingCommentsCommentIdNoContent ", 204)
}

func (o *DeleteV1CommentingCommentsCommentIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /v1/commenting/comments/{comment_id}][%d] deleteV1CommentingCommentsCommentIdNoContent ", 204)
}

func (o *DeleteV1CommentingCommentsCommentIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
