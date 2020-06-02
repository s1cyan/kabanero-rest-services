// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DescribeHandlerFunc turns a function with the right signature into a describe handler
type DescribeHandlerFunc func(DescribeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DescribeHandlerFunc) Handle(params DescribeParams) middleware.Responder {
	return fn(params)
}

// DescribeHandler interface for that can handle valid describe params
type DescribeHandler interface {
	Handle(DescribeParams) middleware.Responder
}

// NewDescribe creates a new http.Handler for the describe operation
func NewDescribe(ctx *middleware.Context, handler DescribeHandler) *Describe {
	return &Describe{Context: ctx, Handler: handler}
}

/*Describe swagger:route GET /v1/describe/stacks/{stackName}/versions/{version} describe

Describe describe API

*/
type Describe struct {
	Context *middleware.Context
	Handler DescribeHandler
}

func (o *Describe) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDescribeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
