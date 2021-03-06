// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindOneWithOptionsParams creates a new FindOneWithOptionsParams object
// no default values defined in spec.
func NewFindOneWithOptionsParams() FindOneWithOptionsParams {

	return FindOneWithOptionsParams{}
}

// FindOneWithOptionsParams contains all the bound params for the find one with options operation
// typically these are obtained from a http.Request
//
// swagger:parameters findOneWithOptions
type FindOneWithOptionsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: path
	*/
	ProductID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewFindOneWithOptionsParams() beforehand.
func (o *FindOneWithOptionsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rProductID, rhkProductID, _ := route.Params.GetOK("productID")
	if err := o.bindProductID(rProductID, rhkProductID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindOneWithOptionsParams) bindProductID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// Parameter is provided by construction from the route

	o.ProductID = &raw

	return nil
}
