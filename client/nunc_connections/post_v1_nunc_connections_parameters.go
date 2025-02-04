// Code generated by go-swagger; DO NOT EDIT.

package nunc_connections

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

// NewPostV1NuncConnectionsParams creates a new PostV1NuncConnectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostV1NuncConnectionsParams() *PostV1NuncConnectionsParams {
	return &PostV1NuncConnectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostV1NuncConnectionsParamsWithTimeout creates a new PostV1NuncConnectionsParams object
// with the ability to set a timeout on a request.
func NewPostV1NuncConnectionsParamsWithTimeout(timeout time.Duration) *PostV1NuncConnectionsParams {
	return &PostV1NuncConnectionsParams{
		timeout: timeout,
	}
}

// NewPostV1NuncConnectionsParamsWithContext creates a new PostV1NuncConnectionsParams object
// with the ability to set a context for a request.
func NewPostV1NuncConnectionsParamsWithContext(ctx context.Context) *PostV1NuncConnectionsParams {
	return &PostV1NuncConnectionsParams{
		Context: ctx,
	}
}

// NewPostV1NuncConnectionsParamsWithHTTPClient creates a new PostV1NuncConnectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostV1NuncConnectionsParamsWithHTTPClient(client *http.Client) *PostV1NuncConnectionsParams {
	return &PostV1NuncConnectionsParams{
		HTTPClient: client,
	}
}

/*
PostV1NuncConnectionsParams contains all the parameters to send to the API endpoint

	for the post v1 nunc connections operation.

	Typically these are written to a http.Request.
*/
type PostV1NuncConnectionsParams struct {

	// ButtonBackgroundColor.
	ButtonBackgroundColor *string

	// ButtonTextColor.
	ButtonTextColor *string

	// CompanyName.
	CompanyName *string

	// CompanyTosURL.
	CompanyTosURL *string

	// CompanyWebsite.
	CompanyWebsite *string

	// CoverImage.
	CoverImage runtime.NamedReadCloser

	// Domain.
	Domain string

	// Favicon.
	Favicon runtime.NamedReadCloser

	// GreetingBody.
	GreetingBody *string

	// GreetingTitle.
	GreetingTitle *string

	// LinkColor.
	LinkColor *string

	// Logo.
	Logo runtime.NamedReadCloser

	// OpenGraphImage.
	OpenGraphImage runtime.NamedReadCloser

	// OperationalMessage.
	OperationalMessage *string

	// PrimaryColor.
	PrimaryColor *string

	// SecondaryColor.
	SecondaryColor *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post v1 nunc connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1NuncConnectionsParams) WithDefaults() *PostV1NuncConnectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post v1 nunc connections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostV1NuncConnectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithTimeout(timeout time.Duration) *PostV1NuncConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithContext(ctx context.Context) *PostV1NuncConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithHTTPClient(client *http.Client) *PostV1NuncConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithButtonBackgroundColor adds the buttonBackgroundColor to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithButtonBackgroundColor(buttonBackgroundColor *string) *PostV1NuncConnectionsParams {
	o.SetButtonBackgroundColor(buttonBackgroundColor)
	return o
}

// SetButtonBackgroundColor adds the buttonBackgroundColor to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetButtonBackgroundColor(buttonBackgroundColor *string) {
	o.ButtonBackgroundColor = buttonBackgroundColor
}

// WithButtonTextColor adds the buttonTextColor to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithButtonTextColor(buttonTextColor *string) *PostV1NuncConnectionsParams {
	o.SetButtonTextColor(buttonTextColor)
	return o
}

// SetButtonTextColor adds the buttonTextColor to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetButtonTextColor(buttonTextColor *string) {
	o.ButtonTextColor = buttonTextColor
}

// WithCompanyName adds the companyName to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithCompanyName(companyName *string) *PostV1NuncConnectionsParams {
	o.SetCompanyName(companyName)
	return o
}

// SetCompanyName adds the companyName to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetCompanyName(companyName *string) {
	o.CompanyName = companyName
}

// WithCompanyTosURL adds the companyTosURL to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithCompanyTosURL(companyTosURL *string) *PostV1NuncConnectionsParams {
	o.SetCompanyTosURL(companyTosURL)
	return o
}

// SetCompanyTosURL adds the companyTosUrl to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetCompanyTosURL(companyTosURL *string) {
	o.CompanyTosURL = companyTosURL
}

// WithCompanyWebsite adds the companyWebsite to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithCompanyWebsite(companyWebsite *string) *PostV1NuncConnectionsParams {
	o.SetCompanyWebsite(companyWebsite)
	return o
}

// SetCompanyWebsite adds the companyWebsite to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetCompanyWebsite(companyWebsite *string) {
	o.CompanyWebsite = companyWebsite
}

// WithCoverImage adds the coverImage to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithCoverImage(coverImage runtime.NamedReadCloser) *PostV1NuncConnectionsParams {
	o.SetCoverImage(coverImage)
	return o
}

// SetCoverImage adds the coverImage to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetCoverImage(coverImage runtime.NamedReadCloser) {
	o.CoverImage = coverImage
}

// WithDomain adds the domain to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithDomain(domain string) *PostV1NuncConnectionsParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetDomain(domain string) {
	o.Domain = domain
}

// WithFavicon adds the favicon to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithFavicon(favicon runtime.NamedReadCloser) *PostV1NuncConnectionsParams {
	o.SetFavicon(favicon)
	return o
}

// SetFavicon adds the favicon to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetFavicon(favicon runtime.NamedReadCloser) {
	o.Favicon = favicon
}

// WithGreetingBody adds the greetingBody to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithGreetingBody(greetingBody *string) *PostV1NuncConnectionsParams {
	o.SetGreetingBody(greetingBody)
	return o
}

// SetGreetingBody adds the greetingBody to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetGreetingBody(greetingBody *string) {
	o.GreetingBody = greetingBody
}

// WithGreetingTitle adds the greetingTitle to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithGreetingTitle(greetingTitle *string) *PostV1NuncConnectionsParams {
	o.SetGreetingTitle(greetingTitle)
	return o
}

// SetGreetingTitle adds the greetingTitle to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetGreetingTitle(greetingTitle *string) {
	o.GreetingTitle = greetingTitle
}

// WithLinkColor adds the linkColor to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithLinkColor(linkColor *string) *PostV1NuncConnectionsParams {
	o.SetLinkColor(linkColor)
	return o
}

// SetLinkColor adds the linkColor to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetLinkColor(linkColor *string) {
	o.LinkColor = linkColor
}

// WithLogo adds the logo to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithLogo(logo runtime.NamedReadCloser) *PostV1NuncConnectionsParams {
	o.SetLogo(logo)
	return o
}

// SetLogo adds the logo to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetLogo(logo runtime.NamedReadCloser) {
	o.Logo = logo
}

// WithOpenGraphImage adds the openGraphImage to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithOpenGraphImage(openGraphImage runtime.NamedReadCloser) *PostV1NuncConnectionsParams {
	o.SetOpenGraphImage(openGraphImage)
	return o
}

// SetOpenGraphImage adds the openGraphImage to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetOpenGraphImage(openGraphImage runtime.NamedReadCloser) {
	o.OpenGraphImage = openGraphImage
}

// WithOperationalMessage adds the operationalMessage to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithOperationalMessage(operationalMessage *string) *PostV1NuncConnectionsParams {
	o.SetOperationalMessage(operationalMessage)
	return o
}

// SetOperationalMessage adds the operationalMessage to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetOperationalMessage(operationalMessage *string) {
	o.OperationalMessage = operationalMessage
}

// WithPrimaryColor adds the primaryColor to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithPrimaryColor(primaryColor *string) *PostV1NuncConnectionsParams {
	o.SetPrimaryColor(primaryColor)
	return o
}

// SetPrimaryColor adds the primaryColor to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetPrimaryColor(primaryColor *string) {
	o.PrimaryColor = primaryColor
}

// WithSecondaryColor adds the secondaryColor to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) WithSecondaryColor(secondaryColor *string) *PostV1NuncConnectionsParams {
	o.SetSecondaryColor(secondaryColor)
	return o
}

// SetSecondaryColor adds the secondaryColor to the post v1 nunc connections params
func (o *PostV1NuncConnectionsParams) SetSecondaryColor(secondaryColor *string) {
	o.SecondaryColor = secondaryColor
}

// WriteToRequest writes these params to a swagger request
func (o *PostV1NuncConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ButtonBackgroundColor != nil {

		// form param button_background_color
		var frButtonBackgroundColor string
		if o.ButtonBackgroundColor != nil {
			frButtonBackgroundColor = *o.ButtonBackgroundColor
		}
		fButtonBackgroundColor := frButtonBackgroundColor
		if fButtonBackgroundColor != "" {
			if err := r.SetFormParam("button_background_color", fButtonBackgroundColor); err != nil {
				return err
			}
		}
	}

	if o.ButtonTextColor != nil {

		// form param button_text_color
		var frButtonTextColor string
		if o.ButtonTextColor != nil {
			frButtonTextColor = *o.ButtonTextColor
		}
		fButtonTextColor := frButtonTextColor
		if fButtonTextColor != "" {
			if err := r.SetFormParam("button_text_color", fButtonTextColor); err != nil {
				return err
			}
		}
	}

	if o.CompanyName != nil {

		// form param company_name
		var frCompanyName string
		if o.CompanyName != nil {
			frCompanyName = *o.CompanyName
		}
		fCompanyName := frCompanyName
		if fCompanyName != "" {
			if err := r.SetFormParam("company_name", fCompanyName); err != nil {
				return err
			}
		}
	}

	if o.CompanyTosURL != nil {

		// form param company_tos_url
		var frCompanyTosURL string
		if o.CompanyTosURL != nil {
			frCompanyTosURL = *o.CompanyTosURL
		}
		fCompanyTosURL := frCompanyTosURL
		if fCompanyTosURL != "" {
			if err := r.SetFormParam("company_tos_url", fCompanyTosURL); err != nil {
				return err
			}
		}
	}

	if o.CompanyWebsite != nil {

		// form param company_website
		var frCompanyWebsite string
		if o.CompanyWebsite != nil {
			frCompanyWebsite = *o.CompanyWebsite
		}
		fCompanyWebsite := frCompanyWebsite
		if fCompanyWebsite != "" {
			if err := r.SetFormParam("company_website", fCompanyWebsite); err != nil {
				return err
			}
		}
	}

	if o.CoverImage != nil {

		if o.CoverImage != nil {
			// form file param cover_image
			if err := r.SetFileParam("cover_image", o.CoverImage); err != nil {
				return err
			}
		}
	}

	// form param domain
	frDomain := o.Domain
	fDomain := frDomain
	if fDomain != "" {
		if err := r.SetFormParam("domain", fDomain); err != nil {
			return err
		}
	}

	if o.Favicon != nil {

		if o.Favicon != nil {
			// form file param favicon
			if err := r.SetFileParam("favicon", o.Favicon); err != nil {
				return err
			}
		}
	}

	if o.GreetingBody != nil {

		// form param greeting_body
		var frGreetingBody string
		if o.GreetingBody != nil {
			frGreetingBody = *o.GreetingBody
		}
		fGreetingBody := frGreetingBody
		if fGreetingBody != "" {
			if err := r.SetFormParam("greeting_body", fGreetingBody); err != nil {
				return err
			}
		}
	}

	if o.GreetingTitle != nil {

		// form param greeting_title
		var frGreetingTitle string
		if o.GreetingTitle != nil {
			frGreetingTitle = *o.GreetingTitle
		}
		fGreetingTitle := frGreetingTitle
		if fGreetingTitle != "" {
			if err := r.SetFormParam("greeting_title", fGreetingTitle); err != nil {
				return err
			}
		}
	}

	if o.LinkColor != nil {

		// form param link_color
		var frLinkColor string
		if o.LinkColor != nil {
			frLinkColor = *o.LinkColor
		}
		fLinkColor := frLinkColor
		if fLinkColor != "" {
			if err := r.SetFormParam("link_color", fLinkColor); err != nil {
				return err
			}
		}
	}

	if o.Logo != nil {

		if o.Logo != nil {
			// form file param logo
			if err := r.SetFileParam("logo", o.Logo); err != nil {
				return err
			}
		}
	}

	if o.OpenGraphImage != nil {

		if o.OpenGraphImage != nil {
			// form file param open_graph_image
			if err := r.SetFileParam("open_graph_image", o.OpenGraphImage); err != nil {
				return err
			}
		}
	}

	if o.OperationalMessage != nil {

		// form param operational_message
		var frOperationalMessage string
		if o.OperationalMessage != nil {
			frOperationalMessage = *o.OperationalMessage
		}
		fOperationalMessage := frOperationalMessage
		if fOperationalMessage != "" {
			if err := r.SetFormParam("operational_message", fOperationalMessage); err != nil {
				return err
			}
		}
	}

	if o.PrimaryColor != nil {

		// form param primary_color
		var frPrimaryColor string
		if o.PrimaryColor != nil {
			frPrimaryColor = *o.PrimaryColor
		}
		fPrimaryColor := frPrimaryColor
		if fPrimaryColor != "" {
			if err := r.SetFormParam("primary_color", fPrimaryColor); err != nil {
				return err
			}
		}
	}

	if o.SecondaryColor != nil {

		// form param secondary_color
		var frSecondaryColor string
		if o.SecondaryColor != nil {
			frSecondaryColor = *o.SecondaryColor
		}
		fSecondaryColor := frSecondaryColor
		if fSecondaryColor != "" {
			if err := r.SetFormParam("secondary_color", fSecondaryColor); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
