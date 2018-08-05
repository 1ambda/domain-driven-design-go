package product

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/persistent"
)

const ImageTypeBIN = "BIN"
const ImageTypeURL = "URL"

type Image struct {
	persistent.BaseModel

	Name string `gorm:"column:name; type:VARCHAR(255); NOT NULL;"`
	Type string `gorm:"column:type; type:VARCHAR(255); NOT NULL;"`
	Path string `gorm:"column:path; type:TEXT; NOT NULL;"`
}

func (Image) TableName() string {
	return "Image"
}
