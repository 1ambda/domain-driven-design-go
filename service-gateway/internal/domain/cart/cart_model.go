package cart

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/persistent"
)

type Cart struct {
	persistent.BaseModel

	TotalPrice uint `gorm:"column:total_price; type:UNSIGNED BIG INT; NOT NULL;"`

	UserID uint `gorm:"column:user_id" sql:"type:UNSIGNED BIG INT REFERENCES User(id) ON DELETE RESTRICT ON UPDATE CASCADE"`
}

func (Cart) TableName() string {
	return "Cart"
}