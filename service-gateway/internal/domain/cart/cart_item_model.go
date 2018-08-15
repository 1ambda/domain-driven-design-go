package cart

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/persistent"
)

type CartItem struct {
	persistent.BaseModel

	Index        uint `gorm:"column:index; type:UNSIGNED BIG INT; NOT NULL;"`
	Quantity     uint `gorm:"column:quantity; type:UNSIGNED BIG INT; NOT NULL;"`
	ProductPrice uint `gorm:"column:product_price; type:UNSIGNED BIG INT; NOT NULL;"`
	TotalPrice   uint `gorm:"column:total_price; type:UNSIGNED BIG INT; NOT NULL;"`

	CartID    uint `gorm:"column:cart_id" sql:"type:UNSIGNED BIG INT REFERENCES Cart(id) ON DELETE RESTRICT ON UPDATE CASCADE"`
	ProductID uint `gorm:"column:product_id" sql:"type:UNSIGNED BIG INT REFERENCES Product(id) ON DELETE RESTRICT ON UPDATE CASCADE"`
}

func (CartItem) TableName() string {
	return "CartItem"
}

type CartItemOption struct {
	persistent.BaseModel

	Quantity           uint `gorm:"column:quantity; type:UNSIGNED BIG INT; NOT NULL;"`
	ProductOptionPrice int  `gorm:"column:product_price; type:UNSIGNED BIG INT; NOT NULL;"`

	CartItemID      uint `gorm:"column:cart_item_id" sql:"type:UNSIGNED BIG INT REFERENCES CartItem(id) ON DELETE RESTRICT ON UPDATE CASCADE"`
	ProductOptionID uint `gorm:"column:product_option_id" sql:"type:UNSIGNED BIG INT REFERENCES ProductOption(id) ON DELETE RESTRICT ON UPDATE CASCADE"`
}

func (CartItemOption) TableName() string {
	return "CartItemOption"
}
