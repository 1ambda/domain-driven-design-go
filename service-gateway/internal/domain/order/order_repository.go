package order

import (
	e "github.com/1ambda/domain-driven-design-go/service-gateway/internal/exception"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Repository interface {
	AddOrder(order *Order) (*Order, e.Exception)
	FindOrderById(id uint) (*Order, e.Exception)
	AddOrderDetail(order *OrderDetail) (*OrderDetail, e.Exception)
	FindOrderDetailById(id uint) (*OrderDetail, e.Exception)
}

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) AddOrder(record *Order) (*Order, e.Exception) {
	err := r.db.Create(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to create Order")
		return nil, e.NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) FindOrderById(id uint) (*Order, e.Exception) {
	record := &Order{}
	err := r.db.Where("id = ?", id).First(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find Order")

		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewNotFoundException(wrap)
		}

		return nil, e.NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) AddOrderDetail(record *OrderDetail) (*OrderDetail, e.Exception) {
	err := r.db.Create(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to create OrderDetail")
		return nil, e.NewInternalServerException(wrap)
	}

	return record, nil
}

func (r *repositoryImpl) FindOrderDetailById(id uint) (*OrderDetail, e.Exception) {
	record := &OrderDetail{}
	err := r.db.Where("id = ?", id).First(record).Error

	if err != nil {
		wrap := errors.Wrap(err, "Failed to find OrderDetail")

		if gorm.IsRecordNotFoundError(err) {
			return nil, e.NewNotFoundException(wrap)
		}

		return nil, e.NewInternalServerException(wrap)
	}

	return record, nil
}
