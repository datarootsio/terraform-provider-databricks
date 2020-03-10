// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewResizeParams creates a new ResizeParams object
// with the default values initialized.
func NewResizeParams() *ResizeParams {
	var ()
	return &ResizeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewResizeParamsWithTimeout creates a new ResizeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResizeParamsWithTimeout(timeout time.Duration) *ResizeParams {
	var ()
	return &ResizeParams{

		timeout: timeout,
	}
}

// NewResizeParamsWithContext creates a new ResizeParams object
// with the default values initialized, and the ability to set a context for a request
func NewResizeParamsWithContext(ctx context.Context) *ResizeParams {
	var ()
	return &ResizeParams{

		Context: ctx,
	}
}

// NewResizeParamsWithHTTPClient creates a new ResizeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResizeParamsWithHTTPClient(client *http.Client) *ResizeParams {
	var ()
	return &ResizeParams{
		HTTPClient: client,
	}
}

/*ResizeParams contains all the parameters to send to the API endpoint
for the resize operation typically these are written to a http.Request
*/
type ResizeParams struct {

	/*Body*/
	Body ResizeBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resize params
func (o *ResizeParams) WithTimeout(timeout time.Duration) *ResizeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resize params
func (o *ResizeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resize params
func (o *ResizeParams) WithContext(ctx context.Context) *ResizeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resize params
func (o *ResizeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resize params
func (o *ResizeParams) WithHTTPClient(client *http.Client) *ResizeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resize params
func (o *ResizeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the resize params
func (o *ResizeParams) WithBody(body ResizeBody) *ResizeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the resize params
func (o *ResizeParams) SetBody(body ResizeBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ResizeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
