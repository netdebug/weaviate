package rooms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveRoomsRemoveLabelHandlerFunc turns a function with the right signature into a weave rooms remove label handler
type WeaveRoomsRemoveLabelHandlerFunc func(WeaveRoomsRemoveLabelParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveRoomsRemoveLabelHandlerFunc) Handle(params WeaveRoomsRemoveLabelParams) middleware.Responder {
	return fn(params)
}

// WeaveRoomsRemoveLabelHandler interface for that can handle valid weave rooms remove label params
type WeaveRoomsRemoveLabelHandler interface {
	Handle(WeaveRoomsRemoveLabelParams) middleware.Responder
}

// NewWeaveRoomsRemoveLabel creates a new http.Handler for the weave rooms remove label operation
func NewWeaveRoomsRemoveLabel(ctx *middleware.Context, handler WeaveRoomsRemoveLabelHandler) *WeaveRoomsRemoveLabel {
	return &WeaveRoomsRemoveLabel{Context: ctx, Handler: handler}
}

/*WeaveRoomsRemoveLabel swagger:route POST /places/{placeId}/rooms/{roomId}/removeLabel rooms weaveRoomsRemoveLabel

Removes a label from the room.

*/
type WeaveRoomsRemoveLabel struct {
	Context *middleware.Context
	Handler WeaveRoomsRemoveLabelHandler
}

func (o *WeaveRoomsRemoveLabel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveRoomsRemoveLabelParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}