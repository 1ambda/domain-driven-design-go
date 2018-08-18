package cart

import (
	"github.com/jinzhu/gorm"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/user"
	e "github.com/1ambda/domain-driven-design-go/service-gateway/internal/exception"
)

type Repository interface {
	CreateCartIfNotExist(tx *gorm.DB, user user.User) (*Cart, e.Exception)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{
		db: db,
	}
}

func (r *repositoryImpl) CreateCartIfNotExist(tx *gorm.DB, user user.User) (*Cart, e.Exception) {
	record := &Cart{
		UserID: user.ID,
		TotalPrice: 0,
	}

	err := tx.Where("user_id = ?", user.ID).First(record).Error
	if err != nil && !gorm.IsRecordNotFoundError(err){
		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewInternalServerException(err, "Failed to get user cart")
		}
	}

	// if exists return it
	if err == nil {
		return record, nil
	}

	if err = tx.Create(record).Error; err != nil {
		return nil, e.NewInternalServerException(err, "Failed to create Cart")
	}

	return record, nil
}
