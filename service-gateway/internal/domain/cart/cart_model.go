package cart

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/persistent"
	dto "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagmodel"
	"time"
	"strconv"
)

type Cart struct {
	persistent.BaseModel

	TotalPrice uint `gorm:"column:total_price; type:UNSIGNED BIG INT; NOT NULL;"`

	// foreign keys
	UserID uint `gorm:"column:user_id" sql:"type:UNSIGNED BIG INT REFERENCES User(id) ON DELETE RESTRICT ON UPDATE CASCADE"`

	// association: external
	CartItem []*CartItem `gorm:"foreignkey:CartID"`
}

func (Cart) TableName() string {
	return "Cart"
}

func (c *Cart) convertToDTO(cartItemCount int) *dto.Cart {

	cartID := int64(c.ID)
	itemCount := strconv.FormatUint(uint64(cartItemCount), 10)
	totalPrice := strconv.FormatUint(uint64(c.TotalPrice), 10)
	updatedAt := c.UpdatedAt.Format(time.RFC3339)

	dto := &dto.Cart{
		CartID: &cartID,
		ItemCount: &itemCount,
		TotalPrice: &totalPrice,
		UpdatedAt: &updatedAt,
	}

	return dto
}