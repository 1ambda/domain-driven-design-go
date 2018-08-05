package user

import (
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/persistent"
)

type User struct {
	persistent.BaseModel

	Email    string `gorm:"column:email; type:VARCHAR(50); NOT NULL; UNIQUE; INDEX;"`
	Phone    string `gorm:"column:phone; type:VARCHAR(50);"`
	Name     string `gorm:"column:name; type:VARCHAR(50);"`
	Birthday string `gorm:"column:birthday; type:VARCHAR(20);"`
	Address  string `gorm:"column:address; type:TEXT;"`
}

func (User) TableName() string {
	return "User"
}
