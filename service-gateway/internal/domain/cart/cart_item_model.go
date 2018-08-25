package cart

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/persistent"
	dto "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagmodel"
	"time"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/product"
)

type CartItem struct {
	persistent.BaseModel

	Index        uint `gorm:"column:index; type:UNSIGNED BIG INT; NOT NULL;"`
	Quantity     uint `gorm:"column:quantity; type:UNSIGNED BIG INT; NOT NULL;"`
	TotalPrice   uint `gorm:"column:total_price; type:UNSIGNED BIG INT; NOT NULL;"`

	// foreign keys
	CartID    uint `gorm:"column:cart_id" sql:"type:UNSIGNED BIG INT REFERENCES Cart(id) ON DELETE RESTRICT ON UPDATE CASCADE"`
	ProductID uint `gorm:"column:product_id" sql:"type:UNSIGNED BIG INT REFERENCES Product(id) ON DELETE RESTRICT ON UPDATE CASCADE"`

	// association
	CartItemOptionList []*CartItemOption `gorm:"foreignkey:CartItemID"`
}

func (CartItem) TableName() string {
	return "CartItem"
}

func (c* CartItem) convertToDTO() *dto.CartItem {
	cartItemID := int64(c.ID)
	index := int64(c.Index)
	updatedAt := c.UpdatedAt.Format(time.RFC3339)

	productID := int64(c.ProductID)
	quantity := int64(c.Quantity)

	// get price for options
	var optionPrice uint = 0
	//for i := range c.CartItemOptionList {
	//	cartItemOption := c.CartItemOptionList[i]
	//	o := cartItemOption.ProductOption
	//	if o != nil && o.Price > 0 {
	//		optionPrice += o.Price
	//	}
	//}

	// get price for product
	var productPrice uint = 0
	//p := c.Product
	//if p != nil && p.Price > 0 {
	//	productPrice += p.Price
	//}

	totalPrice := string(productPrice + optionPrice)

	return &dto.CartItem{
		CartItemID: &cartItemID,

		Index: &index,
		ProductID: &productID,
		Quantity: &quantity,
		TotalPrice: &totalPrice,

		UpdatedAt: &updatedAt,
	}
}

type CartItemOption struct {
	persistent.BaseModel

	Quantity           uint `gorm:"column:quantity; type:UNSIGNED BIG INT; NOT NULL;"`

	// foreign keys
	CartItemID      uint `gorm:"column:cart_item_id" sql:"type:UNSIGNED BIG INT REFERENCES CartItem(id) ON DELETE RESTRICT ON UPDATE CASCADE"`
	ProductOptionID uint `gorm:"column:product_option_id" sql:"type:UNSIGNED BIG INT REFERENCES ProductOption(id) ON DELETE RESTRICT ON UPDATE CASCADE"`

	// association: external
	ProductOption *product.ProductOption `gorm:"save_associations:false"`
}

func (CartItemOption) TableName() string {
	return "CartItemOption"
}

func (c *CartItemOption) convertToDTO(cartItemID uint) *dto.CartItemOption {
	signedCartItemID := int64(cartItemID)
	cartItemOptionID := int64(c.ID)
	productOptionID := int64(c.ProductOptionID)
	quantity := int64(c.Quantity)
	updatedAt := c.UpdatedAt.Format(time.RFC3339)

	return &dto.CartItemOption{
		CartItemID: &signedCartItemID,

		CartItemOptionID: &cartItemOptionID,
		ProductOptionID: &productOptionID,
		Quantity: &quantity,

		UpdatedAt: &updatedAt,
	}
}


