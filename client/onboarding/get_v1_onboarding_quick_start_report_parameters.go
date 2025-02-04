// Code generated by go-swagger; DO NOT EDIT.

package onboarding

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
)

// NewGetV1OnboardingQuickStartReportParams creates a new GetV1OnboardingQuickStartReportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1OnboardingQuickStartReportParams() *GetV1OnboardingQuickStartReportParams {
	return &GetV1OnboardingQuickStartReportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1OnboardingQuickStartReportParamsWithTimeout creates a new GetV1OnboardingQuickStartReportParams object
// with the ability to set a timeout on a request.
func NewGetV1OnboardingQuickStartReportParamsWithTimeout(timeout time.Duration) *GetV1OnboardingQuickStartReportParams {
	return &GetV1OnboardingQuickStartReportParams{
		timeout: timeout,
	}
}

// NewGetV1OnboardingQuickStartReportParamsWithContext creates a new GetV1OnboardingQuickStartReportParams object
// with the ability to set a context for a request.
func NewGetV1OnboardingQuickStartReportParamsWithContext(ctx context.Context) *GetV1OnboardingQuickStartReportParams {
	return &GetV1OnboardingQuickStartReportParams{
		Context: ctx,
	}
}

// NewGetV1OnboardingQuickStartReportParamsWithHTTPClient creates a new GetV1OnboardingQuickStartReportParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1OnboardingQuickStartReportParamsWithHTTPClient(client *http.Client) *GetV1OnboardingQuickStartReportParams {
	return &GetV1OnboardingQuickStartReportParams{
		HTTPClient: client,
	}
}

/*
GetV1OnboardingQuickStartReportParams contains all the parameters to send to the API endpoint

	for the get v1 onboarding quick start report operation.

	Typically these are written to a http.Request.
*/
type GetV1OnboardingQuickStartReportParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 onboarding quick start report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1OnboardingQuickStartReportParams) WithDefaults() *GetV1OnboardingQuickStartReportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 onboarding quick start report params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1OnboardingQuickStartReportParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 onboarding quick start report params
func (o *GetV1OnboardingQuickStartReportParams) WithTimeout(timeout time.Duration) *GetV1OnboardingQuickStartReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 onboarding quick start report params
func (o *GetV1OnboardingQuickStartReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 onboarding quick start report params
func (o *GetV1OnboardingQuickStartReportParams) WithContext(ctx context.Context) *GetV1OnboardingQuickStartReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 onboarding quick start report params
func (o *GetV1OnboardingQuickStartReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 onboarding quick start report params
func (o *GetV1OnboardingQuickStartReportParams) WithHTTPClient(client *http.Client) *GetV1OnboardingQuickStartReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 onboarding quick start report params
func (o *GetV1OnboardingQuickStartReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1OnboardingQuickStartReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
