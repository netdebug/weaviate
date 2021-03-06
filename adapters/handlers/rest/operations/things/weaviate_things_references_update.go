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

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "github.com/creativesoftwarefdn/weaviate/entities/models"
)

// WeaviateThingsReferencesUpdateHandlerFunc turns a function with the right signature into a weaviate things references update handler
type WeaviateThingsReferencesUpdateHandlerFunc func(WeaviateThingsReferencesUpdateParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaviateThingsReferencesUpdateHandlerFunc) Handle(params WeaviateThingsReferencesUpdateParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// WeaviateThingsReferencesUpdateHandler interface for that can handle valid weaviate things references update params
type WeaviateThingsReferencesUpdateHandler interface {
	Handle(WeaviateThingsReferencesUpdateParams, *models.Principal) middleware.Responder
}

// NewWeaviateThingsReferencesUpdate creates a new http.Handler for the weaviate things references update operation
func NewWeaviateThingsReferencesUpdate(ctx *middleware.Context, handler WeaviateThingsReferencesUpdateHandler) *WeaviateThingsReferencesUpdate {
	return &WeaviateThingsReferencesUpdate{Context: ctx, Handler: handler}
}

/*WeaviateThingsReferencesUpdate swagger:route PUT /things/{id}/references/{propertyName} things weaviateThingsReferencesUpdate

Replace all references to a class-property.

Replace all references to a class-property.

*/
type WeaviateThingsReferencesUpdate struct {
	Context *middleware.Context
	Handler WeaviateThingsReferencesUpdateHandler
}

func (o *WeaviateThingsReferencesUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewWeaviateThingsReferencesUpdateParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
