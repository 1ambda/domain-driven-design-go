package order

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/persistent"
)

type OrderDetail struct {
	persistent.BaseModel

	Index    uint `gorm:"column:index; 		type:UNSIGNED BIG INT; 	NOT NULL;"`
	Price    uint `gorm:"column:price; 		type:UNSIGNED BIG INT; 	NOT NULL;"`
	Quantity uint `gorm:"column:quantity; 	type:UNSIGNED BIG INT; 	NOT NULL;"`
	Amount   uint `gorm:"column:amount; 	type:UNSIGNED BIG INT; 	NOT NULL;"`

	ProductID       uint `gorm:"column:product_id;" sql:"type:UNSIGNED BIG INT REFERENCES Product(id) ON DELETE RESTRICT ON UPDATE CASCADE;"`
	ProductOptionID uint `gorm:"column:product_option_id;" sql:"type:UNSIGNED BIG INT REFERENCES ProductOption(id) ON DELETE RESTRICT ON UPDATE CASCADE;"`
}

func (OrderDetail) TableName() string {
	return "OrderDetail"
}
