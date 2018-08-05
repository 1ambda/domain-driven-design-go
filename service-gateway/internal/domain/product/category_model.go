package product

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/persistent"
)

type Category struct {
	persistent.BaseModel

	Path        string `gorm:"column:path; type:VARCHAR(255); NOT NULL; UNIQUE;"`
	Name        string `gorm:"column:name; type:VARCHAR(255); NOT NULL; INDEX;"`
	DisplayName string `gorm:"column:display_name; type:VARCHAR(255); NOT NULL;"`
	Description string `gorm:"column:description; type:TEXT; NOT NULL;"`

	ParentCategoryID *uint `gorm:"column:parent_category_id;" sql:"type:UNSIGNED BIG INT NULL REFERENCES Category(id) ON DELETE SET NULL ON UPDATE CASCADE;"`
}

func (Category) TableName() string {
	return "Category"
}
