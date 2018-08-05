package product

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/persistent"
	dto "github.com/1ambda/domain-driven-design-go/service-gateway/pkg/generated/swagger/swagmodel"
	"strconv"
)

type ProductOption struct {
	persistent.BaseModel

	Name        string `gorm:"column:name; type:VARCHAR(255); NOT NULL;"`
	Price       uint   `gorm:"column:price; type:UNSIGNED BIG INT; NOT NULL;"`
	Description string `gorm:"column:description; type:TEXT; NOT NULL;"`
	OnSale      OnSale `gorm:"column:on_sale; type:VARCHAR(4); NOT NULL;"`

	Product   Product `gorm:"foreignkey:ProductID;"`
	ProductID uint    `gorm:"column:product_id;" sql:"type:UNSIGNED BIG INT REFERENCES Product(id) ON DELETE RESTRICT ON UPDATE CASCADE;"`

	Image   Image `gorm:"foreignkey:ImageID;"`
	ImageID uint  `gorm:"column:image_id;" sql:"type:UNSIGNED BIG INT NULL REFERENCES Image(id) ON DELETE SET NULL ON UPDATE CASCADE"`
}

func (ProductOption) TableName() string {
	return "ProductOption"
}

func (o *ProductOption) convertToDTO() *dto.ProductOption {
	return &dto.ProductOption{
		ID:          strconv.FormatUint(uint64(o.ID), 10),
		Name:        o.Name,
		Price:       strconv.FormatUint(uint64(o.Price), 10),
		Description: o.Description,
		OnSale:      string(o.OnSale),
	}
}
