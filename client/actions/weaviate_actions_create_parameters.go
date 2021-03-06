/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package actions

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

	models "github.com/creativesoftwarefdn/weaviate/entities/models"
)

// NewWeaviateActionsCreateParams creates a new WeaviateActionsCreateParams object
// with the default values initialized.
func NewWeaviateActionsCreateParams() *WeaviateActionsCreateParams {
	var ()
	return &WeaviateActionsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateActionsCreateParamsWithTimeout creates a new WeaviateActionsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateActionsCreateParamsWithTimeout(timeout time.Duration) *WeaviateActionsCreateParams {
	var ()
	return &WeaviateActionsCreateParams{

		timeout: timeout,
	}
}

// NewWeaviateActionsCreateParamsWithContext creates a new WeaviateActionsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateActionsCreateParamsWithContext(ctx context.Context) *WeaviateActionsCreateParams {
	var ()
	return &WeaviateActionsCreateParams{

		Context: ctx,
	}
}

// NewWeaviateActionsCreateParamsWithHTTPClient creates a new WeaviateActionsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateActionsCreateParamsWithHTTPClient(client *http.Client) *WeaviateActionsCreateParams {
	var ()
	return &WeaviateActionsCreateParams{
		HTTPClient: client,
	}
}

/*WeaviateActionsCreateParams contains all the parameters to send to the API endpoint
for the weaviate actions create operation typically these are written to a http.Request
*/
type WeaviateActionsCreateParams struct {

	/*Body*/
	Body *models.Action

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate actions create params
func (o *WeaviateActionsCreateParams) WithTimeout(timeout time.Duration) *WeaviateActionsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate actions create params
func (o *WeaviateActionsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate actions create params
func (o *WeaviateActionsCreateParams) WithContext(ctx context.Context) *WeaviateActionsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate actions create params
func (o *WeaviateActionsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate actions create params
func (o *WeaviateActionsCreateParams) WithHTTPClient(client *http.Client) *WeaviateActionsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate actions create params
func (o *WeaviateActionsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate actions create params
func (o *WeaviateActionsCreateParams) WithBody(body *models.Action) *WeaviateActionsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate actions create params
func (o *WeaviateActionsCreateParams) SetBody(body *models.Action) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateActionsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
