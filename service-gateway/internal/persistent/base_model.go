package persistent

import (
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"PRIMARY_KEY"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"INDEX"`
}
