// Code generated by go-swagger; DO NOT EDIT.

package cart

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// AddCartItemURL generates an URL for the add cart item operation
type AddCartItemURL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AddCartItemURL) WithBasePath(bp string) *AddCartItemURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *AddCartItemURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *AddCartItemURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/cart"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *AddCartItemURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *AddCartItemURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *AddCartItemURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on AddCartItemURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on AddCartItemURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *AddCartItemURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
