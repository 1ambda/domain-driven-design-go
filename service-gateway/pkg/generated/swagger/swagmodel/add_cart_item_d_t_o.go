// Code generated by go-swagger; DO NOT EDIT.

package swagmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddCartItemDTO add cart item d t o
// swagger:model addCartItemDTO
type AddCartItemDTO struct {

	// product ID
	// Required: true
	ProductID *string `json:"productID"`

	// product option ID list
	// Required: true
	ProductOptionIDList []string `json:"productOptionIDList"`

	// quantity
	// Required: true
	Quantity *int64 `json:"quantity"`
}

// Validate validates this add cart item d t o
func (m *AddCartItemDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProductID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductOptionIDList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddCartItemDTO) validateProductID(formats strfmt.Registry) error {

	if err := validate.Required("productID", "body", m.ProductID); err != nil {
		return err
	}

	return nil
}

func (m *AddCartItemDTO) validateProductOptionIDList(formats strfmt.Registry) error {

	if err := validate.Required("productOptionIDList", "body", m.ProductOptionIDList); err != nil {
		return err
	}

	return nil
}

func (m *AddCartItemDTO) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddCartItemDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddCartItemDTO) UnmarshalBinary(b []byte) error {
	var res AddCartItemDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
