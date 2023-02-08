// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/firehydrant/api-client-go/models"
)

// NewPostV1ConversationsConversationIDCommentsParams creates a new PostV1ConversationsConversationIDCommentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1ConversationsConversationIDCommentsParams() *PostV1ConversationsConversationIDCommentsParams {
	return &PostV1ConversationsConversationIDCommentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1ConversationsConversationIDCommentsParamsWithTimeout creates a new PostV1ConversationsConversationIDCommentsParams object
// with the ability to set a timeout on a request.
func NewPostV1ConversationsConversationIDCommentsParamsWithTimeout(timeout time.Duration) *PostV1ConversationsConversationIDCommentsParams {
	return &PostV1ConversationsConversationIDCommentsParams{
		timeout: timeout,
	}
}

// NewPostV1ConversationsConversationIDCommentsParamsWithContext creates a new PostV1ConversationsConversationIDCommentsParams object
// with the ability to set a context for a request.
func NewPostV1ConversationsConversationIDCommentsParamsWithContext(ctx context.Context) *PostV1ConversationsConversationIDCommentsParams {
	return &PostV1ConversationsConversationIDCommentsParams{
		Context: ctx,
	}
}

// NewPostV1ConversationsConversationIDCommentsParamsWithHTTPClient creates a new PostV1ConversationsConversationIDCommentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1ConversationsConversationIDCommentsParamsWithHTTPClient(client *http.Client) *PostV1ConversationsConversationIDCommentsParams {
	return &PostV1ConversationsConversationIDCommentsParams{
		HTTPClient: client,
	}
}

/*
PostV1ConversationsConversationIDCommentsParams contains all the parameters to send to the API endpoint

	for the post v1 conversations conversation Id comments operation.

	Typically these are written to a http.Request.
*/
type PostV1ConversationsConversationIDCommentsParams struct {

	// V1ConversationsConversationIDComments.
	V1ConversationsConversationIDComments *models.PostV1ConversationsConversationIDComments

	// ConversationID.
	//
	// Format: int32
	ConversationID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 conversations conversation Id comments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ConversationsConversationIDCommentsParams) WithDefaults() *PostV1ConversationsConversationIDCommentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 conversations conversation Id comments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1ConversationsConversationIDCommentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 conversations conversation Id comments params
func (o *PostV1ConversationsConversationIDCommentsParams) WithTimeout(timeout time.Duration) *PostV1ConversationsConversationIDCommentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 conversations conversation Id comments params
func (o *PostV1ConversationsConversationIDCommentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 conversations conversation Id comments params
func (o *PostV1ConversationsConversationIDCommentsParams) WithContext(ctx context.Context) *PostV1ConversationsConversationIDCommentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 conversations conversation Id comments params
func (o *PostV1ConversationsConversationIDCommentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 conversations conversation Id comments params
func (o *PostV1ConversationsConversationIDCommentsParams) WithHTTPClient(client *http.Client) *PostV1ConversationsConversationIDCommentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 conversations conversation Id comments params
func (o *PostV1ConversationsConversationIDCommentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithV1ConversationsConversationIDComments adds the v1ConversationsConversationIDComments to the post v1 conversations conversation Id comments params
func (o *PostV1ConversationsConversationIDCommentsParams) WithV1ConversationsConversationIDComments(v1ConversationsConversationIDComments *models.PostV1ConversationsConversationIDComments) *PostV1ConversationsConversationIDCommentsParams {
	o.SetV1ConversationsConversationIDComments(v1ConversationsConversationIDComments)
	return o
}

// SetV1ConversationsConversationIDComments adds the v1ConversationsConversationIdComments to the post v1 conversations conversation Id comments params
func (o *PostV1ConversationsConversationIDCommentsParams) SetV1ConversationsConversationIDComments(v1ConversationsConversationIDComments *models.PostV1ConversationsConversationIDComments) {
	o.V1ConversationsConversationIDComments = v1ConversationsConversationIDComments
}

// WithConversationID adds the conversationID to the post v1 conversations conversation Id comments params
func (o *PostV1ConversationsConversationIDCommentsParams) WithConversationID(conversationID int32) *PostV1ConversationsConversationIDCommentsParams {
	o.SetConversationID(conversationID)
	return o
}

// SetConversationID adds the conversationId to the post v1 conversations conversation Id comments params
func (o *PostV1ConversationsConversationIDCommentsParams) SetConversationID(conversationID int32) {
	o.ConversationID = conversationID
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1ConversationsConversationIDCommentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.V1ConversationsConversationIDComments != nil {
		if err := r.SetBodyParam(o.V1ConversationsConversationIDComments); err != nil {
			return err
		}
	}

	// path param conversation_id
	if err := r.SetPathParam("conversation_id", swag.FormatInt32(o.ConversationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
